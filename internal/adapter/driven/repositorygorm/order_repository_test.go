package repositorygorm

import (
	"context"
	"testing"
	"time"

	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driven/converter"
	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driven/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateOrder(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//db := setupMysql()

	// Migrar o esquema
	err = db.AutoMigrate(&model.Customer{}, &model.Product{}, &model.Order{}, &model.OrderItem{})
	if err != nil {
		panic("failed to migrate database")
	}

	converter := converter.NewOrderConverter()
	repo := NewOrderRepositoryGorm(db, converter)

	// customer
	customer := model.Customer{
		CPF:   "75419654059", //75419654059
		Name:  "John Doe",
		Email: "email@email.com",
	}

	product := model.Product{
		ID:          "1",
		Name:        "Product 1",
		Description: "Description 1",
		Category:    "Category 1",
		Price:       10.0,
	}

	location, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		panic("failed to load location")
	}

	orderItem := model.OrderItem{
		ID:        "111",
		ProductID: "1",
		Product:   product,
		Quantity:  1,
		Price:     10.0,
		Status:    "pending",
	}

	order := model.Order{
		ID: "111",
		Items: []*model.OrderItem{
			&orderItem,
		},
		Total:       10.0,
		Status:      "pending",
		CustomerCPF: &customer.CPF,
		Customer:    &customer,
		CreatedAt:   time.Now().In(location),
	}

	//log.Default().Println(order)

	//result := db.Create(&order)
	//assert.Nil(t, result.Error)
	entity := converter.ToEntity(&order)
	err = repo.Create(context.Background(), entity)
	assert.Nil(t, err)

}

// func setupMysql() *gorm.DB {
// 	host := "localhost"
// 	port := "3306"
// 	user := "root"
// 	password := "root"
// 	dbName := "dbcontrol"

// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)
// 	fmt.Println(dsn)
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("Failed to connect to the database: %v", err)
// 	}

// 	return db
// }
