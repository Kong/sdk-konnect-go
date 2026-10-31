# ForwardToVirtualClusterPolicyConfig

The configuration of the policy.


## Supported Types

### ForwardToClusterByPortMappingConfig

```go
forwardToVirtualClusterPolicyConfig := components.CreateForwardToVirtualClusterPolicyConfigPortMapping(components.ForwardToClusterByPortMappingConfig{/* values here */})
```

### ForwardToClusterBySNIConfig

```go
forwardToVirtualClusterPolicyConfig := components.CreateForwardToVirtualClusterPolicyConfigSni(components.ForwardToClusterBySNIConfig{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch forwardToVirtualClusterPolicyConfig.Type {
	case components.ForwardToVirtualClusterPolicyConfigTypePortMapping:
		// forwardToVirtualClusterPolicyConfig.ForwardToClusterByPortMappingConfig is populated
	case components.ForwardToVirtualClusterPolicyConfigTypeSni:
		// forwardToVirtualClusterPolicyConfig.ForwardToClusterBySNIConfig is populated
}
```
