# AIGatewayEmbeddingsModelConfig

**Pre-release Feature**
This feature is currently in beta and is subject to change.

Configuration for an embeddings model.


## Supported Types

### AIGatewayAzureEmbeddingsModelConfig

```go
aiGatewayEmbeddingsModelConfig := components.CreateAIGatewayEmbeddingsModelConfigAzure(components.AIGatewayAzureEmbeddingsModelConfig{/* values here */})
```

### AIGatewayBedrockEmbeddingsModelConfig

```go
aiGatewayEmbeddingsModelConfig := components.CreateAIGatewayEmbeddingsModelConfigBedrock(components.AIGatewayBedrockEmbeddingsModelConfig{/* values here */})
```

### AIGatewayGeminiEmbeddingsModelConfig

```go
aiGatewayEmbeddingsModelConfig := components.CreateAIGatewayEmbeddingsModelConfigGemini(components.AIGatewayGeminiEmbeddingsModelConfig{/* values here */})
```

### AIGatewayHuggingfaceEmbeddingsModelConfig

```go
aiGatewayEmbeddingsModelConfig := components.CreateAIGatewayEmbeddingsModelConfigHuggingface(components.AIGatewayHuggingfaceEmbeddingsModelConfig{/* values here */})
```

### AIGatewayMistralEmbeddingsModelConfig

```go
aiGatewayEmbeddingsModelConfig := components.CreateAIGatewayEmbeddingsModelConfigMistral(components.AIGatewayMistralEmbeddingsModelConfig{/* values here */})
```

### AIGatewayOllamaEmbeddingsModelConfig

```go
aiGatewayEmbeddingsModelConfig := components.CreateAIGatewayEmbeddingsModelConfigOllama(components.AIGatewayOllamaEmbeddingsModelConfig{/* values here */})
```

### AIGatewayOpenaiEmbeddingsModelConfig

```go
aiGatewayEmbeddingsModelConfig := components.CreateAIGatewayEmbeddingsModelConfigOpenai(components.AIGatewayOpenaiEmbeddingsModelConfig{/* values here */})
```

### AIGatewayVertexEmbeddingsModelConfig

```go
aiGatewayEmbeddingsModelConfig := components.CreateAIGatewayEmbeddingsModelConfigVertex(components.AIGatewayVertexEmbeddingsModelConfig{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayEmbeddingsModelConfig.Type {
	case components.AIGatewayEmbeddingsModelConfigTypeAzure:
		// aiGatewayEmbeddingsModelConfig.AIGatewayAzureEmbeddingsModelConfig is populated
	case components.AIGatewayEmbeddingsModelConfigTypeBedrock:
		// aiGatewayEmbeddingsModelConfig.AIGatewayBedrockEmbeddingsModelConfig is populated
	case components.AIGatewayEmbeddingsModelConfigTypeGemini:
		// aiGatewayEmbeddingsModelConfig.AIGatewayGeminiEmbeddingsModelConfig is populated
	case components.AIGatewayEmbeddingsModelConfigTypeHuggingface:
		// aiGatewayEmbeddingsModelConfig.AIGatewayHuggingfaceEmbeddingsModelConfig is populated
	case components.AIGatewayEmbeddingsModelConfigTypeMistral:
		// aiGatewayEmbeddingsModelConfig.AIGatewayMistralEmbeddingsModelConfig is populated
	case components.AIGatewayEmbeddingsModelConfigTypeOllama:
		// aiGatewayEmbeddingsModelConfig.AIGatewayOllamaEmbeddingsModelConfig is populated
	case components.AIGatewayEmbeddingsModelConfigTypeOpenai:
		// aiGatewayEmbeddingsModelConfig.AIGatewayOpenaiEmbeddingsModelConfig is populated
	case components.AIGatewayEmbeddingsModelConfigTypeVertex:
		// aiGatewayEmbeddingsModelConfig.AIGatewayVertexEmbeddingsModelConfig is populated
}
```
