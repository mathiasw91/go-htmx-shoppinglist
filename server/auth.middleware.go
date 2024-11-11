package main

import (
	"net/http"
	"shoppinglist/user"
)

const LOGIN_URL = "/anmelden"

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authenticated := user.IsAuthenticated(r)
		target := r.URL.String()
		if !authenticated && target != LOGIN_URL {
			if r.Header["Hx-Request"] != nil { //redirect htmx ajax requests client side via htmx
				w.Header().Add("Hx-Redirect", LOGIN_URL)
			} else { //redirect regular browser requests via redirect
				http.Redirect(w, r, LOGIN_URL, http.StatusFound)
			}
			return
		}
		if authenticated && target == LOGIN_URL {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}
