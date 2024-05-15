package main

import (
	"context"
	"github.com/tclutin/classconnect-api/internal/app"
)

func main() {
	app.New().Run(context.Background())
}
