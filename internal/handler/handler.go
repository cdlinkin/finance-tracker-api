package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cdlinkin/finance-tracker-api/internal/domain"
	"github.com/cdlinkin/finance-tracker-api/internal/service"
	"github.com/go-chi/chi/v5"
)

type TransactionHandler struct {
	service *service.TransactionService
}

func NewTransactionHandler(service *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		service: service,
	}
}

// POST
func (h *TransactionHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req domain.CreateTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "failed to POST transiction")
		return
	}
	tx, err := h.service.Create(req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to create transaction")
		return
	}
	writeJSON(w, http.StatusCreated, tx)
}

// GET
func (h *TransactionHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	transactions, err := h.service.GetAll()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to GET transaction")
		return
	}
	writeJSON(w, http.StatusOK, transactions)
}

// DELETE
func (h *TransactionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	reqID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(reqID)
	if err != nil {
		writeError(w, http.StatusBadRequest, "failed convert id")
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		writeError(w, http.StatusNotFound, "transaction not found")
		return
	}
	writeJSON(w, http.StatusNoContent, nil)
}

// GET
func (h *TransactionHandler) Summary(w http.ResponseWriter, r *http.Request) {
	income, expense, balance, err := h.service.Summary()
	if err != nil {
		writeError(w, http.StatusBadRequest, "failed Summary")
	}

	writeJSON(w, http.StatusOK, map[string]float64{
		"income":  income,
		"expense": expense,
		"balance": balance,
	})
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}
