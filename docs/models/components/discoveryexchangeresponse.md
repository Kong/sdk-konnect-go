# DiscoveryExchangeResponse

The organization and user an organization discovery code was minted for.


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        | Example                                                            |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `OrgID`                                                            | `string`                                                           | :heavy_check_mark:                                                 | The organization's unique identifier, from the code's `oid` claim. | 0e6f1b4a-4a2a-4f1e-9c1a-1f2b3c4d5e6f                               |
| `UserID`                                                           | `string`                                                           | :heavy_check_mark:                                                 | The user's unique identifier, from the code's `uid` claim.         | 9b8c7d6e-5f4a-3b2c-1d0e-9f8a7b6c5d4e                               |