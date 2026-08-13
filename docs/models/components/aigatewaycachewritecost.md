# AIGatewayCacheWriteCost

**Pre-release Feature**
This feature is currently in beta and is subject to change.


## Fields

| Field                                               | Type                                                | Required                                            | Description                                         |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| `TTL`                                               | `string`                                            | :heavy_check_mark:                                  | Cache TTL this price applies to, e.g. "5m" or "1h". |
| `Cost`                                              | `float64`                                           | :heavy_check_mark:                                  | Cost per 1M cache-write prompt tokens for this TTL. |