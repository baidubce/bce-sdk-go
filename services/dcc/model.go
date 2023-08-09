package dcc

import bcc "github.com/baidubce/bce-sdk-go/services/bcc/api"

// ListDedicatedHostArgs struct
type ListDedicatedHostArgs struct {
	Marker   string `json:"marker"`
	MaxKeys  int    `json:"maxkeys"`
	ZoneName string `json:"zoneName"`
}

// ListDedicatedHostResult struct
type ListDedicatedHostResult struct {
	Marker         string
	IsTruncated    bool
	NextMarker     string
	MaxKeys        int
	DedicatedHosts []*DedicatedHostModel
}

// DedicatedHostModel -- xx
type DedicatedHostModel struct {
	ID            string              //	专属服务器ID，符合BCE规范，是一个定长字符串，且只允许包含大小写字母、数字、连字号（-）和下划线（_)。
	Name          string              //	专属服务器名称
	Status        DedicatedHostStatus //	专属服务器状态
	FlavorName    string              //	套餐名称
	ResourceUsage ResourceUsage       //	套餐明细
	PaymentTiming string              //	付费方式，包括Postpaid(后付费)，Prepaid(预付费)两种。
	CreateTime    string              //	创建时间
	ExpireTime    string              //	过期时间，只有Prepaid计费资源存在
	Desc          string              //	描述信息
	ZoneName      string              //	可用区名称
	Tags          []TagModel          //	实例当前配置的标签
}

// DedicatedHostStatus string
type DedicatedHostStatus string

const (
	//DedicatedHostStatusStarting DedicatedHostStatus = "Starting"
	DedicatedHostStatusStarting DedicatedHostStatus = "Starting"
	//DedicatedHostStatusRunning  DedicatedHostStatus = "Running"
	DedicatedHostStatusRunning DedicatedHostStatus = "Running"
	//DedicatedHostStatusExpired  DedicatedHostStatus = "Expired"
	DedicatedHostStatusExpired DedicatedHostStatus = "Expired"
	//DedicatedHostStatusError    DedicatedHostStatus = "Error"
	DedicatedHostStatusError DedicatedHostStatus = "Error"
)

// ResourceUsage struct {
type ResourceUsage struct {
	CPUCount               int `json:"cpuCount"`
	FreeCPUCount           int `json:"FreeCpuCount"`
	MemoryCapacityInGB     int
	FreeMemoryCapacityInGB int
	EphemeralDisks         []EphemeralDisk
}

// EphemeralDisk struct { for go-lint
type EphemeralDisk struct {
	StorageType  StorageType `json:"storageType,omitempty"`
	SizeInGB     int         `json:"sizeInGB,omitempty"`
	FreeSizeInGB int         `json:"freesizeInGB,omitempty"`
}

// StorageType string for go-lint
type StorageType string

const (
	//StorageTypeSata StorageType = "sata" for go-lint
	StorageTypeSata StorageType = "sata"
	//StorageTypeSSD  StorageType = "ssd" for go-lint
	StorageTypeSSD StorageType = "ssd"
)

// TagModel struct { for go-lint
type TagModel struct {
	TagKey   string
	TagValue string
}

// GetDedicatedHostDetailResult struct {
type GetDedicatedHostDetailResult struct {
	DedicatedHost DedicatedHostModel
}

// PurchaseReservedArgs struct {
type PurchaseReservedArgs struct {
	ClientToken string  //	是	Query参数	幂等性Token，是一个长度不超过64位的ASCII字符串)。
	Billing     Billing //	是	RequestBody参数	订单信息
}

// Billing struct {
type Billing struct {
	PaymentTiming string      `json:"paymentTiming"` //	付款时间，预支付（Prepaid）和后支付（Postpaid）
	Reservation   Reservation `json:"reservation"`   //	保留信息，支付方式为后支付时不需要设置，预支付时必须设置
}

