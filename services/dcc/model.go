package dcc

import(
	//"github.com/baidubce/bce-sdk-go/services/bbc"
)


// DedicatedHostArgs struct
type DedicatedHostArgs struct{
	Marker     string `json:"marker"`
	MaxKeys    int `json:"maxkeys"`
	ZoneName         string  `json:"zoneName"`
}

// DedicatedHostResult struct 
type DedicatedHostResult struct {
	Marker         string
	IsTruncated    bool
	NextMarker     string
	MaxKeys        int
	DedicatedHosts []DedicatedHost
}

//DedicatedHost struct {
type DedicatedHost struct {
	ID            string
	Name          string
	Status        DedicatedHostStatus
	FlavorName    string
	ResourceUsage ResourceUsage
	PaymentTiming string
	CreateTime    string
	Expiretime    string
	Desc          string
	ZoneName      string
	Tags          []TagModel
}

//DedicatedHostStatus string
type DedicatedHostStatus string

const (
	//DedicatedHostStatusStarting DedicatedHostStatus = "Starting" 
	DedicatedHostStatusStarting DedicatedHostStatus = "Starting" 
	//DedicatedHostStatusRunning  DedicatedHostStatus = "Running"
	DedicatedHostStatusRunning  DedicatedHostStatus = "Running"
	//DedicatedHostStatusExpired  DedicatedHostStatus = "Expired"
	DedicatedHostStatusExpired  DedicatedHostStatus = "Expired"
	//DedicatedHostStatusError    DedicatedHostStatus = "Error"
	DedicatedHostStatusError    DedicatedHostStatus = "Error"
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
	StorageTypeSSD  StorageType = "ssd"
)

//TagModel struct { for go-lint
type TagModel struct {
	TagKey   string
	TagValue string
}
