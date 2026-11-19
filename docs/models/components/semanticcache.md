# SemanticCache

Evidence from the semantic cache lookup for the invocation.


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `Status`                                                                     | `*string`                                                                    | :heavy_minus_sign:                                                           | The semantic cache's status for the invocation, in the producer's own terms. |
| `EmbeddingsProvider`                                                         | `*string`                                                                    | :heavy_minus_sign:                                                           | The provider of the embedding model that the cache lookup used.              |
| `EmbeddingsModel`                                                            | `*string`                                                                    | :heavy_minus_sign:                                                           | The embedding model that the cache lookup used.                              |
| `EmbeddingsTokens`                                                           | `*string`                                                                    | :heavy_minus_sign:                                                           | The number of tokens that the cache lookup embedded.                         |