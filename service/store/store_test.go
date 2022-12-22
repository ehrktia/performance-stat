package store

import (
	"fmt"
	"testing"

	"github.com/ehrktia/performance-stats/service/store/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_store_service(t *testing.T) {
	mockCntrl := gomock.NewController(t)
	mockRepo := mocks.NewMockRepository(mockCntrl)
	mockStoreService := &storeService{
		store: mockRepo,
	}
	t.Run("retreive user by id from store", func(t *testing.T) {
		// input setup
		testUid := 1
		// expect
		gomock.InOrder(
			mockRepo.
				EXPECT().GetByID(testUid).
				MinTimes(1).Return([]byte(t.Name()), nil),
		)
		// test
		got, err := mockStoreService.RetreiveNameByID(testUid)
		// validate
		assert.Nil(t, err)
		assert.Equal(t, got, t.Name())

	})
	t.Run("fail when user is missing in store", func(t *testing.T) {
		// input setup
		testUid := 10
		testErr := fmt.Errorf("%v", t.Name())
		// expect
		gomock.InOrder(
			mockRepo.
				EXPECT().GetByID(testUid).
				MinTimes(1).Return([]byte(defaultUser), testErr),
		)
		// test
		got, err := mockStoreService.RetreiveNameByID(testUid)
		// validate
		assert.NotNil(t, err)
		assert.Equal(t, got, defaultUser)
	})
}
