// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/Kong/sdk-konnect-go/internal/utils"
)

// UpstreamAlgorithm - Which load balancing algorithm to use.
type UpstreamAlgorithm string

const (
	UpstreamAlgorithmConsistentHashing UpstreamAlgorithm = "consistent-hashing"
	UpstreamAlgorithmLeastConnections  UpstreamAlgorithm = "least-connections"
	UpstreamAlgorithmRoundRobin        UpstreamAlgorithm = "round-robin"
	UpstreamAlgorithmLatency           UpstreamAlgorithm = "latency"
	UpstreamAlgorithmStickySessions    UpstreamAlgorithm = "sticky-sessions"
)

func (e UpstreamAlgorithm) ToPointer() *UpstreamAlgorithm {
	return &e
}
func (e *UpstreamAlgorithm) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "consistent-hashing":
		fallthrough
	case "least-connections":
		fallthrough
	case "round-robin":
		fallthrough
	case "latency":
		fallthrough
	case "sticky-sessions":
		*e = UpstreamAlgorithm(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpstreamAlgorithm: %v", v)
	}
}

// UpstreamClientCertificate - If set, the certificate to be used as client certificate while TLS handshaking to the upstream server.
type UpstreamClientCertificate struct {
	ID *string `json:"id,omitempty"`
}

func (o *UpstreamClientCertificate) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// HashFallback - What to use as hashing input if the primary `hash_on` does not return a hash (eg. header is missing, or no Consumer identified). Not available if `hash_on` is set to `cookie`.
type HashFallback string

const (
	HashFallbackNone       HashFallback = "none"
	HashFallbackConsumer   HashFallback = "consumer"
	HashFallbackIP         HashFallback = "ip"
	HashFallbackHeader     HashFallback = "header"
	HashFallbackCookie     HashFallback = "cookie"
	HashFallbackPath       HashFallback = "path"
	HashFallbackQueryArg   HashFallback = "query_arg"
	HashFallbackURICapture HashFallback = "uri_capture"
)

func (e HashFallback) ToPointer() *HashFallback {
	return &e
}
func (e *HashFallback) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "consumer":
		fallthrough
	case "ip":
		fallthrough
	case "header":
		fallthrough
	case "cookie":
		fallthrough
	case "path":
		fallthrough
	case "query_arg":
		fallthrough
	case "uri_capture":
		*e = HashFallback(v)
		return nil
	default:
		return fmt.Errorf("invalid value for HashFallback: %v", v)
	}
}

// HashOn - What to use as hashing input. Using `none` results in a weighted-round-robin scheme with no hashing.
type HashOn string

const (
	HashOnNone       HashOn = "none"
	HashOnConsumer   HashOn = "consumer"
	HashOnIP         HashOn = "ip"
	HashOnHeader     HashOn = "header"
	HashOnCookie     HashOn = "cookie"
	HashOnPath       HashOn = "path"
	HashOnQueryArg   HashOn = "query_arg"
	HashOnURICapture HashOn = "uri_capture"
)

func (e HashOn) ToPointer() *HashOn {
	return &e
}
func (e *HashOn) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "consumer":
		fallthrough
	case "ip":
		fallthrough
	case "header":
		fallthrough
	case "cookie":
		fallthrough
	case "path":
		fallthrough
	case "query_arg":
		fallthrough
	case "uri_capture":
		*e = HashOn(v)
		return nil
	default:
		return fmt.Errorf("invalid value for HashOn: %v", v)
	}
}

type Healthy struct {
	HTTPStatuses []int64  `json:"http_statuses,omitempty"`
	Interval     *float64 `default:"0" json:"interval"`
	Successes    *int64   `default:"0" json:"successes"`
}

func (h Healthy) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(h, "", false)
}

func (h *Healthy) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &h, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Healthy) GetHTTPStatuses() []int64 {
	if o == nil {
		return nil
	}
	return o.HTTPStatuses
}

func (o *Healthy) GetInterval() *float64 {
	if o == nil {
		return nil
	}
	return o.Interval
}

func (o *Healthy) GetSuccesses() *int64 {
	if o == nil {
		return nil
	}
	return o.Successes
}

type UpstreamType string

const (
	UpstreamTypeTCP   UpstreamType = "tcp"
	UpstreamTypeHTTP  UpstreamType = "http"
	UpstreamTypeHTTPS UpstreamType = "https"
	UpstreamTypeGrpc  UpstreamType = "grpc"
	UpstreamTypeGrpcs UpstreamType = "grpcs"
)

