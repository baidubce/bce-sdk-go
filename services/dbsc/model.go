package dbsc

type CreateVolumeClusterArgs struct {
	ZoneName        string      `json:"zoneName,omitempty"`
	ClusterName     string      `json:"clusterName,omitempty"`
	StorageType     StorageType `json:"storageType,omitempty"`
	ClusterSizeInGB int         `json:"clusterSizeInGB,omitempty"`
	PurchaseCount   int         `json:"purchaseCount,omitempty"`
	Billing         *Billing    `json:"billing"`
	RenewTimeUnit   string      `json:"renewTimeUnit"`
	RenewTime       int         `json:"renewTime"`
	UuidFlag        string      `json:"uuidFlag"`
}

type Billing struct {
	PaymentTiming PaymentTimingType `json:"paymentTiming,omitempty"`
	Reservation   *Reservation      `json:"reservation,omitempty"`
}

type Reservation struct {
	ReservationLength   int    `json:"reservationLength"`
	ReservationTimeUnit string `json:"reservationTimeUnit"`
}

type PaymentTimingType string

const (
	PaymentTimingPrePaid  PaymentTimingType = "Prepaid"
	PaymentTimingPostPaid PaymentTimingType = "Postpaid"
)

type StorageType string

const (
	StorageTypeHP1 StorageType = "hp1"
	StorageTypeHdd StorageType = "hdd"
)

type CreateVolumeClusterResult struct {
	ClusterIds   []string `json:"clusterIds"`
	ClusterUuids []string `json:"clusterUuids"`
}

type ListVolumeClusterArgs struct {
	Marker      string
	MaxKeys     int
	ClusterName string
	ZoneName    string
}

type ListVolumeClusterResult struct {
	Marker      string               `json:"marker"`
	IsTruncated bool                 `json:"isTruncated"`
	NextMarker  string               `json:"nextMarker"`
	MaxKeys     int                  `json:"maxKeys"`
	Result      []VolumeClusterModel `json:"result"`
}

type VolumeClusterModel struct {
	ClusterId         string `json:"clusterId"`
	ClusterName       string `json:"clusterName"`
	CreatedTime       string `json:"createdTime"`
	ExpiredTime       string `json:"expiredTime"`
	Status            string `json:"status"`
	ZoneName          string `json:"logicalZone"`
	ProductType       string `json:"productType"`
	ClusterType       string `json:"clusterType"`
	TotalCapacity     int    `json:"totalCapacity"`
	UsedCapacity      int    `json:"usedCapacity"`
	AvailableCapacity int    `json:"availableCapacity"`
	ExpandingCapacity int    `json:"expandingCapacity"`
	CreatedVolumeNum  int    `json:"createdVolumeNum"`
	EnableAutoRenew   bool   `json:"enableAutoRenew"`
	RenewTimeUnit     string `json:"renewTimeUnit"`
	RenewTime         int    `json:"renewTime"`
}

type VolumeClusterDetail struct {
	ClusterId           string   `json:"clusterId"`
	ClusterName         string   `json:"clusterName"`
	CreatedTime         string   `json:"createdTime"`
	ExpiredTime         string   `json:"expiredTime"`
	Status              string   `json:"status"`
	ZoneName            string   `json:"logicalZone"`
	ProductType         string   `json:"productType"`
	ClusterType         string   `json:"clusterType"`
	TotalCapacity       int      `json:"totalCapacity"`
	UsedCapacity        int      `json:"usedCapacity"`
	AvailableCapacity   int      `json:"availableCapacity"`
	ExpandingCapacity   int      `json:"expandingCapacity"`
	CreatedVolumeNum    int      `json:"createdVolumeNum"`
	AffiliatedCDSNumber []string `json:"affiliatedCDSNumber"`
	EnableAutoRenew     bool     `json:"enableAutoRenew"`
	RenewTimeUnit       string   `json:"renewTimeUnit"`
	RenewTime           int      `json:"renewTime"`
}

type ResizeVolumeClusterArgs struct {
	NewClusterSizeInGB int `json:"newClusterSizeInGB"`
}

type PurchaseReservedVolumeClusterArgs struct {
	Billing *Billing `json:"billing"`
}

type AutoRenewVolumeClusterArgs struct {
	ClusterId     string `json:"clusterId"`
	RenewTimeUnit string `json:"renewTimeUnit"`
	RenewTime     int    `json:"renewTime"`
}

type CancelAutoRenewVolumeClusterArgs struct {
	ClusterId string `json:"clusterId"`
}
