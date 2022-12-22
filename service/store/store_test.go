package store

import (
	"testing"

	"github.com/ehrktia/performance-monitor/service/store/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_store_service(t *testing.T) {
	testUid := 1
	mockCntrl := gomock.NewController(t)
	mockRepo := mocks.NewMockRepository(mockCntrl)
	gomock.InOrder(
		mockRepo.EXPECT().GetByID(testUid).MinTimes(1).Return(t.Name(), nil),
	)
	mockStoreService := &storeService{
		store: mockRepo,
	}
	got, err := mockStoreService.RetreiveNameByID(testUid)
	assert.Nil(t, err)
	assert.Equal(t, got, t.Name())

}
