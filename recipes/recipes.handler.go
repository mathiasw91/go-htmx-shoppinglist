package recipes

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func RecipesPage(w http.ResponseWriter, r *http.Request) {
	list, err := Recipelist{}.get()
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	recipesTempl(list).Render(r.Context(), w)
}

func RecipesCreateFragment(w http.ResponseWriter, r *http.Request) {
	recipesCreateFragmentTempl().Render(r.Context(), w)
}

func RecipesCreate(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if name == "" {
		log.Print("no FormValue \"name\" provided")
		http.Error(w, "no FormValue \"name\" provided", http.StatusBadRequest)
		return
	}
	err := Recipelist{}.addRecipe(name)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	list, err := Recipelist{}.get()
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	recipesPageLayoutTempl(list).Render(r.Context(), w)
}

func RecipesDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "invalid Request Var \"id\" provided", http.StatusBadRequest)
		return
	}
	err = Recipelist{}.deleteRecipe(id)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	list, err := Recipelist{}.get()
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	recipesPageLayoutTempl(list).Render(r.Context(), w)
}
