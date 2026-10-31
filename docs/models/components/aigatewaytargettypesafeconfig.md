# AIGatewayTargetTypesafeConfig

Typesafe-specific configuration for a model.

**Requires a minimum runtime version of `2.2`**.


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `Type`                                                   | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `InputCost`                                              | `*float64`                                               | :heavy_minus_sign:                                       | Cost per 1M input tokens for billing and cost tracking.  |
| `OutputCost`                                             | `*float64`                                               | :heavy_minus_sign:                                       | Cost per 1M output tokens for billing and cost tracking. |
| `UpstreamURL`                                            | `*string`                                                | :heavy_minus_sign:                                       | The upstream URL for the model endpoint.                 |