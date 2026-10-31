# Runtime

Runtime configuration for executing the package


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               | Example                                                                   |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `Hint`                                                                    | `*string`                                                                 | :heavy_minus_sign:                                                        | A hint to help clients determine the appropriate runtime for the package. | npx                                                                       |
| `Arguments`                                                               | [][components.MCPArgument](../../models/components/mcpargument.md)        | :heavy_minus_sign:                                                        | A list of arguments to be passed to the package's runtime command.        |                                                                           |