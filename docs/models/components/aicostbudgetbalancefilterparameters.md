# AICostBudgetBalanceFilterParameters

Filters on the balances' fields. Every filter must match, and a bare value means `eq`.

- `budget.id`: `eq`, `neq`, and `oeq`.
- `scope.subject`, `scope.group`, and `scope.team`: `eq` and `oeq`, with complete typed references: subjects as `principal/{uuid}`, `consumer/{uuid}`, or `oidc/{sub}`, groups as `principal/{uuid}`, `consumer_group/{uuid}`, or `oidc/{group}`, and teams as `principal/{uuid}`, where `{uuid}` is never the nil UUID. A Kong Identity group goes under `scope.group`, and a Kong Identity team under `scope.team`. A budget matches when a branch of its scope names a listed reference in a `subject`, `group`, or `team` predicate with `eq` or `in`, or when a listed reference is one of its derived values: a subject on a budget whose `each` is `subject` or `subject_id`, a group on a budget whose `each` is `group`, or a team on a budget whose `each` is `team`. A pooled budget whose scope names no subject, group, or team, such as an organization-wide budget, never matches. `oeq` can't list a reference that contains a comma.

A budget that `scope.subject`, `scope.group`, or `scope.team` matches through a listed derived value lists only the balances of the listed values it admits, with zero spend without usage. Every other matching budget lists all its balances.

Any other filter key fails with `400`. UUIDs compare case-insensitively.


## Fields

| Field                                                                                           | Type                                                                                            | Required                                                                                        | Description                                                                                     |
| ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- |
| `BudgetID`                                                                                      | [*components.UUIDFieldFilter](../../models/components/uuidfieldfilter.md)                       | :heavy_minus_sign:                                                                              | Filter using **one** of the following operators: `eq`, `oeq`, `neq`                             |
| `ScopeSubject`                                                                                  | [*components.AICostReferenceFieldFilter](../../models/components/aicostreferencefieldfilter.md) | :heavy_minus_sign:                                                                              | Filter using **one** of the following operators: `eq`, `oeq`                                    |
| `ScopeGroup`                                                                                    | [*components.AICostReferenceFieldFilter](../../models/components/aicostreferencefieldfilter.md) | :heavy_minus_sign:                                                                              | Filter using **one** of the following operators: `eq`, `oeq`                                    |
| `ScopeTeam`                                                                                     | [*components.AICostReferenceFieldFilter](../../models/components/aicostreferencefieldfilter.md) | :heavy_minus_sign:                                                                              | Filter using **one** of the following operators: `eq`, `oeq`                                    |