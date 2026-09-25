# BillingValidationIssue

A validation issue found while processing a billing resource.


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `Code`                                                                  | `string`                                                                | :heavy_check_mark:                                                      | Machine-readable error code.                                            |
| `Message`                                                               | `string`                                                                | :heavy_check_mark:                                                      | Human-readable description of the error.                                |
| `Attributes`                                                            | map[string]`any`                                                        | :heavy_minus_sign:                                                      | Additional structured context.                                          |
| `Severity`                                                              | [components.Severity](../../models/components/severity.md)              | :heavy_check_mark:                                                      | Severity of the validation issue.                                       |
| `Field`                                                                 | `*string`                                                               | :heavy_minus_sign:                                                      | JSON path to the field that caused the validation issue, if applicable. |
| `Component`                                                             | `*string`                                                               | :heavy_minus_sign:                                                      | Component that reported the validation issue, if applicable.            |