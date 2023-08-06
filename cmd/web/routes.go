package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handler(http.MethodGet, "/static/", http.StripPrefix("/static", fileServer))

	dynamicMiddleware := alice.New(app.sessionManager.LoadAndSave)

	router.Handler(http.MethodGet, "/", dynamicMiddleware.ThenFunc(app.home))

	// Snippets
	router.Handler(http.MethodGet, "/snippet/view/:id", dynamicMiddleware.ThenFunc(app.snippetView))
	router.Handler(http.MethodGet, "/snippet/create", dynamicMiddleware.ThenFunc(app.snippetCreate))
	router.Handler(http.MethodPost, "/snippet/create", dynamicMiddleware.ThenFunc(app.snippetCreatePost))

	// User
	router.Handler(http.MethodGet, "/user/signup", dynamicMiddleware.ThenFunc(app.userSignup))
	router.Handler(http.MethodPost, "/user/signup", dynamicMiddleware.ThenFunc(app.userSignupPost))
	router.Handler(http.MethodGet, "/user/login", dynamicMiddleware.ThenFunc(app.userLogin))
	router.Handler(http.MethodPost, "/user/login", dynamicMiddleware.ThenFunc(app.userLoginPost))
	router.Handler(http.MethodPost, "/user/logout", dynamicMiddleware.ThenFunc(app.userLogoutPost))

	standardMiddleware := alice.New(
		app.recoverPanic,
		app.logRequest,
		secureHeaders,
	)

	return standardMiddleware.Then(router)
}
