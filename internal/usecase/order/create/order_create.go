package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	domainRepository "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository"

	//portsrepository "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository"

	"github.com/jinzhu/copier"
)

// OrderCreateUseCase é a implementação de OrderCreateUseCase.
// Um caso de uso é uma estrutura que contém todas as regras de negócio para uma determinada funcionalidade.
// Nesse cenário, vamos precisar acessar 3 agregados e seus repositorios.
type OrderCreateUseCase struct {
	orderRepository    domainRepository.OrderRepository
	customerRepository domainRepository.CustomerRepository
	productRepository  domainRepository.ProductRepository
}

func NewOrderCreate(orderRepository domainRepository.OrderRepository,
	customerRepository domainRepository.CustomerRepository,
	productRepository domainRepository.ProductRepository) *OrderCreateUseCase {
	return &OrderCreateUseCase{
		orderRepository:    orderRepository,
		customerRepository: customerRepository,
		productRepository:  productRepository,
	}
}

// CreateOrder registra um novo pedido.
func (cr *OrderCreateUseCase) CreateOrder(ctx context.Context, input *OrderCreateInputDTO) (*OrderCreateOutputDTO, error) {

	wrapper := &CreateOrderWrapper{
		dto: input,
	}

	// converte o DTO de input para a entidade Order
	order, err := wrapper.ToEntity()
	if err != nil {
		return nil, err
	}

	fmt.Println("usecase: Criando Order: " + order.CustomerCPF)

	// handle customer
	// se o cpf for empty, indica que o cliente não quis se identificar, e isso esta ok, segundo as regras de negócio
	if order.IsCustomerInformed() {
		// busca o cliente pelo cpf
		customer, err := cr.customerRepository.Find(ctx, order.CustomerCPF)
		if err != nil {
			return nil, err
		}

		// se o cliente for informado, temos q validar o cpf, e cadastra-lo caso não exista
		// só entra aqui se o cliente não existir na base de dados
		// se o cliente não for nulo, ele já existe o cpf é considerado válido
		// em teoria, não existe ordem duplicada. o mesmo cliente pode comprar várias vezes. (não vou validar isso aqui)
		if customer == nil {
			// o cliente não é obrigatório, mas se for informado, ele precisa ser válido.
			// apenas nesse caso, se o cliente não existir, ele será persistido. (apenas o cpf)
			// identifica o cliente pelo cpf
			newCustomer, err := entity.NewCustomerWithCPFInformed(order.CustomerCPF)
			if err != nil {
				return nil, err
			}

			// cria o cliente sem o nome e email
			if err := cr.customerRepository.Create(ctx, newCustomer); err != nil {
				return nil, err
			}
		}

	}

	fmt.Println("usecase: Criando Order: Items: " + order.CustomerCPF)

	// handle items
	// valida se os produtos informados existem e busca o preço
	for _, item := range order.Items {
		product, err := cr.productRepository.Find(ctx, item.ProductID)
		if err != nil {
			return nil, err
		}

		if product == nil {
			return nil, errors.New("product not found")
		}

		// atualiza o preço do produto
		item.ConfirmPrice(product.Price)
	}

	fmt.Println("usecase: Criando Order: ConfirmOrder: " + order.CustomerCPF)
	// toda regra de negócio para criar uma ordem confirmada
	order.Confirm()

	// cria a ordem e usa o cliente (novo ou existente) e o produto existente.
	if err := cr.orderRepository.Create(ctx, order); err != nil {
		fmt.Println("usecase: repo create: " + err.Error())
		return nil, err
	}

	// copia os dados da ordem para o output
	var output OrderCreateOutputDTO
	err = copier.Copy(&output, &order)
	if err != nil {
		return nil, err
	}

	return &output, nil
}

// CreateOrder registra um novo pedido.
func (cr *OrderCreateUseCase) createOrder2(ctx context.Context, input *OrderCreateInputDTO) (*OrderCreateOutputDTO, error) {
	// Copia os dados do DTO para a entidade Order
	var order entity.Order
	if err := copier.Copy(&order, &input); err != nil {
		return nil, fmt.Errorf("erro ao copiar dados para a entidade Order: %w", err)
	}

	// Remove a máscara do CPF
	order.RemoveMaksFromCPF()

	// Log utilizando um logger estruturado (use seu logger preferido)
	log.Default().Println("Criando Order", "CPF", order.CustomerCPF)

	// Se o cliente foi informado, tenta encontrá-lo ou cria um novo
	if err := cr.handleCustomer(ctx, &order); err != nil {
		return nil, err
	}

	// Valida os itens da ordem, verificando se os produtos existem e definindo seus preços
	if err := cr.validateOrderItems(ctx, &order); err != nil {
		return nil, err
	}

	// Confirma a criação da ordem
	order.Confirm()

	// Persiste a ordem no repositório
	if err := cr.orderRepository.Create(ctx, &order); err != nil {
		log.Default().Println("erro ao persistir a ordem", "error", err)
		return nil, fmt.Errorf("erro ao criar a ordem: %w", err)
	}

	// Copia os dados da entidade Order para o DTO de saída
	var output OrderCreateOutputDTO
	if err := copier.Copy(&output, &order); err != nil {
		return nil, fmt.Errorf("erro ao copiar dados da ordem para o DTO de saída: %w", err)
	}

	return &output, nil
}

// handleCustomer lida com a validação e criação do cliente, se necessário
func (cr *OrderCreateUseCase) handleCustomer(ctx context.Context, order *entity.Order) error {
	if !order.IsCustomerInformed() {
		return nil // cliente não informado, não precisa tratar
	}

	// Busca o cliente no repositório
	customer, err := cr.customerRepository.Find(ctx, order.CustomerCPF)
	if err != nil {
		return fmt.Errorf("erro ao buscar cliente: %w", err)
	}

	// Se o cliente já existe, nada mais a ser feito
	if customer != nil {
		return nil
	}

	// Cria um novo cliente com base no CPF informado
	newCustomer, err := entity.NewCustomerWithCPFInformed(order.CustomerCPF)
	if err != nil {
		return fmt.Errorf("erro ao criar cliente com CPF informado: %w", err)
	}

	// Persiste o novo cliente
	if err := cr.customerRepository.Create(ctx, newCustomer); err != nil {
		return fmt.Errorf("erro ao criar novo cliente: %w", err)
	}

	return nil
}

// validateOrderItems valida os itens da ordem e ajusta os preços dos produtos
func (cr *OrderCreateUseCase) validateOrderItems(ctx context.Context, order *entity.Order) error {
	for i, item := range order.Items {
		// Busca o produto pelo ID
		product, err := cr.productRepository.Find(ctx, item.ProductID)
		if err != nil {
			return fmt.Errorf("erro ao buscar produto: %w", err)
		}

		// Se o produto não for encontrado, retorna um erro
		if product == nil {
			return fmt.Errorf("produto não encontrado: %s", item.ProductID)
		}

		// Confirma o preço do produto no item da ordem
		order.Items[i].ConfirmPrice(product.Price)
	}
	return nil
}
