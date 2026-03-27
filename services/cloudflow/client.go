package cloudflow

import (
	"net/http"
	"time"

	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/cloudflow/api"
)

const (
	SERVICE_NAME           = "cloudflow"
	DEFAULT_SERVICE_DOMAIN = SERVICE_NAME + "." + bce.DEFAULT_REGION + ".bcebos.com"
)

// Client of CloudFlow service is a kind of BceClient, so derived from BceClient
type Client struct {
	*bce.BceClient
}

// ClientConfiguration defines the config components structure by user.
type ClientConfiguration struct {
	Ak                    string
	Sk                    string
	Endpoint              string
	RedirectDisabled      bool
	DisableKeepAlives     bool
	NoVerifySSL           bool
	RetryPolicy           bce.RetryPolicy
	DialTimeout           *time.Duration // timeout of building a connection
	KeepAlive             *time.Duration // interval between keep-alive probes for an active connection
	ReadTimeout           *time.Duration // read timeout of net.Conn
	WriteTimeOut          *time.Duration // write timeout of net.Conn
	TLSHandshakeTimeout   *time.Duration // http.Transport.TLSHandshakeTimeout
	IdleConnectionTimeout *time.Duration // http.Transport.IdleConnTimeout
	ResponseHeaderTimeout *time.Duration // http.Transport.ResponseHeaderTimeout
	HTTPClientTimeout     *time.Duration // http.Client.Timeout
	HTTPClient            *http.Client   // customized http client to send request
}

func NewClientConfig(ak, sk, endpoint string) *ClientConfiguration {
	return &ClientConfiguration{
		Ak:                ak,
		Sk:                sk,
		Endpoint:          endpoint,
		RedirectDisabled:  false,
		DisableKeepAlives: false,
		NoVerifySSL:       false,
		RetryPolicy:       bce.DEFAULT_RETRY_POLICY,
	}
}

func (cfg *ClientConfiguration) WithAk(val string) *ClientConfiguration {
	cfg.Ak = val
	return cfg
}

func (cfg *ClientConfiguration) WithSk(val string) *ClientConfiguration {
	cfg.Sk = val
	return cfg
}

func (cfg *ClientConfiguration) WithEndpoint(val string) *ClientConfiguration {
	cfg.Endpoint = val
	return cfg
}

func (cfg *ClientConfiguration) WithRedirectDisabled(val bool) *ClientConfiguration {
	cfg.RedirectDisabled = val
	return cfg
}

func (cfg *ClientConfiguration) WithDisableKeepAlives(val bool) *ClientConfiguration {
	cfg.DisableKeepAlives = val
	return cfg
}

func (cfg *ClientConfiguration) WithNoVerifySSL(val bool) *ClientConfiguration {
	cfg.NoVerifySSL = val
	return cfg
}

func (cfg *ClientConfiguration) WithDialTimeout(val time.Duration) *ClientConfiguration {
	cfg.DialTimeout = &val
	return cfg
}

func (cfg *ClientConfiguration) WithKeepAlive(val time.Duration) *ClientConfiguration {
	cfg.KeepAlive = &val
	return cfg
}

func (cfg *ClientConfiguration) WithReadTimeout(val time.Duration) *ClientConfiguration {
	cfg.ReadTimeout = &val
	return cfg
}

func (cfg *ClientConfiguration) WithWriteTimeout(val time.Duration) *ClientConfiguration {
	cfg.WriteTimeOut = &val
	return cfg
}

func (cfg *ClientConfiguration) WithTLSHandshakeTimeout(val time.Duration) *ClientConfiguration {
	cfg.TLSHandshakeTimeout = &val
	return cfg
}

func (cfg *ClientConfiguration) WithIdleConnectionTimeout(val time.Duration) *ClientConfiguration {
	cfg.IdleConnectionTimeout = &val
	return cfg
}

func (cfg *ClientConfiguration) WithResponseHeaderTimeout(val time.Duration) *ClientConfiguration {
	cfg.ResponseHeaderTimeout = &val
	return cfg
}

func (cfg *ClientConfiguration) WithHttpClientTimeout(val time.Duration) *ClientConfiguration {
	cfg.HTTPClientTimeout = &val
	return cfg
}

