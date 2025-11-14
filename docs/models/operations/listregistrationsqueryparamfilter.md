# ListRegistrationsQueryParamFilter

Filter application registrations returned in the response.


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `DeveloperID`                                                                 | [*components.UUIDFieldFilter](../../models/components/uuidfieldfilter.md)     | :heavy_minus_sign:                                                            | Filters on the given string field value by exact match inequality.            |
| `ApplicationName`                                                             | [*components.StringFieldFilter](../../models/components/stringfieldfilter.md) | :heavy_minus_sign:                                                            | Filters on the given string field value by exact match inequality.            |
| `Status`                                                                      | [*components.StringFieldFilter](../../models/components/stringfieldfilter.md) | :heavy_minus_sign:                                                            | Filters on the given string field value by exact match inequality.            |