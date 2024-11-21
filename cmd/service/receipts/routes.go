package receipts

import (
	"fmt"
	"net/http"
	"takehome/cmd/types"
	"takehome/cmd/utils"
	"github.com/gorilla/mux"
)


type Handler struct {
	rserv types.ReceiptService
}

func NewHandler(receiptService types.ReceiptService) *Handler {
	return &Handler{rserv: receiptService}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/receipts/{id}/points", h.handleReceiptPoints).Methods("GET")
	router.HandleFunc("/receipts/process", h.handleReceiptProcess).Methods("POST")

}

func (h *Handler) handleReceiptPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, exists := vars["id"]

	if !exists {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing id parameter in path"))
		return
	}

	points, err := h.rserv.GetReceiptPointsById(id)
	if err != nil {
        utils.WriteError(w, http.StatusNotFound, err)
        return
	}
	res := types.RetrievePointsResponse{Points: fmt.Sprintf("%f",points)}
	utils.WriteJSON(w, http.StatusAccepted, res)
}

func (h *Handler) handleReceiptProcess(w http.ResponseWriter, r *http.Request) {
	// get payload
	var payload types.RegisterReceiptPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	// validate receipt payload
	if err := utils.ValidateJSON(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	combinedTime, floatTotal := utils.ConvertDateTimeTotal(payload.PurchaseDate, payload.PurchaseTime, payload.Total)
	id, err := h.rserv.CreateReceipt(types.Receipt{Retailer: payload.Retailer, PurchasedOn: combinedTime, Items: payload.Items, Total: floatTotal})
	if err != nil {
		utils.WriteError(w, http.StatusInsufficientStorage, err)
		return
	}
	res := types.RegisterReceiptResponse{ID: id}
	utils.WriteJSON(w, http.StatusCreated, res)
}
