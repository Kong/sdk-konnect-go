# AuthGCPWorkloadIdentityFederationAwsIamOutput

Enables authenticating with GCP via Workload Identity Federation, obtaining temporary
GCP credentials instead of using a static service account key.



## Fields

| Field                                                                                                   | Type                                                                                                    | Required                                                                                                | Description                                                                                             | Example                                                                                                 |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| `Source`                                                                                                | [components.Source](../../models/components/source.md)                                                  | :heavy_check_mark:                                                                                      | The source identity provider/system used to obtain temporary GCP credentials.<br/>                      |                                                                                                         |
| `Aws`                                                                                                   | [*components.AIGatewayUpstreamAuthAWSOutput](../../models/components/aigatewayupstreamauthawsoutput.md) | :heavy_minus_sign:                                                                                      | AWS IAM (SigV4) authentication for the upstream service.<br/>                                           | {<br/>"type": "aws",<br/>"region": "us-east-1",<br/>"access_key_id": "AKIAIOSFODNN7EXAMPLE"<br/>}       |