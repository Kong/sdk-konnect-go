# CatalogMCPToolAnnotations


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `Title`                                                 | `*string`                                               | :heavy_minus_sign:                                      | Human-readable name of the tool.                        |
| `DestructiveHint`                                       | `*bool`                                                 | :heavy_minus_sign:                                      | Indicates if the tool performs destructive actions.     |
| `IdempotentHint`                                        | `*bool`                                                 | :heavy_minus_sign:                                      | Indicates if the tool is idempotent.                    |
| `OpenWorldHint`                                         | `*bool`                                                 | :heavy_minus_sign:                                      | Indicates if the tool interacts with external entities. |
| `ReadOnlyHint`                                          | `*bool`                                                 | :heavy_minus_sign:                                      | Indicates if the tool is read-only.                     |