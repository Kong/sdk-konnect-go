# Optimization

Evidence about cost-saving mechanisms applied to the invocation: a semantic cache and prompt compression. Nothing in it is billable or enters a cost.


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `SemanticCache`                                                       | [*components.SemanticCache](../../models/components/semanticcache.md) | :heavy_minus_sign:                                                    | Evidence from the semantic cache lookup for the invocation.           |
| `Compression`                                                         | [*components.Compression](../../models/components/compression.md)     | :heavy_minus_sign:                                                    | Evidence from prompt compression for the invocation.                  |