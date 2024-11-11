package shoppinglist

import (
	"log"
	"net/http"
	"shoppinglist/recipes"
	"strconv"

	"github.com/gorilla/mux"
)

func ShoppinglistPage(w http.ResponseWriter, r *http.Request) {
	list, err := Shoppinglist{}.get()
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	shoppinglistTempl(list).Render(r.Context(), w)
}

func ShoppinglistToggleItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "invalid Request Var \"id\" provided", http.StatusBadRequest)
		return
	}
	err = Shoppinglist{}.toggleItem(id)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	item, err := Shoppinglist{}.getItemById(id)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	shoppinglistItemTempl(item.id, item.text, item.checked).Render(r.Context(), w)
}

func ShoppinglistRenameItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	text := r.FormValue("text")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if text == "" {
		log.Print("no Request Var \"text\" provided")
		http.Error(w, "no Request Var \"text\" provided", http.StatusBadRequest)
	}
	err = Shoppinglist{}.renameItem(id, text)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	shoppinglistItemTextTempl(id, text).Render(r.Context(), w)
}

func ShoppinglistRenameItemFragment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Print(err.Error())
	}
	item, err := Shoppinglist{}.getItemById(id)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	shoppinglistRenameItemFragmentTempl(id, item.text).Render(r.Context(), w)
}

func ShoppinglistAddItemFragment(w http.ResponseWriter, r *http.Request) {
	shoppinglistAddItemFragmentTempl().Render(r.Context(), w)
}

func ShoppinglistAddItem(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	if text == "" {
		log.Print("no FormValue \"text\" provided")
		http.Error(w, "no FormValue \"text\" provided", http.StatusBadRequest)
		return
	}
	err := Shoppinglist{}.Add(text)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	list, err := Shoppinglist{}.get()
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	pageLayoutTempl(list).Render(r.Context(), w)
}

func ShoppinglistAddRecipeFragment(w http.ResponseWriter, r *http.Request) {
	shoppinglistAddRecipeFragmentTempl().Render(r.Context(), w)
}

func ShoppinglistAddRecipe(w http.ResponseWriter, r *http.Request) {
	idString := r.FormValue("rezeptId")
	anzahlString := r.FormValue("anzahl")
	if idString == "" || anzahlString == "" {
		log.Print("no FormValue \"rezeptId\" or \"anzahl\" provided")
		http.Error(w, "no FormValue \"rezeptId\" or \"anzahl\" provided", http.StatusBadRequest)
	}
	id, _ := strconv.Atoi(idString)
	anzahl, _ := strconv.Atoi(anzahlString)
	recipe, err := recipes.Recipe{}.GetById(id)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = Shoppinglist{}.AddRecipe(r.Context(), recipe, anzahl)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	list, err := Shoppinglist{}.get()
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	pageLayoutTempl(list).Render(r.Context(), w)
}

func ShoppinglistAddRecipeSearch(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if name == "" {
		shoppinglistAddRecipeSearchFragmentTempl(recipes.Recipelist{}).Render(r.Context(), w)
		return
	}
	list, err := recipes.Recipelist{}.SearchByName(name)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	shoppinglistAddRecipeSearchFragmentTempl(list).Render(r.Context(), w)
}

func ShoppinglistDeleteItems(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var err error
	if vars["targets"] == "checked" {
		err = Shoppinglist{}.RemoveCheckedItems()
	} else {
		err = Shoppinglist{}.RemoveAllItems()
	}
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	list, err := Shoppinglist{}.get()
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	pageLayoutTempl(list).Render(r.Context(), w)
}
