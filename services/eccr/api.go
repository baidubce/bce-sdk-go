package eccr

import (
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

type Interface interface {
	ListInstances(args *ListInstancesArgs) (*ListInstancesResponse, error)
	GetInstanceDetail(instanceID string) (*GetInstanceDetailResponse, error)
	CreateInstance(args *CreateInstanceArgs) (*CreateInstanceResponse, error)
	RenewInstance(orderType string, args *RenewInstanceArgs) (*RenewInstanceResponse, error)
	UpdateInstance(instanceID string, args *UpdateInstanceArgs) (*UpdateInstanceResponse, error)
	UpgradeInstance(instanceID string, args *UpgradeInstanceArgs) (*UpgradeInstanceResponse, error)
	ListPrivateNetworks(instanceID string) (*ListPrivateNetworksResponse, error)
	CreatePrivateNetwork(instanceID string, args *CreatePrivateNetworkArgs) (map[string]string, error)
	DeletePrivateNetwork(instanceID string, args *DeletePrivateNetworkArgs) error
	ListPublicNetworks(instanceID string) (*ListPublicNetworksResponse, error)
	UpdatePublicNetwork(instanceID string, args *UpdatePublicNetworkArgs) error
	DeletePublicNetworkWhitelist(instanceID string, args *DeletePublicNetworkWhitelistArgs) error
	AddPublicNetworkWhitelist(instanceID string, args *AddPublicNetworkWhitelistArgs) error
	ResetPassword(instanceID string, args *ResetPasswordArgs) (map[string]string, error)
	CreateTemporaryToken(instanceID string, args *CreateTemporaryTokenArgs) (*CreateTemporaryTokenResponse, error)
	CreateRegistry(instanceID string, args *CreateRegistryArgs) (*CreateRegistryResponse, error)
	GetRegistryDetail(instanceID, registryID string) (*RegistryResponse, error)
	ListRegistries(instanceID string, args *ListRegistriesArgs) (*ListRegistriesResponse, error)
	CheckHealthRegistry(instanceID string, args *RegistryRequestArgs) error
	UpdateRegistry(instanceID, registryID string, args *RegistryRequestArgs) (*RegistryResponse, error)
	DeleteRegistry(instanceID, registryID string) error
}

// ListInstances - list all instance with the specific parameters
//
// PARAMS:
//   - ListInstancesArgs: the arguments to list all instance
//
// RETURNS:
//   - ListInstancesResponse: the result of list Instance
//   - error: nil if success otherwise the specific error
func (c *Client) ListInstances(args *ListInstancesArgs) (*ListInstancesResponse, error) {

	result := &ListInstancesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getInstanceListURI()).
		WithQueryParamFilter("keywordType", args.KeywordType).
		WithQueryParamFilter("keyword", args.Keyword).
		WithQueryParamFilter("pageNo", strconv.Itoa(args.PageNo)).
		WithQueryParamFilter("pageSize", strconv.Itoa(args.PageSize)).
		WithQueryParamFilter("acrossregion", args.Acrossregion).
		WithResult(result).
		Do()

	return result, err
}

// GetInstanceDetail - get a specific instance detail info
//
// PARAMS:
//   - instanceID: the specific instance ID
//
// RETURNS:
//   - *GetInstanceDetailResponse: the result of get instance detail info
//   - error: nil if success otherwise the specific error
func (c *Client) GetInstanceDetail(instanceID string) (*GetInstanceDetailResponse, error) {

	result := &GetInstanceDetailResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getInstanceURI(instanceID)).
		WithResult(result).
		Do()

	return result, err
}

// CreateInstance - create instance with the specific parameters
//
// PARAMS:
//   - CreateInstanceArgs: the arguments to crate Instance
//
// RETURNS:
//   - CreateInstanceResponse: the result of create Instance
//   - error: nil if success otherwise the specific error
func (c *Client) CreateInstance(args *CreateInstanceArgs) (*CreateInstanceResponse, error) {
	result := &CreateInstanceResponse{}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getInstanceCreateURI()).
		WithResult(result).
		WithBody(args).
		Do()

	return result, err
}