func (e UpstreamType) ToPointer() *UpstreamType {
	return &e
}
func (e *UpstreamType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "tcp":
		fallthrough
	case "http":
		fallthrough
	case "https":
		fallthrough
	case "grpc":
		fallthrough
	case "grpcs":
		*e = UpstreamType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpstreamType: %v", v)
	}
}

type Unhealthy struct {
	HTTPFailures *int64   `default:"0" json:"http_failures"`
	HTTPStatuses []int64  `json:"http_statuses,omitempty"`
	Interval     *float64 `default:"0" json:"interval"`
	TCPFailures  *int64   `default:"0" json:"tcp_failures"`
	Timeouts     *int64   `default:"0" json:"timeouts"`
}

func (u Unhealthy) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *Unhealthy) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Unhealthy) GetHTTPFailures() *int64 {
	if o == nil {
		return nil
	}
	return o.HTTPFailures
}

func (o *Unhealthy) GetHTTPStatuses() []int64 {
	if o == nil {
		return nil
	}
	return o.HTTPStatuses
}

func (o *Unhealthy) GetInterval() *float64 {
	if o == nil {
		return nil
	}
	return o.Interval
}

func (o *Unhealthy) GetTCPFailures() *int64 {
	if o == nil {
		return nil
	}
	return o.TCPFailures
}

func (o *Unhealthy) GetTimeouts() *int64 {
	if o == nil {
		return nil
	}
	return o.Timeouts
}

type Active struct {
	Concurrency            *int64            `default:"10" json:"concurrency"`
	Headers                map[string]string `json:"headers,omitempty"`
	Healthy                *Healthy          `json:"healthy,omitempty"`
	HTTPPath               *string           `default:"/" json:"http_path"`
	HTTPSSni               *string           `json:"https_sni,omitempty"`
	HTTPSVerifyCertificate *bool             `default:"true" json:"https_verify_certificate"`
	Timeout                *float64          `default:"1" json:"timeout"`
	Type                   *UpstreamType     `default:"http" json:"type"`
	Unhealthy              *Unhealthy        `json:"unhealthy,omitempty"`
}

func (a Active) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *Active) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Active) GetConcurrency() *int64 {
	if o == nil {
		return nil
	}
	return o.Concurrency
}

func (o *Active) GetHeaders() map[string]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *Active) GetHealthy() *Healthy {
	if o == nil {
		return nil
	}
	return o.Healthy
}

func (o *Active) GetHTTPPath() *string {
	if o == nil {
		return nil
	}
	return o.HTTPPath
}

func (o *Active) GetHTTPSSni() *string {
	if o == nil {
		return nil
	}
	return o.HTTPSSni
}

func (o *Active) GetHTTPSVerifyCertificate() *bool {
	if o == nil {
		return nil
	}
	return o.HTTPSVerifyCertificate
}

func (o *Active) GetTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.Timeout
}

func (o *Active) GetType() *UpstreamType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *Active) GetUnhealthy() *Unhealthy {
	if o == nil {
		return nil
	}
	return o.Unhealthy
}

type UpstreamHealthy struct {
	HTTPStatuses []int64 `json:"http_statuses,omitempty"`
	Successes    *int64  `default:"0" json:"successes"`
}

func (u UpstreamHealthy) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpstreamHealthy) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpstreamHealthy) GetHTTPStatuses() []int64 {
	if o == nil {
		return nil
	}
	return o.HTTPStatuses
}

func (o *UpstreamHealthy) GetSuccesses() *int64 {
	if o == nil {
		return nil
	}
	return o.Successes
}

type UpstreamHealthchecksType string

const (
	UpstreamHealthchecksTypeTCP   UpstreamHealthchecksType = "tcp"
	UpstreamHealthchecksTypeHTTP  UpstreamHealthchecksType = "http"
	UpstreamHealthchecksTypeHTTPS UpstreamHealthchecksType = "https"
	UpstreamHealthchecksTypeGrpc  UpstreamHealthchecksType = "grpc"
	UpstreamHealthchecksTypeGrpcs UpstreamHealthchecksType = "grpcs"
)

func (e UpstreamHealthchecksType) ToPointer() *UpstreamHealthchecksType {
	return &e
}
func (e *UpstreamHealthchecksType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "tcp":
		fallthrough
	case "http":
		fallthrough
	case "https":
		fallthrough
	case "grpc":
		fallthrough
	case "grpcs":
		*e = UpstreamHealthchecksType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpstreamHealthchecksType: %v", v)
	}
}

