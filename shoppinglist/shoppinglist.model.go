package shoppinglist

import (
	"context"
	"database/sql"
	"shoppinglist/database"
	"shoppinglist/recipes"
)

type Shoppinglist []shoppinglistItem

func (_ Shoppinglist) get() (Shoppinglist, error) {
	rows, err := database.DB.Query("SELECT * FROM shoppinglist_items;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	list := Shoppinglist{}
	for rows.Next() {
		var item shoppinglistItem
		err := rows.Scan(&item.id, &item.text, &item.checked, &item.created)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, nil
}

func (_ Shoppinglist) getItemById(id int) (shoppinglistItem, error) {
	row := database.DB.QueryRow("SELECT * FROM shoppinglist_items WHERE id = ?;", id)
	var item shoppinglistItem
	err := row.Scan(&item.id, &item.text, &item.checked, &item.created)
	if err != nil {
		return item, err
	}
	return item, nil
}

func (_ Shoppinglist) Add(item string) error {
	row := database.DB.QueryRow("SELECT MAX(id) as latestId FROM shoppinglist_items;")
	var latestId int32
	var latestIdOrNull sql.NullInt32
	err := row.Scan(&latestIdOrNull)
	if err == sql.ErrNoRows {
		latestId = 0
	} else if err != nil {
		return err
	} else {
		latestId = latestIdOrNull.Int32
	}
	_, insertErr := database.DB.Exec("INSERT INTO shoppinglist_items (id, text, checked) VALUES (?,?,?);", latestId+1, item, false)
	if insertErr != nil {
		return err
	}
	return nil
}

func (_ Shoppinglist) toggleItem(id int) error {
	_, err := database.DB.Exec("UPDATE shoppinglist_items SET checked = (1 - (checked & 1)) WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (_ Shoppinglist) renameItem(id int, text string) error {
	_, err := database.DB.Exec("UPDATE shoppinglist_items SET text = ? WHERE id = ?", text, id)
	if err != nil {
		return err
	}
	return nil
}

func (_ Shoppinglist) RemoveCheckedItems() error {
	_, err := database.DB.Exec("DELETE FROM shoppinglist_items WHERE checked = true")
	if err != nil {
		return err
	}
	return nil
}

func (_ Shoppinglist) RemoveAllItems() error {
	_, err := database.DB.Exec("DELETE FROM shoppinglist_items")
	if err != nil {
		return err
	}
	return nil
}

func (_ Shoppinglist) AddRecipe(ctx context.Context, recipe recipes.Recipe, anzahl int) error {
	tx, err := database.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	var latestId int64
	var latestIdOrNull sql.NullInt64
	if err = tx.QueryRowContext(ctx, "SELECT MAX(id) as latestId FROM shoppinglist_items;").Scan(&latestIdOrNull); err != nil {
		if err == sql.ErrNoRows {
			latestId = 0
		} else {
			return err
		}
	} else {
		latestId = latestIdOrNull.Int64
	}
	for _, item := range recipe.Items {
		result, insertErr := tx.ExecContext(ctx, "INSERT INTO shoppinglist_items (id, text, checked) VALUES (?,?,?);", latestId+1, item.ToStringMultiple(anzahl), false)
		if insertErr != nil {
			return err
		}
		latestId, err = result.LastInsertId()
		if err != nil {
			return err
		}
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}
