# AIGatewayTargetSagemakerConfigTarget


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `Model`                                                                       | `*string`                                                                     | :heavy_minus_sign:                                                            | Sets the X-Amzn-SageMaker-Target-Model header (multi-model endpoints).        |
| `Variant`                                                                     | `*string`                                                                     | :heavy_minus_sign:                                                            | Sets the X-Amzn-SageMaker-Target-Variant header (A/B variant testing).        |
| `ContainerHostname`                                                           | `*string`                                                                     | :heavy_minus_sign:                                                            | Sets the X-Amzn-SageMaker-Target-Container-Hostname header (multi-container). |