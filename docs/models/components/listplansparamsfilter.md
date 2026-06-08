# ListPlansParamsFilter

Filter options for listing plans.


## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `Key`                                                                                        | [*components.StringFieldFilter](../../models/components/stringfieldfilter.md)                | :heavy_minus_sign:                                                                           | Filter using **one** of the following operators: `eq`, `oeq`, `neq`, `contains`, `ocontains` |
| `Name`                                                                                       | [*components.StringFieldFilter](../../models/components/stringfieldfilter.md)                | :heavy_minus_sign:                                                                           | Filter using **one** of the following operators: `eq`, `oeq`, `neq`, `contains`, `ocontains` |
| `Status`                                                                                     | [*components.StringFieldFilterExact](../../models/components/stringfieldfilterexact.md)      | :heavy_minus_sign:                                                                           | Filter using **one** of the following operators: `eq`, `oeq`, `neq`                          |
| `Currency`                                                                                   | [*components.StringFieldFilterExact](../../models/components/stringfieldfilterexact.md)      | :heavy_minus_sign:                                                                           | Filter using **one** of the following operators: `eq`, `oeq`, `neq`                          |