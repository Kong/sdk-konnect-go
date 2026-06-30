# PrivateDNSAttachmentConfig


## Supported Types

### AwsPrivateHostedZoneAttachmentConfig

```go
privateDNSAttachmentConfig := components.CreatePrivateDNSAttachmentConfigAwsPrivateHostedZoneAttachmentConfig(components.AwsPrivateHostedZoneAttachmentConfig{/* values here */})
```

### AwsPrivateDNSResolverAttachmentConfig

```go
privateDNSAttachmentConfig := components.CreatePrivateDNSAttachmentConfigAwsPrivateDNSResolverAttachmentConfig(components.AwsPrivateDNSResolverAttachmentConfig{/* values here */})
```

### GcpPrivateHostedZoneAttachmentConfig

```go
privateDNSAttachmentConfig := components.CreatePrivateDNSAttachmentConfigGcpPrivateHostedZoneAttachmentConfig(components.GcpPrivateHostedZoneAttachmentConfig{/* values here */})
```

### AzurePrivateHostedZoneAttachmentConfig

```go
privateDNSAttachmentConfig := components.CreatePrivateDNSAttachmentConfigAzurePrivateHostedZoneAttachmentConfig(components.AzurePrivateHostedZoneAttachmentConfig{/* values here */})
```

### AzurePrivateDNSResolverAttachmentConfig

```go
privateDNSAttachmentConfig := components.CreatePrivateDNSAttachmentConfigAzurePrivateDNSResolverAttachmentConfig(components.AzurePrivateDNSResolverAttachmentConfig{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch privateDNSAttachmentConfig.Type {
	case components.PrivateDNSAttachmentConfigTypeAwsPrivateHostedZoneAttachmentConfig:
		// privateDNSAttachmentConfig.AwsPrivateHostedZoneAttachmentConfig is populated
	case components.PrivateDNSAttachmentConfigTypeAwsPrivateDNSResolverAttachmentConfig:
		// privateDNSAttachmentConfig.AwsPrivateDNSResolverAttachmentConfig is populated
	case components.PrivateDNSAttachmentConfigTypeGcpPrivateHostedZoneAttachmentConfig:
		// privateDNSAttachmentConfig.GcpPrivateHostedZoneAttachmentConfig is populated
	case components.PrivateDNSAttachmentConfigTypeAzurePrivateHostedZoneAttachmentConfig:
		// privateDNSAttachmentConfig.AzurePrivateHostedZoneAttachmentConfig is populated
	case components.PrivateDNSAttachmentConfigTypeAzurePrivateDNSResolverAttachmentConfig:
		// privateDNSAttachmentConfig.AzurePrivateDNSResolverAttachmentConfig is populated
}
```