// Reservation struct {
type Reservation struct {
	Length   int    `json:"reservationLength"`             //	时长，[1,2,3,4,5,6,7,8,9,12,24,36]
	TimeUnit string `json:"reservationTimeUnit,omitempty"` //	时间单位，month，当前仅支持按月
}

// CreateArgs -- xx
type CreateArgs struct {
	Version       string  //	是	URL参数	API版本号
	ClientToken   string  //	是	Query参数	幂等性Token，是一个长度不超过64位的ASCII字符串)。
	VCPU          int     `json:"vcpu"` //	否	RequestBody参数	待创建专属服务器虚拟CPU核数，数量不能超过物理CPU核数的两倍
	FlavorName    string  //	是	RequestBody参数	套餐类型，可选计算型(calculation)C01/C02，可选大数据机型(storage)S01/S02
	PurchaseCount int     //	否	RequestBody参数	批量创建（购买）的虚专属服务器个数，必须为大于0的整数，可选参数，缺省为1
	Name          string  //	否	RequestBody参数	专属服务器名字（可选）。默认都不指定name。如果指定name：批量时name作为名字的前缀。后端将加上后缀，后缀生成方式：name{ -序号}。如果没有指定name，则自动生成，方式：{instance-八位随机串-序号}。注：随机串从0~9a~z生成；序号按照count的数量级，依次递增，如果count为100，则序号从000~100递增，如果为10，则从00~10递增
	Billing       Billing //	是	RequestBody参数	订单、计费相关参数
	ZoneName      string  //	否	RequestBody参数	指定zone信息，默认为空，由系统自动选择
}

// CreateResult struct {
type CreateResult struct {
	DedicatedHostIds []string
}

// BindTagArgs struct {
type BindTagArgs struct {
	ChangeTags []TagModel //	是	Request Body参数	标签数组，每个标签由tagKey和tagValue组成
}

// ModityInstanceArgs struct {
type ModityInstanceArgs struct {
	Name string //	是	Request Body参数	实例名称
}

// CreateInstanceArgs struct {
type CreateInstanceArgs struct {
	ClientToken           string               `json:"clientToken"`
	ImageID               string               `json:"imageId"`
	Billing               Billing              `json:"billing"`
	CPUCount              int                  `json:"cpuCount"`
	MemoryCapacityInGB    int                  `json:"memoryCapacityInGB"`
	InstanceType          string               `json:"instanceType,omitempty"`
	RootDiskSizeInGb      int                  `json:"rootDiskSizeInGb,omitempty"`
	RootDiskStorageType   string               `json:"rootDiskStorageType,omitempty"`
	LocalDiskSizeInGB     int                  `json:"localDiskSizeInGB,omitempty"`
	EphemeralDisks        []EphemeralDisk      `json:"ephemeralDisks,omitempty"`
	CreateCdsList         []bcc.CreateCdsModel `json:"createCdsList,omitempty"`
	NetworkCapacityInMbps int                  `json:"networkCapacityInMbps"`
	InternetChargeType    string               `json:"internetChargeType,omitempty"`
	DedicatedHostID       string               `json:"dedicatedHostId,omitempty"`
	PurchaseCount         int                  `json:"purchaseCount,omitempty"`
	Name                  string               `json:"name,omitempty"`
	AdminPass             string               `json:"adminPass,omitempty"`
	ZoneName              string               `json:"zoneName,omitempty"`
	SubnetID              string               `json:"subnetId,omitempty"`
	SecurityGroupID       string               `json:"securityGroupId,omitempty"`
	GpuCard               string               `json:"gpuCard,omitempty"`
	FpgaCard              string               `json:"fpgaCard,omitempty"`
	CardCount             string               `json:"cardCount,omitempty"`
}

// CreateInstanceResult struct
type CreateInstanceResult struct {
	InstanceIds []string //	虚机实例ID的集合，其中ID符合BCE规范，必须是一个定长字符串，且只允许包含大小写字母、数字、连字号（-）和下划线（_）。
}
