package api

// ResetAdminPasswordRequest contains the request body for
// PUT /v3/clusters/{clusterId}/users/admins/passwords.
type ResetAdminPasswordRequest struct {
	Request
	ClusterId   string `json:"-"`
	NewPassword string `json:"newPassword"`
}

// ResetAdminPasswordResponse contains the response body for
// PUT /v3/clusters/{clusterId}/users/admins/passwords.
type ResetAdminPasswordResponse struct {
	Success bool `json:"success,omitempty"`
}
