package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/naqvijafar91/focus"
	"github.com/rs/cors"
	"github.com/spf13/viper"

	"github.com/naqvijafar91/focus/email"
	"github.com/naqvijafar91/focus/handlers"
	"github.com/naqvijafar91/focus/mysql"
)

func main() {
	smux := http.NewServeMux()
	// These are memory backed services
	// folderService := &memorybackedservices.DummyFolderService{}
	// userService := &memorybackedservices.DummyUserService{}
	// taskService := &memorybackedservices.DummyTaskService{}
	userLoginService, folderService, taskService, userService := initServices()
	handlers.NewFolderHandler(folderService).RegisterFolderRoutes(smux)
	handlers.NewHandler(userLoginService).RegisterUserRoutes(smux)
	handlers.NewTaskHandler(taskService).RegisterTaskRoutes(smux)
	aggregatorService := focus.NewAggregatorService(taskService,
		folderService,
		userService)
	handlers.NewAggregatorHandler(aggregatorService).RegisterAggregatorRoutes(smux)
	fmt.Println("Server starting")
	handler := cors.AllowAll().Handler(smux)
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func initServices() (focus.UserLoginService, focus.FolderService, focus.TaskService, focus.UserService) {
	viper.SetConfigFile("./config.env")
	viper.SetConfigType("env")  // REQUIRED if the config file does not have the extension in the name
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	viper.SetDefault("dbhost", "localhost")
	viper.SetDefault("dbport", 3306)
	viper.SetDefault("dbname", "focus")
	viper.SetDefault("dbuser", "focus")
	viper.SetDefault("dbpassword", "pwd")

	dbHost := viper.GetString("dbhost")
	dbPwd := viper.GetString("dbpassword")
	dbPort := viper.GetInt("dbport")
	dbName := viper.GetString("dbname")
	dbUser := viper.GetString("dbuser")
	db, err := mysql.NewMysqlConn(dbHost, dbPort, dbUser, dbName, dbPwd)
	if err != nil {
		panic(err)
	}
	userService, err := mysql.NewUserService(db)
	folderService, err := mysql.NewFolderService(db)
	taskService, err := mysql.NewTaskService(db, folderService)
	if err != nil {
		panic(err)
	}
	codeGenerator := focus.NewFourDigitCodeGenerator()
	notificationService := email.NewLoginCodeNotificationSender()
	userLoginService := focus.NewUserLoginService(notificationService, userService, codeGenerator)
	return userLoginService, folderService, taskService, userService
}
