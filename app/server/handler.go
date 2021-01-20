package server

import (
	"encoding/json"
	"github.com/petuhovskiy/yc-healthcheck/app/logic"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handler struct {
	serv *logic.Service
}

func NewHandler(serv *logic.Service) *Handler {
	return &Handler{
		serv: serv,
	}
}

func (h *Handler) writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		logrus.WithError(err).Error("failed to marshal json")
	}
}

func (h *Handler) writeError(w http.ResponseWriter, err error) {
	resp := ErrorResponse{
		Error: err.Error(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	h.writeJSON(w, resp)
}

func (h *Handler) Healthcheck(w http.ResponseWriter, r *http.Request) {
	info, err := h.serv.Healthcheck()
	if err != nil {
		h.writeError(w, err)
		return
	}

	h.writeJSON(w, info)
}
