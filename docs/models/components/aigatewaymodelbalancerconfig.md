# AIGatewayModelBalancerConfig

Configuration for a model's load balancer when multiple target models are configured.


## Supported Types

### AIGatewayModelBalancerConsistentHashingConfig

```go
aiGatewayModelBalancerConfig := components.CreateAIGatewayModelBalancerConfigConsistentHashing(components.AIGatewayModelBalancerConsistentHashingConfig{/* values here */})
```

### AIGatewayModelBalancerLeastConnectionsConfig

```go
aiGatewayModelBalancerConfig := components.CreateAIGatewayModelBalancerConfigLeastConnections(components.AIGatewayModelBalancerLeastConnectionsConfig{/* values here */})
```

### AIGatewayModelBalancerLowestLatencyConfig

```go
aiGatewayModelBalancerConfig := components.CreateAIGatewayModelBalancerConfigLowestLatency(components.AIGatewayModelBalancerLowestLatencyConfig{/* values here */})
```

### AIGatewayModelBalancerLowestUsageConfig

```go
aiGatewayModelBalancerConfig := components.CreateAIGatewayModelBalancerConfigLowestUsage(components.AIGatewayModelBalancerLowestUsageConfig{/* values here */})
```

### AIGatewayModelBalancerPriorityConfig

```go
aiGatewayModelBalancerConfig := components.CreateAIGatewayModelBalancerConfigPriority(components.AIGatewayModelBalancerPriorityConfig{/* values here */})
```

### AIGatewayModelBalancerRoundRobinConfig

```go
aiGatewayModelBalancerConfig := components.CreateAIGatewayModelBalancerConfigRoundRobin(components.AIGatewayModelBalancerRoundRobinConfig{/* values here */})
```

### AIGatewayModelBalancerSemanticConfig

```go
aiGatewayModelBalancerConfig := components.CreateAIGatewayModelBalancerConfigSemantic(components.AIGatewayModelBalancerSemanticConfig{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayModelBalancerConfig.Type {
	case components.AIGatewayModelBalancerConfigTypeConsistentHashing:
		// aiGatewayModelBalancerConfig.AIGatewayModelBalancerConsistentHashingConfig is populated
	case components.AIGatewayModelBalancerConfigTypeLeastConnections:
		// aiGatewayModelBalancerConfig.AIGatewayModelBalancerLeastConnectionsConfig is populated
	case components.AIGatewayModelBalancerConfigTypeLowestLatency:
		// aiGatewayModelBalancerConfig.AIGatewayModelBalancerLowestLatencyConfig is populated
	case components.AIGatewayModelBalancerConfigTypeLowestUsage:
		// aiGatewayModelBalancerConfig.AIGatewayModelBalancerLowestUsageConfig is populated
	case components.AIGatewayModelBalancerConfigTypePriority:
		// aiGatewayModelBalancerConfig.AIGatewayModelBalancerPriorityConfig is populated
	case components.AIGatewayModelBalancerConfigTypeRoundRobin:
		// aiGatewayModelBalancerConfig.AIGatewayModelBalancerRoundRobinConfig is populated
	case components.AIGatewayModelBalancerConfigTypeSemantic:
		// aiGatewayModelBalancerConfig.AIGatewayModelBalancerSemanticConfig is populated
}
```
