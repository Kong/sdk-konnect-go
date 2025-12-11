# ~~PortalOIDCConfig~~

Configuration properties for an OpenID Connect Identity Provider.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       | Example                                                                           |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `Issuer`                                                                          | *string*                                                                          | :heavy_check_mark:                                                                | N/A                                                                               | https://identity.example.com/v2                                                   |
| `ClientID`                                                                        | *string*                                                                          | :heavy_check_mark:                                                                | N/A                                                                               | x7id0o42lklas0blidl2                                                              |
| `Scopes`                                                                          | []*string*                                                                        | :heavy_minus_sign:                                                                | N/A                                                                               | [<br/>"email",<br/>"openid",<br/>"profile"<br/>]                                  |
| `ClaimMappings`                                                                   | [*components.PortalClaimMappings](../../models/components/portalclaimmappings.md) | :heavy_minus_sign:                                                                | Mappings from a portal developer atribute to an Identity Provider claim.          | {<br/>"name": "name",<br/>"email": "email",<br/>"groups": "custom-group-claim"<br/>} |