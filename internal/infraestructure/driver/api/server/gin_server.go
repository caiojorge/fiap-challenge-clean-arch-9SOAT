package server

import (
	"context"
	"log"
	"time"

	"github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driven/converter"
	service "github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driven/gateway"
	repositorygorm "github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driven/repository/gorm"
	controllercheckout "github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driver/api/controller/checkout"
	controllercustomer "github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driver/api/controller/customer"
	controllerkitchen "github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driver/api/controller/kitchen"
	controllerorder "github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driver/api/controller/order"
	controllerproduct "github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driver/api/controller/product"
	usecasecheckout "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/checkout/create"
	usecasecustomerfindall "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/customer/findall"
	usecasecustomerfindbycpf "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/customer/findbycpf"
	usecasecustomerregister "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/customer/register"
	usecasecustomerupdate "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/customer/update"
	usecasekitchen "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/kitchen/findall"
	usecaseordercreate "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/order/create"
	usecaseorderfindall "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/order/findall"
	usecaseorderfindbyid "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/order/findbyid"
	usecaseorderfindbyparam "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/order/findbyparam"
	usecaseproductdelete "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/product/delete"
	usecaseproductfindall "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/product/findall"
	usecaseproductfindbycategory "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/product/findbycategory"
	usecaseproductfindbyid "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/product/findbyid"
	usecaseproductregister "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/product/register"
	usecaseproductupdater "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/product/update"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type GinServer struct {
	router *gin.Engine
	db     *gorm.DB
	logger *zap.Logger
}

func NewServer(db *gorm.DB, logger *zap.Logger) *GinServer {
	r := gin.Default()

	// Configurar CORS
	r.Use(corsMiddleware())

	return &GinServer{
		router: r, db: db,
		logger: logger,
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// Responder diretamente às requisições OPTIONS (pré-voo)
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
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
		registerController := controllercustomer.NewRegisterCustomerController(ctx, usecasecustomerregister.NewCustomerRegister(customerRepo))
		g.POST("/", registerController.PostRegisterCustomer)

		updateController := controllercustomer.NewUpdateCustomerController(ctx, usecasecustomerupdate.NewCustomerUpdate(customerRepo))
		g.PUT("/:cpf", updateController.PutUpdateCustomer)

		findByCPFController := controllercustomer.NewFindCustomerByCPFController(ctx, usecasecustomerfindbycpf.NewCustomerFindByCPF(customerRepo))
		g.GET("/:cpf", findByCPFController.GetCustomerByCPF)

		findAllController := controllercustomer.NewFindAllCustomersController(ctx, usecasecustomerfindall.NewCustomerFindAll(customerRepo))
		g.GET("/", findAllController.GetAllCustomers)
	}

	p := s.router.Group("/kitchencontrol/api/v1/products")
	{
		registerController := controllerproduct.NewRegisterProductController(ctx, usecaseproductregister.NewProductRegister(productRepo))
		p.POST("/", registerController.PostRegisterProduct)

		findAllController := controllerproduct.NewFindAllProductController(ctx, usecaseproductfindall.NewProductFindAll(productRepo))
		p.GET("/", findAllController.GetAllProducts)

		findByIDController := controllerproduct.NewFindProductByIDController(ctx, usecaseproductfindbyid.NewProductFindByID(productRepo))
		p.GET("/:id", findByIDController.GetProductByID)

		findByCategoryController := controllerproduct.NewFindProductByCategoryController(ctx, usecaseproductfindbycategory.NewProductFindByCategory(productRepo))
		p.GET("/category/:id", findByCategoryController.GetProductByCategory)

		updateController := controllerproduct.NewUpdateProductController(ctx, usecaseproductupdater.NewProductUpdate(productRepo))
		p.PUT("/:id", updateController.PutUpdateProduct)

		deleteController := controllerproduct.NewDeleteProductController(ctx, usecaseproductdelete.NewProductDelete(productRepo))
		p.DELETE("/:id", deleteController.DeleteProduct)

	}

	o := s.router.Group("/kitchencontrol/api/v1/orders")
	{

		orderController := controllerorder.NewCreateOrderController(ctx, usecaseordercreate.NewOrderCreate(orderRepo, customerRepo, productRepo))
		o.POST("/", orderController.PostCreateOrder)

		findAllOrdersController := controllerorder.NewFindAllController(ctx, usecaseorderfindall.NewOrderFindAll(orderRepo))
		o.GET("/", findAllOrdersController.GetAllOrders)

		findByIDController := controllerorder.NewFindOrderByIDController(ctx, usecaseorderfindbyid.NewOrderFindByID(orderRepo))
		o.GET("/:id", findByIDController.GetOrderByID)

		findByParamsOrdersController := controllerorder.NewFindByParamsController(ctx, usecaseorderfindbyparam.NewOrderFindByParams(orderRepo))
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
