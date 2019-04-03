package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/naqvijafar91/focus"
)

type AggregatorHandler struct {
	agservice focus.AggregatorService
}

func (agh *AggregatorHandler) GetAllDataForUser(w http.ResponseWriter, req *http.Request) {

	response, err := agh.agservice.GetAllData(req.Context().Value("userID").(string))
	if err != nil {
		fmt.Fprintf(w, "Failed to get details %s", err)
	}
	json.NewEncoder(w).Encode(response)
}

func (agh *AggregatorHandler) RegisterAggregatorRoutes(mux *http.ServeMux) {
	middlewares := chainMiddleware(withUserParsing)
	mux.HandleFunc("/", middlewares(func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			agh.GetAllDataForUser(w, req)
			break
		}
	}))
}

func NewAggregatorHandler(agservice focus.AggregatorService) *AggregatorHandler {
	return &AggregatorHandler{agservice}
}
