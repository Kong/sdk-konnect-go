# PostPortalEmailConfig

Create a portal email configuration.


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `DomainName`                                                   | **string*                                                      | :heavy_minus_sign:                                             | The domain name to use for sending emails. Null means default. |
| `FromName`                                                     | **string*                                                      | :heavy_minus_sign:                                             | The name to display in the 'From' field of emails.             |
| `FromEmail`                                                    | **string*                                                      | :heavy_minus_sign:                                             | The email address to use in the 'From' field.                  |
| `ReplyToEmail`                                                 | **string*                                                      | :heavy_minus_sign:                                             | The email address to use in the 'Reply-To' field.              |