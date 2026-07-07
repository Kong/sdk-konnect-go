# UpdateAIGatewayModelProviderRequest


## Supported Types

### AIGatewayModelProviderAnthropic

```go
updateAIGatewayModelProviderRequest := components.CreateUpdateAIGatewayModelProviderRequestAnthropic(components.AIGatewayModelProviderAnthropic{/* values here */})
```

### AIGatewayModelProviderAzure

```go
updateAIGatewayModelProviderRequest := components.CreateUpdateAIGatewayModelProviderRequestAzure(components.AIGatewayModelProviderAzure{/* values here */})
```

### AIGatewayModelProviderBedrock

```go
updateAIGatewayModelProviderRequest := components.CreateUpdateAIGatewayModelProviderRequestBedrock(components.AIGatewayModelProviderBedrock{/* values here */})
```

### AIGatewayModelProviderCerebras

```go
updateAIGatewayModelProviderRequest := components.CreateUpdateAIGatewayModelProviderRequestCerebras(components.AIGatewayModelProviderCerebras{/* values here */})
```

### AIGatewayModelProviderCohere

```go
updateAIGatewayModelProviderRequest := components.CreateUpdateAIGatewayModelProviderRequestCohere(components.AIGatewayModelProviderCohere{/* values here */})
```

### AIGatewayModelProviderDashscope

```go
updateAIGatewayModelProviderRequest := components.CreateUpdateAIGatewayModelProviderRequestDashscope(components.AIGatewayModelProviderDashscope{/* values here */})
```

### AIGatewayModelProviderDatabricks

```go
updateAIGatewayModelProviderRequest := components.CreateUpdateAIGatewayModelProviderRequestDatabricks(components.AIGatewayModelProviderDatabricks{/* values here */})
```

### AIGatewayModelProviderDeepseek

```go
updateAIGatewayModelProviderRequest := components.CreateUpdateAIGatewayModelProviderRequestDeepseek(components.AIGatewayModelProviderDeepseek{/* values here */})
```

### AIGatewayModelProviderGemini

```go
updateAIGatewayModelProviderRequest := components.CreateUpdateAIGatewayModelProviderRequestGemini(components.AIGatewayModelProviderGemini{/* values here */})
```

### AIGatewayModelProviderHuggingface

```go
updateAIGatewayModelProviderRequest := components.CreateUpdateAIGatewayModelProviderRequestHuggingface(components.AIGatewayModelProviderHuggingface{/* values here */})
```

### AIGatewayModelProviderKimi

```go
updateAIGatewayModelProviderRequest := components.CreateUpdateAIGatewayModelProviderRequestKimi(components.AIGatewayModelProviderKimi{/* values here */})
```

### AIGatewayModelProviderLlama2

```go
updateAIGatewayModelProviderRequest := components.CreateUpdateAIGatewayModelProviderRequestLlama2(components.AIGatewayModelProviderLlama2{/* values here */})
```

### AIGatewayModelProviderMistral

```go
updateAIGatewayModelProviderRequest := components.CreateUpdateAIGatewayModelProviderRequestMistral(components.AIGatewayModelProviderMistral{/* values here */})
```

### AIGatewayModelProviderOllama

```go
updateAIGatewayModelProviderRequest := components.CreateUpdateAIGatewayModelProviderRequestOllama(components.AIGatewayModelProviderOllama{/* values here */})
```

### AIGatewayModelProviderOpenai

```go
updateAIGatewayModelProviderRequest := components.CreateUpdateAIGatewayModelProviderRequestOpenai(components.AIGatewayModelProviderOpenai{/* values here */})
```

### AIGatewayModelProviderVercel

```go
updateAIGatewayModelProviderRequest := components.CreateUpdateAIGatewayModelProviderRequestVercel(components.AIGatewayModelProviderVercel{/* values here */})
```

### AIGatewayModelProviderVllm

