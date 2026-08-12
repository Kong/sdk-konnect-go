# ListTeamsQueryParamFilter

Filter teams returned in the response. Supports filtering by label value using
dot-notation, e.g. `filter[labels.<key>][<op>]=<value>`, where `<op>` is one of
`eq`, `contains`, or `exists`.


## Fields

| Field                                                                                                                   | Type                                                                                                                    | Required                                                                                                                | Description                                                                                                             |
| ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- |
| `Name`                                                                                                                  | [*components.LegacyStringFieldFilter](../../models/components/legacystringfieldfilter.md)                               | :heavy_minus_sign:                                                                                                      | Filter using **one** of the following operators: `eq`, `contains`                                                       |
| `Labels`                                                                                                                | map[string][components.LegacyStringFieldFilterWithExists](../../models/components/legacystringfieldfilterwithexists.md) | :heavy_minus_sign:                                                                                                      | N/A                                                                                                                     |