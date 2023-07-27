package main

import (
	"log"
	"main/pkg/config"
	di "main/pkg/di"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}

	// //db:=db.ConnectDatabase()
	// db.Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic(err)
	// }
	// db.Db.AutoMigrate(&models.User{},&models.Admin{},&models.Product{},&models.Category{})
	// //DB.Db.Model(&models.Product{}).BelongsTo(&models.Category{}, "Category_ID")

	// //User
	// //user signup(with validation),user login , admin signin , OTP login , list Products , Products details page
	// router.POST("/signup", handlers.Signup)
	// router.POST("/login", handlers.Login)
	// router.POST("/loginwithotp", handlers.Loginwithotp)
	// router.POST("/loginverification",handlers.OtpValidation)
	// router.GET("/home", handlers.Home)
	// router.GET("/product",handlers.Product)
	// router.POST("/logout", handlers.Logout)

	// //Admin
	// //List users, user management(block/unblock), category management(add,edit,delete), product  management(add,edit,delete)
	// router.POST("/adminregistration",handlers.AdminRegistration)
	// router.POST("/adminlogin",handlers.AdminLogin)
	// router.POST("/adminloginwithotp", handlers.AdminLoginwithotp)
	// router.POST("/adminloginverification",handlers.AdminOtpValidation)
	// router.GET("/adminpanel", handlers.AdminPanel)

	// router.POST("/addproduct",handlers.AddProduct)
	// router.PUT("/editproduct",handlers.EditProduct)
	// router.DELETE("/deleteproduct",handlers.DeleteProduct)

	// router.POST("/addcategory",handlers.AddCategory)
	// router.PUT("/editcategory",handlers.EditCategory)
	// router.DELETE("/deletecategory",handlers.DeleteCategory)
	// router.POST("/permissiontoggle",handlers.PermissionToggle)

	// router.Run(":1243")
}
