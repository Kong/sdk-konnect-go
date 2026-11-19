# AICostNotificationThreshold

A share of a budget's limit at which its notification rule fires. Spend reaches the threshold when it equals or exceeds that share of the limit.


## Fields

| Field                                                                                                      | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `Percent`                                                                                                  | `int64`                                                                                                    | :heavy_check_mark:                                                                                         | The share of the limit, as a whole percent. A value above 100 fires on spend beyond the limit.<br/><br/>1 to 1000. |