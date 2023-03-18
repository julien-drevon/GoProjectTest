package gateway

type Repo[T any] struct {
	Context []T
}

func (this *Repo[T]) Get() []T {
	return this.Context
}

func (this *Repo[T]) Add(data T) {
	this.Context = append(this.Context, data)
}

func NewRepo[T any]() Repo[T] {
	return Repo[T]{Context: []T{}}
}

func NewRepoWithContext[T any](context []T) Repo[T] {
	return Repo[T]{Context: context}
}
