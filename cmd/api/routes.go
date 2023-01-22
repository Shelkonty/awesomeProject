package main

// я хочу оставлять комментарии чтобы не забыть где что и чтобы не забыть зачем мне это нужно надеюсь вы поймете меня
import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)
	// Add the route for the PUT /v1/movies/:id endpoint.
	router.HandlerFunc(http.MethodPut, "/v1/movies/:id", app.updateMovieHandler)
	return router
}
