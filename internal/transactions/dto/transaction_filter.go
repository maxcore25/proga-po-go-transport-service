package dto

// TransactionFilter represents query parameters for filtering transactions.
// swagger:parameters listTransactions
type TransactionFilter struct {
	// Status is an optional filter by transaction status.
	// in: query
	Status *string `form:"status" json:"status"`
}
