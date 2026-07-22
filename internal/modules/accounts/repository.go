package accounts

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetAll(
	ctx context.Context,
) ([]AccountBalance, error) {

	query := `
		SELECT
			id,
			name,
			type,
			opening_balance,
			net_transaction_amount,
			calculated_balance
		FROM v_account_balances
		ORDER BY name;
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	accounts := make([]AccountBalance, 0)

	for rows.Next() {
		var account AccountBalance

		err := rows.Scan(
			&account.ID,
			&account.Name,
			&account.Type,
			&account.OpeningBalance,
			&account.NetTransactionAmount,
			&account.CalculatedBalance,
		)

		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return accounts, nil
}
