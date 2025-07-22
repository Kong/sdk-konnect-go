# Filter

Filter application auth strategies returned in the response.


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `StrategyType`                                                                | [*components.StringFieldFilter](../../models/components/stringfieldfilter.md) | :heavy_minus_sign:                                                            | Filters on the given string field value by either exact or fuzzy match.       |
| `Name`                                                                        | [*components.StringFieldFilter](../../models/components/stringfieldfilter.md) | :heavy_minus_sign:                                                            | Filters on the given string field value by either exact or fuzzy match.       |
| `DisplayName`                                                                 | [*components.StringFieldFilter](../../models/components/stringfieldfilter.md) | :heavy_minus_sign:                                                            | Filters on the given string field value by either exact or fuzzy match.       |
| `DcrProviderID`                                                               | [*components.UUIDFieldFilter](../../models/components/uuidfieldfilter.md)     | :heavy_minus_sign:                                                            | Filters on the given UUID field value by exact match.                         |
| `DcrProviderName`                                                             | [*components.StringFieldFilter](../../models/components/stringfieldfilter.md) | :heavy_minus_sign:                                                            | Filters on the given string field value by either exact or fuzzy match.       |
| `DcrProviderType`                                                             | [*components.StringFieldFilter](../../models/components/stringfieldfilter.md) | :heavy_minus_sign:                                                            | Filters on the given string field value by either exact or fuzzy match.       |