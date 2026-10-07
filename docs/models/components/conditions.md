# Conditions

A tokens will only be exchange when it matches all these criteria. To exchanging tokens issued from a different issuer, conditions must not be defined; On the contrary, to exchange tokens issued from the target issuer itself, conditions must be defined.


## Fields

| Field              | Type               | Required           | Description        |
| ------------------ | ------------------ | ------------------ | ------------------ |
| `HasAudience`      | []`string`         | :heavy_minus_sign: | N/A                |
| `HasScopes`        | []`string`         | :heavy_minus_sign: | N/A                |
| `MissingAudience`  | []`string`         | :heavy_minus_sign: | N/A                |
| `MissingScopes`    | []`string`         | :heavy_minus_sign: | N/A                |