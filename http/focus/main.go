package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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
	userLoginService, folderService, taskService, userService, port := initServices()
	handlers.NewFolderHandler(folderService).RegisterFolderRoutes(smux)
	handlers.NewHandler(userLoginService).RegisterUserRoutes(smux)
	handlers.NewTaskHandler(taskService).RegisterTaskRoutes(smux)
	aggregatorService := focus.NewAggregatorService(taskService,
		folderService,
		userService)
	handlers.NewAggregatorHandler(aggregatorService).RegisterAggregatorRoutes(smux)
	fmt.Println("Server starting")
	handler := cors.AllowAll().Handler(smux)
	log.Fatal(http.ListenAndServe(port, handler))
}

func initServices() (focus.UserLoginService, focus.FolderService, focus.TaskService, focus.UserService, string) {
	viper.SetConfigFile("./config.env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		fmt.Println(fmt.Errorf(`Failed to read config file, please make sure all properties are present
		in the environment variables: %s`, err))
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
	emailForSending := viper.GetString("email")
	emailPwd := viper.GetString("password")
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
	notificationService, err := email.NewLoginCodeNotificationSender(emailForSending, emailPwd)
	if err != nil {
		panic(err)
	}
	userLoginService := focus.NewUserLoginService(notificationService, userService, codeGenerator, folderService)
	// Figure out port now
	port := viper.GetInt("port")
	if port == 0 {
		port = 8080
	}
	fmt.Printf("Starting on port %d\n", port)
	portStr := ":" + strconv.Itoa(port)
	return userLoginService, folderService, taskService, userService, portStr
}
