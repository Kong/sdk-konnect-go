# Message

The wording of the rule's alerts on every channel, or `null` for the built-in wording.

`subject` and `body` can contain placeholders, such as `{{budget.name}}`. A placeholder is `{{`, optional whitespace, a name from the list below, optional whitespace, and `}}`. There's no escaping, and a lone `}}` is text. Each alert fills the placeholders with its own values. Amounts appear in US dollars, such as `$4,125.50`.

- `{{budget.name}}`: the budget's name.
- `{{budget.url}}`: a link to the budget in Konnect.
- `{{threshold}}`: the threshold that was crossed, such as `80%`.
- `{{limit}}`: the limit that spend is measured against. With `each`, the limit of one derived budget.
- `{{spend}}`: the spend so far in the budget window.
- `{{percent}}`: the spend as a share of the limit, such as `82.51%`.
- `{{remaining}}`: what is left of the limit, never below `$0.00`.
- `{{projected}}`: where spend ends the budget window at the current pace, as an amount and a share of the limit.
- `{{period}}`: the budget window, in the time zone of its anchor.
- `{{period.start}}`: when the budget window started.
- `{{period.end}}`: when the budget window ends and the budget resets.
- `{{period.remaining}}`: how long is left in the budget window when the alert is sent, in its largest whole unit, such as `8 days` or `12 minutes`. Under a minute, it's `less than a minute`, and after the window ends, `no time`.
- `{{entity.field}}`: the `each` field as a noun, such as `group`, `user` for `subject`, or `env tag` for `tags.env`.
- `{{entity.value}}`: the value whose derived budget crossed the threshold. A Kong Identity principal or a consumer shows as its display name when the service can resolve it. Any other value, such as a Kong Identity group or team, shows as it is.

`{{entity.field}}` and `{{entity.value}}` work only on a budget with `each`. Defaults to `null`.


## Fields

| Field                                                                                                                                                                                                                  | Type                                                                                                                                                                                                                   | Required                                                                                                                                                                                                               | Description                                                                                                                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Subject`                                                                                                                                                                                                              | `string`                                                                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                                     | The subject line of the rule's alerts, as plain text that can contain placeholders such as `{{budget.name}}`.<br/><br/>A single line of 1 to 200 characters after the server removes leading and trailing whitespace.  |
| `Body`                                                                                                                                                                                                                 | `string`                                                                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                                     | The text of the rule's alerts, as plain text that can contain placeholders such as `{{budget.name}}`. Blank lines separate paragraphs.<br/><br/>1 to 2000 characters after the server removes leading and trailing whitespace. |