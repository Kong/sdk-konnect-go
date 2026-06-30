# ResourceBoundIntegrationAPISpecProviderPayload

API spec provider registered by a catalog integration.

These providers differ from `IntegrationApiSpecProvider` in that they
denote a binding relationship between a resource type and the API spec.
This means that the API Spec will automatically be created/deleted for/from a service
as resources of the given type are mapped/unmapped.

Consult integration documentation to learn more about available API spec providers.



## Fields

| Field                                                                                                                                                                                     | Type                                                                                                                                                                                      | Required                                                                                                                                                                                  | Description                                                                                                                                                                               |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Type`                                                                                                                                                                                    | `string`                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                        | The globally unique API spec provider type that is registered by a given catalog integration.<br/>Resource-bound providers create a 1-to-1 mapping between a resource type and the API Spec.<br/> |
| `Config`                                                                                                                                                                                  | [components.ResourceBoundIntegrationAPISpecProviderPayloadConfig](../../models/components/resourceboundintegrationapispecproviderpayloadconfig.md)                                        | :heavy_check_mark:                                                                                                                                                                        | N/A                                                                                                                                                                                       |