# PartialAppAuthStrategyConfigOpenIDConnect

A more advanced mode to configure an API Product Version’s Application Auth Strategy. 
Using this mode will allow developers to use API credentials issued from an external IdP that will authenticate their application requests. 
Once authenticated, an application will be granted access to any Product Version it is registered for that is configured for the same Auth Strategy. 
An OIDC strategy may be used in conjunction with a DCR provider to automatically create the IdP application.



## Fields

| Field                  | Type                   | Required               | Description            |
| ---------------------- | ---------------------- | ---------------------- | ---------------------- |
| `Issuer`               | **string*              | :heavy_minus_sign:     | N/A                    |
| `CredentialClaim`      | []*string*             | :heavy_minus_sign:     | N/A                    |
| `Scopes`               | []*string*             | :heavy_minus_sign:     | N/A                    |
| `AuthMethods`          | []*string*             | :heavy_minus_sign:     | N/A                    |
| `AdditionalProperties` | map[string]*any*       | :heavy_minus_sign:     | N/A                    |