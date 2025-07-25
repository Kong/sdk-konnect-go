// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkkonnectgo

// Generated from OpenAPI doc version 3.0.22 and generator version 2.638.5

import (
	"context"
	"fmt"
	"github.com/Kong/sdk-konnect-go/internal/config"
	"github.com/Kong/sdk-konnect-go/internal/hooks"
	"github.com/Kong/sdk-konnect-go/internal/utils"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/retry"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	"https://global.api.konghq.com",
	"https://us.api.konghq.com",
	"https://eu.api.konghq.com",
	"https://au.api.konghq.com",
}

// HTTPClient provides an interface for supplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

// Pointer provides a helper function to return a pointer to a type
func Pointer[T any](v T) *T { return &v }

// SDK - Konnect API - Go SDK: The Konnect platform API
//
// https://developer.konghq.com - Documentation for Kong Gateway and its APIs
type SDK struct {
	SDKVersion string
	// Operations related to notifications
	Notifications *Notifications
	// Application Auth Strategies are sets of plugin configurations that represent how the gateway will perform authentication and authorization for a Product Version.
	// Called “Auth Strategy” for short in the context of portals/applications.
	// The plugins are synced to any Gateway Service that is currently linked or becomes linked to the Product Version.
	//
	AppAuthStrategies *AppAuthStrategies
	CloudGateways     *CloudGateways
	ControlPlanes     *ControlPlanes
	// Config Stores
	ConfigStores *ConfigStores
	// Config Store Secrets
	ConfigStoreSecrets   *ConfigStoreSecrets
	ACLs                 *ACLs
	BasicAuthCredentials *BasicAuthCredentials
	// A CA certificate object represents a trusted certificate authority.
	// These objects are used by Kong Gateway to verify the validity of a client or server certificate.
	CACertificates *CACertificates
	// A certificate object represents a public certificate, and can be optionally paired with the corresponding private key. These objects are used by Kong Gateway to handle SSL/TLS termination for encrypted requests, or for use as a trusted CA store when validating peer certificate of client/service.
	// <br><br>
	// Certificates are optionally associated with SNI objects to tie a cert/key pair to one or more hostnames.
	// <br><br>
	// If intermediate certificates are required in addition to the main certificate, they should be concatenated together into one string.
	//
	Certificates *Certificates
	// An SNI object represents a many-to-one mapping of hostnames to a certificate.
	// <br><br>
	// A certificate object can have many hostnames associated with it. When Kong Gateway receives an SSL request, it uses the SNI field in the Client Hello to look up the certificate object based on the SNI associated with the certificate.
	SNIs *SNIs
	// Consumer groups enable the organization and categorization of consumers (users or applications) within an API ecosystem.
	// By grouping consumers together, you eliminate the need to manage them individually, providing a scalable, efficient approach to managing configurations.
	ConsumerGroups *ConsumerGroups
	// A plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. Plugins let you add functionality to services that run behind a Kong Gateway instance, like authentication or rate limiting.
	// You can find more information about available plugins and which values each plugin accepts at the [Plugin Hub](https://developer.konghq.com/plugins/).
	// <br><br>
	// When adding a plugin configuration to a service, the plugin will run on every request made by a client to that service. If a plugin needs to be tuned to different values for some specific consumers, you can do so by creating a separate plugin instance that specifies both the service and the consumer, through the service and consumer fields.
	Plugins *Plugins
	// The consumer object represents a consumer - or a user - of a service.
	// You can either rely on Kong Gateway as the primary datastore, or you can map the consumer list with your database to keep consistency between Kong Gateway and your existing primary datastore.
	//
	Consumers           *Consumers
	HMACAuthCredentials *HMACAuthCredentials
	JWTs                *JWTs
	APIKeys             *APIKeys
	MTLSAuthCredentials *MTLSAuthCredentials
	CustomPlugins       *CustomPlugins
	DegraphqlRoutes     *DegraphqlRoutes
	// A JSON Web key set. Key sets are the preferred way to expose keys to plugins because they tell the plugin where to look for keys or have a scoping mechanism to restrict plugins to specific keys.
	//
	KeySets *KeySets
	// A key object holds a representation of asymmetric keys in various formats. When Kong Gateway or a Kong plugin requires a specific public or private key to perform certain operations, it can use this entity.
	//
	Keys *Keys
	// Some entities in Kong Gateway share common configuration settings that often need to be repeated. For example, multiple plugins that connect to Redis may require the same connection settings. Without Partials, you would need to replicate this configuration across all plugins. If the settings change, you would need to update each plugin individually.
	Partials     *Partials
	PartialLinks *PartialLinks
	// Custom Plugin Schemas
	CustomPluginSchemas *CustomPluginSchemas
	// Route entities define rules to match client requests. Each route is associated with a service, and a service may have multiple routes associated to it. Every request matching a given route will be proxied to the associated service. You need at least one matching rule that applies to the protocol being matched by the route.
	// <br><br>
	// The combination of routes and services, and the separation of concerns between them, offers a powerful routing mechanism with which it is possible to define fine-grained entrypoints in Kong Gateway leading to different upstream services of your infrastructure.
	// <br><br>
	// Depending on the protocol, one of the following attributes must be set:
	// <br>
	//
	// - `http`: At least one of `methods`, `hosts`, `headers`, or `paths`
	// - `https`: At least one of `methods`, `hosts`, `headers`, `paths`, or `snis`
	// - `tcp`: At least one of `sources` or `destinations`
	// - `tls`: at least one of `sources`, `destinations`, or `snis`
	// - `tls_passthrough`: set `snis`
	// - `grpc`: At least one of `hosts`, `headers`, or `paths`
	// - `grpcs`: At least one of `hosts`, `headers`, `paths`, or `snis`
	// - `ws`: At least one of `hosts`, `headers`, or `paths`
	// - `wss`: At least one of `hosts`, `headers`, `paths`, or `snis`
	//
	//
	//
	//   <br>
	//   A route can't have both `tls` and `tls_passthrough` protocols at same time.
	//   <br><br>
	//   Learn more about the router:
	// - [Configure routes using expressions](https://developer.konghq.com/gateway/routing/expressions/)
	//
	Routes  *Routes
	Schemas *Schemas
	// Service entities are abstractions of your microservice interfaces or formal APIs. For example, a service could be a data transformation microservice or a billing API.
	// <br><br>
	// The main attribute of a service is the destination URL for proxying traffic. This URL can be set as a single string or by specifying its protocol, host, port and path individually.
	// <br><br>
	// Services are associated to routes, and a single service can have many routes associated with it. Routes are entrypoints in Kong Gateway which define rules to match client requests. Once a route is matched, Kong Gateway proxies the request to its associated service. See the [Route documentation](https://developer.konghq.com/gateway/entities/route/) for a detailed explanation of how Kong proxies traffic.
	// <br><br>
	// Services can be both [tagged and filtered by tags](https://developer.konghq.com/admin-api/).
	//
	Services *Services
	// The upstream object represents a virtual hostname and can be used to load balance incoming requests over multiple services (targets).
	// <br><br>
	// An upstream also includes a [health checker](https://developer.konghq.com/gateway/traffic-control/health-checks-circuit-breakers/), which can enable and disable targets based on their ability or inability to serve requests.
	// The configuration for the health checker is stored in the upstream object, and applies to all of its targets.
	Upstreams *Upstreams
	// A target is an IP address or hostname with a port that identifies an instance of a backend service. Every upstream can have many targets, and the targets can be dynamically added, modified, or deleted. Changes take effect on the fly.
	// <br><br>
	// To disable a target, post a new one with `weight=0`, or use the `DELETE` method to accomplish the same.
	//
	Targets *Targets
	// Vault objects are used to configure different vault connectors for [managing secrets](https://developer.konghq.com/gateway/secrets-management/).
	// Configuring a vault lets you reference secrets from other entities.
	// This allows for a proper separation of secrets and configuration and prevents secret sprawl.
	// <br><br>
	// For example, you could store a certificate and a key in a vault, then reference them from a certificate entity. This way, the certificate and key are not stored in the entity directly and are more secure.
	// <br><br>
	// Secrets rotation can be managed using [TTLs](https://developer.konghq.com/gateway/entities/vault/).
	//
	Vaults *Vaults
	// DP Certificates
	DPCertificates *DPCertificates
	// DP Nodes
	DPNodes            *DPNodes
	ControlPlaneGroups *ControlPlaneGroups
	// Dynamic Client Registration Providers are configurations representing an external Identity Provider whose clients (i.e. Applications) Konnect will be authorized to manage.
	// For instance, they will be able to perform dynamic client registration (DCR) with the provider.
	// The DCR provider provides credentials to each DCR-enabled application in Konnect that can be used to access Product Versions that the app is registered for.
	//
	DCRProviders      *DCRProviders
	APIAttributes     *APIAttributes
	APIImplementation *APIImplementation
	APIPublication    *APIPublication
	API               *API
	APIDocumentation  *APIDocumentation
	APISpecification  *APISpecification
	APIVersion        *APIVersion
	// APIs related to Konnect Developer Portal Applications.
	Applications          *Applications
	Authentication        *Authentication
	AuthSettings          *AuthSettings
	Invites               *Invites
	ImpersonationSettings *ImpersonationSettings
	Me                    *Me
	// APIs related to Konnect Developer Portal developer team roles.
	PortalTeamRoles *PortalTeamRoles
	// APIs related to configuration of Konnect Developer Portals.
	Portals *Portals
	// APIs related to Konnect Developer Portal Application Registrations.
	ApplicationRegistrations *ApplicationRegistrations
	// APIs for managing static assets for Konnect Developer Portals.
	Assets          *Assets
	PortalAuditLogs *PortalAuditLogs
	// APIs related to configuration of Konnect Developer Portal auth settings.
	PortalAuthSettings *PortalAuthSettings
	// APIs related to configuration of Konnect Developer Portals custom domains.
	PortalCustomDomains *PortalCustomDomains
	// APIs related to customization of Konnect Developer Portals.
	PortalCustomization *PortalCustomization
	// APIs related to Konnect Developer Portal Custom Pages.
	Pages *Pages
	// APIs related to Konnect Developer Portal developers.
	PortalDevelopers *PortalDevelopers
	// APIs related to Konnect Developer Portal developer team membership.
	PortalTeamMembership *PortalTeamMembership
	// APIs related to Konnect Developer Portal Emails.
	PortalEmails *PortalEmails
	// APIs related to Konnect Developer Portal Custom Snippets.
	Snippets *Snippets
	// APIs related to configuration of Konnect Developer Portal developer teams.
	PortalTeams                  *PortalTeams
	Roles                        *Roles
	SystemAccounts               *SystemAccounts
	SystemAccountsAccessTokens   *SystemAccountsAccessTokens
	SystemAccountsRoles          *SystemAccountsRoles
	SystemAccountsTeamMembership *SystemAccountsTeamMembership
	Teams                        *Teams
	TeamMembership               *TeamMembership
	Users                        *Users

	sdkConfiguration config.SDKConfiguration
	hooks            *hooks.Hooks
}

