# ProduceValueValidationAction

Defines a behavior when record value is not valid.
* reject - rejects a batch for topic partition. Only available for produce.
* mark - marks a record with kong/server header and client ID value


  to help to identify the clients violating schema.



## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `ProduceValueValidationActionReject` | reject                               |
| `ProduceValueValidationActionMark`   | mark                                 |