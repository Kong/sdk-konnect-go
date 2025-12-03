# MutableCondition

Defines the condition under which this configuration field is allowed to be modified.
When specified, the platform will restrict updates to this field unless the integration
meets the given condition.

For example, setting `mutable_condition: unauthorized` means the field can only be
changed while the integration is in an unauthorized state.



## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `MutableConditionUnauthorized` | unauthorized                   |