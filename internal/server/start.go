package server

import (
	"net/http"

	"github.com/gabrielluizsf/qrpix/internal/pix"
	"github.com/i9si-sistemas/nine"
	i9 "github.com/i9si-sistemas/nine/pkg/server"
)

func Start(server nine.Server) {
	server.ServeFiles("/", "./static")
	qrPix := server.Group("/qrpix")
	qrPix.Get("/default", func(c *i9.Context) error {
		cp, err := pix.Default()
		if err != nil {
			return internalServerError(c)
		}
		qrCode, err := pix.New(cp)
		if err != nil {
			return internalServerError(c)
		}
		return c.Send(qrCode)
	})
	qrPix.Post("/generate", func(c *i9.Context) error {
		var body struct {
			ReceiverName  string  `json:"receiver_name"`
			Key           string  `json:"key"`
			City          string  `json:"city"`
			Amount        float64 `json:"amount,omitempty"`
			Description   string  `json:"description,omitempty"`
			TransactionID string  `json:"transaction_id,omitempty"`
		}
		if err := c.BodyParser(&body); err != nil {
			return c.SendStatus(http.StatusBadRequest)
		}
		cp, err := pix.Generate(pix.Options{
			Name:          body.ReceiverName,
			Key:           body.Key,
			City:          body.City,
			Amount:        body.Amount,
			Description:   body.Description,
			TransactionID: body.TransactionID,
		})
		if err != nil {
			return internalServerError(c)
		}
		qrCode, err := pix.New(cp)
		if err != nil {
			return internalServerError(c)
		}
		return c.Send(qrCode)
	})
	server.Listen()
}

func internalServerError(c *i9.Context) error {
	return c.SendStatus(http.StatusInternalServerError)
}
