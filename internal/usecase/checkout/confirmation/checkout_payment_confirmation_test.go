package usecase

import (
	"context"
	"testing"

	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	mocks "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestConfirmPayment_Success(t *testing.T) {
	ctrl := gomock.NewController(t) // Cria um controlador do gomock
	defer ctrl.Finish()             // Libera os mocks após o teste

	// Mocks
	orderRepoMock := mocks.NewMockOrderRepository(ctrl)
	checkoutRepoMock := mocks.NewMockCheckoutRepository(ctrl)
	transactionManagerMock := mocks.NewMockTransactionManager(ctrl)

	// Use Case
	checkoutConfirmationUseCase := NewCheckoutConfirmation(orderRepoMock, checkoutRepoMock, transactionManagerMock)

	// Dados de entrada e entidades simuladas
	ctx := context.Background()
	input := &CheckoutConfirmationInputDTO{
		OrderID: "order123",
	}
	order := &entity.Order{
		ID:     "order123",
		Status: "pending",
	}
	checkout := &entity.Checkout{
		ID:      "checkout123",
		OrderID: "order123",
		Status:  "pending",
	}

	// Configuração dos mocks
	orderRepoMock.EXPECT().Find(ctx, input.OrderID).Return(order, nil)
	checkoutRepoMock.EXPECT().FindbyOrderID(ctx, input.OrderID).Return(checkout, nil)
	orderRepoMock.EXPECT().Update(ctx, gomock.Any()).Return(nil)
	checkoutRepoMock.EXPECT().Update(ctx, gomock.Any()).Return(nil)
	transactionManagerMock.EXPECT().
		RunInTransaction(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, fn func(ctx context.Context) error) error {
			return fn(ctx) // Executa a função transacional passada
		})

	// Execução do teste
	output, err := checkoutConfirmationUseCase.ConfirmPayment(ctx, input)

	// Asserções
	assert.NoError(t, err)
	assert.NotNil(t, output)
}
