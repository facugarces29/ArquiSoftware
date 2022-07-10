package services

import (
	addressCliente "proyecto/ArquiSoftware/clients/address"
	dto "proyecto/ArquiSoftware/dto/address"
)

type addressService struct{}

type addressServiceInterface interface {
	GetAddressByUserId(id int) (dto.AddressDto, error)
}

var (
	AddressService addressServiceInterface
)

func init() {
	AddressService = &addressService{}
}

func (s *addressService) GetAddressByUserId(id int) (dto.AddressDto, error) {

	address, err := addressCliente.GetAddressById(id)
	var addressDto dto.AddressDto

	if err != nil {
		return addressDto, err
	}
	if address.UserID == 0 {
		return addressDto, nil
	}

	addressDto.Id = address.AddressID
	addressDto.UserId = address.UserID
	addressDto.AddressLine = address.Addressline
	addressDto.City = address.City
	addressDto.State = address.State
	addressDto.Zip = address.Zip

	return addressDto, nil
}
