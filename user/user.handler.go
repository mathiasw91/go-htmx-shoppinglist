package user

import (
	"log"
	"net/http"
	"net/smtp"

	"github.com/gorilla/sessions"
)

var cookieStore *sessions.CookieStore

const SESSION_STORE_NAME = "go-session"
const SESSION_COOKIE_NAME = "usrId"
const DEFAULT_LOGIN_FAIL_MSG = "Anmeldung fehlgeschlagen"

func init() {
	//TODO move key to secrets
	cookieStore = sessions.NewCookieStore([]byte("mysuperdupersecret"))
	cookieStore.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 15,
		HttpOnly: true,
	}
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	LoginTempl().Render(r.Context(), w)
}

func Login(w http.ResponseWriter, r *http.Request) {
	mail := r.FormValue("mail")
	if mail == "" {
		log.Print("no FormValue \"mail\" provided")
		http.Error(w, "Bitte Mail Adresse angeben", http.StatusUnauthorized)
		return
	}
	id, err := getIdByMail(mail)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, DEFAULT_LOGIN_FAIL_MSG, http.StatusUnauthorized)
		return
	}
	session, sessionGetErr := cookieStore.Get(r, SESSION_STORE_NAME)
	if sessionGetErr != nil {
		log.Print(sessionGetErr.Error())
		http.Error(w, DEFAULT_LOGIN_FAIL_MSG, http.StatusUnauthorized)
		return
	}
	session.Values[SESSION_COOKIE_NAME] = id
	sessErr := session.Save(r, w)
	if sessErr != nil {
		log.Print(sessErr.Error())
		http.Error(w, DEFAULT_LOGIN_FAIL_MSG, http.StatusUnauthorized)
		return
	}
	w.Header().Add("HX-Redirect", "/")
}

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	RegisterTempl().Render(r.Context(), w)
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	to := r.FormValue("mail")
	if to == "" {
		log.Print("no FormValue \"mail\" provided")
		return
	}
	//TODO: generate code from mail and salt
	//TODO write code and mail in database

	gmailPasswd := "TODO generate app password in gmail"
	auth := smtp.PlainAuth("", "TODO Email", gmailPasswd, "smtp.gmail.com")
	msg := []byte("From: Shoppinglist App <TODO Email>\r\n" +
		"To: recipient@example.net\r\n" +
		"Subject: Registrierung\r\n" +
		"\r\n" +
		"Du hast einen Account zur Nutzung der Shoppinglist App angelegt. Bestätige deine Email-Adresse unter <a href=\"http://localhost/anmelden/\">diesem Link</a>.\r\n")
	err := smtp.SendMail("smtp.gmail.com:25", auth, "TODO Email", []string{to}, msg)
	if err != nil {
		log.Print(err.Error())
	}
	RegisterTempl().Render(r.Context(), w)
}

func IsAuthenticated(r *http.Request) bool {
	session, err := cookieStore.Get(r, SESSION_STORE_NAME)
	if err != nil {
		log.Print(err.Error())
		return false
	}
	if session.Values[SESSION_COOKIE_NAME] == nil {
		log.Print("Session not found")
		return false
	}
	return true
}
