# AIGatewayModelProvider

**Pre-release Feature**
This feature is currently in beta and is subject to change.


## Supported Types

### AIGatewayModelProviderAIGatewayModelProviderAnthropic

```go
aiGatewayModelProvider := components.CreateAIGatewayModelProviderAnthropic(components.AIGatewayModelProviderAIGatewayModelProviderAnthropic{/* values here */})
```

### AIGatewayModelProviderAIGatewayModelProviderAzure

```go
aiGatewayModelProvider := components.CreateAIGatewayModelProviderAzure(components.AIGatewayModelProviderAIGatewayModelProviderAzure{/* values here */})
```

### AIGatewayModelProviderAIGatewayModelProviderBedrock

```go
aiGatewayModelProvider := components.CreateAIGatewayModelProviderBedrock(components.AIGatewayModelProviderAIGatewayModelProviderBedrock{/* values here */})
```

### AIGatewayModelProviderAIGatewayModelProviderCerebras

```go
aiGatewayModelProvider := components.CreateAIGatewayModelProviderCerebras(components.AIGatewayModelProviderAIGatewayModelProviderCerebras{/* values here */})
```

### AIGatewayModelProviderAIGatewayModelProviderCohere

```go
aiGatewayModelProvider := components.CreateAIGatewayModelProviderCohere(components.AIGatewayModelProviderAIGatewayModelProviderCohere{/* values here */})
```

### AIGatewayModelProviderAIGatewayModelProviderDashscope

```go
aiGatewayModelProvider := components.CreateAIGatewayModelProviderDashscope(components.AIGatewayModelProviderAIGatewayModelProviderDashscope{/* values here */})
```

### AIGatewayModelProviderAIGatewayModelProviderDatabricks

```go
aiGatewayModelProvider := components.CreateAIGatewayModelProviderDatabricks(components.AIGatewayModelProviderAIGatewayModelProviderDatabricks{/* values here */})
```

### AIGatewayModelProviderAIGatewayModelProviderDeepseek

```go
aiGatewayModelProvider := components.CreateAIGatewayModelProviderDeepseek(components.AIGatewayModelProviderAIGatewayModelProviderDeepseek{/* values here */})
```

### AIGatewayModelProviderAIGatewayModelProviderGemini

```go
aiGatewayModelProvider := components.CreateAIGatewayModelProviderGemini(components.AIGatewayModelProviderAIGatewayModelProviderGemini{/* values here */})
```

### AIGatewayModelProviderAIGatewayModelProviderHuggingface

```go
aiGatewayModelProvider := components.CreateAIGatewayModelProviderHuggingface(components.AIGatewayModelProviderAIGatewayModelProviderHuggingface{/* values here */})
```

### AIGatewayModelProviderAIGatewayModelProviderKimi

```go
aiGatewayModelProvider := components.CreateAIGatewayModelProviderKimi(components.AIGatewayModelProviderAIGatewayModelProviderKimi{/* values here */})
```

### AIGatewayModelProviderAIGatewayModelProviderLlama2

```go
aiGatewayModelProvider := components.CreateAIGatewayModelProviderLlama2(components.AIGatewayModelProviderAIGatewayModelProviderLlama2{/* values here */})
```

### AIGatewayModelProviderAIGatewayModelProviderMistral

```go
aiGatewayModelProvider := components.CreateAIGatewayModelProviderMistral(components.AIGatewayModelProviderAIGatewayModelProviderMistral{/* values here */})
```

### AIGatewayModelProviderAIGatewayModelProviderOllama

```go
aiGatewayModelProvider := components.CreateAIGatewayModelProviderOllama(components.AIGatewayModelProviderAIGatewayModelProviderOllama{/* values here */})
```

### AIGatewayModelProviderAIGatewayModelProviderOpenai

```go
aiGatewayModelProvider := components.CreateAIGatewayModelProviderOpenai(components.AIGatewayModelProviderAIGatewayModelProviderOpenai{/* values here */})
```

### AIGatewayModelProviderAIGatewayModelProviderVercel

```go
aiGatewayModelProvider := components.CreateAIGatewayModelProviderVercel(components.AIGatewayModelProviderAIGatewayModelProviderVercel{/* values here */})
```

### AIGatewayModelProviderAIGatewayModelProviderVllm

```go
aiGatewayModelProvider := components.CreateAIGatewayModelProviderVllm(components.AIGatewayModelProviderAIGatewayModelProviderVllm{/* values here */})
```

### AIGatewayModelProviderAIGatewayModelProviderXai

```go
aiGatewayModelProvider := components.CreateAIGatewayModelProviderXai(components.AIGatewayModelProviderAIGatewayModelProviderXai{/* values here */})
```

### AIGatewayModelProviderAIGatewayModelProviderVertex

```go
aiGatewayModelProvider := components.CreateAIGatewayModelProviderVertex(components.AIGatewayModelProviderAIGatewayModelProviderVertex{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayModelProvider.Type {
	case components.AIGatewayModelProviderTypeAnthropic:
		// aiGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderAnthropic is populated
	case components.AIGatewayModelProviderTypeAzure:
		// aiGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderAzure is populated
	case components.AIGatewayModelProviderTypeBedrock:
		// aiGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderBedrock is populated
	case components.AIGatewayModelProviderTypeCerebras:
		// aiGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderCerebras is populated
	case components.AIGatewayModelProviderTypeCohere:
		// aiGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderCohere is populated
	case components.AIGatewayModelProviderTypeDashscope:
		// aiGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderDashscope is populated
	case components.AIGatewayModelProviderTypeDatabricks:
		// aiGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderDatabricks is populated
	case components.AIGatewayModelProviderTypeDeepseek:
		// aiGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderDeepseek is populated
	case components.AIGatewayModelProviderTypeGemini:
		// aiGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderGemini is populated
	case components.AIGatewayModelProviderTypeHuggingface:
		// aiGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderHuggingface is populated
	case components.AIGatewayModelProviderTypeKimi:
		// aiGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderKimi is populated
	case components.AIGatewayModelProviderTypeLlama2:
		// aiGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderLlama2 is populated
	case components.AIGatewayModelProviderTypeMistral:
		// aiGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderMistral is populated
	case components.AIGatewayModelProviderTypeOllama:
		// aiGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderOllama is populated
	case components.AIGatewayModelProviderTypeOpenai:
		// aiGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderOpenai is populated
	case components.AIGatewayModelProviderTypeVercel:
		// aiGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderVercel is populated
	case components.AIGatewayModelProviderTypeVllm:
		// aiGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderVllm is populated
	case components.AIGatewayModelProviderTypeXai:
		// aiGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderXai is populated
	case components.AIGatewayModelProviderTypeVertex:
		// aiGatewayModelProvider.AIGatewayModelProviderAIGatewayModelProviderVertex is populated
}
```
