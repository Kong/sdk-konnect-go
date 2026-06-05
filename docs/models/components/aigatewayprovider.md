# AIGatewayProvider


## Supported Types

### AIGatewayProviderAIGatewayProviderAnthropic

```go
aiGatewayProvider := components.CreateAIGatewayProviderAnthropic(components.AIGatewayProviderAIGatewayProviderAnthropic{/* values here */})
```

### AIGatewayProviderAIGatewayProviderAzure

```go
aiGatewayProvider := components.CreateAIGatewayProviderAzure(components.AIGatewayProviderAIGatewayProviderAzure{/* values here */})
```

### AIGatewayProviderAIGatewayProviderBedrock

```go
aiGatewayProvider := components.CreateAIGatewayProviderBedrock(components.AIGatewayProviderAIGatewayProviderBedrock{/* values here */})
```

### AIGatewayProviderAIGatewayProviderCerebras

```go
aiGatewayProvider := components.CreateAIGatewayProviderCerebras(components.AIGatewayProviderAIGatewayProviderCerebras{/* values here */})
```

### AIGatewayProviderAIGatewayProviderCohere

```go
aiGatewayProvider := components.CreateAIGatewayProviderCohere(components.AIGatewayProviderAIGatewayProviderCohere{/* values here */})
```

### AIGatewayProviderAIGatewayProviderDashscope

```go
aiGatewayProvider := components.CreateAIGatewayProviderDashscope(components.AIGatewayProviderAIGatewayProviderDashscope{/* values here */})
```

### AIGatewayProviderAIGatewayProviderDatabricks

```go
aiGatewayProvider := components.CreateAIGatewayProviderDatabricks(components.AIGatewayProviderAIGatewayProviderDatabricks{/* values here */})
```

### AIGatewayProviderAIGatewayProviderDeepseek

```go
aiGatewayProvider := components.CreateAIGatewayProviderDeepseek(components.AIGatewayProviderAIGatewayProviderDeepseek{/* values here */})
```

### AIGatewayProviderAIGatewayProviderGemini

```go
aiGatewayProvider := components.CreateAIGatewayProviderGemini(components.AIGatewayProviderAIGatewayProviderGemini{/* values here */})
```

### AIGatewayProviderAIGatewayProviderHuggingface

```go
aiGatewayProvider := components.CreateAIGatewayProviderHuggingface(components.AIGatewayProviderAIGatewayProviderHuggingface{/* values here */})
```

### AIGatewayProviderAIGatewayProviderKimi

```go
aiGatewayProvider := components.CreateAIGatewayProviderKimi(components.AIGatewayProviderAIGatewayProviderKimi{/* values here */})
```

### AIGatewayProviderAIGatewayProviderLlama2

```go
aiGatewayProvider := components.CreateAIGatewayProviderLlama2(components.AIGatewayProviderAIGatewayProviderLlama2{/* values here */})
```

### AIGatewayProviderAIGatewayProviderMistral

```go
aiGatewayProvider := components.CreateAIGatewayProviderMistral(components.AIGatewayProviderAIGatewayProviderMistral{/* values here */})
```

### AIGatewayProviderAIGatewayProviderOllama

```go
aiGatewayProvider := components.CreateAIGatewayProviderOllama(components.AIGatewayProviderAIGatewayProviderOllama{/* values here */})
```

### AIGatewayProviderAIGatewayProviderOpenai

```go
aiGatewayProvider := components.CreateAIGatewayProviderOpenai(components.AIGatewayProviderAIGatewayProviderOpenai{/* values here */})
```

### AIGatewayProviderAIGatewayProviderVercel

```go
aiGatewayProvider := components.CreateAIGatewayProviderVercel(components.AIGatewayProviderAIGatewayProviderVercel{/* values here */})
```

### AIGatewayProviderAIGatewayProviderVllm

```go
aiGatewayProvider := components.CreateAIGatewayProviderVllm(components.AIGatewayProviderAIGatewayProviderVllm{/* values here */})
```

### AIGatewayProviderAIGatewayProviderXai

```go
aiGatewayProvider := components.CreateAIGatewayProviderXai(components.AIGatewayProviderAIGatewayProviderXai{/* values here */})
```

### AIGatewayProviderAIGatewayProviderVertex

```go
aiGatewayProvider := components.CreateAIGatewayProviderVertex(components.AIGatewayProviderAIGatewayProviderVertex{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayProvider.Type {
	case components.AIGatewayProviderTypeAnthropic:
		// aiGatewayProvider.AIGatewayProviderAIGatewayProviderAnthropic is populated
	case components.AIGatewayProviderTypeAzure:
		// aiGatewayProvider.AIGatewayProviderAIGatewayProviderAzure is populated
	case components.AIGatewayProviderTypeBedrock:
		// aiGatewayProvider.AIGatewayProviderAIGatewayProviderBedrock is populated
	case components.AIGatewayProviderTypeCerebras:
		// aiGatewayProvider.AIGatewayProviderAIGatewayProviderCerebras is populated
	case components.AIGatewayProviderTypeCohere:
		// aiGatewayProvider.AIGatewayProviderAIGatewayProviderCohere is populated
	case components.AIGatewayProviderTypeDashscope:
		// aiGatewayProvider.AIGatewayProviderAIGatewayProviderDashscope is populated
	case components.AIGatewayProviderTypeDatabricks:
		// aiGatewayProvider.AIGatewayProviderAIGatewayProviderDatabricks is populated
	case components.AIGatewayProviderTypeDeepseek:
		// aiGatewayProvider.AIGatewayProviderAIGatewayProviderDeepseek is populated
	case components.AIGatewayProviderTypeGemini:
		// aiGatewayProvider.AIGatewayProviderAIGatewayProviderGemini is populated
	case components.AIGatewayProviderTypeHuggingface:
		// aiGatewayProvider.AIGatewayProviderAIGatewayProviderHuggingface is populated
	case components.AIGatewayProviderTypeKimi:
		// aiGatewayProvider.AIGatewayProviderAIGatewayProviderKimi is populated
	case components.AIGatewayProviderTypeLlama2:
		// aiGatewayProvider.AIGatewayProviderAIGatewayProviderLlama2 is populated
	case components.AIGatewayProviderTypeMistral:
		// aiGatewayProvider.AIGatewayProviderAIGatewayProviderMistral is populated
	case components.AIGatewayProviderTypeOllama:
		// aiGatewayProvider.AIGatewayProviderAIGatewayProviderOllama is populated
	case components.AIGatewayProviderTypeOpenai:
		// aiGatewayProvider.AIGatewayProviderAIGatewayProviderOpenai is populated
	case components.AIGatewayProviderTypeVercel:
		// aiGatewayProvider.AIGatewayProviderAIGatewayProviderVercel is populated
	case components.AIGatewayProviderTypeVllm:
		// aiGatewayProvider.AIGatewayProviderAIGatewayProviderVllm is populated
	case components.AIGatewayProviderTypeXai:
		// aiGatewayProvider.AIGatewayProviderAIGatewayProviderXai is populated
	case components.AIGatewayProviderTypeVertex:
		// aiGatewayProvider.AIGatewayProviderAIGatewayProviderVertex is populated
}
```
