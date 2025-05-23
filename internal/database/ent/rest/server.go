// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// DO NOT EDIT, CODE GENERATED BY entc.

package rest

import (
	"bytes"
	_ "embed"
	"encoding"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/url"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/form/v4"
	"github.com/lrstanley/liam.sh/internal/database/ent"
	"github.com/lrstanley/liam.sh/internal/database/ent/githubasset"
	"github.com/lrstanley/liam.sh/internal/database/ent/githubevent"
	"github.com/lrstanley/liam.sh/internal/database/ent/githubgist"
	"github.com/lrstanley/liam.sh/internal/database/ent/githubrelease"
	"github.com/lrstanley/liam.sh/internal/database/ent/githubrepository"
	"github.com/lrstanley/liam.sh/internal/database/ent/label"
	"github.com/lrstanley/liam.sh/internal/database/ent/post"
	"github.com/lrstanley/liam.sh/internal/database/ent/privacy"
	"github.com/lrstanley/liam.sh/internal/database/ent/user"
)

//go:embed openapi.json
var OpenAPI []byte // OpenAPI contains the JSON schema of the API.

// Operation represents the CRUD operation(s).
type Operation string

const (
	// OperationCreate represents the create operation (method: POST).
	OperationCreate Operation = "create"
	// OperationRead represents the read operation (method: GET).
	OperationRead Operation = "read"
	// OperationUpdate represents the update operation (method: PATCH).
	OperationUpdate Operation = "update"
	// OperationDelete represents the delete operation (method: DELETE).
	OperationDelete Operation = "delete"
	// OperationList represents the list operation (method: GET).
	OperationList Operation = "list"
)

// ErrorResponse is the response structure for errors.
type ErrorResponse struct {
	Error     string `json:"error"`                // The underlying error, which may be masked when debugging is disabled.
	Type      string `json:"type"`                 // A summary of the error code based off the HTTP status code or application error code.
	Code      int    `json:"code"`                 // The HTTP status code or other internal application error code.
	RequestID string `json:"request_id,omitempty"` // The unique request ID for this error.
	Timestamp string `json:"timestamp,omitempty"`  // The timestamp of the error, in RFC3339 format.
}

type ErrBadRequest struct {
	Err error
}

func (e ErrBadRequest) Error() string {
	return fmt.Sprintf("bad request: %s", e.Err)
}

func (e ErrBadRequest) Unwrap() error {
	return e.Err
}

// IsBadRequest returns true if the unwrapped/underlying error is of type ErrBadRequest.
func IsBadRequest(err error) bool {
	var target *ErrBadRequest
	return errors.As(err, &target)
}

var ErrEndpointNotFound = errors.New("endpoint not found")

// IsEndpointNotFound returns true if the unwrapped/underlying error is of type ErrEndpointNotFound.
func IsEndpointNotFound(err error) bool {
	return errors.Is(err, ErrEndpointNotFound)
}

var ErrMethodNotAllowed = errors.New("method not allowed")

// IsMethodNotAllowed returns true if the unwrapped/underlying error is of type ErrMethodNotAllowed.
func IsMethodNotAllowed(err error) bool {
	return errors.Is(err, ErrMethodNotAllowed)
}

type ErrInvalidID struct {
	ID  string
	Err error
}

func (e ErrInvalidID) Error() string {
	return fmt.Sprintf("invalid ID provided: %q: %v", e.ID, e.Err)
}

func (e ErrInvalidID) Unwrap() error {
	return e.Err
}

// IsInvalidID returns true if the unwrapped/underlying error is of type ErrInvalidID.
func IsInvalidID(err error) bool {
	var target *ErrInvalidID
	return errors.As(err, &target)
}

// JSON marshals 'v' to JSON, and setting the Content-Type as application/json.
// Note that this does NOT auto-escape HTML. If 'v' cannot be marshalled to JSON,
// this will panic.
//
// JSON also supports prettification when the origin request has a query parameter
// of "pretty" set to true.
func JSON(w http.ResponseWriter, r *http.Request, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	enc := json.NewEncoder(w)

	if pretty, _ := strconv.ParseBool(r.FormValue("pretty")); pretty {
		enc.SetIndent("", "    ")
	}

	if err := enc.Encode(v); err != nil && err != io.EOF {
		panic(fmt.Sprintf("failed to marshal response: %v", err))
	}
}

// M is an alias for map[string]any, which makes it easier to respond with generic JSON data structures.
type M map[string]any

