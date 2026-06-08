# LLMCostOverrideCreatePricing

Token pricing data.


## Fields

| Field                                   | Type                                    | Required                                | Description                             |
| --------------------------------------- | --------------------------------------- | --------------------------------------- | --------------------------------------- |
| `InputPerToken`                         | `string`                                | :heavy_check_mark:                      | Input price per token (USD).            |
| `OutputPerToken`                        | `string`                                | :heavy_check_mark:                      | Output price per token (USD).           |
| `CacheReadPerToken`                     | `*string`                               | :heavy_minus_sign:                      | Cache read price per token (USD).       |
| `CacheWritePerToken`                    | `*string`                               | :heavy_minus_sign:                      | Cache write price per token (USD).      |
| `ReasoningPerToken`                     | `*string`                               | :heavy_minus_sign:                      | Reasoning output price per token (USD). |