package itemsRepository

import (
	"BE-shop/models/dto/itemsDto"
	"BE-shop/src/items"
	"database/sql"
	"errors"
	"github.com/google/uuid"
)

type itemsRepository struct {
	db *sql.DB
}

func NewItemsRepository(db *sql.DB) items.ItemsRepository {
	return itemsRepository{db}
}

func (i itemsRepository) CreateItem(items itemsDto.Items) error {
	tx, err := i.db.Begin()
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

	itemsQuery := `INSERT INTO
	  items (
		items_id,
		code,
		name,
		amount,
		description,
		statusactive,
		created_at,
		updated_at
	  )
	VALUES
	  ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err = tx.Exec(itemsQuery, items.ItemsID, items.Code, items.Name, items.Amount, items.Description, items.StatusActive, items.CreatedAt, items.UpdatedAt)
	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"items_code_key\"" {
			return errors.New("01")
		}
		return err
	}

	tiq, err := uuid.NewRandom()
	transactionItemQuery := `INSERT INTO
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

	_, err = tx.Exec(transactionItemQuery, tiq, items.Code, items.TransactionType, items.Amount, items.CreatedAt, items.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (i itemsRepository) RetrieveAllItems() ([]itemsDto.ItemsResponse, error) {
	var items []itemsDto.ItemsResponse

	query := `
        SELECT items_id, code, name, amount, description, statusActive, created_at, updated_at
		FROM items
		WHERE isdeleted = FALSE
    `
	rows, err := i.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var item itemsDto.ItemsResponse
		err := rows.Scan(&item.ItemsID, &item.Code, &item.Name, &item.Amount, &item.Description, &item.StatusActive, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (i itemsRepository) RetrieveItemsByCode(code string) (itemsDto.Items, error) {
	var item itemsDto.Items

	query := `
		SELECT items_id, code, name, amount, description, statusActive, created_at, updated_at
		FROM items
		WHERE isdeleted = FALSE
		  AND code = $1
	`

	row := i.db.QueryRow(query, code)
	err := row.Scan(&item.ItemsID, &item.Code, &item.Name, &item.Amount, &item.Description, &item.StatusActive, &item.CreatedAt, &item.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return itemsDto.Items{}, errors.New("01")
		}
		return itemsDto.Items{}, err
	}

	return item, nil
}

func (i itemsRepository) UpdateItemsByCode(items itemsDto.ItemsUpdate) error {
	tx, err := i.db.Begin()
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
	  statusActive = $2,
	  updated_at = NOW()
	WHERE
	  code = $3`

	_, err = tx.Exec(itemsQuery, items.Amount, items.StatusActive, items.Code)
	if err != nil {
		return err
	}

	transactionItemsQuery := `UPDATE
	  transaction_items
	SET
	  quantity = $1,
	  transaction_type = $2,
	  updated_at = NOW()
	WHERE
	  items_code = $3`
	_, err = tx.Exec(transactionItemsQuery, items.Amount, items.TransactionType, items.Code)
	if err != nil {
		return err
	}

	return nil
}

func (i itemsRepository) DeleteItemsByCode(code string) error {
	tx, err := i.db.Begin()
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
	  isdeleted = TRUE,
	  updated_at = NOW()
	WHERE
	  code = $1`

	_, err = tx.Exec(itemsQuery, code)
	if err != nil {
		return err
	}

	return nil
}
