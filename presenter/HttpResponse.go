package presenter

type HttpResponse[T any] struct {
	Data   T
	Status int
	Error  string
}
