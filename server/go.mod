module shoppinglist/server

go 1.22.1

require (
	github.com/gorilla/mux v1.8.1
	shoppinglist/database v0.0.0-00010101000000-000000000000
	shoppinglist/recipes v0.0.0-00010101000000-000000000000
	shoppinglist/shoppinglist v0.0.0-00010101000000-000000000000
	shoppinglist/user v0.0.0-00010101000000-000000000000
)

require (
	github.com/a-h/templ v0.2.771 // indirect
	github.com/gorilla/securecookie v1.1.2 // indirect
	github.com/gorilla/sessions v1.3.0 // indirect
	github.com/gosimple/slug v1.14.0 // indirect
	github.com/gosimple/unidecode v1.0.1 // indirect
	github.com/mattn/go-sqlite3 v1.14.23 // indirect
	shoppinglist/layout v0.0.0-00010101000000-000000000000 // indirect
)

replace shoppinglist/recipes => ../recipes

replace shoppinglist/shoppinglist => ../shoppinglist

replace shoppinglist/layout => ../layout

replace shoppinglist/database => ../database

replace shoppinglist/user => ../user
