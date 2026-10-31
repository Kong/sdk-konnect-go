# EventGatewayMaskStrategyKeepChars

Keeps a number of leading and/or trailing characters and replaces the middle with a fixed phrase.
If `first + last` is greater than or equal to the value length, the whole value is replaced with the phrase, so it never shows more characters than the original value. The phrase is fixed, so the masked output does not leak the value length.


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `Type`                                                   | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `First`                                                  | `*int64`                                                 | :heavy_minus_sign:                                       | Number of leading characters to keep unmasked.           |
| `Last`                                                   | `*int64`                                                 | :heavy_minus_sign:                                       | Number of trailing characters to keep unmasked.          |
| `Phrase`                                                 | `string`                                                 | :heavy_check_mark:                                       | The phrase that replaces the masked middle of the value. |