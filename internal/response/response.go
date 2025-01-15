package response

type HTTPResponse[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
}
