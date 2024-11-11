module shoppinglist/shoppinglist

go 1.22.1

replace shoppinglist/recipes => ../recipes

replace shoppinglist/layout => ../layout

replace shoppinglist/database => ../database

require (
	github.com/a-h/templ v0.2.771
	github.com/gorilla/mux v1.8.1
	shoppinglist/database v0.0.0-00010101000000-000000000000
	shoppinglist/layout v0.0.0-00010101000000-000000000000
	shoppinglist/recipes v0.0.0-00010101000000-000000000000
)

require (
	github.com/gosimple/slug v1.14.0 // indirect
	github.com/gosimple/unidecode v1.0.1 // indirect
	github.com/mattn/go-sqlite3 v1.14.22 // indirect
)
