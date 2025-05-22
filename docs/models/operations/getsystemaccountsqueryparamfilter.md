# GetSystemAccountsQueryParamFilter

Filter system accounts returned in the response.


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   | Example                                                                       |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `Name`                                                                        | [*components.StringFieldFilter](../../models/components/stringfieldfilter.md) | :heavy_minus_sign:                                                            | Filters on the given string field value by exact match inequality.            |                                                                               |
| `Description`                                                                 | [*components.StringFieldFilter](../../models/components/stringfieldfilter.md) | :heavy_minus_sign:                                                            | Filters on the given string field value by exact match inequality.            |                                                                               |
| `KonnectManaged`                                                              | **bool*                                                                       | :heavy_minus_sign:                                                            | Filter by a boolean value (true/false).                                       | true                                                                          |