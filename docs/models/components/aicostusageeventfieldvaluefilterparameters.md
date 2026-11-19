# AICostUsageEventFieldValueFilterParameters

Filters on the returned values. Every filter must match, and a bare value means `eq`.

- `value`: `eq`, `neq`, `oeq`, `contains`, and `ocontains`, for every `field`, including the gateway ID fields. With `field=tags`, each value, including each part of an `oeq` or `ocontains` list, has at most 256 characters, none of them a control character.
- `type`: `eq`, `neq`, and `oeq`, with a subject type. With any `field` other than `subject`, a `type` filter fails with `400`.

Any other filter key fails with `400`. `eq`, `neq`, and `oeq` compare strings case-sensitively, and `contains` and `ocontains` ignore case. For a UUID field, `eq`, `neq`, and `oeq` compare full UUIDs case-insensitively.


## Fields

| Field                                                                                                                          | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `Value`                                                                                                                        | [*components.StringFieldFilter](../../models/components/stringfieldfilter.md)                                                  | :heavy_minus_sign:                                                                                                             | Filter using **one** of the following operators: `eq`, `oeq`, `neq`, `contains`, `ocontains`                                   |
| `Type`                                                                                                                         | [*components.AICostObservedSubjectTypeFilter](../../models/components/aicostobservedsubjecttypefilter.md)                      | :heavy_minus_sign:                                                                                                             | Filters a field by its value with `eq`, `neq`, or `oeq`. Accepts `principal`, `consumer`, and `oidc`. A bare value means `eq`. |