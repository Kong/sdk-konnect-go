# ConsumeValueValidationAction

Defines a behavior when record value is not valid.
* mark - marks a record with kong/server header and client ID value
  to help to identify the clients violating schema.
* skip - skips delivering a record.



## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `ConsumeValueValidationActionMark` | mark                               |
| `ConsumeValueValidationActionSkip` | skip                               |