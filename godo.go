package godo

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"sync"

	"golang.org/x/time/rate"
)

const (
	libraryVersion = "1.192.0"
	defaultBaseURL = "https://api.digitalocean.com/"
	userAgent      = "godo/" + libraryVersion
	mediaType      = "application/json"

	headerRateLimit             = "RateLimit-Limit"
	headerRateRemaining         = "RateLimit-Remaining"
	headerRateReset             = "RateLimit-Reset"
	headerRequestID             = "x-request-id"
	internalHeaderRetryAttempts = "X-Godo-Retry-Attempts"

	defaultRetryMax     = 4
	defaultRetryWaitMax = 30
	defaultRetryWaitMin = 1
)

// Client manages communication with DigitalOcean V2 API.
type Client struct {
	// HTTP client used to communicate with the DO API.
	HTTPClient *http.Client

	// Base URL for API requests.
	BaseURL *url.URL

	// User agent for client
	UserAgent string

	// Rate contains the current rate limit for the client as determined by the most recent
	// API call. It is not thread-safe. Please consider using GetRate() instead.
	Rate    Rate
	ratemtx sync.Mutex

	// Services used for communicating with the API
	Account             AccountService
	Actions             ActionsService
	Apps                AppsService
	Balance             BalanceService
	BillingHistory      BillingHistoryService
	CDNs                CDNService
	Certificates        CertificatesService
	Databases           DatabasesService
	Domains             DomainsService
	Droplets            DropletsService
	DropletActions      DropletActionsService
	DropletAutoscale    DropletAutoscaleService
	VPCNATGateways      VPCNATGatewaysService
	Firewalls           FirewallsService
	FloatingIPs         FloatingIPsService
	FloatingIPActions   FloatingIPActionsService
	Functions           FunctionsService
	Images              ImagesService
	ImageActions        ImageActionsService
	Invoices            InvoicesService
	Keys                KeysService
	Kubernetes          KubernetesService
	LoadBalancers       LoadBalancersService
	Monitoring          MonitoringService
	Security            SecurityService
	Nfs                 NfsService
	NfsActions          NfsActionsService
	OneClick            OneClickService
	Projects            ProjectsService
	Regions             RegionsService
	Registry            RegistryService
	Registries          RegistriesService
	ReservedIPs         ReservedIPsService
	ReservedIPV6s       ReservedIPV6sService
	ReservedIPActions   ReservedIPActionsService
	ReservedIPV6Actions ReservedIPV6ActionsService
	Sizes               SizesService
	Snapshots           SnapshotsService
	SpacesKeys          SpacesKeysService
	Storage             StorageService
	StorageActions      StorageActionsService
	Tags                TagsService
	UptimeChecks        UptimeChecksService
	VectorDBs           VectorDBsService
	VPCs                VPCsService
	PartnerAttachment   PartnerAttachmentService
	GradientAI          GradientAIService
	DedicatedInference  DedicatedInferenceService
	BatchInference      BatchInferenceService
	BYOIPPrefixes       BYOIPPrefixesService

	// Serverless Inference resources at https://inference.do-ai.run.
	Chat             *ChatService
	Embeddings       *EmbeddingService
	ImageGenerations *ImageGenerationService
	Messages         *MessageService
	Models           *ModelService
	Responses        *ResponseService
	AsyncInvocations *AsyncInvocationService
	// Optional function called after every successful request made to the DO APIs
	onRequestCompleted RequestCompletionCallback

	// Optional extra HTTP headers to set on every request to the API.
	headers map[string]string

	// Optional rate limiter to ensure QoS.
	rateLimiter *rate.Limiter

	// Optional retry values. Setting the RetryConfig.RetryMax value enables automatically retrying requests
	// that fail with 429 or 500-level response codes using the go-retryablehttp client
	RetryConfig RetryConfig
}

// RetryConfig sets the values used for enabling retries and backoffs for
// requests that fail with 429 or 500-level response codes using the go-retryablehttp client.
// RetryConfig.RetryMax must be configured to enable this behavior. RetryConfig.RetryWaitMin and
// RetryConfig.RetryWaitMax are optional, with the default values being 1.0 and 30.0, respectively.
//
// You can use
//
//	godo.PtrTo(1.0)
//
// to explicitly set the RetryWaitMin and RetryWaitMax values.
//
// Note: Opting to use the go-retryablehttp client will overwrite any custom HTTP client passed into New().
// Only the oauth2.TokenSource and Timeout will be maintained.
type RetryConfig struct {
	RetryMax     int
	RetryWaitMin *float64    // Minimum time to wait
	RetryWaitMax *float64    // Maximum time to wait
	Logger       interface{} // Customer logger instance. Must implement either go-retryablehttp.Logger or go-retryablehttp.LeveledLogger
}

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

