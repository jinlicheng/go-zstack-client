package zstack

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type ZStackHttpClient struct {
	client    *http.Client
	baseURL   string // The base URL of the API
	apiKey    string // Api key
	secretKey string // Secret key
	timeout   int64  // Max waiting timeout in seconds for async jobs to finish; defaults to 300 seconds
	asyncURL  string
	syncURL   string
	queryURL  string
}

func (zs *ZStackHttpClient) SyncApi(msg ApiMessage, resp interface{}, callback func(), error func(err interface{})) {
	msg.SetApiKey(zs.apiKey)
	msg.SetSecretKey(zs.secretKey)

	jsonStr, err := msg.toApiMessage()
	if err != nil {
		error(err)
		return
	}
	response, err := zs.client.Post(zs.syncURL, "application/json;charset=UTF-8", strings.NewReader(jsonStr))
	if err != nil {
		error(err)
		return
	}
	defer response.Body.Close()
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		error(err)
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		error(err)
		return
	}
	callback()
}

func (zs *ZStackHttpClient) AsyncApi(msg ApiMessage, resp interface{}, callback func(), error func(err interface{})) {
	msg.SetApiKey(zs.apiKey)
	msg.SetSecretKey(zs.secretKey)

	jsonStr, err := msg.toApiMessage()
	if err != nil {
		error(err)
		return
	}
	response, err := zs.client.Post(zs.asyncURL, "application/json;charset=UTF-8", strings.NewReader(jsonStr))
	if err != nil {
		error(err)
		return
	}
	defer response.Body.Close()
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		error(err)
		return
	}
	var jobResp asyncJobResponse
	err = json.Unmarshal(b, &jobResp)
	if err != nil {
		error(err)
		log.Println("error in conver async job call response.", err, string(b))
		return
	}
	go zs.queryJobResult(jobResp.Id, resp, callback, error)
}

func (zs *ZStackHttpClient) queryJobResult(id string, rtn interface{}, callback func(), error func(err interface{})) {
	for {
		response, err := zs.client.Post(zs.queryURL, "text/xml;charset=UTF-8", strings.NewReader(id))
		if err != nil {
			error(err)
			return
		}
		defer response.Body.Close()
		b, err := ioutil.ReadAll(response.Body)
		if err != nil {
			error(err)
			return
		}
		var ajResp asyncJobResponse
		err = json.Unmarshal(b, &ajResp)
		if err != nil {
			log.Println("Error occur when unmarshal asyncjob query result.", string(b), ajResp)
			error(err)
			return
		}
		log.Println(string(b))
		var respStr string
		if ajResp.Status == 1 {
			log.Println("Loop Run Query ZStack Job Result.")
			time.Sleep(time.Second)
		} else if ajResp.Status == 2 {
			bb, err := ajResp.Rsp.MarshalJSON()
			respStr = string(bb)
			if strings.Contains(respStr, "org.zstack.header.message.APIEvent") {
				var apiEvent ZStackAPIEvent
				err = json.Unmarshal(ajResp.Rsp, &apiEvent)
				if err != nil {
					error(err)
				} else {
					error(apiEvent)
					return
				}
			}
			err = json.Unmarshal(bb, &rtn)
			log.Println(rtn)
			if err != nil {
				error(err)
			} else {
				callback()
			}
			break
		}
	}
}

func newClient(apiurl string, apikey string, secret string, verifyssl bool) *ZStackHttpClient {
	zs := &ZStackHttpClient{
		client: &http.Client{
			Transport: &http.Transport{
				Proxy:           http.ProxyFromEnvironment,
				TLSClientConfig: &tls.Config{InsecureSkipVerify: !verifyssl}, // If verifyssl is true, skipping the verify should be false and vice versa
			},
			Timeout: time.Duration(60 * time.Second),
		},
		baseURL:   apiurl,
		apiKey:    apikey,
		secretKey: secret,
		timeout:   300,
		asyncURL:  apiurl + ASYNC_CALL_PATH,
		syncURL:   apiurl + SYNC_CALL_PATH,
		queryURL:  apiurl + QUERY_PATH,
	}
	return zs
}

func NewClient(apiurl string, apikey string, secret string, verifyssl bool) *ZStackHttpClient {
	zs := newClient(apiurl, apikey, secret, verifyssl)
	return zs
}