var (
	// DefaultDecoder is the default decoder used by Bind. You can either override
	// this, or provide your own. Make sure it is set before Bind is called.
	DefaultDecoder = form.NewDecoder()

	// DefaultDecodeMaxMemory is the maximum amount of memory in bytes that will be
	// used for decoding multipart/form-data requests.
	DefaultDecodeMaxMemory int64 = 8 << 20
)

// Bind decodes the request body to the given struct. At this time the only supported
// content-types are application/json, application/x-www-form-urlencoded, as well as
// GET parameters.
func Bind(r *http.Request, v any) error {
	err := r.ParseForm()
	if err != nil {
		return &ErrBadRequest{Err: fmt.Errorf("parsing form parameters: %w", err)}
	}

	switch r.Method {
	case http.MethodGet, http.MethodHead:
		err = DefaultDecoder.Decode(v, r.Form)
	case http.MethodPost, http.MethodPut, http.MethodPatch:
		switch {
		case strings.HasPrefix(r.Header.Get("Content-Type"), "application/json"):
			dec := json.NewDecoder(r.Body)
			dec.DisallowUnknownFields()
			defer r.Body.Close()
			err = dec.Decode(v)
		case strings.HasPrefix(r.Header.Get("Content-Type"), "multipart/form-data"):
			err = r.ParseMultipartForm(DefaultDecodeMaxMemory)
			if err == nil {
				err = DefaultDecoder.Decode(v, r.MultipartForm.Value)
			}
		default:
			err = DefaultDecoder.Decode(v, r.PostForm)
		}
	default:
		return &ErrBadRequest{Err: fmt.Errorf("unsupported method %s", r.Method)}
	}

	if err != nil {
		return &ErrBadRequest{Err: fmt.Errorf("error decoding %s request into required format (%T): %w", r.Method, v, err)}
	}
	return nil
}

// Req simplifies making an HTTP handler that returns a single result, and an error.
// The result, if not nil, must be JSON-marshalable. If result is nil, [http.StatusNoContent]
// will be returned.
func Req[Resp any](s *Server, op Operation, fn func(*http.Request) (*Resp, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		results, err := fn(r)
		handleResponse(s, w, r, op, results, err)
	}
}

// resolveID resolves the ID from the request path, and unmarshals it into the provided type.
// Only supports string, int, and types that support UnmarshalText, UnmarshalJSON, or UnmarshalBinary
// (in that order).
func resolveID[T any](r *http.Request) (id T, err error) {
	value := r.PathValue("id")

	switch any(id).(type) {
	case string:
		id = any(value).(T)
	case int:
		rid, err := strconv.Atoi(value)
		if err == nil {
			id = any(rid).(T)
		}
	default:
		hasUnmarshal := false

		// Check if the underlying type supports UnmarshalText, UnmarshalJSON, or UnmarshalBinary.
		if u, ok := any(&id).(encoding.TextUnmarshaler); ok {
			hasUnmarshal = true
			err = u.UnmarshalText([]byte(value))
		} else if u, ok := any(&id).(json.Unmarshaler); ok {
			hasUnmarshal = true
			err = u.UnmarshalJSON([]byte(value))
		} else if u, ok := any(&id).(encoding.BinaryUnmarshaler); ok {
			hasUnmarshal = true
			err = u.UnmarshalBinary([]byte(value))
		}

		if !hasUnmarshal {
			panic(fmt.Sprintf("unsupported ID type (cannot unmarshal): %T", id))
		}
	}

	if err != nil {
		return id, &ErrInvalidID{ID: value, Err: err}
	}
	return id, nil
}

// ReqID is similar to Req, but also processes an "id" path parameter and provides it to the
// handler function.
func ReqID[Resp, I any](s *Server, op Operation, fn func(*http.Request, I) (*Resp, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := resolveID[I](r)
		if err != nil {
			handleResponse[Resp](s, w, r, op, nil, err)
			return
		}
		results, err := fn(r, id)
		handleResponse(s, w, r, op, results, err)
	}
}

// ReqParam is similar to Req, but also processes a request body/query params and provides it
// to the handler function.
func ReqParam[Params, Resp any](s *Server, op Operation, fn func(*http.Request, *Params) (*Resp, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := new(Params)
		if err := Bind(r, params); err != nil {
			handleResponse[Resp](s, w, r, op, nil, err)
			return
		}
		results, err := fn(r, params)
		handleResponse(s, w, r, op, results, err)
	}
}

