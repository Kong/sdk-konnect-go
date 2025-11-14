## How to use this SDK

> **READ THIS SECTION!**

This repository is a private version of <https://github.com/Kong/sdk-konnect-go>.

It mostly serves a purpose for internal development and testing of the SDK as the
underlying APIs evolve towards their stable versions.

Typical usage of this SDK is to import `github.com/Kong/sdk-konnect-go` in your
Go project and add a replace directorive in your `go.mod` file to point to this
private repository.

For example, having the following Go file using the SDK:

```go
package main

import (
	"context"
	"fmt"
	"os"

	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
)

func main() {
	s := sdkkonnectgo.New(
		sdkkonnectgo.WithSecurity(components.Security{
			PersonalAccessToken: sdkkonnectgo.Pointer(os.Getenv("KONNECT_TOKEN")),
		}),
	)
	_ = s
}
```

We use the replace directive in `go.mod` as follows:

```
# Replace `main` with your desired branch or tag
go mod edit -replace=github.com/Kong/sdk-konnect-go=github.com/Kong/sdk-konnect-go-internal@main
go mod tidy
```

> **NOTE**: Make sure that you have `GOPRIVATE` set to `github.com/Kong/sdk-konnect-go-internal`
> or `github.com/Kong`. You might need to tweak your `git` authentication settings
> accordingly to be able to access the private repository.

This way allows users of the SDK to easily switch between the public and private
versions of the SDK as needed without the need for changing any code in their projects.

<!-- Start Summary [summary] -->
## Summary

Konnect API - Go SDK: The Konnect platform API

