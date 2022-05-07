package main

import (
	"github.com/wartinich/web/pkg/handler"
)

func main() {
	r := handler.Handler()

	r.Run("127.0.0.1:8000")
}
