package services

import (
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"

	model "proyecto/ArquiSoftware/model/address"
	"testing"
)

type AddressClientInterface struct {
	mock.Mock
}

func (m *AddressClientInterface) GetAddressById(id int) model.Address {
	ret := m.Called(id)
	return ret.Get(0).(model.Address)
}

func (m *AddressClientInterface) GetAddressByUserId(id int) model.Addresses {
	ret := m.Called(id)
	return ret.Get(0).(model.Addresses)
}

func TestGetAddressByUserId(t *testing.T) {

	mockClient := new(AddressClientInterface)

	var addresses model.Addresses
	var address model.Address
	addresses = append(addresses, address)

	var empty model.Addresses

	mockClient.On("GetAddressByUserId", 1).Return(addresses) // Not empty addresses
	mockClient.On("GetAddressByUserId", 2).Return(empty)     // Empty addresses
	service := initAddressService(mockClient)
	res, err := service.GetAddressByUserId(1)
	assert.Nil(t, err, "Error should be nil")
	assert.NotEqual(t, 0, res)

	res2, err2 := service.GetAddressByUserId(2)
	assert.NotNil(t, err2, "Error should not be nil")
	assert.Equal(t, 0, res2)

}

func TestGetAddressById(t *testing.T) {

	mockClient := new(AddressClientInterface)

	var addresses model.Addresses
	var address model.Address
	var empty model.Addresses

	address.AddressID = 1
	var addressBad model.Address
	addressBad.AddressID = 0

	mockClient.On("GetAddressesByUserId", 1).Return(addresses) // Not empty addresses
	mockClient.On("GetAddressesByUserId", 2).Return(empty)     // Empty addresses
	service := initAddressService(mockClient)
	res, err := service.GetAddressByUserId(1)
	assert.Nil(t, err, "Error should be nil")
	assert.NotEqual(t, 0, len(res))

	res2, err2 := service.GetAddressByUserId(2)
	assert.NotNil(t, err2, "Error should not be nil")
	assert.Equal(t, 0, len(res2))

}
