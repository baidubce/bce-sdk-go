package bbc

import (
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// UnbindTag -- xx
func (c *Client) UnbindTag(instanceID string, args *ChangeTagsArgs) (err error) {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getURLforTagwithID(Version1, instanceID)).
		WithQueryParam("unbind", "").
		WithBody(args).
		Do()
}
