# TargetModel

An upstream LLM target the AI Model proxies.


## Fields

| Field                                           | Type                                            | Required                                        | Description                                     | Example                                         |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| `Provider`                                      | `string`                                        | :heavy_check_mark:                              | The upstream model provider.                    | openai                                          |
| `Name`                                          | `string`                                        | :heavy_check_mark:                              | The upstream model name for the given provider. | gpt-4o                                          |