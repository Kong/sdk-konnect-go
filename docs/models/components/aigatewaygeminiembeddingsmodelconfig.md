# AIGatewayGeminiEmbeddingsModelConfig

Google Gemini-specific configuration for a model.


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `UpstreamURL`                                                           | `*string`                                                               | :heavy_minus_sign:                                                      | The URL of the embeddings model.                                        |
| `Type`                                                                  | `string`                                                                | :heavy_check_mark:                                                      | N/A                                                                     |
| `GcpEnvironment`                                                        | [*components.GCPModelConfig](../../models/components/gcpmodelconfig.md) | :heavy_minus_sign:                                                      | Configuration for a model hosted on Google Cloud Project.               |