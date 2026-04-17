package api

type DetachHpasVolumeReq struct {
	VolumeIds []string `json:"volumeIds,omitempty"`
}