type SDKOption func(*SDK)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *SDK) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *SDK) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.Client = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security components.Security) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.Security = utils.AsSecuritySource(security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (components.Security, error)) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig retry.Config) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// WithTimeout Optional request timeout applied to each operation
func WithTimeout(timeout time.Duration) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.Timeout = &timeout
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *SDK {
	sdk := &SDK{
		SDKVersion: "0.7.1",
		sdkConfiguration: config.SDKConfiguration{
			UserAgent:  "speakeasy-sdk/go 0.7.1 2.638.5 3.0.22 github.com/Kong/sdk-konnect-go",
			ServerList: ServerList,
		},
		hooks: hooks.New(),
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.Client == nil {
		sdk.sdkConfiguration.Client = &http.Client{Timeout: 60 * time.Second}
	}

	currentServerURL, _ := sdk.sdkConfiguration.GetServerDetails()
	serverURL := currentServerURL
	serverURL, sdk.sdkConfiguration.Client = sdk.hooks.SDKInit(currentServerURL, sdk.sdkConfiguration.Client)
	if currentServerURL != serverURL {
		sdk.sdkConfiguration.ServerURL = serverURL
	}

	sdk.Notifications = newNotifications(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.AppAuthStrategies = newAppAuthStrategies(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.CloudGateways = newCloudGateways(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.ControlPlanes = newControlPlanes(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.ConfigStores = newConfigStores(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.ConfigStoreSecrets = newConfigStoreSecrets(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.ACLs = newACLs(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.BasicAuthCredentials = newBasicAuthCredentials(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.CACertificates = newCACertificates(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Certificates = newCertificates(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.SNIs = newSNIs(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.ConsumerGroups = newConsumerGroups(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Plugins = newPlugins(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Consumers = newConsumers(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.HMACAuthCredentials = newHMACAuthCredentials(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.JWTs = newJWTs(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.APIKeys = newAPIKeys(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.MTLSAuthCredentials = newMTLSAuthCredentials(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.CustomPlugins = newCustomPlugins(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.DegraphqlRoutes = newDegraphqlRoutes(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.KeySets = newKeySets(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Keys = newKeys(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Partials = newPartials(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.PartialLinks = newPartialLinks(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.CustomPluginSchemas = newCustomPluginSchemas(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Routes = newRoutes(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Schemas = newSchemas(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Services = newServices(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Upstreams = newUpstreams(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Targets = newTargets(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Vaults = newVaults(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.DPCertificates = newDPCertificates(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.DPNodes = newDPNodes(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.ControlPlaneGroups = newControlPlaneGroups(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.DCRProviders = newDCRProviders(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.APIAttributes = newAPIAttributes(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.APIImplementation = newAPIImplementation(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.APIPublication = newAPIPublication(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.API = newAPI(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.APIDocumentation = newAPIDocumentation(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.APISpecification = newAPISpecification(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.APIVersion = newAPIVersion(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Applications = newApplications(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Authentication = newAuthentication(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.AuthSettings = newAuthSettings(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Invites = newInvites(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.ImpersonationSettings = newImpersonationSettings(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Me = newMe(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.PortalTeamRoles = newPortalTeamRoles(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Portals = newPortals(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.ApplicationRegistrations = newApplicationRegistrations(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Assets = newAssets(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.PortalAuditLogs = newPortalAuditLogs(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.PortalAuthSettings = newPortalAuthSettings(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.PortalCustomDomains = newPortalCustomDomains(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.PortalCustomization = newPortalCustomization(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Pages = newPages(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.PortalDevelopers = newPortalDevelopers(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.PortalTeamMembership = newPortalTeamMembership(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.PortalEmails = newPortalEmails(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Snippets = newSnippets(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.PortalTeams = newPortalTeams(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Roles = newRoles(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.SystemAccounts = newSystemAccounts(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.SystemAccountsAccessTokens = newSystemAccountsAccessTokens(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.SystemAccountsRoles = newSystemAccountsRoles(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.SystemAccountsTeamMembership = newSystemAccountsTeamMembership(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Teams = newTeams(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.TeamMembership = newTeamMembership(sdk, sdk.sdkConfiguration, sdk.hooks)
	sdk.Users = newUsers(sdk, sdk.sdkConfiguration, sdk.hooks)

	return sdk
}
