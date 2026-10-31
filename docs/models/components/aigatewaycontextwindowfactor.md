# AIGatewayContextWindowFactor


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `Above`                                                                   | `string`                                                                  | :heavy_check_mark:                                                        | Input-token threshold above which the factors apply, e.g. "128k" or "1m". |
| `InputFactor`                                                             | `float64`                                                                 | :heavy_check_mark:                                                        | Multiplier applied to input pricing above the threshold.                  |
| `OutputFactor`                                                            | `float64`                                                                 | :heavy_check_mark:                                                        | Multiplier applied to output pricing above the threshold.                 |