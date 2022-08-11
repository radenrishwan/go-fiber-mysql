package repository

import (
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go-fiber-mysql/database"
	"go-fiber-mysql/model"
	"gorm.io/gorm"
	"testing"
)

type ArticleRepositorySuite struct {
	suite.Suite
	*gorm.DB
	ArticleRepository
}

func (suite *ArticleRepositorySuite) SetupTest() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	suite.DB = database.NewDB()

	err = suite.DB.AutoMigrate(&model.Article{})
	if err != nil {
		panic(err)
	}

	suite.ArticleRepository = NewArticleRepository(suite.DB)

	suite.DB.Exec("delete from articles")
}

func (suite *ArticleRepositorySuite) TestCreate() {
	result := suite.ArticleRepository.Create(model.Article{
		Id:          uuid.NewString(),
		Title:       "What is React JS ?",
		Description: "React. js is an open-source JavaScript library that is used for building user interfaces specifically for single-page applications",
	})

	article, err := suite.ArticleRepository.FindById(result)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), result, article)
}

func (suite *ArticleRepositorySuite) TestUpdate() {
	result := suite.ArticleRepository.Create(model.Article{
		Id:          uuid.NewString(),
		Title:       "What is React JS ?",
		Description: "React. js is an open-source JavaScript library that is used for building user interfaces specifically for single-page applications",
	})

	suite.ArticleRepository.Update(model.Article{
		Id:          result.Id,
		Title:       "What is Go Fiber ?",
		Description: "Fiber is a Go web framework built on top of Fasthttp, the fastest HTTP engine for Go. It's designed to ease things up for fast development with zero memory allocation and performance in mind.",
	})

	article, err := suite.ArticleRepository.FindById(result)

	assert.Nil(suite.T(), err)
	assert.NotEqual(suite.T(), result, article)
}

func (suite *ArticleRepositorySuite) TestDelete() {
	result := suite.ArticleRepository.Create(model.Article{
		Id:          uuid.NewString(),
		Title:       "What is React JS ?",
		Description: "React. js is an open-source JavaScript library that is used for building user interfaces specifically for single-page applications",
	})

	suite.ArticleRepository.Delete(result)

	_, err := suite.ArticleRepository.FindById(result)

	assert.NotNil(suite.T(), err)
}

func (suite *ArticleRepositorySuite) TearDownSuite() {
	suite.DB.Exec("delete from urls")
}

func TestArticleRepsitory(t *testing.T) {
	suite.Run(t, new(ArticleRepositorySuite))
}
