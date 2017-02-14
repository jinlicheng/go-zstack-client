package zstack

import (
	"encoding/json"
)

type QueryVipParams struct {
	baseQueryParams
}

func NewQueryVipParams() *QueryVipParams {
	rtn := &QueryVipParams{}
	rtn.p = make(map[string]interface{})
	return rtn
}

func (p *QueryVipParams) toApiMessage() (string, error) {
	pp := make(map[string]interface{})
	p.fillQueryStruct()
	pp["org.zstack.network.service.vip.APIQueryVipMsg"] = p.p
	b, err := json.Marshal(pp)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

type QueryVipResponse struct {
	Reply struct {
		Id          string `json:"id"`
		Success     bool   `json:"success"`
		ServiceId   string `json:"serviceId"`
		CreatedTime uint   `json:"createdTime"`
		Total       uint16 `json:"total"`
		Vips        Vip    `json:"inventories"`
	} `json:"org.zstack.network.service.vip.APIQueryVipReply"`
}

type CreateVipParams struct {
	baseCreateParams
}

func (p *CreateVipParams) SetName(name string) {
	p.makeSureNotNil()
	p.p["name"] = name
}

func (p *CreateVipParams) SetDescription(description string) {
	p.makeSureNotNil()
	p.p["description"] = description
}

func (p *CreateVipParams) SetL3NetworkUuid(l3NetworkUuid string) {
	p.makeSureNotNil()
	p.p["l3NetworkUuid"] = l3NetworkUuid
}

func (p *CreateVipParams) SetAllocatorStrategy(allocatorStrategy string) {
	p.makeSureNotNil()
	p.p["allocatorStrategy"] = allocatorStrategy
}

func (p *CreateVipParams) SetRequiredIp(requiredIp string) {
	p.makeSureNotNil()
	p.p["requiredIp"] = requiredIp
}

func (p *CreateVipParams) toApiMessage() (string, error) {
	pp := make(map[string]interface{})
	pp["org.zstack.network.service.vip.APICreateVipMsg"] = p.p
	b, err := json.Marshal(pp)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

type CreateVipResponse struct {
	Reply struct {
		id          string      `json:"id"`
		ApiId       string      `json:"apiId"`
		CreatedTime uint64      `json:"createdTime"`
		Success     bool        `json:"success"`
		Error       ZStackError `json:"error"`
		Type        ZStackType  `json:"type"`
		Vip         Vip         `json:"inventory"`
	} `json:"org.zstack.network.service.vip.APICreateVipEvent"`
}

type Vip struct {
	Uuid              string `json:"uuid"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	L3NetworkUuid     string `json:"l3NetworkUuid"`
	Ip                string `json:"ip"`
	State             string `json:"state"`
	Gateway           string `json:"gateway"`
	Netmask           string `json:"netmask"`
	ServiceProvider   string `json:"serviceProvider"`
	PeerL3NetworkUuid string `json:"peerL3NetworkUuid"`
	UseFor            string `json:"useFor"`
	CreateDate        string `json:"createDate"`
	LastOpDate        string `json:"lastOpDate"`
}

type VipService struct {
	zs ZStackClient
}

func NewVipService(zs ZStackClient) *VipService {
	return &VipService{zs: zs}
}

func (s *VipService) QueryVip(param *QueryVipParams, callback func(data QueryVipResponse), error func(err interface{})) {
	var resp QueryVipResponse
	s.zs.SyncApi(param, &resp, func() {
		callback(resp)
	}, error)
}

func (s *VipService) CreateVip(param *CreateVipParams, callback func(data CreateVipResponse), error func(err interface{})) {
	var resp CreateVipResponse
	s.zs.AsyncApi(param, &resp, func() {
		callback(resp)
	}, error)
}
