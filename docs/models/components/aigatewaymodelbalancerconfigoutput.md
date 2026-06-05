# AIGatewayModelBalancerConfigOutput

Configuration for a model's load balancer when multiple target models are configured.


## Supported Types

### AIGatewayModelBalancerConsistentHashingConfig

```go
aiGatewayModelBalancerConfigOutput := components.CreateAIGatewayModelBalancerConfigOutputConsistentHashing(components.AIGatewayModelBalancerConsistentHashingConfig{/* values here */})
```

### AIGatewayModelBalancerLeastConnectionsConfig

```go
aiGatewayModelBalancerConfigOutput := components.CreateAIGatewayModelBalancerConfigOutputLeastConnections(components.AIGatewayModelBalancerLeastConnectionsConfig{/* values here */})
```

### AIGatewayModelBalancerLowestLatencyConfig

```go
aiGatewayModelBalancerConfigOutput := components.CreateAIGatewayModelBalancerConfigOutputLowestLatency(components.AIGatewayModelBalancerLowestLatencyConfig{/* values here */})
```

### AIGatewayModelBalancerLowestUsageConfig

```go
aiGatewayModelBalancerConfigOutput := components.CreateAIGatewayModelBalancerConfigOutputLowestUsage(components.AIGatewayModelBalancerLowestUsageConfig{/* values here */})
```

### AIGatewayModelBalancerPriorityConfig

```go
aiGatewayModelBalancerConfigOutput := components.CreateAIGatewayModelBalancerConfigOutputPriority(components.AIGatewayModelBalancerPriorityConfig{/* values here */})
```

### AIGatewayModelBalancerRoundRobinConfig

```go
aiGatewayModelBalancerConfigOutput := components.CreateAIGatewayModelBalancerConfigOutputRoundRobin(components.AIGatewayModelBalancerRoundRobinConfig{/* values here */})
```

### AIGatewayModelBalancerSemanticConfigOutput

```go
aiGatewayModelBalancerConfigOutput := components.CreateAIGatewayModelBalancerConfigOutputSemantic(components.AIGatewayModelBalancerSemanticConfigOutput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayModelBalancerConfigOutput.Type {
	case components.AIGatewayModelBalancerConfigOutputTypeConsistentHashing:
		// aiGatewayModelBalancerConfigOutput.AIGatewayModelBalancerConsistentHashingConfig is populated
	case components.AIGatewayModelBalancerConfigOutputTypeLeastConnections:
		// aiGatewayModelBalancerConfigOutput.AIGatewayModelBalancerLeastConnectionsConfig is populated
	case components.AIGatewayModelBalancerConfigOutputTypeLowestLatency:
		// aiGatewayModelBalancerConfigOutput.AIGatewayModelBalancerLowestLatencyConfig is populated
	case components.AIGatewayModelBalancerConfigOutputTypeLowestUsage:
		// aiGatewayModelBalancerConfigOutput.AIGatewayModelBalancerLowestUsageConfig is populated
	case components.AIGatewayModelBalancerConfigOutputTypePriority:
		// aiGatewayModelBalancerConfigOutput.AIGatewayModelBalancerPriorityConfig is populated
	case components.AIGatewayModelBalancerConfigOutputTypeRoundRobin:
		// aiGatewayModelBalancerConfigOutput.AIGatewayModelBalancerRoundRobinConfig is populated
	case components.AIGatewayModelBalancerConfigOutputTypeSemantic:
		// aiGatewayModelBalancerConfigOutput.AIGatewayModelBalancerSemanticConfigOutput is populated
}
```