// RenewInstance - create instance with the specific parameters
//
// PARAMS:
//   - orderType: the operation type, value requires renew
//   - ConfirmOrderRequest: the arguments to crate Instance
//
// RETURNS:
//   - CreateInstanceResponse: the result of create Instance
//   - error: nil if success otherwise the specific error
func (c *Client) RenewInstance(orderType string, args *RenewInstanceArgs) (*RenewInstanceResponse, error) {
	result := &RenewInstanceResponse{}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithQueryParamFilter("orderType", orderType).
		WithURL(getInstanceRenewURI()).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// UpdateInstance - update instance info
//
// PARAMS:
//   - instanceId: the specific instance ID
//   - UpdateInstanceArgs: parameters required to update instance info
//
// RETURNS:
//   - *UpdateInstanceResponse: the result of updated instance info
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateInstance(instanceID string, args *UpdateInstanceArgs) (*UpdateInstanceResponse, error) {
	result := &UpdateInstanceResponse{}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getInstanceURI(instanceID)).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// UpgradeInstance - upgrade instance by specific parameters
//
// PARAMS:
//   - instanceId: the specific instance ID
//   - UpgradeInstanceArgs: parameters required to upgrade instance information
//
// RETURNS:
//   = UpgradeInstanceResponse: the result of upgrade instance
//   - error: nil if success otherwise the specific error
func (c *Client) UpgradeInstance(instanceID string, args *UpgradeInstanceArgs) (*UpgradeInstanceResponse, error) {
	result := &UpgradeInstanceResponse{}

	err := bce.NewRequestBuilder(c).WithMethod(http.PUT).
		WithURL(getInstanceUpgradeURI(instanceID)).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// ListPrivateNetworks - list all Privatelinks in an instance with the specific parameters
//
// PARAMS:
//   - instanceID: the specific instance ID
//
// RETURNS:
//   - *ListPrivateNetworksResponse: the result of list Privatelinks
//   - error: nil if success otherwise the specific error
func (c *Client) ListPrivateNetworks(instanceID string) (*ListPrivateNetworksResponse, error) {
	result := &ListPrivateNetworksResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getPrivateNetworkListResponseURI(instanceID)).
		WithResult(result).
		Do()

	return result, err
}

// CreatePrivateNetwork - create private Network with the specific parameters
//
// PARAMS:
//   - instanceID: the specific instance ID
//   - CreateInstanceArgs: the arguments to crate private network
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) CreatePrivateNetwork(instanceID string, args *CreatePrivateNetworkArgs) (map[string]string, error) {
	result := make(map[string]string)

	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getPrivateNetworkResponseURI(instanceID)).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// DeletePrivateNetwork - delete private Network with the specific parameters
//
// PARAMS:
//   - instanceID: the specific instance ID
//   - DeletePrivateNetworkArgs: the arguments to delete private network
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeletePrivateNetwork(instanceID string, args *DeletePrivateNetworkArgs) error {

	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getPrivateNetworkResponseURI(instanceID)).
		WithBody(args).
		Do()

	return err
}

// ListPublicNetworks - list all Publiclinks in an instance with the specific parameters
//
// PARAMS:
//   - instanceID: the specific instance ID
//
// RETURNS:
//   - *ListPublicNetworksResponse: the result of list Publiclinks
//   - error: nil if success otherwise the specific error
func (c *Client) ListPublicNetworks(instanceID string) (*ListPublicNetworksResponse, error) {
	result := &ListPublicNetworksResponse{}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getPublicNetworkResponseURI(instanceID)).
		WithResult(result).Do()

	return result, err
}

// UpdatePublicNetwork - update Publiclink
//
// PARAMS:
//   - instanceId: the specific instance ID
//   - UpdatePublicNetworkArgs: parameters required to update publiclink info
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdatePublicNetwork(instanceID string, args *UpdatePublicNetworkArgs) error {

	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getPublicNetworkResponseURI(instanceID)).
		WithBody(args).
		Do()

	return err
}

// DeletePublicNetworkWhitelist - delete Publiclink white list
//
// PARAMS:
//   - instanceId: the specific instance ID
//   - DeletePublicNetworkWhiteListArgs: delete publiclinks list
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeletePublicNetworkWhitelist(instanceID string, args *DeletePublicNetworkWhitelistArgs) error {

	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getPublicNetworkWhitelistURI(instanceID)).
		WithBody(args).
		Do()
	return err
}

