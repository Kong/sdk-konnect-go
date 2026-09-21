# KeepChars


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `First`                                                  | `*int64`                                                 | :heavy_minus_sign:                                       | Number of leading characters to keep unmasked.           |
| `Last`                                                   | `*int64`                                                 | :heavy_minus_sign:                                       | Number of trailing characters to keep unmasked.          |
| `Phrase`                                                 | `string`                                                 | :heavy_check_mark:                                       | The phrase that replaces the masked middle of the value. |