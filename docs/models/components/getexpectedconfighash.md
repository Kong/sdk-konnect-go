# GetExpectedConfigHash

Response body for retrieving the expected config hash of the control plane.


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `ExpectedHash`                                         | **string*                                              | :heavy_minus_sign:                                     | The expected configuration hash.                       |
| `CreatedAt`                                            | **int64*                                               | :heavy_minus_sign:                                     | Date the control plane configuration was created.      |
| `UpdatedAt`                                            | **int64*                                               | :heavy_minus_sign:                                     | Date the control plane configuration was last updated. |