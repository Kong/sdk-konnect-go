# ListApplicationsQueryParamFilter

Filter applications returned in the response.


## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `DeveloperID`                                                                                | [*components.UUIDFieldFilter](../../models/components/uuidfieldfilter.md)                    | :heavy_minus_sign:                                                                           | Filter using **one** of the following operators: `eq`, `oeq`, `neq`                          |
| `Name`                                                                                       | [*components.StringFieldFilter](../../models/components/stringfieldfilter.md)                | :heavy_minus_sign:                                                                           | Filter using **one** of the following operators: `eq`, `oeq`, `neq`, `contains`, `ocontains` |