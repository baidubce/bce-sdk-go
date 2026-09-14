package api

import (
	"errors"

	nethttp "net/http"

	"github.com/baidubce/bce-sdk-go/bce"
)

// ResetAdminPassword resets the admin (superuser) password of a cluster.
func ResetAdminPassword(cli bce.Client, region string, request *ResetAdminPasswordRequest) (*ResetAdminPasswordResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("reset admin password request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("reset admin password request clusterId should not be empty")
	}
	if request.NewPassword == "" {
		return nil, errors.New("reset admin password request newPassword should not be empty")
	}
	region = setRegion(region, request.Region)

	credentials, err := credentialsOf(cli)
	if err != nil {
		return nil, err
	}
	encryptedPassword, err := aes128EncryptWithFirst16Char(request.NewPassword, credentials.SecretAccessKey)
	if err != nil {
		return nil, err
	}
	body := *request
	body.NewPassword = encryptedPassword
	headers := map[string]string{HEADER_X_BCE_ACCESSKEY: credentials.AccessKeyId}

	uri := clusterURI(request.ClusterId, URI_USERS, URI_ADMINS, URI_PASSWORDS)

	result := &ResetAdminPasswordResponse{}
	if err := createJSONRequestWithHeaders(cli, region, nethttp.MethodPut, uri, nil, headers, &body, result); err != nil {
		return nil, err
	}
	return result, nil
}
