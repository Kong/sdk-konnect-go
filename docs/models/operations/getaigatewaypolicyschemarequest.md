# GetAiGatewayPolicySchemaRequest


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           | Example                                                               |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `GatewayID`                                                           | `string`                                                              | :heavy_check_mark:                                                    | The unique ID of the AI Gateway.                                      | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                                  |
| `PolicySchemaName`                                                    | `string`                                                              | :heavy_check_mark:                                                    | The type of the Policy. This is equivalent to the Kong 3 plugin name. | ai-sanitizer                                                          |