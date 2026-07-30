# GcpEnvironment

**Pre-release Feature**
This feature is currently in beta and is subject to change.

Configuration for a model hosted on Google Cloud Project.


## Fields

| Field                                                                                                          | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `APIEndpoint`                                                                                                  | `string`                                                                                                       | :heavy_check_mark:                                                                                             | The custom API endpoint for the Gemini model.                                                                  |
| `LocationID`                                                                                                   | `string`                                                                                                       | :heavy_check_mark:                                                                                             | The Google Cloud location ID for the model endpoint.                                                           |
| `ProjectID`                                                                                                    | `string`                                                                                                       | :heavy_check_mark:                                                                                             | The Google Cloud project ID for the model endpoint.                                                            |
| `EndpointID`                                                                                                   | `*string`                                                                                                      | :heavy_minus_sign:                                                                                             | The endpoint ID for the model.<br/>This must be set when running a target model on Gemini on Vertex Model Garden.<br/> |