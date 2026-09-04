# CreateAIGatewayConfigStoreSecretRequest


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    | Example                                                        |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `Key`                                                          | `string`                                                       | :heavy_check_mark:                                             | The unique key identifying the secret within the Config Store. | my-secret-key                                                  |
| `Value`                                                        | `string`                                                       | :heavy_check_mark:                                             | The secret value. Once stored, this value cannot be retrieved. | my-secret-value                                                |