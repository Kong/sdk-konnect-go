# ListRegistrationsQueryParamFilter

Filter application registrations returned in the response.


## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `DeveloperID`                                                                                | [*components.UUIDFieldFilter](../../models/components/uuidfieldfilter.md)                    | :heavy_minus_sign:                                                                           | Filter using **one** of the following operators: `eq`, `oeq`, `neq`                          |
| `ApplicationName`                                                                            | [*components.StringFieldFilter](../../models/components/stringfieldfilter.md)                | :heavy_minus_sign:                                                                           | Filter using **one** of the following operators: `eq`, `oeq`, `neq`, `contains`, `ocontains` |
| `Status`                                                                                     | [*components.StringFieldFilter](../../models/components/stringfieldfilter.md)                | :heavy_minus_sign:                                                                           | Filter using **one** of the following operators: `eq`, `oeq`, `neq`, `contains`, `ocontains` |