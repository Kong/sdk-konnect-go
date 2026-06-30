# URLMetadataValueTemplate

Contains the name and link value associated with a URL Metadata Field. Supports JQ templates which resolves to a string.


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        | Example                                            |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `Name`                                             | `string`                                           | :heavy_check_mark:                                 | The human-readable name for the url metadata value | #public-slack-channel                              |
| `Link`                                             | `string`                                           | :heavy_check_mark:                                 | The url link associated with the metadata value.   | #public-slack-channel                              |