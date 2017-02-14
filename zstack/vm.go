package zstack

import "encoding/json"

type QueryVmInstanceParams struct {
    baseQueryParams
}

func NewQueryVmInstanceParams() *QueryVmInstanceParams {
	rtn := &QueryVmInstanceParams{}
	rtn.p = make(map[string]interface{})
	return rtn
}

func (p *QueryVmInstanceParams) toApiMessage() (string, error) {
	pp := make(map[string]interface{})
	p.fillQueryStruct()
	pp["org.zstack.header.vm.APIQueryVmInstanceMsg"] = p.p
	b, err := json.Marshal(pp)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

type QueryVmNicParmas struct {
    baseQueryParams
}

func NewQueryVmNicParmas() *QueryVmNicParmas {
	rtn := &QueryVmNicParmas{}
	rtn.p = make(map[string]interface{})
	return rtn
}

func (p *QueryVmNicParmas) toApiMessage() (string, error) {
	pp := make(map[string]interface{})
	p.fillQueryStruct()
	pp["org.zstack.header.vm.APIQueryVmNicMsg"] = p.p
	b, err := json.Marshal(pp)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

type VmService struct {
	zs ZStackClient
}

func NewVmService(zs ZStackClient) *VmService {
	return &VmService{zs: zs}
}