type UpstreamUnhealthy struct {
	HTTPFailures *int64  `default:"0" json:"http_failures"`
	HTTPStatuses []int64 `json:"http_statuses,omitempty"`
	TCPFailures  *int64  `default:"0" json:"tcp_failures"`
	Timeouts     *int64  `default:"0" json:"timeouts"`
}

func (u UpstreamUnhealthy) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpstreamUnhealthy) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpstreamUnhealthy) GetHTTPFailures() *int64 {
	if o == nil {
		return nil
	}
	return o.HTTPFailures
}

func (o *UpstreamUnhealthy) GetHTTPStatuses() []int64 {
	if o == nil {
		return nil
	}
	return o.HTTPStatuses
}

func (o *UpstreamUnhealthy) GetTCPFailures() *int64 {
	if o == nil {
		return nil
	}
	return o.TCPFailures
}

func (o *UpstreamUnhealthy) GetTimeouts() *int64 {
	if o == nil {
		return nil
	}
	return o.Timeouts
}

type Passive struct {
	Healthy   *UpstreamHealthy          `json:"healthy,omitempty"`
	Type      *UpstreamHealthchecksType `default:"http" json:"type"`
	Unhealthy *UpstreamUnhealthy        `json:"unhealthy,omitempty"`
}

func (p Passive) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *Passive) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Passive) GetHealthy() *UpstreamHealthy {
	if o == nil {
		return nil
	}
	return o.Healthy
}

func (o *Passive) GetType() *UpstreamHealthchecksType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *Passive) GetUnhealthy() *UpstreamUnhealthy {
	if o == nil {
		return nil
	}
	return o.Unhealthy
}

type Healthchecks struct {
	Active    *Active  `json:"active,omitempty"`
	Passive   *Passive `json:"passive,omitempty"`
	Threshold *float64 `default:"0" json:"threshold"`
}

func (h Healthchecks) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(h, "", false)
}

func (h *Healthchecks) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &h, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Healthchecks) GetActive() *Active {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *Healthchecks) GetPassive() *Passive {
	if o == nil {
		return nil
	}
	return o.Passive
}

func (o *Healthchecks) GetThreshold() *float64 {
	if o == nil {
		return nil
	}
	return o.Threshold
}

