package di

import (
	"context"
	"errors"
	"time"

	"github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driven/converter"
	service "github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driven/gateway"
	repositorygorm "github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driven/repository/gorm"
	controllercheckout "github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driver/api/controller/checkout"
	controllercustomer "github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driver/api/controller/customer"
	controllerkitchen "github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driver/api/controller/kitchen"
	controllerorder "github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driver/api/controller/order"
	controllerproduct "github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driver/api/controller/product"
	usecasecheckoutcheck "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/checkout/checkpayment"
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
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Container struct {
	DB     *gorm.DB
	Logger *zap.Logger

	/* repositories */
	ProductRepo      *repositorygorm.ProductRepositoryGorm
	OrderRepo        *repositorygorm.OrderRepositoryGorm
	CustomerRepo     *repositorygorm.CustomerRepositoryGorm
	CheckoutRepo     *repositorygorm.CheckoutRepositoryGorm
	KitchenRepo      *repositorygorm.KitchenRepositoryGorm
	GatewayService   *service.FakePaymentService
	ProductConverter *converter.ProductConverter
	OrderConverter   *converter.OrderConverter

	/* Use cases */
	// Customer
	CustomerRegisterUseCase  *usecasecustomerregister.CustomerRegisterUseCase
	CustomerUpdateUseCase    *usecasecustomerupdate.CustomerUpdateUseCase
	CustomerFindByCPFUseCase *usecasecustomerfindbycpf.CustomerFindByCPFUseCase
	CustomerFindAllUseCase   *usecasecustomerfindall.CustomerFindAllUseCase

	// Product
	ProductRegisterUseCase       *usecaseproductregister.ProductRegisterUseCase
	ProductUpdateUseCase         *usecaseproductupdater.ProductUpdateUseCase
	ProductFindByIDUseCase       *usecaseproductfindbyid.ProductFindByIDUseCase
	ProductFindAllUseCase        *usecaseproductfindall.ProductFindAllUseCase
	ProductFindByCategoryUseCase *usecaseproductfindbycategory.ProductFindByCategoryUseCase
	ProductDeleteUseCase         *usecaseproductdelete.ProductDeleteUseCase

	// Order
	OrderCreateUseCase       *usecaseordercreate.OrderCreateUseCase
	OrderFindAllUseCase      *usecaseorderfindall.OrderFindAllUseCase
	OrderFindByIDUseCase     *usecaseorderfindbyid.OrderFindByIDUseCase
	OrderFindByParamsUseCase *usecaseorderfindbyparam.OrderFindByParamsUseCase

	// Checkout
	CheckoutCreateUseCase *usecasecheckout.CheckoutCreateUseCase
	CheckoutCheckUseCase  *usecasecheckoutcheck.CheckPaymentUseCase

	// Kitchen
	KitchenFindAllUseCase *usecasekitchen.KitchenFindAllUseCase

	/*Controllers*/

	// Customer
	RegisterCustomerController  *controllercustomer.RegisterCustomerController
	UpdateCustomerController    *controllercustomer.UpdateCustomerController
	FindCustomerByCPFController *controllercustomer.FindCustomerByCPFController
	FindAllCustomersController  *controllercustomer.FindAllCustomersController

	// Product
	RegisterProductController       *controllerproduct.RegisterProductController
	UpdateProductController         *controllerproduct.UpdateProductController
	FindProductByIDController       *controllerproduct.FindProductByIDController
	FindAllProductController        *controllerproduct.FindAllProductController
	FindProductByCategoryController *controllerproduct.FindProductByCategoryController
	DeleteProductController         *controllerproduct.DeleteProductController

	// Order
	CreateOrderController                       *controllerorder.CreateOrderController
	FindAllOrdersController                     *controllerorder.FindAllController
	FindOrderByIDController                     *controllerorder.FindOrderByIDController
	FindByParamsOrdersNotConfirmedController    *controllerorder.FindByParamsNotConfirmedController
	FindByParamsOrdersConfirmedController       *controllerorder.FindByParamsConfirmedController
	FindByParamsOrdersPaymentApprovedController *controllerorder.FindByParamsPaymentApprovedController

	// Checkout
	CreateCheckoutController *controllercheckout.CreateCheckoutController
	CheckoutCheckController  *controllercheckout.CheckPaymentCheckoutController

	// Kitchen
	FindKitchenAllController *controllerkitchen.FindKitchenAllController
}

func NewContainer(db *gorm.DB, logger *zap.Logger) *Container {
	// Initialize converters
	productConverter := converter.NewProductConverter()
	orderConverter := converter.NewOrderConverter()

	// Initialize repositories
	customerRepo := repositorygorm.NewCustomerRepositoryGorm(db)
	productRepo := repositorygorm.NewProductRepositoryGorm(db, productConverter)
	orderRepo := repositorygorm.NewOrderRepositoryGorm(db, orderConverter)
	checkoutRepo := repositorygorm.NewCheckoutRepositoryGorm(db)
	kitchenRepo := repositorygorm.NewKitchenRepositoryGorm(db)
	gatewayService := service.NewFakePaymentService()

	// Initialize use cases
	customerRegisterUseCase := usecasecustomerregister.NewCustomerRegister(customerRepo)
	customerUpdateUseCase := usecasecustomerupdate.NewCustomerUpdate(customerRepo)
	customerFindByCPFUseCase := usecasecustomerfindbycpf.NewCustomerFindByCPF(customerRepo)
	customerFindAllUseCase := usecasecustomerfindall.NewCustomerFindAll(customerRepo)
	productRegisterUseCase := usecaseproductregister.NewProductRegister(productRepo)
	productUpdateUseCase := usecaseproductupdater.NewProductUpdate(productRepo)
	productFindByIDUseCase := usecaseproductfindbyid.NewProductFindByID(productRepo)
	productFindAllUseCase := usecaseproductfindall.NewProductFindAll(productRepo)
	productFindByCategoryUseCase := usecaseproductfindbycategory.NewProductFindByCategory(productRepo)
	productDeleteUseCase := usecaseproductdelete.NewProductDelete(productRepo)
	orderCreateUseCase := usecaseordercreate.NewOrderCreate(orderRepo, customerRepo, productRepo)
	orderFindAllUseCase := usecaseorderfindall.NewOrderFindAll(orderRepo)
	orderFindByIDUseCase := usecaseorderfindbyid.NewOrderFindByID(orderRepo)
	orderFindByParamsUseCase := usecaseorderfindbyparam.NewOrderFindByParams(orderRepo)
	checkoutCreateUseCase := usecasecheckout.NewCheckoutCreate(orderRepo, checkoutRepo, gatewayService, kitchenRepo, productRepo)
	checkoutCheckUseCase := usecasecheckoutcheck.NewCheckPaymentUseCase(checkoutRepo, orderRepo)
	kitchenFindAllUseCase := usecasekitchen.NewKitchenFindAll(kitchenRepo)

	// Initialize controllers
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	registerCustomerController := controllercustomer.NewRegisterCustomerController(ctx, customerRegisterUseCase)
	updateCustomerController := controllercustomer.NewUpdateCustomerController(ctx, customerUpdateUseCase)
	findCustomerByCPFController := controllercustomer.NewFindCustomerByCPFController(ctx, customerFindByCPFUseCase)
	findAllCustomersController := controllercustomer.NewFindAllCustomersController(ctx, customerFindAllUseCase)
	registerProductController := controllerproduct.NewRegisterProductController(ctx, productRegisterUseCase)
	updateProductController := controllerproduct.NewUpdateProductController(ctx, productUpdateUseCase)
	findProductByIDController := controllerproduct.NewFindProductByIDController(ctx, productFindByIDUseCase)
	findAllProductController := controllerproduct.NewFindAllProductController(ctx, productFindAllUseCase)
	findProductByCategoryController := controllerproduct.NewFindProductByCategoryController(ctx, productFindByCategoryUseCase)
	deleteProductController := controllerproduct.NewDeleteProductController(ctx, productDeleteUseCase)
	createOrderController := controllerorder.NewCreateOrderController(ctx, orderCreateUseCase)
	findAllOrdersController := controllerorder.NewFindAllController(ctx, orderFindAllUseCase)
	findOrderByIDController := controllerorder.NewFindOrderByIDController(ctx, orderFindByIDUseCase)
	findByParamsOrdersNotConfirmedController := controllerorder.NewFindByParamsNotConfirmedController(ctx, orderFindByParamsUseCase)
	findByParamsOrdersConfirmedController := controllerorder.NewFindByParamsConfirmedController(ctx, orderFindByParamsUseCase)
	findByParamsOrdersPaymentApprovedController := controllerorder.NewFindByParamsPaymentApprovedController(ctx, orderFindByParamsUseCase)
	createCheckoutController := controllercheckout.NewCreateCheckoutController(ctx, checkoutCreateUseCase)
	checkCheckoutController := controllercheckout.NewCheckPaymentCheckoutController(ctx, checkoutCheckUseCase)
	findKitchenAllController := controllerkitchen.NewFindKitchenAllController(ctx, kitchenFindAllUseCase)

	return &Container{
		DB:                                          db,
		Logger:                                      logger,
		ProductRepo:                                 productRepo,
		OrderRepo:                                   orderRepo,
		CustomerRepo:                                customerRepo,
		CheckoutRepo:                                checkoutRepo,
		KitchenRepo:                                 kitchenRepo,
		GatewayService:                              gatewayService,
		ProductConverter:                            productConverter,
		OrderConverter:                              orderConverter,
		CustomerRegisterUseCase:                     customerRegisterUseCase,
		CustomerUpdateUseCase:                       customerUpdateUseCase,
		CustomerFindByCPFUseCase:                    customerFindByCPFUseCase,
		CustomerFindAllUseCase:                      customerFindAllUseCase,
		ProductRegisterUseCase:                      productRegisterUseCase,
		ProductUpdateUseCase:                        productUpdateUseCase,
		ProductFindByIDUseCase:                      productFindByIDUseCase,
		ProductFindAllUseCase:                       productFindAllUseCase,
		ProductFindByCategoryUseCase:                productFindByCategoryUseCase,
		ProductDeleteUseCase:                        productDeleteUseCase,
		OrderCreateUseCase:                          orderCreateUseCase,
		OrderFindAllUseCase:                         orderFindAllUseCase,
		OrderFindByIDUseCase:                        orderFindByIDUseCase,
		OrderFindByParamsUseCase:                    orderFindByParamsUseCase,
		CheckoutCreateUseCase:                       checkoutCreateUseCase,
		KitchenFindAllUseCase:                       kitchenFindAllUseCase,
		RegisterCustomerController:                  registerCustomerController,
		UpdateCustomerController:                    updateCustomerController,
		FindCustomerByCPFController:                 findCustomerByCPFController,
		FindAllCustomersController:                  findAllCustomersController,
		RegisterProductController:                   registerProductController,
		UpdateProductController:                     updateProductController,
		FindProductByIDController:                   findProductByIDController,
		FindAllProductController:                    findAllProductController,
		FindProductByCategoryController:             findProductByCategoryController,
		DeleteProductController:                     deleteProductController,
		CreateOrderController:                       createOrderController,
		FindAllOrdersController:                     findAllOrdersController,
		FindOrderByIDController:                     findOrderByIDController,
		FindByParamsOrdersNotConfirmedController:    findByParamsOrdersNotConfirmedController,
		FindByParamsOrdersConfirmedController:       findByParamsOrdersConfirmedController,
		FindByParamsOrdersPaymentApprovedController: findByParamsOrdersPaymentApprovedController,
		CreateCheckoutController:                    createCheckoutController,
		CheckoutCheckController:                     checkCheckoutController,
		FindKitchenAllController:                    findKitchenAllController,
	}
}

func (c *Container) Validate() error {
	// check if all dependencies are set
	if c.DB == nil {
		c.Logger.Error("DB is not set")
		return errors.New("db is not set")
	}
	if c.Logger == nil {
		//c.Logger.Error("Logger is not set")
		return errors.New("logger is not set")
	}
	if c.ProductRepo == nil {
		c.Logger.Error("ProductRepo is not set")
		return errors.New("productRepo is not set")
	}
	if c.OrderRepo == nil {
		c.Logger.Error("OrderRepo is not set")
		return errors.New("orderRepo is not set")
	}
	if c.CustomerRepo == nil {
		c.Logger.Error("CustomerRepo is not set")
		return errors.New("customerRepo is not set")
	}
	if c.CheckoutRepo == nil {
		c.Logger.Error("CheckoutRepo is not set")
		return errors.New("checkoutRepo is not set")
	}
	if c.KitchenRepo == nil {
		c.Logger.Error("KitchenRepo is not set")
		return errors.New("kitchenRepo is not set")
	}
	if c.GatewayService == nil {
		c.Logger.Error("GatewayService is not set")
		return errors.New("gatewayService is not set")
	}
	if c.ProductConverter == nil {
		c.Logger.Error("ProductConverter is not set")
		return errors.New("productConverter is not set")
	}
	if c.OrderConverter == nil {
		c.Logger.Error("OrderConverter is not set")
		return errors.New("orderConverter is not set")
	}
	if c.CustomerRegisterUseCase == nil {
		c.Logger.Error("CustomerRegisterUseCase is not set")
		return errors.New("customerRegisterUseCase is not set")
	}
	if c.CustomerUpdateUseCase == nil {
		c.Logger.Error("CustomerUpdateUseCase is not set")
		return errors.New("customerUpdateUseCase is not set")
	}
	if c.CustomerFindByCPFUseCase == nil {
		c.Logger.Error("CustomerFindByCPFUseCase is not set")
		return errors.New("customerFindByCPFUseCase is not set")
	}
	if c.CustomerFindAllUseCase == nil {
		c.Logger.Error("CustomerFindAllUseCase is not set")
		return errors.New("customerFindAllUseCase is not set")
	}
	if c.ProductRegisterUseCase == nil {
		c.Logger.Error("ProductRegisterUseCase is not set")
		return errors.New("productRegisterUseCase is not set")
	}
	if c.ProductUpdateUseCase == nil {
		c.Logger.Error("ProductUpdateUseCase is not set")
		return errors.New("productUpdateUseCase is not set")
	}
	if c.ProductFindByIDUseCase == nil {
		c.Logger.Error("ProductFindByIDUseCase is not set")
		return errors.New("productFindByIDUseCase is not set")
	}
	if c.ProductFindAllUseCase == nil {
		c.Logger.Error("ProductFindAllUseCase is not set")
		return errors.New("productFindAllUseCase is not set")
	}
	if c.ProductFindByCategoryUseCase == nil {
		c.Logger.Error("ProductFindByCategoryUseCase is not set")
		return errors.New("productFindByCategoryUseCase is not set")
	}
	if c.ProductDeleteUseCase == nil {
		c.Logger.Error("ProductDeleteUseCase is not set")
		return errors.New("productDeleteUseCase is not set")
	}
	if c.OrderCreateUseCase == nil {
		c.Logger.Error("OrderCreateUseCase is not set")
		return errors.New("orderCreateUseCase is not set")
	}
	if c.OrderFindAllUseCase == nil {
		c.Logger.Error("OrderFindAllUseCase is not set")
		return errors.New("orderFindAllUseCase is not set")
	}
	if c.OrderFindByIDUseCase == nil {
		c.Logger.Error("OrderFindByIDUseCase is not set")
		return errors.New("orderFindByIDUseCase is not set")
	}
	if c.OrderFindByParamsUseCase == nil {
		c.Logger.Error("OrderFindByParamsUseCase is not set")
		return errors.New("orderFindByParamsUseCase is not set")
	}
	if c.CheckoutCreateUseCase == nil {
		c.Logger.Error("CheckoutCreateUseCase is not set")
		return errors.New("checkoutCreateUseCase is not set")
	}
	if c.KitchenFindAllUseCase == nil {
		c.Logger.Error("KitchenFindAllUseCase is not set")
		return errors.New("kitchenFindAllUseCase is not set")
	}
	if c.RegisterCustomerController == nil {
		c.Logger.Error("RegisterCustomerController is not set")
		return errors.New("registerCustomerController is not set")
	}
	if c.UpdateCustomerController == nil {
		c.Logger.Error("UpdateCustomerController is not set")
		return errors.New("updateCustomerController is not set")
	}
	if c.FindCustomerByCPFController == nil {
		c.Logger.Error("FindCustomerByCPFController is not set")
		return errors.New("findCustomerByCPFController is not set")
	}
	if c.FindAllCustomersController == nil {
		c.Logger.Error("FindAllCustomersController is not set")
		return errors.New("findAllCustomersController is not set")
	}
	if c.RegisterProductController == nil {
		c.Logger.Error("RegisterProductController is not set")
		return errors.New("registerProductController is not set")
	}
	if c.UpdateProductController == nil {
		c.Logger.Error("UpdateProductController is not set")
		return errors.New("updateProductController is not set")
	}
	if c.FindProductByIDController == nil {
		c.Logger.Error("FindProductByIDController is not set")
		return errors.New("findProductByIDController is not set")
	}
	if c.FindAllProductController == nil {
		c.Logger.Error("FindAllProductController is not set")
		return errors.New("findAllProductController is not set")
	}
	if c.FindProductByCategoryController == nil {
		c.Logger.Error("FindProductByCategoryController is not set")
		return errors.New("findProductByCategoryController is not set")
	}
	if c.DeleteProductController == nil {
		c.Logger.Error("DeleteProductController is not set")
		return errors.New("deleteProductController is not set")
	}
	if c.CreateOrderController == nil {
		c.Logger.Error("CreateOrderController is not set")
		return errors.New("createOrderController is not set")
	}
	if c.FindAllOrdersController == nil {
		c.Logger.Error("FindAllOrdersController is not set")
		return errors.New("findAllOrdersController is not set")
	}
	if c.FindOrderByIDController == nil {
		c.Logger.Error("FindOrderByIDController is not set")
		return errors.New("findOrderByIDController is not set")
	}
	if c.FindByParamsOrdersNotConfirmedController == nil {
		c.Logger.Error("FindByParamsOrdersNotConfirmedController is not set")
		return errors.New("findByParamsOrdersNotConfirmedController is not set")
	}
	if c.FindByParamsOrdersConfirmedController == nil {
		c.Logger.Error("FindByParamsOrdersConfirmedController is not set")
		return errors.New("findByParamsOrdersConfirmedController is not set")
	}
	if c.FindByParamsOrdersPaymentApprovedController == nil {
		c.Logger.Error("FindByParamsOrdersPaymentApprovedController is not set")
		return errors.New("findByParamsOrdersPaymentApprovedController is not set")
	}
	if c.CreateCheckoutController == nil {
		c.Logger.Error("CreateCheckoutController is not set")
		return errors.New("createCheckoutController is not set")
	}
	if c.FindKitchenAllController == nil {
		c.Logger.Error("FindKitchenAllController is not set")
		return errors.New("findKitchenAllController is not set")
	}
	return nil
}
