package itemsRepository

import (
	"BE-shop/models/dto/itemsDto"
	"BE-shop/src/items"
	"database/sql"
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

func (i itemsRepository) RetrieveAllItems(transactionType string) ([]itemsDto.Items, error) {
	var items []itemsDto.Items

	query := `
        SELECT i.items_id, i.code, i.name, i.amount, i.description, i.statusActive,
               ti.transaction_type, ti.created_at AS transaction_created_at
        FROM items i
        JOIN transaction_items ti ON i.code = ti.items_code
        WHERE i.isdeleted = FALSE
    `

	var rows *sql.Rows
	var err error

	if transactionType != "" {
		query += " AND ti.transaction_type = $1"
		query += " ORDER BY ti.transaction_type;"
		rows, err = i.db.Query(query, transactionType)
	} else {
		query += " ORDER BY ti.transaction_type;"
		rows, err = i.db.Query(query)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item itemsDto.Items
		err = rows.Scan(&item.ItemsID, &item.Code, &item.Name, &item.Amount, &item.Description, &item.StatusActive, &item.TransactionType, &item.CreatedAt)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (i itemsRepository) RetrieveItemsByCode(code string) (itemsDto.Items, error) {
	//TODO implement me
	panic("implement me")
}
