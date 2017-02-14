package zstack

import (
	"encoding/json"
)

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

type QueryVmInstanceResponse struct {
	Reply struct {
		Id          string       `json:"id"`
		Success     bool         `json:"success"`
		ServiceId   string       `json:"serviceId"`
		CreatedTime uint         `json:"createdTime"`
		Total       uint16       `json:"total"`
		Vms         []VmInstance `json:"inventories"`
	} `json:"org.zstack.header.vm.APIQueryVmInstanceReply"`
}

type VmInstance struct {
	Uuid                 string   `json:"uuid"`
	Name                 string   `json:"name"`
	Description          string   `json:"description"`
	Zoneuuid             string   `json:"zoneUuid"`
	Clusteruuid          string   `json:"clusterUuid"`
	Imageuuid            string   `json:"imageUuid"`
	Hostuuid             string   `json:"hostUuid"`
	Lasthostuuid         string   `json:"lastHostUuid"`
	Instanceofferinguuid string   `json:"instanceOfferingUuid"`
	Rootvolumeuuid       string   `json:"rootVolumeUuid"`
	Platform             string   `json:"platform"`
	Defaultl3networkuuid string   `json:"defaultL3NetworkUuid"`
	Type                 string   `json:"type"`
	Hypervisortype       string   `json:"hypervisorType"`
	Memorysize           uint64   `json:"memorySize"`
	Cpunum               uint16   `json:"cpuNum"`
	Cpuspeed             uint64   `json:"cpuSpeed"`
	Allocatorstrategy    string   `json:"allocatorStrategy"`
	Createdate           string   `json:"createDate"`
	Lastopdate           string   `json:"lastOpDate"`
	State                string   `json:"state"`
	Vmnics               []VmNic  `json:"vmNics"`
	Allvolumes           []Volume `json:"allVolumes"`
	Hostinfo             Host     `json:"hostInfontor"`
}

type VmNic struct {
	Uuid           string `json:"uuid"`
	VmInstanceUuid string `json:"vmInstanceUuid"`
	L3NetworkUuid  string `json:"l3NetworkUuid"`
	Ip             string `json:"ip"`
	Mac            string `json:"mac"`
	Netmask        string `json:"netmask"`
	Gateway        string `json:"gateway"`
	MetaData       string `json:"metaData"`
	DeviceId       uint64 `json:"deviceId"`
	CreateDate     string `json:"createDate"`
	LastOpDate     string `json:"lastOpDate"`
}

type Volume struct {
	Uuid               string `json:"uuid"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	PrimaryStorageUuid string `json:"primaryStorageUuid"`
	VmInstanceUuid     string `json:"vmInstanceUuid"`
	DiskOfferingUuid   string `json:"diskOfferingUuid"`
	RootImageUuid      string `json:"rootImageUuid"`
	InstallPath        string `json:"installPath"`
	Type               string `json:"type"`
	Format             string `json:"format"`
	Size               uint64 `json:"size"`
	DeviceId           uint64 `json:"deviceId"`
	State              string `json:"state"`
	Status             string `json:"status"`
	CreateDate         string `json:"createDate"`
	LastOpDate         string `json:"lastOpDate"`
}

type Host struct {
	ZoneUuid                string `json:"zoneUuid"`
	Name                    string `json:"name"`
	Uuid                    string `json:"uuid"`
	ClusterUuid             string `json:"clusterUuid"`
	Description             string `json:"description"`
	ManagementIp            string `json:"managementIp"`
	HypervisorType          string `json:"hypervisorType"`
	SshPort                 uint16 `json:"sshPort"`
	State                   string `json:"state"`
	Status                  string `json:"status"`
	TotalCpuCapacity        uint64 `json:"totalCpuCapacity"`
	AvailableCpuCapacity    uint64 `json:"availableCpuCapacity"`
	TotalMemoryCapacity     uint64 `json:"totalMemoryCapacity"`
	AvailableMemoryCapacity uint64 `json:"availableMemoryCapacity"`
	CreateDate              string `json:"createDate"`
	LastOpDate              string `json:"lastOpDate"`
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

type QueryVmNicResponse struct {
	Reply struct {
		Id          string  `json:"id"`
		Success     bool    `json:"success"`
		ServiceId   string  `json:"serviceId"`
		CreatedTime uint    `json:"createdTime"`
		Total       uint16  `json:"total"`
		VmNics      []VmNic `json:"inventories"`
	} `json:"org.zstack.header.vm.APIQueryVmNicReply"`
}

type VmService struct {
	zs ZStackClient
}

func NewVmService(zs ZStackClient) *VmService {
	return &VmService{zs: zs}
}

func (s *VmService) QueryVmInstance(param *QueryVmInstanceParams, callback func(data QueryVmInstanceResponse), error func(err interface{})) {
	var resp QueryVmInstanceResponse
	s.zs.SyncApi(param, &resp, func() {
		callback(resp)
	}, error)
}

func (s *VmService) QueryVmNic(param *QueryVmNicParmas, callback func(data QueryVmNicResponse), error func(err interface{})) {
	var resp QueryVmNicResponse
	s.zs.SyncApi(param, &resp, func() {
		callback(resp)
	}, error)
}
