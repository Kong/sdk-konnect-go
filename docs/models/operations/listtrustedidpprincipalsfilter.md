# ListTrustedIdpPrincipalsFilter

Filter principals by field. Supported filter fields:
- `display_name`: exact match or contains



## Fields

| Field                                                                                     | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `DisplayName`                                                                             | [*components.LegacyStringFieldFilter](../../models/components/legacystringfieldfilter.md) | :heavy_minus_sign:                                                                        | Filter using **one** of the following operators: `eq`, `contains`                         |