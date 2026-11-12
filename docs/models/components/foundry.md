# Foundry

Endpoint configuration for Azure AI Foundry hosted models. Required when
`service` is `azure-foundry`.



## Fields

| Field                                          | Type                                           | Required                                       | Description                                    | Example                                        |
| ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- |
| `Resource`                                     | `string`                                       | :heavy_check_mark:                             | The Azure AI Foundry resource name.            | kong-foundry-east                              |
| `Domain`                                       | `*string`                                      | :heavy_minus_sign:                             | The domain for Azure AI Foundry hosted models. |                                                |