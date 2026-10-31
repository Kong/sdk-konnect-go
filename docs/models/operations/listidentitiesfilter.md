# ListIdentitiesFilter

Filter identities by field. Supported filter field:
- `type`: exact match on identity type. Accepted values: `oidc`, `consumer`, `custom`, `auth_server_client`.

Example: `filter[type]="oidc"`



## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `Type`                                                                  | [*operations.QueryParamType](../../models/operations/queryparamtype.md) | :heavy_minus_sign:                                                      | Filter using **one** of the following operators: `eq`, `oeq`            |