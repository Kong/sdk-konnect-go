# AIGatewayPolicySchema

A successful response returning the AI Gateway Policy Schema.


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              | Example                                                                  |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `Name`                                                                   | `string`                                                                 | :heavy_check_mark:                                                       | The plugin name (e.g. "key-auth", "cors").                               | key-auth                                                                 |
| `Fields`                                                                 | [][components.Fields](../../models/components/fields.md)                 | :heavy_check_mark:                                                       | List of the plugin's fields                                              |                                                                          |
| `EntityChecks`                                                           | [][components.EntityChecks](../../models/components/entitychecks.md)     | :heavy_minus_sign:                                                       | Entity-level (cross-field) validation rules applied to the whole record. |                                                                          |
| `AdditionalProperties`                                                   | map[string]`any`                                                         | :heavy_minus_sign:                                                       | N/A                                                                      |                                                                          |