# AuthGCPWorkloadIdentityFederation

Config for GCP Workload Identity Federation.



## Supported Types

### AuthGCPWorkloadIdentityFederationAwsIam

```go
authGCPWorkloadIdentityFederation := components.CreateAuthGCPWorkloadIdentityFederationAwsIam(components.AuthGCPWorkloadIdentityFederationAwsIam{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch authGCPWorkloadIdentityFederation.Type {
	case components.AuthGCPWorkloadIdentityFederationTypeAwsIam:
		// authGCPWorkloadIdentityFederation.AuthGCPWorkloadIdentityFederationAwsIam is populated
}
```
