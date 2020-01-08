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

//DedicatedHostModel -- xx
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

//DedicatedHostStatus string
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

//ResourceUsage struct {
type ResourceUsage struct {
	CPUCount               int `json:"cpuCount"`
	FreeCPUCount           int `json:"FreeCpuCount"`
	MemoryCapacityInGB     int
	FreeMemoryCapacityInGB int
	EphemeralDisks         []EphemeralDisk
}

//EphemeralDisk struct { for go-lint
type EphemeralDisk struct {
	StorageType  StorageType
	SizeInGB     int
	FreeSizeInGB int
}

//StorageType string for go-lint
type StorageType string

const (
	//StorageTypeSata StorageType = "sata" for go-lint
	StorageTypeSata StorageType = "sata"
	//StorageTypeSSD  StorageType = "ssd" for go-lint
	StorageTypeSSD StorageType = "ssd"
)

//TagModel struct { for go-lint
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
	PaymentTiming string      //	付款时间，预支付（Prepaid）和后支付（Postpaid）
	Reservation   Reservation //	保留信息，支付方式为后支付时不需要设置，预支付时必须设置
}

// Reservation struct {
type Reservation struct {
	ReservationLength   int    //	时长，[1,2,3,4,5,6,7,8,9,12,24,36]
	ReservationTimeUnit string //	时间单位，month，当前仅支持按月
}

//CreateArgs -- xx
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
	Version               string               //	是	URL参数	API版本号
	ClientToken           string               //	是	Query参数	幂等性Token，是一个长度不超过64位的ASCII字符串。
	ImageID               string               `json:"imageId"` //	是	RequestBody参数	待创建虚拟机实例的镜像ID，可通过调用查询镜像列表接口选择获取所需镜像ID。
	Billing               Billing              //	是	RequestBody参数	订单、计费相关参数
	InstanceType          string               //	否	RequestBody参数	待创建虚拟机实例的类型，具体可选类型参见下述InstanceType,为空时使用默认虚机类型。
	CPUCount              int                  `json:"cpuCount"` //	是	RequestBody参数	待创建虚拟机实例的CPU核数，可选配置请参考区域机型以及可选配置。
	MemoryCapacityInGB    int                  //	是	RequestBody参数	待创建虚拟机实例的内存容量，单位GB，可选配置请参考区域机型以及可选配置。
	RootDiskSizeInGb      int                  //	否	RequestBody参数	待创建虚拟机实例的系统盘大小，单位GB，默认是40GB，范围为[40, 100]GB，超过40GB按照云磁盘价格收费。注意指定的系统盘大小需要满足所使用镜像最小磁盘空间限制。
	RootDiskStorageType   string               //	否	RequestBody参数	待创建虚拟机实例系统盘介质，默认使用SSD型云磁盘，可指定系统盘磁盘类型可参见StorageType。
	LocalDiskSizeInGB     int                  //	否	RequestBody参数	[已废弃]待创建虚拟机实例的临时数据盘大小（不含系统盘，系统盘为免费赠送），单位为GB，大小为0~500G，请采用ephemeralDisks字段。
	EphemeralDisks        []EphemeralDisk      //	否	RequestBody参数	DCC实例可以创建多块本地盘，需要指定磁盘类型以及大小。其他类型BCC最多只能使用一块本地盘，使用默认磁盘类型，需要指定磁盘大小。FPGA实例以及GPU实例默认使用一块本地磁盘，根据配置指定本地盘大小，具体请参考GPU型BCC可选规格配置 以及FPGA型BCC可选规格配置 。
	CreateCdsList         []bcc.CreateCdsModel //	否	RequestBody参数	待创建的CDS磁盘列表，具体数据格式参见下述CreateCdsModel
	NetworkCapacityInMbps int                  //	否	RequestBody参数	公网带宽，单位为Mbps。必须为0~200之间的整数，为0表示不分配公网IP，默认为0Mbps
	InternetChargeType    string               //	否	RequestBody参数	公网带宽计费方式，可选参数详见internetChargeType，若不指定internetChargeType，默认付费方式同BCC，预付费默认为包年包月按带宽，后付费默认为按使用带宽计费。
	DedicatedHostID       string               `json:"dedicatedHostId"` //	否	RequestBody参数	专属服务器id，指定虚机置放位置时指定该值。
	PurchaseCount         int                  //	否	RequestBody参数	批量创建（购买）的虚拟机实例个数，必须为大于0的整数，可选参数，缺省为1
	Name                  string               //	否	RequestBody参数	虚拟机名字（可选）。默认都不指定name。如果指定name：批量时name作为名字的前缀。后端将加上后缀，后缀生成方式：name{ -序号}。如果没有指定name，则自动生成，方式：{instance-八位随机串-序号}。注：随机串从0~9a~z生成；序号按照count的数量级，依次递增，如果count为100，则序号从000~100递增，如果为10，则从00~10递增。支持大小写字母、数字、中文以及-_ /.特殊字符，必须以字母开头，长度1-65。
	AdminPass             string               //	否	RequestBody参数	待指定的实例管理员密码，8-16位字符，英文，数字和符号必须同时存在，符号仅限!@#$%*()，密码需要加密传输，详见链接
	ZoneName              string               //	否	RequestBody参数	zoneName命名规范是“国家-region-可用区序列"，小写，例如北京可用区A为"cn-bj-a"。专属实例使用专属服务器所在zone,无需指定该字段。
	SubnetID              string               `json:"subnetId"`        //	否	RequestBody参数	指定subnet信息，为空时将使用默认子网
	SecurityGroupID       string               `json:"SecurityGroupId"` //	否	RequestBody参数	指定securityGroup信息，为空时将使用默认安全组
	GpuCard               string               //	否	RequestBody参数	待创建实例所要携带的GPU卡信息，具体可选信息参照GpuType，非GPU型实例无需指定此字段
	FpgaCard              string               //	否	RequestBody参数	待创建实例所要携带的FPGA卡信息，具体可选信息参照FpgaType，非FPGA型实例无需指定此字段
	CardCount             string               //	否	RequestBody参数	待创建实例所要携带的GPU卡FPGA卡数量，仅在gpuCard或fpgaCard字段不为空时有效,且需要满足GPU型BCC可选规格配置 以及FPGA型BCC可选规格配置
}

// CreateInstanceResult struct
type CreateInstanceResult struct {
	InstanceIds []string //	虚机实例ID的集合，其中ID符合BCE规范，必须是一个定长字符串，且只允许包含大小写字母、数字、连字号（-）和下划线（_）。
}