// ListOptions specifies the optional parameters to various List methods that
// support pagination.
type ListOptions struct {
	// For paginated result sets, page of results to retrieve.
	Page int `url:"page,omitempty"`

	// For paginated result sets, the number of results to include per page.
	PerPage int `url:"per_page,omitempty"`

	// Whether App responses should include project_id fields. The field will be empty if false or if omitted. (ListApps)
	WithProjects bool `url:"with_projects,omitempty"`

	// This parameter is used to only list agents that are deployed in the response.
	Deployed bool `url:"only_deployed,omitempty"`

	// This parameter is used to include models that are publicly available.
	PublicOnly bool `url:"public_only,omitempty"`

	// This parameter is used to include models according to the use cases.
	Usecases []string `url:"usecases,omitempty"`
}

// TokenListOptions specifies the optional parameters to various List methods that support token pagination.
type TokenListOptions struct {
	// For paginated result sets, page of results to retrieve.
	Page int `url:"page,omitempty"`

	// For paginated result sets, the number of results to include per page.
	PerPage int `url:"per_page,omitempty"`

	// For paginated result sets which support tokens, the token provided by the last set
	// of results in order to retrieve the next set of results. This is expected to be faster
	// than incrementing or decrementing the page number.
	Token string `url:"page_token,omitempty"`
}

// Response is a DigitalOcean response. This wraps the standard http.Response returned from DigitalOcean.
type Response struct {
	*http.Response

	// Links that were returned with the response. These are parsed from
	// request body and not the header.
	Links *Links

	// Meta describes generic information about the response.
	Meta *Meta

	// Monitoring URI
	// Deprecated: This field is not populated. To poll for the status of a
	// newly created Droplet, use Links.Actions[0].HREF
	Monitor string

	Rate
}

// An ErrorResponse reports the error caused by an API request
type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response

	// Error message
	Message string `json:"message"`

	// RequestID returned from the API, useful to contact support.
	RequestID string `json:"request_id"`

	// Attempts is the number of times the request was attempted when retries are enabled.
	Attempts int
}

// Rate contains the rate limit for the current client.
type Rate struct {
	// The number of request per hour the client is currently limited to.
	Limit int `json:"limit"`

	// The number of remaining requests the client can make this hour.
	Remaining int `json:"remaining"`

	// The time at which the current rate limit will reset.
	Reset Timestamp `json:"reset"`
}

