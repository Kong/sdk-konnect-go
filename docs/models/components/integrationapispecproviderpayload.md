# IntegrationAPISpecProviderPayload

API spec provider registered by a catalog integration.

Integrations can function as API spec providers where they register a globally unique
`type` and config schema which defines the shape of `config`.

Consult integration documentation to learn more about available API spec providers.



## Fields

| Field                                                                                         | Type                                                                                          | Required                                                                                      | Description                                                                                   | Example                                                                                       |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `Type`                                                                                        | `string`                                                                                      | :heavy_check_mark:                                                                            | The globally unique API spec provider type that is registered by a given catalog integration. | swaggerhub_api_version                                                                        |
| `Config`                                                                                      | map[string]`any`                                                                              | :heavy_check_mark:                                                                            | JSON object containing values as defined by integration provider's config schema.             | {<br/>"owner": "petco",<br/>"api": "pet_store",<br/>"version": "v3"<br/>}                     |
| `IntegrationInstance`                                                                         | `string`                                                                                      | :heavy_check_mark:                                                                            | The integration instance id or name                                                           |                                                                                               |