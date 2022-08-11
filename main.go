package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"go-fiber-mysql/database"
	"go-fiber-mysql/handler"
	"go-fiber-mysql/model"
	"go-fiber-mysql/repository"
	"go-fiber-mysql/web"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	app := fiber.New()

	app.Use(logger.New())  // digunakan untuk menampilkan log app
	app.Use(recover.New()) // digunakan untuk recover app jika break pada app

	// DB
	db := database.NewDB()

	err = db.AutoMigrate(&model.Article{}) // database migration
	if err != nil {
		panic(err)
	}

	// repository
	articleRepository := repository.NewArticleRepository(db)

	// handler
	handler.NewArticleHandler(articleRepository).Bind(app)

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(web.CommonResponse[string]{
			Data:    "200 OK",
			Message: "Hello, World!",
		})
	})

	err = app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}
