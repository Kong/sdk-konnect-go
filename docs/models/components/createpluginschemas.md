# CreatePluginSchemas


## Fields

| Field                                                                                         | Type                                                                                          | Required                                                                                      | Description                                                                                   | Example                                                                                       |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `LuaSchema`                                                                                   | *string*                                                                                      | :heavy_check_mark:                                                                            | The custom plugin schema; `jq -Rs '.' schema.lua`.<br/>                                       | return { name = \"myplugin\", fields = { { config = { type = \"record\", fields = { } } } } } |