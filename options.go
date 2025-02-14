package flashbotsrpc

import (
	"io"
	"net/http"
)

type httpClient interface {
	Post(url string, contentType string, body io.Reader) (*http.Response, error)
	Do(req *http.Request) (*http.Response, error)
}

type logger interface {
	Println(v ...interface{})
}

// WithHttpClient set custom http client
func WithHttpClient(client httpClient) func(rpc *FlashbotsRPC) {
	return func(rpc *FlashbotsRPC) {
		rpc.client = client
	}
}

// WithLogger set custom logger
func WithLogger(l logger) func(rpc *FlashbotsRPC) {
	return func(rpc *FlashbotsRPC) {
		rpc.log = l
	}
}

// WithDebug set debug flag
func WithDebug(enabled bool) func(rpc *FlashbotsRPC) {
	return func(rpc *FlashbotsRPC) {
		rpc.Debug = enabled
	}
}

// BroadcastRPC WithLogger set custom logger
func BroadcastRPCWithLogger(l logger) func(rpc *BuilderBroadcastRPC) {
	return func(rpc *BuilderBroadcastRPC) {
		rpc.log = l
	}
}

// BroadcastRPC WithHttpClient set custom http client
func BroadcastRPCWithHttpClient(client httpClient) func(rpc *BuilderBroadcastRPC) {
	return func(rpc *BuilderBroadcastRPC) {
		rpc.client = client
	}
}

// BroadcastRPC WithDebug set debug flag
func BroadcastRPCWithDebug(enabled bool) func(rpc *BuilderBroadcastRPC) {
	return func(rpc *BuilderBroadcastRPC) {
		rpc.Debug = enabled
	}
}
