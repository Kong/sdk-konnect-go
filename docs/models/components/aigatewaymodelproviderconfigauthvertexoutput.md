# AIGatewayModelProviderConfigAuthVertexOutput

**Pre-release Feature**
This feature is currently in beta and is subject to change.

Configuration for Vertex model provider.


## Fields

| Field                                                                                                                          | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `Type`                                                                                                                         | [components.AIGatewayModelProviderConfigAuthVertexType](../../models/components/aigatewaymodelproviderconfigauthvertextype.md) | :heavy_check_mark:                                                                                                             | N/A                                                                                                                            |
| `UseGcpServiceAccount`                                                                                                         | `*bool`                                                                                                                        | :heavy_minus_sign:                                                                                                             | Use the Google Cloud Service Account (or user-assigned identity) to authenticate with Vertex-provider models.                  |