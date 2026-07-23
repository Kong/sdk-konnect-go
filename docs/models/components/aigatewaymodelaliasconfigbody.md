# AIGatewayModelAliasConfigBody

Configuration for routing requests to a specific model using a request body property.



## Fields

| Field                                                                                               | Type                                                                                                | Required                                                                                            | Description                                                                                         | Example                                                                                             |
| --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| `Body`                                                                                              | map[string]`any`                                                                                    | :heavy_check_mark:                                                                                  | Value indexed by property name that will cause this route to match if present in the request body.<br/> | {<br/>"model": [<br/>"@azure/claude-sonnet-5"<br/>]<br/>}                                           |