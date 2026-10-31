# GetAiGatewayAgentResponse


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `ContentType`                                                           | `string`                                                                | :heavy_check_mark:                                                      | HTTP response content type for this operation                           |
| `StatusCode`                                                            | `int`                                                                   | :heavy_check_mark:                                                      | HTTP response status code for this operation                            |
| `RawResponse`                                                           | [*http.Response](https://pkg.go.dev/net/http#Response)                  | :heavy_check_mark:                                                      | Raw HTTP response; suitable for custom response parsing                 |
| `AIGatewayAgent`                                                        | [*components.AIGatewayAgent](../../models/components/aigatewayagent.md) | :heavy_minus_sign:                                                      | A successful response returning an AI Gateway Agent.                    |