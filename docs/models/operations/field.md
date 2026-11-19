# Field

The field whose values to list:

- `subject`: subject references, such as `principal/{uuid}` or `oidc/{sub}`.
- `tags`: the tag keys, or with `key`, the values of that tag.
- `groups`, `teams`, `consumer_group_ids`: the individual group references, team references, or consumer group IDs from the events' lists.
- `source`, `provider`, `model`, `model_id`, `region`, `client`, `service_id`, `route_id`, `consumer_id`, `application_id`: the field's values.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/operations"
)

value := operations.FieldSubject

// Open enum: custom values can be created with a direct type cast
custom := operations.Field("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `FieldSubject`          | subject                 |
| `FieldTags`             | tags                    |
| `FieldSource`           | source                  |
| `FieldProvider`         | provider                |
| `FieldModel`            | model                   |
| `FieldModelID`          | model_id                |
| `FieldRegion`           | region                  |
| `FieldClient`           | client                  |
| `FieldGroups`           | groups                  |
| `FieldTeams`            | teams                   |
| `FieldServiceID`        | service_id              |
| `FieldRouteID`          | route_id                |
| `FieldConsumerID`       | consumer_id             |
| `FieldApplicationID`    | application_id          |
| `FieldConsumerGroupIds` | consumer_group_ids      |