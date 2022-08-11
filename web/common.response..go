package web

type CommonResponse[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
}
