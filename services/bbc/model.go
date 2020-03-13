package bbc

import "github.com/baidubce/bce-sdk-go/model"

// for instance.go

// CreateInstanceArgs -- xx
type CreateInstanceArgs struct {
	Version          int     `json:"version"`
	ClientToken      string  `json:"clientToken"`
	FlavorID         string  `json:"flavorId"`
	ImageID          string  `json:"imageId"`
	RaidID           string  `json:"raidId"`
	RootDiskSizeInGb int     `json:"rootDiskSizeInGb"`
	PurchaseCount    int     `json:"purchaseCount"`
	ZoneName         string  `json:"zoneName"`
	SubnetID         string  `json:"subnetId"`
	Billing          Billing `json:"billing"`
	Name             string  `json:"name"`
	AdminPass        string  `json:"adminPass"`
}

// Billing -- xx
type Billing struct {
	PaymentTiming string      `json:"paymentTiming"`
	Reservation   Reservation `json:"reservation"`
}

// Reservation -- xx
type Reservation struct {
	Length   int    `json:"reservationLength"`
	TimeUnit string `json:"reservationTimeUnit"`
}

// CreateInstanceResult -- xx
type CreateInstanceResult struct {
	InstanceIDs []string
}

// ListInstancesArgs -- xx
type ListInstancesArgs struct {
	model.ArgsMeta
	InternalIP string //	否	Query参数	内网ip
}

// ListInstancesResult -- xx
type ListInstancesResult struct {
	model.ResultMeta
	Instances []*InstanceModel //	实例信息，由 InstanceModel 组成的集合。
}

// InstanceModel -- xx
type InstanceModel struct {
	ID                    string
	Name                  string
	Status                string
	Desc                  string
	PaymentTiming         string
	CreateTime            string
	ExpireTime            string
	InternalIP            string
	PublicIP              string
	ImageID               string
	FlavorID              string
	Zone                  string
	Region                string
	NetworkCapacityInMbps int
	Tags                  []*model.Tag
}

// StopInstanceArgs -- xx
type StopInstanceArgs struct {
	ForceStop bool
}

// RenameInstanceArgs -- xx
type RenameInstanceArgs struct {
	Name string
}

// RebuildInstanceArgs -- xx
type RebuildInstanceArgs struct {
	Version        int    //	是	URL参数	API版本号
	InstanceID     string //	是	URL参数	指定的实例ID
	Action         string //	是	Query参数	对实例执行的动作，当前取值rebuild
	ImageID        string //	是	Request Body参数	待指定的镜像ID
	AdminPass      string //	是	Request Body参数	机器密码，密码需要加密传输
	IsPreserveData bool   //	否	Request Body参数	是否保留数据，默认为true。当值为true时，raidId和sysRootSize字段不生效
	RaidID         string //	否	Request Body参数	raid配置Id，可通过查询RAID接口获得。此参数在isPreserveData为false时为必填，在isPreserveData为true时不生效
	SysRootSize    int    //	否	Request Body参数	系统盘根分区大小，默认为20G，取值范围为20-100。此参数在isPreserveData为true时不生效
}

// NetworkModel -- xx
type NetworkModel struct {
	BbcID  string      //	bbc实例uuid
	Vpc    VpcModel    //	Vpc信息
	Subnet SubnetModel //	subnet信息
}

// VpcModel -- xx
type VpcModel struct {
	VpcID       string //	网络id
	Cidr        string //	cidr
	Name        string //	网络名称
	IsDefault   bool   //	是否默认网络
	Description string //	描述信息
}

// SubnetModel -- xx
type SubnetModel struct {
	VpcID      string //	网络id
	Name       string //	子网名称
	SubnetType string //	子网类型
	SubnetID   string //	子网id
	Cidr       string //	cidr
	ZoneName   string //	可用区
}

//for tag

// ChangeTagsArgs -- xx
type ChangeTagsArgs struct {
	ChangeTags []*model.Tag
}

// for flavors

// ListFlavorsResult -- xx
type ListFlavorsResult struct {
	Flavors []FlavorModel
}

// FlavorModel -- xx
type FlavorModel struct {
	FlavorID           string //	套餐ID
	CPUCount           int    //	cpu核数
	CPUType            string //	cpu类型
	MemoryCapacityInGB int    //	内存容量，单位为GB
	Disk               string //	磁盘信息，包括SSD盘和SATA盘
	NetworkCard        string //	网络设备信息
	Others             string //	套餐包含的其他信息
}

// GetRaidofFlavorResult -- xx
type GetRaidofFlavorResult struct {
	FlavorID string      //	套餐ID
	Raids    []RaidModel //	RAID信息列表，由 RaidModel 组成的集合
}