// ReqIDParam is similar to ReqParam, but also processes an "id" path parameter and request
// body/query params, and provides it to the handler function.
func ReqIDParam[Params, Resp, I any](s *Server, op Operation, fn func(*http.Request, I, *Params) (*Resp, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := resolveID[I](r)
		if err != nil {
			handleResponse[Resp](s, w, r, op, nil, err)
			return
		}
		params := new(Params)
		err = Bind(r, params)
		if err != nil {
			handleResponse[Resp](s, w, r, op, nil, err)
			return
		}
		results, err := fn(r, id, params)
		handleResponse(s, w, r, op, results, err)
	}
}

// Links represents a set of linkable-relationsips that can be represented through
// the "Link" header. Note that all urls must be url-encoded already.
type Links map[string]string

func (l Links) String() string {
	var links []string
	var keys []string
	for k := range l {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	for _, k := range keys {
		links = append(links, fmt.Sprintf(`<%s>; rel=%q`, l[k], k))
	}
	return strings.Join(links, ", ")
}

type linkablePagedResource interface {
	GetPage() int
	GetIsLastPage() bool
}

// Spec returns the OpenAPI spec for the server implementation.
func (s *Server) Spec(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if !s.config.DisableSpecInjectServer && s.config.BaseURL != "" {
		spec := map[string]any{}
		err := json.Unmarshal(OpenAPI, &spec)
		if err != nil {
			panic(fmt.Sprintf("failed to unmarshal spec: %v", err))
		}

		type Server struct {
			URL string `json:"url"`
		}

		if _, ok := spec["servers"]; !ok {
			spec["servers"] = []Server{{URL: s.config.BaseURL}}
			JSON(w, r, http.StatusOK, spec)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(OpenAPI)
}

var scalarTemplate = template.Must(template.New("docs").Parse(`<!DOCTYPE html>
<html>
  <head>
    <title>API Reference</title>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <meta name="darkreader-lock">
    <link rel="icon" type="image/svg+xml"
      href="data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='32' height='32' viewBox='0 0 1024 1024'%3E%3Cpath fill='currentColor' d='m917.7 148.8l-42.4-42.4c-1.6-1.6-3.6-2.3-5.7-2.3s-4.1.8-5.7 2.3l-76.1 76.1a199.27 199.27 0 0 0-112.1-34.3c-51.2 0-102.4 19.5-141.5 58.6L432.3 308.7a8.03 8.03 0 0 0 0 11.3L704 591.7c1.6 1.6 3.6 2.3 5.7 2.3c2 0 4.1-.8 5.7-2.3l101.9-101.9c68.9-69 77-175.7 24.3-253.5l76.1-76.1c3.1-3.2 3.1-8.3 0-11.4M769.1 441.7l-59.4 59.4l-186.8-186.8l59.4-59.4c24.9-24.9 58.1-38.7 93.4-38.7s68.4 13.7 93.4 38.7c24.9 24.9 38.7 58.1 38.7 93.4s-13.8 68.4-38.7 93.4m-190.2 105a8.03 8.03 0 0 0-11.3 0L501 613.3L410.7 523l66.7-66.7c3.1-3.1 3.1-8.2 0-11.3L441 408.6a8.03 8.03 0 0 0-11.3 0L363 475.3l-43-43a7.85 7.85 0 0 0-5.7-2.3c-2 0-4.1.8-5.7 2.3L206.8 534.2c-68.9 69-77 175.7-24.3 253.5l-76.1 76.1a8.03 8.03 0 0 0 0 11.3l42.4 42.4c1.6 1.6 3.6 2.3 5.7 2.3s4.1-.8 5.7-2.3l76.1-76.1c33.7 22.9 72.9 34.3 112.1 34.3c51.2 0 102.4-19.5 141.5-58.6l101.9-101.9c3.1-3.1 3.1-8.2 0-11.3l-43-43l66.7-66.7c3.1-3.1 3.1-8.2 0-11.3zM441.7 769.1a131.32 131.32 0 0 1-93.4 38.7c-35.3 0-68.4-13.7-93.4-38.7a131.32 131.32 0 0 1-38.7-93.4c0-35.3 13.7-68.4 38.7-93.4l59.4-59.4l186.8 186.8z'/%3E%3C/svg%3E" />
  </head>
  <body>
    <script id="api-reference"></script>
    <script>
      document.getElementById("api-reference").dataset.configuration = JSON.stringify({
        spec: {
          url: "{{ $.SpecPath }}",
        },
        {{- if $.DisableSpecInjectServer }}
        servers: [
            {url: window.location.origin + window.location.pathname.replace(/\/docs$/g, "")}
        ],
        {{- end }}
        theme: "kepler",
        isEditable: false,
        hideDownloadButton: true,
        customCss: ".darklight-reference-promo { visibility: hidden !important; height: 0 !important; } .open-api-client-button { display: none !important; }",
      });
    </script>
    <script
      src="https://cdn.jsdelivr.net/npm/@scalar/api-reference@1.25.99"
      integrity="sha256-uuOSJ5AN5uhftgJPc4yYZ3PZpdYcsolIRamzCKTiJnE="
      crossorigin="anonymous"
    ></script>
  </body>
</html>`))

func (s *Server) Docs(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	err := scalarTemplate.Execute(&buf, map[string]any{
		"SpecPath":                s.config.BasePath + "/openapi.json",
		"DisableSpecInjectServer": s.config.DisableSpecInjectServer,
	})
	if err != nil {
		handleResponse[struct{}](s, w, r, "", nil, err)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Content-Security-Policy", "default-src 'self' cdn.jsdelivr.net fonts.scalar.com 'unsafe-inline' 'unsafe-eval' data: blob:")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Referrer-Policy", "no-referrer-when-downgrade")
	w.Header().Set("Permissions-Policy", "clipboard-write=(self)")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(buf.Bytes())
}

type ServerConfig struct {
	// BaseURL is similar to [ServerConfig.BasePath], however, only the path of the URL is used
	// to prefill BasePath. This is not required if BasePath is provided.
	BaseURL string

	// BasePath if provided, and the /openapi.json endpoint is enabled, will allow annotating
	// API responses with "Link" headers. See [ServerConfig.EnableLinks] for more information.
	BasePath string

	// DisableSpecHandler if set to true, will disable the /openapi.json endpoint. This will also
	// disable the embedded API reference documentation, see [ServerConfig.DisableDocs] for more
	// information.
	DisableSpecHandler bool

	// DisableSpecInjectServer if set to true, will disable the automatic injection of the
	// server URL into the spec. This only applies if [ServerConfig.BaseURL] is provided.
	DisableSpecInjectServer bool

	// DisableDocsHandler if set to true, will disable the embedded API reference documentation
	// endpoint at /docs. Use this if you want to provide your own documentation functionality.
	// This is disabled by default if [ServerConfig.DisableSpecHandler] is true.
	DisableDocsHandler bool

	// EnableLinks if set to true, will enable the "Link" response header, which can be used to hint
	// to clients about the location of the OpenAPI spec, API documentation, how to auto-paginate
	// through results, and more.
	EnableLinks bool

	// MaskErrors if set to true, will mask the error message returned to the client,
	// returning a generic error message based on the HTTP status code.
	MaskErrors bool

	// ErrorHandler is invoked when an error occurs. If not provided, the default
	// error handling logic will be used. If you want to run logic on errors, but
	// not actually handle the error yourself, you can still call [Server.DefaultErrorHandler]
	// after your logic.
	ErrorHandler func(w http.ResponseWriter, r *http.Request, op Operation, err error)

	// GetReqID returns the request ID for the given request. If not provided, the
	// default implementation will use the X-Request-Id header, otherwise an empty
	// string will be returned. If using go-chi, middleware.GetReqID will be used.
	GetReqID func(r *http.Request) string
}

type Server struct {
	db     *ent.Client
	config *ServerConfig
}

// NewServer returns a new auto-generated server implementation for your ent schema.
// [Server.Handler] returns a ready-to-use http.Handler that mounts all of the
// necessary endpoints.
func NewServer(db *ent.Client, config *ServerConfig) (*Server, error) {
	s := &Server{
		db:     db,
		config: config,
	}
	if s.config == nil {
		s.config = &ServerConfig{}
	}
	if s.config.BaseURL != "" && s.config.BasePath == "" {
		uri, err := url.Parse(s.config.BaseURL)
		if err != nil {
			return nil, fmt.Errorf("failed to parse BaseURL: %w", err)
		}
		s.config.BasePath = uri.Path
	}
	if s.config.BaseURL == "" {
		s.config.DisableSpecInjectServer = true
	}
	if s.config.BasePath != "" {
		if !strings.HasPrefix(s.config.BasePath, "/") {
			s.config.BasePath = "/" + s.config.BasePath
		}
		s.config.BasePath = strings.TrimRight(s.config.BasePath, "/")
	}
	return s, nil
}

// DefaultErrorHandler is the default error handler for the Server.
func (s *Server) DefaultErrorHandler(w http.ResponseWriter, r *http.Request, op Operation, err error) {
	ts := time.Now().UTC().Format(time.RFC3339)

	resp := ErrorResponse{
		Error:     err.Error(),
		Timestamp: ts,
	}

	var numErr *strconv.NumError

	switch {
	case IsEndpointNotFound(err):
		resp.Code = http.StatusNotFound
	case IsMethodNotAllowed(err):
		resp.Code = http.StatusMethodNotAllowed
	case IsBadRequest(err):
		resp.Code = http.StatusBadRequest
	case IsInvalidID(err):
		resp.Code = http.StatusBadRequest
	case errors.Is(err, privacy.Deny):
		resp.Code = http.StatusForbidden
	case ent.IsNotFound(err):
		resp.Code = http.StatusNotFound
	case ent.IsConstraintError(err), ent.IsNotSingular(err):
		resp.Code = http.StatusConflict
	case ent.IsValidationError(err):
		resp.Code = http.StatusBadRequest
	case errors.As(err, &numErr):
		resp.Code = http.StatusBadRequest
		resp.Error = fmt.Sprintf("invalid ID provided: %v", err)
	default:
		resp.Code = http.StatusInternalServerError
	}

	if resp.Type == "" {
		resp.Type = http.StatusText(resp.Code)
	}
	if s.config.MaskErrors {
		resp.Error = http.StatusText(resp.Code)
	}
	if s.config.GetReqID != nil {
		resp.RequestID = s.config.GetReqID(r)
	} else {
		resp.RequestID = middleware.GetReqID(r.Context())
	}
	JSON(w, r, resp.Code, resp)
}

func handleResponse[Resp any](s *Server, w http.ResponseWriter, r *http.Request, op Operation, resp *Resp, err error) {
	if s.config.EnableLinks {
		links := Links{}
		if !s.config.DisableSpecHandler {
			links["service-desc"] = s.config.BasePath + "/openapi.json"
			links["describedby"] = s.config.BasePath + "/openapi.json"
		}

		if err == nil && resp != nil && op == OperationList {
			if lr, ok := any(resp).(linkablePagedResource); ok {
				query := r.URL.Query()
				if page := lr.GetPage(); page > 1 {
					query.Set("page", strconv.Itoa(page-1))
					r.URL.RawQuery = query.Encode()
					links["prev"] = r.URL.String()
					if !strings.HasPrefix(links["prev"], s.config.BasePath) {
						links["prev"] = s.config.BasePath + links["prev"]
					}
				}
				if !lr.GetIsLastPage() {
					query.Set("page", strconv.Itoa(lr.GetPage()+1))
					r.URL.RawQuery = query.Encode()
					links["next"] = r.URL.String()
					if !strings.HasPrefix(links["next"], s.config.BasePath) {
						links["next"] = s.config.BasePath + links["next"]
					}
				}
			}
		}

		if v := links.String(); v != "" {
			w.Header().Set("Link", v)
		}
	}
	if err != nil {
		if s.config.ErrorHandler != nil {
			s.config.ErrorHandler(w, r, op, err)
			return
		}
		s.DefaultErrorHandler(w, r, op, err)
		return
	}
	if resp != nil {
		type pagedResp interface {
			GetTotalCount() int
		}
		if r.Method == http.MethodPost {
			JSON(w, r, http.StatusCreated, resp)
			return
		}
		JSON(w, r, http.StatusOK, resp)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// UseEntContext can be used to inject an [ent.Client] into the context for use
// by other middleware, or ent privacy layers. Note that the server will do this
// by default, so you don't need to do this manually, unless it's a context that's
// not being passed to the server and is being consumed elsewhere.
func UseEntContext(db *ent.Client) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r.WithContext(ent.NewContext(r.Context(), db)))
		})
	}
}

