package buyer_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/buyer"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/buyer/mocks"
)

var expectBuyer = &buyer.Buyer{
	Id:           1,
	CardNumberId: 402323,
	FirstName:    "FirstNameTest",
	LastName:     "LastNameTest",
}

var expectBuyerList = []buyer.Buyer{
	{
		CardNumberId: 402323,
		FirstName:    "FirstNameTest",
		LastName:     "LastNameTest",
	},
	{
		CardNumberId: 402300,
		FirstName:    "FirstNameTest 2",
		LastName:     "LastNameTestTest2",
	},
}

func TestService_Create(t *testing.T) {

	t.Run("crete_ok: Se contiver os campos necessários, será criado", func(t *testing.T) {
		repo := mocks.NewRepository(t)

		repo.On("Create", int64(402323), "FirstNameTest", "LastNameTest").Return(expectBuyer, nil)

		service := buyer.NewService(repo)

		result, _ := service.Create(int64(402323), "FirstNameTest", "LastNameTest")

		assert.Equal(t, expectBuyer, result)
	})

	t.Run("create_conflict: Se o card_number_id já existir, ele não pode ser criado", func(t *testing.T) {

		repo := mocks.NewRepository(t)

		repo.On("Create", int64(402323), "FirstNameTest", "LastNameTest").Return(expectBuyer, fmt.Errorf("Card number id is not unique."))

		service := buyer.NewService(repo)

		buyer, err := service.Create(int64(402323), "FirstNameTest", "LastNameTest")

		assert.NotNil(t, err)
		assert.Empty(t, buyer)
		assert.Equal(t, err.Error(), "Card number id is not unique.")

	})
}

func TestService_GetAll(t *testing.T) {

	t.Run("find_all: Se a lista tiver elementos, retornará o número total de elementos", func(t *testing.T) {
		repo := mocks.NewRepository(t)

		repo.On("GetAll").Return(expectBuyerList, nil)

		service := buyer.NewService(repo)

		buyerList, _ := service.GetAll()

		assert.Equal(t, expectBuyerList, buyerList)
	})
}

func TestService_GetId(t *testing.T) {
	t.Run("find_by_id_non_existent: Se o elemento procurado por id não existir, retorna null", func(t *testing.T) {
		repo := mocks.NewRepository(t)

		repo.On("GetId", int64(4)).Return(nil, fmt.Errorf("Buyer not found."))

		service := buyer.NewService(repo)

		buyer, err := service.GetId(int64(4))

		assert.Nil(t, buyer)
		assert.NotNil(t, err)
	})

	t.Run("find_by_id_existent: Se o elemento procurado por id existir, ele retornará as informações do elemento solicitado", func(t *testing.T) {

		repo := mocks.NewRepository(t)

		repo.On("GetId", int64(1)).Return(expectBuyer, nil)
		service := buyer.NewService(repo)

		buyer, err := service.GetId(int64(1))

		assert.Nil(t, err)
		assert.Equal(t, buyer, expectBuyer)

	})
}

func TestService_Update(t *testing.T) {

	t.Run("update_existent: Quando a atualização dos dados for bem sucedida, o comprador será devolvido com as informações atualizadas", func(t *testing.T) {
		repo := mocks.NewRepository(t)

		repo.On("Update", int64(1), int64(456), "LastNameTest 2").Return(expectBuyer, nil)

		service := buyer.NewService(repo)

		buyer, err := service.Update(int64(1), int64(456), "LastNameTest 2")

		assert.Equal(t, expectBuyer, buyer)
		assert.Nil(t, err)

	})

	t.Run("update_non_existent: Se o comprador a ser atualizado não existir, será   retornado null.", func(t *testing.T) {
		repo := mocks.NewRepository(t)
		repo.On("Update", int64(1), int64(456), "LastNameTest 2").Return(nil, fmt.Errorf("Buyer not found."))
		service := buyer.NewService(repo)

		buyer, err := service.Update(int64(1), int64(456), "LastNameTest 2")
		assert.Nil(t, buyer)
		assert.NotNil(t, err)
	})

}

func TestService_Delete(t *testing.T) {
	t.Run("delete_non_existent: Quando o comprador não existir, será retornado null.", func(t *testing.T) {
		repo := mocks.NewRepository(t)
		repo.On("Delete", int64(1)).Return(fmt.Errorf("buyer not found."))
		service := buyer.NewService(repo)

		err := service.Delete(int64(1))

		assert.NotNil(t, err)
	})

	t.Run("delete_ok: Se a exclusão for bem-sucedida, o item não aparecerá na lista.", func(t *testing.T) {
		repo := mocks.NewRepository(t)
		repo.On("Delete", int64(1)).Return(nil)
		service := buyer.NewService(repo)

		err := service.Delete(int64(1))

		assert.Nil(t, err)
	})
}
