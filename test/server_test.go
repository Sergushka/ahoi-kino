package test

import (
	"github.com/gorilla/mux"
	"github.com/sergushka/ahoi-kino/model"
	"github.com/sergushka/ahoi-kino/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/movies", Get).Methods("GET")
	return router
}

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := model.Response{
		Result: "OK",
	}

	utils.JSON(w, http.StatusOK, response)
}

func TestGraphQLEndpoint(t *testing.T) {
	request, _ := http.NewRequest("GET", "/movies", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}
