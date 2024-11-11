package server

import (
	"context"
	"log"
	"time"

	usecasecheckout "github.com/caiojorge/fiap-challenge-ddd/internal/core/usecase/checkout/create"
	usecasecustomerfindall "github.com/caiojorge/fiap-challenge-ddd/internal/core/usecase/customer/findall"
	usecasecustomerfindbycpf "github.com/caiojorge/fiap-challenge-ddd/internal/core/usecase/customer/findbycpf"
	usecasecustomerregister "github.com/caiojorge/fiap-challenge-ddd/internal/core/usecase/customer/register"
	usecasecustomerupdate "github.com/caiojorge/fiap-challenge-ddd/internal/core/usecase/customer/update"
	usecasekitchen "github.com/caiojorge/fiap-challenge-ddd/internal/core/usecase/kitchen/findall"
	usecaseordercreate "github.com/caiojorge/fiap-challenge-ddd/internal/core/usecase/order/create"
	usecaseorderfindall "github.com/caiojorge/fiap-challenge-ddd/internal/core/usecase/order/findall"
	usecaseorderfindbyid "github.com/caiojorge/fiap-challenge-ddd/internal/core/usecase/order/findbyid"
	usecaseorderfindbyparam "github.com/caiojorge/fiap-challenge-ddd/internal/core/usecase/order/findbyparam"
	usecaseproductdelete "github.com/caiojorge/fiap-challenge-ddd/internal/core/usecase/product/delete"
	usecaseproductfindall "github.com/caiojorge/fiap-challenge-ddd/internal/core/usecase/product/findall"
	usecaseproductfindbycategory "github.com/caiojorge/fiap-challenge-ddd/internal/core/usecase/product/findbycategory"
	usecaseproductfindbyid "github.com/caiojorge/fiap-challenge-ddd/internal/core/usecase/product/findbyid"
	usecaseproductregister "github.com/caiojorge/fiap-challenge-ddd/internal/core/usecase/product/register"
	usecaseproductupdater "github.com/caiojorge/fiap-challenge-ddd/internal/core/usecase/product/update"
	"github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driven/converter"
	service "github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driven/gateway"
	"github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driven/repositorygorm"
	controllercheckout "github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driver/api/controller/checkout"
	controllercustomer "github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driver/api/controller/customer"
	controllerkitchen "github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driver/api/controller/kitchen"
	controllerorder "github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driver/api/controller/order"
	controllerproduct "github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driver/api/controller/product"
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
	return &GinServer{
		router: r, db: db,
		logger: logger,
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
