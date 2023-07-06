package eccr

import (
	"github.com/baidubce/bce-sdk-go/bce"
)

const (
	DEFAULT_ENDPOINT = "ccr." + bce.DEFAULT_REGION + ".baidubce.com"

	URI_PREFIX = bce.URI_PREFIX + "v1"

	REQUEST_INSTANCE_URL = "/instances"

	REQUEST_PRIVATELINK_URL = "/privatelinks"

	REQUEST_PUBLICLINK_URL = "/publiclinks"

	REQUEST_PUBLICLINK_WITHLIST_URL = "/whitelist"

	REQUEST_CREDENTIAL_URL = "/credential"

	REQUEST_REGISTRY_URL = "/registries"

	REQUEST_REPOSITORIES_URL = "/repositories"

	REQUEST_PROJECT_URL = "/projects"

	REQUEST_IMAGEBUILD_URL = "/imagebuilds"
)

// Client ccr enterprise interface.Interface
type Client struct {
	*bce.BceClient
}

func NewClient(ak, sk, endPoint string) (*Client, error) {
	if len(endPoint) == 0 {
		endPoint = DEFAULT_ENDPOINT
	}
	client, err := bce.NewBceClientWithAkSk(ak, sk, endPoint)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

func getInstanceListURI() string {
	return URI_PREFIX + REQUEST_INSTANCE_URL
}

func getInstanceURI(instanceID string) string {
	return URI_PREFIX + REQUEST_INSTANCE_URL + "/" + instanceID
}

func getInstanceCreateURI() string {
	return URI_PREFIX + REQUEST_INSTANCE_URL
}

func getInstanceRenewURI() string {
	return URI_PREFIX + REQUEST_INSTANCE_URL + "/renew"
}

func getInstanceUpgradeURI(instanceID string) string {
	return URI_PREFIX + REQUEST_INSTANCE_URL + "/" + instanceID + "/upgrade"
}

func getPrivateNetworkListResponseURI(instanceID string) string {
	return URI_PREFIX + REQUEST_INSTANCE_URL + "/" + instanceID + REQUEST_PRIVATELINK_URL
}

func getPrivateNetworkResponseURI(instanceID string) string {
	return URI_PREFIX + REQUEST_INSTANCE_URL + "/" + instanceID + REQUEST_PRIVATELINK_URL
}

func getPublicNetworkResponseURI(instanceID string) string {
	return URI_PREFIX + REQUEST_INSTANCE_URL + "/" + instanceID + REQUEST_PUBLICLINK_URL
}

func getPublicNetworkWhitelistURI(instanceID string) string {
	return URI_PREFIX + REQUEST_INSTANCE_URL + "/" + instanceID + REQUEST_PUBLICLINK_URL + REQUEST_PUBLICLINK_WITHLIST_URL
}

func getInstanceCredentialURI(instanceID string) string {
	return URI_PREFIX + REQUEST_INSTANCE_URL + "/" + instanceID + REQUEST_CREDENTIAL_URL
}

func getInstanceRegistryURI(instanceID string) string {
	return URI_PREFIX + REQUEST_INSTANCE_URL + "/" + instanceID + REQUEST_REGISTRY_URL
}

func getInstanceRegistryIDURI(instanceID, registryID string) string {
	return URI_PREFIX + REQUEST_INSTANCE_URL + "/" + instanceID + REQUEST_REGISTRY_URL + "/" + registryID
}

func getCheckHealthRegistryURI(instanceID string) string {
	return URI_PREFIX + REQUEST_INSTANCE_URL + "/" + instanceID + REQUEST_REGISTRY_URL + "/ping"
}

func getImageBuildURI(instanceID, projectName, repositoryName string) string {
	return URI_PREFIX + REQUEST_INSTANCE_URL + "/" + instanceID + REQUEST_PROJECT_URL + "/" + projectName +
		REQUEST_REPOSITORIES_URL + "/" + repositoryName + REQUEST_IMAGEBUILD_URL
}

func getImageBuildInfoURI(instanceID, projectName, repositoryName, imageBuildID string) string {
	return URI_PREFIX + REQUEST_INSTANCE_URL + "/" + instanceID + REQUEST_PROJECT_URL + "/" + projectName +
		REQUEST_REPOSITORIES_URL + "/" + repositoryName + REQUEST_IMAGEBUILD_URL + "/" + imageBuildID
}
