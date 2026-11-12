# BillingCreditGrantPlanFilterVersion

Omission matches all versions, including future versions.


## Fields

| Field                                    | Type                                     | Required                                 | Description                              |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| `Eq`                                     | `*int`                                   | :heavy_minus_sign:                       | Match this exact version.                |
| `Oeq`                                    | []`int`                                  | :heavy_minus_sign:                       | Match one of these versions.             |
| `Gte`                                    | `*int`                                   | :heavy_minus_sign:                       | Match this version and later versions.   |
| `Lte`                                    | `*int`                                   | :heavy_minus_sign:                       | Match this version and earlier versions. |