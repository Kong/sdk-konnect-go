# DataPlaneClientCertificate


## Fields

| Field                                   | Type                                    | Required                                | Description                             |
| --------------------------------------- | --------------------------------------- | --------------------------------------- | --------------------------------------- |
| `ID`                                    | **string*                               | :heavy_minus_sign:                      | Unique ID of the certificate entity.    |
| `CreatedAt`                             | **int64*                                | :heavy_minus_sign:                      | Date certificate was created.           |
| `UpdatedAt`                             | **int64*                                | :heavy_minus_sign:                      | Date certificate was last updated.      |
| `Cert`                                  | **string*                               | :heavy_minus_sign:                      | JSON escaped string of the certificate. |