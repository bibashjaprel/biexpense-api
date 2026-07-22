package accounts

import "github.com/google/uuid"

type AccountBalance struct {
	ID                   uuid.UUID `json:"id"`
	Name                 string    `json:"name"`
	Type                 string    `json:"type"`
	OpeningBalance       float64   `json:"opening_balance"`
	NetTransactionAmount float64   `json:"net_transaction_amount"`
	CalculatedBalance    float64   `json:"calculated_balance"`
}
