# ListRegistrationsQueryParamFilter

Filter application registrations returned in the response.


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `DeveloperID`                                                                 | [*components.UUIDFieldFilter](../../models/components/uuidfieldfilter.md)     | :heavy_minus_sign:                                                            | Filters on the given UUID field value by exact match.                         |
| `ApplicationName`                                                             | [*components.StringFieldFilter](../../models/components/stringfieldfilter.md) | :heavy_minus_sign:                                                            | Filters on the given string field value by either exact or fuzzy match.       |
| `Status`                                                                      | [*components.StringFieldFilter](../../models/components/stringfieldfilter.md) | :heavy_minus_sign:                                                            | Filters on the given string field value by either exact or fuzzy match.       |