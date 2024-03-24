package middleware

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/justinas/alice"
)

var (
	Chain = alice.New
)

// Go Chi v5
var (
	// Enforces a whitelist of request Content-Encoding headers
	AllowContentEncoding = middleware.AllowContentEncoding

	// Explicit whitelist of accepted request Content-Types
	AllowContentType = middleware.AllowContentType

	// Basic HTTP authentication
	BasicAuth = middleware.BasicAuth

	// Gzip compression for clients that accept compressed responses
	Compress = middleware.Compress

	// Ensure charset for Content-Type request headers
	ContentCharset = middleware.ContentCharset

	// Clean double slashes from request path
	CleanPath = middleware.CleanPath

	// Automatically route undefined HEAD requests to GET handlers
	GetHead = middleware.GetHead

	// Monitoring endpoint to check the servers pulse
	Heartbeat = middleware.Heartbeat

	// Logs the start and end of each request with the elapsed processing time
	Logger = middleware.Logger

	// Sets response headers to prevent clients from caching
	NoCache = middleware.NoCache

	// Easily attach net/http/pprof to your routers
	Profiler = middleware.Profiler

	// Sets a http.Request's RemoteAddr to either X-Real-IP or X-Forwarded-For
	RealIP = middleware.RealIP

	// Gracefully absorb panics and prints the stack trace
	Recoverer = middleware.Recoverer

	// Injects a request ID into the context of each request
	RequestID = middleware.RequestID

	// Redirect slashes on routing paths
	RedirectSlashes = middleware.RedirectSlashes

	// Route handling for request headers
	RouteHeaders = middleware.RouteHeaders

	// Short-hand middleware to set a response header key/value
	SetHeader = middleware.SetHeader

	// Strip slashes on routing paths
	StripSlashes = middleware.StripSlashes

	// Sunset set Deprecation/Sunset header to response
	Sunset = middleware.Sunset

	// Puts a ceiling on the number of concurrent requests
	Throttle = middleware.Throttle

	// Signals to the request context when the timeout deadline is reached
	Timeout = middleware.Timeout

	// Parse extension from url and put it on request context
	URLFormat = middleware.URLFormat

	// Short-hand middleware to set a key/value on the request context
	WithValue = middleware.WithValue
)