// Upstream - The upstream object represents a virtual hostname and can be used to loadbalance incoming requests over multiple services (targets). So for example an upstream named `service.v1.xyz` for a Service object whose `host` is `service.v1.xyz`. Requests for this Service would be proxied to the targets defined within the upstream. An upstream also includes a [health checker][healthchecks], which is able to enable and disable targets based on their ability or inability to serve requests. The configuration for the health checker is stored in the upstream object, and applies to all of its targets.
type Upstream struct {
	// Which load balancing algorithm to use.
	Algorithm *UpstreamAlgorithm `default:"round-robin" json:"algorithm"`
	// If set, the certificate to be used as client certificate while TLS handshaking to the upstream server.
	ClientCertificate *UpstreamClientCertificate `json:"client_certificate,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// What to use as hashing input if the primary `hash_on` does not return a hash (eg. header is missing, or no Consumer identified). Not available if `hash_on` is set to `cookie`.
	HashFallback *HashFallback `default:"none" json:"hash_fallback"`
	// The header name to take the value from as hash input. Only required when `hash_fallback` is set to `header`.
	HashFallbackHeader *string `json:"hash_fallback_header,omitempty"`
	// The name of the query string argument to take the value from as hash input. Only required when `hash_fallback` is set to `query_arg`.
	HashFallbackQueryArg *string `json:"hash_fallback_query_arg,omitempty"`
	// The name of the route URI capture to take the value from as hash input. Only required when `hash_fallback` is set to `uri_capture`.
	HashFallbackURICapture *string `json:"hash_fallback_uri_capture,omitempty"`
	// What to use as hashing input. Using `none` results in a weighted-round-robin scheme with no hashing.
	HashOn *HashOn `default:"none" json:"hash_on"`
	// The cookie name to take the value from as hash input. Only required when `hash_on` or `hash_fallback` is set to `cookie`. If the specified cookie is not in the request, Kong will generate a value and set the cookie in the response.
	HashOnCookie *string `json:"hash_on_cookie,omitempty"`
	// The cookie path to set in the response headers. Only required when `hash_on` or `hash_fallback` is set to `cookie`.
	HashOnCookiePath *string `default:"/" json:"hash_on_cookie_path"`
	// The header name to take the value from as hash input. Only required when `hash_on` is set to `header`.
	HashOnHeader *string `json:"hash_on_header,omitempty"`
	// The name of the query string argument to take the value from as hash input. Only required when `hash_on` is set to `query_arg`.
	HashOnQueryArg *string `json:"hash_on_query_arg,omitempty"`
	// The name of the route URI capture to take the value from as hash input. Only required when `hash_on` is set to `uri_capture`.
	HashOnURICapture *string       `json:"hash_on_uri_capture,omitempty"`
	Healthchecks     *Healthchecks `json:"healthchecks,omitempty"`
	// The hostname to be used as `Host` header when proxying requests through Kong.
	HostHeader *string `json:"host_header,omitempty"`
	ID         *string `json:"id,omitempty"`
	// This is a hostname, which must be equal to the `host` of a Service.
	Name string `json:"name"`
	// The number of slots in the load balancer algorithm. If `algorithm` is set to `round-robin`, this setting determines the maximum number of slots. If `algorithm` is set to `consistent-hashing`, this setting determines the actual number of slots in the algorithm. Accepts an integer in the range `10`-`65536`.
	Slots                    *int64  `default:"10000" json:"slots"`
	StickySessionsCookie     *string `json:"sticky_sessions_cookie,omitempty"`
	StickySessionsCookiePath *string `default:"/" json:"sticky_sessions_cookie_path"`
	// An optional set of strings associated with the Upstream for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// If set, the balancer will use SRV hostname(if DNS Answer has SRV record) as the proxy upstream `Host`.
	UseSrvName *bool `default:"false" json:"use_srv_name"`
}

func (u Upstream) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *Upstream) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Upstream) GetAlgorithm() *UpstreamAlgorithm {
	if o == nil {
		return nil
	}
	return o.Algorithm
}

func (o *Upstream) GetClientCertificate() *UpstreamClientCertificate {
	if o == nil {
		return nil
	}
	return o.ClientCertificate
}

func (o *Upstream) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Upstream) GetHashFallback() *HashFallback {
	if o == nil {
		return nil
	}
	return o.HashFallback
}

func (o *Upstream) GetHashFallbackHeader() *string {
	if o == nil {
		return nil
	}
	return o.HashFallbackHeader
}

func (o *Upstream) GetHashFallbackQueryArg() *string {
	if o == nil {
		return nil
	}
	return o.HashFallbackQueryArg
}

func (o *Upstream) GetHashFallbackURICapture() *string {
	if o == nil {
		return nil
	}
	return o.HashFallbackURICapture
}

func (o *Upstream) GetHashOn() *HashOn {
	if o == nil {
		return nil
	}
	return o.HashOn
}

func (o *Upstream) GetHashOnCookie() *string {
	if o == nil {
		return nil
	}
	return o.HashOnCookie
}

func (o *Upstream) GetHashOnCookiePath() *string {
	if o == nil {
		return nil
	}
	return o.HashOnCookiePath
}

func (o *Upstream) GetHashOnHeader() *string {
	if o == nil {
		return nil
	}
	return o.HashOnHeader
}

func (o *Upstream) GetHashOnQueryArg() *string {
	if o == nil {
		return nil
	}
	return o.HashOnQueryArg
}

func (o *Upstream) GetHashOnURICapture() *string {
	if o == nil {
		return nil
	}
	return o.HashOnURICapture
}

func (o *Upstream) GetHealthchecks() *Healthchecks {
	if o == nil {
		return nil
	}
	return o.Healthchecks
}

func (o *Upstream) GetHostHeader() *string {
	if o == nil {
		return nil
	}
	return o.HostHeader
}

func (o *Upstream) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Upstream) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Upstream) GetSlots() *int64 {
	if o == nil {
		return nil
	}
	return o.Slots
}

func (o *Upstream) GetStickySessionsCookie() *string {
	if o == nil {
		return nil
	}
	return o.StickySessionsCookie
}

func (o *Upstream) GetStickySessionsCookiePath() *string {
	if o == nil {
		return nil
	}
	return o.StickySessionsCookiePath
}

func (o *Upstream) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *Upstream) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *Upstream) GetUseSrvName() *bool {
	if o == nil {
		return nil
	}
	return o.UseSrvName
}
