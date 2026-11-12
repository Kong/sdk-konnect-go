# MeteringIngestedEventValidationError

Event validation errors.


## Fields

| Field                                        | Type                                         | Required                                     | Description                                  |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| `Code`                                       | `string`                                     | :heavy_check_mark:                           | The machine readable code of the error.      |
| `Message`                                    | `string`                                     | :heavy_check_mark:                           | The human readable description of the error. |
| `Attributes`                                 | map[string]`any`                             | :heavy_minus_sign:                           | Additional attributes.                       |