package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var (
	addr = flag.Int("addr", 8000, "TCP Address to listen to")
)

func main() {
	flag.Parse()
	app := fiber.New()
	// middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(compress.New())
	app.Use(cors.New())

	// start server
	log.Println(fmt.Sprintf("Listening on PORT %d", *addr))
	if err := app.Listen(fmt.Sprintf(":%d", *addr)); err != nil {
		log.Fatal(err.Error())
	}
}
