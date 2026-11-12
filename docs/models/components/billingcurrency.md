# BillingCurrency

Fiat or custom currency.


## Supported Types

### BillingCurrencyFiat

```go
billingCurrency := components.CreateBillingCurrencyFiat(components.BillingCurrencyFiat{/* values here */})
```

### BillingCurrencyCustom

```go
billingCurrency := components.CreateBillingCurrencyCustom(components.BillingCurrencyCustom{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch billingCurrency.Type {
	case components.BillingCurrencyUnionTypeFiat:
		// billingCurrency.BillingCurrencyFiat is populated
	case components.BillingCurrencyUnionTypeCustom:
		// billingCurrency.BillingCurrencyCustom is populated
}
```
