# CreateAIGatewayModelProviderRequest

**Pre-release Feature**
This feature is currently in beta and is subject to change.


## Supported Types

### AIGatewayModelProviderAnthropic

```go
createAIGatewayModelProviderRequest := components.CreateCreateAIGatewayModelProviderRequestAnthropic(components.AIGatewayModelProviderAnthropic{/* values here */})
```

### AIGatewayModelProviderAzure

```go
createAIGatewayModelProviderRequest := components.CreateCreateAIGatewayModelProviderRequestAzure(components.AIGatewayModelProviderAzure{/* values here */})
```

### AIGatewayModelProviderBedrock

```go
createAIGatewayModelProviderRequest := components.CreateCreateAIGatewayModelProviderRequestBedrock(components.AIGatewayModelProviderBedrock{/* values here */})
```

### AIGatewayModelProviderCerebras

```go
createAIGatewayModelProviderRequest := components.CreateCreateAIGatewayModelProviderRequestCerebras(components.AIGatewayModelProviderCerebras{/* values here */})
```

### AIGatewayModelProviderCohere

```go
createAIGatewayModelProviderRequest := components.CreateCreateAIGatewayModelProviderRequestCohere(components.AIGatewayModelProviderCohere{/* values here */})
```

### AIGatewayModelProviderDashscope

```go
createAIGatewayModelProviderRequest := components.CreateCreateAIGatewayModelProviderRequestDashscope(components.AIGatewayModelProviderDashscope{/* values here */})
```

### AIGatewayModelProviderDatabricks

```go
createAIGatewayModelProviderRequest := components.CreateCreateAIGatewayModelProviderRequestDatabricks(components.AIGatewayModelProviderDatabricks{/* values here */})
```

### AIGatewayModelProviderDeepseek

```go
createAIGatewayModelProviderRequest := components.CreateCreateAIGatewayModelProviderRequestDeepseek(components.AIGatewayModelProviderDeepseek{/* values here */})
```

### AIGatewayModelProviderGemini

```go
createAIGatewayModelProviderRequest := components.CreateCreateAIGatewayModelProviderRequestGemini(components.AIGatewayModelProviderGemini{/* values here */})
```

### AIGatewayModelProviderHuggingface

```go
createAIGatewayModelProviderRequest := components.CreateCreateAIGatewayModelProviderRequestHuggingface(components.AIGatewayModelProviderHuggingface{/* values here */})
```

### AIGatewayModelProviderKimi

```go
createAIGatewayModelProviderRequest := components.CreateCreateAIGatewayModelProviderRequestKimi(components.AIGatewayModelProviderKimi{/* values here */})
```

### AIGatewayModelProviderLlama2

```go
createAIGatewayModelProviderRequest := components.CreateCreateAIGatewayModelProviderRequestLlama2(components.AIGatewayModelProviderLlama2{/* values here */})
```

### AIGatewayModelProviderMistral

```go
createAIGatewayModelProviderRequest := components.CreateCreateAIGatewayModelProviderRequestMistral(components.AIGatewayModelProviderMistral{/* values here */})
```

### AIGatewayModelProviderOllama

```go
createAIGatewayModelProviderRequest := components.CreateCreateAIGatewayModelProviderRequestOllama(components.AIGatewayModelProviderOllama{/* values here */})
```

### AIGatewayModelProviderOpenai

```go
createAIGatewayModelProviderRequest := components.CreateCreateAIGatewayModelProviderRequestOpenai(components.AIGatewayModelProviderOpenai{/* values here */})
```

### AIGatewayModelProviderVercel

```go
createAIGatewayModelProviderRequest := components.CreateCreateAIGatewayModelProviderRequestVercel(components.AIGatewayModelProviderVercel{/* values here */})
```

### AIGatewayModelProviderVllm

```go
createAIGatewayModelProviderRequest := components.CreateCreateAIGatewayModelProviderRequestVllm(components.AIGatewayModelProviderVllm{/* values here */})
```

### AIGatewayModelProviderXai

```go
createAIGatewayModelProviderRequest := components.CreateCreateAIGatewayModelProviderRequestXai(components.AIGatewayModelProviderXai{/* values here */})
```

### AIGatewayModelProviderVertex

```go
createAIGatewayModelProviderRequest := components.CreateCreateAIGatewayModelProviderRequestVertex(components.AIGatewayModelProviderVertex{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createAIGatewayModelProviderRequest.Type {
	case components.CreateAIGatewayModelProviderRequestTypeAnthropic:
		// createAIGatewayModelProviderRequest.AIGatewayModelProviderAnthropic is populated
	case components.CreateAIGatewayModelProviderRequestTypeAzure:
		// createAIGatewayModelProviderRequest.AIGatewayModelProviderAzure is populated
	case components.CreateAIGatewayModelProviderRequestTypeBedrock:
		// createAIGatewayModelProviderRequest.AIGatewayModelProviderBedrock is populated
	case components.CreateAIGatewayModelProviderRequestTypeCerebras:
		// createAIGatewayModelProviderRequest.AIGatewayModelProviderCerebras is populated
	case components.CreateAIGatewayModelProviderRequestTypeCohere:
		// createAIGatewayModelProviderRequest.AIGatewayModelProviderCohere is populated
	case components.CreateAIGatewayModelProviderRequestTypeDashscope:
		// createAIGatewayModelProviderRequest.AIGatewayModelProviderDashscope is populated
	case components.CreateAIGatewayModelProviderRequestTypeDatabricks:
		// createAIGatewayModelProviderRequest.AIGatewayModelProviderDatabricks is populated
	case components.CreateAIGatewayModelProviderRequestTypeDeepseek:
		// createAIGatewayModelProviderRequest.AIGatewayModelProviderDeepseek is populated
	case components.CreateAIGatewayModelProviderRequestTypeGemini:
		// createAIGatewayModelProviderRequest.AIGatewayModelProviderGemini is populated
	case components.CreateAIGatewayModelProviderRequestTypeHuggingface:
		// createAIGatewayModelProviderRequest.AIGatewayModelProviderHuggingface is populated
	case components.CreateAIGatewayModelProviderRequestTypeKimi:
		// createAIGatewayModelProviderRequest.AIGatewayModelProviderKimi is populated
	case components.CreateAIGatewayModelProviderRequestTypeLlama2:
		// createAIGatewayModelProviderRequest.AIGatewayModelProviderLlama2 is populated
	case components.CreateAIGatewayModelProviderRequestTypeMistral:
		// createAIGatewayModelProviderRequest.AIGatewayModelProviderMistral is populated
	case components.CreateAIGatewayModelProviderRequestTypeOllama:
		// createAIGatewayModelProviderRequest.AIGatewayModelProviderOllama is populated
	case components.CreateAIGatewayModelProviderRequestTypeOpenai:
		// createAIGatewayModelProviderRequest.AIGatewayModelProviderOpenai is populated
	case components.CreateAIGatewayModelProviderRequestTypeVercel:
		// createAIGatewayModelProviderRequest.AIGatewayModelProviderVercel is populated
	case components.CreateAIGatewayModelProviderRequestTypeVllm:
		// createAIGatewayModelProviderRequest.AIGatewayModelProviderVllm is populated
	case components.CreateAIGatewayModelProviderRequestTypeXai:
		// createAIGatewayModelProviderRequest.AIGatewayModelProviderXai is populated
	case components.CreateAIGatewayModelProviderRequestTypeVertex:
		// createAIGatewayModelProviderRequest.AIGatewayModelProviderVertex is populated
}
```