// Handler mounts all of the necessary endpoints onto the provided chi.Router.
func (s *Server) Handler(r chi.Router) {
	r.Use(UseEntContext(s.db))
	r.Get("/github-assets", ReqParam(s, OperationList, s.ListGithubAssets))
	r.Get("/github-assets/{id}", ReqID(s, OperationRead, s.GetGithubAsset))
	r.Get("/github-assets/{id}/release", ReqID(s, OperationRead, s.GetGithubAssetRelease))
	r.Get("/github-events", ReqParam(s, OperationList, s.ListGithubEvents))
	r.Get("/github-events/{id}", ReqID(s, OperationRead, s.GetGithubEvent))
	r.Get("/github-gists", ReqParam(s, OperationList, s.ListGithubGists))
	r.Get("/github-gists/{id}", ReqID(s, OperationRead, s.GetGithubGist))
	r.Get("/github-releases", ReqParam(s, OperationList, s.ListGithubReleases))
	r.Get("/github-releases/{id}", ReqID(s, OperationRead, s.GetGithubRelease))
	r.Get("/github-releases/{id}/repository", ReqID(s, OperationRead, s.GetGithubReleaseRepository))
	r.Get("/github-releases/{id}/assets", ReqIDParam(s, OperationList, s.ListGithubReleaseAssets))
	r.Get("/github-repositories", ReqParam(s, OperationList, s.ListGithubRepositories))
	r.Get("/github-repositories/{id}", ReqID(s, OperationRead, s.GetGithubRepository))
	r.Get("/github-repositories/{id}/labels", ReqIDParam(s, OperationList, s.ListGithubRepositoryLabels))
	r.Get("/github-repositories/{id}/releases", ReqIDParam(s, OperationList, s.ListGithubRepositoryReleases))
	r.Get("/labels", ReqParam(s, OperationList, s.ListLabels))
	r.Get("/labels/{id}", ReqID(s, OperationRead, s.GetLabel))
	r.Get("/labels/{id}/posts", ReqIDParam(s, OperationList, s.ListLabelPosts))
	r.Get("/labels/{id}/github-repositories", ReqIDParam(s, OperationList, s.ListLabelGithubRepositories))
	r.Post("/labels", ReqParam(s, OperationCreate, s.CreateLabel))
	r.Patch("/labels/{id}", ReqIDParam(s, OperationUpdate, s.UpdateLabel))
	r.Delete("/labels/{id}", ReqID(s, OperationDelete, s.DeleteLabel))
	r.Get("/posts", ReqParam(s, OperationList, s.ListPosts))
	r.Get("/posts/{id}", ReqID(s, OperationRead, s.GetPost))
	r.Get("/posts/{id}/author", ReqID(s, OperationRead, s.GetPostAuthor))
	r.Get("/posts/{id}/labels", ReqIDParam(s, OperationList, s.ListPostLabels))
	r.Post("/posts", ReqParam(s, OperationCreate, s.CreatePost))
	r.Patch("/posts/{id}", ReqIDParam(s, OperationUpdate, s.UpdatePost))
	r.Delete("/posts/{id}", ReqID(s, OperationDelete, s.DeletePost))
	r.Get("/users", ReqParam(s, OperationList, s.ListUsers))
	r.Get("/users/{id}", ReqID(s, OperationRead, s.GetUser))
	r.Get("/users/{id}/posts", ReqIDParam(s, OperationList, s.ListUserPosts))

	if !s.config.DisableSpecHandler {
		r.Get("/openapi.json", s.Spec)
	}

	if !s.config.DisableSpecHandler && !s.config.DisableDocsHandler {
		r.Get("/docs", s.Docs)
	}

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		handleResponse[struct{}](s, w, r, "", nil, ErrEndpointNotFound)
	})
	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		handleResponse[struct{}](s, w, r, "", nil, ErrMethodNotAllowed)
	})
}

