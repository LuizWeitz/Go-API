package models

type Success struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}

type SuccessData[T any] struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   T      `json:"data"`
}

type SuccessList[T any] struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   []T    `json:"data"`
}
