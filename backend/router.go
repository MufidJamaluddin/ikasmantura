package main

import (
	_ "backend/docs"
	aboutHandler "backend/handlers/about"
	albumHandler "backend/handlers/album"
	albumPhotoHandler "backend/handlers/albumphoto"
	articleHandler "backend/handlers/article"
	articleTopicHandler "backend/handlers/articletopic"
	authHandler "backend/handlers/auth"
	classroomHandler "backend/handlers/classroom"
	departmentHandler "backend/handlers/department"
	eventHandler "backend/handlers/event"
	tempUserHandler "backend/handlers/temp_user"
	userHandler "backend/handlers/user"
	"fmt"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	jwtWare "github.com/gofiber/jwt/v2"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
	"time"
)

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

	publicHandler := func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	}

	secretHandler := jwtWare.New(jwtWare.Config{
		SigningKey:  []byte(os.Getenv("SECRET_KEY")),
		TokenLookup: "header:Authorization,cookie:web_ika_id",
		SuccessHandler: func(ctx *fiber.Ctx) error {
			return authHandler.AuthenticationHandler(ctx, db)
		},
	})

	secretOrPublicHandler := jwtWare.New(jwtWare.Config{
		SigningKey:  []byte(os.Getenv("SECRET_KEY")),
		TokenLookup: "header:Authorization,cookie:web_ika_id",
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			log.Println(err.Error())
			ctx.Locals("user", nil)
			ctx.Locals("db", db)
			return ctx.Next()
		},
		SuccessHandler: func(ctx *fiber.Ctx) error {
			return authHandler.AuthenticationHandler(ctx, db)
		},
	})

	cacheDuration, _ := strconv.Atoi(os.Getenv("CACHE_DURATION"))
	cacheHandler := cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("refresh") == "true"
		},
		Expiration:   time.Duration(cacheDuration) * time.Minute,
		CacheControl: true,
	})

	app.Use(recover.New())

	app.Use("/swagger", swagger.Handler)

	app.Static(assetUri, "."+assetUri)

	app.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        10,
		Expiration: time.Second,
	}))

	api := app.Group("/api", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return c.Next()
	})

	api.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("CORS_URI"),
	}))

	apiV1 := api.Group("/v1")

	about := apiV1.Group("/about")
	about.Get("/:id", cacheHandler, publicHandler, aboutHandler.GetAbout)
	about.Put("/:id", secretHandler, aboutHandler.UpdateAbout)

	album := apiV1.Group("/albums")
	album.Get("", publicHandler, albumHandler.SearchAlbum)
	album.Get("/:id", publicHandler, albumHandler.GetOneAlbum)
	album.Post("", secretHandler, albumHandler.SaveAlbum)
	album.Put("/:id", secretHandler, albumHandler.UpdateAlbum)
	album.Delete("/:id", secretHandler, albumHandler.DeleteAlbum)

	albumPhoto := apiV1.Group("/photos")
	albumPhoto.Get("", publicHandler, albumPhotoHandler.SearchAlbumPhoto)
	albumPhoto.Get("/:id", publicHandler, albumPhotoHandler.GetOneAlbumPhoto)
	albumPhoto.Post("", secretHandler, albumPhotoHandler.SaveAlbumPhoto)
	albumPhoto.Put("/:id", secretHandler, albumPhotoHandler.UpdateAlbumPhoto)
	albumPhoto.Delete("/:id", secretHandler, albumPhotoHandler.DeleteAlbumPhoto)

	article := apiV1.Group("/articles")
	article.Get("", publicHandler, articleHandler.SearchArticle)
	article.Get("/:id", publicHandler, articleHandler.GetOneArticle)
	article.Post("", secretHandler, articleHandler.SaveArticle)
	article.Put("/:id", secretHandler, articleHandler.UpdateArticle)
	article.Delete("/:id", secretHandler, articleHandler.DeleteArticle)

	topics := apiV1.Group("/article_topics")
	topics.Get("", publicHandler, articleTopicHandler.SearchArticleTopic)
	topics.Get("/:id", publicHandler, articleTopicHandler.GetOneArticleTopic)
	topics.Post("", secretHandler, articleTopicHandler.SaveArticleTopic)
	topics.Put("/:id", secretHandler, articleTopicHandler.UpdateArticleTopic)
	topics.Delete("/:id", secretHandler, articleTopicHandler.DeleteArticleTopic)

	auth := apiV1.Group("/auth")
	auth.Post("", publicHandler, authHandler.Login)
	auth.Delete("", publicHandler, authHandler.Logout)

	department := apiV1.Group("/departments")
	department.Get("", publicHandler, departmentHandler.SearchDepartment)
	department.Get("/:id", publicHandler, departmentHandler.GetOneDepartment)
	department.Post("", secretHandler, departmentHandler.SaveDepartment)
	department.Put("/:id", secretHandler, departmentHandler.UpdateDepartment)
	department.Delete("/:id", secretHandler, departmentHandler.DeleteDepartment)

	event := apiV1.Group("/events")
	event.Get("", secretOrPublicHandler, eventHandler.SearchEvent)
	event.Get("/:id", secretOrPublicHandler, eventHandler.GetOneEvent)
	event.Post("", secretHandler, eventHandler.SaveEvent)
	event.Put("/:id", secretHandler, eventHandler.UpdateEvent)
	event.Delete("/:id", secretHandler, eventHandler.DeleteEvent)

	eventReg := apiV1.Group("/eventregister")
	eventReg.Post("/:id", secretHandler, eventHandler.RegisterEvent)

	eventDownload := apiV1.Group("/eventsdownload")
	eventDownload.Post("/:id", secretHandler, eventHandler.DownloadEventTicket)

	// secretHandler, timeout.New(eventHandler.DownloadEventTicket, 2*time.Minute))
	// timeout framework bisa race condition

	user := apiV1.Group("/users")
	user.Get("", publicHandler, userHandler.SearchUser)
	user.Get("/:id", publicHandler, userHandler.GetOneUser)
	user.Post("", secretHandler, userHandler.SaveUser)
	user.Put("/:id", secretHandler, userHandler.UpdateUser)
	user.Delete("/:id", secretHandler, userHandler.DeleteUser)

	classrooms := apiV1.Group("/classrooms")
	classrooms.Get("", publicHandler, classroomHandler.SearchClassroom)
	classrooms.Get("/:id", publicHandler, classroomHandler.GetOneClassroom)
	classrooms.Post("", secretHandler, classroomHandler.SaveClassroom)
	classrooms.Put("/:id", secretHandler, classroomHandler.UpdateClassroom)
	classrooms.Delete("/:id", secretHandler, classroomHandler.DeleteClassroom)

	tempUser := apiV1.Group("/temp_users")
	tempUser.Get("", publicHandler, tempUserHandler.SearchTempUser)
	tempUser.Get("/:id", publicHandler, tempUserHandler.GetOneTempUser)
	tempUser.Post("", publicHandler, tempUserHandler.SaveTempUser)
	tempUser.Put("/:id", secretHandler, tempUserHandler.UpdateTempUser)
	tempUser.Delete("/:id", secretHandler, tempUserHandler.DeleteTempUser)

	verifyUser := apiV1.Group("/verify_user")
	verifyUser.Post("/:id", secretHandler, tempUserHandler.VerifyUser)

	register := apiV1.Group("/register")
	register.Post("/availability", publicHandler, tempUserHandler.CheckAvailabilityUser)
}
