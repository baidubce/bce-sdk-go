package bbc

import (
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
	"strconv"
)

// ListCDSVolume - list all cds volumes with the given parameters
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - queryArgs: the optional arguments to list cds volumes
// RETURNS:
//     - *ListCDSVolumeResult: the result of cds volume list
//     - error: nil if success otherwise the specific error
func ListCDSVolume(cli bce.Client, queryArgs *ListCDSVolumeArgs) (*ListCDSVolumeResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getVolumeUri())
	req.SetMethod(http.GET)

	if queryArgs != nil {
		if len(queryArgs.InstanceId) != 0 {
			req.SetParam("instanceId", queryArgs.InstanceId)
		}
		if len(queryArgs.ZoneName) != 0 {
			req.SetParam("zoneName", queryArgs.ZoneName)
		}
		if len(queryArgs.Marker) != 0 {
			req.SetParam("marker", queryArgs.Marker)
		}
		if queryArgs.MaxKeys != 0 {
			req.SetParam("maxKeys", strconv.Itoa(queryArgs.MaxKeys))
		}
	}

	if queryArgs == nil || queryArgs.MaxKeys == 0 {
		req.SetParam("maxKeys", "1000")
	}

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ListCDSVolumeResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}