// ListGithubAssets maps to "GET /github-assets".
func (s *Server) ListGithubAssets(r *http.Request, p *ListGithubAssetParams) (*PagedResponse[ent.GithubAsset], error) {
	return p.Exec(r.Context(), s.db.GithubAsset.Query())
}

// GetGithubAsset maps to "GET /github-assets/{id}".
func (s *Server) GetGithubAsset(r *http.Request, githubassetID int) (*ent.GithubAsset, error) {
	return EagerLoadGithubAsset(s.db.GithubAsset.Query().Where(githubasset.ID(githubassetID))).Only(r.Context())
}

// GetGithubAssetRelease maps to "GET /github-assets/{id}/release".
func (s *Server) GetGithubAssetRelease(r *http.Request, githubassetID int) (*ent.GithubRelease, error) {
	return EagerLoadGithubRelease(s.db.GithubAsset.Query().Where(githubasset.ID(githubassetID)).QueryRelease()).Only(r.Context())
}

// ListGithubEvents maps to "GET /github-events".
func (s *Server) ListGithubEvents(r *http.Request, p *ListGithubEventParams) (*PagedResponse[ent.GithubEvent], error) {
	return p.Exec(r.Context(), s.db.GithubEvent.Query())
}

// GetGithubEvent maps to "GET /github-events/{id}".
func (s *Server) GetGithubEvent(r *http.Request, githubeventID int) (*ent.GithubEvent, error) {
	return EagerLoadGithubEvent(s.db.GithubEvent.Query().Where(githubevent.ID(githubeventID))).Only(r.Context())
}

