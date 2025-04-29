package main

import (
	"os"

	"github.com/gabrielluizsf/qrpix/internal/server"
	"github.com/i9si-sistemas/nine"
)

func main() {
	app := nine.NewServer(os.Getenv("PORT"))
	server.Start(app)
}
