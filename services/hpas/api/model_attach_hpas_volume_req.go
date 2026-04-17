package api

type AttachHpasVolumeReq struct {
	HpasId    string   `json:"hpasId,omitempty"`
	VolumeIds []string `json:"volumeIds,omitempty"`
}
