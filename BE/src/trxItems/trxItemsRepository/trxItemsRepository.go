package trxItemsRepository

import (
	"BE-shop/models/dto/trxItemsDto"
	"BE-shop/src/trxItems"
	"database/sql"
	"errors"
)

type trxItemsRepository struct {
	db *sql.DB
}

func NewTrxItemsRepository(db *sql.DB) trxItems.TrxItemsRepository {
	return &trxItemsRepository{db}
}

func (t trxItemsRepository) CreateTrxItems(trx trxItemsDto.TrxItems) error {
	tx, err := t.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	itemsQuery := `UPDATE
	  items
	SET
	  amount = $1,
	  updated_at = NOW()
	WHERE
	  code = $2`

	_, err = tx.Exec(itemsQuery, trx.Amount, trx.ItemCode)
	if err != nil {
		return err
	}

	trxItemsQuery := `INSERT INTO
	  transaction_items (
		transaction_id,
		items_code,
		transaction_type,
		quantity,
		created_at,
		updated_at
	  )
	VALUES
	  ($1, $2, $3, $4, $5, $6)`

	_, err = tx.Exec(trxItemsQuery, trx.TrxID, trx.ItemCode, trx.TrxType, trx.Quantity, trx.CreatedAt, trx.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (t trxItemsRepository) RetrieveAllTrxItems(transactionType string) (trxItems []trxItemsDto.TrxItemsRes, err error) {
	var rows *sql.Rows

	if transactionType != "" {
		trxItemsQuery := `SELECT
            transaction_id,
            items_code,
            transaction_type,
            quantity,
            created_at,
            updated_at
        FROM
            transaction_items
        WHERE
            transaction_type = $1`

		rows, err = t.db.Query(trxItemsQuery, transactionType)
	} else {
		trxItemsQuery := `SELECT
            transaction_id,
            items_code,
            transaction_type,
            quantity,
            created_at,
            updated_at
        FROM
            transaction_items`

		rows, err = t.db.Query(trxItemsQuery)
	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var trxItem trxItemsDto.TrxItemsRes
		err = rows.Scan(&trxItem.TrxID, &trxItem.ItemCode, &trxItem.TrxType, &trxItem.Quantity, &trxItem.CreatedAt, &trxItem.UpdatedAt)
		if err != nil {
			return nil, err
		}

		trxItems = append(trxItems, trxItem)
	}

	return trxItems, nil
}

func (t trxItemsRepository) DeleteTrxItems(trxID string) error {
	tx, err := t.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	trxItemsQuery := `UPDATE
	  transaction_items
	SET
	  isdeleted = TRUE,
	  updated_at = NOW()
	WHERE
	  transaction_id = $1`

	_, err = tx.Exec(trxItemsQuery, trxID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("01")
		}
		return err
	}

	return nil
}
