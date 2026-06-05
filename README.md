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

* [ListACLInWorkspace](docs/sdks/acls/README.md#listaclinworkspace) - List all ACLs in a workspace
* [GetACLInWorkspace](docs/sdks/acls/README.md#getaclinworkspace) - Get an ACL in a workspace
* [ListACLWithConsumerInWorkspace](docs/sdks/acls/README.md#listaclwithconsumerinworkspace) - List all ACLs associated with a Consumer in a workspace
* [CreateACLWithConsumerInWorkspace](docs/sdks/acls/README.md#createaclwithconsumerinworkspace) - Create a new ACL associated with a Consumer in a workspace
* [DeleteACLWithConsumerInWorkspace](docs/sdks/acls/README.md#deleteaclwithconsumerinworkspace) - Delete a an ACL associated with a Consumer in a workspace
* [GetACLWithConsumerInWorkspace](docs/sdks/acls/README.md#getaclwithconsumerinworkspace) - Get an ACL associated with a Consumer in a workspace
* [UpsertACLWithConsumerInWorkspace](docs/sdks/acls/README.md#upsertaclwithconsumerinworkspace) - Upsert an ACL associated with a Consumer in a workspace
* [ListACL](docs/sdks/acls/README.md#listacl) - List all ACLs
* [GetACL](docs/sdks/acls/README.md#getacl) - Get an ACL
* [ListACLWithConsumer](docs/sdks/acls/README.md#listaclwithconsumer) - List all ACLs associated with a Consumer
* [CreateACLWithConsumer](docs/sdks/acls/README.md#createaclwithconsumer) - Create a new ACL associated with a Consumer
* [DeleteACLWithConsumer](docs/sdks/acls/README.md#deleteaclwithconsumer) - Delete a an ACL associated with a Consumer
* [GetACLWithConsumer](docs/sdks/acls/README.md#getaclwithconsumer) - Get an ACL associated with a Consumer
* [UpsertACLWithConsumer](docs/sdks/acls/README.md#upsertaclwithconsumer) - Upsert an ACL associated with a Consumer

### [AIGatewayAgents](docs/sdks/aigatewayagents/README.md)

* [ListAiGatewayAgents](docs/sdks/aigatewayagents/README.md#listaigatewayagents) - List AI Gateway Agents
* [CreateAiGatewayAgent](docs/sdks/aigatewayagents/README.md#createaigatewayagent) - Create an AI Gateway Agent
* [GetAiGatewayAgent](docs/sdks/aigatewayagents/README.md#getaigatewayagent) - Get an AI Gateway Agent
* [UpdateAiGatewayAgent](docs/sdks/aigatewayagents/README.md#updateaigatewayagent) - Update an AI Gateway Agent
* [DeleteAiGatewayAgent](docs/sdks/aigatewayagents/README.md#deleteaigatewayagent) - Delete an AI Gateway Agent

### [AIGatewayConsumerGroups](docs/sdks/aigatewayconsumergroups/README.md)

* [ListAiGatewayConsumerGroups](docs/sdks/aigatewayconsumergroups/README.md#listaigatewayconsumergroups) - List AI Gateway Consumer Groups
* [CreateAiGatewayConsumerGroup](docs/sdks/aigatewayconsumergroups/README.md#createaigatewayconsumergroup) - Create an AI Gateway Consumer Group
* [GetAiGatewayConsumerGroup](docs/sdks/aigatewayconsumergroups/README.md#getaigatewayconsumergroup) - Get an AI Gateway Consumer Group
* [UpdateAiGatewayConsumerGroup](docs/sdks/aigatewayconsumergroups/README.md#updateaigatewayconsumergroup) - Update an AI Gateway Consumer Group
* [DeleteAiGatewayConsumerGroup](docs/sdks/aigatewayconsumergroups/README.md#deleteaigatewayconsumergroup) - Delete an AI Gateway Consumer Group
* [ListAiGatewayConsumersInConsumerGroup](docs/sdks/aigatewayconsumergroups/README.md#listaigatewayconsumersinconsumergroup) - List AI Gateway Consumers in a Consumer Group
* [AddAiGatewayConsumerToConsumerGroup](docs/sdks/aigatewayconsumergroups/README.md#addaigatewayconsumertoconsumergroup) - Add a Consumer to a Consumer Group
* [RemoveAiGatewayConsumerFromConsumerGroup](docs/sdks/aigatewayconsumergroups/README.md#removeaigatewayconsumerfromconsumergroup) - Remove a Consumer from a Consumer Group

### [AIGatewayConsumers](docs/sdks/aigatewayconsumers/README.md)

* [ListAiGatewayConsumers](docs/sdks/aigatewayconsumers/README.md#listaigatewayconsumers) - List AI Gateway Consumers
* [CreateAiGatewayConsumer](docs/sdks/aigatewayconsumers/README.md#createaigatewayconsumer) - Create an AI Gateway Consumer
* [GetAiGatewayConsumer](docs/sdks/aigatewayconsumers/README.md#getaigatewayconsumer) - Get an AI Gateway Consumer
* [UpdateAiGatewayConsumer](docs/sdks/aigatewayconsumers/README.md#updateaigatewayconsumer) - Update an AI Gateway Consumer
* [DeleteAiGatewayConsumer](docs/sdks/aigatewayconsumers/README.md#deleteaigatewayconsumer) - Delete an AI Gateway Consumer
* [ListAiGatewayConsumerCredentials](docs/sdks/aigatewayconsumers/README.md#listaigatewayconsumercredentials) - List AI Gateway Consumer Credentials
* [CreateAiGatewayConsumerCredential](docs/sdks/aigatewayconsumers/README.md#createaigatewayconsumercredential) - Create an AI Gateway Consumer Credential
* [GetAiGatewayConsumerCredential](docs/sdks/aigatewayconsumers/README.md#getaigatewayconsumercredential) - Get an AI Gateway Consumer Credential
* [DeleteAiGatewayConsumerCredential](docs/sdks/aigatewayconsumers/README.md#deleteaigatewayconsumercredential) - Delete an AI Gateway Consumer Credential

### [AIGatewayDataPlaneCertificates](docs/sdks/aigatewaydataplanecertificates/README.md)

* [ListAiGatewayDataPlaneCertificates](docs/sdks/aigatewaydataplanecertificates/README.md#listaigatewaydataplanecertificates) - List AI Gateway DataPlane Certificates
* [CreateAiGatewayDataPlaneCertificate](docs/sdks/aigatewaydataplanecertificates/README.md#createaigatewaydataplanecertificate) - Create New AI Gateway DataPlane Certificate
* [GetAiGatewayDataPlaneCertificate](docs/sdks/aigatewaydataplanecertificates/README.md#getaigatewaydataplanecertificate) - Get a DataPlane Certificate
* [DeleteAiGatewayDataPlaneCertificate](docs/sdks/aigatewaydataplanecertificates/README.md#deleteaigatewaydataplanecertificate) - Delete AI Gateway DataPlane Certificate

### [AIGatewayMCPServers](docs/sdks/aigatewaymcpservers/README.md)

* [ListAiGatewayMcpServers](docs/sdks/aigatewaymcpservers/README.md#listaigatewaymcpservers) - List MCP Servers
* [CreateAiGatewayMcpServer](docs/sdks/aigatewaymcpservers/README.md#createaigatewaymcpserver) - Create an MCP Server
* [GetAiGatewayMcpServer](docs/sdks/aigatewaymcpservers/README.md#getaigatewaymcpserver) - Get an MCP Server
* [UpdateAiGatewayMcpServer](docs/sdks/aigatewaymcpservers/README.md#updateaigatewaymcpserver) - Update an MCP Server
* [DeleteAiGatewayMcpServer](docs/sdks/aigatewaymcpservers/README.md#deleteaigatewaymcpserver) - Delete an MCP Server

### [AIGatewayModels](docs/sdks/aigatewaymodels/README.md)

* [ListAiGatewayModels](docs/sdks/aigatewaymodels/README.md#listaigatewaymodels) - List AI Gateway Models
* [CreateAiGatewayModel](docs/sdks/aigatewaymodels/README.md#createaigatewaymodel) - Create an AI Gateway Model
* [GetAiGatewayModel](docs/sdks/aigatewaymodels/README.md#getaigatewaymodel) - Get an AI Gateway Model
* [UpdateAiGatewayModel](docs/sdks/aigatewaymodels/README.md#updateaigatewaymodel) - Update an AI Gateway Model
* [DeleteAiGatewayModel](docs/sdks/aigatewaymodels/README.md#deleteaigatewaymodel) - Delete an AI Gateway Model

### [AIGatewayNodes](docs/sdks/aigatewaynodes/README.md)

* [ListAiGatewayNodes](docs/sdks/aigatewaynodes/README.md#listaigatewaynodes) - List Nodes
* [GetAiGatewayNode](docs/sdks/aigatewaynodes/README.md#getaigatewaynode) - Get a Node

### [AIGatewayPolicies](docs/sdks/aigatewaypolicies/README.md)

* [ListAiGatewayAvailablePolicies](docs/sdks/aigatewaypolicies/README.md#listaigatewayavailablepolicies) - List AI Gateway Available Policies
* [GetAiGatewayPolicySchema](docs/sdks/aigatewaypolicies/README.md#getaigatewaypolicyschema) - Get an AI Gateway Policy Schema
* [ListAiGatewayPolicies](docs/sdks/aigatewaypolicies/README.md#listaigatewaypolicies) - List AI Gateway Policies
* [CreateAiGatewayPolicy](docs/sdks/aigatewaypolicies/README.md#createaigatewaypolicy) - Create an AI Gateway Policy
* [GetAiGatewayPolicy](docs/sdks/aigatewaypolicies/README.md#getaigatewaypolicy) - Get an AI Gateway Policy
* [UpdateAiGatewayPolicy](docs/sdks/aigatewaypolicies/README.md#updateaigatewaypolicy) - Update an AI Gateway Policy
* [DeleteAiGatewayPolicy](docs/sdks/aigatewaypolicies/README.md#deleteaigatewaypolicy) - Delete an AI Gateway Policy
* [ListAiGatewayPolicyUsage](docs/sdks/aigatewaypolicies/README.md#listaigatewaypolicyusage) - List AI Gateway Policy Usage

### [AIGatewayProviders](docs/sdks/aigatewayproviders/README.md)

* [ListAiGatewayProviders](docs/sdks/aigatewayproviders/README.md#listaigatewayproviders) - List AI Gateway Providers
* [CreateAiGatewayProvider](docs/sdks/aigatewayproviders/README.md#createaigatewayprovider) - Create an AI Gateway Provider
* [GetAiGatewayProvider](docs/sdks/aigatewayproviders/README.md#getaigatewayprovider) - Get an AI Gateway Provider
* [UpdateAiGatewayProvider](docs/sdks/aigatewayproviders/README.md#updateaigatewayprovider) - Update an AI Gateway Provider
* [DeleteAiGatewayProvider](docs/sdks/aigatewayproviders/README.md#deleteaigatewayprovider) - Delete an AI Gateway Provider

### [AIGatewayVaults](docs/sdks/aigatewayvaults/README.md)

* [ListAiGatewayVaults](docs/sdks/aigatewayvaults/README.md#listaigatewayvaults) - List AI Gateway Vaults
* [CreateAiGatewayVault](docs/sdks/aigatewayvaults/README.md#createaigatewayvault) - Create an AI Gateway Vault
* [GetAiGatewayVault](docs/sdks/aigatewayvaults/README.md#getaigatewayvault) - Get an AI Gateway Vault
* [UpdateAiGatewayVault](docs/sdks/aigatewayvaults/README.md#updateaigatewayvault) - Update an AI Gateway Vault
* [DeleteAiGatewayVault](docs/sdks/aigatewayvaults/README.md#deleteaigatewayvault) - Delete an AI Gateway Vault

### [AIGateways](docs/sdks/aigateways/README.md)

* [ListAiGateways](docs/sdks/aigateways/README.md#listaigateways) - List AI Gateways
* [CreateAiGateway](docs/sdks/aigateways/README.md#createaigateway) - Create an AI Gateway
* [GetAiGateway](docs/sdks/aigateways/README.md#getaigateway) - Get an AI Gateway
* [UpdateAiGateway](docs/sdks/aigateways/README.md#updateaigateway) - Update an AI Gateway
* [DeleteAiGateway](docs/sdks/aigateways/README.md#deleteaigateway) - Delete an AI Gateway

### [AIManager](docs/sdks/aimanager/README.md)

* [ListVirtualKeys](docs/sdks/aimanager/README.md#listvirtualkeys) - Get Control Plane Virtual Keys

### [Api](docs/sdks/api/README.md)

* [CreateAPI](docs/sdks/api/README.md#createapi) - Create API
* [ListApis](docs/sdks/api/README.md#listapis) - List APIs
* [FetchAPI](docs/sdks/api/README.md#fetchapi) - Get an API
* [UpdateAPI](docs/sdks/api/README.md#updateapi) - Update API
* [DeleteAPI](docs/sdks/api/README.md#deleteapi) - Delete API
* [ListApisComputed](docs/sdks/api/README.md#listapiscomputed) - List APIs computed
* [ListAPIRegistrations](docs/sdks/api/README.md#listapiregistrations) - List API Registrations

### [APIAttributes](docs/sdks/apiattributes/README.md)

* [ListAPIAttributes](docs/sdks/apiattributes/README.md#listapiattributes) - List API Attributes

### [APIDocumentation](docs/sdks/apidocumentation/README.md)

* [CreateAPIDocument](docs/sdks/apidocumentation/README.md#createapidocument) - Create API Document
* [ListAPIDocuments](docs/sdks/apidocumentation/README.md#listapidocuments) - List API Documents
* [FetchAPIDocument](docs/sdks/apidocumentation/README.md#fetchapidocument) - Get an API Document
* [UpdateAPIDocument](docs/sdks/apidocumentation/README.md#updateapidocument) - Update API Document
* [DeleteAPIDocument](docs/sdks/apidocumentation/README.md#deleteapidocument) - Delete API Documentation
* [MoveAPIDocument](docs/sdks/apidocumentation/README.md#moveapidocument) - Move API Documentation

### [APIImage](docs/sdks/apiimage/README.md)

* [UpsertAPIImage](docs/sdks/apiimage/README.md#upsertapiimage) - Create or Replace an API Image
* [FetchAPIImage](docs/sdks/apiimage/README.md#fetchapiimage) - Get API Image
* [DeleteAPIImage](docs/sdks/apiimage/README.md#deleteapiimage) - Delete an API Image
* [FetchAPIRawImage](docs/sdks/apiimage/README.md#fetchapirawimage) - Get an API Raw Image

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

### [APIPackageImage](docs/sdks/apipackageimage/README.md)

* [UpsertAPIPackageImage](docs/sdks/apipackageimage/README.md#upsertapipackageimage) - Create or Replace an API Package Image
* [FetchAPIPackageImage](docs/sdks/apipackageimage/README.md#fetchapipackageimage) - Get API Package Image
* [DeleteAPIPackageImage](docs/sdks/apipackageimage/README.md#deleteapipackageimage) - Delete an API Package Image
* [FetchAPIPackageRawImage](docs/sdks/apipackageimage/README.md#fetchapipackagerawimage) - Get an API Package Raw Image

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

### [APISpecsPreview](docs/sdks/apispecspreview/README.md)

* [PreviewAPISpec](docs/sdks/apispecspreview/README.md#previewapispec) - Preview API Spec

### [APIVersion](docs/sdks/apiversion/README.md)

* [CreateAPIVersion](docs/sdks/apiversion/README.md#createapiversion) - Create API Version
* [ListAPIVersions](docs/sdks/apiversion/README.md#listapiversions) - List API Versions
* [FetchAPIVersion](docs/sdks/apiversion/README.md#fetchapiversion) - Get an API Version
* [UpdateAPIVersion](docs/sdks/apiversion/README.md#updateapiversion) - Update API Version
* [DeleteAPIVersion](docs/sdks/apiversion/README.md#deleteapiversion) - Delete API Version

### [APIKeys](docs/sdks/apikeys/README.md)

* [ListKeyAuthWithConsumerInWorkspace](docs/sdks/apikeys/README.md#listkeyauthwithconsumerinworkspace) - List all API-keys associated with a Consumer in a workspace
* [CreateKeyAuthWithConsumerInWorkspace](docs/sdks/apikeys/README.md#createkeyauthwithconsumerinworkspace) - Create a new API-key associated with a Consumer in a workspace
* [DeleteKeyAuthWithConsumerInWorkspace](docs/sdks/apikeys/README.md#deletekeyauthwithconsumerinworkspace) - Delete a an API-key associated with a Consumer in a workspace
* [GetKeyAuthWithConsumerInWorkspace](docs/sdks/apikeys/README.md#getkeyauthwithconsumerinworkspace) - Get an API-key associated with a Consumer in a workspace
* [UpsertKeyAuthWithConsumerInWorkspace](docs/sdks/apikeys/README.md#upsertkeyauthwithconsumerinworkspace) - Upsert an API-key associated with a Consumer in a workspace
* [ListKeyAuthInWorkspace](docs/sdks/apikeys/README.md#listkeyauthinworkspace) - List all API-keys in a workspace
* [GetKeyAuthInWorkspace](docs/sdks/apikeys/README.md#getkeyauthinworkspace) - Get an API-key in a workspace
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
* [CreateApplicationRegistration](docs/sdks/applicationregistrations/README.md#createapplicationregistration) - Create Registration
* [ListRegistrationsByApplication](docs/sdks/applicationregistrations/README.md#listregistrationsbyapplication) - List Registrations by Application
* [GetApplicationRegistration](docs/sdks/applicationregistrations/README.md#getapplicationregistration) - Get a Registration
* [UpdateApplicationRegistration](docs/sdks/applicationregistrations/README.md#updateapplicationregistration) - Update Registration
* [DeleteApplicationRegistration](docs/sdks/applicationregistrations/README.md#deleteapplicationregistration) - Delete Registration
* [UpdateApplicationRegistrationConsumer](docs/sdks/applicationregistrations/README.md#updateapplicationregistrationconsumer) - Update application consumer mapping
* [DeleteApplicationRegistrationConsumer](docs/sdks/applicationregistrations/README.md#deleteapplicationregistrationconsumer) - Delete application consumer mapping

### [Applications](docs/sdks/applications/README.md)

* [CreateApplication](docs/sdks/applications/README.md#createapplication) - Create Application
* [ListApplications](docs/sdks/applications/README.md#listapplications) - List Applications
* [GetApplication](docs/sdks/applications/README.md#getapplication) - Get an Application by Portal
* [UpdateApplication](docs/sdks/applications/README.md#updateapplication) - Update Application
* [DeleteApplication](docs/sdks/applications/README.md#deleteapplication) - Delete Application by Portal
* [ListDevelopersByApplication](docs/sdks/applications/README.md#listdevelopersbyapplication) - List Developers by Application
* [AddDeveloperToApplication](docs/sdks/applications/README.md#adddevelopertoapplication) - Add Developer to Application
* [RemoveDeveloperFromApplication](docs/sdks/applications/README.md#removedeveloperfromapplication) - Remove Developer from Application
* [GetApplicationUnscoped](docs/sdks/applications/README.md#getapplicationunscoped) - Get an Application
* [ListCredentialsByApplication](docs/sdks/applications/README.md#listcredentialsbyapplication) - List Credentials by Application

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
* [ListIdpTeamGroupMappings](docs/sdks/authsettings/README.md#listidpteamgroupmappings) - List Team Group Mappings
* [CreateIdpTeamGroupMapping](docs/sdks/authsettings/README.md#createidpteamgroupmapping) - Create Team Group Mapping
* [GetIdpTeamGroupMapping](docs/sdks/authsettings/README.md#getidpteamgroupmapping) - Get Team Group Mapping
* [DeleteIdpTeamGroupMapping](docs/sdks/authsettings/README.md#deleteidpteamgroupmapping) - Delete Team Group Mapping
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

* [ListBasicAuthInWorkspace](docs/sdks/basicauthcredentials/README.md#listbasicauthinworkspace) - List all Basic-auth credentials in a workspace
* [GetBasicAuthInWorkspace](docs/sdks/basicauthcredentials/README.md#getbasicauthinworkspace) - Get a Basic-auth credential in a workspace
* [ListBasicAuthWithConsumerInWorkspace](docs/sdks/basicauthcredentials/README.md#listbasicauthwithconsumerinworkspace) - List all Basic-auth credentials associated with a Consumer in a workspace
* [CreateBasicAuthWithConsumerInWorkspace](docs/sdks/basicauthcredentials/README.md#createbasicauthwithconsumerinworkspace) - Create a new Basic-auth credential associated with a Consumer in a workspace
* [DeleteBasicAuthWithConsumerInWorkspace](docs/sdks/basicauthcredentials/README.md#deletebasicauthwithconsumerinworkspace) - Delete a a Basic-auth credential associated with a Consumer in a workspace
* [GetBasicAuthWithConsumerInWorkspace](docs/sdks/basicauthcredentials/README.md#getbasicauthwithconsumerinworkspace) - Get a Basic-auth credential associated with a Consumer in a workspace
* [UpsertBasicAuthWithConsumerInWorkspace](docs/sdks/basicauthcredentials/README.md#upsertbasicauthwithconsumerinworkspace) - Upsert a Basic-auth credential associated with a Consumer in a workspace
* [ListBasicAuth](docs/sdks/basicauthcredentials/README.md#listbasicauth) - List all Basic-auth credentials
* [GetBasicAuth](docs/sdks/basicauthcredentials/README.md#getbasicauth) - Get a Basic-auth credential
* [ListBasicAuthWithConsumer](docs/sdks/basicauthcredentials/README.md#listbasicauthwithconsumer) - List all Basic-auth credentials associated with a Consumer
* [CreateBasicAuthWithConsumer](docs/sdks/basicauthcredentials/README.md#createbasicauthwithconsumer) - Create a new Basic-auth credential associated with a Consumer
* [DeleteBasicAuthWithConsumer](docs/sdks/basicauthcredentials/README.md#deletebasicauthwithconsumer) - Delete a a Basic-auth credential associated with a Consumer
* [GetBasicAuthWithConsumer](docs/sdks/basicauthcredentials/README.md#getbasicauthwithconsumer) - Get a Basic-auth credential associated with a Consumer
* [UpsertBasicAuthWithConsumer](docs/sdks/basicauthcredentials/README.md#upsertbasicauthwithconsumer) - Upsert a Basic-auth credential associated with a Consumer

### [CACertificates](docs/sdks/cacertificates/README.md)

* [ListCaCertificateInWorkspace](docs/sdks/cacertificates/README.md#listcacertificateinworkspace) - List all CA Certificates in a workspace
* [CreateCaCertificateInWorkspace](docs/sdks/cacertificates/README.md#createcacertificateinworkspace) - Create a new CA Certificate in a workspace
* [DeleteCaCertificateInWorkspace](docs/sdks/cacertificates/README.md#deletecacertificateinworkspace) - Delete a CA Certificate in a workspace
* [GetCaCertificateInWorkspace](docs/sdks/cacertificates/README.md#getcacertificateinworkspace) - Get a CA Certificate in a workspace
* [UpsertCaCertificateInWorkspace](docs/sdks/cacertificates/README.md#upsertcacertificateinworkspace) - Upsert a CA Certificate in a workspace
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

* [ListCertificateInWorkspace](docs/sdks/certificates/README.md#listcertificateinworkspace) - List all Certificates in a workspace
* [CreateCertificateInWorkspace](docs/sdks/certificates/README.md#createcertificateinworkspace) - Create a new Certificate in a workspace
* [DeleteCertificateInWorkspace](docs/sdks/certificates/README.md#deletecertificateinworkspace) - Delete a Certificate in a workspace
* [GetCertificateInWorkspace](docs/sdks/certificates/README.md#getcertificateinworkspace) - Get a Certificate in a workspace
* [UpsertCertificateInWorkspace](docs/sdks/certificates/README.md#upsertcertificateinworkspace) - Upsert a Certificate in a workspace
* [ListCertificate](docs/sdks/certificates/README.md#listcertificate) - List all Certificates
* [CreateCertificate](docs/sdks/certificates/README.md#createcertificate) - Create a new Certificate
* [DeleteCertificate](docs/sdks/certificates/README.md#deletecertificate) - Delete a Certificate
* [GetCertificate](docs/sdks/certificates/README.md#getcertificate) - Get a Certificate
* [UpsertCertificate](docs/sdks/certificates/README.md#upsertcertificate) - Upsert a Certificate

### [ClonedPlugins](docs/sdks/clonedplugins/README.md)

* [ListClonedPlugin](docs/sdks/clonedplugins/README.md#listclonedplugin) - List all Cloned Plugins
* [CreateClonedPlugin](docs/sdks/clonedplugins/README.md#createclonedplugin) - Create a new Cloned Plugin
* [GetClonedPlugin](docs/sdks/clonedplugins/README.md#getclonedplugin) - Get a Cloned Plugin
* [UpsertClonedPlugin](docs/sdks/clonedplugins/README.md#upsertclonedplugin) - Upsert a Cloned Plugin
* [DeleteClonedPlugin](docs/sdks/clonedplugins/README.md#deleteclonedplugin) - Delete a Cloned Plugin

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
* [UpdateAddOn](docs/sdks/cloudgateways/README.md#updateaddon) - Update Add-On

### [ConfigStoreSecrets](docs/sdks/configstoresecrets/README.md)

* [CreateConfigStoreSecret](docs/sdks/configstoresecrets/README.md#createconfigstoresecret) - Create Config Store Secret
* [ListConfigStoreSecrets](docs/sdks/configstoresecrets/README.md#listconfigstoresecrets) - List Config Store Secrets
* [GetConfigStoreSecret](docs/sdks/configstoresecrets/README.md#getconfigstoresecret) - Get a Config Store Secret
* [UpdateConfigStoreSecret](docs/sdks/configstoresecrets/README.md#updateconfigstoresecret) - Update Config Store Secret
* [DeleteConfigStoreSecret](docs/sdks/configstoresecrets/README.md#deleteconfigstoresecret) - Delete Config Store Secret
* [CreateConfigStoreSecretInWorkspace](docs/sdks/configstoresecrets/README.md#createconfigstoresecretinworkspace) - Create Config Store Secret in a workspace
* [ListConfigStoreSecretsInWorkspace](docs/sdks/configstoresecrets/README.md#listconfigstoresecretsinworkspace) - List Config Store Secrets in a workspace
* [GetConfigStoreSecretInWorkspace](docs/sdks/configstoresecrets/README.md#getconfigstoresecretinworkspace) - Get a Config Store Secret in a workspace
* [UpdateConfigStoreSecretInWorkspace](docs/sdks/configstoresecrets/README.md#updateconfigstoresecretinworkspace) - Update Config Store Secret in a workspace
* [DeleteConfigStoreSecretInWorkspace](docs/sdks/configstoresecrets/README.md#deleteconfigstoresecretinworkspace) - Delete Config Store Secret in a workspace

### [ConfigStores](docs/sdks/configstores/README.md)

* [ListConfigStores](docs/sdks/configstores/README.md#listconfigstores) - List all config stores for a control plane
* [CreateConfigStore](docs/sdks/configstores/README.md#createconfigstore) - Create Config Store
* [GetConfigStore](docs/sdks/configstores/README.md#getconfigstore) - Get a Config Store
* [UpdateConfigStore](docs/sdks/configstores/README.md#updateconfigstore) - Update an individual Config Store
* [DeleteConfigStore](docs/sdks/configstores/README.md#deleteconfigstore) - Delete Config Store
* [ListConfigStoresInWorkspace](docs/sdks/configstores/README.md#listconfigstoresinworkspace) - List all Config Stores for a workspace
* [CreateConfigStoreInWorkspace](docs/sdks/configstores/README.md#createconfigstoreinworkspace) - Create Config Store in a workspace
* [GetConfigStoreInWorkspace](docs/sdks/configstores/README.md#getconfigstoreinworkspace) - Get a Config Store in a workspace
* [UpdateConfigStoreInWorkspace](docs/sdks/configstores/README.md#updateconfigstoreinworkspace) - Update a Config Store in a workspace
* [DeleteConfigStoreInWorkspace](docs/sdks/configstores/README.md#deleteconfigstoreinworkspace) - Delete Config Store in a workspace

### [ConsumerGroups](docs/sdks/consumergroups/README.md)

* [ListConsumerGroupInWorkspace](docs/sdks/consumergroups/README.md#listconsumergroupinworkspace) - List all Consumer Groups in a workspace
* [CreateConsumerGroupInWorkspace](docs/sdks/consumergroups/README.md#createconsumergroupinworkspace) - Create a new Consumer Group in a workspace
* [DeleteConsumerGroupInWorkspace](docs/sdks/consumergroups/README.md#deleteconsumergroupinworkspace) - Delete a Consumer Group in a workspace
* [GetConsumerGroupInWorkspace](docs/sdks/consumergroups/README.md#getconsumergroupinworkspace) - Get a Consumer Group in a workspace
* [UpsertConsumerGroupInWorkspace](docs/sdks/consumergroups/README.md#upsertconsumergroupinworkspace) - Upsert a Consumer Group in a workspace
* [RemoveAllConsumersFromConsumerGroupInWorkspace](docs/sdks/consumergroups/README.md#removeallconsumersfromconsumergroupinworkspace) - Remove consumers from consumer group in a workspace
* [ListConsumersForConsumerGroupInWorkspace](docs/sdks/consumergroups/README.md#listconsumersforconsumergroupinworkspace) - List all Consumers in a Consumer Group in a workspace
* [AddConsumerToGroupInWorkspace](docs/sdks/consumergroups/README.md#addconsumertogroupinworkspace) - Add consumer to consumer group in a workspace
* [RemoveConsumerFromGroupInWorkspace](docs/sdks/consumergroups/README.md#removeconsumerfromgroupinworkspace) - Remove consumer from consumer group in a workspace
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

* [ListConsumerInWorkspace](docs/sdks/consumers/README.md#listconsumerinworkspace) - List all Consumers in a workspace
* [CreateConsumerInWorkspace](docs/sdks/consumers/README.md#createconsumerinworkspace) - Create a new Consumer in a workspace
* [DeleteConsumerInWorkspace](docs/sdks/consumers/README.md#deleteconsumerinworkspace) - Delete a Consumer in a workspace
* [GetConsumerInWorkspace](docs/sdks/consumers/README.md#getconsumerinworkspace) - Get a Consumer in a workspace
* [UpsertConsumerInWorkspace](docs/sdks/consumers/README.md#upsertconsumerinworkspace) - Upsert a Consumer in a workspace
* [RemoveConsumerFromAllConsumerGroupsInWorkspace](docs/sdks/consumers/README.md#removeconsumerfromallconsumergroupsinworkspace) - Remove consumer from all consumer groups in a workspace
* [ListConsumerGroupsForConsumerInWorkspace](docs/sdks/consumers/README.md#listconsumergroupsforconsumerinworkspace) - List all Consumer Groups a Consumer belongs to in a workspace
* [AddConsumerToSpecificConsumerGroupInWorkspace](docs/sdks/consumers/README.md#addconsumertospecificconsumergroupinworkspace) - Add consumer to a specific consumer group in a workspace
* [RemoveConsumerFromConsumerGroupInWorkspace](docs/sdks/consumers/README.md#removeconsumerfromconsumergroupinworkspace) - Remove consumer from consumer group in a workspace
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
* [ListPluginSchemasInWorkspace](docs/sdks/custompluginschemas/README.md#listpluginschemasinworkspace) - List Custom Plugin Schemas in a workspace
* [CreatePluginSchemasInWorkspace](docs/sdks/custompluginschemas/README.md#createpluginschemasinworkspace) - Upload custom plugin schema in a workspace
* [GetPluginSchemaInWorkspace](docs/sdks/custompluginschemas/README.md#getpluginschemainworkspace) - Get a custom plugin schema in a workspace
* [DeletePluginSchemasInWorkspace](docs/sdks/custompluginschemas/README.md#deletepluginschemasinworkspace) - Delete custom plugin schema in a workspace
* [UpdatePluginSchemasInWorkspace](docs/sdks/custompluginschemas/README.md#updatepluginschemasinworkspace) - Create or update a custom plugin schema in a workspace

### [CustomPlugins](docs/sdks/customplugins/README.md)

* [ListCustomPluginInWorkspace](docs/sdks/customplugins/README.md#listcustomplugininworkspace) - List all CustomPlugins in a workspace
* [CreateCustomPluginInWorkspace](docs/sdks/customplugins/README.md#createcustomplugininworkspace) - Create a new CustomPlugin in a workspace
* [DeleteCustomPluginInWorkspace](docs/sdks/customplugins/README.md#deletecustomplugininworkspace) - Delete a CustomPlugin in a workspace
* [GetCustomPluginInWorkspace](docs/sdks/customplugins/README.md#getcustomplugininworkspace) - Get a CustomPlugin in a workspace
* [UpsertCustomPluginInWorkspace](docs/sdks/customplugins/README.md#upsertcustomplugininworkspace) - Upsert a CustomPlugin in a workspace
* [ListCustomPlugin](docs/sdks/customplugins/README.md#listcustomplugin) - List all CustomPlugins
* [CreateCustomPlugin](docs/sdks/customplugins/README.md#createcustomplugin) - Create a new CustomPlugin
* [DeleteCustomPlugin](docs/sdks/customplugins/README.md#deletecustomplugin) - Delete a CustomPlugin
* [GetCustomPlugin](docs/sdks/customplugins/README.md#getcustomplugin) - Get a CustomPlugin
* [UpsertCustomPlugin](docs/sdks/customplugins/README.md#upsertcustomplugin) - Upsert a CustomPlugin

### [Dashboards](docs/sdks/dashboards/README.md)

* [DashboardsList](docs/sdks/dashboards/README.md#dashboardslist) - List dashboards
* [DashboardsCreate](docs/sdks/dashboards/README.md#dashboardscreate) - Create a new dashboard
* [DashboardsGet](docs/sdks/dashboards/README.md#dashboardsget) - Get a single dashboard
* [DashboardsUpdate](docs/sdks/dashboards/README.md#dashboardsupdate) - Update an existing dashboard
* [DashboardsDelete](docs/sdks/dashboards/README.md#dashboardsdelete) - Delete an existing dashboard

### [DCRProviders](docs/sdks/dcrproviders/README.md)

* [CreateDcrProvider](docs/sdks/dcrproviders/README.md#createdcrprovider) - Create DCR provider
* [ListDcrProviders](docs/sdks/dcrproviders/README.md#listdcrproviders) - List DCR Providers
* [GetDcrProvider](docs/sdks/dcrproviders/README.md#getdcrprovider) - Get a DCR provider
* [UpdateDcrProvider](docs/sdks/dcrproviders/README.md#updatedcrprovider) - Update DCR provider
* [DeleteDcrProvider](docs/sdks/dcrproviders/README.md#deletedcrprovider) - Delete DCR provider
* [VerifyDcrProvider](docs/sdks/dcrproviders/README.md#verifydcrprovider) - Verify DCR provider configuration

### [~~DebugSessions~~](docs/sdks/debugsessions/README.md)

* [~~ListDebugSessions~~](docs/sdks/debugsessions/README.md#listdebugsessions) - List all debug sessions for a control plane :warning: **Deprecated**
* [~~CreateDebugSession~~](docs/sdks/debugsessions/README.md#createdebugsession) - Create Debug Session :warning: **Deprecated**
* [~~GetDebugSession~~](docs/sdks/debugsessions/README.md#getdebugsession) - Fetch a Debug Session :warning: **Deprecated**
* [~~DeleteDebugSession~~](docs/sdks/debugsessions/README.md#deletedebugsession) - Delete a Debug Session :warning: **Deprecated**
* [~~StopDebugSession~~](docs/sdks/debugsessions/README.md#stopdebugsession) - Stops an active Debug Session :warning: **Deprecated**

### [DeclarativeConfiguration](docs/sdks/declarativeconfiguration/README.md)

* [UpsertDeclarativeConfig](docs/sdks/declarativeconfiguration/README.md#upsertdeclarativeconfig) - Create or Update the declarative config
* [GetDeclarativeConfig](docs/sdks/declarativeconfiguration/README.md#getdeclarativeconfig) - Get the declarative configuration
* [DeleteDeclarativeConfig](docs/sdks/declarativeconfiguration/README.md#deletedeclarativeconfig) - Delete the declarative config
* [GetNativeEventProxyDeclarativeConfig](docs/sdks/declarativeconfiguration/README.md#getnativeeventproxydeclarativeconfig) - Get Event Proxy Configuration
* [GetHTTPGatewayDeclarativeConfig](docs/sdks/declarativeconfiguration/README.md#gethttpgatewaydeclarativeconfig) - Get HTTP Proxy Configuration
* [GetHoudiniEventGatewayDeclarativeConfig](docs/sdks/declarativeconfiguration/README.md#gethoudinieventgatewaydeclarativeconfig) - Get Houdini Event Gateway Configuration

### [DegraphqlRoutes](docs/sdks/degraphqlroutes/README.md)

* [ListDegraphqlRouteInWorkspace](docs/sdks/degraphqlroutes/README.md#listdegraphqlrouteinworkspace) - List all Degraphql_routes in a workspace
* [GetDegraphqlRouteInWorkspace](docs/sdks/degraphqlroutes/README.md#getdegraphqlrouteinworkspace) - Get a Degraphql_route in a workspace
* [ListDegraphqlRouteWithServiceInWorkspace](docs/sdks/degraphqlroutes/README.md#listdegraphqlroutewithserviceinworkspace) - List all Degraphql_routes associated with a Service in a workspace
* [CreateDegraphqlRouteWithServiceInWorkspace](docs/sdks/degraphqlroutes/README.md#createdegraphqlroutewithserviceinworkspace) - Create a new Degraphql_route associated with a Service in a workspace
* [DeleteDegraphqlRouteWithServiceInWorkspace](docs/sdks/degraphqlroutes/README.md#deletedegraphqlroutewithserviceinworkspace) - Delete a a Degraphql_route associated with a Service in a workspace
* [GetDegraphqlRouteWithServiceInWorkspace](docs/sdks/degraphqlroutes/README.md#getdegraphqlroutewithserviceinworkspace) - Get a Degraphql_route associated with a Service in a workspace
* [UpsertDegraphqlRouteWithServiceInWorkspace](docs/sdks/degraphqlroutes/README.md#upsertdegraphqlroutewithserviceinworkspace) - Upsert a Degraphql_route associated with a Service in a workspace
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
* [UpdateDataplaneCertificate](docs/sdks/dpcertificates/README.md#updatedataplanecertificate) - Update DP Client Certificate
* [DeleteDataplaneCertificate](docs/sdks/dpcertificates/README.md#deletedataplanecertificate) - Delete DP Client Certificate

### [DPNodes](docs/sdks/dpnodes/README.md)

* [GetExpectedConfigHash](docs/sdks/dpnodes/README.md#getexpectedconfighash) - Get an Expected Config Hash
* [GetExpectedConfigVersion](docs/sdks/dpnodes/README.md#getexpectedconfigversion) - Get an Expected Config Version
* [ListDataplaneNodes](docs/sdks/dpnodes/README.md#listdataplanenodes) - List Data Plane Node Records
* [GetNodesEol](docs/sdks/dpnodes/README.md#getnodeseol) - List End-of-Life Data Plane Node Records
* [GetNodesNodeID](docs/sdks/dpnodes/README.md#getnodesnodeid) - Get a Data Plane Node Record
* [DeleteNodesNodeID](docs/sdks/dpnodes/README.md#deletenodesnodeid) - Delete Data Plane Node Record

### [GraphQLCostDecorations](docs/sdks/graphqlcostdecorations/README.md)

* [ListGraphqlRateLimitingAdvancedCostInWorkspace](docs/sdks/graphqlcostdecorations/README.md#listgraphqlratelimitingadvancedcostinworkspace) - List all GraphQL Cost Decorations in a workspace
* [GetGraphqlRateLimitingAdvancedCostInWorkspace](docs/sdks/graphqlcostdecorations/README.md#getgraphqlratelimitingadvancedcostinworkspace) - Get a GraphQL Cost Decoration in a workspace
* [ListGraphqlRateLimitingAdvancedCostWithServiceInWorkspace](docs/sdks/graphqlcostdecorations/README.md#listgraphqlratelimitingadvancedcostwithserviceinworkspace) - List all GraphQL Cost Decorations associated with a Service in a workspace
* [CreateGraphqlRateLimitingAdvancedCostWithServiceInWorkspace](docs/sdks/graphqlcostdecorations/README.md#creategraphqlratelimitingadvancedcostwithserviceinworkspace) - Create a new GraphQL Cost Decoration associated with a Service in a workspace
* [DeleteGraphqlRateLimitingAdvancedCostWithServiceInWorkspace](docs/sdks/graphqlcostdecorations/README.md#deletegraphqlratelimitingadvancedcostwithserviceinworkspace) - Delete a a GraphQL Cost Decoration associated with a Service in a workspace
* [GetGraphqlRateLimitingAdvancedCostWithServiceInWorkspace](docs/sdks/graphqlcostdecorations/README.md#getgraphqlratelimitingadvancedcostwithserviceinworkspace) - Get a GraphQL Cost Decoration associated with a Service in a workspace
* [UpsertGraphqlRateLimitingAdvancedCostWithServiceInWorkspace](docs/sdks/graphqlcostdecorations/README.md#upsertgraphqlratelimitingadvancedcostwithserviceinworkspace) - Upsert a GraphQL Cost Decoration associated with a Service in a workspace
* [ListGraphqlRateLimitingAdvancedCost](docs/sdks/graphqlcostdecorations/README.md#listgraphqlratelimitingadvancedcost) - List all GraphQL Cost Decorations
* [GetGraphqlRateLimitingAdvancedCost](docs/sdks/graphqlcostdecorations/README.md#getgraphqlratelimitingadvancedcost) - Get a GraphQL Cost Decoration
* [ListGraphqlRateLimitingAdvancedCostWithService](docs/sdks/graphqlcostdecorations/README.md#listgraphqlratelimitingadvancedcostwithservice) - List all GraphQL Cost Decorations associated with a Service
* [CreateGraphqlRateLimitingAdvancedCostWithService](docs/sdks/graphqlcostdecorations/README.md#creategraphqlratelimitingadvancedcostwithservice) - Create a new GraphQL Cost Decoration associated with a Service
* [DeleteGraphqlRateLimitingAdvancedCostWithService](docs/sdks/graphqlcostdecorations/README.md#deletegraphqlratelimitingadvancedcostwithservice) - Delete a a GraphQL Cost Decoration associated with a Service
* [GetGraphqlRateLimitingAdvancedCostWithService](docs/sdks/graphqlcostdecorations/README.md#getgraphqlratelimitingadvancedcostwithservice) - Get a GraphQL Cost Decoration associated with a Service
* [UpsertGraphqlRateLimitingAdvancedCostWithService](docs/sdks/graphqlcostdecorations/README.md#upsertgraphqlratelimitingadvancedcostwithservice) - Upsert a GraphQL Cost Decoration associated with a Service

### [HMACAuthCredentials](docs/sdks/hmacauthcredentials/README.md)

* [ListHmacAuthWithConsumerInWorkspace](docs/sdks/hmacauthcredentials/README.md#listhmacauthwithconsumerinworkspace) - List all HMAC-auth credentials associated with a Consumer in a workspace
* [CreateHmacAuthWithConsumerInWorkspace](docs/sdks/hmacauthcredentials/README.md#createhmacauthwithconsumerinworkspace) - Create a new HMAC-auth credential associated with a Consumer in a workspace
* [DeleteHmacAuthWithConsumerInWorkspace](docs/sdks/hmacauthcredentials/README.md#deletehmacauthwithconsumerinworkspace) - Delete a a HMAC-auth credential associated with a Consumer in a workspace
* [GetHmacAuthWithConsumerInWorkspace](docs/sdks/hmacauthcredentials/README.md#gethmacauthwithconsumerinworkspace) - Get a HMAC-auth credential associated with a Consumer in a workspace
* [UpsertHmacAuthWithConsumerInWorkspace](docs/sdks/hmacauthcredentials/README.md#upserthmacauthwithconsumerinworkspace) - Upsert a HMAC-auth credential associated with a Consumer in a workspace
* [ListHmacAuthInWorkspace](docs/sdks/hmacauthcredentials/README.md#listhmacauthinworkspace) - List all HMAC-auth credentials in a workspace
* [GetHmacAuthInWorkspace](docs/sdks/hmacauthcredentials/README.md#gethmacauthinworkspace) - Get a HMAC-auth credential in a workspace
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

* [ListJwtWithConsumerInWorkspace](docs/sdks/jwts/README.md#listjwtwithconsumerinworkspace) - List all JWTs associated with a Consumer in a workspace
* [CreateJwtWithConsumerInWorkspace](docs/sdks/jwts/README.md#createjwtwithconsumerinworkspace) - Create a new JWT associated with a Consumer in a workspace
* [DeleteJwtWithConsumerInWorkspace](docs/sdks/jwts/README.md#deletejwtwithconsumerinworkspace) - Delete a a JWT associated with a Consumer in a workspace
* [GetJwtWithConsumerInWorkspace](docs/sdks/jwts/README.md#getjwtwithconsumerinworkspace) - Get a JWT associated with a Consumer in a workspace
* [UpsertJwtWithConsumerInWorkspace](docs/sdks/jwts/README.md#upsertjwtwithconsumerinworkspace) - Upsert a JWT associated with a Consumer in a workspace
* [ListJwtInWorkspace](docs/sdks/jwts/README.md#listjwtinworkspace) - List all JWTs in a workspace
* [GetJwtInWorkspace](docs/sdks/jwts/README.md#getjwtinworkspace) - Get a JWT in a workspace
* [ListJwtWithConsumer](docs/sdks/jwts/README.md#listjwtwithconsumer) - List all JWTs associated with a Consumer
* [CreateJwtWithConsumer](docs/sdks/jwts/README.md#createjwtwithconsumer) - Create a new JWT associated with a Consumer
* [DeleteJwtWithConsumer](docs/sdks/jwts/README.md#deletejwtwithconsumer) - Delete a a JWT associated with a Consumer
* [GetJwtWithConsumer](docs/sdks/jwts/README.md#getjwtwithconsumer) - Get a JWT associated with a Consumer
* [UpsertJwtWithConsumer](docs/sdks/jwts/README.md#upsertjwtwithconsumer) - Upsert a JWT associated with a Consumer
* [ListJwt](docs/sdks/jwts/README.md#listjwt) - List all JWTs
* [GetJwt](docs/sdks/jwts/README.md#getjwt) - Get a JWT

### [Keys](docs/sdks/keys/README.md)

* [ListKeyWithKeySetInWorkspace](docs/sdks/keys/README.md#listkeywithkeysetinworkspace) - List all Keys associated with a KeySet in a workspace
* [CreateKeyWithKeySetInWorkspace](docs/sdks/keys/README.md#createkeywithkeysetinworkspace) - Create a new Key associated with a KeySet in a workspace
* [DeleteKeyWithKeySetInWorkspace](docs/sdks/keys/README.md#deletekeywithkeysetinworkspace) - Delete a a Key associated with a KeySet in a workspace
* [GetKeyWithKeySetInWorkspace](docs/sdks/keys/README.md#getkeywithkeysetinworkspace) - Get a Key associated with a KeySet in a workspace
* [UpsertKeyWithKeySetInWorkspace](docs/sdks/keys/README.md#upsertkeywithkeysetinworkspace) - Upsert a Key associated with a KeySet in a workspace
* [ListKeyInWorkspace](docs/sdks/keys/README.md#listkeyinworkspace) - List all Keys in a workspace
* [CreateKeyInWorkspace](docs/sdks/keys/README.md#createkeyinworkspace) - Create a new Key in a workspace
* [DeleteKeyInWorkspace](docs/sdks/keys/README.md#deletekeyinworkspace) - Delete a Key in a workspace
* [GetKeyInWorkspace](docs/sdks/keys/README.md#getkeyinworkspace) - Get a Key in a workspace
* [UpsertKeyInWorkspace](docs/sdks/keys/README.md#upsertkeyinworkspace) - Upsert a Key in a workspace
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

* [ListKeySetInWorkspace](docs/sdks/keysets/README.md#listkeysetinworkspace) - List all KeySets in a workspace
* [CreateKeySetInWorkspace](docs/sdks/keysets/README.md#createkeysetinworkspace) - Create a new KeySet in a workspace
* [DeleteKeySetInWorkspace](docs/sdks/keysets/README.md#deletekeysetinworkspace) - Delete a KeySet in a workspace
* [GetKeySetInWorkspace](docs/sdks/keysets/README.md#getkeysetinworkspace) - Get a KeySet in a workspace
* [UpsertKeySetInWorkspace](docs/sdks/keysets/README.md#upsertkeysetinworkspace) - Upsert a KeySet in a workspace
* [ListKeySet](docs/sdks/keysets/README.md#listkeyset) - List all KeySets
* [CreateKeySet](docs/sdks/keysets/README.md#createkeyset) - Create a new KeySet
* [DeleteKeySet](docs/sdks/keysets/README.md#deletekeyset) - Delete a KeySet
* [GetKeySet](docs/sdks/keysets/README.md#getkeyset) - Get a KeySet
* [UpsertKeySet](docs/sdks/keysets/README.md#upsertkeyset) - Upsert a KeySet

### [ManagedSystemAccountsRoles](docs/sdks/managedsystemaccountsroles/README.md)

* [GetSystemAccountsAssignedRolesInternal](docs/sdks/managedsystemaccountsroles/README.md#getsystemaccountsassignedrolesinternal) - List Roles (Internal)
* [CreateSystemAccountsAssignedRolesInternal](docs/sdks/managedsystemaccountsroles/README.md#createsystemaccountsassignedrolesinternal) - Assign a role to a managed System Account

### [MCPServers](docs/sdks/mcpservers/README.md)

* [ListMcpServerConfigs](docs/sdks/mcpservers/README.md#listmcpserverconfigs) - List all MCP Servers
* [CreateMcpServerConfig](docs/sdks/mcpservers/README.md#createmcpserverconfig) - Create an MCP Server
* [GetMcpServerConfig](docs/sdks/mcpservers/README.md#getmcpserverconfig) - Get MCP Server
* [UpdateMcpServerConfig](docs/sdks/mcpservers/README.md#updatemcpserverconfig) - Update MCP Server
* [PatchMcpServerConfig](docs/sdks/mcpservers/README.md#patchmcpserverconfig) - Partially Update MCP Server
* [DeleteMcpServerConfig](docs/sdks/mcpservers/README.md#deletemcpserverconfig) - Delete MCP Server
* [GetMcpServerStatus](docs/sdks/mcpservers/README.md#getmcpserverstatus) - Get MCP Server deployment status
* [GetMcpServerGeneratedCode](docs/sdks/mcpservers/README.md#getmcpservergeneratedcode) - Get generated Python code for an MCP Server
* [ListMcpResources](docs/sdks/mcpservers/README.md#listmcpresources) - List MCP Resources
* [CreateMcpResource](docs/sdks/mcpservers/README.md#createmcpresource) - Create an MCP Resource
* [GetMcpResource](docs/sdks/mcpservers/README.md#getmcpresource) - Get an MCP Resource
* [DeleteMcpResource](docs/sdks/mcpservers/README.md#deletemcpresource) - Delete an MCP Resource
* [UpdateMcpResource](docs/sdks/mcpservers/README.md#updatemcpresource) - Update an MCP Resource
* [GetMcpResourceSpec](docs/sdks/mcpservers/README.md#getmcpresourcespec) - Get MCP Resource Spec
* [GetMcpServerSignals](docs/sdks/mcpservers/README.md#getmcpserversignals) - Get MCP Server Signals
* [ListMcpServersByControlPlane](docs/sdks/mcpservers/README.md#listmcpserversbycontrolplane) - List MCP Servers by Control Plane
* [GetMcpServerByControlPlane](docs/sdks/mcpservers/README.md#getmcpserverbycontrolplane) - Get MCP Server by Control Plane
* [GetMcpServerKongEntities](docs/sdks/mcpservers/README.md#getmcpserverkongentities) - Get Kong entities for the MCP Server Gateway
* [PostMcpServerStatus](docs/sdks/mcpservers/README.md#postmcpserverstatus) - Report MCP Server deployment status
* [GetMcpServerCode](docs/sdks/mcpservers/README.md#getmcpservercode) - Get generated Python code for the MCP Server

### [Me](docs/sdks/me/README.md)

* [GetUsersMe](docs/sdks/me/README.md#getusersme) - Get My User Account
* [DeleteUsersMe](docs/sdks/me/README.md#deleteusersme) - Delete My User Account
* [PatchUsersMe](docs/sdks/me/README.md#patchusersme) - Update My User Account
* [GetUsersMePermissions](docs/sdks/me/README.md#getusersmepermissions) - Get My Permissions
* [GetOrganizationsMe](docs/sdks/me/README.md#getorganizationsme) - Get My Organization
* [UpdateOrganizationsMe](docs/sdks/me/README.md#updateorganizationsme) - Update My Organization

### [MeteringEvents](docs/sdks/meteringevents/README.md)

* [ListMeteringEvents](docs/sdks/meteringevents/README.md#listmeteringevents) - List metering events
* [IngestMeteringEvents](docs/sdks/meteringevents/README.md#ingestmeteringevents) - Ingest metering events

### [Meters](docs/sdks/meters/README.md)

* [CreateMeter](docs/sdks/meters/README.md#createmeter) - Create meter
* [ListMeters](docs/sdks/meters/README.md#listmeters) - List meters
* [GetMeter](docs/sdks/meters/README.md#getmeter) - Get meter
* [UpdateMeter](docs/sdks/meters/README.md#updatemeter) - Update meter
* [DeleteMeter](docs/sdks/meters/README.md#deletemeter) - Delete meter
* [QueryMeter](docs/sdks/meters/README.md#querymeter) - Query meter

### [MTLSAuthCredentials](docs/sdks/mtlsauthcredentials/README.md)

* [ListMtlsAuthWithConsumerInWorkspace](docs/sdks/mtlsauthcredentials/README.md#listmtlsauthwithconsumerinworkspace) - List all MTLS-auth credentials associated with a Consumer in a workspace
* [CreateMtlsAuthWithConsumerInWorkspace](docs/sdks/mtlsauthcredentials/README.md#createmtlsauthwithconsumerinworkspace) - Create a new MTLS-auth credential associated with a Consumer in a workspace
* [DeleteMtlsAuthWithConsumerInWorkspace](docs/sdks/mtlsauthcredentials/README.md#deletemtlsauthwithconsumerinworkspace) - Delete a a MTLS-auth credential associated with a Consumer in a workspace
* [GetMtlsAuthWithConsumerInWorkspace](docs/sdks/mtlsauthcredentials/README.md#getmtlsauthwithconsumerinworkspace) - Get a MTLS-auth credential associated with a Consumer in a workspace
* [UpsertMtlsAuthWithConsumerInWorkspace](docs/sdks/mtlsauthcredentials/README.md#upsertmtlsauthwithconsumerinworkspace) - Upsert a MTLS-auth credential associated with a Consumer in a workspace
* [ListMtlsAuthInWorkspace](docs/sdks/mtlsauthcredentials/README.md#listmtlsauthinworkspace) - List all MTLS-auth credentials in a workspace
* [GetMtlsAuthInWorkspace](docs/sdks/mtlsauthcredentials/README.md#getmtlsauthinworkspace) - Get a MTLS-auth credential in a workspace
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

### [OpenMeterApps](docs/sdks/openmeterapps/README.md)

* [ListApps](docs/sdks/openmeterapps/README.md#listapps) - List apps
* [GetApp](docs/sdks/openmeterapps/README.md#getapp) - Get app

### [OpenMeterBilling](docs/sdks/openmeterbilling/README.md)

* [ListCurrencies](docs/sdks/openmeterbilling/README.md#listcurrencies) - List currencies
* [CreateCustomCurrency](docs/sdks/openmeterbilling/README.md#createcustomcurrency) - Create custom currency
* [ListCostBases](docs/sdks/openmeterbilling/README.md#listcostbases) - List cost bases
* [CreateCostBasis](docs/sdks/openmeterbilling/README.md#createcostbasis) - Create cost basis
* [ListCustomerCharges](docs/sdks/openmeterbilling/README.md#listcustomercharges) - List customer charges
* [ListBillingProfiles](docs/sdks/openmeterbilling/README.md#listbillingprofiles) - List billing profiles
* [CreateBillingProfile](docs/sdks/openmeterbilling/README.md#createbillingprofile) - Create a new billing profile
* [GetBillingProfile](docs/sdks/openmeterbilling/README.md#getbillingprofile) - Get a billing profile
* [UpdateBillingProfile](docs/sdks/openmeterbilling/README.md#updatebillingprofile) - Update a billing profile
* [DeleteBillingProfile](docs/sdks/openmeterbilling/README.md#deletebillingprofile) - Delete a billing profile

### [OpenMeterCustomers](docs/sdks/openmetercustomers/README.md)

* [CreateCustomer](docs/sdks/openmetercustomers/README.md#createcustomer) - Create customer
* [ListCustomers](docs/sdks/openmetercustomers/README.md#listcustomers) - List customers
* [GetCustomer](docs/sdks/openmetercustomers/README.md#getcustomer) - Get customer
* [UpsertCustomer](docs/sdks/openmetercustomers/README.md#upsertcustomer) - Upsert customer
* [DeleteCustomer](docs/sdks/openmetercustomers/README.md#deletecustomer) - Delete customer
* [GetCustomerBilling](docs/sdks/openmetercustomers/README.md#getcustomerbilling) - Get customer billing data
* [UpdateCustomerBilling](docs/sdks/openmetercustomers/README.md#updatecustomerbilling) - Update customer billing data
* [UpdateCustomerBillingAppData](docs/sdks/openmetercustomers/README.md#updatecustomerbillingappdata) - Update customer billing app data
* [CreateCustomerStripeCheckoutSession](docs/sdks/openmetercustomers/README.md#createcustomerstripecheckoutsession) - Create Stripe Checkout Session
* [CreateCustomerStripePortalSession](docs/sdks/openmetercustomers/README.md#createcustomerstripeportalsession) - Create Stripe customer portal session
* [CreateCreditAdjustment](docs/sdks/openmetercustomers/README.md#createcreditadjustment) - Create a credit adjustment
* [GetCustomerCreditBalance](docs/sdks/openmetercustomers/README.md#getcustomercreditbalance) - Get a customer's credit balance
* [CreateCreditGrant](docs/sdks/openmetercustomers/README.md#createcreditgrant) - Create a new credit grant
* [ListCreditGrants](docs/sdks/openmetercustomers/README.md#listcreditgrants) - List credit grants
* [GetCreditGrant](docs/sdks/openmetercustomers/README.md#getcreditgrant) - Get a credit grant
* [UpdateCreditGrantExternalSettlement](docs/sdks/openmetercustomers/README.md#updatecreditgrantexternalsettlement) - Update credit grant external settlement status
* [ListCreditTransactions](docs/sdks/openmetercustomers/README.md#listcredittransactions) - List credit transactions

### [OpenMeterDefaults](docs/sdks/openmeterdefaults/README.md)

* [GetOrganizationDefaultTaxCodes](docs/sdks/openmeterdefaults/README.md#getorganizationdefaulttaxcodes) - Get organization default tax codes
* [UpdateOrganizationDefaultTaxCodes](docs/sdks/openmeterdefaults/README.md#updateorganizationdefaulttaxcodes) - Update organization default tax codes

### [OpenMeterEntitlements](docs/sdks/openmeterentitlements/README.md)

* [ListCustomerEntitlementAccess](docs/sdks/openmeterentitlements/README.md#listcustomerentitlementaccess) - List customer entitlement access

### [OpenMeterFeatures](docs/sdks/openmeterfeatures/README.md)

* [ListFeatures](docs/sdks/openmeterfeatures/README.md#listfeatures) - List features
* [CreateFeature](docs/sdks/openmeterfeatures/README.md#createfeature) - Create feature
* [GetFeature](docs/sdks/openmeterfeatures/README.md#getfeature) - Get feature
* [UpdateFeature](docs/sdks/openmeterfeatures/README.md#updatefeature) - Update feature
* [DeleteFeature](docs/sdks/openmeterfeatures/README.md#deletefeature) - Delete feature
* [QueryFeatureCost](docs/sdks/openmeterfeatures/README.md#queryfeaturecost) - Query feature cost

### [OpenMeterGovernance](docs/sdks/openmetergovernance/README.md)

* [QueryGovernanceAccess](docs/sdks/openmetergovernance/README.md#querygovernanceaccess) - Query governance access

### [OpenMeterLLMCost](docs/sdks/openmeterllmcost/README.md)

* [ListLlmCostOverrides](docs/sdks/openmeterllmcost/README.md#listllmcostoverrides) - List LLM cost overrides
* [CreateLlmCostOverride](docs/sdks/openmeterllmcost/README.md#createllmcostoverride) - Create LLM cost override
* [DeleteLlmCostOverride](docs/sdks/openmeterllmcost/README.md#deletellmcostoverride) - Delete LLM cost override
* [ListLlmCostPrices](docs/sdks/openmeterllmcost/README.md#listllmcostprices) - List LLM cost prices
* [GetLlmCostPrice](docs/sdks/openmeterllmcost/README.md#getllmcostprice) - Get LLM cost price

### [OpenMeterProductCatalog](docs/sdks/openmeterproductcatalog/README.md)

* [ListOpenmeterAddons](docs/sdks/openmeterproductcatalog/README.md#listopenmeteraddons) - List add-ons
* [CreateOpenmeterAddon](docs/sdks/openmeterproductcatalog/README.md#createopenmeteraddon) - Create add-on
* [UpdateOpenmeterAddon](docs/sdks/openmeterproductcatalog/README.md#updateopenmeteraddon) - Update add-on
* [GetOpenmeterAddon](docs/sdks/openmeterproductcatalog/README.md#getopenmeteraddon) - Get add-on
* [DeleteOpenmeterAddon](docs/sdks/openmeterproductcatalog/README.md#deleteopenmeteraddon) - Soft delete add-on
* [ArchiveOpenmeterAddon](docs/sdks/openmeterproductcatalog/README.md#archiveopenmeteraddon) - Archive add-on version
* [PublishAddon](docs/sdks/openmeterproductcatalog/README.md#publishaddon) - Publish add-on version
* [ListPlans](docs/sdks/openmeterproductcatalog/README.md#listplans) - List plans
* [CreatePlan](docs/sdks/openmeterproductcatalog/README.md#createplan) - Create plan
* [UpdatePlan](docs/sdks/openmeterproductcatalog/README.md#updateplan) - Update plan
* [GetPlan](docs/sdks/openmeterproductcatalog/README.md#getplan) - Get plan
* [DeletePlan](docs/sdks/openmeterproductcatalog/README.md#deleteplan) - Delete plan
* [ListPlanAddons](docs/sdks/openmeterproductcatalog/README.md#listplanaddons) - List add-ons for plan
* [CreatePlanAddon](docs/sdks/openmeterproductcatalog/README.md#createplanaddon) - Add add-on to plan
* [GetPlanAddon](docs/sdks/openmeterproductcatalog/README.md#getplanaddon) - Get add-on association for plan
* [UpdatePlanAddon](docs/sdks/openmeterproductcatalog/README.md#updateplanaddon) - Update add-on association for plan
* [DeletePlanAddon](docs/sdks/openmeterproductcatalog/README.md#deleteplanaddon) - Remove add-on from plan
* [ArchivePlan](docs/sdks/openmeterproductcatalog/README.md#archiveplan) - Archive plan version
* [PublishPlan](docs/sdks/openmeterproductcatalog/README.md#publishplan) - Publish plan version

### [OpenMeterSubscriptions](docs/sdks/openmetersubscriptions/README.md)

* [CreateSubscription](docs/sdks/openmetersubscriptions/README.md#createsubscription) - Create subscription
* [ListSubscriptions](docs/sdks/openmetersubscriptions/README.md#listsubscriptions) - List subscriptions
* [GetSubscription](docs/sdks/openmetersubscriptions/README.md#getsubscription) - Get subscription
* [ListSubscriptionAddons](docs/sdks/openmetersubscriptions/README.md#listsubscriptionaddons) - List subscription addons
* [GetSubscriptionAddon](docs/sdks/openmetersubscriptions/README.md#getsubscriptionaddon) - Get add-on association for subscription
* [CancelSubscription](docs/sdks/openmetersubscriptions/README.md#cancelsubscription) - Cancel subscription
* [ChangeSubscription](docs/sdks/openmetersubscriptions/README.md#changesubscription) - Change subscription
* [UnscheduleCancelation](docs/sdks/openmetersubscriptions/README.md#unschedulecancelation) - Unschedule subscription cancelation

### [OpenMeterTax](docs/sdks/openmetertax/README.md)

* [CreateTaxCode](docs/sdks/openmetertax/README.md#createtaxcode) - Create tax code
* [ListTaxCodes](docs/sdks/openmetertax/README.md#listtaxcodes) - List tax codes
* [GetTaxCode](docs/sdks/openmetertax/README.md#gettaxcode) - Get tax code
* [UpsertTaxCode](docs/sdks/openmetertax/README.md#upserttaxcode) - Upsert tax code
* [DeleteTaxCode](docs/sdks/openmetertax/README.md#deletetaxcode) - Delete tax code

### [OrganizationFeature](docs/sdks/organizationfeature/README.md)

* [GetOrganizationFeature](docs/sdks/organizationfeature/README.md#getorganizationfeature) - Get Feature Configuration
* [UpsertOrganizationFeature](docs/sdks/organizationfeature/README.md#upsertorganizationfeature) - Upsert Feature Configuration

### [Pages](docs/sdks/pages/README.md)

* [CreateDefaultContent](docs/sdks/pages/README.md#createdefaultcontent) - Create Boilerplate Content

### [PartialLinks](docs/sdks/partiallinks/README.md)

* [ListPartialLinkInWorkspace](docs/sdks/partiallinks/README.md#listpartiallinkinworkspace) - List partial links in a workspace
* [ListPartialLink](docs/sdks/partiallinks/README.md#listpartiallink) - List partial links

### [Partials](docs/sdks/partials/README.md)

* [ListPartialInWorkspace](docs/sdks/partials/README.md#listpartialinworkspace) - List all Partials in a workspace
* [CreatePartialInWorkspace](docs/sdks/partials/README.md#createpartialinworkspace) - Create a new Partial in a workspace
* [DeletePartialInWorkspace](docs/sdks/partials/README.md#deletepartialinworkspace) - Delete a Partial in a workspace
* [GetPartialInWorkspace](docs/sdks/partials/README.md#getpartialinworkspace) - Get a Partial in a workspace
* [UpsertPartialInWorkspace](docs/sdks/partials/README.md#upsertpartialinworkspace) - Upsert a Partial in a workspace
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

* [SearchPlugin](docs/sdks/plugins/README.md#searchplugin) - Search for Plugins
* [ListPluginWithConsumerGroupInWorkspace](docs/sdks/plugins/README.md#listpluginwithconsumergroupinworkspace) - List all Plugins associated with a Consumer Group in a workspace
* [CreatePluginWithConsumerGroupInWorkspace](docs/sdks/plugins/README.md#createpluginwithconsumergroupinworkspace) - Create a new Plugin associated with a Consumer Group in a workspace
* [DeletePluginWithConsumerGroupInWorkspace](docs/sdks/plugins/README.md#deletepluginwithconsumergroupinworkspace) - Delete a a Plugin associated with a Consumer Group in a workspace
* [GetPluginWithConsumerGroupInWorkspace](docs/sdks/plugins/README.md#getpluginwithconsumergroupinworkspace) - Get a Plugin associated with a Consumer Group in a workspace
* [UpsertPluginWithConsumerGroupInWorkspace](docs/sdks/plugins/README.md#upsertpluginwithconsumergroupinworkspace) - Upsert a Plugin associated with a Consumer Group in a workspace
* [ListPluginWithConsumerInWorkspace](docs/sdks/plugins/README.md#listpluginwithconsumerinworkspace) - List all Plugins associated with a Consumer in a workspace
* [CreatePluginWithConsumerInWorkspace](docs/sdks/plugins/README.md#createpluginwithconsumerinworkspace) - Create a new Plugin associated with a Consumer in a workspace
* [DeletePluginWithConsumerInWorkspace](docs/sdks/plugins/README.md#deletepluginwithconsumerinworkspace) - Delete a a Plugin associated with a Consumer in a workspace
* [GetPluginWithConsumerInWorkspace](docs/sdks/plugins/README.md#getpluginwithconsumerinworkspace) - Get a Plugin associated with a Consumer in a workspace
* [UpsertPluginWithConsumerInWorkspace](docs/sdks/plugins/README.md#upsertpluginwithconsumerinworkspace) - Upsert a Plugin associated with a Consumer in a workspace
* [ListPluginInWorkspace](docs/sdks/plugins/README.md#listplugininworkspace) - List all Plugins in a workspace
* [CreatePluginInWorkspace](docs/sdks/plugins/README.md#createplugininworkspace) - Create a new Plugin in a workspace
* [DeletePluginInWorkspace](docs/sdks/plugins/README.md#deleteplugininworkspace) - Delete a Plugin in a workspace
* [GetPluginInWorkspace](docs/sdks/plugins/README.md#getplugininworkspace) - Get a Plugin in a workspace
* [UpsertPluginInWorkspace](docs/sdks/plugins/README.md#upsertplugininworkspace) - Upsert a Plugin in a workspace
* [ListPluginWithRouteInWorkspace](docs/sdks/plugins/README.md#listpluginwithrouteinworkspace) - List all Plugins associated with a Route in a workspace
* [CreatePluginWithRouteInWorkspace](docs/sdks/plugins/README.md#createpluginwithrouteinworkspace) - Create a new Plugin associated with a Route in a workspace
* [DeletePluginWithRouteInWorkspace](docs/sdks/plugins/README.md#deletepluginwithrouteinworkspace) - Delete a a Plugin associated with a Route in a workspace
* [GetPluginWithRouteInWorkspace](docs/sdks/plugins/README.md#getpluginwithrouteinworkspace) - Get a Plugin associated with a Route in a workspace
* [UpsertPluginWithRouteInWorkspace](docs/sdks/plugins/README.md#upsertpluginwithrouteinworkspace) - Upsert a Plugin associated with a Route in a workspace
* [ListPluginWithServiceInWorkspace](docs/sdks/plugins/README.md#listpluginwithserviceinworkspace) - List all Plugins associated with a Service in a workspace
* [CreatePluginWithServiceInWorkspace](docs/sdks/plugins/README.md#createpluginwithserviceinworkspace) - Create a new Plugin associated with a Service in a workspace
* [DeletePluginWithServiceInWorkspace](docs/sdks/plugins/README.md#deletepluginwithserviceinworkspace) - Delete a a Plugin associated with a Service in a workspace
* [GetPluginWithServiceInWorkspace](docs/sdks/plugins/README.md#getpluginwithserviceinworkspace) - Get a Plugin associated with a Service in a workspace
* [UpsertPluginWithServiceInWorkspace](docs/sdks/plugins/README.md#upsertpluginwithserviceinworkspace) - Upsert a Plugin associated with a Service in a workspace
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
* [ListPortalIdpTeamGroupMappings](docs/sdks/portalauthsettings/README.md#listportalidpteamgroupmappings) - List Team Group Mappings
* [CreatePortalIdpTeamGroupMapping](docs/sdks/portalauthsettings/README.md#createportalidpteamgroupmapping) - Create Team Group Mapping
* [GetPortalIdpTeamGroupMapping](docs/sdks/portalauthsettings/README.md#getportalidpteamgroupmapping) - Get Team Group Mapping
* [DeletePortalIdpTeamGroupMapping](docs/sdks/portalauthsettings/README.md#deleteportalidpteamgroupmapping) - Delete Team Group Mapping

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

* [CreateDeveloper](docs/sdks/portaldevelopers/README.md#createdeveloper) - Create Developer Account
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

### [PortalForms](docs/sdks/portalforms/README.md)

* [CreatePortalForm](docs/sdks/portalforms/README.md#createportalform) - Create Form
* [ListPortalForms](docs/sdks/portalforms/README.md#listportalforms) - List Forms
* [GetPortalForm](docs/sdks/portalforms/README.md#getportalform) - Get Form
* [ReplacePortalForm](docs/sdks/portalforms/README.md#replaceportalform) - Replace Form
* [DeletePortalForm](docs/sdks/portalforms/README.md#deleteportalform) - Delete Form

### [PortalIntegrations](docs/sdks/portalintegrations/README.md)

* [GetPortalIntegrations](docs/sdks/portalintegrations/README.md#getportalintegrations) - Get Portal Integration Configurations
* [UpsertPortalIntegrations](docs/sdks/portalintegrations/README.md#upsertportalintegrations) - Replace Integration Configurations
* [UpdatePortalIntegrations](docs/sdks/portalintegrations/README.md#updateportalintegrations) - Update Integration Configurations

### [PortalMCPRegistryPublications](docs/sdks/portalmcpregistrypublications/README.md)

* [ListPortalMcpRegistryPublications](docs/sdks/portalmcpregistrypublications/README.md#listportalmcpregistrypublications) - List portal MCP registry publications
* [ListPortalMcpRegistries](docs/sdks/portalmcpregistrypublications/README.md#listportalmcpregistries) - List MCP Registries
* [ListPortalComputedMcpRegistries](docs/sdks/portalmcpregistrypublications/README.md#listportalcomputedmcpregistries) - List computed MCP Registries

### [PortalPages](docs/sdks/portalpages/README.md)

* [ListPortalPages](docs/sdks/portalpages/README.md#listportalpages) - List Pages
* [CreatePortalPage](docs/sdks/portalpages/README.md#createportalpage) - Create Page
* [GetPortalPage](docs/sdks/portalpages/README.md#getportalpage) - Get a Page
* [UpdatePortalPage](docs/sdks/portalpages/README.md#updateportalpage) - Update Page
* [DeletePortalPage](docs/sdks/portalpages/README.md#deleteportalpage) - Delete Page
* [MovePortalPages](docs/sdks/portalpages/README.md#moveportalpages) - Move Page

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

### [PortalsIPAllowList](docs/sdks/portalsipallowlist/README.md)

* [CreatePortalIPAllowList](docs/sdks/portalsipallowlist/README.md#createportalipallowlist) - Create an IP allow list for a portal
* [ListPortalIPAllowList](docs/sdks/portalsipallowlist/README.md#listportalipallowlist) - List the IP allow list for portal
* [GetPortalIPAllowList](docs/sdks/portalsipallowlist/README.md#getportalipallowlist) - Get an IP allow list for a portal
* [PutPortalIPAllowList](docs/sdks/portalsipallowlist/README.md#putportalipallowlist) - Replace an IP allow list for a portal
* [UpdatePortalIPAllowList](docs/sdks/portalsipallowlist/README.md#updateportalipallowlist) - Update an IP allow list for a portal
* [DeletePortalIPAllowList](docs/sdks/portalsipallowlist/README.md#deleteportalipallowlist) - Delete an IP allow list from a portal

### [Quotas](docs/sdks/quotas/README.md)

* [ListDefaultQuotas](docs/sdks/quotas/README.md#listdefaultquotas) - List Default Quotas
* [ListQuotas](docs/sdks/quotas/README.md#listquotas) - List Quotas
* [CreateQuota](docs/sdks/quotas/README.md#createquota) - Create Quota
* [FetchQuota](docs/sdks/quotas/README.md#fetchquota) - Get a Quota
* [UpdateQuota](docs/sdks/quotas/README.md#updatequota) - Update Quota

### [ResourceActions](docs/sdks/resourceactions/README.md)

* [ListResourceActions](docs/sdks/resourceactions/README.md#listresourceactions) - List Resource Actions

### [ResourceIngestion](docs/sdks/resourceingestion/README.md)

* [ScheduleResourceIngestion](docs/sdks/resourceingestion/README.md#scheduleresourceingestion) - Schedule Resource Ingestion
* [FetchResourceIngestion](docs/sdks/resourceingestion/README.md#fetchresourceingestion) - Get a Resource Ingestion

### [Roles](docs/sdks/roles/README.md)

* [GetPredefinedRoles](docs/sdks/roles/README.md#getpredefinedroles) - Get Predefined Roles
* [ListTeamRoles](docs/sdks/roles/README.md#listteamroles) - List Team Roles
* [TeamsAssignRole](docs/sdks/roles/README.md#teamsassignrole) - Assign Team Role
* [GetTeamRole](docs/sdks/roles/README.md#getteamrole) - Get Team Role
* [TeamsRemoveRole](docs/sdks/roles/README.md#teamsremoverole) - Remove Team Role
* [ListUserRoles](docs/sdks/roles/README.md#listuserroles) - List User Roles
* [UsersAssignRole](docs/sdks/roles/README.md#usersassignrole) - Assign Role
* [GetUserRole](docs/sdks/roles/README.md#getuserrole) - Get User Role
* [UsersRemoveRole](docs/sdks/roles/README.md#usersremoverole) - Remove Role

### [Routes](docs/sdks/routes/README.md)

* [ListRouteInWorkspace](docs/sdks/routes/README.md#listrouteinworkspace) - List all Routes in a workspace
* [CreateRouteInWorkspace](docs/sdks/routes/README.md#createrouteinworkspace) - Create a new Route in a workspace
* [DeleteRouteInWorkspace](docs/sdks/routes/README.md#deleterouteinworkspace) - Delete a Route in a workspace
* [GetRouteInWorkspace](docs/sdks/routes/README.md#getrouteinworkspace) - Get a Route in a workspace
* [UpsertRouteInWorkspace](docs/sdks/routes/README.md#upsertrouteinworkspace) - Upsert a Route in a workspace
* [ListRouteWithServiceInWorkspace](docs/sdks/routes/README.md#listroutewithserviceinworkspace) - List all Routes associated with a Service in a workspace
* [CreateRouteWithServiceInWorkspace](docs/sdks/routes/README.md#createroutewithserviceinworkspace) - Create a new Route associated with a Service in a workspace
* [DeleteRouteWithServiceInWorkspace](docs/sdks/routes/README.md#deleteroutewithserviceinworkspace) - Delete a a Route associated with a Service in a workspace
* [GetRouteWithServiceInWorkspace](docs/sdks/routes/README.md#getroutewithserviceinworkspace) - Get a Route associated with a Service in a workspace
* [UpsertRouteWithServiceInWorkspace](docs/sdks/routes/README.md#upsertroutewithserviceinworkspace) - Upsert a Route associated with a Service in a workspace
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

* [ListServiceInWorkspace](docs/sdks/services/README.md#listserviceinworkspace) - List all Services in a workspace
* [CreateServiceInWorkspace](docs/sdks/services/README.md#createserviceinworkspace) - Create a new Service in a workspace
* [DeleteServiceInWorkspace](docs/sdks/services/README.md#deleteserviceinworkspace) - Delete a Service in a workspace
* [GetServiceInWorkspace](docs/sdks/services/README.md#getserviceinworkspace) - Get a Service in a workspace
* [UpsertServiceInWorkspace](docs/sdks/services/README.md#upsertserviceinworkspace) - Upsert a Service in a workspace
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

* [ListSniWithCertificateInWorkspace](docs/sdks/snis/README.md#listsniwithcertificateinworkspace) - List all SNIs associated with a Certificate in a workspace
* [CreateSniWithCertificateInWorkspace](docs/sdks/snis/README.md#createsniwithcertificateinworkspace) - Create a new SNI associated with a Certificate in a workspace
* [DeleteSniWithCertificateInWorkspace](docs/sdks/snis/README.md#deletesniwithcertificateinworkspace) - Delete a an SNI associated with a Certificate in a workspace
* [GetSniWithCertificateInWorkspace](docs/sdks/snis/README.md#getsniwithcertificateinworkspace) - Get an SNI associated with a Certificate in a workspace
* [UpsertSniWithCertificateInWorkspace](docs/sdks/snis/README.md#upsertsniwithcertificateinworkspace) - Upsert an SNI associated with a Certificate in a workspace
* [ListSniInWorkspace](docs/sdks/snis/README.md#listsniinworkspace) - List all SNIs in a workspace
* [CreateSniInWorkspace](docs/sdks/snis/README.md#createsniinworkspace) - Create a new SNI in a workspace
* [DeleteSniInWorkspace](docs/sdks/snis/README.md#deletesniinworkspace) - Delete an SNI in a workspace
* [GetSniInWorkspace](docs/sdks/snis/README.md#getsniinworkspace) - Get an SNI in a workspace
* [UpsertSniInWorkspace](docs/sdks/snis/README.md#upsertsniinworkspace) - Upsert a SNI in a workspace
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
* [GetSystemAccountRole](docs/sdks/systemaccountsroles/README.md#getsystemaccountrole) - Get System Account Role
* [DeleteSystemAccountsAccountIDAssignedRolesRoleID](docs/sdks/systemaccountsroles/README.md#deletesystemaccountsaccountidassignedrolesroleid) - Delete Assigned Role from System Account

### [SystemAccountsTeamMembership](docs/sdks/systemaccountsteammembership/README.md)

* [GetTeamsTeamIDSystemAccounts](docs/sdks/systemaccountsteammembership/README.md#getteamsteamidsystemaccounts) - List System Accounts on a Team
* [PostTeamsTeamIDSystemAccounts](docs/sdks/systemaccountsteammembership/README.md#postteamsteamidsystemaccounts) - Add System Account to a Team
* [DeleteTeamsTeamIDSystemAccountsAccountID](docs/sdks/systemaccountsteammembership/README.md#deleteteamsteamidsystemaccountsaccountid) - Remove System Account From Team
* [GetSystemAccountsAccountIDTeams](docs/sdks/systemaccountsteammembership/README.md#getsystemaccountsaccountidteams) - List Teams for a System Account

### [Targets](docs/sdks/targets/README.md)

* [ListTargets](docs/sdks/targets/README.md#listtargets) - List all targets for a control plane
* [ListTargetsInWorkspace](docs/sdks/targets/README.md#listtargetsinworkspace) - List all targets for a control plane in a workspace
* [ListTargetWithUpstreamInWorkspace](docs/sdks/targets/README.md#listtargetwithupstreaminworkspace) - List all Targets associated with an Upstream in a workspace
* [CreateTargetWithUpstreamInWorkspace](docs/sdks/targets/README.md#createtargetwithupstreaminworkspace) - Create a new Target associated with an Upstream in a workspace
* [DeleteTargetWithUpstreamInWorkspace](docs/sdks/targets/README.md#deletetargetwithupstreaminworkspace) - Delete a a Target associated with an Upstream in a workspace
* [GetTargetWithUpstreamInWorkspace](docs/sdks/targets/README.md#gettargetwithupstreaminworkspace) - Get a Target associated with an Upstream in a workspace
* [UpsertTargetWithUpstreamInWorkspace](docs/sdks/targets/README.md#upserttargetwithupstreaminworkspace) - Upsert a Target associated with an Upstream in a workspace
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

* [ListUpstreamInWorkspace](docs/sdks/upstreams/README.md#listupstreaminworkspace) - List all Upstreams in a workspace
* [CreateUpstreamInWorkspace](docs/sdks/upstreams/README.md#createupstreaminworkspace) - Create a new Upstream in a workspace
* [DeleteUpstreamInWorkspace](docs/sdks/upstreams/README.md#deleteupstreaminworkspace) - Delete an Upstream in a workspace
* [GetUpstreamInWorkspace](docs/sdks/upstreams/README.md#getupstreaminworkspace) - Get an Upstream in a workspace
* [UpsertUpstreamInWorkspace](docs/sdks/upstreams/README.md#upsertupstreaminworkspace) - Upsert a Upstream in a workspace
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

* [ListVaultInWorkspace](docs/sdks/vaults/README.md#listvaultinworkspace) - List all Vaults in a workspace
* [CreateVaultInWorkspace](docs/sdks/vaults/README.md#createvaultinworkspace) - Create a new Vault in a workspace
* [DeleteVaultInWorkspace](docs/sdks/vaults/README.md#deletevaultinworkspace) - Delete a Vault in a workspace
* [GetVaultInWorkspace](docs/sdks/vaults/README.md#getvaultinworkspace) - Get a Vault in a workspace
* [UpsertVaultInWorkspace](docs/sdks/vaults/README.md#upsertvaultinworkspace) - Upsert a Vault in a workspace
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

### [Workspaces](docs/sdks/workspaces/README.md)

* [ListWorkspaces](docs/sdks/workspaces/README.md#listworkspaces) - List all Workspaces
* [CreateWorkspace](docs/sdks/workspaces/README.md#createworkspace) - Create a Workspace
* [GetWorkspace](docs/sdks/workspaces/README.md#getworkspace) - Get a Workspace
* [DeleteWorkspace](docs/sdks/workspaces/README.md#deleteworkspace) - Delete a Workspace
* [UpsertWorkspace](docs/sdks/workspaces/README.md#upsertworkspace) - Upsert a Workspace

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
