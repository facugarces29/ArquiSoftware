package services

import (
	addressCliente "proyecto/ArquiSoftware/clients/address"
	dto "proyecto/ArquiSoftware/dto/address"
	model "proyecto/ArquiSoftware/model/address"
)

type addressService struct {
	addressClient addressCliente.AddressClientInterface
}

type addressServiceInterface interface {
	GetAddressByUserId(id int) (dto.AddressDto, error)
	GetAddressById(id int) (dto.AddressDto, error)
}

var (
	AddressService addressServiceInterface
)

func init() {
	AddressService = &addressService{}
}

func initAddressService(addressClient addressCliente.AddressClientInterface) addressServiceInterface {
	service := new(addressService)
	service.addressClient = addressClient
	return service
}

func (s *addressService) GetAddressByUserId(id int) (dto.AddressDto, error) {

	var address model.Address
	var addressDto dto.AddressDto
	address, err := addressCliente.GetAddressByUserId(id)

	if address.AddressID == 0 {
		return addressDto, err
	}

	addressDto.Id = address.AddressID
	addressDto.UserId = address.UserID
	addressDto.AddressLine = address.Addressline
	addressDto.State = address.State
	addressDto.City = address.City
	addressDto.Zip = address.Zip

	return addressDto, nil
}

func (s *addressService) GetAddressById(id int) (dto.AddressDto, error) {

	var address model.Address
	var addressDto dto.AddressDto
	address, err := addressCliente.GetAddressById(id)

	if address.AddressID == 0 {
		return addressDto, err
	}

	addressDto.Id = address.AddressID
	addressDto.UserId = address.UserID
	addressDto.AddressLine = address.Addressline
	addressDto.City = address.City
	addressDto.State = address.State
	addressDto.Zip = address.Zip

	return addressDto, nil
}
