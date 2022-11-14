package main

import (
	"log"
	"pustaka-api/book"
	"pustaka-api/config"
	"pustaka-api/handler"
	"pustaka-api/jwtUse"
	"pustaka-api/middleware"
	"pustaka-api/users"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func init() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("error load env")
	}
}

var (
	jwtService     jwtUse.JWTService = jwtUse.NewJWTService()
	db             *gorm.DB          = config.SetupDatabaseConnection()
	bookRepository book.Repository   = book.NewRepository(db)
	bookService    book.Service      = book.NewService(bookRepository)
	userRepository users.Repository  = users.NewRepository(db)
	userService    users.Service     = users.NewService(userRepository)
)

func main() {
	//book
	db.AutoMigrate(book.Book{})
	bookHandler := handler.NewBookHandler(bookService)
	//users
	db.AutoMigrate(users.User{})
	userHandler := handler.NewUserHandler(userService)

	//Router
	r := gin.Default()
	r.SetTrustedProxies([]string{"0.0.0.0:8080"})
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{MaxAge: 60 * 60 * 24}) // expire in a day
	r.Use(sessions.Sessions("mysession", store))

	bookRoutes := r.Group("book/", middleware.AuthorizeJWT(jwtService))
	{
		bookRoutes.GET("/", bookHandler.GetBookHandler)
		bookRoutes.GET("/:id", bookHandler.GeBookID)
		bookRoutes.POST("/", bookHandler.PostBookHandler)
	}
	r.POST("/login", userHandler.Login)
	r.POST("/SignUp", userHandler.SignUp)

	r.Run() // listen and serve on 0.0.0.0:8080
}

// result
// handler
// service
// repository
// gorm connection
// mysql
