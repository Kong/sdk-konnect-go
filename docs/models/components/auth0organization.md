# Auth0Organization

A single auth0 organization.


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            | Example                                                                |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `Name`                                                                 | `*string`                                                              | :heavy_minus_sign:                                                     | The name of the organization.                                          | My Organization                                                        |
| `LoginPath`                                                            | `*string`                                                              | :heavy_minus_sign:                                                     | The login path of the organization.                                    | /v2/authenticate/federated?org_id=7f9fd312-a987-4628-b4c5-bb4f4fddd5f7 |