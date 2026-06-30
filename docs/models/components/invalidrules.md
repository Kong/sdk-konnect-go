# InvalidRules

invalid parameters rules

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.InvalidRulesRequired

// Open enum: custom values can be created with a direct type cast
custom := components.InvalidRules("custom_value")
```


## Values

| Name                                                 | Value                                                |
| ---------------------------------------------------- | ---------------------------------------------------- |
| `InvalidRulesRequired`                               | required                                             |
| `InvalidRulesIsArray`                                | is_array                                             |
| `InvalidRulesIsBase64`                               | is_base64                                            |
| `InvalidRulesIsBoolean`                              | is_boolean                                           |
| `InvalidRulesIsDateTime`                             | is_date_time                                         |
| `InvalidRulesIsInteger`                              | is_integer                                           |
| `InvalidRulesIsNull`                                 | is_null                                              |
| `InvalidRulesIsNumber`                               | is_number                                            |
| `InvalidRulesIsObject`                               | is_object                                            |
| `InvalidRulesIsString`                               | is_string                                            |
| `InvalidRulesIsUUID`                                 | is_uuid                                              |
| `InvalidRulesIsFqdn`                                 | is_fqdn                                              |
| `InvalidRulesIsArn`                                  | is_arn                                               |
| `InvalidRulesUnknownProperty`                        | unknown_property                                     |
| `InvalidRulesMissingReference`                       | missing_reference                                    |
| `InvalidRulesIsLabel`                                | is_label                                             |
| `InvalidRulesMatchesRegex`                           | matches_regex                                        |
| `InvalidRulesInvalid`                                | invalid                                              |
| `InvalidRulesIsSupportedNetworkAvailabilityZoneList` | is_supported_network_availability_zone_list          |
| `InvalidRulesIsSupportedNetworkCidrBlock`            | is_supported_network_cidr_block                      |
| `InvalidRulesIsSupportedProviderRegion`              | is_supported_provider_region                         |
| `InvalidRulesType`                                   | type                                                 |