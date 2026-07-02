# TestSuggestionRule

Contains evaluation results of a suggestion rule configuration against an integration record.


## Fields

| Field                                                                                                       | Type                                                                                                        | Required                                                                                                    | Description                                                                                                 |
| ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- |
| `Evaluation`                                                                                                | [components.Evaluation](../../models/components/evaluation.md)                                              | :heavy_check_mark:                                                                                          | Evaluation result of suggesetion rule configuration against given integration record.                       |
| `Errors`                                                                                                    | [][components.SuggestionRuleErrorDetail](../../models/components/suggestionruleerrordetail.md)              | :heavy_check_mark:                                                                                          | List of errors that occured when evaluating suggestion rule configuration against given integration record. |