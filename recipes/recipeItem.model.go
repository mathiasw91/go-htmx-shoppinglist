package recipes

import (
	"fmt"
	"shoppinglist/database"
)

type recipeItem struct {
	id       int
	recipeId int
	name     string
	amount   int
	created  string
}

func (_ recipeItem) GetById(id int) (recipeItem, error) {
	row := database.DB.QueryRow("SELECT * FROM recipe_items WHERE id = ?;", id)
	item := recipeItem{}
	err := row.Scan(&item.id, &item.recipeId, &item.name, &item.amount, &item.created)
	if err != nil {
		return item, err
	}
	return item, nil
}

func (item *recipeItem) ToString() string {
	return fmt.Sprintf("%d x %s", item.amount, item.name)
}

func (item *recipeItem) ToStringMultiple(anzahl int) string {
	return fmt.Sprintf("%d x %s", anzahl*item.amount, item.name)
}

func (_ recipeItem) edit(id int, name string, amount int) error {
	_, err := database.DB.Exec("UPDATE recipe_items SET name = ?, amount = ? WHERE id=?;", name, amount, id)
	if err != nil {
		return err
	}
	return nil
}
