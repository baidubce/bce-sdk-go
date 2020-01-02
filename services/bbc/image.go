package bbc

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// CreateImage -- xx
func (c *Client) CreateImage(args *CreateImageArgs) (ret *CreateImageResult, err error) {
	err = bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getURLforImage(args.Version)).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithResult(&ret).
		Do()
	return
}

// ListImage -- xx
func (c *Client) ListImage(args *ListImagesArgs) (list *ListImagesResult, err error) {
	err = bce.NewRequestBuilder(c).
		WithURL(getURLforImage(1)).
		WithMethod(http.GET).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", fmt.Sprintf("%d", args.MaxKeys)).
		WithQueryParamFilter("imageType", args.ImageType).
		WithResult(&list).
		Do()
	return
}

// GetImageDetail -- xx
func (c *Client) GetImageDetail(imageID string) (ret *GetImageDetailResult, err error) {
	err = bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getURLforImagewithID(1, imageID)).
		WithResult(&ret).
		Do()
	return
}

//DeleteImage -- xx
func (c *Client) DeleteImage(imageID string) (err error) {
	err = bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getURLforImagewithID(1, imageID)).
		Do()
	return

}
