package main

import (
	_ "backend/docs"
	"backend/handlers"
	"backend/services/impl"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	jwtware "github.com/gofiber/jwt/v2"
	"gorm.io/gorm"
	"os"
	"time"
)

var (
	AboutHandler        handlers.AboutHandler
	AlbumHandler        handlers.AlbumHandler
	AlbumPhotoHandler   handlers.AlbumPhotoHandler
	ArticleHandler      handlers.ArticleHandler
	ArticleTopicHandler handlers.ArticleTopicHandler
	AuthHandler         handlers.AuthHandler
	DepartmentHandler   handlers.DepartmentHandler
	EventHandler        handlers.EventHandler
	UserHandler         handlers.UserHandler
)

func InitService(db *gorm.DB) {
	AboutHandler.Service = &impl.AboutServiceImpl{DB: db}
	AlbumHandler.Service = &impl.AlbumServiceImpl{DB: db}
	AlbumPhotoHandler.Service = &impl.AlbumPhotoServiceImpl{DB: db}
	ArticleHandler.Service = &impl.ArticleServiceImpl{DB: db}
	ArticleTopicHandler.Service = &impl.ArticleTopicServiceImpl{DB: db}
	AuthHandler.Service = &impl.AuthServiceImpl{DB: db}
	DepartmentHandler.Service = &impl.DepartmentServiceImpl{DB: db}
	EventHandler.Service = &impl.EventServiceImpl{DB: db}
	UserHandler.Service = &impl.UserServiceImpl{DB: db}

	AuthHandler.SecretKey = []byte(os.Getenv("SECRET_KEY"))
}

// @title IKA SMANTURA
// @version 1.0
// @description Web Ikatan Alumni SMA Negeri Situraja
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email mufid.jamaluddin@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /
func Route(app *fiber.App, db *gorm.DB) {

	secretHandler := jwtware.New(jwtware.Config{
		SigningKey:  []byte(os.Getenv("SECRET_KEY")),
		TokenLookup: "header:Authorization,cookie:web_ika_id",
	})

	/*
	secretOrPublicHandler  := jwtware.New(jwtware.Config{
		SigningKey:  []byte(os.Getenv("SECRET_KEY")),
		TokenLookup: "header:Authorization,cookie:web_ika_id",
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Next()
		},
	})
	 */

	InitService(db)

	app.Use(recover.New())

	app.Use("/swagger", swagger.Handler)

	api := app.Group("/api", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return c.Next()
	})

	api.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("CORS_URI"),
	}))

	apiV1 := api.Group("/v1")

	about := apiV1.Group("/about")
	about.Get("/:id", AboutHandler.GetAbout)
	about.Put("/:id", secretHandler, AboutHandler.UpdateAbout)

	album := apiV1.Group("/albums")
	album.Get("", AlbumHandler.SearchAlbum)
	album.Get("/:id", AlbumHandler.GetOneAlbum)
	album.Post("", secretHandler, AlbumHandler.SaveAlbum)
	album.Put("/:id", secretHandler, AlbumHandler.UpdateAlbum)
	album.Delete("/:id", secretHandler, AlbumHandler.DeleteAlbum)

	albumPhoto := apiV1.Group("/photos")
	albumPhoto.Get("", AlbumPhotoHandler.SearchAlbumPhoto)
	albumPhoto.Get("/:id", AlbumPhotoHandler.GetOneAlbumPhoto)
	albumPhoto.Post("", secretHandler, AlbumPhotoHandler.SaveAlbumPhoto)
	albumPhoto.Put("/:id", secretHandler, AlbumPhotoHandler.UpdateAlbumPhoto)
	albumPhoto.Delete("/:id", secretHandler, AlbumPhotoHandler.DeleteAlbumPhoto)

	article := apiV1.Group("/articles")
	article.Get("", ArticleHandler.SearchArticle)
	article.Get("/:id", ArticleHandler.GetOneArticle)
	article.Post("", secretHandler, ArticleHandler.SaveArticle)
	article.Put("/:id", secretHandler, ArticleHandler.UpdateArticle)
	article.Delete("/:id", secretHandler, ArticleHandler.DeleteArticle)

	topics := apiV1.Group("/article_topics")
	topics.Get("", ArticleTopicHandler.SearchArticle)
	topics.Get("/:id", ArticleTopicHandler.GetOneArticle)
	topics.Post("", secretHandler, ArticleTopicHandler.SaveArticle)
	topics.Put("/:id", secretHandler, ArticleTopicHandler.UpdateArticle)
	topics.Delete("/:id", secretHandler, ArticleTopicHandler.DeleteArticle)

	auth := apiV1.Group("/auth")
	auth.Post("", AuthHandler.Login)
	auth.Delete("", AuthHandler.Logout)

	department := apiV1.Group("/departments")
	department.Get("", DepartmentHandler.SearchDepartment)
	department.Get("/:id", DepartmentHandler.GetOneDepartment)
	department.Post("", secretHandler, DepartmentHandler.SaveDepartment)
	department.Put("/:id", secretHandler, DepartmentHandler.UpdateDepartment)
	department.Delete("/:id", secretHandler, DepartmentHandler.DeleteDepartment)

	event := apiV1.Group("/events")
	event.Get("", EventHandler.SearchEvent)
	event.Get("/:id", EventHandler.GetOneEvent)
	event.Post("", secretHandler, EventHandler.SaveEvent)
	event.Put("/:id", secretHandler, EventHandler.UpdateEvent)
	event.Delete("/:id", secretHandler, EventHandler.DeleteEvent)

	eventReg := apiV1.Group("/eventregister")
	eventReg.Post("/:id", secretHandler, EventHandler.RegisterEvent)

	eventDownload := apiV1.Group("/eventsdownload")
	eventDownload.Post("/:id",
		secretHandler, timeout.New(EventHandler.DownloadEventTicket, 10 * time.Second))

	user := apiV1.Group("/users")
	user.Get("", UserHandler.SearchUser)
	user.Get("/:id", UserHandler.GetOneUser)
	user.Post("", secretHandler, UserHandler.SaveUser)
	user.Put("/:id", secretHandler, UserHandler.UpdateUser)
	user.Delete("/:id", secretHandler, UserHandler.DeleteUser)
}
