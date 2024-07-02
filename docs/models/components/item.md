# Item


## Fields

| Field                                                                                         | Type                                                                                          | Required                                                                                      | Description                                                                                   | Example                                                                                       |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `LuaSchema`                                                                                   | **string*                                                                                     | :heavy_minus_sign:                                                                            | The custom plugin schema; `jq -Rs '.' schema.lua`.                                            | return { name = \"myplugin\", fields = { { config = { type = \"record\", fields = { } } } } } |
| `Name`                                                                                        | **string*                                                                                     | :heavy_minus_sign:                                                                            | The custom plugin name determined by the custom plugin schema.                                | myplugin                                                                                      |
| `CreatedAt`                                                                                   | **int64*                                                                                      | :heavy_minus_sign:                                                                            | An ISO-8604 timestamp representation of custom plugin schema creation date.                   | 1422386534                                                                                    |
| `UpdatedAt`                                                                                   | **int64*                                                                                      | :heavy_minus_sign:                                                                            | An ISO-8604 timestamp representation of custom plugin schema update date.                     | 1422412345                                                                                    |