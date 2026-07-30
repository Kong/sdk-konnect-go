# AIGatewayTargetSagemakerConfigAws

**Pre-release Feature**
This feature is currently in beta and is subject to change.


## Fields

| Field                                                                                       | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `Region`                                                                                    | `*string`                                                                                   | :heavy_minus_sign:                                                                          | Overrides the AWS_REGION environment variable for SageMaker requests.                       |
| `AssumeRoleArn`                                                                             | `*string`                                                                                   | :heavy_minus_sign:                                                                          | Assume a different IAM role after authenticating; mutually required with role_session_name. |
| `RoleSessionName`                                                                           | `*string`                                                                                   | :heavy_minus_sign:                                                                          | Session identifier for the assumed role; mutually required with assume_role_arn.            |
| `StsEndpointURL`                                                                            | `*string`                                                                                   | :heavy_minus_sign:                                                                          | Overrides the STS endpoint when assuming a role.                                            |