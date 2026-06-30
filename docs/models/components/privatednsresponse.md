# PrivateDNSResponse


## Supported Types

### AwsPrivateHostedZoneResponse

```go
privateDNSResponse := components.CreatePrivateDNSResponseAwsPrivateHostedZoneResponse(components.AwsPrivateHostedZoneResponse{/* values here */})
```

### AwsPrivateDNSResolverResponse

```go
privateDNSResponse := components.CreatePrivateDNSResponseAwsPrivateDNSResolverResponse(components.AwsPrivateDNSResolverResponse{/* values here */})
```

### GcpPrivateHostedZoneResponse

```go
privateDNSResponse := components.CreatePrivateDNSResponseGcpPrivateHostedZoneResponse(components.GcpPrivateHostedZoneResponse{/* values here */})
```

### AzurePrivateHostedZoneResponse

```go
privateDNSResponse := components.CreatePrivateDNSResponseAzurePrivateHostedZoneResponse(components.AzurePrivateHostedZoneResponse{/* values here */})
```

### AzurePrivateDNSResolverResponse

```go
privateDNSResponse := components.CreatePrivateDNSResponseAzurePrivateDNSResolverResponse(components.AzurePrivateDNSResolverResponse{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch privateDNSResponse.Type {
	case components.PrivateDNSResponseTypeAwsPrivateHostedZoneResponse:
		// privateDNSResponse.AwsPrivateHostedZoneResponse is populated
	case components.PrivateDNSResponseTypeAwsPrivateDNSResolverResponse:
		// privateDNSResponse.AwsPrivateDNSResolverResponse is populated
	case components.PrivateDNSResponseTypeGcpPrivateHostedZoneResponse:
		// privateDNSResponse.GcpPrivateHostedZoneResponse is populated
	case components.PrivateDNSResponseTypeAzurePrivateHostedZoneResponse:
		// privateDNSResponse.AzurePrivateHostedZoneResponse is populated
	case components.PrivateDNSResponseTypeAzurePrivateDNSResolverResponse:
		// privateDNSResponse.AzurePrivateDNSResolverResponse is populated
}
```
