# ScorecardScore

The current score for the given Scorecard.
A `null` value indicates the scorecard has not yet been evaluated and therefore no score has been computed.



## Fields

| Field                                                                                               | Type                                                                                                | Required                                                                                            | Description                                                                                         | Example                                                                                             |
| --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| `Value`                                                                                             | *string*                                                                                            | :heavy_check_mark:                                                                                  | The human-readable score value coverted to the Scorecard's assigned grading system.                 | 88%                                                                                                 |
| `RawValue`                                                                                          | *float64*                                                                                           | :heavy_check_mark:                                                                                  | The raw, numeric score calculated during evaluation of the scorecard.<br/>Rounded to 3 decimal places.<br/> | 87.5                                                                                                |