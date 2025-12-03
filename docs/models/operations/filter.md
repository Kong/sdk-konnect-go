# Filter

Filter documents returned in the response.


## Fields

| Field                                                                                         | Type                                                                                          | Required                                                                                      | Description                                                                                   |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `BackendClusterID`                                                                            | [*components.StringFieldEqualsFilter](../../models/components/stringfieldequalsfilter.md)     | :heavy_minus_sign:                                                                            | Filters on the given string field value by exact match.                                       |
| `Name`                                                                                        | [*components.StringFieldContainsFilter](../../models/components/stringfieldcontainsfilter.md) | :heavy_minus_sign:                                                                            | Filters on the given string field value by fuzzy match.                                       |