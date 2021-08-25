package handler

import (
	"encoding/json"
	"net/http"

	"github.com/danisbagus/edagang-product/internal/core/port"
	"github.com/danisbagus/edagang-product/internal/dto"
)

type TransactionHandler struct {
	Service port.ITransactionService
}

func (rc TransactionHandler) NewTransaction(w http.ResponseWriter, r *http.Request) {
	var request dto.NewTransactionRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	data, err := rc.Service.NewTransaction(&request)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}
	writeResponse(w, http.StatusCreated, data)
}