// ListGithubGists maps to "GET /github-gists".
func (s *Server) ListGithubGists(r *http.Request, p *ListGithubGistParams) (*PagedResponse[ent.GithubGist], error) {
	return p.Exec(r.Context(), s.db.GithubGist.Query())
}

// GetGithubGist maps to "GET /github-gists/{id}".
func (s *Server) GetGithubGist(r *http.Request, githubgistID int) (*ent.GithubGist, error) {
	return EagerLoadGithubGist(s.db.GithubGist.Query().Where(githubgist.ID(githubgistID))).Only(r.Context())
}

// ListGithubReleases maps to "GET /github-releases".
func (s *Server) ListGithubReleases(r *http.Request, p *ListGithubReleaseParams) (*PagedResponse[ent.GithubRelease], error) {
	return p.Exec(r.Context(), s.db.GithubRelease.Query())
}

// GetGithubRelease maps to "GET /github-releases/{id}".
func (s *Server) GetGithubRelease(r *http.Request, githubreleaseID int) (*ent.GithubRelease, error) {
	return EagerLoadGithubRelease(s.db.GithubRelease.Query().Where(githubrelease.ID(githubreleaseID))).Only(r.Context())
}

// GetGithubReleaseRepository maps to "GET /github-releases/{id}/repository".
func (s *Server) GetGithubReleaseRepository(r *http.Request, githubreleaseID int) (*ent.GithubRepository, error) {
	return EagerLoadGithubRepository(s.db.GithubRelease.Query().Where(githubrelease.ID(githubreleaseID)).QueryRepository()).Only(r.Context())
}

