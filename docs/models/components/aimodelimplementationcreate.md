# AiModelImplementationCreate

Request body for linking an AI Model to an AI Gateway model.


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  | Example                                                      |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `GatewayControlPlaneID`                                      | `string`                                                     | :heavy_check_mark:                                           | The AI Gateway control plane uuid the model belongs to.      | 223e4567-e89b-12d3-a456-426614174999                         |
| `GatewayModelID`                                             | `string`                                                     | :heavy_check_mark:                                           | The text or uuid identifier of the AI Gateway model to link. | gw-model-abc                                                 |