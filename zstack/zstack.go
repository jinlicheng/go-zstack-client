package zstack

import (
	"encoding/json"
	"strings"
)

const (
	ASYNC_CALL_PATH = "/api/async"
	SYNC_CALL_PATH  = "/api/sync"
	QUERY_PATH      = "/api/query"
)

type ZStackClient interface {
	SyncApi(msg ApiMessage, resp interface{}, callback func(), error func(err interface{}))
	AsyncApi(msg ApiMessage, resp interface{}, callback func(), error func(err interface{}))
}

type QueryCondition struct {
	Name  string `json:"name"`
	Op    string `json:"op"`
	Value string `json:"value"`
}

type ApiMessage interface {
	toApiMessage() (string, error)
	SetApiKey(apiKey string)
	SetSecretKey(secretKey string)
}

type baseParams struct {
	p map[string]interface{}
}

func (p *baseParams) makeSureNotNil() {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
}

func (p *baseParams) SetApiKey(apiKey string) {
	p.makeSureNotNil()
	p.p["apiKey"] = apiKey
}

func (p *baseParams) SetSecretKey(secretKey string) {
	p.makeSureNotNil()
	p.p["secretKey"] = secretKey
}

type baseDeleteParams struct {
    baseParams
}

func (p *baseDeleteParams) SetUuid(uuid string){
    p.makeSureNotNil()
    p.p["uuid"] = uuid
}

type baseCreateParams struct {
	baseParams
}

func (p *baseCreateParams) setResourceUuid(resourceUuid string) {
	p.makeSureNotNil()
	p.p["resouceUuid"] = resourceUuid
}

func (p *baseCreateParams) setSystemTags(tags []string) {
	p.makeSureNotNil()
	p.p["systemTags"] = strings.Join(tags, ",")
}

func (p *baseCreateParams) setUserTags(tags []string) {
	p.makeSureNotNil()
	p.p["userTags"] = strings.Join(tags, ",")
}

type baseQueryParams struct {
	baseParams
}

func (p *baseQueryParams) fillQueryStruct() {
	if _, found := p.p["conditions"]; !found {
		p.p["conditions"] = make([]QueryCondition, 0)
	}

	if _, found := p.p["count"]; !found {
		p.p["count"] = false
	}
	if _, found := p.p["replyWithCount"]; !found {
		p.p["replyWithCount"] = true
	}
}

func (p *baseQueryParams) SetConditions(conditions []QueryCondition) {
	p.makeSureNotNil()
	p.p["conditions"] = conditions
}

func (p *baseQueryParams) SetLimit(limit uint16) {
	p.makeSureNotNil()
	p.p["limit"] = limit
}

func (p *baseQueryParams) SetStart(start uint16) {
	p.makeSureNotNil()
	p.p["start"] = start
}

func (p *baseQueryParams) SetCount(count bool) {
	p.makeSureNotNil()
	p.p["count"] = count
}

func (p *baseQueryParams) SetReplyWithCount(replyWithCount bool) {
	p.makeSureNotNil()
	p.p["replyWithCount"] = replyWithCount
}

type asyncJobResponse struct {
	Id     string          `json:"id"`
	Status uint8           `json:"status"`
	Rsp    json.RawMessage `json:"rsp"`
}

type ZStackError struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	Details     string `json:"details"`
}

type ZStackAPIReply struct {
	Success     string      `json:"success"`
	ServiceId   string      `json:"serviceId"`
	CreatedTime uint64      `json:"createdTime"`
	Id          string      `json:"id"`
	Error       ZStackError `json:"error"`
}

type ZStackType struct {
    Name string `json:"_name"`
}

type ZStackAPIEvent struct {
	Reply struct {
		Id          string      `json:"id"`
		ApiId       string      `json:"apiId"`
		CreatedTime uint64      `json:"createdTime"`
		Success     bool        `json:"success"`
		Error       ZStackError `json:"error"`
		Type        ZStackType  `json:"type"`
	} `json:"org.zstack.header.message.APIEvent"`
}