func addOptions(s string, opt interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// NewFromToken returns a new DigitalOcean API client with the given API
// token.
func NewFromToken(token string) *Client { _ = "STUB: not implemented"; return nil }

// NewClient returns a new DigitalOcean API client, using the given
// http.Client to perform all requests.
//
// Users who wish to pass their own http.Client should use this method. If
// you're in need of further customization, the godo.New method allows more
// options, such as setting a custom URL or a custom user agent string.
func NewClient(httpClient *http.Client) *Client { _ = "STUB: not implemented"; return nil }

// ClientOpt are options for New.
type ClientOpt func(*Client) error

// New returns a new DigitalOcean API client instance.
func New(httpClient *http.Client, opts ...ClientOpt) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if retryMax is set it will use the retryablehttp client.

// By default this is nil and does not log.

// if timeout is set, it is maintained before overwriting client with StandardClient()

// This custom ErrorHandler is required to provide errors that are consistent
// with a *godo.ErrorResponse and a non-nil *godo.Response while providing
// insight into retries using an internal header.

// In addition to the default retry policy, we also retry HTTP/2 INTERNAL_ERROR errors.
// See: https://github.com/golang/go/issues/51323

// SetBaseURL is a client option for setting the base URL.
func SetBaseURL(bu string) ClientOpt { _ = "STUB: not implemented"; return *new(ClientOpt) }

// SetUserAgent is a client option for setting the user agent.
func SetUserAgent(ua string) ClientOpt { _ = "STUB: not implemented"; return *new(ClientOpt) }

// SetRequestHeaders sets optional HTTP headers on the client that are
// sent on each HTTP request.
func SetRequestHeaders(headers map[string]string) ClientOpt {
	_ = "STUB: not implemented"
	return *new(ClientOpt)
}

// SetStaticRateLimit sets an optional client-side rate limiter that restricts
// the number of queries per second that the client can send to enforce QoS.
func SetStaticRateLimit(rps float64) ClientOpt { _ = "STUB: not implemented"; return *new(ClientOpt) }

// WithRetryAndBackoffs sets retry values. Setting the RetryConfig.RetryMax value enables automatically retrying requests
// that fail with 429 or 500-level response codes using the go-retryablehttp client
func WithRetryAndBackoffs(retryConfig RetryConfig) ClientOpt {
	_ = "STUB: not implemented"
	return *new(ClientOpt)
}

// NewRequest creates an API request. A relative URL can be provided in urlStr, which will be resolved to the
// BaseURL of the Client. Relative URLS should always be specified without a preceding slash. If specified, the
// value pointed to by body is JSON encoded and included in as the request body.
func (c *Client) NewRequest(ctx context.Context, method, urlStr string, body interface{}) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// OnRequestCompleted sets the DO API request completion callback
func (c *Client) OnRequestCompleted(rc RequestCompletionCallback) {
	_ = "STUB: not implemented"
	return
}

// GetRate returns the current rate limit for the client as determined by the most recent
// API call. It is thread-safe.
func (c *Client) GetRate() Rate { _ = "STUB: not implemented"; return *new(Rate) }

// newResponse creates a new Response for the provided http.Response
func newResponse(r *http.Response) *Response { _ = "STUB: not implemented"; return nil }

// populateRate parses the rate related headers and populates the response Rate.
func (r *Response) populateRate() { _ = "STUB: not implemented"; return }

// Do sends an API request and returns the API response. The API response is JSON decoded and stored in the value
// pointed to by v, or returned as an error if an API error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Ensure the response body is fully read and closed
// before we reconnect, so that we reuse the same TCPConnection.
// Close the previous response's body. But read at least some of
// the body so if it's small the underlying TCP connection will be
// re-used. No need to check for errors: if it fails, the Transport
// won't reuse it anyway.

// DoStream sends an API request and returns the response with its body
// left open for streaming consumption (e.g. text/event-stream). On 2xx,
// the caller owns resp.Body and must close it. On non-2xx, the body is
// drained and closed and the typed *ErrorResponse is returned.
func (c *Client) DoStream(ctx context.Context, req *http.Request) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DoRequest submits an HTTP request.
func DoRequest(ctx context.Context, req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DoRequestWithClient submits an HTTP request using the specified client.
func DoRequestWithClient(
	ctx context.Context,
	client *http.Client,
	req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ErrorResponse) Error() string { _ = "STUB: not implemented"; return "" }

// CheckResponse checks the API response for errors, and returns them if present. A response is considered an
// error if it has a status code outside the 200 range. API error responses are expected to have either no response
// body, or a JSON response body that maps to ErrorResponse. Any other response body will be silently ignored.
// If the API error response does not include the request ID in its body, the one from its header will be used.
func CheckResponse(r *http.Response) error { _ = "STUB: not implemented"; return nil }

func (r Rate) String() string { _ = "STUB: not implemented"; return "" }

// PtrTo returns a pointer to the provided input.
func PtrTo[T any](v T) *T {
	_ = "STUB: not implemented"

	// String is a helper routine that allocates a new string value
	// to store v and returns a pointer to it.
	//
	// Deprecated: Use PtrTo instead.
	return nil
}

func String(v string) *string { _ = "STUB: not implemented"; return nil }

// Int is a helper routine that allocates a new int32 value
// to store v and returns a pointer to it, but unlike Int32
// its argument value is an int.
//
// Deprecated: Use PtrTo instead.
func Int(v int) *int { _ = "STUB: not implemented"; return nil }

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
//
// Deprecated: Use PtrTo instead.
func Bool(v bool) *bool { _ = "STUB: not implemented"; return nil }

// StreamToString converts a reader to a string
func StreamToString(stream io.Reader) string { _ = "STUB: not implemented"; return "" }
