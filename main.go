package main

import (
	"crowdfunding/auth"
	"crowdfunding/campaign"
	"crowdfunding/handler"
	"crowdfunding/helper"
	"crowdfunding/user"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/db_crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)

	userService := user.NewService(userRepository)
	campaignService := campaign.NewService(campaignRepository)

	authService := auth.NewService()
	userHandler := handler.NewUserHandler(userService, authService)

	test, err := campaignService.GetCampaign(0)

	fmt.Println(len(test))

	router := gin.Default()
	api := router.Group("api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	router.Run("localhost:8080")

}

//ambil nilai header authorization : Bearer tokentokentoken
//dari header authorization, kita ambil nilai tokennya saja
//kita validasi token
//kita ambil user_id
//ambil user dari db berdasarkan user_id lewat service
//kita set context dari isinya user

func authMiddleware(authService auth.Servive, userService user.Service) gin.HandlerFunc {

	return func(c *gin.Context) {
		auhtHeader := c.GetHeader("Authorization")

		if !strings.Contains(auhtHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(auhtHeader, " ")

		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)

		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserById(userID)

		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}
