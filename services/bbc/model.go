package bbc

// for instance.go

// CreateInstanceArgs -- xx
type CreateInstanceArgs struct {
	Version          int
	ClientToken      string
	FlavorID         string //物理机套餐Id。
	ImageID          string //镜像Id。
	RaidID           string //raid配置Id，可通过查询RAID接口获得。
	RootDiskSizeInGb int    //待创建物理机的系统盘大小。
	PurchaseCount    int    //批量创建（购买）的实例个数
	ZoneName         string //可通过调用查询可用区列表接口查询可用区列表。

	SubnetID  string  //指定subnet信息，为空时将使用默认子网。
	Billing   Billing //订单、计费相关参数。
	Name      string  //物理机名字（可选）。默认都不指定name
	AdminPass string  //机器密码，密码需要加密传输。
}

// Billing -- xx
type Billing struct {
	PaymentTiming string      //付款时间，预支付（Prepaid）和后支付（Postpaid）
	Reservation   Reservation //保留信息，支付方式为后支付时不需要设置，预支付时必须设置
}

// Reservation -- xx
type Reservation struct {
	ReservationLength   int    //时长，[1,2,3,4,5,6,7,8,9,12,24,36]
	ReservationTimeUnit string //时间单位，month，当前仅支持按月
}

// CreateInstanceResult -- xx
type CreateInstanceResult struct {
	InstanceIDs []string
}

// ListInstancesArgs -- xx
type ListInstancesArgs struct {
	//Version    int    //	是	URL参数	API版本号
	Marker     string //	否	Query参数	批量获取列表的查询的起始位置，是一个由系统生成的字符串。
	MaxKeys    int    //	否	Query参数	每页包含的最大数量，最大数量通常不超过1000，缺省值为1000。
	InternalIP string //	否	Query参数	内网ip
}

// ListInstancesResult -- xx
type ListInstancesResult struct {
	Marker      string           //	标记查询的起始位置。
	IsTruncated bool             //	true表示后面还有数据，false表示已经是最后一页。
	NextMarker  string           //	获取下一页所需要传递的marker值。当isTruncated为false时，该域不出现。
	MaxKeys     int              //	 每页包含的最大数量。
	Instances   []*InstanceModel //	实例信息，由 InstanceModel 组成的集合。
}

// InstanceModel -- xx
type InstanceModel struct {
	ID                    string      //	实例ID，符合BCE规范，是一个定长字符串，且只允许包含大小写字母、数字、连字号（-）和下划线（_）
	Name                  string      //	实例名称,支持大小写字母、数字、中文以及-_ /.特殊字符，必须以字母开头，长度1-65
	Status                string      //	实例状态
	Desc                  string      //	实例描述信息
	PaymentTiming         string      //	付费方式，包括Postpaid(后付费)，Prepaid(预付费)两种。
	CreateTime            string      //	创建时间
	ExpireTime            string      //	过期时间
	InternalIP            string      //	内网IP
	PublicIP              string      //	外网IP
	ImageID               string      //	镜像ID
	FlavorID              string      //	套餐ID
	Zone                  string      //	可用区名称
	Region                string      //	区域名称
	NetworkCapacityInMbps int         //	公网带宽，单位为Mb
	Tags                  []*TagModel //	标签信息，由Tag组成的集合
}

// TagModel -- xx
type TagModel struct {
	TagKey   string //	标签键
	TagValue string //	标签值
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
	ChangeTags []*TagModel
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
	Version   string //	是	URL参数	API版本号
	Marker    string //	否	Query参数	批量获取列表的查询的起始位置，是一个由系统生成的字符串
	MaxKeys   int    //	否	Query参数	每页包含的最大数量，最大数量通常不超过1000。缺省值为1000
	ImageType string //	否	Query参数	指定要查询何种类型的镜像，包括All(所有)，System(系统镜像/公共镜像)，Custom(自定义镜像)，Integration(服务集成镜像)，缺省值为All
}

// ListImagesResult -- xx
type ListImagesResult struct {
	Marker      string       //	标记查询的起始位置
	IsTruncated bool         //	true表示后面还有数据，false表示已经是最后一页。
	NextMarker  string       //	获取下一页所需要传递的marker值。当isTruncated为false时，该域不出现。
	MaxKeys     int          //	每页包含的最大数量
	Images      []ImageModel //	返回的镜像列表
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
	Version   string //	是	URL参数	API版本号
	Marker    string //	否	Query参数	批量获取列表的查询的起始位置，是一个由系统生成的字符串
	MaxKeys   int    //	否	Query参数	每页包含的最大数量，最大数量通常不超过1000。缺省值为100
	StartTime string //	否	Query参数	需查询物理机操作的起始时间（UTC时间），格式 yyyy-MM-dd'T'HH:mm:ss'Z' ，为空则查询当日操作日志
	EndTime   string //	否	Query参数	需查询物理机操作的终止时间（UTC时间），格式 yyyy-MM-dd'T'HH:mm:ss'Z' ，为空则查询当日操作日志
}

// ListOperationLogResult -- xx
type ListOperationLogResult struct {
	Marker        string              //	标记查询的起始位置
	IsTruncated   bool                //	true表示后面还有数据，false表示已经是最后一页
	NextMarker    string              //	获取下一页所需要传递的marker值。当isTruncated为false时，该域不出现
	MaxKeys       int                 //	每页包含的最大数量
	OperationLogs []OperationLogModel //	操作日志信息，由 OperationLogModel 组成的集合
}

// OperationLogModel -- xx
type OperationLogModel struct {
	OperationStatus bool   //	操作状态，包括true(成功)，false(失败)两种
	OperationTime   string //	操作日期
	OperationDesc   string //	操作描述
	OperationIP     string //	操作来源Ip
}