func (cfg *ClientConfiguration) WithRetryPolicy(val bce.RetryPolicy) *ClientConfiguration {
	cfg.RetryPolicy = val
	return cfg
}

func (cfg *ClientConfiguration) WithHttpClient(val http.Client) *ClientConfiguration {
	cfg.HTTPClient = &val
	return cfg
}

// NewClient make the CloudFlow service client with default configuration.
// Use `cli.Config.xxx` to access the config or change it to non-default value.
func NewClient(ak, sk, endpoint string) (*Client, error) {
	return NewClientWithConfig(NewClientConfig(ak, sk, endpoint))
}

func NewClientWithConfig(config *ClientConfiguration) (*Client, error) {
	var credentials *auth.BceCredentials
	var err error
	ak, sk, endpoint := config.Ak, config.Sk, config.Endpoint
	if len(ak) == 0 && len(sk) == 0 { // to support public-read-write request
		credentials, err = nil, nil
	} else {
		credentials, err = auth.NewBceCredentials(ak, sk)
		if err != nil {
			return nil, err
		}
	}
	if len(endpoint) == 0 {
		endpoint = DEFAULT_SERVICE_DOMAIN
	}
	if config.RetryPolicy == nil {
		config.RetryPolicy = bce.DEFAULT_RETRY_POLICY
	}
	defaultSignOptions := &auth.SignOptions{
		HeadersToSign: auth.DEFAULT_HEADERS_TO_SIGN,
		ExpireSeconds: auth.DEFAULT_EXPIRE_SECONDS}
	defaultConf := &bce.BceClientConfiguration{
		Endpoint:                  endpoint,
		Region:                    bce.DEFAULT_REGION,
		UserAgent:                 bce.DEFAULT_USER_AGENT,
		Credentials:               credentials,
		SignOption:                defaultSignOptions,
		Retry:                     config.RetryPolicy,
		ConnectionTimeoutInMillis: bce.DEFAULT_CONNECTION_TIMEOUT_IN_MILLIS,
		RedirectDisabled:          config.RedirectDisabled,
		DisableKeepAlives:         config.DisableKeepAlives,
		NoVerifySSL:               config.NoVerifySSL,
		DialTimeout:               config.DialTimeout,
		KeepAlive:                 config.KeepAlive,
		ReadTimeout:               config.ReadTimeout,
		WriteTimeOut:              config.WriteTimeOut,
		TLSHandshakeTimeout:       config.TLSHandshakeTimeout,
		IdleConnectionTimeout:     config.IdleConnectionTimeout,
		ResponseHeaderTimeout:     config.ResponseHeaderTimeout,
		HTTPClientTimeout:         config.HTTPClientTimeout,
		HTTPClient:                config.HTTPClient,
	}
	v1Signer := &auth.BceV1Signer{}
	client := &Client{bce.NewBceClientWithTimeout(defaultConf, v1Signer)}
	return client, nil
}

func (c *Client) PostMigration(args *api.PostMigrationArgs) (*api.PostMigrationResult, error) {
	return api.PostMigration(c, args)
}

func (c *Client) PostMigrationFromList(args *api.PostMigrationFromListArgs) (*api.PostMigrationResult, error) {
	return api.PostMigrationFromList(c, args)
}

func (c *Client) GetMigration(taskId string) (*api.GetMigrationInfo, error) {
	return api.GetMigration(c, taskId)
}

func (c *Client) ListMigration() (*api.ListMigrationInfo, error) {
	return api.ListMigration(c)
}

func (c *Client) GetMigrationResult(taskId string) (*api.MigrationResult, error) {
	return api.GetMigrationResult(c, taskId)
}

func (c *Client) PauseMigration(taskId string) (*api.MigrationResultCommon, error) {
	return api.PauseMigration(c, taskId)
}

func (c *Client) ResumeMigration(taskId string) (*api.MigrationResultCommon, error) {
	return api.ResumeMigration(c, taskId)
}

func (c *Client) RetryMigration(taskId string) (*api.MigrationResultCommon, error) {
	return api.RetryMigration(c, taskId)
}

func (c *Client) DeleteMigration(taskId string) (*api.MigrationResultCommon, error) {
	return api.DeleteMigration(c, taskId)
}
