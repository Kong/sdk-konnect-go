# AIGatewayTargetConfig

Configuration for a target model.


## Supported Types

### AIGatewayTargetAnthropicConfig

```go
aiGatewayTargetConfig := components.CreateAIGatewayTargetConfigAnthropic(components.AIGatewayTargetAnthropicConfig{/* values here */})
```

### AIGatewayTargetAzureConfig

```go
aiGatewayTargetConfig := components.CreateAIGatewayTargetConfigAzure(components.AIGatewayTargetAzureConfig{/* values here */})
```

### AIGatewayTargetBedrockConfig

```go
aiGatewayTargetConfig := components.CreateAIGatewayTargetConfigBedrock(components.AIGatewayTargetBedrockConfig{/* values here */})
```

### AIGatewayTargetCerebrasConfig

```go
aiGatewayTargetConfig := components.CreateAIGatewayTargetConfigCerebras(components.AIGatewayTargetCerebrasConfig{/* values here */})
```

### AIGatewayTargetCohereConfig

```go
aiGatewayTargetConfig := components.CreateAIGatewayTargetConfigCohere(components.AIGatewayTargetCohereConfig{/* values here */})
```

### AIGatewayTargetDashscopeConfig

```go
aiGatewayTargetConfig := components.CreateAIGatewayTargetConfigDashscope(components.AIGatewayTargetDashscopeConfig{/* values here */})
```

### AIGatewayTargetDatabricksConfig

```go
aiGatewayTargetConfig := components.CreateAIGatewayTargetConfigDatabricks(components.AIGatewayTargetDatabricksConfig{/* values here */})
```

### AIGatewayTargetDeepseekConfig

```go
aiGatewayTargetConfig := components.CreateAIGatewayTargetConfigDeepseek(components.AIGatewayTargetDeepseekConfig{/* values here */})
```

### AIGatewayTargetGeminiConfig

```go
aiGatewayTargetConfig := components.CreateAIGatewayTargetConfigGemini(components.AIGatewayTargetGeminiConfig{/* values here */})
```

### AIGatewayTargetHuggingfaceConfig

```go
aiGatewayTargetConfig := components.CreateAIGatewayTargetConfigHuggingface(components.AIGatewayTargetHuggingfaceConfig{/* values here */})
```

### AIGatewayTargetKimiConfig

```go
aiGatewayTargetConfig := components.CreateAIGatewayTargetConfigKimi(components.AIGatewayTargetKimiConfig{/* values here */})
```

### AIGatewayTargetLlama2Config

```go
aiGatewayTargetConfig := components.CreateAIGatewayTargetConfigLlama2(components.AIGatewayTargetLlama2Config{/* values here */})
```

### AIGatewayTargetMistralConfig

```go
aiGatewayTargetConfig := components.CreateAIGatewayTargetConfigMistral(components.AIGatewayTargetMistralConfig{/* values here */})
```

### AIGatewayTargetOllamaConfig

```go
aiGatewayTargetConfig := components.CreateAIGatewayTargetConfigOllama(components.AIGatewayTargetOllamaConfig{/* values here */})
```

### AIGatewayTargetOpenaiConfig

```go
aiGatewayTargetConfig := components.CreateAIGatewayTargetConfigOpenai(components.AIGatewayTargetOpenaiConfig{/* values here */})
```

### AIGatewayTargetVercelConfig

```go
aiGatewayTargetConfig := components.CreateAIGatewayTargetConfigVercel(components.AIGatewayTargetVercelConfig{/* values here */})
```

### AIGatewayTargetVertexConfig

```go
aiGatewayTargetConfig := components.CreateAIGatewayTargetConfigVertex(components.AIGatewayTargetVertexConfig{/* values here */})
```

### AIGatewayTargetVllmConfig

```go
aiGatewayTargetConfig := components.CreateAIGatewayTargetConfigVllm(components.AIGatewayTargetVllmConfig{/* values here */})
```

### AIGatewayTargetXaiConfig

```go
aiGatewayTargetConfig := components.CreateAIGatewayTargetConfigXai(components.AIGatewayTargetXaiConfig{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayTargetConfig.Type {
	case components.AIGatewayTargetConfigTypeAnthropic:
		// aiGatewayTargetConfig.AIGatewayTargetAnthropicConfig is populated
	case components.AIGatewayTargetConfigTypeAzure:
		// aiGatewayTargetConfig.AIGatewayTargetAzureConfig is populated
	case components.AIGatewayTargetConfigTypeBedrock:
		// aiGatewayTargetConfig.AIGatewayTargetBedrockConfig is populated
	case components.AIGatewayTargetConfigTypeCerebras:
		// aiGatewayTargetConfig.AIGatewayTargetCerebrasConfig is populated
	case components.AIGatewayTargetConfigTypeCohere:
		// aiGatewayTargetConfig.AIGatewayTargetCohereConfig is populated
	case components.AIGatewayTargetConfigTypeDashscope:
		// aiGatewayTargetConfig.AIGatewayTargetDashscopeConfig is populated
	case components.AIGatewayTargetConfigTypeDatabricks:
		// aiGatewayTargetConfig.AIGatewayTargetDatabricksConfig is populated
	case components.AIGatewayTargetConfigTypeDeepseek:
		// aiGatewayTargetConfig.AIGatewayTargetDeepseekConfig is populated
	case components.AIGatewayTargetConfigTypeGemini:
		// aiGatewayTargetConfig.AIGatewayTargetGeminiConfig is populated
	case components.AIGatewayTargetConfigTypeHuggingface:
		// aiGatewayTargetConfig.AIGatewayTargetHuggingfaceConfig is populated
	case components.AIGatewayTargetConfigTypeKimi:
		// aiGatewayTargetConfig.AIGatewayTargetKimiConfig is populated
	case components.AIGatewayTargetConfigTypeLlama2:
		// aiGatewayTargetConfig.AIGatewayTargetLlama2Config is populated
	case components.AIGatewayTargetConfigTypeMistral:
		// aiGatewayTargetConfig.AIGatewayTargetMistralConfig is populated
	case components.AIGatewayTargetConfigTypeOllama:
		// aiGatewayTargetConfig.AIGatewayTargetOllamaConfig is populated
	case components.AIGatewayTargetConfigTypeOpenai:
		// aiGatewayTargetConfig.AIGatewayTargetOpenaiConfig is populated
	case components.AIGatewayTargetConfigTypeVercel:
		// aiGatewayTargetConfig.AIGatewayTargetVercelConfig is populated
	case components.AIGatewayTargetConfigTypeVertex:
		// aiGatewayTargetConfig.AIGatewayTargetVertexConfig is populated
	case components.AIGatewayTargetConfigTypeVllm:
		// aiGatewayTargetConfig.AIGatewayTargetVllmConfig is populated
	case components.AIGatewayTargetConfigTypeXai:
		// aiGatewayTargetConfig.AIGatewayTargetXaiConfig is populated
}
```
