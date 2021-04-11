package main

import (
	"github.com/cwxyz007/x-cms/controller"
	"github.com/cwxyz007/x-cms/core"
	"github.com/cwxyz007/x-cms/database"
	"github.com/cwxyz007/x-cms/service"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AppContext struct {
	Logger   *logrus.Logger
	Services *service.AllServices
}

func main() {
	conf := core.GetConfig()

	db := database.New("test.db")

	app := iris.Default()

	services := setupService(db)

	appContext := &AppContext{
		Logger:   core.GetLogger(),
		Services: services,
	}

	if conf.Debug {
		app.Logger().SetLevel("debug")
	}

	routeApi := app.Party("/api")
	{
		mvc.Configure(routeApi.Party("/articles"), setupArticleMVC(appContext))
	}

	app.Listen(conf.Port)
}

func setupArticleMVC(ctx *AppContext) func(app *mvc.Application) {
	return func(app *mvc.Application) {
		app.Register(
			ctx.Services,
			ctx.Logger,
		)

		app.Handle(new(controller.ArticleController))
	}
}

func setupService(db *gorm.DB) *service.AllServices {
	tag := service.TagService{}
	tag.SetDB(db)

	category := service.CategoryService{}
	category.SetDB(db)

	post := service.PostService{}
	post.SetDB(db)

	postTag := service.PostTagService{}
	postTag.SetDB(db)

	postCategory := service.PostCategoryService{}
	postCategory.SetDB(db)

	article := service.ArticleService{
		TagService:      tag,
		CategoryService: category,

		PostTagService:      postTag,
		PostCategoryService: postCategory,
	}
	article.SetDB(db)

	return &service.AllServices{
		Tag:      tag,
		Category: category,

		Post:         post,
		PostTag:      postTag,
		PostCategory: postCategory,

		Article: article,
	}
}
