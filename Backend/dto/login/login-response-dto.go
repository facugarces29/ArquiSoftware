package dto

type LoginResponseDto struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
}
