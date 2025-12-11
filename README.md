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

Konnect API - Go Internal SDK: The Konnect platform API

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
| `ClientToken`              | http | HTTP Bearer |
| `ServiceAccessToken`       | http | HTTP Bearer |

You can set the security parameters through the `WithSecurity` option when initializing the SDK client instance. The selected scheme will be used by default to authenticate with the API for all operations that support it. For example:
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

### [AIManager](docs/sdks/aimanager/README.md)

* [ListVirtualKeys](docs/sdks/aimanager/README.md#listvirtualkeys) - Get Control Plane Virtual Keys

### [Api](docs/sdks/api/README.md)

* [CreateAPI](docs/sdks/api/README.md#createapi) - Create API
* [ListApis](docs/sdks/api/README.md#listapis) - List APIs
* [FetchAPI](docs/sdks/api/README.md#fetchapi) - Get an API
* [UpdateAPI](docs/sdks/api/README.md#updateapi) - Update API
* [DeleteAPI](docs/sdks/api/README.md#deleteapi) - Delete API
* [ListApisComputed](docs/sdks/api/README.md#listapiscomputed) - List APIs computed

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

* [CreateAPIImplementation](docs/sdks/apiimplementation/README.md#createapiimplementation) - Create API Implementation
* [FetchAPIImplementation](docs/sdks/apiimplementation/README.md#fetchapiimplementation) - Get an API Implementation
* [DeleteAPIImplementation](docs/sdks/apiimplementation/README.md#deleteapiimplementation) - Delete API Implementation
* [ListAPIImplementations](docs/sdks/apiimplementation/README.md#listapiimplementations) - List API Implementations

### [APIOperations](docs/sdks/apioperations/README.md)

* [ListAPIOperations](docs/sdks/apioperations/README.md#listapioperations) - List API Operations
* [FetchAPIOperation](docs/sdks/apioperations/README.md#fetchapioperation) - Get an API Operation

### [APIPackageDocumentation](docs/sdks/apipackagedocumentation/README.md)

* [CreateAPIPackageDocument](docs/sdks/apipackagedocumentation/README.md#createapipackagedocument) - Create API Package Document
* [ListAPIPackageDocuments](docs/sdks/apipackagedocumentation/README.md#listapipackagedocuments) - List API Package Documents
* [FetchAPIPackageDocument](docs/sdks/apipackagedocumentation/README.md#fetchapipackagedocument) - Get an API Package Document
* [UpdateAPIPackageDocument](docs/sdks/apipackagedocumentation/README.md#updateapipackagedocument) - Update API Package Document
* [DeleteAPIPackageDocument](docs/sdks/apipackagedocumentation/README.md#deleteapipackagedocument) - Delete API Package Documentation
* [MoveAPIPackageDocument](docs/sdks/apipackagedocumentation/README.md#moveapipackagedocument) - Move API Package Documentation

### [APIPackageOperations](docs/sdks/apipackageoperations/README.md)

* [ListAPIPackagesOperations](docs/sdks/apipackageoperations/README.md#listapipackagesoperations) - List API Packages Operations
* [UpdateAPIPackageOperations](docs/sdks/apipackageoperations/README.md#updateapipackageoperations) - Update API Package Operation
* [GetAPIPackagesOperation](docs/sdks/apipackageoperations/README.md#getapipackagesoperation) - Get API Packages Operation
* [RemoveAPIPackageOperation](docs/sdks/apipackageoperations/README.md#removeapipackageoperation) - Remove API Package Operation

### [APIPackageSpecification](docs/sdks/apipackagespecification/README.md)

* [UpdateAPIPackageCurrentSpecification](docs/sdks/apipackagespecification/README.md#updateapipackagecurrentspecification) - Update API package current specification
* [GetAPIPackageCurrentSpecification](docs/sdks/apipackagespecification/README.md#getapipackagecurrentspecification) - Get the API package current specification
* [DeleteAPIPackageCurrentSpecification](docs/sdks/apipackagespecification/README.md#deleteapipackagecurrentspecification) - Delete API Package current Specification
* [GetAPIPackageComputedSpecification](docs/sdks/apipackagespecification/README.md#getapipackagecomputedspecification) - Get the API package computed specification

### [APIPackages](docs/sdks/apipackages/README.md)

* [ListAPIPackages](docs/sdks/apipackages/README.md#listapipackages) - List API Packages
* [CreateAPIPackage](docs/sdks/apipackages/README.md#createapipackage) - Create API Package
* [GetAPIPackage](docs/sdks/apipackages/README.md#getapipackage) - Get API Package
* [PatchAPIPackage](docs/sdks/apipackages/README.md#patchapipackage) - Patch API Package
* [DeleteAPIPackage](docs/sdks/apipackages/README.md#deleteapipackage) - Delete API Package

### [APIPublication](docs/sdks/apipublication/README.md)

* [PublishAPIToPortal](docs/sdks/apipublication/README.md#publishapitoportal) - Publish API
* [FetchPublication](docs/sdks/apipublication/README.md#fetchpublication) - Get a Publication
* [DeletePublication](docs/sdks/apipublication/README.md#deletepublication) - Delete Publication
* [ListAPIPublications](docs/sdks/apipublication/README.md#listapipublications) - List Publications
* [ListPortalAPIPublications](docs/sdks/apipublication/README.md#listportalapipublications) - List Portal API Publications
* [PublishAPIPackageToPortal](docs/sdks/apipublication/README.md#publishapipackagetoportal) - Publish API Package
* [FetchAPIPackagePublication](docs/sdks/apipublication/README.md#fetchapipackagepublication) - Get an API Package Publication
* [DeleteAPIPackagePublication](docs/sdks/apipublication/README.md#deleteapipackagepublication) - Delete API Package Publication

### [APISpecification](docs/sdks/apispecification/README.md)

* [ValidateSpecification](docs/sdks/apispecification/README.md#validatespecification) - Validate API Specification
* [~~CreateAPISpec~~](docs/sdks/apispecification/README.md#createapispec) - Create API Specification :warning: **Deprecated**
* [~~ListAPISpecs~~](docs/sdks/apispecification/README.md#listapispecs) - List API Specifications :warning: **Deprecated**
* [~~FetchAPISpec~~](docs/sdks/apispecification/README.md#fetchapispec) - Get API Specification :warning: **Deprecated**
* [~~UpdateAPISpec~~](docs/sdks/apispecification/README.md#updateapispec) - Update API Specification :warning: **Deprecated**
* [~~DeleteAPISpec~~](docs/sdks/apispecification/README.md#deleteapispec) - Delete API Specification :warning: **Deprecated**

### [APIVersion](docs/sdks/apiversion/README.md)

* [CreateAPIVersion](docs/sdks/apiversion/README.md#createapiversion) - Create API Version
* [ListAPIVersions](docs/sdks/apiversion/README.md#listapiversions) - List API Versions
* [FetchAPIVersion](docs/sdks/apiversion/README.md#fetchapiversion) - Get an API Version
* [UpdateAPIVersion](docs/sdks/apiversion/README.md#updateapiversion) - Update API Version
* [DeleteAPIVersion](docs/sdks/apiversion/README.md#deleteapiversion) - Delete API Version

### [APIKeys](docs/sdks/apikeys/README.md)

* [ListKeyAuthWithConsumer](docs/sdks/apikeys/README.md#listkeyauthwithconsumer) - List all API-keys associated with a Consumer
* [CreateKeyAuthWithConsumer](docs/sdks/apikeys/README.md#createkeyauthwithconsumer) - Create a new API-key associated with a Consumer
* [DeleteKeyAuthWithConsumer](docs/sdks/apikeys/README.md#deletekeyauthwithconsumer) - Delete a an API-key associated with a Consumer
* [GetKeyAuthWithConsumer](docs/sdks/apikeys/README.md#getkeyauthwithconsumer) - Get an API-key associated with a Consumer
* [UpsertKeyAuthWithConsumer](docs/sdks/apikeys/README.md#upsertkeyauthwithconsumer) - Upsert an API-key associated with a Consumer
* [ListKeyAuth](docs/sdks/apikeys/README.md#listkeyauth) - List all API-keys
* [GetKeyAuth](docs/sdks/apikeys/README.md#getkeyauth) - Get an API-key

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

* [ListApplications](docs/sdks/applications/README.md#listapplications) - List Applications
* [GetApplication](docs/sdks/applications/README.md#getapplication) - Get an Application by Portal
* [DeleteApplication](docs/sdks/applications/README.md#deleteapplication) - Delete Application by Portal
* [ListDevelopersByApplication](docs/sdks/applications/README.md#listdevelopersbyapplication) - List Developers by Application
* [AddDeveloperToApplication](docs/sdks/applications/README.md#adddevelopertoapplication) - Add Developer to Application
* [RemoveDeveloperFromApplication](docs/sdks/applications/README.md#removedeveloperfromapplication) - Remove Developer from Application
* [GetApplicationUnscoped](docs/sdks/applications/README.md#getapplicationunscoped) - Get an Application

### [Assets](docs/sdks/assets/README.md)

* [GetPortalAssetLogo](docs/sdks/assets/README.md#getportalassetlogo) - Get Logo
* [ReplacePortalAssetLogo](docs/sdks/assets/README.md#replaceportalassetlogo) - Replace Logo
* [GetPortalAssetLogoRaw](docs/sdks/assets/README.md#getportalassetlogoraw) - Get Logo (Raw)
* [GetPortalAssetFavicon](docs/sdks/assets/README.md#getportalassetfavicon) - Get Favicon
* [ReplacePortalAssetFavicon](docs/sdks/assets/README.md#replaceportalassetfavicon) - Replace Favicon
* [GetPortalAssetFaviconRaw](docs/sdks/assets/README.md#getportalassetfaviconraw) - Get Favicon (Raw)

### [AuthSettings](docs/sdks/authsettings/README.md)

* [GetAuthenticationSettings](docs/sdks/authsettings/README.md#getauthenticationsettings) - Get Auth Settings
* [UpdateAuthenticationSettings](docs/sdks/authsettings/README.md#updateauthenticationsettings) - Update Auth Settings
* [GetIdentityProviders](docs/sdks/authsettings/README.md#getidentityproviders) - List Identity Providers
* [CreateIdentityProvider](docs/sdks/authsettings/README.md#createidentityprovider) - Create Identity Provider
* [GetIdentityProvider](docs/sdks/authsettings/README.md#getidentityprovider) - Get Identity Provider
* [UpdateIdentityProvider](docs/sdks/authsettings/README.md#updateidentityprovider) - Update Identity Provider
* [DeleteIdentityProvider](docs/sdks/authsettings/README.md#deleteidentityprovider) - Delete Identity Provider
* [GetIdpConfiguration](docs/sdks/authsettings/README.md#getidpconfiguration) - Get the IdP Configuration
* [UpdateIdpConfiguration](docs/sdks/authsettings/README.md#updateidpconfiguration) - Update IdP Configuration
* [UpdateIdpTeamMappings](docs/sdks/authsettings/README.md#updateidpteammappings) - Update Team Mappings
* [GetIdpTeamMappings](docs/sdks/authsettings/README.md#getidpteammappings) - Get a Team Mapping
* [GetTeamGroupMappings](docs/sdks/authsettings/README.md#getteamgroupmappings) - Get a Team Group Mappings
* [PatchTeamGroupMappings](docs/sdks/authsettings/README.md#patchteamgroupmappings) - Patch Mappings by Team ID

### [Auth0](docs/sdks/auth0/README.md)

* [PostAuth0UserMfaSettingsInternal](docs/sdks/auth0/README.md#postauth0usermfasettingsinternal) - Get Auth0 User MFA Settings (Internal)

### [Authentication](docs/sdks/authentication/README.md)

* [RefreshToken](docs/sdks/authentication/README.md#refreshtoken) - Refresh Token
* [Logout](docs/sdks/authentication/README.md#logout) - Log Out
* [AuthenticateSso](docs/sdks/authentication/README.md#authenticatesso) - SSO Callback

### [Aws](docs/sdks/aws/README.md)

* [ResolveCustomer](docs/sdks/aws/README.md#resolvecustomer) - Resolves an AWS Customer

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

### [CatalogIntegrations](docs/sdks/catalogintegrations/README.md)

* [CreateCatalogIntegration](docs/sdks/catalogintegrations/README.md#createcatalogintegration) - Create Integration
* [ListCatalogIntegrations](docs/sdks/catalogintegrations/README.md#listcatalogintegrations) - List Integrations
* [GetCatalogIntegration](docs/sdks/catalogintegrations/README.md#getcatalogintegration) - Get Integration
* [UpdateCatalogIntegration](docs/sdks/catalogintegrations/README.md#updatecatalogintegration) - Update Integration
* [DeleteCatalogIntegration](docs/sdks/catalogintegrations/README.md#deletecatalogintegration) - Delete Integration

### [CatalogResourceMappings](docs/sdks/catalogresourcemappings/README.md)

* [CreateResourceMapping](docs/sdks/catalogresourcemappings/README.md#createresourcemapping) - Create Resource Mapping
* [ListResourceMappings](docs/sdks/catalogresourcemappings/README.md#listresourcemappings) - List Resource Mappings
* [FetchResourceMapping](docs/sdks/catalogresourcemappings/README.md#fetchresourcemapping) - Get a Resource Mapping
* [DeleteResourceMapping](docs/sdks/catalogresourcemappings/README.md#deleteresourcemapping) - Delete Resource Mapping

### [CatalogResourceServices](docs/sdks/catalogresourceservices/README.md)

* [ListCatalogResourceServices](docs/sdks/catalogresourceservices/README.md#listcatalogresourceservices) - List Resource Services

### [CatalogResources](docs/sdks/catalogresources/README.md)

* [ListResources](docs/sdks/catalogresources/README.md#listresources) - List Resources
* [FetchResource](docs/sdks/catalogresources/README.md#fetchresource) - Get a Resource
* [InitializeResource](docs/sdks/catalogresources/README.md#initializeresource) - Initialize Resource
* [UpsertResources](docs/sdks/catalogresources/README.md#upsertresources) - Upsert Resources
* [DeleteResources](docs/sdks/catalogresources/README.md#deleteresources) - Delete Resources
* [UpdateResource](docs/sdks/catalogresources/README.md#updateresource) - Update Resource

### [CatalogServiceAPIMappings](docs/sdks/catalogserviceapimappings/README.md)

* [ListCatalogServiceAPIMappings](docs/sdks/catalogserviceapimappings/README.md#listcatalogserviceapimappings) - List API Mappings for a Service
* [CreateCatalogServiceAPIMapping](docs/sdks/catalogserviceapimappings/README.md#createcatalogserviceapimapping) - Create API Mapping for a Service
* [GetCatalogServiceAPIMapping](docs/sdks/catalogserviceapimappings/README.md#getcatalogserviceapimapping) - Get API Mapping for a Service
* [DeleteCatalogServiceAPIMapping](docs/sdks/catalogserviceapimappings/README.md#deletecatalogserviceapimapping) - Delete API Mapping for a Service
* [ListServiceMappingsForAPI](docs/sdks/catalogserviceapimappings/README.md#listservicemappingsforapi) - List Service Mappings for an API

### [CatalogServiceAPISpecs](docs/sdks/catalogserviceapispecs/README.md)

* [CreateCatalogServiceAPISpec](docs/sdks/catalogserviceapispecs/README.md#createcatalogserviceapispec) - Create API spec
* [ListCatalogServiceAPISpecs](docs/sdks/catalogserviceapispecs/README.md#listcatalogserviceapispecs) - List Catalog Service API Specs
* [PreviewCatalogServiceAPISpec](docs/sdks/catalogserviceapispecs/README.md#previewcatalogserviceapispec) - Preview API Spec
* [FetchCatalogServiceAPISpec](docs/sdks/catalogserviceapispecs/README.md#fetchcatalogserviceapispec) - Get an API spec
* [UpdateCatalogServiceAPISpec](docs/sdks/catalogserviceapispecs/README.md#updatecatalogserviceapispec) - Update API Spec
* [DeleteCatalogServiceAPISpec](docs/sdks/catalogserviceapispecs/README.md#deletecatalogserviceapispec) - Delete API Spec
* [FetchCatalogServiceAPISpecContents](docs/sdks/catalogserviceapispecs/README.md#fetchcatalogserviceapispeccontents) - Get an API Spec Contents

### [CatalogServiceCustomFields](docs/sdks/catalogservicecustomfields/README.md)

* [ListCatalogCustomFieldSchemas](docs/sdks/catalogservicecustomfields/README.md#listcatalogcustomfieldschemas) - List Catalog Custom Field Schemas

### [CatalogServiceDocuments](docs/sdks/catalogservicedocuments/README.md)

* [CreateCatalogServiceDocument](docs/sdks/catalogservicedocuments/README.md#createcatalogservicedocument) - Create Service Document
* [ListCatalogServiceDocuments](docs/sdks/catalogservicedocuments/README.md#listcatalogservicedocuments) - List Catalog Service Documents
* [FetchCatalogServiceDocument](docs/sdks/catalogservicedocuments/README.md#fetchcatalogservicedocument) - Get a Catalog Service Document
* [UpdateCatalogServiceDocument](docs/sdks/catalogservicedocuments/README.md#updatecatalogservicedocument) - Update Catalog Service Document
* [DeleteCatalogServiceDocument](docs/sdks/catalogservicedocuments/README.md#deletecatalogservicedocument) - Delete Catalog Service Document
* [MoveCatalogServiceDocument](docs/sdks/catalogservicedocuments/README.md#movecatalogservicedocument) - Move Catalog Service Document

### [CatalogServiceResources](docs/sdks/catalogserviceresources/README.md)

* [ListCatalogServiceResources](docs/sdks/catalogserviceresources/README.md#listcatalogserviceresources) - List Service Resources

### [CatalogServices](docs/sdks/catalogservices/README.md)

* [CreateCatalogService](docs/sdks/catalogservices/README.md#createcatalogservice) - Create Service
* [ListCatalogServices](docs/sdks/catalogservices/README.md#listcatalogservices) - List Services
* [FetchCatalogService](docs/sdks/catalogservices/README.md#fetchcatalogservice) - Get a Service
* [UpdateCatalogService](docs/sdks/catalogservices/README.md#updatecatalogservice) - Update Service
* [DeleteCatalogService](docs/sdks/catalogservices/README.md#deletecatalogservice) - Delete Service

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
* [ListNetworks](docs/sdks/cloudgateways/README.md#listnetworks) - List Networks
* [CreateNetwork](docs/sdks/cloudgateways/README.md#createnetwork) - Create Network
* [GetNetwork](docs/sdks/cloudgateways/README.md#getnetwork) - Get Network
* [UpdateNetwork](docs/sdks/cloudgateways/README.md#updatenetwork) - Update Network
* [DeleteNetwork](docs/sdks/cloudgateways/README.md#deletenetwork) - Delete Network
* [ListTransitGateways](docs/sdks/cloudgateways/README.md#listtransitgateways) - List Transit Gateways
* [CreateTransitGateway](docs/sdks/cloudgateways/README.md#createtransitgateway) - Create Transit Gateway
* [GetTransitGateway](docs/sdks/cloudgateways/README.md#gettransitgateway) - Get Transit Gateway
* [UpdateTransitGateway](docs/sdks/cloudgateways/README.md#updatetransitgateway) - Update Transit Gateway
* [DeleteTransitGateway](docs/sdks/cloudgateways/README.md#deletetransitgateway) - Delete Transit Gateway
* [ListPrivateDNS](docs/sdks/cloudgateways/README.md#listprivatedns) - List Private DNS
* [CreatePrivateDNS](docs/sdks/cloudgateways/README.md#createprivatedns) - Create Private DNS
* [GetPrivateDNS](docs/sdks/cloudgateways/README.md#getprivatedns) - Get Private DNS
* [UpdatePrivateDNS](docs/sdks/cloudgateways/README.md#updateprivatedns) - Update Private DNS
* [DeletePrivateDNS](docs/sdks/cloudgateways/README.md#deleteprivatedns) - Delete Private DNS
* [ListNetworkConfigurations](docs/sdks/cloudgateways/README.md#listnetworkconfigurations) - List Network Configuration References
* [ListProviderAccounts](docs/sdks/cloudgateways/README.md#listprovideraccounts) - List Provider Accounts
* [CreateProviderAccount](docs/sdks/cloudgateways/README.md#createprovideraccount) - Create Provider Account
* [GetProviderAccount](docs/sdks/cloudgateways/README.md#getprovideraccount) - Get Provider Account
* [DeleteProviderAccount](docs/sdks/cloudgateways/README.md#deleteprovideraccount) - Delete Provider Account
* [ListCustomDomains](docs/sdks/cloudgateways/README.md#listcustomdomains) - List Custom Domains
* [CreateCustomDomains](docs/sdks/cloudgateways/README.md#createcustomdomains) - Create Custom Domain
* [GetCustomDomain](docs/sdks/cloudgateways/README.md#getcustomdomain) - Get Custom Domain
* [DeleteCustomDomain](docs/sdks/cloudgateways/README.md#deletecustomdomain) - Delete Custom Domain
* [GetCustomDomainOnlineStatus](docs/sdks/cloudgateways/README.md#getcustomdomainonlinestatus) - Get Custom Domain Online Status
* [ListDefaultResourceQuotas](docs/sdks/cloudgateways/README.md#listdefaultresourcequotas) - List Default Resource Quotas
* [ListResourceQuotas](docs/sdks/cloudgateways/README.md#listresourcequotas) - List Resource Quotas
* [CreateResourceQuota](docs/sdks/cloudgateways/README.md#createresourcequota) - Create Resource Quota
* [GetResourceQuota](docs/sdks/cloudgateways/README.md#getresourcequota) - Get Resource Quota
* [UpdateResourceQuota](docs/sdks/cloudgateways/README.md#updateresourcequota) - Update Resource Quota
* [ListDefaultResourceConfigurations](docs/sdks/cloudgateways/README.md#listdefaultresourceconfigurations) - List Default Resource Configurations
* [ListResourceConfigurations](docs/sdks/cloudgateways/README.md#listresourceconfigurations) - List Resource Configurations
* [CreateResourceConfiguration](docs/sdks/cloudgateways/README.md#createresourceconfiguration) - Create Resource Configuration
* [GetResourceConfiguration](docs/sdks/cloudgateways/README.md#getresourceconfiguration) - Get Resource Configuration
* [UpdateResourceConfiguration](docs/sdks/cloudgateways/README.md#updateresourceconfiguration) - Update Resource Configuration
* [CreateAddOn](docs/sdks/cloudgateways/README.md#createaddon) - Create Add-On
* [ListAddOns](docs/sdks/cloudgateways/README.md#listaddons) - List Add-Ons
* [GetAddOn](docs/sdks/cloudgateways/README.md#getaddon) - Get Add-On
* [DeleteAddOn](docs/sdks/cloudgateways/README.md#deleteaddon) - Delete Add-On

### [ConfigStoreSecrets](docs/sdks/configstoresecrets/README.md)

* [CreateConfigStoreSecret](docs/sdks/configstoresecrets/README.md#createconfigstoresecret) - Create Config Store Secret
* [ListConfigStoreSecrets](docs/sdks/configstoresecrets/README.md#listconfigstoresecrets) - List Config Store Secrets
* [GetConfigStoreSecret](docs/sdks/configstoresecrets/README.md#getconfigstoresecret) - Get a Config Store Secret
* [UpdateConfigStoreSecret](docs/sdks/configstoresecrets/README.md#updateconfigstoresecret) - Update Config Store Secret
* [DeleteConfigStoreSecret](docs/sdks/configstoresecrets/README.md#deleteconfigstoresecret) - Delete Config Store Secret

### [ConfigStores](docs/sdks/configstores/README.md)

* [ListConfigStores](docs/sdks/configstores/README.md#listconfigstores) - List all config stores for a control plane
* [CreateConfigStore](docs/sdks/configstores/README.md#createconfigstore) - Create Config Store
* [GetConfigStore](docs/sdks/configstores/README.md#getconfigstore) - Get a Config Store
* [UpdateConfigStore](docs/sdks/configstores/README.md#updateconfigstore) - Update an individual Config Store
* [DeleteConfigStore](docs/sdks/configstores/README.md#deleteconfigstore) - Delete Config Store

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

* [GetControlPlanesIDGroupMemberships](docs/sdks/controlplanegroups/README.md#getcontrolplanesidgroupmemberships) - List Control Plane Group Memberships
* [PutControlPlanesIDGroupMemberships](docs/sdks/controlplanegroups/README.md#putcontrolplanesidgroupmemberships) - Upsert Control Plane Group Members
* [PostControlPlanesIDGroupMembershipsAdd](docs/sdks/controlplanegroups/README.md#postcontrolplanesidgroupmembershipsadd) - Add Control Plane Group Members
* [PostControlPlanesIDGroupMembershipsRemove](docs/sdks/controlplanegroups/README.md#postcontrolplanesidgroupmembershipsremove) - Remove Control Plane Group Members
* [GetControlPlanesIDGroupMemberStatus](docs/sdks/controlplanegroups/README.md#getcontrolplanesidgroupmemberstatus) - Get Control Plane Group Member Status
* [GetControlPlanesIDGroupStatus](docs/sdks/controlplanegroups/README.md#getcontrolplanesidgroupstatus) - Get Control Plane Group Status
* [GetControlPlaneGroupSettings](docs/sdks/controlplanegroups/README.md#getcontrolplanegroupsettings) - Get Control Plane Group Settings
* [PutControlPlaneGroupSettings](docs/sdks/controlplanegroups/README.md#putcontrolplanegroupsettings) - Upsert Control Plane Group Settings

### [ControlPlaneResourceQuotas](docs/sdks/controlplaneresourcequotas/README.md)

* [ListControlPlaneDefaultResourceQuotas](docs/sdks/controlplaneresourcequotas/README.md#listcontrolplanedefaultresourcequotas) - List Default Quotas
* [ListControlPlaneResourceQuota](docs/sdks/controlplaneresourcequotas/README.md#listcontrolplaneresourcequota) - List Quota Overrides
* [CreateControlPlaneResourceQuota](docs/sdks/controlplaneresourcequotas/README.md#createcontrolplaneresourcequota) - Create a control plane resource quota
* [GetControlPlaneResourceQuota](docs/sdks/controlplaneresourcequotas/README.md#getcontrolplaneresourcequota) - Get Control Plane Resource Quota
* [UpdateControlPlaneResourceQuota](docs/sdks/controlplaneresourcequotas/README.md#updatecontrolplaneresourcequota) - Update Resource Quota
* [DeleteControlPlaneResourceQuota](docs/sdks/controlplaneresourcequotas/README.md#deletecontrolplaneresourcequota) - Delete the control plane resource quota

### [ControlPlanes](docs/sdks/controlplanes/README.md)

* [ListControlPlanes](docs/sdks/controlplanes/README.md#listcontrolplanes) - List Control Planes
* [CreateControlPlane](docs/sdks/controlplanes/README.md#createcontrolplane) - Create Control Plane
* [GetControlPlane](docs/sdks/controlplanes/README.md#getcontrolplane) - Get a Control Plane
* [UpdateControlPlane](docs/sdks/controlplanes/README.md#updatecontrolplane) - Update Control Plane
* [DeleteControlPlane](docs/sdks/controlplanes/README.md#deletecontrolplane) - Delete Control Plane

### [CriteriaTemplates](docs/sdks/criteriatemplates/README.md)

* [ListCriteriaTemplates](docs/sdks/criteriatemplates/README.md#listcriteriatemplates) - List Criteria Templates

### [CustomPluginSchemas](docs/sdks/custompluginschemas/README.md)

* [ListPluginSchemas](docs/sdks/custompluginschemas/README.md#listpluginschemas) - List Custom Plugin Schemas
* [CreatePluginSchemas](docs/sdks/custompluginschemas/README.md#createpluginschemas) - Upload custom plugin schema
* [GetPluginSchema](docs/sdks/custompluginschemas/README.md#getpluginschema) - Get a custom plugin schema
* [DeletePluginSchemas](docs/sdks/custompluginschemas/README.md#deletepluginschemas) - Delete custom plugin schema
* [UpdatePluginSchemas](docs/sdks/custompluginschemas/README.md#updatepluginschemas) - Create or update a custom plugin schema

### [CustomPlugins](docs/sdks/customplugins/README.md)

* [ListCustomPlugin](docs/sdks/customplugins/README.md#listcustomplugin) - List all CustomPlugins
* [CreateCustomPlugin](docs/sdks/customplugins/README.md#createcustomplugin) - Create a new CustomPlugin
* [DeleteCustomPlugin](docs/sdks/customplugins/README.md#deletecustomplugin) - Delete a CustomPlugin
* [GetCustomPlugin](docs/sdks/customplugins/README.md#getcustomplugin) - Get a CustomPlugin
* [UpsertCustomPlugin](docs/sdks/customplugins/README.md#upsertcustomplugin) - Upsert a CustomPlugin

### [DCRProviders](docs/sdks/dcrproviders/README.md)

* [CreateDcrProvider](docs/sdks/dcrproviders/README.md#createdcrprovider) - Create DCR provider
* [ListDcrProviders](docs/sdks/dcrproviders/README.md#listdcrproviders) - List DCR Providers
* [GetDcrProvider](docs/sdks/dcrproviders/README.md#getdcrprovider) - Get a DCR provider
* [UpdateDcrProvider](docs/sdks/dcrproviders/README.md#updatedcrprovider) - Update DCR provider
* [DeleteDcrProvider](docs/sdks/dcrproviders/README.md#deletedcrprovider) - Delete DCR provider
* [VerifyDcrProvider](docs/sdks/dcrproviders/README.md#verifydcrprovider) - Verify DCR provider configuration

### [DebugSessions](docs/sdks/debugsessions/README.md)

* [ListDebugSessions](docs/sdks/debugsessions/README.md#listdebugsessions) - List all debug sessions for a control plane
* [CreateDebugSession](docs/sdks/debugsessions/README.md#createdebugsession) - Create Debug Session
* [GetDebugSession](docs/sdks/debugsessions/README.md#getdebugsession) - Fetch a Debug Session
* [DeleteDebugSession](docs/sdks/debugsessions/README.md#deletedebugsession) - Delete a Debug Session
* [StopDebugSession](docs/sdks/debugsessions/README.md#stopdebugsession) - Stops an active Debug Session

### [DeclarativeConfiguration](docs/sdks/declarativeconfiguration/README.md)

* [UpsertDeclarativeConfig](docs/sdks/declarativeconfiguration/README.md#upsertdeclarativeconfig) - Create or Update the declarative config
* [GetDeclarativeConfig](docs/sdks/declarativeconfiguration/README.md#getdeclarativeconfig) - Get the declarative configuration
* [DeleteDeclarativeConfig](docs/sdks/declarativeconfiguration/README.md#deletedeclarativeconfig) - Delete the declarative config
* [GetNativeEventProxyDeclarativeConfig](docs/sdks/declarativeconfiguration/README.md#getnativeeventproxydeclarativeconfig) - Get Event Proxy Configuration
* [GetHTTPGatewayDeclarativeConfig](docs/sdks/declarativeconfiguration/README.md#gethttpgatewaydeclarativeconfig) - Get HTTP Proxy Configuration
* [GetHoudiniEventGatewayDeclarativeConfig](docs/sdks/declarativeconfiguration/README.md#gethoudinieventgatewaydeclarativeconfig) - Get Houdini Event Gateway Configuration

### [DegraphqlRoutes](docs/sdks/degraphqlroutes/README.md)

* [ListDegraphqlRoute](docs/sdks/degraphqlroutes/README.md#listdegraphqlroute) - List all Degraphql_routes
* [GetDegraphqlRoute](docs/sdks/degraphqlroutes/README.md#getdegraphqlroute) - Get a Degraphql_route
* [ListDegraphqlRouteWithService](docs/sdks/degraphqlroutes/README.md#listdegraphqlroutewithservice) - List all Degraphql_routes associated with a Service
* [CreateDegraphqlRouteWithService](docs/sdks/degraphqlroutes/README.md#createdegraphqlroutewithservice) - Create a new Degraphql_route associated with a Service
* [DeleteDegraphqlRouteWithService](docs/sdks/degraphqlroutes/README.md#deletedegraphqlroutewithservice) - Delete a a Degraphql_route associated with a Service
* [GetDegraphqlRouteWithService](docs/sdks/degraphqlroutes/README.md#getdegraphqlroutewithservice) - Get a Degraphql_route associated with a Service
* [UpsertDegraphqlRouteWithService](docs/sdks/degraphqlroutes/README.md#upsertdegraphqlroutewithservice) - Upsert a Degraphql_route associated with a Service

### [DeviceAuthorizationGrant](docs/sdks/deviceauthorizationgrant/README.md)

* [PostOauthDeviceAuthorize](docs/sdks/deviceauthorizationgrant/README.md#postoauthdeviceauthorize) - Device authorization request
* [PostOauthDeviceToken](docs/sdks/deviceauthorizationgrant/README.md#postoauthdevicetoken) - Device access token request
* [PostOauthDeviceAuthorizeUser](docs/sdks/deviceauthorizationgrant/README.md#postoauthdeviceauthorizeuser) - User device authorization request
* [PatchOauthDeviceConfirm](docs/sdks/deviceauthorizationgrant/README.md#patchoauthdeviceconfirm) - Device confirmation request

### [DPCertificates](docs/sdks/dpcertificates/README.md)

* [ListDpClientCertificates](docs/sdks/dpcertificates/README.md#listdpclientcertificates) - List DP Client Certificates
* [CreateDataplaneCertificate](docs/sdks/dpcertificates/README.md#createdataplanecertificate) - Pin New DP Client Certificate
* [GetDataplaneCertificate](docs/sdks/dpcertificates/README.md#getdataplanecertificate) - Get a DP Client Certificate
* [DeleteDataplaneCertificate](docs/sdks/dpcertificates/README.md#deletedataplanecertificate) - Delete DP Client Certificate

### [DPNodes](docs/sdks/dpnodes/README.md)

* [GetExpectedConfigHash](docs/sdks/dpnodes/README.md#getexpectedconfighash) - Get an Expected Config Hash
* [GetExpectedConfigVersion](docs/sdks/dpnodes/README.md#getexpectedconfigversion) - Get an Expected Config Version
* [ListDataplaneNodes](docs/sdks/dpnodes/README.md#listdataplanenodes) - List Data Plane Node Records
* [GetNodesEol](docs/sdks/dpnodes/README.md#getnodeseol) - List End-of-Life Data Plane Node Records
* [GetNodesNodeID](docs/sdks/dpnodes/README.md#getnodesnodeid) - Get a Data Plane Node Record
* [DeleteNodesNodeID](docs/sdks/dpnodes/README.md#deletenodesnodeid) - Delete Data Plane Node Record

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

### [IntegrationEvents](docs/sdks/integrationevents/README.md)

* [ListIntegrationEvents](docs/sdks/integrationevents/README.md#listintegrationevents) - List Integration Events
* [CreateIntegrationEvents](docs/sdks/integrationevents/README.md#createintegrationevents) - Bulk Create Integration Events
* [ListCatalogServiceIntegrationEvents](docs/sdks/integrationevents/README.md#listcatalogserviceintegrationevents) - List a catalog service integration Events

### [IntegrationInstanceAuthConfig](docs/sdks/integrationinstanceauthconfig/README.md)

* [GetIntegrationInstanceAuthConfig](docs/sdks/integrationinstanceauthconfig/README.md#getintegrationinstanceauthconfig) - Get Integration Instance Auth Config
* [UpsertIntegrationInstanceAuthConfig](docs/sdks/integrationinstanceauthconfig/README.md#upsertintegrationinstanceauthconfig) - Upsert Integration Instance Auth Config
* [DeleteIntegrationInstanceAuthConfig](docs/sdks/integrationinstanceauthconfig/README.md#deleteintegrationinstanceauthconfig) - Delete Integration Instance Auth Config

### [IntegrationInstanceAuthCredentials](docs/sdks/integrationinstanceauthcredentials/README.md)

* [CreateIntegrationInstanceAuthCredential](docs/sdks/integrationinstanceauthcredentials/README.md#createintegrationinstanceauthcredential) - Create Integration Instance Auth Credential
* [GetIntegrationInstanceAuthCredential](docs/sdks/integrationinstanceauthcredentials/README.md#getintegrationinstanceauthcredential) - Get Integration Instance Auth Credential
* [DeleteIntegrationInstanceAuthCredential](docs/sdks/integrationinstanceauthcredentials/README.md#deleteintegrationinstanceauthcredential) - Delete Integration Instance Auth Credential

### [IntegrationInstanceProxy](docs/sdks/integrationinstanceproxy/README.md)

* [IntegrationInstanceProxyRequest](docs/sdks/integrationinstanceproxy/README.md#integrationinstanceproxyrequest) - Integration Instance Proxy Request

### [IntegrationInstances](docs/sdks/integrationinstances/README.md)

* [CreateIntegrationInstance](docs/sdks/integrationinstances/README.md#createintegrationinstance) - Create Integration Instance
* [ListIntegrationInstances](docs/sdks/integrationinstances/README.md#listintegrationinstances) - List Integration Instances
* [FetchIntegrationInstance](docs/sdks/integrationinstances/README.md#fetchintegrationinstance) - Get an Integration Instance
* [UpdateIntegrationInstance](docs/sdks/integrationinstances/README.md#updateintegrationinstance) - Update Integration Instance
* [DeleteIntegrationInstance](docs/sdks/integrationinstances/README.md#deleteintegrationinstance) - Delete Integration Instance

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

### [ManagedSystemAccountsRoles](docs/sdks/managedsystemaccountsroles/README.md)

* [GetSystemAccountsAssignedRolesInternal](docs/sdks/managedsystemaccountsroles/README.md#getsystemaccountsassignedrolesinternal) - List Roles (Internal)
* [CreateSystemAccountsAssignedRolesInternal](docs/sdks/managedsystemaccountsroles/README.md#createsystemaccountsassignedrolesinternal) - Assign a role to a managed System Account

### [Me](docs/sdks/me/README.md)

* [GetUsersMe](docs/sdks/me/README.md#getusersme) - Get My User Account
* [DeleteUsersMe](docs/sdks/me/README.md#deleteusersme) - Delete My User Account
* [PatchUsersMe](docs/sdks/me/README.md#patchusersme) - Update My User Account
* [GetUsersMePermissions](docs/sdks/me/README.md#getusersmepermissions) - Get My Permissions
* [GetOrganizationsMe](docs/sdks/me/README.md#getorganizationsme) - Get My Organization
* [UpdateOrganizationsMe](docs/sdks/me/README.md#updateorganizationsme) - Update My Organization

### [MTLSAuthCredentials](docs/sdks/mtlsauthcredentials/README.md)

* [ListMtlsAuthWithConsumer](docs/sdks/mtlsauthcredentials/README.md#listmtlsauthwithconsumer) - List all MTLS-auth credentials associated with a Consumer
* [CreateMtlsAuthWithConsumer](docs/sdks/mtlsauthcredentials/README.md#createmtlsauthwithconsumer) - Create a new MTLS-auth credential associated with a Consumer
* [DeleteMtlsAuthWithConsumer](docs/sdks/mtlsauthcredentials/README.md#deletemtlsauthwithconsumer) - Delete a a MTLS-auth credential associated with a Consumer
* [GetMtlsAuthWithConsumer](docs/sdks/mtlsauthcredentials/README.md#getmtlsauthwithconsumer) - Get a MTLS-auth credential associated with a Consumer
* [UpsertMtlsAuthWithConsumer](docs/sdks/mtlsauthcredentials/README.md#upsertmtlsauthwithconsumer) - Upsert a MTLS-auth credential associated with a Consumer
* [ListMtlsAuth](docs/sdks/mtlsauthcredentials/README.md#listmtlsauth) - List all MTLS-auth credentials
* [GetMtlsAuth](docs/sdks/mtlsauthcredentials/README.md#getmtlsauth) - Get a MTLS-auth credential

### [Nodes](docs/sdks/nodes/README.md)

* [UpsertNode](docs/sdks/nodes/README.md#upsertnode) - Create or update a node

### [Notifications](docs/sdks/notifications/README.md)

* [ListNotifications](docs/sdks/notifications/README.md#listnotifications) - List available notifications
* [GetNotificationDetails](docs/sdks/notifications/README.md#getnotificationdetails) - Get notification details
* [UpdateNotification](docs/sdks/notifications/README.md#updatenotification) - Update notification
* [DeleteNotification](docs/sdks/notifications/README.md#deletenotification) - Delete notification
* [BulkNotifications](docs/sdks/notifications/README.md#bulknotifications) - Mark a list of notifications to a status
* [ListUserConfigurations](docs/sdks/notifications/README.md#listuserconfigurations) - List available user configurations
* [ListEventSubscriptions](docs/sdks/notifications/README.md#listeventsubscriptions) - List event subscriptions
* [CreateEventSubscription](docs/sdks/notifications/README.md#createeventsubscription) - Create a new subscription for an event
* [GetEventSubscription](docs/sdks/notifications/README.md#geteventsubscription) - Get subscription for an event
* [UpdateEventSubscription](docs/sdks/notifications/README.md#updateeventsubscription) - Update subscription for an event
* [DeleteEventSubscription](docs/sdks/notifications/README.md#deleteeventsubscription) - Delete subscription associated with event

### [OrganizationFeature](docs/sdks/organizationfeature/README.md)

* [GetOrganizationFeature](docs/sdks/organizationfeature/README.md#getorganizationfeature) - Get Feature Configuration
* [UpsertOrganizationFeature](docs/sdks/organizationfeature/README.md#upsertorganizationfeature) - Upsert Feature Configuration

### [Pages](docs/sdks/pages/README.md)

* [ListPortalPages](docs/sdks/pages/README.md#listportalpages) - List Pages
* [CreatePortalPage](docs/sdks/pages/README.md#createportalpage) - Create Page
* [GetPortalPage](docs/sdks/pages/README.md#getportalpage) - Get a Page
* [UpdatePortalPage](docs/sdks/pages/README.md#updateportalpage) - Update Page
* [DeletePortalPage](docs/sdks/pages/README.md#deleteportalpage) - Delete Page
* [MovePortalPages](docs/sdks/pages/README.md#moveportalpages) - Move Page
* [CreateDefaultContent](docs/sdks/pages/README.md#createdefaultcontent) - Creates Default Pages

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

* [ListEmailDomains](docs/sdks/portalemails/README.md#listemaildomains) - List email domains
* [CreateEmailDomain](docs/sdks/portalemails/README.md#createemaildomain) - Create an email domain
* [GetEmailDomain](docs/sdks/portalemails/README.md#getemaildomain) - Get an email domain
* [DeleteEmailDomain](docs/sdks/portalemails/README.md#deleteemaildomain) - Delete an email domain
* [~~GetEmailDelivery~~](docs/sdks/portalemails/README.md#getemaildelivery) - Get the email delivery for the portal :warning: **Deprecated**
* [~~UpdateEmailDelivery~~](docs/sdks/portalemails/README.md#updateemaildelivery) - Setup the email delivery for a portal :warning: **Deprecated**
* [~~DeleteEmailDelivery~~](docs/sdks/portalemails/README.md#deleteemaildelivery) - Delete email delivery :warning: **Deprecated**
* [GetEmailConfig](docs/sdks/portalemails/README.md#getemailconfig) - Get the email config for the portal
* [CreatePortalEmailConfig](docs/sdks/portalemails/README.md#createportalemailconfig) - Create the email config for a portal
* [UpdatePortalEmailConfig](docs/sdks/portalemails/README.md#updateportalemailconfig) - Setup the email config for a portal
* [DeletePortalEmailConfig](docs/sdks/portalemails/README.md#deleteportalemailconfig) - Delete portal email config
* [ListDefaultEmailTemplates](docs/sdks/portalemails/README.md#listdefaultemailtemplates) - List default email templates
* [ListEmailTemplateVariables](docs/sdks/portalemails/README.md#listemailtemplatevariables) - List email template variables
* [GetDefaultEmailTemplate](docs/sdks/portalemails/README.md#getdefaultemailtemplate) - Get default email template
* [ListPortalCustomEmailTemplates](docs/sdks/portalemails/README.md#listportalcustomemailtemplates) - List custom email templates for a portal
* [GetPortalCustomEmailTemplate](docs/sdks/portalemails/README.md#getportalcustomemailtemplate) - Get custom email template used in a portal
* [UpdatePortalCustomEmailTemplate](docs/sdks/portalemails/README.md#updateportalcustomemailtemplate) - Update custom email template for a portal
* [DeletePortalCustomEmailTemplate](docs/sdks/portalemails/README.md#deleteportalcustomemailtemplate) - Delete custom email template
* [PostPortalCustomEmailTestSend](docs/sdks/portalemails/README.md#postportalcustomemailtestsend) - Send Test Email

### [PortalTeamMembership](docs/sdks/portalteammembership/README.md)

* [ListPortalTeamDevelopers](docs/sdks/portalteammembership/README.md#listportalteamdevelopers) - List Team Developers
* [AddDeveloperToPortalTeam](docs/sdks/portalteammembership/README.md#adddevelopertoportalteam) - Add Developer to Team
* [RemoveDeveloperFromPortalTeam](docs/sdks/portalteammembership/README.md#removedeveloperfromportalteam) - Remove Developer from Team
* [ListPortalDeveloperTeams](docs/sdks/portalteammembership/README.md#listportaldeveloperteams) - List Developer Teams

### [PortalTeamRoles](docs/sdks/portalteamroles/README.md)

* [ListPortalTeamRoles](docs/sdks/portalteamroles/README.md#listportalteamroles) - List Team Roles
* [AssignRoleToPortalTeams](docs/sdks/portalteamroles/README.md#assignroletoportalteams) - Assign Role
* [RemoveRoleFromPortalTeam](docs/sdks/portalteamroles/README.md#removerolefromportalteam) - Remove Role
* [ListPortalRoles](docs/sdks/portalteamroles/README.md#listportalroles) - List Portal Roles

### [PortalTeams](docs/sdks/portalteams/README.md)

* [ListPortalTeams](docs/sdks/portalteams/README.md#listportalteams) - List Teams
* [CreatePortalTeam](docs/sdks/portalteams/README.md#createportalteam) - Create Team
* [GetPortalTeam](docs/sdks/portalteams/README.md#getportalteam) - Get Team
* [UpdatePortalTeam](docs/sdks/portalteams/README.md#updateportalteam) - Update Team
* [DeletePortalTeam](docs/sdks/portalteams/README.md#deleteportalteam) - Delete Team

### [Portals](docs/sdks/portals/README.md)

* [ListPortals](docs/sdks/portals/README.md#listportals) - List Portals
* [CreatePortal](docs/sdks/portals/README.md#createportal) - Create Portal
* [GetPortal](docs/sdks/portals/README.md#getportal) - Get a Portal
* [UpdatePortal](docs/sdks/portals/README.md#updateportal) - Update Portal
* [DeletePortal](docs/sdks/portals/README.md#deleteportal) - Delete Portal

### [ResourceActions](docs/sdks/resourceactions/README.md)

* [ListResourceActions](docs/sdks/resourceactions/README.md#listresourceactions) - List Resource Actions

### [ResourceIngestion](docs/sdks/resourceingestion/README.md)

* [ScheduleResourceIngestion](docs/sdks/resourceingestion/README.md#scheduleresourceingestion) - Schedule Resource Ingestion
* [FetchResourceIngestion](docs/sdks/resourceingestion/README.md#fetchresourceingestion) - Get a Resource Ingestion

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

### [Scorecards](docs/sdks/scorecards/README.md)

* [ListScorecardTemplates](docs/sdks/scorecards/README.md#listscorecardtemplates) - List Scorecard Templates
* [CreateScorecard](docs/sdks/scorecards/README.md#createscorecard) - Create Scorecard
* [ListScorecards](docs/sdks/scorecards/README.md#listscorecards) - List Scorecards
* [FetchScorecard](docs/sdks/scorecards/README.md#fetchscorecard) - Get a Scorecard
* [UpdateScorecard](docs/sdks/scorecards/README.md#updatescorecard) - Update Scorecard
* [DeleteScorecard](docs/sdks/scorecards/README.md#deletescorecard) - Delete Scorecard
* [ScheduleScorecardEvaluation](docs/sdks/scorecards/README.md#schedulescorecardevaluation) - Schedule Scorecard Evaluation
* [ListScorecardServices](docs/sdks/scorecards/README.md#listscorecardservices) - List Scorecard Services
* [ListScorecardCriteria](docs/sdks/scorecards/README.md#listscorecardcriteria) - List Scorecard Criteria
* [ListScorecardCriteriaServices](docs/sdks/scorecards/README.md#listscorecardcriteriaservices) - List Scorecard Criteria Services
* [ListCatalogServiceScorecards](docs/sdks/scorecards/README.md#listcatalogservicescorecards) - List Catalog Service Scorecards
* [FetchCatalogServiceScorecard](docs/sdks/scorecards/README.md#fetchcatalogservicescorecard) - Get a Catalog Service Scorecard

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

### [SSOAuth0](docs/sdks/ssoauth0/README.md)

* [PostAuth0RegisterInternal](docs/sdks/ssoauth0/README.md#postauth0registerinternal) - Create Organization (Internal)
* [GetAuth0OrganizationsInternal](docs/sdks/ssoauth0/README.md#getauth0organizationsinternal) - List Organizations (Internal)

### [SuggestedResourceActions](docs/sdks/suggestedresourceactions/README.md)

* [ListSuggestedResourceAction](docs/sdks/suggestedresourceactions/README.md#listsuggestedresourceaction) - List Suggested Resource Actions
* [GetSuggestedResourceAction](docs/sdks/suggestedresourceactions/README.md#getsuggestedresourceaction) - Get a Suggested Resource Action
* [UpdateSuggestedResourceAction](docs/sdks/suggestedresourceactions/README.md#updatesuggestedresourceaction) - Update Suggested Resource Action

### [SuggestionRuleErrors](docs/sdks/suggestionruleerrors/README.md)

* [ListSuggestionRuleErrors](docs/sdks/suggestionruleerrors/README.md#listsuggestionruleerrors) - List Suggestion Rule Errors

### [SuggestionRules](docs/sdks/suggestionrules/README.md)

* [ListSystemIntegrationSuggestionRule](docs/sdks/suggestionrules/README.md#listsystemintegrationsuggestionrule) - List System Suggestion Rules
* [CreateIntegrationSuggestionRule](docs/sdks/suggestionrules/README.md#createintegrationsuggestionrule) - Create Suggestion Rule
* [ListIntegrationSuggestionRule](docs/sdks/suggestionrules/README.md#listintegrationsuggestionrule) - List Suggestion Rules
* [GetIntegrationSuggestionRule](docs/sdks/suggestionrules/README.md#getintegrationsuggestionrule) - Get a Suggestion Rule
* [UpdateIntegrationSuggestionRule](docs/sdks/suggestionrules/README.md#updateintegrationsuggestionrule) - Update Suggestion Rule
* [DeleteIntegrationSuggestionRule](docs/sdks/suggestionrules/README.md#deleteintegrationsuggestionrule) - Delete Suggestion Rule
* [TestSuggestionRule](docs/sdks/suggestionrules/README.md#testsuggestionrule) - Test a Suggestion Rule Configuration

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

* [GetTeamsTeamIDSystemAccounts](docs/sdks/systemaccountsteammembership/README.md#getteamsteamidsystemaccounts) - List System Accounts on a Team
* [PostTeamsTeamIDSystemAccounts](docs/sdks/systemaccountsteammembership/README.md#postteamsteamidsystemaccounts) - Add System Account to a Team
* [DeleteTeamsTeamIDSystemAccountsAccountID](docs/sdks/systemaccountsteammembership/README.md#deleteteamsteamidsystemaccountsaccountid) - Remove System Account From Team
* [GetSystemAccountsAccountIDTeams](docs/sdks/systemaccountsteammembership/README.md#getsystemaccountsaccountidteams) - List Teams for a System Account

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

* [GetUsersInternal](docs/sdks/users/README.md#getusersinternal) - List Users (Internal)
* [ListUsers](docs/sdks/users/README.md#listusers) - List Users
* [GetUser](docs/sdks/users/README.md#getuser) - Get a User
* [UpdateUser](docs/sdks/users/README.md#updateuser) - Update User
* [DeleteUser](docs/sdks/users/README.md#deleteuser) - Delete User
* [DeleteUserMfas](docs/sdks/users/README.md#deleteusermfas) - Delete User MFA

### [Vaults](docs/sdks/vaults/README.md)

* [ListVault](docs/sdks/vaults/README.md#listvault) - List all Vaults
* [CreateVault](docs/sdks/vaults/README.md#createvault) - Create a new Vault
* [DeleteVault](docs/sdks/vaults/README.md#deletevault) - Delete a Vault
* [GetVault](docs/sdks/vaults/README.md#getvault) - Get a Vault
* [UpsertVault](docs/sdks/vaults/README.md#upsertvault) - Upsert a Vault

### [Vulnerabilities](docs/sdks/vulnerabilities/README.md)

* [ListVulnerabilityScans](docs/sdks/vulnerabilities/README.md#listvulnerabilityscans) - List Vulnerability Scans
* [CreateVulnerabilityScan](docs/sdks/vulnerabilities/README.md#createvulnerabilityscan) - Create Vulnerability Scan
* [FetchVulnerabilityScan](docs/sdks/vulnerabilities/README.md#fetchvulnerabilityscan) - Get a Vulnerability Scan
* [ListVulnerabilityScanVulnerabilities](docs/sdks/vulnerabilities/README.md#listvulnerabilityscanvulnerabilities) - List Vulnerability Scan Vulnerabilities
* [ListServiceVulnerabilityScans](docs/sdks/vulnerabilities/README.md#listservicevulnerabilityscans) - List Service Vulnerability Scans
* [FetchServiceVulnerabilityScan](docs/sdks/vulnerabilities/README.md#fetchservicevulnerabilityscan) - Get a Service Vulnerability Scan
* [ListVulnerabilities](docs/sdks/vulnerabilities/README.md#listvulnerabilities) - List Vulnerabilities
* [FetchVulnerability](docs/sdks/vulnerabilities/README.md#fetchvulnerability) - Get Vulnerability
* [ListVulnerabilityServices](docs/sdks/vulnerabilities/README.md#listvulnerabilityservices) - List Vulnerability Services
* [ListVulnerabilityInstances](docs/sdks/vulnerabilities/README.md#listvulnerabilityinstances) - List Vulnerability Instances
* [ListServiceVulnerabilities](docs/sdks/vulnerabilities/README.md#listservicevulnerabilities) - List Service Vulnerabilities
* [FetchServiceVulnerability](docs/sdks/vulnerabilities/README.md#fetchservicevulnerability) - Get Service Vulnerability
* [ListServiceVulnerabilityInstances](docs/sdks/vulnerabilities/README.md#listservicevulnerabilityinstances) - List Service Vulnerability Instances
* [PutServiceVulnerabilitySeverityOverride](docs/sdks/vulnerabilities/README.md#putservicevulnerabilityseverityoverride) - Put Service Vulnerability Severity Override
* [DeleteServiceVulnerabilitySeverityOverride](docs/sdks/vulnerabilities/README.md#deleteservicevulnerabilityseverityoverride) - Delete Vulnerability Severity Override
* [PutServiceVulnerabilityDismissal](docs/sdks/vulnerabilities/README.md#putservicevulnerabilitydismissal) - Put Service Vulnerability Dismissal
* [DeleteServiceVulnerabilityDismissal](docs/sdks/vulnerabilities/README.md#deleteservicevulnerabilitydismissal) - Delete Service Vulnerability Dismissal
* [ListCatalogVulnerabilityServices](docs/sdks/vulnerabilities/README.md#listcatalogvulnerabilityservices) - List Catalog Vulnerability-Services
* [QueryVulnerabilitiesMetrics](docs/sdks/vulnerabilities/README.md#queryvulnerabilitiesmetrics) - Query Vulnerabilities Metrics

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
	"github.com/Kong/sdk-konnect-go/models/operations"
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

	res, err := s.ControlPlanes.ListControlPlanes(ctx, operations.ListControlPlanesRequest{
		PageSize:     sdkkonnectgo.Pointer[int64](10),
		PageNumber:   sdkkonnectgo.Pointer[int64](1),
		FilterLabels: sdkkonnectgo.Pointer("key:value,existCheck"),
		Sort:         sdkkonnectgo.Pointer("created_at desc"),
	}, operations.WithRetries(
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

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
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
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `sdkerrors.SDKError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `ListControlPlanes` function may return the following errors:

| Error Type                   | Status Code | Content Type             |
| ---------------------------- | ----------- | ------------------------ |
| sdkerrors.BadRequestError    | 400         | application/problem+json |
| sdkerrors.UnauthorizedError  | 401         | application/problem+json |
| sdkerrors.ForbiddenError     | 403         | application/problem+json |
| sdkerrors.BaseError          | 500         | application/problem+json |
| sdkerrors.ServiceUnavailable | 503         | application/problem+json |
| sdkerrors.SDKError           | 4XX, 5XX    | \*/\*                    |

### Example

```go
package main

import (
	"context"
	"errors"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
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

	res, err := s.ControlPlanes.ListControlPlanes(ctx, operations.ListControlPlanesRequest{
		PageSize:     sdkkonnectgo.Pointer[int64](10),
		PageNumber:   sdkkonnectgo.Pointer[int64](1),
		FilterLabels: sdkkonnectgo.Pointer("key:value,existCheck"),
		Sort:         sdkkonnectgo.Pointer("created_at desc"),
	})
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

		var e *sdkerrors.BaseError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.ServiceUnavailable
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
	"github.com/Kong/sdk-konnect-go/models/operations"
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

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
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
		sdkkonnectgo.WithServerURL("https://au.api.konghq.com"),
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

### Override Server URL Per-Operation

The server URL can also be overridden on a per-operation basis, provided a server list was specified for the operation. For example:
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

	res, err := s.SSOAuth0.PostAuth0RegisterInternal(ctx, nil, operations.WithServerURL("https://global.api.konghq.com/"))
	if err != nil {
		log.Fatal(err)
	}
	if res.Auth0Registration != nil {
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
