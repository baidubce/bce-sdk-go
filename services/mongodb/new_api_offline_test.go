package mongodb

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"testing"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

type capturedMongodbRequest struct {
	Method string
	Path   string
	Query  url.Values
	Body   map[string]interface{}
}

func newOfflineMongodbClient(t *testing.T, handler func(*http.Request) string) *Client {
	t.Helper()

	client, err := NewClient("ak", "1234567890abcdef1234567890abcdef", "mongodb.test")
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	client.HTTPClient = &http.Client{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			response := handler(r)
			return &http.Response{
				StatusCode: http.StatusOK,
				Status:     http.StatusText(http.StatusOK),
				Header:     make(http.Header),
				Body:       io.NopCloser(bytes.NewBufferString(response)),
			}, nil
		}),
	}

	return client
}

func captureMongodbRequest(t *testing.T, r *http.Request) capturedMongodbRequest {
	t.Helper()

	captured := capturedMongodbRequest{
		Method: r.Method,
		Path:   r.URL.Path,
		Query:  r.URL.Query(),
	}
	if r.Body != nil {
		defer r.Body.Close()
		if err := json.NewDecoder(r.Body).Decode(&captured.Body); err != nil {
			t.Fatalf("decode request body error = %v", err)
		}
	}

	return captured
}

