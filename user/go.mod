module shoppinglist/user

go 1.22.1

replace shoppinglist/layout => ../layout

require (
	github.com/a-h/templ v0.2.771
	github.com/gorilla/sessions v1.3.0
	shoppinglist/database v0.0.0-00010101000000-000000000000
	shoppinglist/layout v0.0.0-00010101000000-000000000000
)

require (
	github.com/gorilla/securecookie v1.1.2 // indirect
	github.com/mattn/go-sqlite3 v1.14.23 // indirect
)

replace shoppinglist/database => ../database
