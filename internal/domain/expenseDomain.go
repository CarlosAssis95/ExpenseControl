package domain

type Expense struct {
	DescriptionExpense   string  `json:"ds_expense"`
	QuantityInstallments int     `json:"qt_installments"`
	ValueInstallments    float64 `json:"vl_instalments"`
	DatePaymentExpense   string  `json:"dt_payment"`
}
