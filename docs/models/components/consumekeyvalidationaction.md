# ConsumeKeyValidationAction

Defines a behavior when record key is not valid.
* mark - marks a record with kong/server header and client ID value
  to help to identify the clients violating schema.
* skip - skips delivering a record.



## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `ConsumeKeyValidationActionMark` | mark                             |
| `ConsumeKeyValidationActionSkip` | skip                             |