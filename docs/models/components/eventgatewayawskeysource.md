# EventGatewayAWSKeySource

A key source that uses an AWS KMS to find a symmetric key. Load KMS credentials from the environment.

See [aws docs](https://docs.aws.amazon.com/sdk-for-rust/latest/dg/credproviders.html#credproviders-default-credentials-provider-chain)
for more information about how credential retrieval.



## Fields

| Field              | Type               | Required           | Description        |
| ------------------ | ------------------ | ------------------ | ------------------ |
| `Type`             | *string*           | :heavy_check_mark: | N/A                |