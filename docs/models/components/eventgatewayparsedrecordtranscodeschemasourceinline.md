# EventGatewayParsedRecordTranscodeSchemaSourceInline

A schema embedded directly in the policy configuration.


## Fields

| Field                                                                                 | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `Type`                                                                                | `string`                                                                              | :heavy_check_mark:                                                                    | N/A                                                                                   |
| `Inline`                                                                              | `string`                                                                              | :heavy_check_mark:                                                                    | The raw schema text (e.g. an Avro JSON schema) to use for the transcoded output data. |