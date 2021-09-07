package application_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	application "github.com/javielrezende/go-hexagonal/application"
	mock_application "github.com/javielrezende/go-hexagonal/application/mocks"
	"github.com/stretchr/testify/require"
)

// Utilizado o mockgen para gerar os mocks em cima das nossas interfaces. Pasta mock dentro de application
// Para gerar, rodar dentro do container:
// mockgen -destination=application/mocks/application.go -source=application/product.go application

//defer funciona como se fosse um await. É executado quando o resto finalizar
func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)

	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Get("1")
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)

	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Create("Product 1", 10)
	require.Nil(t, err)
	require.Equal(t, product, result)
}
