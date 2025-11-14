# PortalEmailConfig

Portal email config.


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   | Example                                                       |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `ID`                                                          | *string*                                                      | :heavy_check_mark:                                            | N/A                                                           | 7f9fd312-a987-4628-b4c5-bb4f4fddd5f7                          |
| `DomainName`                                                  | **string*                                                     | :heavy_minus_sign:                                            | The domain name used for sending emails. Null means default.  |                                                               |
| `FromName`                                                    | *string*                                                      | :heavy_check_mark:                                            | The name displayed in the 'From' field of emails.             |                                                               |
| `FromEmail`                                                   | *string*                                                      | :heavy_check_mark:                                            | The email address used in the 'From' field.                   |                                                               |
| `ReplyToEmail`                                                | *string*                                                      | :heavy_check_mark:                                            | The email address used in the 'Reply-To' field.               |                                                               |
| `CreatedAt`                                                   | [time.Time](https://pkg.go.dev/time#Time)                     | :heavy_check_mark:                                            | An ISO-8601 timestamp representation of entity creation date. | 2022-11-04T20:10:06.927Z                                      |
| `UpdatedAt`                                                   | [time.Time](https://pkg.go.dev/time#Time)                     | :heavy_check_mark:                                            | An ISO-8601 timestamp representation of entity update date.   | 2022-11-04T20:10:06.927Z                                      |