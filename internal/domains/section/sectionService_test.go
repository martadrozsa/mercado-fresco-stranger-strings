package section_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/section"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/section/mocks"
)

var expectedSection = section.Section{
	SectionNumber:      int64(1),
	CurrentTemperature: int64(1),
	MinimumTemperature: int64(1),
	CurrentCapacity:    int64(1),
	MinimumCapacity:    int64(1),
	MaximumCapacity:    int64(1),
	WarehouseId:        int64(1),
	ProductTypeId:      int64(1),
}

var expectedUpdatedSection = section.Section{
	Id:                 int64(1),
	SectionNumber:      int64(1),
	CurrentTemperature: int64(1),
	MinimumTemperature: int64(1),
	CurrentCapacity:    int64(5),
	MinimumCapacity:    int64(1),
	MaximumCapacity:    int64(1),
	WarehouseId:        int64(1),
	ProductTypeId:      int64(1),
}

func TestSectionService_Create(t *testing.T) {
	mockRepository := mocks.NewRepository(t)

	t.Run("create_ok: If it contains the required fields, it will be created", func(t *testing.T) {
		mockRepository.
			On("Create",
				expectedSection.SectionNumber,
				expectedSection.CurrentTemperature,
				expectedSection.MinimumTemperature,
				expectedSection.CurrentCapacity,
				expectedSection.MinimumCapacity,
				expectedSection.MaximumCapacity,
				expectedSection.WarehouseId,
				expectedSection.ProductTypeId,
			).
			Return(expectedSection, nil).
			Once()

		service := section.NewService(mockRepository)
		_, err := service.Create(1, 1, 1, 1, 1, 1, 1, 1)

		assert.Nil(t, err)
	})

	t.Run("create_conflict: If section_number already exists it cannot be created", func(t *testing.T) {
		errorConflict := fmt.Errorf("Already a section with the code: %d", expectedSection.SectionNumber)

		mockRepository.
			On("Create",
				expectedSection.SectionNumber,
				expectedSection.CurrentTemperature,
				expectedSection.MinimumTemperature,
				expectedSection.CurrentCapacity,
				expectedSection.MinimumCapacity,
				expectedSection.MaximumCapacity,
				expectedSection.WarehouseId,
				expectedSection.ProductTypeId,
			).
			Return(section.Section{}, errorConflict).
			Once()

		service := section.NewService(mockRepository)
		result, err := service.Create(1, 1, 1, 1, 1, 1, 1, 1)

		assert.Equal(t, section.Section{}, result)
		assert.Equal(t, errorConflict, err)
	})
}

func TestSectionService_GetAll(t *testing.T) {
	t.Run("get_all: If the list has elements, it will return an amount of the total elements", func(t *testing.T) {
		mockRepository := mocks.NewRepository(t)
		mockRepository.
			On("GetAll").
			Return([]section.Section{expectedSection}, nil).
			Once()

		service := section.NewService(mockRepository)
		result, err := service.GetAll()

		assert.Nil(t, err)
		assert.Equal(t, []section.Section{expectedSection}, result)
	})

	t.Run("get_all: ", func(t *testing.T) {
		mockRepository := mocks.NewRepository(t)
		mockRepository.
			On("GetAll").
			Return([]section.Section{}, fmt.Errorf("any error")).
			Once()

		service := section.NewService(mockRepository)
		_, err := service.GetAll()

		assert.NotNil(t, err)
	})

}

func TestSectionService_GetById(t *testing.T) {
	mockRepository := mocks.NewRepository(t)

	t.Run("find_by_id_existent: If the element searched for by id exists, it will be return", func(t *testing.T) {
		mockRepository.
			On("GetById", int64(1)).
			Return(expectedSection, nil).
			Once()

		service := section.NewService(mockRepository)
		result, err := service.GetById(1)

		assert.Nil(t, err)
		assert.Equal(t, expectedSection, result)

	})

	t.Run("find_by_id_non_existent: If the element searched for by id does not exist, return null", func(t *testing.T) {
		id := int64(3)
		errorNotFound := fmt.Errorf("Section %d not found", id)

		mockRepository.
			On("GetById", id).
			Return(section.Section{}, errorNotFound).
			Once()

		service := section.NewService(mockRepository)
		result, err := service.GetById(id)

		assert.Equal(t, section.Section{}, result)
		assert.Equal(t, errorNotFound, err)
	})
}

func TestSectionService_Delete(t *testing.T) {
	mockRepository := mocks.NewRepository(t)

	t.Run("delete_ok: If the deletion is successful, the item will not appear in the list", func(t *testing.T) {
		mockRepository.
			On("Delete", int64(1)).
			Return(nil).
			Once()

		service := section.NewService(mockRepository)
		err := service.Delete(1)

		assert.Nil(t, err)
	})

	t.Run("delete_non_existent: When the section does not exist, null will be returned", func(t *testing.T) {
		id := int64(3)
		errorNotFound := fmt.Errorf("Section %d not found", id)
		mockRepository.
			On("Delete", id).
			Return(errorNotFound).
			Once()

		service := section.NewService(mockRepository)
		err := service.Delete(id)

		assert.Equal(t, errorNotFound, err)
	})
}

func TestSectionService_Update(t *testing.T) {
	mockRepository := mocks.NewRepository(t)

	t.Run("update_existent: When the data update is successful, the section with the updated information will be returned", func(t *testing.T) {
		mockRepository.
			On("UpdateCurrentCapacity", int64(1), int64(5)).
			Return(expectedUpdatedSection, nil).
			Once()

		service := section.NewService(mockRepository)
		result, err := service.UpdateCurrentCapacity(int64(1), int64(5))

		assert.Nil(t, err)
		assert.Equal(t, expectedUpdatedSection.CurrentCapacity, result.CurrentCapacity)
	})

	t.Run("update_non_existent: If the section to be updated does not exist, null will be returned.", func(t *testing.T) {
		id := int64(3)
		errorNotFound := fmt.Errorf("Section %d not found", id)
		mockRepository.
			On("UpdateCurrentCapacity", id, int64(5)).
			Return(section.Section{}, errorNotFound).
			Once()

		service := section.NewService(mockRepository)
		result, err := service.UpdateCurrentCapacity(id, int64(5))

		assert.Equal(t, errorNotFound, err)
		assert.Equal(t, section.Section{}, result)
	})

}