// RaidModel -- xx
type RaidModel struct {
	RaidID       string //	RAID的ID
	Raid         string //	RAID名称
	SysSwapSize  int    //	系统盘swap分区默认大小，单位为GIB
	SysRootSize  int    //	系统盘根分区默认大小，单位为GIB
	SysHomeSize  int    //	系统盘/home分区默认大小，单位为GIB
	SysDiskSize  int    //	系统盘总大小，单位为GIB
	DataDiskSize int    //	数据盘总大小，单位为GIB
}

// for image

// CreateImageArgs -- xx
type CreateImageArgs struct {
	Version     int    //	是	URL参数	API版本号
	ClientToken string //	是	Query参数	幂等性Token，是一个长度不超过64位的ASCII字符串
	ImageName   string //	是	Request Body参数	待创建的自定义镜像名称，不能为空，且长度1~65，只能有字母、数字和中划线
	InstanceID  string //	否	Request Body参数	当从实例创建镜像时，此参数是指用于创建镜像的实例ID
}

// CreateImageResult -- xx
type CreateImageResult struct {
	imageID string
}

// ListImagesArgs -- xx
type ListImagesArgs struct {
	model.ArgsMeta
	Version   string //	是	URL参数	API版本号
	ImageType string //	否	Query参数	指定要查询何种类型的镜像，包括All(所有)，System(系统镜像/公共镜像)，Custom(自定义镜像)，Integration(服务集成镜像)，缺省值为All
}

// ListImagesResult -- xx
type ListImagesResult struct {
	model.ResultMeta
	Images []ImageModel //	返回的镜像列表
}

// ImageModel -- xx
type ImageModel struct {
	ID         string      //	镜像ID
	Name       string      //	镜像名称
	Type       ImageType   //	镜像类型
	OsType     string      //	操作系统类型
	OsVersion  string      //	操作系统版本
	OsArch     string      //	操作系统位数
	OsName     string      //	操作系统名称
	OsBuild    string      //	镜像操作系统的构建时间
	CreateTime string      //	镜像的创建时间，符合BCE规范的日期格式
	Status     ImageStatus //	镜像状态
	Desc       string      //	镜像描述信息
}

// GetImageDetailResult -- xx
type GetImageDetailResult struct {
	ID         string      //	镜像ID
	Name       string      //	镜像名称
	Type       ImageType   //	镜像类型
	OsType     string      //	操作系统类型
	OsVersion  string      //	操作系统版本
	OsArch     string      //	操作系统位数
	OsName     string      //	操作系统名称
	OsBuild    string      //	镜像操作系统的构建时间
	CreateTime string      //	镜像的创建时间，符合BCE规范的日期格式
	Status     ImageStatus //	镜像状态
	Desc       string      //	镜像描述信息
}

// ImageType -- xx
type ImageType string

const (
	// ImageTypeSystem -- 系统镜像/公共镜像
	ImageTypeSystem ImageType = "System"
	// ImageTypeCustom -- 自定义镜像
	ImageTypeCustom ImageType = "Custom"
	// ImageTypeIntegration -- 服务集成镜像
	ImageTypeIntegration ImageType = "Integration"
)

// ImageStatus -- xx
type ImageStatus string

const (
	//ImageStatusCreating -- 创建中
	ImageStatusCreating ImageStatus = "Creating"
	//ImageStatusCreatedFailed -- 创建失败
	ImageStatusCreatedFailed ImageStatus = "CreateFailed"
	//ImageStatusAvailable -- 可用
	ImageStatusAvailable ImageStatus = "Available"
	//ImageStatusNotAvailable -- 不可用
	ImageStatusNotAvailable ImageStatus = "NotAvailable"
	//ImageStatusError -- 错误
	ImageStatusError ImageStatus = "Error"
)

// ListOperationLogArgs -- xx
type ListOperationLogArgs struct {
	model.ArgsMeta
	Version   string //	是	URL参数	API版本号
	StartTime string //	否	Query参数	需查询物理机操作的起始时间（UTC时间），格式 yyyy-MM-dd'T'HH:mm:ss'Z' ，为空则查询当日操作日志
	EndTime   string //	否	Query参数	需查询物理机操作的终止时间（UTC时间），格式 yyyy-MM-dd'T'HH:mm:ss'Z' ，为空则查询当日操作日志
}

// ListOperationLogResult -- xx
type ListOperationLogResult struct {
	model.ResultMeta
	OperationLogs []OperationLogModel //	操作日志信息，由 OperationLogModel 组成的集合
}

// OperationLogModel -- xx
type OperationLogModel struct {
	OperationStatus bool   //	操作状态，包括true(成功)，false(失败)两种
	OperationTime   string //	操作日期
	OperationDesc   string //	操作描述
	OperationIP     string //	操作来源Ip
}
