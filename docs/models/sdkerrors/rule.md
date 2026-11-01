# Rule

invalid parameters rules

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/sdkerrors"
)

value := sdkerrors.RuleRequired
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `RuleRequired`                               | required                                     |
| `RuleUnique`                                 | unique                                       |
| `RuleIsArray`                                | is_array                                     |
| `RuleIsBase64`                               | is_base64                                    |
| `RuleIsBoolean`                              | is_boolean                                   |
| `RuleIsDateTime`                             | is_date_time                                 |
| `RuleIsInteger`                              | is_integer                                   |
| `RuleIsNull`                                 | is_null                                      |
| `RuleIsNumber`                               | is_number                                    |
| `RuleIsObject`                               | is_object                                    |
| `RuleIsString`                               | is_string                                    |
| `RuleIsUUID`                                 | is_uuid                                      |
| `RuleIsFqdn`                                 | is_fqdn                                      |
| `RuleIsArn`                                  | is_arn                                       |
| `RuleUnknownProperty`                        | unknown_property                             |
| `RuleMissingReference`                       | missing_reference                            |
| `RuleIsLabel`                                | is_label                                     |
| `RuleKeyInvalid`                             | key_invalid                                  |
| `RuleMatchesRegex`                           | matches_regex                                |
| `RuleInvalid`                                | invalid                                      |
| `RuleIsSupportedNetworkAvailabilityZoneList` | is_supported_network_availability_zone_list  |
| `RuleIsSupportedNetworkCidrBlock`            | is_supported_network_cidr_block              |
| `RuleIsSupportedProviderRegion`              | is_supported_provider_region                 |
| `RuleType`                                   | type                                         |