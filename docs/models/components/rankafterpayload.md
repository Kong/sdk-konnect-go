# RankAfterPayload

Determines how this suggestion rule ranks against existing suggestion rules for the integration.
When rank is not specified, this suggestion rule will placed as the lowest priority rule.



## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `AfterSuggestionRuleID`                                                      | `*string`                                                                    | :heavy_check_mark:                                                           | When set to `null`, this rule is indicated to be the highest priority rule.<br/> |