# SystemAccount

Schema of the system account.


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `ID`                                                   | **string*                                              | :heavy_minus_sign:                                     | ID of the system account.                              |
| `Name`                                                 | **string*                                              | :heavy_minus_sign:                                     | Name of the system account.                            |
| `Description`                                          | **string*                                              | :heavy_minus_sign:                                     | Description of the system account.                     |
| `CreatedAt`                                            | [*time.Time](https://pkg.go.dev/time#Time)             | :heavy_minus_sign:                                     | Timestamp of when the system account was created.      |
| `UpdatedAt`                                            | [*time.Time](https://pkg.go.dev/time#Time)             | :heavy_minus_sign:                                     | Timestamp of when the system account was last updated. |
| `KonnectManaged`                                       | **bool*                                                | :heavy_minus_sign:                                     | The system account is managed by Konnect (true/false). |