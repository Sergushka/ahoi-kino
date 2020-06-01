package utils

import (
	"encoding/json"
	"fmt"
	"github.com/sergushka/ahoi-kino/log"
	"github.com/sergushka/ahoi-kino/model"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		log.GetLogger().Println(err)
		JSON(w, statusCode, model.Response{
			Err:    err.Error(),
			Result: nil,
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}
