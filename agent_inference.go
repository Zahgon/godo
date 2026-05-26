package godo

import (
	"context"
	"net/http"
	"net/url"
)

// Agent Inference API (https://docs.digitalocean.com/reference/api/reference/agent-inference/).
// Each provisioned agent has its own base URL (e.g. https://abc123.agents.do-ai.run)
// and an agent_access_key that is NOT a dop_v1_* token, so this API lives on its
// own *AgentInferenceClient (one per agent endpoint)

const (
	agentInferenceChatCompletionsPath = "api/v1/chat/completions"
	agentInferenceQueryAgent          = "agent"
)

type AgentInferenceClient struct {
	Chat *AgentChatService

	baseURL    *url.URL
	accessKey  string
	httpClient *http.Client
	userAgent  string

	headers            map[string]string
	onRequestCompleted RequestCompletionCallback
}

// AgentInferenceClientOpt is a functional option for an AgentInferenceClient.
type AgentInferenceClientOpt func(*AgentInferenceClient) error

type AgentChatService struct {
	Completions *AgentChatCompletionService
}

// AgentChatCompletionService exposes POST /api/v1/chat/completions?agent=true.
type AgentChatCompletionService struct {
	parent *AgentInferenceClient
}

func NewAgentInferenceClient(baseURL, accessKey string, opts ...AgentInferenceClientOpt) (*AgentInferenceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetAgentHTTPClient overrides the http.Client used for all requests.
func SetAgentHTTPClient(hc *http.Client) AgentInferenceClientOpt {
	_ = "STUB: not implemented"
	return *new(AgentInferenceClientOpt)
}

// SetAgentUserAgent prepends ua to the User-Agent header.
func SetAgentUserAgent(ua string) AgentInferenceClientOpt {
	_ = "STUB: not implemented"
	return *new(AgentInferenceClientOpt)
}

// SetAgentRequestHeaders adds default request headers. Authorization, Content-Type,
// Accept, and User-Agent are reserved and overwritten by the client.
func SetAgentRequestHeaders(headers map[string]string) AgentInferenceClientOpt {
	_ = "STUB: not implemented"
	return *new(AgentInferenceClientOpt)
}

// OnRequestCompleted registers a callback fired after each HTTP request.
func (c *AgentInferenceClient) OnRequestCompleted(rc RequestCompletionCallback) {
	_ = "STUB: not implemented"
	return
}

// BaseURL returns the (normalised) agent base URL.
func (c *AgentInferenceClient) BaseURL() *url.URL { _ = "STUB: not implemented"; return nil }

// New creates a non-streaming agent chat completion.
func (s *AgentChatCompletionService) New(ctx context.Context, body *ChatCompletionNewParams) (*ChatCompletion, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// NewStreaming opens an SSE stream of chat completion chunks; body.Stream is forced to true.
// Callers MUST Close the returned stream.
func (s *AgentChatCompletionService) NewStreaming(ctx context.Context, body *ChatCompletionNewParams) (*ChatCompletionStream, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// newRequest builds a request against the agent base URL with ?agent=true and bearer auth.
func (c *AgentInferenceClient) newRequest(ctx context.Context, method, path string, body interface{}) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// do executes a non-streaming request and decodes a 2xx body into v.
func (c *AgentInferenceClient) do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// doStream executes a streaming request and wraps the open body in an *InferenceStream.
func (c *AgentInferenceClient) doStream(ctx context.Context, req *http.Request) (*InferenceStream, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