// AddPublicNetworkWhitelist - add Publiclink white list
//
// PARAMS:
//   - instanceId: the specific instance ID
//   - DeletePublicNetworkWhiteListArgs: the arguments to delete publiclinks withlist
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) AddPublicNetworkWhitelist(instanceID string, args *AddPublicNetworkWhitelistArgs) error {

	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getPublicNetworkWhitelistURI(instanceID)).
		WithBody(args).
		Do()

	return err
}

// ResetPassword - reset login password
//
// PARAMS:
//   - instanceId: the specific instance ID
//   - ResetPasswordArgs: the arguments to reset password
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - map[string]string: the result of reset password
func (c *Client) ResetPassword(instanceID string, args *ResetPasswordArgs) (map[string]string, error) {
	result := make(map[string]string)

	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getInstanceCredentialURI(instanceID)).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// CreateTemporaryToken - create temporary token
//
// PARAMS:
//   - instanceID: the specific instance ID
//   - CreateTemporaryTokenArgs: the arguments to crate temporary token
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) CreateTemporaryToken(instanceID string, args *CreateTemporaryTokenArgs) (*CreateTemporaryTokenResponse, error) {
	result := &CreateTemporaryTokenResponse{}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getInstanceCredentialURI(instanceID)).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// CreateRegistry - create registry
//
// PARAMS:
//   - instanceID: the specific instance ID
//   - CreateRegistryArgs: the arguments to crate registry
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - CreateRegistryResponse: the result od create registry
func (c *Client) CreateRegistry(instanceID string, args *CreateRegistryArgs) (*CreateRegistryResponse, error) {
	result := &CreateRegistryResponse{}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getInstanceRegistryURI(instanceID)).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// GetRegistryDetail - get a specific registry detail info
//
// PARAMS:
//   - instanceID: the specific instance ID
//   - registryID: the specific registry ID
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - RegistryResponse: the result of create registry
func (c *Client) GetRegistryDetail(instanceID, registryID string) (*RegistryResponse, error) {
	result := &RegistryResponse{}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getInstanceRegistryIDURI(instanceID, registryID)).
		WithResult(result).
		Do()

	return result, err
}

// ListRegistries - get a registry list of instance
//
// PARAMS:
//   - instanceID: the specific instance ID
//   - ListRegistriesArgs: parameters required to list registry information
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - ListRegistriesResponse: the result of list registry
func (c *Client) ListRegistries(instanceID string, args *ListRegistriesArgs) (*ListRegistriesResponse, error) {
	result := &ListRegistriesResponse{}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithQueryParamFilter("registryName", args.RegistryName).
		WithQueryParamFilter("registryType", args.RegistryType).
		WithQueryParamFilter("pageNo", strconv.Itoa(args.PageNo)).
		WithQueryParamFilter("pageSize", strconv.Itoa(args.PageSize)).
		WithURL(getInstanceRegistryURI(instanceID)).
		WithResult(result).
		Do()

	return result, err
}

// CheckHealthRegistry - check if the registry is healthy
//
// PARAMS:
//   - instanceID: the specific instance ID
//   - RegistryRequestArgs: parameters required to check registry
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) CheckHealthRegistry(instanceID string, args *RegistryRequestArgs) error {

	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCheckHealthRegistryURI(instanceID)).
		WithBody(args).
		Do()

	return err
}

// UpdateRegistry - update the registry info
//
// PARAMS:
//   - instanceID: the specific instance ID
//   - registryID: the specific registry ID
//   - RegistryRequestArgs: parameters required to update registry
//
// RETURNS:
//   - error: nil if success otherwise the specific error
//   - RegistryResponse: the result of update registry
func (c *Client) UpdateRegistry(instanceID, registryID string, args *RegistryRequestArgs) (*RegistryResponse, error) {
	result := &RegistryResponse{}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getInstanceRegistryIDURI(instanceID, registryID)).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) DeleteRegistry(instanceID, registryID string) error {

	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getInstanceRegistryIDURI(instanceID, registryID)).
		Do()

	return err
}
