# AuthGCPWorkloadIdentityFederationOutput

Config for GCP Workload Identity Federation.



## Supported Types

### AuthGCPWorkloadIdentityFederationAwsIamOutput

```go
authGCPWorkloadIdentityFederationOutput := components.CreateAuthGCPWorkloadIdentityFederationOutputAwsIam(components.AuthGCPWorkloadIdentityFederationAwsIamOutput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch authGCPWorkloadIdentityFederationOutput.Type {
	case components.AuthGCPWorkloadIdentityFederationOutputTypeAwsIam:
		// authGCPWorkloadIdentityFederationOutput.AuthGCPWorkloadIdentityFederationAwsIamOutput is populated
}
```
