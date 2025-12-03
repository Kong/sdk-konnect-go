# VirtualClusterAuthenticationClaimsMapping

Maps JWT claims in the case when sub and scope are presented as different claims in your JWT token.


## Fields

| Field                   | Type                    | Required                | Description             |
| ----------------------- | ----------------------- | ----------------------- | ----------------------- |
| `Sub`                   | **string*               | :heavy_minus_sign:      | Maps the subject claim. |
| `Scope`                 | **string*               | :heavy_minus_sign:      | Maps the scope claim.   |