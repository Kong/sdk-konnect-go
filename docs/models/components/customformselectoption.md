# CustomFormSelectOption

A single option in a select field.


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                | Example                                                                    |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `Value`                                                                    | `string`                                                                   | :heavy_check_mark:                                                         | The value submitted when this option is chosen.                            | sales                                                                      |
| `Label`                                                                    | `string`                                                                   | :heavy_check_mark:                                                         | Display label for the option.                                              | Sales                                                                      |
| `Selected`                                                                 | `*bool`                                                                    | :heavy_minus_sign:                                                         | When `true`, this option is selected by default when the form is rendered. | false                                                                      |