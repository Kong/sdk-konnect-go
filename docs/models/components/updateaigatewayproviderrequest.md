# UpdateAIGatewayProviderRequest


## Supported Types

### AIGatewayProviderAnthropic

```go
updateAIGatewayProviderRequest := components.CreateUpdateAIGatewayProviderRequestAnthropic(components.AIGatewayProviderAnthropic{/* values here */})
```

### AIGatewayProviderAzure

```go
updateAIGatewayProviderRequest := components.CreateUpdateAIGatewayProviderRequestAzure(components.AIGatewayProviderAzure{/* values here */})
```

### AIGatewayProviderBedrock

```go
updateAIGatewayProviderRequest := components.CreateUpdateAIGatewayProviderRequestBedrock(components.AIGatewayProviderBedrock{/* values here */})
```

### AIGatewayProviderCerebras

```go
updateAIGatewayProviderRequest := components.CreateUpdateAIGatewayProviderRequestCerebras(components.AIGatewayProviderCerebras{/* values here */})
```

### AIGatewayProviderCohere

```go
updateAIGatewayProviderRequest := components.CreateUpdateAIGatewayProviderRequestCohere(components.AIGatewayProviderCohere{/* values here */})
```

### AIGatewayProviderDashscope

```go
updateAIGatewayProviderRequest := components.CreateUpdateAIGatewayProviderRequestDashscope(components.AIGatewayProviderDashscope{/* values here */})
```

### AIGatewayProviderDatabricks

```go
updateAIGatewayProviderRequest := components.CreateUpdateAIGatewayProviderRequestDatabricks(components.AIGatewayProviderDatabricks{/* values here */})
```

### AIGatewayProviderDeepseek

```go
updateAIGatewayProviderRequest := components.CreateUpdateAIGatewayProviderRequestDeepseek(components.AIGatewayProviderDeepseek{/* values here */})
```

### AIGatewayProviderGemini

```go
updateAIGatewayProviderRequest := components.CreateUpdateAIGatewayProviderRequestGemini(components.AIGatewayProviderGemini{/* values here */})
```

### AIGatewayProviderHuggingface

```go
updateAIGatewayProviderRequest := components.CreateUpdateAIGatewayProviderRequestHuggingface(components.AIGatewayProviderHuggingface{/* values here */})
```

### AIGatewayProviderKimi

```go
updateAIGatewayProviderRequest := components.CreateUpdateAIGatewayProviderRequestKimi(components.AIGatewayProviderKimi{/* values here */})
```

### AIGatewayProviderLlama2

```go
updateAIGatewayProviderRequest := components.CreateUpdateAIGatewayProviderRequestLlama2(components.AIGatewayProviderLlama2{/* values here */})
```

### AIGatewayProviderMistral

```go
updateAIGatewayProviderRequest := components.CreateUpdateAIGatewayProviderRequestMistral(components.AIGatewayProviderMistral{/* values here */})
```

### AIGatewayProviderOllama

```go
updateAIGatewayProviderRequest := components.CreateUpdateAIGatewayProviderRequestOllama(components.AIGatewayProviderOllama{/* values here */})
```

### AIGatewayProviderOpenai

```go
updateAIGatewayProviderRequest := components.CreateUpdateAIGatewayProviderRequestOpenai(components.AIGatewayProviderOpenai{/* values here */})
```

### AIGatewayProviderVercel

```go
updateAIGatewayProviderRequest := components.CreateUpdateAIGatewayProviderRequestVercel(components.AIGatewayProviderVercel{/* values here */})
```

### AIGatewayProviderVllm

```go
updateAIGatewayProviderRequest := components.CreateUpdateAIGatewayProviderRequestVllm(components.AIGatewayProviderVllm{/* values here */})
```

### AIGatewayProviderXai

```go
updateAIGatewayProviderRequest := components.CreateUpdateAIGatewayProviderRequestXai(components.AIGatewayProviderXai{/* values here */})
```

### AIGatewayProviderVertex

```go
updateAIGatewayProviderRequest := components.CreateUpdateAIGatewayProviderRequestVertex(components.AIGatewayProviderVertex{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch updateAIGatewayProviderRequest.Type {
	case components.UpdateAIGatewayProviderRequestTypeAnthropic:
		// updateAIGatewayProviderRequest.AIGatewayProviderAnthropic is populated
	case components.UpdateAIGatewayProviderRequestTypeAzure:
		// updateAIGatewayProviderRequest.AIGatewayProviderAzure is populated
	case components.UpdateAIGatewayProviderRequestTypeBedrock:
		// updateAIGatewayProviderRequest.AIGatewayProviderBedrock is populated
	case components.UpdateAIGatewayProviderRequestTypeCerebras:
		// updateAIGatewayProviderRequest.AIGatewayProviderCerebras is populated
	case components.UpdateAIGatewayProviderRequestTypeCohere:
		// updateAIGatewayProviderRequest.AIGatewayProviderCohere is populated
	case components.UpdateAIGatewayProviderRequestTypeDashscope:
		// updateAIGatewayProviderRequest.AIGatewayProviderDashscope is populated
	case components.UpdateAIGatewayProviderRequestTypeDatabricks:
		// updateAIGatewayProviderRequest.AIGatewayProviderDatabricks is populated
	case components.UpdateAIGatewayProviderRequestTypeDeepseek:
		// updateAIGatewayProviderRequest.AIGatewayProviderDeepseek is populated
	case components.UpdateAIGatewayProviderRequestTypeGemini:
		// updateAIGatewayProviderRequest.AIGatewayProviderGemini is populated
	case components.UpdateAIGatewayProviderRequestTypeHuggingface:
		// updateAIGatewayProviderRequest.AIGatewayProviderHuggingface is populated
	case components.UpdateAIGatewayProviderRequestTypeKimi:
		// updateAIGatewayProviderRequest.AIGatewayProviderKimi is populated
	case components.UpdateAIGatewayProviderRequestTypeLlama2:
		// updateAIGatewayProviderRequest.AIGatewayProviderLlama2 is populated
	case components.UpdateAIGatewayProviderRequestTypeMistral:
		// updateAIGatewayProviderRequest.AIGatewayProviderMistral is populated
	case components.UpdateAIGatewayProviderRequestTypeOllama:
		// updateAIGatewayProviderRequest.AIGatewayProviderOllama is populated
	case components.UpdateAIGatewayProviderRequestTypeOpenai:
		// updateAIGatewayProviderRequest.AIGatewayProviderOpenai is populated
	case components.UpdateAIGatewayProviderRequestTypeVercel:
		// updateAIGatewayProviderRequest.AIGatewayProviderVercel is populated
	case components.UpdateAIGatewayProviderRequestTypeVllm:
		// updateAIGatewayProviderRequest.AIGatewayProviderVllm is populated
	case components.UpdateAIGatewayProviderRequestTypeXai:
		// updateAIGatewayProviderRequest.AIGatewayProviderXai is populated
	case components.UpdateAIGatewayProviderRequestTypeVertex:
		// updateAIGatewayProviderRequest.AIGatewayProviderVertex is populated
}
```
