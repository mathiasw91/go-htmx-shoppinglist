package recipes

import (
	"database/sql"
	"shoppinglist/database"

	"github.com/gosimple/slug"
)

type Recipelist []Recipe

func (_ Recipelist) get() (Recipelist, error) {
	rows, err := database.DB.Query("SELECT * FROM recipes;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	newList := Recipelist{}
	for rows.Next() {
		var item Recipe
		err := rows.Scan(&item.Id, &item.Name, &item.slug, &item.notice, &item.created, &item.Persons)
		if err != nil {
			return nil, err
		}
		newList = append(newList, item)
	}
	return newList, nil
}

func (_ Recipelist) addRecipe(name string) error {
	row := database.DB.QueryRow("SELECT MAX(id) as latestId FROM recipes;")
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
	slug := slug.MakeLang(name, "de")
	_, err = database.DB.Exec("INSERT INTO recipes (id, name, slug, notice) VALUES (?,?,?,?);", latestId+1, name, slug, "")
	if err != nil {
		return err
	}
	return nil
}

func (_ Recipelist) deleteRecipe(id int) error {
	_, err := database.DB.Exec("DELETE FROM recipes WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (_ Recipelist) SearchByName(name string) (Recipelist, error) {
	rows, err := database.DB.Query("SELECT * FROM recipes WHERE name LIKE ?", "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	newList := Recipelist{}
	for rows.Next() {
		var item Recipe
		err := rows.Scan(&item.Id, &item.Name, &item.slug, &item.notice, &item.created, &item.Persons)
		if err != nil {
			return nil, err
		}
		newList = append(newList, item)
	}
	return newList, nil
}
