# UserDefinedRules

User-defined payload sanitization rules


## Fields

| Field                                                                                    | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `Headers`                                                                                | [][components.HeaderSanitizationRule](../../models/components/headersanitizationrule.md) | :heavy_minus_sign:                                                                       | Rules for sanitizing HTTP headers                                                        |
| `Body`                                                                                   | [][components.BodySanitizationRule](../../models/components/bodysanitizationrule.md)     | :heavy_minus_sign:                                                                       | Rules for sanitizing request/response body content                                       |