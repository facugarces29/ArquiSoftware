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
	GetAddressByUserId(id int) (dto.AddressesDto, error)
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

func (s *addressService) GetAddressByUserId(id int) (dto.AddressesDto, error) {

	addresses, err := addressCliente.GetAddressByUserId(id)
	var addressesDto dto.AddressesDto

	if err != nil {
		return addressesDto, err
	}

	for _, address := range addresses {
		var adressDto dto.AddressDto
		adressDto.Id = address.AddressID
		adressDto.UserId = address.UserID
		adressDto.AddressLine = address.Addressline
		adressDto.State = address.State
		adressDto.City = address.City
		adressDto.Zip = address.Zip

		addressesDto = append(addressesDto, adressDto)
	}

	return addressesDto, nil
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
