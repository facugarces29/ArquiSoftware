package services

import (
	model "proyecto/ArquiSoftware/model/address"
	"testing"

	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

type AddressClientInterface struct {
	mock.Mock
}

func (m *AddressClientInterface) GetAddressById(id int) model.Address {
	ret := m.Called(id)
	return ret.Get(0).(model.Address)
}

func (m *AddressClientInterface) GetAddressByUserId(id int) model.Address {
	ret := m.Called(id)
	return ret.Get(0).(model.Address)
}

func TestGetAddressByUserId(t *testing.T) {

	mockClient := new(AddressClientInterface)
	var address model.Address
	address.AddressID = 1

	var addressBad model.Address
	addressBad.AddressID = 0

	mockClient.On("GetAddressByUserId", 1).Return(address)
	mockClient.On("GetAddressByUserId", -1).Return(addressBad)
	service := initAddressService(mockClient)

	_, err := service.GetAddressById(1)
	assert.Nil(t, err, "Error should be nil")

	_, err2 := service.GetAddressById(-1)
	assert.NotNil(t, err2, "Error should not be nil")

}

func TestGetAddressById(t *testing.T) {

	mockClient := new(AddressClientInterface)
	var address model.Address
	address.AddressID = 1

	var addressBad model.Address
	addressBad.AddressID = 0

	mockClient.On("GetAddressById", 1).Return(address)
	mockClient.On("GetAddressById", -1).Return(addressBad)
	service := initAddressService(mockClient)

	_, err := service.GetAddressById(1)
	assert.Nil(t, err, "Error should be nil")

	_, err2 := service.GetAddressById(-1)
	assert.NotNil(t, err2, "Error should not be nil")

}
