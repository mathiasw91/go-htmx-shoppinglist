package recipes

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func RecipePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	recipe, err := Recipe{}.getBySlug(vars["slug"])
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	recipeTempl(recipe).Render(r.Context(), w)
}

func RecipeAddFragment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "invalid Request Var \"id\" provided", http.StatusBadRequest)
		return
	}
	recipeAddFragmentTempl(id).Render(r.Context(), w)
}

func RecipeAdd(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "invalid Request Var \"id\" provided", http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	amountString := r.FormValue("amount")
	amount, err := strconv.Atoi(amountString)
	if name == "" {
		log.Print("no FormValue \"name\" provided")
		http.Error(w, "no FormValue \"name\" provided", http.StatusBadRequest)
		return
	}
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "invalid FormValue \"name\" provided", http.StatusBadRequest)
		return
	}
	err = Recipe{}.addItem(id, name, amount)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	recipe, err := Recipe{}.GetById(id)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	recipePageLayout(recipe).Render(r.Context(), w)
}

func RecipeSetPersons(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "invalid Request Var \"id\" provided", http.StatusBadRequest)
		return
	}
	personsStr := r.FormValue("persons")
	persons, err := strconv.Atoi(personsStr)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "invalid FormValue \"persons\" provided", http.StatusBadRequest)
		return
	}
	err = Recipe{}.setPersons(id, persons)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	recipePersonsTempl(id, persons).Render(r.Context(), w)
}

func RecipeSetNotice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "invalid Request Var \"id\" provided", http.StatusBadRequest)
		return
	}
	err = Recipe{}.setNotice(id, r.FormValue("notiz"))
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	recipeNoticeTempl(id, r.FormValue("notiz")).Render(r.Context(), w)
}

func RecipeEditItemFragment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "invalid Request Var \"id\" provided", http.StatusBadRequest)
		return
	}
	item, err := recipeItem{}.GetById(id)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	RecipeEditItemFragmentTempl(&item).Render(r.Context(), w)
}

func RecipeEditItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "invalid Request Var \"id\" provided", http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "invalid FormValue \"name\" provided", http.StatusBadRequest)
		return
	}
	amountStr := r.FormValue("amount")
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "invalid FormValue \"amount\" provided", http.StatusBadRequest)
		return
	}
	err = recipeItem{}.edit(id, name, amount)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	item, err := recipeItem{}.GetById(id)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	recipeItemTempl(id, item.ToString()).Render(r.Context(), w)
}
