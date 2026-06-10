# ProductCatalogValidationError

Validation errors providing detailed description of the issue.


## Fields

| Field                                    | Type                                     | Required                                 | Description                              | Example                                  |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| `Code`                                   | `string`                                 | :heavy_check_mark:                       | Machine-readable error code.             |                                          |
| `Message`                                | `string`                                 | :heavy_check_mark:                       | Human-readable description of the error. |                                          |
| `Attributes`                             | map[string]`any`                         | :heavy_minus_sign:                       | Additional structured context.           |                                          |
| `Field`                                  | `string`                                 | :heavy_check_mark:                       | The path to the field.                   | addons/pro/ratecards/token/featureKey    |