package server

import (
	"context"
	"log"
	"time"

	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driven/converter"
	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driven/repositorygorm"
	controllercheckout "github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driver/api/controller/checkout"
	controllercustomer "github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driver/api/controller/customer"
	controllerkitchen "github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driver/api/controller/kitchen"
	controllerorder "github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driver/api/controller/order"
	controllerproduct "github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driver/api/controller/product"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/application/service"
	usecasecheckout "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/usecase/checkout"
	usecasecustomer "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/usecase/customer"
	usecasekitchen "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/usecase/kitchen"
	usecaseorder "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/usecase/order"
	usecaseproduct "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/usecase/product"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GinServer struct {
	router *gin.Engine
	db     *gorm.DB
}

func NewServer(db *gorm.DB) *GinServer {
	r := gin.Default()
	return &GinServer{router: r, db: db}
}

func (s *GinServer) GetDB() *gorm.DB {
	return s.db
}

func (s *GinServer) Initialization() *GinServer {

	//db := setupSQLite()
	//s.db = setupDB()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	customerRepo := repositorygorm.NewCustomerRepositoryGorm(s.db)
	productConverter := converter.NewProductConverter()
	productRepo := repositorygorm.NewProductRepositoryGorm(s.db, productConverter)
	orderConverter := converter.NewOrderConverter()
	orderRepo := repositorygorm.NewOrderRepositoryGorm(s.db, orderConverter)
	checkoutRepo := repositorygorm.NewCheckoutRepositoryGorm(s.db)
	kitchenRepo := repositorygorm.NewKitchenRepositoryGorm(s.db)
	gatewayService := service.NewFakePaymentService()

	g := s.router.Group("/kitchencontrol/api/v1/customers")
	{
		registerController := controllercustomer.NewRegisterCustomerController(ctx, usecasecustomer.NewCustomerRegister(customerRepo))
		g.POST("/", registerController.PostRegisterCustomer)

		updateController := controllercustomer.NewUpdateCustomerController(ctx, usecasecustomer.NewCustomerUpdate(customerRepo))
		g.PUT("/:cpf", updateController.PutUpdateCustomer)

		findByCPFController := controllercustomer.NewFindCustomerByCPFController(ctx, usecasecustomer.NewCustomerFindByCPF(customerRepo))
		g.GET("/:cpf", findByCPFController.GetCustomerByCPF)

		findAllController := controllercustomer.NewFindAllCustomersController(ctx, usecasecustomer.NewCustomerFindAll(customerRepo))
		g.GET("/", findAllController.GetAllCustomers)
	}

	p := s.router.Group("/kitchencontrol/api/v1/products")
	{
		registerController := controllerproduct.NewRegisterProductController(ctx, usecaseproduct.NewProductRegister(productRepo))
		p.POST("/", registerController.PostRegisterProduct)

		findAllController := controllerproduct.NewFindAllProductController(ctx, usecaseproduct.NewProductFindAll(productRepo))
		p.GET("/", findAllController.GetAllProducts)

		findByIDController := controllerproduct.NewFindProductByIDController(ctx, usecaseproduct.NewProductFindByID(productRepo))
		p.GET("/:id", findByIDController.GetProductByID)

		findByCategoryController := controllerproduct.NewFindProductByCategoryController(ctx, usecaseproduct.NewProductFindByCategory(productRepo))
		p.GET("/category/:id", findByCategoryController.GetProductByCategory)

		updateController := controllerproduct.NewUpdateProductController(ctx, usecaseproduct.NewProductUpdate(productRepo))
		p.PUT("/:id", updateController.PutUpdateProduct)

		deleteController := controllerproduct.NewDeleteProductController(ctx, usecaseproduct.NewProductDelete(productRepo))
		p.DELETE("/:id", deleteController.DeleteProduct)

	}

	o := s.router.Group("/kitchencontrol/api/v1/orders")
	{

		orderController := controllerorder.NewCreateOrderController(ctx, usecaseorder.NewOrderCreate(orderRepo, customerRepo, productRepo))
		o.POST("/", orderController.PostCreateOrder)

		findAllOrdersController := controllerorder.NewFindAllController(ctx, usecaseorder.NewOrderFindAll(orderRepo))
		o.GET("/", findAllOrdersController.GetAllOrders)

		findByIDController := controllerorder.NewFindOrderByIDController(ctx, usecaseorder.NewOrderFindByID(orderRepo))
		o.GET("/:id", findByIDController.GetOrderByID)

		findByParamsOrdersController := controllerorder.NewFindByParamsController(ctx, usecaseorder.NewOrderFindByParams(orderRepo))
		o.GET("/paid", findByParamsOrdersController.GetByParamsOrders)

	}

	c := s.router.Group("/kitchencontrol/api/v1/checkouts")
	{
		checkoutController := controllercheckout.NewCreateCheckoutController(ctx, usecasecheckout.NewCheckoutCreate(orderRepo, checkoutRepo, gatewayService, kitchenRepo, productRepo))
		c.POST("/", checkoutController.PostCreateCheckout)
	}

	k := s.router.Group("/kitchencontrol/api/v1/kitchens")
	{
		ktController := controllerkitchen.NewFindKitchenAllController(ctx, usecasekitchen.NewKitchenFindAll(kitchenRepo))
		k.GET("/orders", ktController.GetAllOrdersInKitchen)
	}

	return s
}

func (s *GinServer) Run(port string) {
	if err := s.router.Run(port); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}

func (s *GinServer) GetRouter() *gin.Engine {
	return s.router
}
