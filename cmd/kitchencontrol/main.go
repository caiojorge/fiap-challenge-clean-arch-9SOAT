package main

import (
	"fmt"
	"log"
	"os"

	"github.com/caiojorge/fiap-challenge-ddd/docs"
	infra "github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driven/db"
	"github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driven/model"
	"github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driver/api/server"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// @title Fiap Fase 2 Challenge Clean Arch API - 9SOAT
// @version 1.0
// @description This is fiap fase 2 challenge project.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /kitchencontrol/api/v1

func main() {

	_ = godotenv.Load() // Carrega o .env se não estiver definido em variáveis de ambiente

	hostname := os.Getenv("HOST_NAME")
	hostport := os.Getenv("HOST_PORT")

	// @host localhost:8083
	// @host localhost:30080

	gin.SetMode(gin.ReleaseMode)

	//logger, err := zap.NewProduction()
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer logger.Sync() // flushes buffer, if any

	db := setupDB(logger)
	server := server.NewServer(db, logger)
	//server.Initialization()

	logger.Info("Server Initialized")

	// Migrate the schema
	logger.Info("Startin Migration")
	setupMigration(server)
	logger.Info("Migration executed successfully")

	//swaggerURL := fmt.Sprintf("http://%s:%s/kitchencontrol/api/v1/docs/*any", hostname, hostport)
	//server.GetRouter().GET("/kitchencontrol/api/v1/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.URL(swaggerURL)))
	swaggerURL := setupSwagger(hostname, hostport, server)

	logger.Info("Server running on " + hostname + ":" + hostport)
	logger.Info("swagger running on " + swaggerURL)

	//server.Run(":8083")
	server.Run(":" + hostport)

}

func setupMigration(server *server.GinServer) {
	if err := server.GetDB().AutoMigrate(
		&model.Customer{},
		&model.Product{},
		&model.Order{},
		&model.OrderItem{},
		&model.Checkout{},
		&model.Kitchen{}); err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}
}

func setupSwagger(hostname string, hostport string, server *server.GinServer) string {

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", hostname, hostport)
	docs.SwaggerInfo.BasePath = "/kitchencontrol/api/v1"

	swaggerURL := fmt.Sprintf("http://%s:%s/kitchencontrol/api/v1/docs/doc.json", hostname, hostport)
	server.GetRouter().GET("/kitchencontrol/api/v1/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return swaggerURL
}

func setupDB(logger *zap.Logger) *gorm.DB {

	// host := "localhost"
	// port := "3306"
	// user := "root"
	// password := "root"
	// dbName := "dbcontrolf2"

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	db := infra.NewDB(host, port, user, password, dbName)

	logger.Info("Database connection established")
	logger.Info(dbName, zap.String("host", host), zap.String("port", port), zap.String("user", user))
	// get a connection
	connection := db.GetConnection("mysql")
	if connection == nil {
		log.Fatal("Expected a non-nil MySQL connection, but got nil")
	}

	return connection
}
