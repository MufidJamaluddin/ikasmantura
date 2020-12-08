package main

import (
	"backend/controllers"
	_ "backend/docs"
	"backend/services/impl"
	"fmt"
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
	AboutController        controllers.AboutController
	AlbumController        controllers.AlbumController
	AlbumPhotoController   controllers.AlbumPhotoController
	ArticleController      controllers.ArticleController
	ArticleTopicController controllers.ArticleTopicController
	AuthController         controllers.AuthController
	DepartmentController   controllers.DepartmentController
	EventController        controllers.EventController
	UserController         controllers.UserController
)

func InitService(db *gorm.DB) {
	AboutController.Service = &impl.AboutServiceImpl{DB: db}
	AlbumController.Service = &impl.AlbumServiceImpl{DB: db}
	AlbumPhotoController.Service = &impl.AlbumPhotoServiceImpl{DB: db}
	ArticleController.Service = &impl.ArticleServiceImpl{DB: db}
	ArticleTopicController.Service = &impl.ArticleTopicServiceImpl{DB: db}
	AuthController.Service = &impl.AuthServiceImpl{DB: db}
	DepartmentController.Service = &impl.DepartmentServiceImpl{DB: db}
	EventController.Service = &impl.EventServiceImpl{DB: db}
	UserController.Service = &impl.UserServiceImpl{DB: db}

	AuthController.SecretKey = []byte(os.Getenv("SECRET_KEY"))
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

	assetUri := fmt.Sprintf("/%s", os.Getenv("ASSET_PATH"))

	secretHandler := jwtware.New(jwtware.Config{
		SigningKey:  []byte(os.Getenv("SECRET_KEY")),
		TokenLookup: "header:Authorization,cookie:web_ika_id",
	})

	secretOrPublicHandler := jwtware.New(jwtware.Config{
		SigningKey:  []byte(os.Getenv("SECRET_KEY")),
		TokenLookup: "header:Authorization,cookie:web_ika_id",
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Next()
		},
	})

	InitService(db)

	app.Use(recover.New())

	app.Use("/swagger", swagger.Handler)

	app.Static(assetUri, assetUri)

	api := app.Group("/api", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return c.Next()
	})

	api.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("CORS_URI"),
	}))

	apiV1 := api.Group("/v1")

	about := apiV1.Group("/about")
	about.Get("/:id", AboutController.GetAbout)
	about.Put("/:id", secretHandler, AboutController.UpdateAbout)

	album := apiV1.Group("/albums")
	album.Get("", AlbumController.SearchAlbum)
	album.Get("/:id", AlbumController.GetOneAlbum)
	album.Post("", secretHandler, AlbumController.SaveAlbum)
	album.Put("/:id", secretHandler, AlbumController.UpdateAlbum)
	album.Delete("/:id", secretHandler, AlbumController.DeleteAlbum)

	albumPhoto := apiV1.Group("/photos")
	albumPhoto.Get("", AlbumPhotoController.SearchAlbumPhoto)
	albumPhoto.Get("/:id", AlbumPhotoController.GetOneAlbumPhoto)
	albumPhoto.Post("", secretHandler, AlbumPhotoController.SaveAlbumPhoto)
	albumPhoto.Put("/:id", secretHandler, AlbumPhotoController.UpdateAlbumPhoto)
	albumPhoto.Delete("/:id", secretHandler, AlbumPhotoController.DeleteAlbumPhoto)

	article := apiV1.Group("/articles")
	article.Get("", secretOrPublicHandler, ArticleController.SearchArticle)
	article.Get("/:id", ArticleController.GetOneArticle)
	article.Post("", secretHandler, ArticleController.SaveArticle)
	article.Put("/:id", secretHandler, ArticleController.UpdateArticle)
	article.Delete("/:id", secretHandler, ArticleController.DeleteArticle)

	topics := apiV1.Group("/article_topics")
	topics.Get("", ArticleTopicController.SearchArticle)
	topics.Get("/:id", ArticleTopicController.GetOneArticle)
	topics.Post("", secretHandler, ArticleTopicController.SaveArticle)
	topics.Put("/:id", secretHandler, ArticleTopicController.UpdateArticle)
	topics.Delete("/:id", secretHandler, ArticleTopicController.DeleteArticle)

	auth := apiV1.Group("/auth")
	auth.Post("", AuthController.Login)
	auth.Delete("", AuthController.Logout)

	department := apiV1.Group("/departments")
	department.Get("", DepartmentController.SearchDepartment)
	department.Get("/:id", DepartmentController.GetOneDepartment)
	department.Post("", secretHandler, DepartmentController.SaveDepartment)
	department.Put("/:id", secretHandler, DepartmentController.UpdateDepartment)
	department.Delete("/:id", secretHandler, DepartmentController.DeleteDepartment)

	event := apiV1.Group("/events")
	event.Get("", EventController.SearchEvent)
	event.Get("/:id", EventController.GetOneEvent)
	event.Post("", secretHandler, EventController.SaveEvent)
	event.Put("/:id", secretHandler, EventController.UpdateEvent)
	event.Delete("/:id", secretHandler, EventController.DeleteEvent)

	eventReg := apiV1.Group("/eventregister")
	eventReg.Post("/:id", secretHandler, EventController.RegisterEvent)

	eventDownload := apiV1.Group("/eventsdownload")
	eventDownload.Post("/:id",
		secretHandler, timeout.New(EventController.DownloadEventTicket, 10*time.Second))

	user := apiV1.Group("/users")
	user.Get("", UserController.SearchUser)
	user.Get("/:id", UserController.GetOneUser)
	user.Post("", secretHandler, UserController.SaveUser)
	user.Put("/:id", secretHandler, UserController.UpdateUser)
	user.Delete("/:id", secretHandler, UserController.DeleteUser)
}
