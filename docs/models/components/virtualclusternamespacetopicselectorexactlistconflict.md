# VirtualClusterNamespaceTopicSelectorExactListConflict

How to inform the user about conflicts where multiple backend topics would map to the same virtual topic name.
* warn - log in the Event Gateway logs. Additionally, it sets knep_namespace_topic_conflict to 1.
* ignore - do not do anything. It does not cause knep_namespace_topic_conflict metric to be set to 1.



## Values

| Name                                                          | Value                                                         |
| ------------------------------------------------------------- | ------------------------------------------------------------- |
| `VirtualClusterNamespaceTopicSelectorExactListConflictWarn`   | warn                                                          |
| `VirtualClusterNamespaceTopicSelectorExactListConflictIgnore` | ignore                                                        |