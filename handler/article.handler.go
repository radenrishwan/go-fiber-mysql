package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go-fiber-mysql/model"
	"go-fiber-mysql/repository"
	"go-fiber-mysql/web"
	"net/http"
)

type ArticleHandler interface {
	Create(ctx *fiber.Ctx) error
	FindById(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	Bind(app *fiber.App)
}

type articleHandler struct {
	repository.ArticleRepository
}

func NewArticleHandler(articleRepository repository.ArticleRepository) ArticleHandler {
	return &articleHandler{ArticleRepository: articleRepository}
}

func (handler *articleHandler) Create(ctx *fiber.Ctx) error {
	var article model.Article

	err := ctx.BodyParser(&article)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(web.CommonResponse[any]{
			Message: "Bad Request",
			Data:    nil,
		})
	}

	article.Id = uuid.NewString()

	result := handler.ArticleRepository.Create(article)

	return ctx.Status(http.StatusCreated).JSON(web.CommonResponse[web.ArticleResponse]{
		Message: "201 Created",
		Data: web.ArticleResponse{
			Id:          result.Id,
			Title:       result.Title,
			Description: result.Description,
		},
	})
}

func (handler *articleHandler) FindById(ctx *fiber.Ctx) error {
	id := ctx.Query("id", "")

	result, err := handler.ArticleRepository.FindById(model.Article{
		Id: id,
	})

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(web.CommonResponse[any]{
			Message: "Article Not Found",
			Data:    nil,
		})
	}

	return ctx.Status(http.StatusCreated).JSON(web.CommonResponse[web.ArticleResponse]{
		Message: "201 Created",
		Data: web.ArticleResponse{
			Id:          result.Id,
			Title:       result.Title,
			Description: result.Description,
		},
	})
}

func (handler *articleHandler) Update(ctx *fiber.Ctx) error {
	var article model.Article

	err := ctx.BodyParser(&article)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(web.CommonResponse[any]{
			Message: "Bad Request",
			Data:    nil,
		})
	}

	_, err = handler.ArticleRepository.FindById(model.Article{
		Id: article.Id,
	})

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(web.CommonResponse[any]{
			Message: "Article Not Found",
			Data:    nil,
		})
	}

	handler.ArticleRepository.Update(article)

	return ctx.Status(http.StatusOK).JSON(web.CommonResponse[web.ArticleResponse]{
		Message: "200 OK",
		Data: web.ArticleResponse{
			Id:          article.Id,
			Title:       article.Title,
			Description: article.Description,
		},
	})
}

func (handler *articleHandler) Delete(ctx *fiber.Ctx) error {
	id := ctx.Query("id", "")

	_, err := handler.ArticleRepository.FindById(model.Article{
		Id: id,
	})

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(web.CommonResponse[any]{
			Message: "Article Not Found",
			Data:    nil,
		})
	}

	handler.ArticleRepository.Delete(model.Article{
		Id: id,
	})

	return ctx.Status(http.StatusOK).JSON(web.CommonResponse[string]{
		Message: "200 OK",
		Data:    "Delete Successfully",
	})
}

func (handler *articleHandler) Bind(app *fiber.App) {
	app.Post("api/article", handler.Create)
	app.Get("api/article", handler.FindById)
	app.Put("api/article", handler.Update)
	app.Delete("api/article", handler.Delete)
}
