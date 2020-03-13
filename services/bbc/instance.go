package bbc

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// CreateInstance - create an instance with the specific parameters
//
// PARAMS:
//     - args: the arguments to create instance
// RETURNS:
//     - *CreateInstanceResult: the result of create Instance, contains new Instance ID
//     - error: nil if success otherwise the specific error
func (c *Client) CreateInstance(args *CreateInstanceArgs) (ret *CreateInstanceResult, err error) {
	if args == nil {
		err = fmt.Errorf("args cannot be nil")
		return
	}
	if args.Version == 0 {
		args.Version = Version1
	}
	err = bce.NewRequestBuilder(c).
		WithURL(getURL(args.Version)).
		WithMethod(http.POST).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithResult(&ret).
		Do()
	return
}

// ListInstances - xx
func (c *Client) ListInstances(args *ListInstancesArgs) (list *ListInstancesResult, err error) {
	if args == nil {
		args = &ListInstancesArgs{}
	}
	err = bce.NewRequestBuilder(c).
		WithURL(getURL(Version1)).
		WithMethod(http.GET).
		WithQueryParamFilter("marker", args.Marker).
		WithResult(&list).
		Do()
	return
}

// GetInstanceDetail - xx
func (c *Client) GetInstanceDetail(instanceid string) (ret InstanceModel, err error) {
	err = bce.NewRequestBuilder(c).
		WithURL(getURLwithID(Version1, instanceid)).
		WithMethod(http.GET).
		WithResult(&ret).
		Do()
	return
}

// ActionInstance -- xx
func (c *Client) ActionInstance(instanceID string, action string, args *StopInstanceArgs) (err error) {
	cli := bce.NewRequestBuilder(c).
		WithURL(getURLwithID(Version1, instanceID)).
		WithMethod(http.PUT).
		WithQueryParam(action, "")
	if args != nil {
		cli.WithBody(args)
	}
	cli.Do()
	return
}

// StartInstance -- xx
func (c *Client) StartInstance(instanceID string) error {
	return c.ActionInstance(instanceID, "start", nil)
}

// StopInstance -- xx
func (c *Client) StopInstance(instanceID string, force bool) error {
	return c.ActionInstance(instanceID, "stop", &StopInstanceArgs{force})
}

// RebootInstance -- xx
func (c *Client) RebootInstance(instanceID string, force bool) error {
	return c.ActionInstance(instanceID, "reboot", &StopInstanceArgs{force})
}

// RenameInstance -- xx
func (c *Client) RenameInstance(instanceID, name string) error {
	return bce.NewRequestBuilder(c).WithMethod(http.PUT).
		WithURL(getURLwithID(Version1, instanceID)).
		WithQueryParam("rename", "").
		WithBody(&RenameInstanceArgs{name}).
		Do()
}

// UpdateDescInstance -- xx
func (c *Client) UpdateDescInstance(instanceID, desc string) error {
	return bce.NewRequestBuilder(c).WithMethod(http.PUT).
		WithURL(getURLwithID(Version1, instanceID)).
		WithQueryParam("updateDesc", "").
		WithBody(map[string]string{"desc": desc}).
		Do()
}

// RebuildInstance -- xx
func (c *Client) RebuildInstance(instanceID string, args *RebuildInstanceArgs) error {
	if args.Version == 0 {
		args.Version = Version1
	}
	return bce.NewRequestBuilder(c).WithMethod(http.PUT).
		WithURL(getURLwithID(args.Version, instanceID)).
		WithQueryParam("rebuild", "").
		WithBody(args).Do()
}

// OfflineInstance -- xx
func (c *Client) OfflineInstance(instanceID string) (err error) {
	return bce.NewRequestBuilder(c).
		WithURL(getURLwithID(Version2, instanceID)).
		WithMethod(http.DELETE).
		Do()
}

// ChangePassInstance -- xx
func (c *Client) ChangePassInstance(instanceID, adminpass string) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getURLwithID(Version1, instanceID)).
		WithQueryParam("changePass", "").
		Do()
}

// GetSubnetofInstance -- xx
func (c *Client) GetSubnetofInstance(instanceID string, args []string) (ret []*NetworkModel, err error) {
	err = bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getURLforVPC(Version1)).
		WithResult(&ret).
		WithBody(args).Do()
	return
}