// ListGithubReleaseAssets maps to "GET /github-releases/{id}/assets".
func (s *Server) ListGithubReleaseAssets(r *http.Request, githubreleaseID int, p *ListGithubAssetParams) (*PagedResponse[ent.GithubAsset], error) {
	return p.Exec(r.Context(), s.db.GithubRelease.Query().Where(githubrelease.ID(githubreleaseID)).QueryAssets())
}

// ListGithubRepositories maps to "GET /github-repositories".
func (s *Server) ListGithubRepositories(r *http.Request, p *ListGithubRepositoryParams) (*PagedResponse[ent.GithubRepository], error) {
	return p.Exec(r.Context(), s.db.GithubRepository.Query())
}

// GetGithubRepository maps to "GET /github-repositories/{id}".
func (s *Server) GetGithubRepository(r *http.Request, githubrepositoryID int) (*ent.GithubRepository, error) {
	return EagerLoadGithubRepository(s.db.GithubRepository.Query().Where(githubrepository.ID(githubrepositoryID))).Only(r.Context())
}

// ListGithubRepositoryLabels maps to "GET /github-repositories/{id}/labels".
func (s *Server) ListGithubRepositoryLabels(r *http.Request, githubrepositoryID int, p *ListLabelParams) (*PagedResponse[ent.Label], error) {
	return p.Exec(r.Context(), s.db.GithubRepository.Query().Where(githubrepository.ID(githubrepositoryID)).QueryLabels())
}

// ListGithubRepositoryReleases maps to "GET /github-repositories/{id}/releases".
func (s *Server) ListGithubRepositoryReleases(r *http.Request, githubrepositoryID int, p *ListGithubReleaseParams) (*PagedResponse[ent.GithubRelease], error) {
	return p.Exec(r.Context(), s.db.GithubRepository.Query().Where(githubrepository.ID(githubrepositoryID)).QueryReleases())
}

// ListLabels maps to "GET /labels".
func (s *Server) ListLabels(r *http.Request, p *ListLabelParams) (*PagedResponse[ent.Label], error) {
	return p.Exec(r.Context(), s.db.Label.Query())
}

