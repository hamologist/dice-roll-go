package router

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/hamologist/dice-roll/pkg/evaluator"
	"github.com/hamologist/dice-roll/pkg/model"
	"net/http"
)

func RollRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", func(writer http.ResponseWriter, request *http.Request) {
		rollPayload := model.RollPayload{}
		err := json.NewDecoder(request.Body).Decode(&rollPayload)
		if err != nil {
			http.Error(writer, http.StatusText(500), 500)
			return
		}

		rollResponse, err := evaluator.EvaluateRoll(rollPayload)
		if err != nil {
			http.Error(writer, http.StatusText(500), 500)
			return
		}

		response, err := json.Marshal(rollResponse)
		if err != nil {
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.Write(response)
	})

	return router
}