func TestNewMongoDBManagementAPIsOffline(t *testing.T) {
	tests := []struct {
		name     string
		response string
		call     func(*Client) error
		assert   func(capturedMongodbRequest)
	}{
		{
			name:     "ListUsers",
			response: `{"marker":"","maxKeys":10,"isTruncated":false,"result":[{"name":"testUser","roles":[{"dbName":"testdb","role":"readWrite"}]}]}`,
			call: func(client *Client) error {
				result, err := client.ListUsers("m-test", &ListUsersArgs{Marker: "m1", MaxKeys: 10})
				if err != nil {
					return err
				}
				if len(result.Result) != 1 || result.Result[0].Name != "testUser" {
					t.Fatalf("ListUsers() result = %+v", result)
				}
				return nil
			},
			assert: func(req capturedMongodbRequest) {
				if req.Method != http.MethodGet || req.Path != "/v1/instance/m-test" {
					t.Fatalf("unexpected request line: %s %s", req.Method, req.Path)
				}
				if _, ok := req.Query["listUsers"]; !ok {
					t.Fatalf("missing listUsers query: %v", req.Query)
				}
				if got := req.Query.Get("marker"); got != "m1" {
					t.Fatalf("marker query = %q", got)
				}
				if got := req.Query.Get("maxKeys"); got != "10" {
					t.Fatalf("maxKeys query = %q", got)
				}
			},
		},
		{
			name:     "CreateUser",
			response: `{}`,
			call: func(client *Client) error {
				return client.CreateUser("m-test", &CreateUserArgs{
					Name:        "testUser",
					Password:    "PlainPassword123",
					Description: "test user",
					Roles:       []RoleInfo{{DbName: "testdb", Role: "readWrite"}},
				})
			},
			assert: func(req capturedMongodbRequest) {
				if req.Method != http.MethodPut || req.Path != "/v1/instance/m-test" {
					t.Fatalf("unexpected request line: %s %s", req.Method, req.Path)
				}
				if _, ok := req.Query["createUser"]; !ok {
					t.Fatalf("missing createUser query: %v", req.Query)
				}
				if got := req.Body["name"]; got != "testUser" {
					t.Fatalf("name body = %v", got)
				}
				if got := req.Body["password"]; got == "" || got == "PlainPassword123" {
					t.Fatalf("password was not encrypted: %v", got)
				}
				roles, ok := req.Body["roles"].([]interface{})
				if !ok || len(roles) != 1 {
					t.Fatalf("roles body = %#v", req.Body["roles"])
				}
			},
		},
		{
			name:     "UpdateUserRoles",
			response: `{}`,
			call: func(client *Client) error {
				return client.UpdateUserRoles("m-test", &UpdateUserRolesArgs{
					Name:  "testUser",
					Roles: []RoleInfo{{DbName: "testdb", Role: "read"}},
				})
			},
			assert: func(req capturedMongodbRequest) {
				if req.Method != http.MethodPut || req.Path != "/v1/instance/m-test" {
					t.Fatalf("unexpected request line: %s %s", req.Method, req.Path)
				}
				if _, ok := req.Query["updateRole"]; !ok {
					t.Fatalf("missing updateRole query: %v", req.Query)
				}
				if got := req.Body["name"]; got != "testUser" {
					t.Fatalf("name body = %v", got)
				}
			},
		},
		{
			name:     "DropUser",
			response: `{}`,
			call: func(client *Client) error {
				return client.DropUser("m-test", &DropUserArgs{Name: "testUser"})
			},
			assert: func(req capturedMongodbRequest) {
				if req.Method != http.MethodPut || req.Path != "/v1/instance/m-test" {
					t.Fatalf("unexpected request line: %s %s", req.Method, req.Path)
				}
				if _, ok := req.Query["dropUser"]; !ok {
					t.Fatalf("missing dropUser query: %v", req.Query)
				}
				if got := req.Body["name"]; got != "testUser" {
					t.Fatalf("name body = %v", got)
				}
			},
		},
		{
			name:     "ListDatabases",
			response: `{"marker":"","maxKeys":10,"isTruncated":false,"result":[{"name":"testdb","description":"test database"}]}`,
			call: func(client *Client) error {
				result, err := client.ListDatabases("m-test", &ListDatabasesArgs{MaxKeys: 10})
				if err != nil {
					return err
				}
				if len(result.Result) != 1 || result.Result[0].Name != "testdb" {
					t.Fatalf("ListDatabases() result = %+v", result)
				}
				return nil
			},
			assert: func(req capturedMongodbRequest) {
				if req.Method != http.MethodGet || req.Path != "/v1/instance/m-test" {
					t.Fatalf("unexpected request line: %s %s", req.Method, req.Path)
				}
				if _, ok := req.Query["listDatabases"]; !ok {
					t.Fatalf("missing listDatabases query: %v", req.Query)
				}
				if got := req.Query.Get("maxKeys"); got != "10" {
					t.Fatalf("maxKeys query = %q", got)
				}
			},
		},
		{
			name:     "CreateDatabase",
			response: `{}`,
			call: func(client *Client) error {
				return client.CreateDatabase("m-test", &CreateDatabaseArgs{
					Name:           "testdb",
					CollectionName: "testcol",
					Description:    "test database",
				})
			},
			assert: func(req capturedMongodbRequest) {
				if req.Method != http.MethodPut || req.Path != "/v1/instance/m-test" {
					t.Fatalf("unexpected request line: %s %s", req.Method, req.Path)
				}
				if _, ok := req.Query["createDatabase"]; !ok {
					t.Fatalf("missing createDatabase query: %v", req.Query)
				}
				if got := req.Body["name"]; got != "testdb" {
					t.Fatalf("name body = %v", got)
				}
				if got := req.Body["collectionName"]; got != "testcol" {
					t.Fatalf("collectionName body = %v", got)
				}
			},
		},
		{
			name:     "DropDatabase",
			response: `{}`,
			call: func(client *Client) error {
				return client.DropDatabase("m-test", &DropDatabaseArgs{Name: "testdb"})
			},
			assert: func(req capturedMongodbRequest) {
				if req.Method != http.MethodPut || req.Path != "/v1/instance/m-test" {
					t.Fatalf("unexpected request line: %s %s", req.Method, req.Path)
				}
				if _, ok := req.Query["dropDatabase"]; !ok {
					t.Fatalf("missing dropDatabase query: %v", req.Query)
				}
				if got := req.Body["name"]; got != "testdb" {
					t.Fatalf("name body = %v", got)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var captured capturedMongodbRequest
			client := newOfflineMongodbClient(t, func(r *http.Request) string {
				captured = captureMongodbRequest(t, r)
				return tt.response
			})

			if err := tt.call(client); err != nil {
				t.Fatalf("%s error = %v", tt.name, err)
			}
			tt.assert(captured)
		})
	}
}

func TestNewMongoDBManagementAPIsValidationOffline(t *testing.T) {
	client, err := NewClient("ak", "1234567890abcdef1234567890abcdef", "127.0.0.1")
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	tests := []struct {
		name string
		call func() error
	}{
		{
			name: "CreateUserRequiresArgs",
			call: func() error { return client.CreateUser("m-test", nil) },
		},
		{
			name: "CreateUserRequiresName",
			call: func() error { return client.CreateUser("m-test", &CreateUserArgs{Password: "password"}) },
		},
		{
			name: "CreateUserRequiresPassword",
			call: func() error { return client.CreateUser("m-test", &CreateUserArgs{Name: "testUser"}) },
		},
		{
			name: "UpdateUserRolesRequiresName",
			call: func() error { return client.UpdateUserRoles("m-test", &UpdateUserRolesArgs{}) },
		},
		{
			name: "DropUserRequiresName",
			call: func() error { return client.DropUser("m-test", &DropUserArgs{}) },
		},
		{
			name: "CreateDatabaseRequiresCollectionName",
			call: func() error { return client.CreateDatabase("m-test", &CreateDatabaseArgs{Name: "testdb"}) },
		},
		{
			name: "DropDatabaseRequiresName",
			call: func() error { return client.DropDatabase("m-test", &DropDatabaseArgs{}) },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.call(); err == nil {
				t.Fatalf("%s expected validation error", tt.name)
			}
		})
	}
}

func TestNewMongoDBManagementAPIsDefaultMaxKeysOffline(t *testing.T) {
	tests := []struct {
		name string
		call func(*Client) error
	}{
		{
			name: "ListUsers",
			call: func(client *Client) error {
				_, err := client.ListUsers("m-test", &ListUsersArgs{MaxKeys: 0})
				return err
			},
		},
		{
			name: "ListDatabases",
			call: func(client *Client) error {
				_, err := client.ListDatabases("m-test", &ListDatabasesArgs{MaxKeys: 1001})
				return err
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := newOfflineMongodbClient(t, func(r *http.Request) string {
				if got := r.URL.Query().Get("maxKeys"); got != "1000" {
					t.Fatalf("default maxKeys = %q", got)
				}
				return `{}`
			})

			if err := tt.call(client); err != nil {
				t.Fatalf("%s error = %v", tt.name, err)
			}
		})
	}
}
