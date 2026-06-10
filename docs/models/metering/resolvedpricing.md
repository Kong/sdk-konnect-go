# ResolvedPricing

Resolved per-token pricing from the LLM cost database. Populated in responses
when the provider and model can be determined, either from static values or from
meter group-by filters with exact matches.


## Fields

| Field                              | Type                               | Required                           | Description                        |
| ---------------------------------- | ---------------------------------- | ---------------------------------- | ---------------------------------- |
| `InputPerToken`                    | `string`                           | :heavy_check_mark:                 | Cost per input token in USD.       |
| `OutputPerToken`                   | `string`                           | :heavy_check_mark:                 | Cost per output token in USD.      |
| `CacheReadPerToken`                | `*string`                          | :heavy_minus_sign:                 | Cost per cache read token in USD.  |
| `ReasoningPerToken`                | `*string`                          | :heavy_minus_sign:                 | Cost per reasoning token in USD.   |
| `CacheWritePerToken`               | `*string`                          | :heavy_minus_sign:                 | Cost per cache write token in USD. |