package recipes

import (
	"database/sql"
	"errors"
	"shoppinglist/database"
)

type Recipe struct {
	Id      int
	Name    string
	slug    string
	Items   []recipeItem
	notice  string
	created string
	Persons int
}

func (recipe Recipe) get() (Recipe, error) {
	newRecipe := Recipe{}
	var row *sql.Row
	if recipe.slug == "" && recipe.Id == 0 {
		return newRecipe, errors.New("Recipe.get() called without providing an \"id\" or \"slug\"")
	}
	if recipe.slug != "" {
		row = database.DB.QueryRow("SELECT * FROM recipes WHERE slug = ?;", recipe.slug)
	} else {
		row = database.DB.QueryRow("SELECT * FROM recipes WHERE id = ?;", recipe.Id)
	}
	row.Scan(&newRecipe.Id, &newRecipe.Name, &newRecipe.slug, &newRecipe.notice, &newRecipe.created, &newRecipe.Persons)
	rows, err := database.DB.Query("SELECT * FROM recipe_items WHERE recipe_id = ?;", newRecipe.Id)
	if err != nil {
		return newRecipe, err
	}
	defer rows.Close()
	for rows.Next() {
		var item recipeItem
		err := rows.Scan(&item.id, &item.recipeId, &item.name, &item.amount, &item.created)
		if err != nil {
			return newRecipe, err
		}
		newRecipe.Items = append(newRecipe.Items, item)
	}
	return newRecipe, nil
}

func (_ Recipe) addItem(recipeId int, name string, amount int) error {
	row := database.DB.QueryRow("SELECT MAX(id) as latestId FROM recipe_items;")
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
	_, err = database.DB.Exec("INSERT INTO recipe_items (id, recipe_id, name, amount) VALUES (?,?,?,?);", latestId+1, recipeId, name, amount)
	if err != nil {
		return err
	}
	return nil
}

func (_ Recipe) setPersons(recipeId int, persons int) error {
	_, err := database.DB.Exec("UPDATE recipes SET personen = ? WHERE id=?;", persons, recipeId)
	if err != nil {
		return err
	}
	return nil
}

func (_ Recipe) setNotice(recipeId int, text string) error {
	_, err := database.DB.Exec("UPDATE recipes SET notice = ? WHERE id=?;", text, recipeId)
	if err != nil {
		return err
	}
	return nil
}

func (_ Recipe) GetById(id int) (Recipe, error) {
	return Recipe{Id: id}.get()
}

func (_ Recipe) getBySlug(slug string) (Recipe, error) {
	return Recipe{slug: slug}.get()
}
