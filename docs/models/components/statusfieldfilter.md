# StatusFieldFilter

Filter using **one** of the following operators: `eq`, `neq`, `contains`


## Fields

| Field                                         | Type                                          | Required                                      | Description                                   | Example                                       |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| `Eq`                                          | `*string`                                     | :heavy_minus_sign:                            | The field exactly matches the provided value. | ?filter[status][eq]=failed                    |
| `Neq`                                         | `*string`                                     | :heavy_minus_sign:                            | The field does not match the provided value.  | ?filter[status][neq]=failed                   |
| `Contains`                                    | `*string`                                     | :heavy_minus_sign:                            | The field contains the provided value.        | ?filter[status][contains]=fail                |