```go
updateAIGatewayModelProviderRequest := components.CreateUpdateAIGatewayModelProviderRequestVllm(components.AIGatewayModelProviderVllm{/* values here */})
```

### AIGatewayModelProviderXai

```go
updateAIGatewayModelProviderRequest := components.CreateUpdateAIGatewayModelProviderRequestXai(components.AIGatewayModelProviderXai{/* values here */})
```

### AIGatewayModelProviderVertex

```go
updateAIGatewayModelProviderRequest := components.CreateUpdateAIGatewayModelProviderRequestVertex(components.AIGatewayModelProviderVertex{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch updateAIGatewayModelProviderRequest.Type {
	case components.UpdateAIGatewayModelProviderRequestTypeAnthropic:
		// updateAIGatewayModelProviderRequest.AIGatewayModelProviderAnthropic is populated
	case components.UpdateAIGatewayModelProviderRequestTypeAzure:
		// updateAIGatewayModelProviderRequest.AIGatewayModelProviderAzure is populated
	case components.UpdateAIGatewayModelProviderRequestTypeBedrock:
		// updateAIGatewayModelProviderRequest.AIGatewayModelProviderBedrock is populated
	case components.UpdateAIGatewayModelProviderRequestTypeCerebras:
		// updateAIGatewayModelProviderRequest.AIGatewayModelProviderCerebras is populated
	case components.UpdateAIGatewayModelProviderRequestTypeCohere:
		// updateAIGatewayModelProviderRequest.AIGatewayModelProviderCohere is populated
	case components.UpdateAIGatewayModelProviderRequestTypeDashscope:
		// updateAIGatewayModelProviderRequest.AIGatewayModelProviderDashscope is populated
	case components.UpdateAIGatewayModelProviderRequestTypeDatabricks:
		// updateAIGatewayModelProviderRequest.AIGatewayModelProviderDatabricks is populated
	case components.UpdateAIGatewayModelProviderRequestTypeDeepseek:
		// updateAIGatewayModelProviderRequest.AIGatewayModelProviderDeepseek is populated
	case components.UpdateAIGatewayModelProviderRequestTypeGemini:
		// updateAIGatewayModelProviderRequest.AIGatewayModelProviderGemini is populated
	case components.UpdateAIGatewayModelProviderRequestTypeHuggingface:
		// updateAIGatewayModelProviderRequest.AIGatewayModelProviderHuggingface is populated
	case components.UpdateAIGatewayModelProviderRequestTypeKimi:
		// updateAIGatewayModelProviderRequest.AIGatewayModelProviderKimi is populated
	case components.UpdateAIGatewayModelProviderRequestTypeLlama2:
		// updateAIGatewayModelProviderRequest.AIGatewayModelProviderLlama2 is populated
	case components.UpdateAIGatewayModelProviderRequestTypeMistral:
		// updateAIGatewayModelProviderRequest.AIGatewayModelProviderMistral is populated
	case components.UpdateAIGatewayModelProviderRequestTypeOllama:
		// updateAIGatewayModelProviderRequest.AIGatewayModelProviderOllama is populated
	case components.UpdateAIGatewayModelProviderRequestTypeOpenai:
		// updateAIGatewayModelProviderRequest.AIGatewayModelProviderOpenai is populated
	case components.UpdateAIGatewayModelProviderRequestTypeVercel:
		// updateAIGatewayModelProviderRequest.AIGatewayModelProviderVercel is populated
	case components.UpdateAIGatewayModelProviderRequestTypeVllm:
		// updateAIGatewayModelProviderRequest.AIGatewayModelProviderVllm is populated
	case components.UpdateAIGatewayModelProviderRequestTypeXai:
		// updateAIGatewayModelProviderRequest.AIGatewayModelProviderXai is populated
	case components.UpdateAIGatewayModelProviderRequestTypeVertex:
		// updateAIGatewayModelProviderRequest.AIGatewayModelProviderVertex is populated
}
```
