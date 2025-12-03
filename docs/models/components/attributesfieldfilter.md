# AttributesFieldFilter

Filter using **one** of the following operators: `eq`, `oeq`, `neq`, `contains`, `ocontains`


## Fields

| Field                                          | Type                                           | Required                                       | Description                                    |
| ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- |
| `Eq`                                           | **string*                                      | :heavy_minus_sign:                             | The field exactly matches the provided value.  |
| `Contains`                                     | **string*                                      | :heavy_minus_sign:                             | The field contains the provided value.         |
| `Ocontains`                                    | **string*                                      | :heavy_minus_sign:                             | The field contains any of the provided values. |
| `Oeq`                                          | **string*                                      | :heavy_minus_sign:                             | The field matches any of the provided values.  |
| `Neq`                                          | **string*                                      | :heavy_minus_sign:                             | The field does not match the provided value.   |