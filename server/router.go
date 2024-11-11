package main

import (
	"net/http"
	"shoppinglist/recipes"
	"shoppinglist/shoppinglist"
	"shoppinglist/user"

	"github.com/gorilla/mux"
)

func GenerateRouter() *mux.Router {
	router := mux.NewRouter()

	routes := router.PathPrefix("/").Subrouter()

	routes.HandleFunc("/", shoppinglist.ShoppinglistPage)
	routes.HandleFunc("/toggleItem/{id}", shoppinglist.ShoppinglistToggleItem)
	routes.HandleFunc("/renameItem/fragment/{id}", shoppinglist.ShoppinglistRenameItemFragment)
	routes.HandleFunc("/renameItem/{id}", shoppinglist.ShoppinglistRenameItem)
	routes.HandleFunc("/addItem/fragment", shoppinglist.ShoppinglistAddItemFragment)
	routes.HandleFunc("/addItem", shoppinglist.ShoppinglistAddItem)
	routes.HandleFunc("/addRecipe/fragment", shoppinglist.ShoppinglistAddRecipeFragment)
	routes.HandleFunc("/addRecipe", shoppinglist.ShoppinglistAddRecipe)
	routes.HandleFunc("/addRecipe/search", shoppinglist.ShoppinglistAddRecipeSearch)
	routes.HandleFunc("/delete/{targets}", shoppinglist.ShoppinglistDeleteItems)

	routes.HandleFunc("/rezepte", recipes.RecipesPage)
	routes.HandleFunc("/rezepte/anlegen/fragment", recipes.RecipesCreateFragment)
	routes.HandleFunc("/rezepte/anlegen", recipes.RecipesCreate)
	routes.HandleFunc("/rezepte/loeschen/{id}", recipes.RecipesDelete)

	routes.HandleFunc("/rezept/{slug}", recipes.RecipePage)
	routes.HandleFunc("/rezept/hinzufuegen/fragment/{id}", recipes.RecipeAddFragment)
	routes.HandleFunc("/rezept/hinzufuegen/{id}", recipes.RecipeAdd)
	routes.HandleFunc("/rezept/eintragEditieren/fragment/{id}", recipes.RecipeEditItemFragment)
	routes.HandleFunc("/rezept/eintragEditieren/{id}", recipes.RecipeEditItem)
	routes.HandleFunc("/rezept/{id}/personen", recipes.RecipeSetPersons).Methods("PUT")
	routes.HandleFunc("/rezept/{id}/notiz", recipes.RecipeSetNotice).Methods("PUT")

	routes.HandleFunc("/anmelden", user.LoginPage).Methods("GET")
	routes.HandleFunc("/anmelden", user.Login).Methods("POST")
	routes.HandleFunc("/registrieren", user.RegisterPage).Methods("GET")
	routes.HandleFunc("/registrieren", user.RegisterUser).Methods("PUT")

	fs := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))

	router.Use(loggingMiddleware)
	routes.Use(authMiddleware)

	return router
}