// GetLabel maps to "GET /labels/{id}".
func (s *Server) GetLabel(r *http.Request, labelID int) (*ent.Label, error) {
	return EagerLoadLabel(s.db.Label.Query().Where(label.ID(labelID))).Only(r.Context())
}

// ListLabelPosts maps to "GET /labels/{id}/posts".
func (s *Server) ListLabelPosts(r *http.Request, labelID int, p *ListPostParams) (*PagedResponse[ent.Post], error) {
	return p.Exec(r.Context(), s.db.Label.Query().Where(label.ID(labelID)).QueryPosts())
}

// ListLabelGithubRepositories maps to "GET /labels/{id}/github-repositories".
func (s *Server) ListLabelGithubRepositories(r *http.Request, labelID int, p *ListGithubRepositoryParams) (*PagedResponse[ent.GithubRepository], error) {
	return p.Exec(r.Context(), s.db.Label.Query().Where(label.ID(labelID)).QueryGithubRepositories())
}

// CreateLabel maps to "POST /labels".
func (s *Server) CreateLabel(r *http.Request, p *CreateLabelParams) (*ent.Label, error) {
	return p.Exec(r.Context(), s.db.Label.Create(), s.db.Label.Query())
}

// UpdateLabel maps to "PATCH /labels/{id}".
func (s *Server) UpdateLabel(r *http.Request, labelID int, p *UpdateLabelParams) (*ent.Label, error) {
	return p.Exec(r.Context(), s.db.Label.UpdateOneID(labelID), s.db.Label.Query())
}

// DeleteLabel maps to "DELETE /labels/{id}".
func (s *Server) DeleteLabel(r *http.Request, labelID int) (*struct{}, error) {
	return nil, s.db.Label.DeleteOneID(labelID).Exec(r.Context())
}

// ListPosts maps to "GET /posts".
func (s *Server) ListPosts(r *http.Request, p *ListPostParams) (*PagedResponse[ent.Post], error) {
	return p.Exec(r.Context(), s.db.Post.Query())
}

// GetPost maps to "GET /posts/{id}".
func (s *Server) GetPost(r *http.Request, postID int) (*ent.Post, error) {
	return EagerLoadPost(s.db.Post.Query().Where(post.ID(postID))).Only(r.Context())
}

// GetPostAuthor maps to "GET /posts/{id}/author".
func (s *Server) GetPostAuthor(r *http.Request, postID int) (*ent.User, error) {
	return EagerLoadUser(s.db.Post.Query().Where(post.ID(postID)).QueryAuthor()).Only(r.Context())
}

// ListPostLabels maps to "GET /posts/{id}/labels".
func (s *Server) ListPostLabels(r *http.Request, postID int, p *ListLabelParams) (*PagedResponse[ent.Label], error) {
	return p.Exec(r.Context(), s.db.Post.Query().Where(post.ID(postID)).QueryLabels())
}

// CreatePost maps to "POST /posts".
func (s *Server) CreatePost(r *http.Request, p *CreatePostParams) (*ent.Post, error) {
	return p.Exec(r.Context(), s.db.Post.Create(), s.db.Post.Query())
}

// UpdatePost maps to "PATCH /posts/{id}".
func (s *Server) UpdatePost(r *http.Request, postID int, p *UpdatePostParams) (*ent.Post, error) {
	return p.Exec(r.Context(), s.db.Post.UpdateOneID(postID), s.db.Post.Query())
}

// DeletePost maps to "DELETE /posts/{id}".
func (s *Server) DeletePost(r *http.Request, postID int) (*struct{}, error) {
	return nil, s.db.Post.DeleteOneID(postID).Exec(r.Context())
}

// ListUsers maps to "GET /users".
func (s *Server) ListUsers(r *http.Request, p *ListUserParams) (*PagedResponse[ent.User], error) {
	return p.Exec(r.Context(), s.db.User.Query())
}

// GetUser maps to "GET /users/{id}".
func (s *Server) GetUser(r *http.Request, userID int) (*ent.User, error) {
	return EagerLoadUser(s.db.User.Query().Where(user.ID(userID))).Only(r.Context())
}

// ListUserPosts maps to "GET /users/{id}/posts".
func (s *Server) ListUserPosts(r *http.Request, userID int, p *ListPostParams) (*PagedResponse[ent.Post], error) {
	return p.Exec(r.Context(), s.db.User.Query().Where(user.ID(userID)).QueryPosts())
}
