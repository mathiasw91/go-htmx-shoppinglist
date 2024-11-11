Stack:
- go
- sqlite
- templ
- htmx
- tailwind

Error handling & Validation:
- validation happens at handler level
- model functions should return errors
- error logging happens at handler level
- validation errors return bad request
- other errors return internal server error

DB Layer:
- methods on models are used for db interaction
- usually the struct is not accessed in these methods. they are just used for namespace reasons
- there should be refactoring with a better solution for the database layer, while still being domain specific

TODO:
- how to secure requests against unauthorized data fetches by changing "id" parameter in requests