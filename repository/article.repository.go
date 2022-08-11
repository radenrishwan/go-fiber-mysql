package repository

import (
	"errors"
	"go-fiber-mysql/model"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	Create(article model.Article) model.Article
	FindById(model.Article) (model.Article, error)
	Update(model.Article)
	Delete(model.Article)
}

type articleRepository struct {
	*gorm.DB
}

func NewArticleRepository(DB *gorm.DB) ArticleRepository {
	return &articleRepository{DB: DB}
}

func (repository *articleRepository) Create(article model.Article) model.Article {
	repository.DB.Create(&article)

	return article
}

func (repository *articleRepository) FindById(article model.Article) (model.Article, error) {
	result := repository.DB.Where("id = ?", article.Id).First(&article)

	if result.RowsAffected < 1 {
		return article, errors.New("article not found")
	}

	return article, nil
}

func (repository *articleRepository) Update(article model.Article) {
	repository.DB.Where("id = ?", article.Id).Updates(&article)
}

func (repository *articleRepository) Delete(article model.Article) {
	repository.DB.Where("id = ?", article.Id).Delete(&article)
}
