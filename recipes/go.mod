module shoppinglist/recipes

go 1.22.1

require (
	github.com/a-h/templ v0.2.771
	github.com/gosimple/slug v1.14.0
	shoppinglist/database v0.0.0-00010101000000-000000000000
	shoppinglist/layout v0.0.0-00010101000000-000000000000
)

require (
	github.com/gosimple/unidecode v1.0.1 // indirect
	github.com/mattn/go-sqlite3 v1.14.22 // indirect
)

replace shoppinglist/layout => ../layout

replace shoppinglist/database => ../database
