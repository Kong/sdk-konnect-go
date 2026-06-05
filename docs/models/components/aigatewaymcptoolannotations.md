# AIGatewayMCPToolAnnotations


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `DestructiveHint`                                                | `*bool`                                                          | :heavy_minus_sign:                                               | If true, the tool may perform destructive updates                |
| `IdempotentHint`                                                 | `*bool`                                                          | :heavy_minus_sign:                                               | If true, repeated calls with same args have no additional effect |
| `OpenWorldHint`                                                  | `*bool`                                                          | :heavy_minus_sign:                                               | If true, tool interacts with external entities                   |
| `ReadOnlyHint`                                                   | `*bool`                                                          | :heavy_minus_sign:                                               | If true, the tool does not modify its environment                |
| `Title`                                                          | `*string`                                                        | :heavy_minus_sign:                                               | Human-readable title for the tool                                |