For more information about the API: [Documentation for Kong Gateway and its APIs](https://developer.konghq.com)
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
  * [How to use this SDK](#how-to-use-this-sdk)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Pagination](#pagination)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/Kong/sdk-konnect-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := sdkkonnectgo.New(
		sdkkonnectgo.WithSecurity(components.Security{
			PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
		}),
	)

	res, err := s.Notifications.ListUserConfigurations(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.UserConfigurationListResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security schemes globally:

| Name                       | Type | Scheme      |
| -------------------------- | ---- | ----------- |
| `PersonalAccessToken`      | http | HTTP Bearer |
| `SystemAccountAccessToken` | http | HTTP Bearer |
| `KonnectAccessToken`       | http | HTTP Bearer |
| `ServiceAccessToken`       | http | HTTP Bearer |

You can set the security parameters through the `WithSecurity` option when initializing the SDK client instance. The selected scheme will be used by default to authenticate with the API for all operations that support it. For example:
```go
package main

import (
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := sdkkonnectgo.New(
		sdkkonnectgo.WithSecurity(components.Security{
			PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
		}),
	)

	res, err := s.Notifications.ListUserConfigurations(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.UserConfigurationListResponse != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [ACLs](docs/sdks/acls/README.md)

* [ListACL](docs/sdks/acls/README.md#listacl) - List all ACLs
* [GetACL](docs/sdks/acls/README.md#getacl) - Get an ACL
* [ListACLWithConsumer](docs/sdks/acls/README.md#listaclwithconsumer) - List all ACLs associated with a Consumer
* [CreateACLWithConsumer](docs/sdks/acls/README.md#createaclwithconsumer) - Create a new ACL associated with a Consumer
* [DeleteACLWithConsumer](docs/sdks/acls/README.md#deleteaclwithconsumer) - Delete a an ACL associated with a Consumer
* [GetACLWithConsumer](docs/sdks/acls/README.md#getaclwithconsumer) - Get an ACL associated with a Consumer
* [UpsertACLWithConsumer](docs/sdks/acls/README.md#upsertaclwithconsumer) - Upsert an ACL associated with a Consumer

### [API](docs/sdks/api/README.md)

* [CreateAPI](docs/sdks/api/README.md#createapi) - Create API
* [ListApis](docs/sdks/api/README.md#listapis) - List APIs
* [ListApisComputed](docs/sdks/api/README.md#listapiscomputed) - List APIs computed
* [FetchAPI](docs/sdks/api/README.md#fetchapi) - Get an API
* [UpdateAPI](docs/sdks/api/README.md#updateapi) - Update API
* [DeleteAPI](docs/sdks/api/README.md#deleteapi) - Delete API

### [APIAttributes](docs/sdks/apiattributes/README.md)

* [ListAPIAttributes](docs/sdks/apiattributes/README.md#listapiattributes) - List API Attributes

### [APIDocumentation](docs/sdks/apidocumentation/README.md)

* [CreateAPIDocument](docs/sdks/apidocumentation/README.md#createapidocument) - Create API Document
* [ListAPIDocuments](docs/sdks/apidocumentation/README.md#listapidocuments) - List API Documents
* [FetchAPIDocument](docs/sdks/apidocumentation/README.md#fetchapidocument) - Get an API Document
* [UpdateAPIDocument](docs/sdks/apidocumentation/README.md#updateapidocument) - Update API Document
* [DeleteAPIDocument](docs/sdks/apidocumentation/README.md#deleteapidocument) - Delete API Documentation
* [MoveAPIDocument](docs/sdks/apidocumentation/README.md#moveapidocument) - Move API Documentation

### [APIImplementation](docs/sdks/apiimplementation/README.md)

* [ListAPIImplementations](docs/sdks/apiimplementation/README.md#listapiimplementations) - List API Implementations
* [CreateAPIImplementation](docs/sdks/apiimplementation/README.md#createapiimplementation) - Create API Implementation
* [FetchAPIImplementation](docs/sdks/apiimplementation/README.md#fetchapiimplementation) - Get an API Implementation
* [DeleteAPIImplementation](docs/sdks/apiimplementation/README.md#deleteapiimplementation) - Delete API Implementation

### [APIKeys](docs/sdks/apikeys/README.md)

* [ListKeyAuthWithConsumer](docs/sdks/apikeys/README.md#listkeyauthwithconsumer) - List all API-keys associated with a Consumer
* [CreateKeyAuthWithConsumer](docs/sdks/apikeys/README.md#createkeyauthwithconsumer) - Create a new API-key associated with a Consumer
* [DeleteKeyAuthWithConsumer](docs/sdks/apikeys/README.md#deletekeyauthwithconsumer) - Delete a an API-key associated with a Consumer
* [GetKeyAuthWithConsumer](docs/sdks/apikeys/README.md#getkeyauthwithconsumer) - Get an API-key associated with a Consumer
* [UpsertKeyAuthWithConsumer](docs/sdks/apikeys/README.md#upsertkeyauthwithconsumer) - Upsert an API-key associated with a Consumer
* [ListKeyAuth](docs/sdks/apikeys/README.md#listkeyauth) - List all API-keys
* [GetKeyAuth](docs/sdks/apikeys/README.md#getkeyauth) - Get an API-key

### [APIPublication](docs/sdks/apipublication/README.md)

* [ListAPIPublications](docs/sdks/apipublication/README.md#listapipublications) - List Publications
* [PublishAPIToPortal](docs/sdks/apipublication/README.md#publishapitoportal) - Publish API
* [FetchPublication](docs/sdks/apipublication/README.md#fetchpublication) - Get a Publication
* [DeletePublication](docs/sdks/apipublication/README.md#deletepublication) - Delete Publication

### [APISpecification](docs/sdks/apispecification/README.md)

* [~~CreateAPISpec~~](docs/sdks/apispecification/README.md#createapispec) - Create API Specification :warning: **Deprecated**
* [~~ListAPISpecs~~](docs/sdks/apispecification/README.md#listapispecs) - List API Specifications :warning: **Deprecated**
* [~~FetchAPISpec~~](docs/sdks/apispecification/README.md#fetchapispec) - Get API Specification :warning: **Deprecated**
* [~~UpdateAPISpec~~](docs/sdks/apispecification/README.md#updateapispec) - Update API Specification :warning: **Deprecated**
* [~~DeleteAPISpec~~](docs/sdks/apispecification/README.md#deleteapispec) - Delete API Specification :warning: **Deprecated**
* [ValidateSpecification](docs/sdks/apispecification/README.md#validatespecification) - Validate API Specification

### [APIVersion](docs/sdks/apiversion/README.md)

* [CreateAPIVersion](docs/sdks/apiversion/README.md#createapiversion) - Create API Version
* [ListAPIVersions](docs/sdks/apiversion/README.md#listapiversions) - List API Versions
* [FetchAPIVersion](docs/sdks/apiversion/README.md#fetchapiversion) - Get an API Version
* [UpdateAPIVersion](docs/sdks/apiversion/README.md#updateapiversion) - Update API Version
* [DeleteAPIVersion](docs/sdks/apiversion/README.md#deleteapiversion) - Delete API Version

### [AppAuthStrategies](docs/sdks/appauthstrategies/README.md)

* [CreateAppAuthStrategy](docs/sdks/appauthstrategies/README.md#createappauthstrategy) - Create App Auth Strategy
* [ListAppAuthStrategies](docs/sdks/appauthstrategies/README.md#listappauthstrategies) - List App Auth Strategies
* [GetAppAuthStrategy](docs/sdks/appauthstrategies/README.md#getappauthstrategy) - Get App Auth Strategy
* [ReplaceAppAuthStrategy](docs/sdks/appauthstrategies/README.md#replaceappauthstrategy) - Replace App Auth Strategy
* [UpdateAppAuthStrategy](docs/sdks/appauthstrategies/README.md#updateappauthstrategy) - Update App Auth Strategy
* [DeleteAppAuthStrategy](docs/sdks/appauthstrategies/README.md#deleteappauthstrategy) - Delete App Auth Strategy

### [ApplicationRegistrations](docs/sdks/applicationregistrations/README.md)

* [ListRegistrations](docs/sdks/applicationregistrations/README.md#listregistrations) - List Registrations by Portal
* [ListRegistrationsByApplication](docs/sdks/applicationregistrations/README.md#listregistrationsbyapplication) - List Registrations by Application
* [GetApplicationRegistration](docs/sdks/applicationregistrations/README.md#getapplicationregistration) - Get a Registration
* [UpdateApplicationRegistration](docs/sdks/applicationregistrations/README.md#updateapplicationregistration) - Update Registration
* [DeleteApplicationRegistration](docs/sdks/applicationregistrations/README.md#deleteapplicationregistration) - Delete Registration

### [Applications](docs/sdks/applications/README.md)

* [GetApplicationUnscoped](docs/sdks/applications/README.md#getapplicationunscoped) - Get an Application
* [ListApplications](docs/sdks/applications/README.md#listapplications) - List Applications
* [GetApplication](docs/sdks/applications/README.md#getapplication) - Get an Application by Portal
* [DeleteApplication](docs/sdks/applications/README.md#deleteapplication) - Delete Application by Portal
* [ListDevelopersByApplication](docs/sdks/applications/README.md#listdevelopersbyapplication) - List Developers by Application

### [Assets](docs/sdks/assets/README.md)

* [GetPortalAssetFavicon](docs/sdks/assets/README.md#getportalassetfavicon) - Get Favicon
* [ReplacePortalAssetFavicon](docs/sdks/assets/README.md#replaceportalassetfavicon) - Replace Favicon
* [GetPortalAssetFaviconRaw](docs/sdks/assets/README.md#getportalassetfaviconraw) - Get Favicon (Raw)
* [GetPortalAssetLogo](docs/sdks/assets/README.md#getportalassetlogo) - Get Logo
* [ReplacePortalAssetLogo](docs/sdks/assets/README.md#replaceportalassetlogo) - Replace Logo
* [GetPortalAssetLogoRaw](docs/sdks/assets/README.md#getportalassetlogoraw) - Get Logo (Raw)

### [Authentication](docs/sdks/authentication/README.md)

* [AuthenticateSso](docs/sdks/authentication/README.md#authenticatesso) - SSO Callback

### [AuthSettings](docs/sdks/authsettings/README.md)

* [GetAuthenticationSettings](docs/sdks/authsettings/README.md#getauthenticationsettings) - Get Auth Settings
* [UpdateAuthenticationSettings](docs/sdks/authsettings/README.md#updateauthenticationsettings) - Update Auth Settings
* [GetIdpConfiguration](docs/sdks/authsettings/README.md#getidpconfiguration) - Get the IdP Configuration
* [UpdateIdpConfiguration](docs/sdks/authsettings/README.md#updateidpconfiguration) - Update IdP Configuration
* [GetTeamGroupMappings](docs/sdks/authsettings/README.md#getteamgroupmappings) - Get a Team Group Mappings
* [PatchTeamGroupMappings](docs/sdks/authsettings/README.md#patchteamgroupmappings) - Patch Mappings by Team ID
* [UpdateIdpTeamMappings](docs/sdks/authsettings/README.md#updateidpteammappings) - Update Team Mappings
* [GetIdpTeamMappings](docs/sdks/authsettings/README.md#getidpteammappings) - Get a Team Mapping
* [GetIdentityProviders](docs/sdks/authsettings/README.md#getidentityproviders) - List Identity Providers
* [CreateIdentityProvider](docs/sdks/authsettings/README.md#createidentityprovider) - Create Identity Provider
* [GetIdentityProvider](docs/sdks/authsettings/README.md#getidentityprovider) - Get Identity Provider
* [UpdateIdentityProvider](docs/sdks/authsettings/README.md#updateidentityprovider) - Update Identity Provider
* [DeleteIdentityProvider](docs/sdks/authsettings/README.md#deleteidentityprovider) - Delete Identity Provider

### [BasicAuthCredentials](docs/sdks/basicauthcredentials/README.md)

* [ListBasicAuth](docs/sdks/basicauthcredentials/README.md#listbasicauth) - List all Basic-auth credentials
* [GetBasicAuth](docs/sdks/basicauthcredentials/README.md#getbasicauth) - Get a Basic-auth credential
* [ListBasicAuthWithConsumer](docs/sdks/basicauthcredentials/README.md#listbasicauthwithconsumer) - List all Basic-auth credentials associated with a Consumer
* [CreateBasicAuthWithConsumer](docs/sdks/basicauthcredentials/README.md#createbasicauthwithconsumer) - Create a new Basic-auth credential associated with a Consumer
* [DeleteBasicAuthWithConsumer](docs/sdks/basicauthcredentials/README.md#deletebasicauthwithconsumer) - Delete a a Basic-auth credential associated with a Consumer
* [GetBasicAuthWithConsumer](docs/sdks/basicauthcredentials/README.md#getbasicauthwithconsumer) - Get a Basic-auth credential associated with a Consumer
* [UpsertBasicAuthWithConsumer](docs/sdks/basicauthcredentials/README.md#upsertbasicauthwithconsumer) - Upsert a Basic-auth credential associated with a Consumer

### [CACertificates](docs/sdks/cacertificates/README.md)

* [ListCaCertificate](docs/sdks/cacertificates/README.md#listcacertificate) - List all CA Certificates
* [CreateCaCertificate](docs/sdks/cacertificates/README.md#createcacertificate) - Create a new CA Certificate
* [DeleteCaCertificate](docs/sdks/cacertificates/README.md#deletecacertificate) - Delete a CA Certificate
* [GetCaCertificate](docs/sdks/cacertificates/README.md#getcacertificate) - Get a CA Certificate
* [UpsertCaCertificate](docs/sdks/cacertificates/README.md#upsertcacertificate) - Upsert a CA Certificate

### [Certificates](docs/sdks/certificates/README.md)

* [ListCertificate](docs/sdks/certificates/README.md#listcertificate) - List all Certificates
* [CreateCertificate](docs/sdks/certificates/README.md#createcertificate) - Create a new Certificate
* [DeleteCertificate](docs/sdks/certificates/README.md#deletecertificate) - Delete a Certificate
* [GetCertificate](docs/sdks/certificates/README.md#getcertificate) - Get a Certificate
* [UpsertCertificate](docs/sdks/certificates/README.md#upsertcertificate) - Upsert a Certificate

### [CloudGateways](docs/sdks/cloudgateways/README.md)

* [GetAvailabilityJSON](docs/sdks/cloudgateways/README.md#getavailabilityjson) - Get Resource Availability JSON
* [ListConfigurations](docs/sdks/cloudgateways/README.md#listconfigurations) - List Configurations
* [CreateConfiguration](docs/sdks/cloudgateways/README.md#createconfiguration) - Create Configuration
* [GetConfiguration](docs/sdks/cloudgateways/README.md#getconfiguration) - Get Configuration
* [ListCustomDomains](docs/sdks/cloudgateways/README.md#listcustomdomains) - List Custom Domains
* [CreateCustomDomains](docs/sdks/cloudgateways/README.md#createcustomdomains) - Create Custom Domain
* [GetCustomDomain](docs/sdks/cloudgateways/README.md#getcustomdomain) - Get Custom Domain
* [DeleteCustomDomain](docs/sdks/cloudgateways/README.md#deletecustomdomain) - Delete Custom Domain
* [GetCustomDomainOnlineStatus](docs/sdks/cloudgateways/README.md#getcustomdomainonlinestatus) - Get Custom Domain Online Status
* [ListDefaultResourceConfigurations](docs/sdks/cloudgateways/README.md#listdefaultresourceconfigurations) - List Default Resource Configurations
* [ListDefaultResourceQuotas](docs/sdks/cloudgateways/README.md#listdefaultresourcequotas) - List Default Resource Quotas
* [ListNetworks](docs/sdks/cloudgateways/README.md#listnetworks) - List Networks
* [CreateNetwork](docs/sdks/cloudgateways/README.md#createnetwork) - Create Network
* [GetNetwork](docs/sdks/cloudgateways/README.md#getnetwork) - Get Network
* [UpdateNetwork](docs/sdks/cloudgateways/README.md#updatenetwork) - Update Network
* [DeleteNetwork](docs/sdks/cloudgateways/README.md#deletenetwork) - Delete Network
* [ListNetworkConfigurations](docs/sdks/cloudgateways/README.md#listnetworkconfigurations) - List Network Configuration References
* [ListPrivateDNS](docs/sdks/cloudgateways/README.md#listprivatedns) - List Private DNS
* [CreatePrivateDNS](docs/sdks/cloudgateways/README.md#createprivatedns) - Create Private DNS
* [GetPrivateDNS](docs/sdks/cloudgateways/README.md#getprivatedns) - Get Private DNS
* [UpdatePrivateDNS](docs/sdks/cloudgateways/README.md#updateprivatedns) - Update Private DNS
* [DeletePrivateDNS](docs/sdks/cloudgateways/README.md#deleteprivatedns) - Delete Private DNS
* [ListTransitGateways](docs/sdks/cloudgateways/README.md#listtransitgateways) - List Transit Gateways
* [CreateTransitGateway](docs/sdks/cloudgateways/README.md#createtransitgateway) - Create Transit Gateway
* [GetTransitGateway](docs/sdks/cloudgateways/README.md#gettransitgateway) - Get Transit Gateway
* [UpdateTransitGateway](docs/sdks/cloudgateways/README.md#updatetransitgateway) - Update Transit Gateway
* [DeleteTransitGateway](docs/sdks/cloudgateways/README.md#deletetransitgateway) - Delete Transit Gateway
* [ListProviderAccounts](docs/sdks/cloudgateways/README.md#listprovideraccounts) - List Provider Accounts
* [GetProviderAccount](docs/sdks/cloudgateways/README.md#getprovideraccount) - Get Provider Account
* [ListResourceConfigurations](docs/sdks/cloudgateways/README.md#listresourceconfigurations) - List Resource Configurations
* [GetResourceConfiguration](docs/sdks/cloudgateways/README.md#getresourceconfiguration) - Get Resource Configuration
* [ListResourceQuotas](docs/sdks/cloudgateways/README.md#listresourcequotas) - List Resource Quotas
* [GetResourceQuota](docs/sdks/cloudgateways/README.md#getresourcequota) - Get Resource Quota

### [ConfigStores](docs/sdks/configstores/README.md)

* [ListConfigStores](docs/sdks/configstores/README.md#listconfigstores) - List all config stores for a control plane
* [CreateConfigStore](docs/sdks/configstores/README.md#createconfigstore) - Create Config Store
* [GetConfigStore](docs/sdks/configstores/README.md#getconfigstore) - Get a Config Store
* [UpdateConfigStore](docs/sdks/configstores/README.md#updateconfigstore) - Update an individual Config Store
* [DeleteConfigStore](docs/sdks/configstores/README.md#deleteconfigstore) - Delete Config Store

### [ConfigStoreSecrets](docs/sdks/configstoresecrets/README.md)

* [CreateConfigStoreSecret](docs/sdks/configstoresecrets/README.md#createconfigstoresecret) - Create Config Store Secret
* [ListConfigStoreSecrets](docs/sdks/configstoresecrets/README.md#listconfigstoresecrets) - List Config Store Secrets
* [GetConfigStoreSecret](docs/sdks/configstoresecrets/README.md#getconfigstoresecret) - Get a Config Store Secret
* [UpdateConfigStoreSecret](docs/sdks/configstoresecrets/README.md#updateconfigstoresecret) - Update Config Store Secret
* [DeleteConfigStoreSecret](docs/sdks/configstoresecrets/README.md#deleteconfigstoresecret) - Delete Config Store Secret

### [ConsumerGroups](docs/sdks/consumergroups/README.md)

* [ListConsumerGroup](docs/sdks/consumergroups/README.md#listconsumergroup) - List all Consumer Groups
* [CreateConsumerGroup](docs/sdks/consumergroups/README.md#createconsumergroup) - Create a new Consumer Group
* [DeleteConsumerGroup](docs/sdks/consumergroups/README.md#deleteconsumergroup) - Delete a Consumer Group
* [GetConsumerGroup](docs/sdks/consumergroups/README.md#getconsumergroup) - Get a Consumer Group
* [UpsertConsumerGroup](docs/sdks/consumergroups/README.md#upsertconsumergroup) - Upsert a Consumer Group
* [RemoveAllConsumersFromConsumerGroup](docs/sdks/consumergroups/README.md#removeallconsumersfromconsumergroup) - Remove consumers from consumer group
* [ListConsumersForConsumerGroup](docs/sdks/consumergroups/README.md#listconsumersforconsumergroup) - List all Consumers in a Consumer Group
* [AddConsumerToGroup](docs/sdks/consumergroups/README.md#addconsumertogroup) - Add consumer to consumer group
* [RemoveConsumerFromGroup](docs/sdks/consumergroups/README.md#removeconsumerfromgroup) - Remove consumer from consumer group

### [Consumers](docs/sdks/consumers/README.md)

* [ListConsumer](docs/sdks/consumers/README.md#listconsumer) - List all Consumers
* [CreateConsumer](docs/sdks/consumers/README.md#createconsumer) - Create a new Consumer
* [DeleteConsumer](docs/sdks/consumers/README.md#deleteconsumer) - Delete a Consumer
* [GetConsumer](docs/sdks/consumers/README.md#getconsumer) - Get a Consumer
* [UpsertConsumer](docs/sdks/consumers/README.md#upsertconsumer) - Upsert a Consumer
* [RemoveConsumerFromAllConsumerGroups](docs/sdks/consumers/README.md#removeconsumerfromallconsumergroups) - Remove consumer from all consumer groups
* [ListConsumerGroupsForConsumer](docs/sdks/consumers/README.md#listconsumergroupsforconsumer) - List all Consumer Groups a Consumer belongs to
* [AddConsumerToSpecificConsumerGroup](docs/sdks/consumers/README.md#addconsumertospecificconsumergroup) - Add consumer to a specific consumer group
* [RemoveConsumerFromConsumerGroup](docs/sdks/consumers/README.md#removeconsumerfromconsumergroup) - Remove consumer from consumer group

### [ControlPlaneGroups](docs/sdks/controlplanegroups/README.md)

* [GetControlPlanesIDGroupMemberStatus](docs/sdks/controlplanegroups/README.md#getcontrolplanesidgroupmemberstatus) - Get Control Plane Group Member Status
* [GetControlPlanesIDGroupMemberships](docs/sdks/controlplanegroups/README.md#getcontrolplanesidgroupmemberships) - List Control Plane Group Memberships
* [PutControlPlanesIDGroupMemberships](docs/sdks/controlplanegroups/README.md#putcontrolplanesidgroupmemberships) - Upsert Control Plane Group Members
* [PostControlPlanesIDGroupMembershipsAdd](docs/sdks/controlplanegroups/README.md#postcontrolplanesidgroupmembershipsadd) - Add Control Plane Group Members
* [PostControlPlanesIDGroupMembershipsRemove](docs/sdks/controlplanegroups/README.md#postcontrolplanesidgroupmembershipsremove) - Remove Control Plane Group Members
* [GetControlPlanesIDGroupStatus](docs/sdks/controlplanegroups/README.md#getcontrolplanesidgroupstatus) - Get Control Plane Group Status

### [ControlPlanes](docs/sdks/controlplanes/README.md)

* [ListControlPlanes](docs/sdks/controlplanes/README.md#listcontrolplanes) - List Control Planes
* [CreateControlPlane](docs/sdks/controlplanes/README.md#createcontrolplane) - Create Control Plane
* [GetControlPlane](docs/sdks/controlplanes/README.md#getcontrolplane) - Get a Control Plane
* [UpdateControlPlane](docs/sdks/controlplanes/README.md#updatecontrolplane) - Update Control Plane
* [DeleteControlPlane](docs/sdks/controlplanes/README.md#deletecontrolplane) - Delete Control Plane

### [CustomPlugins](docs/sdks/customplugins/README.md)

* [ListCustomPlugin](docs/sdks/customplugins/README.md#listcustomplugin) - List all CustomPlugins
* [CreateCustomPlugin](docs/sdks/customplugins/README.md#createcustomplugin) - Create a new CustomPlugin
* [DeleteCustomPlugin](docs/sdks/customplugins/README.md#deletecustomplugin) - Delete a CustomPlugin
* [GetCustomPlugin](docs/sdks/customplugins/README.md#getcustomplugin) - Get a CustomPlugin
* [UpsertCustomPlugin](docs/sdks/customplugins/README.md#upsertcustomplugin) - Upsert a CustomPlugin

### [CustomPluginSchemas](docs/sdks/custompluginschemas/README.md)

* [ListPluginSchemas](docs/sdks/custompluginschemas/README.md#listpluginschemas) - List Custom Plugin Schemas
* [CreatePluginSchemas](docs/sdks/custompluginschemas/README.md#createpluginschemas) - Upload custom plugin schema
* [GetPluginSchema](docs/sdks/custompluginschemas/README.md#getpluginschema) - Get a custom plugin schema
* [DeletePluginSchemas](docs/sdks/custompluginschemas/README.md#deletepluginschemas) - Delete custom plugin schema
* [UpdatePluginSchemas](docs/sdks/custompluginschemas/README.md#updatepluginschemas) - Create or update a custom plugin schema

### [DCRProviders](docs/sdks/dcrproviders/README.md)

* [CreateDcrProvider](docs/sdks/dcrproviders/README.md#createdcrprovider) - Create DCR provider
* [ListDcrProviders](docs/sdks/dcrproviders/README.md#listdcrproviders) - List DCR Providers
* [GetDcrProvider](docs/sdks/dcrproviders/README.md#getdcrprovider) - Get a DCR provider
* [UpdateDcrProvider](docs/sdks/dcrproviders/README.md#updatedcrprovider) - Update DCR provider
* [DeleteDcrProvider](docs/sdks/dcrproviders/README.md#deletedcrprovider) - Delete DCR provider
* [VerifyDcrProvider](docs/sdks/dcrproviders/README.md#verifydcrprovider) - Verify DCR provider configuration

### [DegraphqlRoutes](docs/sdks/degraphqlroutes/README.md)

* [ListDegraphqlRoute](docs/sdks/degraphqlroutes/README.md#listdegraphqlroute) - List all Degraphql_routes
* [GetDegraphqlRoute](docs/sdks/degraphqlroutes/README.md#getdegraphqlroute) - Get a Degraphql_route
* [ListDegraphqlRouteWithService](docs/sdks/degraphqlroutes/README.md#listdegraphqlroutewithservice) - List all Degraphql_routes associated with a Service
* [CreateDegraphqlRouteWithService](docs/sdks/degraphqlroutes/README.md#createdegraphqlroutewithservice) - Create a new Degraphql_route associated with a Service
* [DeleteDegraphqlRouteWithService](docs/sdks/degraphqlroutes/README.md#deletedegraphqlroutewithservice) - Delete a a Degraphql_route associated with a Service
* [GetDegraphqlRouteWithService](docs/sdks/degraphqlroutes/README.md#getdegraphqlroutewithservice) - Get a Degraphql_route associated with a Service
* [UpsertDegraphqlRouteWithService](docs/sdks/degraphqlroutes/README.md#upsertdegraphqlroutewithservice) - Upsert a Degraphql_route associated with a Service

### [DPCertificates](docs/sdks/dpcertificates/README.md)

* [ListDpClientCertificates](docs/sdks/dpcertificates/README.md#listdpclientcertificates) - List DP Client Certificates
* [CreateDataplaneCertificate](docs/sdks/dpcertificates/README.md#createdataplanecertificate) - Pin New DP Client Certificate
* [GetDataplaneCertificate](docs/sdks/dpcertificates/README.md#getdataplanecertificate) - Get a DP Client Certificate
* [DeleteDataplaneCertificate](docs/sdks/dpcertificates/README.md#deletedataplanecertificate) - Delete DP Client Certificate

### [DPNodes](docs/sdks/dpnodes/README.md)

* [GetExpectedConfigHash](docs/sdks/dpnodes/README.md#getexpectedconfighash) - Get an Expected Config Hash
* [ListDataplaneNodes](docs/sdks/dpnodes/README.md#listdataplanenodes) - List Data Plane Node Records
* [GetNodesNodeID](docs/sdks/dpnodes/README.md#getnodesnodeid) - Get a Data Plane Node Record
* [DeleteNodesNodeID](docs/sdks/dpnodes/README.md#deletenodesnodeid) - Delete Data Plane Node Record
* [GetNodesEol](docs/sdks/dpnodes/README.md#getnodeseol) - List End-of-Life Data Plane Node Records

### [HMACAuthCredentials](docs/sdks/hmacauthcredentials/README.md)

* [ListHmacAuthWithConsumer](docs/sdks/hmacauthcredentials/README.md#listhmacauthwithconsumer) - List all HMAC-auth credentials associated with a Consumer
* [CreateHmacAuthWithConsumer](docs/sdks/hmacauthcredentials/README.md#createhmacauthwithconsumer) - Create a new HMAC-auth credential associated with a Consumer
* [DeleteHmacAuthWithConsumer](docs/sdks/hmacauthcredentials/README.md#deletehmacauthwithconsumer) - Delete a a HMAC-auth credential associated with a Consumer
* [GetHmacAuthWithConsumer](docs/sdks/hmacauthcredentials/README.md#gethmacauthwithconsumer) - Get a HMAC-auth credential associated with a Consumer
* [UpsertHmacAuthWithConsumer](docs/sdks/hmacauthcredentials/README.md#upserthmacauthwithconsumer) - Upsert a HMAC-auth credential associated with a Consumer
* [ListHmacAuth](docs/sdks/hmacauthcredentials/README.md#listhmacauth) - List all HMAC-auth credentials
* [GetHmacAuth](docs/sdks/hmacauthcredentials/README.md#gethmacauth) - Get a HMAC-auth credential

### [ImpersonationSettings](docs/sdks/impersonationsettings/README.md)

* [GetImpersonationSettings](docs/sdks/impersonationsettings/README.md#getimpersonationsettings) - Get Impersonation Settings
* [UpdateImpersonationSettings](docs/sdks/impersonationsettings/README.md#updateimpersonationsettings) - Update Impersonation Settings

### [Invites](docs/sdks/invites/README.md)

* [InviteUser](docs/sdks/invites/README.md#inviteuser) - Invite User

### [JWTs](docs/sdks/jwts/README.md)

* [ListJwtWithConsumer](docs/sdks/jwts/README.md#listjwtwithconsumer) - List all JWTs associated with a Consumer
* [CreateJwtWithConsumer](docs/sdks/jwts/README.md#createjwtwithconsumer) - Create a new JWT associated with a Consumer
* [DeleteJwtWithConsumer](docs/sdks/jwts/README.md#deletejwtwithconsumer) - Delete a a JWT associated with a Consumer
* [GetJwtWithConsumer](docs/sdks/jwts/README.md#getjwtwithconsumer) - Get a JWT associated with a Consumer
* [UpsertJwtWithConsumer](docs/sdks/jwts/README.md#upsertjwtwithconsumer) - Upsert a JWT associated with a Consumer
* [ListJwt](docs/sdks/jwts/README.md#listjwt) - List all JWTs
* [GetJwt](docs/sdks/jwts/README.md#getjwt) - Get a JWT

### [Keys](docs/sdks/keys/README.md)

* [ListKeyWithKeySet](docs/sdks/keys/README.md#listkeywithkeyset) - List all Keys associated with a KeySet
* [CreateKeyWithKeySet](docs/sdks/keys/README.md#createkeywithkeyset) - Create a new Key associated with a KeySet
* [DeleteKeyWithKeySet](docs/sdks/keys/README.md#deletekeywithkeyset) - Delete a a Key associated with a KeySet
* [GetKeyWithKeySet](docs/sdks/keys/README.md#getkeywithkeyset) - Get a Key associated with a KeySet
* [UpsertKeyWithKeySet](docs/sdks/keys/README.md#upsertkeywithkeyset) - Upsert a Key associated with a KeySet
* [ListKey](docs/sdks/keys/README.md#listkey) - List all Keys
* [CreateKey](docs/sdks/keys/README.md#createkey) - Create a new Key
* [DeleteKey](docs/sdks/keys/README.md#deletekey) - Delete a Key
* [GetKey](docs/sdks/keys/README.md#getkey) - Get a Key
* [UpsertKey](docs/sdks/keys/README.md#upsertkey) - Upsert a Key

### [KeySets](docs/sdks/keysets/README.md)

* [ListKeySet](docs/sdks/keysets/README.md#listkeyset) - List all KeySets
* [CreateKeySet](docs/sdks/keysets/README.md#createkeyset) - Create a new KeySet
* [DeleteKeySet](docs/sdks/keysets/README.md#deletekeyset) - Delete a KeySet
* [GetKeySet](docs/sdks/keysets/README.md#getkeyset) - Get a KeySet
* [UpsertKeySet](docs/sdks/keysets/README.md#upsertkeyset) - Upsert a KeySet

### [Me](docs/sdks/me/README.md)

* [GetOrganizationsMe](docs/sdks/me/README.md#getorganizationsme) - Get My Organization
* [GetUsersMe](docs/sdks/me/README.md#getusersme) - Get My User Account

### [MTLSAuthCredentials](docs/sdks/mtlsauthcredentials/README.md)

* [ListMtlsAuthWithConsumer](docs/sdks/mtlsauthcredentials/README.md#listmtlsauthwithconsumer) - List all MTLS-auth credentials associated with a Consumer
* [CreateMtlsAuthWithConsumer](docs/sdks/mtlsauthcredentials/README.md#createmtlsauthwithconsumer) - Create a new MTLS-auth credential associated with a Consumer
* [DeleteMtlsAuthWithConsumer](docs/sdks/mtlsauthcredentials/README.md#deletemtlsauthwithconsumer) - Delete a a MTLS-auth credential associated with a Consumer
* [GetMtlsAuthWithConsumer](docs/sdks/mtlsauthcredentials/README.md#getmtlsauthwithconsumer) - Get a MTLS-auth credential associated with a Consumer
* [UpsertMtlsAuthWithConsumer](docs/sdks/mtlsauthcredentials/README.md#upsertmtlsauthwithconsumer) - Upsert a MTLS-auth credential associated with a Consumer
* [ListMtlsAuth](docs/sdks/mtlsauthcredentials/README.md#listmtlsauth) - List all MTLS-auth credentials
* [GetMtlsAuth](docs/sdks/mtlsauthcredentials/README.md#getmtlsauth) - Get a MTLS-auth credential

### [Notifications](docs/sdks/notifications/README.md)

* [ListUserConfigurations](docs/sdks/notifications/README.md#listuserconfigurations) - List available user configurations
* [ListEventSubscriptions](docs/sdks/notifications/README.md#listeventsubscriptions) - List event subscriptions
* [CreateEventSubscription](docs/sdks/notifications/README.md#createeventsubscription) - Create a new subscription for an event
* [GetEventSubscription](docs/sdks/notifications/README.md#geteventsubscription) - Get subscription for an event
* [UpdateEventSubscription](docs/sdks/notifications/README.md#updateeventsubscription) - Update subscription for an event
* [DeleteEventSubscription](docs/sdks/notifications/README.md#deleteeventsubscription) - Delete subscription associated with event
* [ListNotifications](docs/sdks/notifications/README.md#listnotifications) - List available notifications
* [GetNotificationDetails](docs/sdks/notifications/README.md#getnotificationdetails) - Get notification details
* [UpdateNotification](docs/sdks/notifications/README.md#updatenotification) - Update notification
* [DeleteNotification](docs/sdks/notifications/README.md#deletenotification) - Delete notification
* [BulkNotifications](docs/sdks/notifications/README.md#bulknotifications) - Mark a list of notifications to a status

### [Pages](docs/sdks/pages/README.md)

* [CreateDefaultContent](docs/sdks/pages/README.md#createdefaultcontent) - Creates Default Pages
* [ListPortalPages](docs/sdks/pages/README.md#listportalpages) - List Pages
* [CreatePortalPage](docs/sdks/pages/README.md#createportalpage) - Create Page
* [GetPortalPage](docs/sdks/pages/README.md#getportalpage) - Get a Page
* [UpdatePortalPage](docs/sdks/pages/README.md#updateportalpage) - Update Page
* [DeletePortalPage](docs/sdks/pages/README.md#deleteportalpage) - Delete Page
* [MovePortalPages](docs/sdks/pages/README.md#moveportalpages) - Move Page

### [PartialLinks](docs/sdks/partiallinks/README.md)

* [ListPartialLink](docs/sdks/partiallinks/README.md#listpartiallink) - List partial links

### [Partials](docs/sdks/partials/README.md)

* [ListPartial](docs/sdks/partials/README.md#listpartial) - List all Partials
* [CreatePartial](docs/sdks/partials/README.md#createpartial) - Create a new Partial
* [DeletePartial](docs/sdks/partials/README.md#deletepartial) - Delete a Partial
* [GetPartial](docs/sdks/partials/README.md#getpartial) - Get a Partial
* [UpsertPartial](docs/sdks/partials/README.md#upsertpartial) - Upsert a Partial

### [PersonalAccessTokens](docs/sdks/personalaccesstokens/README.md)

* [ListUsersPersonalAccessTokens](docs/sdks/personalaccesstokens/README.md#listuserspersonalaccesstokens) - List PATs
* [CreatePersonalAccessToken](docs/sdks/personalaccesstokens/README.md#createpersonalaccesstoken) - Create a new personal access token
* [GetPersonalAccessTokenDetails](docs/sdks/personalaccesstokens/README.md#getpersonalaccesstokendetails) - Get Personal Access Token details
* [UpdatePersonalAccessTokenDetails](docs/sdks/personalaccesstokens/README.md#updatepersonalaccesstokendetails) - Update personal access token details
* [DeletePersonalAccessToken](docs/sdks/personalaccesstokens/README.md#deletepersonalaccesstoken) - Delete personal access token
* [RevokePersonalAccessToken](docs/sdks/personalaccesstokens/README.md#revokepersonalaccesstoken) - Revoke Personal Access Token

### [Plugins](docs/sdks/plugins/README.md)

* [ListPluginWithConsumerGroup](docs/sdks/plugins/README.md#listpluginwithconsumergroup) - List all Plugins associated with a Consumer Group
* [CreatePluginWithConsumerGroup](docs/sdks/plugins/README.md#createpluginwithconsumergroup) - Create a new Plugin associated with a Consumer Group
* [DeletePluginWithConsumerGroup](docs/sdks/plugins/README.md#deletepluginwithconsumergroup) - Delete a a Plugin associated with a Consumer Group
* [GetPluginWithConsumerGroup](docs/sdks/plugins/README.md#getpluginwithconsumergroup) - Get a Plugin associated with a Consumer Group
* [UpsertPluginWithConsumerGroup](docs/sdks/plugins/README.md#upsertpluginwithconsumergroup) - Upsert a Plugin associated with a Consumer Group
* [ListPluginWithConsumer](docs/sdks/plugins/README.md#listpluginwithconsumer) - List all Plugins associated with a Consumer
* [CreatePluginWithConsumer](docs/sdks/plugins/README.md#createpluginwithconsumer) - Create a new Plugin associated with a Consumer
* [DeletePluginWithConsumer](docs/sdks/plugins/README.md#deletepluginwithconsumer) - Delete a a Plugin associated with a Consumer
* [GetPluginWithConsumer](docs/sdks/plugins/README.md#getpluginwithconsumer) - Get a Plugin associated with a Consumer
* [UpsertPluginWithConsumer](docs/sdks/plugins/README.md#upsertpluginwithconsumer) - Upsert a Plugin associated with a Consumer
* [ListPlugin](docs/sdks/plugins/README.md#listplugin) - List all Plugins
* [CreatePlugin](docs/sdks/plugins/README.md#createplugin) - Create a new Plugin
* [DeletePlugin](docs/sdks/plugins/README.md#deleteplugin) - Delete a Plugin
* [GetPlugin](docs/sdks/plugins/README.md#getplugin) - Get a Plugin
* [UpsertPlugin](docs/sdks/plugins/README.md#upsertplugin) - Upsert a Plugin
* [ListPluginWithRoute](docs/sdks/plugins/README.md#listpluginwithroute) - List all Plugins associated with a Route
* [CreatePluginWithRoute](docs/sdks/plugins/README.md#createpluginwithroute) - Create a new Plugin associated with a Route
* [DeletePluginWithRoute](docs/sdks/plugins/README.md#deletepluginwithroute) - Delete a a Plugin associated with a Route
* [GetPluginWithRoute](docs/sdks/plugins/README.md#getpluginwithroute) - Get a Plugin associated with a Route
* [UpsertPluginWithRoute](docs/sdks/plugins/README.md#upsertpluginwithroute) - Upsert a Plugin associated with a Route
* [FetchPluginSchema](docs/sdks/plugins/README.md#fetchpluginschema) - Get plugin schema
* [ListPluginWithService](docs/sdks/plugins/README.md#listpluginwithservice) - List all Plugins associated with a Service
* [CreatePluginWithService](docs/sdks/plugins/README.md#createpluginwithservice) - Create a new Plugin associated with a Service
* [DeletePluginWithService](docs/sdks/plugins/README.md#deletepluginwithservice) - Delete a a Plugin associated with a Service
* [GetPluginWithService](docs/sdks/plugins/README.md#getpluginwithservice) - Get a Plugin associated with a Service
* [UpsertPluginWithService](docs/sdks/plugins/README.md#upsertpluginwithservice) - Upsert a Plugin associated with a Service

### [PortalAuditLogs](docs/sdks/portalauditlogs/README.md)

* [UpdatePortalAuditLogReplayJob](docs/sdks/portalauditlogs/README.md#updateportalauditlogreplayjob) - Update Portal Audit Log Replay Job
* [GetPortalAuditLogReplayJob](docs/sdks/portalauditlogs/README.md#getportalauditlogreplayjob) - Get Portal Audit Log Replay Job
* [UpdatePortalAuditLogWebhook](docs/sdks/portalauditlogs/README.md#updateportalauditlogwebhook) - Update Portal Audit Log Webhook
* [GetPortalAuditLogWebhook](docs/sdks/portalauditlogs/README.md#getportalauditlogwebhook) - Get Portal Audit Log Webhook
* [DeletePortalAuditLogWebhook](docs/sdks/portalauditlogs/README.md#deleteportalauditlogwebhook) - Delete Portal Audit Log Webhook
* [GetPortalAuditLogWebhookStatus](docs/sdks/portalauditlogs/README.md#getportalauditlogwebhookstatus) - Get Portal Audit Log Webhook Status

### [PortalAuthSettings](docs/sdks/portalauthsettings/README.md)

* [GetPortalAuthenticationSettings](docs/sdks/portalauthsettings/README.md#getportalauthenticationsettings) - Get Auth Settings
* [UpdatePortalAuthenticationSettings](docs/sdks/portalauthsettings/README.md#updateportalauthenticationsettings) - Update Auth Settings
* [ListPortalTeamGroupMappings](docs/sdks/portalauthsettings/README.md#listportalteamgroupmappings) - List Team Group Mappings
* [UpdatePortalTeamGroupMappings](docs/sdks/portalauthsettings/README.md#updateportalteamgroupmappings) - Update Team Group Mappings
* [GetPortalIdentityProviders](docs/sdks/portalauthsettings/README.md#getportalidentityproviders) - List Identity Providers
* [CreatePortalIdentityProvider](docs/sdks/portalauthsettings/README.md#createportalidentityprovider) - Create Identity Provider
* [GetPortalIdentityProvider](docs/sdks/portalauthsettings/README.md#getportalidentityprovider) - Get Identity Provider
* [UpdatePortalIdentityProvider](docs/sdks/portalauthsettings/README.md#updateportalidentityprovider) - Update Identity Provider
* [DeletePortalIdentityProvider](docs/sdks/portalauthsettings/README.md#deleteportalidentityprovider) - Delete Identity Provider

### [PortalCustomDomains](docs/sdks/portalcustomdomains/README.md)

* [GetPortalCustomDomain](docs/sdks/portalcustomdomains/README.md#getportalcustomdomain) - Get Custom Domain
* [CreatePortalCustomDomain](docs/sdks/portalcustomdomains/README.md#createportalcustomdomain) - Create Custom Domain
* [UpdatePortalCustomDomain](docs/sdks/portalcustomdomains/README.md#updateportalcustomdomain) - Enable or Disable Domain
* [DeletePortalCustomDomain](docs/sdks/portalcustomdomains/README.md#deleteportalcustomdomain) - Remove Domain

### [PortalCustomization](docs/sdks/portalcustomization/README.md)

* [GetPortalCustomization](docs/sdks/portalcustomization/README.md#getportalcustomization) - Get Customization
* [ReplacePortalCustomization](docs/sdks/portalcustomization/README.md#replaceportalcustomization) - Replace Customization
* [UpdatePortalCustomization](docs/sdks/portalcustomization/README.md#updateportalcustomization) - Update Customization

### [PortalDevelopers](docs/sdks/portaldevelopers/README.md)

* [ListPortalDevelopers](docs/sdks/portaldevelopers/README.md#listportaldevelopers) - List Developers
* [GetDeveloper](docs/sdks/portaldevelopers/README.md#getdeveloper) - Get a Developer
* [UpdateDeveloper](docs/sdks/portaldevelopers/README.md#updatedeveloper) - Update Developer
* [DeleteDeveloper](docs/sdks/portaldevelopers/README.md#deletedeveloper) - Delete Developer

### [PortalEmails](docs/sdks/portalemails/README.md)

* [GetEmailConfig](docs/sdks/portalemails/README.md#getemailconfig) - Get the email config for the portal
* [CreatePortalEmailConfig](docs/sdks/portalemails/README.md#createportalemailconfig) - Create the email config for a portal
* [UpdatePortalEmailConfig](docs/sdks/portalemails/README.md#updateportalemailconfig) - Setup the email config for a portal
* [DeletePortalEmailConfig](docs/sdks/portalemails/README.md#deleteportalemailconfig) - Delete portal email config
* [~~GetEmailDelivery~~](docs/sdks/portalemails/README.md#getemaildelivery) - Get the email delivery for the portal :warning: **Deprecated**
* [~~UpdateEmailDelivery~~](docs/sdks/portalemails/README.md#updateemaildelivery) - Setup the email delivery for a portal :warning: **Deprecated**
* [~~DeleteEmailDelivery~~](docs/sdks/portalemails/README.md#deleteemaildelivery) - Delete email delivery :warning: **Deprecated**
* [ListPortalCustomEmailTemplates](docs/sdks/portalemails/README.md#listportalcustomemailtemplates) - List custom email templates for a portal
* [GetPortalCustomEmailTemplate](docs/sdks/portalemails/README.md#getportalcustomemailtemplate) - Get custom email template used in a portal
* [UpdatePortalCustomEmailTemplate](docs/sdks/portalemails/README.md#updateportalcustomemailtemplate) - Update custom email template for a portal
* [DeletePortalCustomEmailTemplate](docs/sdks/portalemails/README.md#deleteportalcustomemailtemplate) - Delete custom email template
* [PostPortalCustomEmailTestSend](docs/sdks/portalemails/README.md#postportalcustomemailtestsend) - Send Test Email
* [ListEmailDomains](docs/sdks/portalemails/README.md#listemaildomains) - List email domains
* [CreateEmailDomain](docs/sdks/portalemails/README.md#createemaildomain) - Create an email domain
* [GetEmailDomain](docs/sdks/portalemails/README.md#getemaildomain) - Get an email domain
* [DeleteEmailDomain](docs/sdks/portalemails/README.md#deleteemaildomain) - Delete an email domain
* [ListDefaultEmailTemplates](docs/sdks/portalemails/README.md#listdefaultemailtemplates) - List default email templates
* [GetDefaultEmailTemplate](docs/sdks/portalemails/README.md#getdefaultemailtemplate) - Get default email template
* [ListEmailTemplateVariables](docs/sdks/portalemails/README.md#listemailtemplatevariables) - List email template variables

### [Portals](docs/sdks/portals/README.md)

* [ListPortals](docs/sdks/portals/README.md#listportals) - List Portals
* [CreatePortal](docs/sdks/portals/README.md#createportal) - Create Portal
* [GetPortal](docs/sdks/portals/README.md#getportal) - Get a Portal
* [UpdatePortal](docs/sdks/portals/README.md#updateportal) - Update Portal
* [DeletePortal](docs/sdks/portals/README.md#deleteportal) - Delete Portal

### [PortalTeamMembership](docs/sdks/portalteammembership/README.md)

* [ListPortalDeveloperTeams](docs/sdks/portalteammembership/README.md#listportaldeveloperteams) - List Developer Teams
* [ListPortalTeamDevelopers](docs/sdks/portalteammembership/README.md#listportalteamdevelopers) - List Team Developers
* [AddDeveloperToPortalTeam](docs/sdks/portalteammembership/README.md#adddevelopertoportalteam) - Add Developer to Team
* [RemoveDeveloperFromPortalTeam](docs/sdks/portalteammembership/README.md#removedeveloperfromportalteam) - Remove Developer from Team

### [PortalTeamRoles](docs/sdks/portalteamroles/README.md)

* [ListPortalRoles](docs/sdks/portalteamroles/README.md#listportalroles) - List Portal Roles
* [ListPortalTeamRoles](docs/sdks/portalteamroles/README.md#listportalteamroles) - List Team Roles
* [AssignRoleToPortalTeams](docs/sdks/portalteamroles/README.md#assignroletoportalteams) - Assign Role
* [RemoveRoleFromPortalTeam](docs/sdks/portalteamroles/README.md#removerolefromportalteam) - Remove Role

### [PortalTeams](docs/sdks/portalteams/README.md)

* [ListPortalTeams](docs/sdks/portalteams/README.md#listportalteams) - List Teams
* [CreatePortalTeam](docs/sdks/portalteams/README.md#createportalteam) - Create Team
* [GetPortalTeam](docs/sdks/portalteams/README.md#getportalteam) - Get Team
* [UpdatePortalTeam](docs/sdks/portalteams/README.md#updateportalteam) - Update Team
* [DeletePortalTeam](docs/sdks/portalteams/README.md#deleteportalteam) - Delete Team

### [Roles](docs/sdks/roles/README.md)

* [GetPredefinedRoles](docs/sdks/roles/README.md#getpredefinedroles) - Get Predefined Roles
* [ListTeamRoles](docs/sdks/roles/README.md#listteamroles) - List Team Roles
* [TeamsAssignRole](docs/sdks/roles/README.md#teamsassignrole) - Assign Team Role
* [TeamsRemoveRole](docs/sdks/roles/README.md#teamsremoverole) - Remove Team Role
* [ListUserRoles](docs/sdks/roles/README.md#listuserroles) - List User Roles
* [UsersAssignRole](docs/sdks/roles/README.md#usersassignrole) - Assign Role
* [UsersRemoveRole](docs/sdks/roles/README.md#usersremoverole) - Remove Role

### [Routes](docs/sdks/routes/README.md)

* [ListRoute](docs/sdks/routes/README.md#listroute) - List all Routes
* [CreateRoute](docs/sdks/routes/README.md#createroute) - Create a new Route
* [DeleteRoute](docs/sdks/routes/README.md#deleteroute) - Delete a Route
* [GetRoute](docs/sdks/routes/README.md#getroute) - Get a Route
* [UpsertRoute](docs/sdks/routes/README.md#upsertroute) - Upsert a Route
* [ListRouteWithService](docs/sdks/routes/README.md#listroutewithservice) - List all Routes associated with a Service
* [CreateRouteWithService](docs/sdks/routes/README.md#createroutewithservice) - Create a new Route associated with a Service
* [DeleteRouteWithService](docs/sdks/routes/README.md#deleteroutewithservice) - Delete a a Route associated with a Service
* [GetRouteWithService](docs/sdks/routes/README.md#getroutewithservice) - Get a Route associated with a Service
* [UpsertRouteWithService](docs/sdks/routes/README.md#upsertroutewithservice) - Upsert a Route associated with a Service

### [Schemas](docs/sdks/schemas/README.md)

* [ValidateEntitySchema](docs/sdks/schemas/README.md#validateentityschema) - Validate entity schema
* [FetchPartialSchema](docs/sdks/schemas/README.md#fetchpartialschema) - Get partial schema

### [Services](docs/sdks/services/README.md)

* [ListService](docs/sdks/services/README.md#listservice) - List all Services
* [CreateService](docs/sdks/services/README.md#createservice) - Create a new Service
* [DeleteService](docs/sdks/services/README.md#deleteservice) - Delete a Service
* [GetService](docs/sdks/services/README.md#getservice) - Get a Service
* [UpsertService](docs/sdks/services/README.md#upsertservice) - Upsert a Service

### [Snippets](docs/sdks/snippets/README.md)

* [ListPortalSnippets](docs/sdks/snippets/README.md#listportalsnippets) - List Snippets
* [CreatePortalSnippet](docs/sdks/snippets/README.md#createportalsnippet) - Create Snippet
* [GetPortalSnippet](docs/sdks/snippets/README.md#getportalsnippet) - Get a Snippet
* [UpdatePortalSnippet](docs/sdks/snippets/README.md#updateportalsnippet) - Update Snippet
* [DeletePortalSnippet](docs/sdks/snippets/README.md#deleteportalsnippet) - Delete Snippet

### [SNIs](docs/sdks/snis/README.md)

* [ListSniWithCertificate](docs/sdks/snis/README.md#listsniwithcertificate) - List all SNIs associated with a Certificate
* [CreateSniWithCertificate](docs/sdks/snis/README.md#createsniwithcertificate) - Create a new SNI associated with a Certificate
* [DeleteSniWithCertificate](docs/sdks/snis/README.md#deletesniwithcertificate) - Delete a an SNI associated with a Certificate
* [GetSniWithCertificate](docs/sdks/snis/README.md#getsniwithcertificate) - Get an SNI associated with a Certificate
* [UpsertSniWithCertificate](docs/sdks/snis/README.md#upsertsniwithcertificate) - Upsert an SNI associated with a Certificate
* [ListSni](docs/sdks/snis/README.md#listsni) - List all SNIs
* [CreateSni](docs/sdks/snis/README.md#createsni) - Create a new SNI
* [DeleteSni](docs/sdks/snis/README.md#deletesni) - Delete an SNI
* [GetSni](docs/sdks/snis/README.md#getsni) - Get an SNI
* [UpsertSni](docs/sdks/snis/README.md#upsertsni) - Upsert a SNI

### [SystemAccounts](docs/sdks/systemaccounts/README.md)

* [GetSystemAccounts](docs/sdks/systemaccounts/README.md#getsystemaccounts) - List System Accounts
* [PostSystemAccounts](docs/sdks/systemaccounts/README.md#postsystemaccounts) - Create System Account
* [GetSystemAccountsID](docs/sdks/systemaccounts/README.md#getsystemaccountsid) - Get a System Account
* [PatchSystemAccountsID](docs/sdks/systemaccounts/README.md#patchsystemaccountsid) - Update System Account
* [DeleteSystemAccountsID](docs/sdks/systemaccounts/README.md#deletesystemaccountsid) - Delete System Account

### [SystemAccountsAccessTokens](docs/sdks/systemaccountsaccesstokens/README.md)

* [GetSystemAccountIDAccessTokens](docs/sdks/systemaccountsaccesstokens/README.md#getsystemaccountidaccesstokens) - List System Account Access Tokens
* [PostSystemAccountsIDAccessTokens](docs/sdks/systemaccountsaccesstokens/README.md#postsystemaccountsidaccesstokens) - Create System Account Access Token
* [GetSystemAccountsIDAccessTokensID](docs/sdks/systemaccountsaccesstokens/README.md#getsystemaccountsidaccesstokensid) - Get a System Account Access Token
* [PatchSystemAccountsIDAccessTokensID](docs/sdks/systemaccountsaccesstokens/README.md#patchsystemaccountsidaccesstokensid) - Update System Account Access Token
* [DeleteSystemAccountsIDAccessTokensID](docs/sdks/systemaccountsaccesstokens/README.md#deletesystemaccountsidaccesstokensid) - Delete System Account Access Token

### [SystemAccountsRoles](docs/sdks/systemaccountsroles/README.md)

* [GetSystemAccountsAccountIDAssignedRoles](docs/sdks/systemaccountsroles/README.md#getsystemaccountsaccountidassignedroles) - List Assigned Roles for System Account
* [PostSystemAccountsAccountIDAssignedRoles](docs/sdks/systemaccountsroles/README.md#postsystemaccountsaccountidassignedroles) - Create Assigned Role for System Account
* [DeleteSystemAccountsAccountIDAssignedRolesRoleID](docs/sdks/systemaccountsroles/README.md#deletesystemaccountsaccountidassignedrolesroleid) - Delete Assigned Role from System Account

### [SystemAccountsTeamMembership](docs/sdks/systemaccountsteammembership/README.md)

* [GetSystemAccountsAccountIDTeams](docs/sdks/systemaccountsteammembership/README.md#getsystemaccountsaccountidteams) - List Teams for a System Account
* [GetTeamsTeamIDSystemAccounts](docs/sdks/systemaccountsteammembership/README.md#getteamsteamidsystemaccounts) - List System Accounts on a Team
* [PostTeamsTeamIDSystemAccounts](docs/sdks/systemaccountsteammembership/README.md#postteamsteamidsystemaccounts) - Add System Account to a Team
* [DeleteTeamsTeamIDSystemAccountsAccountID](docs/sdks/systemaccountsteammembership/README.md#deleteteamsteamidsystemaccountsaccountid) - Remove System Account From Team

### [Targets](docs/sdks/targets/README.md)

* [ListTargetWithUpstream](docs/sdks/targets/README.md#listtargetwithupstream) - List all Targets associated with an Upstream
* [CreateTargetWithUpstream](docs/sdks/targets/README.md#createtargetwithupstream) - Create a new Target associated with an Upstream
* [DeleteTargetWithUpstream](docs/sdks/targets/README.md#deletetargetwithupstream) - Delete a a Target associated with an Upstream
* [GetTargetWithUpstream](docs/sdks/targets/README.md#gettargetwithupstream) - Get a Target associated with an Upstream
* [UpsertTargetWithUpstream](docs/sdks/targets/README.md#upserttargetwithupstream) - Upsert a Target associated with an Upstream

### [TeamMembership](docs/sdks/teammembership/README.md)

* [ListTeamUsers](docs/sdks/teammembership/README.md#listteamusers) - List Team Users
* [AddUserToTeam](docs/sdks/teammembership/README.md#addusertoteam) - Add User
* [RemoveUserFromTeam](docs/sdks/teammembership/README.md#removeuserfromteam) - Remove User
* [ListUserTeams](docs/sdks/teammembership/README.md#listuserteams) - List User Teams

### [Teams](docs/sdks/teams/README.md)

* [ListTeams](docs/sdks/teams/README.md#listteams) - List Teams
* [CreateTeam](docs/sdks/teams/README.md#createteam) - Create Team
* [GetTeam](docs/sdks/teams/README.md#getteam) - Get a Team
* [UpdateTeam](docs/sdks/teams/README.md#updateteam) - Update Team
* [DeleteTeam](docs/sdks/teams/README.md#deleteteam) - Delete Team

### [Upstreams](docs/sdks/upstreams/README.md)

* [ListUpstream](docs/sdks/upstreams/README.md#listupstream) - List all Upstreams
* [CreateUpstream](docs/sdks/upstreams/README.md#createupstream) - Create a new Upstream
* [DeleteUpstream](docs/sdks/upstreams/README.md#deleteupstream) - Delete an Upstream
* [GetUpstream](docs/sdks/upstreams/README.md#getupstream) - Get an Upstream
* [UpsertUpstream](docs/sdks/upstreams/README.md#upsertupstream) - Upsert a Upstream

### [Users](docs/sdks/users/README.md)

* [ListUsers](docs/sdks/users/README.md#listusers) - List Users
* [GetUser](docs/sdks/users/README.md#getuser) - Get a User
* [UpdateUser](docs/sdks/users/README.md#updateuser) - Update User
* [DeleteUser](docs/sdks/users/README.md#deleteuser) - Delete User

### [Vaults](docs/sdks/vaults/README.md)

* [ListVault](docs/sdks/vaults/README.md#listvault) - List all Vaults
* [CreateVault](docs/sdks/vaults/README.md#createvault) - Create a new Vault
* [DeleteVault](docs/sdks/vaults/README.md#deletevault) - Delete a Vault
* [GetVault](docs/sdks/vaults/README.md#getvault) - Get a Vault
* [UpsertVault](docs/sdks/vaults/README.md#upsertvault) - Upsert a Vault

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Pagination [pagination] -->
## Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
```go
package main

import (
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := sdkkonnectgo.New(
		sdkkonnectgo.WithSecurity(components.Security{
			PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
		}),
	)

	res, err := s.ControlPlanes.ListControlPlanes(ctx, operations.ListControlPlanesRequest{
		PageSize:     sdkkonnectgo.Pointer[int64](10),
		PageNumber:   sdkkonnectgo.Pointer[int64](1),
		FilterLabels: sdkkonnectgo.Pointer("key:value,existCheck"),
		Sort:         sdkkonnectgo.Pointer("created_at desc"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.ListControlPlanesResponse != nil {
		for {
			// handle items

			res, err = res.Next()

			if err != nil {
				// handle error
			}

			if res == nil {
				break
			}
		}
	}
}

```
<!-- End Pagination [pagination] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/retry"
	"log"
	"models/operations"
)

func main() {
	ctx := context.Background()

	s := sdkkonnectgo.New(
		sdkkonnectgo.WithSecurity(components.Security{
			PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
		}),
	)

	res, err := s.Notifications.ListUserConfigurations(ctx, nil, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.UserConfigurationListResponse != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/retry"
	"log"
)

func main() {
	ctx := context.Background()

	s := sdkkonnectgo.New(
		sdkkonnectgo.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		sdkkonnectgo.WithSecurity(components.Security{
			PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
		}),
	)

	res, err := s.Notifications.ListUserConfigurations(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.UserConfigurationListResponse != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `sdkerrors.SDKError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `ListUserConfigurations` function may return the following errors:

| Error Type                  | Status Code | Content Type             |
| --------------------------- | ----------- | ------------------------ |
| sdkerrors.BadRequestError   | 400         | application/problem+json |
| sdkerrors.UnauthorizedError | 401         | application/problem+json |
| sdkerrors.ForbiddenError    | 403         | application/problem+json |
| sdkerrors.SDKError          | 4XX, 5XX    | \*/\*                    |

### Example

```go
package main

import (
	"context"
	"errors"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/sdkerrors"
	"log"
)

func main() {
	ctx := context.Background()

	s := sdkkonnectgo.New(
		sdkkonnectgo.WithSecurity(components.Security{
			PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
		}),
	)

	res, err := s.Notifications.ListUserConfigurations(ctx, nil)
	if err != nil {

		var e *sdkerrors.BadRequestError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.UnauthorizedError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.ForbiddenError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex(serverIndex int)` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| #   | Server                          | Description |
| --- | ------------------------------- | ----------- |
| 0   | `https://global.api.konghq.com` |             |
| 1   | `https://us.api.konghq.com`     |             |
| 2   | `https://eu.api.konghq.com`     |             |
| 3   | `https://au.api.konghq.com`     |             |

#### Example

```go
package main

import (
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := sdkkonnectgo.New(
		sdkkonnectgo.WithServerIndex(0),
		sdkkonnectgo.WithSecurity(components.Security{
			PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
		}),
	)

	res, err := s.Notifications.ListUserConfigurations(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.UserConfigurationListResponse != nil {
		// handle response
	}
}

```

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := sdkkonnectgo.New(
		sdkkonnectgo.WithServerURL("https://au.api.konghq.com"),
		sdkkonnectgo.WithSecurity(components.Security{
			PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
		}),
	)

	res, err := s.Notifications.ListUserConfigurations(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.UserConfigurationListResponse != nil {
		// handle response
	}
}

```

### Override Server URL Per-Operation

The server URL can also be overridden on a per-operation basis, provided a server list was specified for the operation. For example:
```go
package main

import (
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := sdkkonnectgo.New()

	res, err := s.CloudGateways.GetAvailabilityJSON(ctx, operations.WithServerURL("https://global.api.konghq.com/"))
	if err != nil {
		log.Fatal(err)
	}
	if res.AvailabilityDocument != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"

	"github.com/Kong/sdk-konnect-go"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdkkonnectgo.New(sdkkonnectgo.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->
