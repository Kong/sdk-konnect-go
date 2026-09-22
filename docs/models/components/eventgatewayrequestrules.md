# EventGatewayRequestRules

The rules to apply to one Kafka request type.


## Supported Types

### EventGatewayCreateTopicsRequestRules

```go
eventGatewayRequestRules := components.CreateEventGatewayRequestRulesCreateTopics(components.EventGatewayCreateTopicsRequestRules{/* values here */})
```

### EventGatewayCreatePartitionsRequestRules

```go
eventGatewayRequestRules := components.CreateEventGatewayRequestRulesCreatePartitions(components.EventGatewayCreatePartitionsRequestRules{/* values here */})
```

### EventGatewayAlterConfigsRequestRules

```go
eventGatewayRequestRules := components.CreateEventGatewayRequestRulesAlterConfigs(components.EventGatewayAlterConfigsRequestRules{/* values here */})
```

### EventGatewayIncrementalAlterConfigsRequestRules

```go
eventGatewayRequestRules := components.CreateEventGatewayRequestRulesIncrementalAlterConfigs(components.EventGatewayIncrementalAlterConfigsRequestRules{/* values here */})
```

### EventGatewayAlterClientQuotasRequestRules

```go
eventGatewayRequestRules := components.CreateEventGatewayRequestRulesAlterClientQuotas(components.EventGatewayAlterClientQuotasRequestRules{/* values here */})
```

### EventGatewayAlterUserScramCredentialsRequestRules

```go
eventGatewayRequestRules := components.CreateEventGatewayRequestRulesAlterUserScramCredentials(components.EventGatewayAlterUserScramCredentialsRequestRules{/* values here */})
```

### EventGatewayProduceRequestRules

```go
eventGatewayRequestRules := components.CreateEventGatewayRequestRulesProduce(components.EventGatewayProduceRequestRules{/* values here */})
```

### EventGatewayFetchRequestRules

```go
eventGatewayRequestRules := components.CreateEventGatewayRequestRulesFetch(components.EventGatewayFetchRequestRules{/* values here */})
```

### EventGatewayJoinGroupRequestRules

```go
eventGatewayRequestRules := components.CreateEventGatewayRequestRulesJoinGroup(components.EventGatewayJoinGroupRequestRules{/* values here */})
```

### EventGatewayConsumerGroupHeartbeatRequestRules

```go
eventGatewayRequestRules := components.CreateEventGatewayRequestRulesConsumerGroupHeartbeat(components.EventGatewayConsumerGroupHeartbeatRequestRules{/* values here */})
```

### EventGatewayOffsetCommitRequestRules

```go
eventGatewayRequestRules := components.CreateEventGatewayRequestRulesOffsetCommit(components.EventGatewayOffsetCommitRequestRules{/* values here */})
```

### EventGatewayOffsetFetchRequestRules

```go
eventGatewayRequestRules := components.CreateEventGatewayRequestRulesOffsetFetch(components.EventGatewayOffsetFetchRequestRules{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventGatewayRequestRules.Type {
	case components.EventGatewayRequestRulesTypeCreateTopics:
		// eventGatewayRequestRules.EventGatewayCreateTopicsRequestRules is populated
	case components.EventGatewayRequestRulesTypeCreatePartitions:
		// eventGatewayRequestRules.EventGatewayCreatePartitionsRequestRules is populated
	case components.EventGatewayRequestRulesTypeAlterConfigs:
		// eventGatewayRequestRules.EventGatewayAlterConfigsRequestRules is populated
	case components.EventGatewayRequestRulesTypeIncrementalAlterConfigs:
		// eventGatewayRequestRules.EventGatewayIncrementalAlterConfigsRequestRules is populated
	case components.EventGatewayRequestRulesTypeAlterClientQuotas:
		// eventGatewayRequestRules.EventGatewayAlterClientQuotasRequestRules is populated
	case components.EventGatewayRequestRulesTypeAlterUserScramCredentials:
		// eventGatewayRequestRules.EventGatewayAlterUserScramCredentialsRequestRules is populated
	case components.EventGatewayRequestRulesTypeProduce:
		// eventGatewayRequestRules.EventGatewayProduceRequestRules is populated
	case components.EventGatewayRequestRulesTypeFetch:
		// eventGatewayRequestRules.EventGatewayFetchRequestRules is populated
	case components.EventGatewayRequestRulesTypeJoinGroup:
		// eventGatewayRequestRules.EventGatewayJoinGroupRequestRules is populated
	case components.EventGatewayRequestRulesTypeConsumerGroupHeartbeat:
		// eventGatewayRequestRules.EventGatewayConsumerGroupHeartbeatRequestRules is populated
	case components.EventGatewayRequestRulesTypeOffsetCommit:
		// eventGatewayRequestRules.EventGatewayOffsetCommitRequestRules is populated
	case components.EventGatewayRequestRulesTypeOffsetFetch:
		// eventGatewayRequestRules.EventGatewayOffsetFetchRequestRules is populated
}
```
