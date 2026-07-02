# CreateAIGatewayProviderRequest


## Supported Types

### AIGatewayProviderAnthropic

```go
createAIGatewayProviderRequest := components.CreateCreateAIGatewayProviderRequestAnthropic(components.AIGatewayProviderAnthropic{/* values here */})
```

### AIGatewayProviderAzure

```go
createAIGatewayProviderRequest := components.CreateCreateAIGatewayProviderRequestAzure(components.AIGatewayProviderAzure{/* values here */})
```

### AIGatewayProviderBedrock

```go
createAIGatewayProviderRequest := components.CreateCreateAIGatewayProviderRequestBedrock(components.AIGatewayProviderBedrock{/* values here */})
```

### AIGatewayProviderCerebras

```go
createAIGatewayProviderRequest := components.CreateCreateAIGatewayProviderRequestCerebras(components.AIGatewayProviderCerebras{/* values here */})
```

### AIGatewayProviderCohere

```go
createAIGatewayProviderRequest := components.CreateCreateAIGatewayProviderRequestCohere(components.AIGatewayProviderCohere{/* values here */})
```

### AIGatewayProviderDashscope

```go
createAIGatewayProviderRequest := components.CreateCreateAIGatewayProviderRequestDashscope(components.AIGatewayProviderDashscope{/* values here */})
```

### AIGatewayProviderDatabricks

```go
createAIGatewayProviderRequest := components.CreateCreateAIGatewayProviderRequestDatabricks(components.AIGatewayProviderDatabricks{/* values here */})
```

### AIGatewayProviderDeepseek

```go
createAIGatewayProviderRequest := components.CreateCreateAIGatewayProviderRequestDeepseek(components.AIGatewayProviderDeepseek{/* values here */})
```

### AIGatewayProviderGemini

```go
createAIGatewayProviderRequest := components.CreateCreateAIGatewayProviderRequestGemini(components.AIGatewayProviderGemini{/* values here */})
```

### AIGatewayProviderHuggingface

```go
createAIGatewayProviderRequest := components.CreateCreateAIGatewayProviderRequestHuggingface(components.AIGatewayProviderHuggingface{/* values here */})
```

### AIGatewayProviderKimi

```go
createAIGatewayProviderRequest := components.CreateCreateAIGatewayProviderRequestKimi(components.AIGatewayProviderKimi{/* values here */})
```

### AIGatewayProviderLlama2

```go
createAIGatewayProviderRequest := components.CreateCreateAIGatewayProviderRequestLlama2(components.AIGatewayProviderLlama2{/* values here */})
```

### AIGatewayProviderMistral

```go
createAIGatewayProviderRequest := components.CreateCreateAIGatewayProviderRequestMistral(components.AIGatewayProviderMistral{/* values here */})
```

### AIGatewayProviderOllama

```go
createAIGatewayProviderRequest := components.CreateCreateAIGatewayProviderRequestOllama(components.AIGatewayProviderOllama{/* values here */})
```

### AIGatewayProviderOpenai

```go
createAIGatewayProviderRequest := components.CreateCreateAIGatewayProviderRequestOpenai(components.AIGatewayProviderOpenai{/* values here */})
```

### AIGatewayProviderVercel

```go
createAIGatewayProviderRequest := components.CreateCreateAIGatewayProviderRequestVercel(components.AIGatewayProviderVercel{/* values here */})
```

### AIGatewayProviderVllm

```go
createAIGatewayProviderRequest := components.CreateCreateAIGatewayProviderRequestVllm(components.AIGatewayProviderVllm{/* values here */})
```

### AIGatewayProviderXai

```go
createAIGatewayProviderRequest := components.CreateCreateAIGatewayProviderRequestXai(components.AIGatewayProviderXai{/* values here */})
```

### AIGatewayProviderVertex

```go
createAIGatewayProviderRequest := components.CreateCreateAIGatewayProviderRequestVertex(components.AIGatewayProviderVertex{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createAIGatewayProviderRequest.Type {
	case components.CreateAIGatewayProviderRequestTypeAnthropic:
		// createAIGatewayProviderRequest.AIGatewayProviderAnthropic is populated
	case components.CreateAIGatewayProviderRequestTypeAzure:
		// createAIGatewayProviderRequest.AIGatewayProviderAzure is populated
	case components.CreateAIGatewayProviderRequestTypeBedrock:
		// createAIGatewayProviderRequest.AIGatewayProviderBedrock is populated
	case components.CreateAIGatewayProviderRequestTypeCerebras:
		// createAIGatewayProviderRequest.AIGatewayProviderCerebras is populated
	case components.CreateAIGatewayProviderRequestTypeCohere:
		// createAIGatewayProviderRequest.AIGatewayProviderCohere is populated
	case components.CreateAIGatewayProviderRequestTypeDashscope:
		// createAIGatewayProviderRequest.AIGatewayProviderDashscope is populated
	case components.CreateAIGatewayProviderRequestTypeDatabricks:
		// createAIGatewayProviderRequest.AIGatewayProviderDatabricks is populated
	case components.CreateAIGatewayProviderRequestTypeDeepseek:
		// createAIGatewayProviderRequest.AIGatewayProviderDeepseek is populated
	case components.CreateAIGatewayProviderRequestTypeGemini:
		// createAIGatewayProviderRequest.AIGatewayProviderGemini is populated
	case components.CreateAIGatewayProviderRequestTypeHuggingface:
		// createAIGatewayProviderRequest.AIGatewayProviderHuggingface is populated
	case components.CreateAIGatewayProviderRequestTypeKimi:
		// createAIGatewayProviderRequest.AIGatewayProviderKimi is populated
	case components.CreateAIGatewayProviderRequestTypeLlama2:
		// createAIGatewayProviderRequest.AIGatewayProviderLlama2 is populated
	case components.CreateAIGatewayProviderRequestTypeMistral:
		// createAIGatewayProviderRequest.AIGatewayProviderMistral is populated
	case components.CreateAIGatewayProviderRequestTypeOllama:
		// createAIGatewayProviderRequest.AIGatewayProviderOllama is populated
	case components.CreateAIGatewayProviderRequestTypeOpenai:
		// createAIGatewayProviderRequest.AIGatewayProviderOpenai is populated
	case components.CreateAIGatewayProviderRequestTypeVercel:
		// createAIGatewayProviderRequest.AIGatewayProviderVercel is populated
	case components.CreateAIGatewayProviderRequestTypeVllm:
		// createAIGatewayProviderRequest.AIGatewayProviderVllm is populated
	case components.CreateAIGatewayProviderRequestTypeXai:
		// createAIGatewayProviderRequest.AIGatewayProviderXai is populated
	case components.CreateAIGatewayProviderRequestTypeVertex:
		// createAIGatewayProviderRequest.AIGatewayProviderVertex is populated
}
```
