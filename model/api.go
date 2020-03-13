package model

type ResultMeta struct {
	Marker      string `json:"marker"`
	IsTruncated bool   `json:"isTruncated"`
	NextMarker  string `json:"nextMarker"`
	MaxKeys     int    `json:"maxKeys"`
}

type ArgsMeta struct {
	Marker  string `json:"marker"`
	MaxKeys int    `json:"maxKeys"`
}
