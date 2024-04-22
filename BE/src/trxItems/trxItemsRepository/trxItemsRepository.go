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
            transaction_type = $1
		AND 
            isdeleted = false`

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
            transaction_items
        WHERE
            isdeleted = false`

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

func (t trxItemsRepository) UpdateTrxItemsByID(items trxItemsDto.TrxUpdateReq) error {
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
	  quantity = $1,
	  transaction_type = $2,
	  updated_at = NOW()
	WHERE
	  transaction_id = $3`

	_, err = tx.Exec(trxItemsQuery, items.Quantity, items.TrxType, items.TrxID)
	if err != nil {
		return err
	}

	itemsQuery := `UPDATE
	  items
	SET
	  amount = $1,
	  updated_at = NOW()
	WHERE
	  code = $2`

	_, err = tx.Exec(itemsQuery, items.Amount, items.ItemCode)
	if err != nil {
		return err
	}

	return nil
}

func (t trxItemsRepository) RetrieveTrxItemsByID(trxID string) (trxItemsDto.TrxItemsRes, error) {
	var items trxItemsDto.TrxItemsRes

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
            transaction_id = $1
		AND 
            isdeleted = false`

	row := t.db.QueryRow(trxItemsQuery, trxID)
	err := row.Scan(&items.TrxID, &items.ItemCode, &items.TrxType, &items.Quantity, &items.CreatedAt, &items.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return trxItemsDto.TrxItemsRes{}, errors.New("01")
		}
		return trxItemsDto.TrxItemsRes{}, err
	}

	return items, nil
}
