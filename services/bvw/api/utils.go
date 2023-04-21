package api

import "github.com/baidubce/bce-sdk-go/bce"

const (
	BVW_PREFIX     = bce.URI_PREFIX + "v1/"
	MATERIALIBRARY = "materialLibrary"
	MATLIB         = "matlib"
	PRESET         = "/preset/"
	CONFIG         = "/config"
	POLLINGVIDEO   = "videoEdit/pollingVideo/"
	CREATEVIDEO    = "videoEdit/createNewVideo"
)

func getMaterialLibrary() string {
	return BVW_PREFIX + MATERIALIBRARY
}

func getUploadMaterial() string {
	return BVW_PREFIX + MATLIB
}

func getMateriaPresrtURL() string {
	return BVW_PREFIX + MATERIALIBRARY + PRESET
}

func getMatlibConfigURL() string {
	return BVW_PREFIX + MATLIB + CONFIG
}

func getDraftURL() string {
	return BVW_PREFIX + MATLIB
}

func getPollingVideoURL() string {
	return BVW_PREFIX + POLLINGVIDEO
}

func getCreateVideoURL() string {
	return BVW_PREFIX + CREATEVIDEO
}
