# CatalogMCPPrompt

MCP Prompt definition.


## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `Name`                                                                                       | `string`                                                                                     | :heavy_check_mark:                                                                           | Machine/programmatic name of the prompt.                                                     |
| `Title`                                                                                      | `*string`                                                                                    | :heavy_minus_sign:                                                                           | Human-readable name of the prompt.                                                           |
| `Description`                                                                                | `*string`                                                                                    | :heavy_minus_sign:                                                                           | Detailed description of the prompt's purpose and usage.                                      |
| `Arguments`                                                                                  | [][components.CatalogMCPPromptArgument](../../models/components/catalogmcppromptargument.md) | :heavy_minus_sign:                                                                           | List of arguments to use for templating the prompt.                                          |