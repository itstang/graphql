package resolver

import "context"

type QueryResolver struct {
}

func NewRoot() (*QueryResolver, error) {
	return &QueryResolver{}, nil
}

func (r QueryResolver) Title(ctx context.Context) *string {
	resp := "hello world"
	return &resp
}
