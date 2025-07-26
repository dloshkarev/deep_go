package hw7

import "errors"

var (
	ErrNonCallableType    = errors.New("constructor is not of type Callable")
	ErrNoConstructorFound = errors.New("no constructor found")
)

type UserService struct {
	// not need to implement
	NotEmptyStruct bool
}
type MessageService struct {
	// not need to implement
	NotEmptyStruct bool
}

type Callable = func() interface{}

type Container struct {
	types map[string]Callable
}

func NewContainer() *Container {
	return &Container{
		types: make(map[string]Callable),
	}
}

func (c *Container) RegisterType(name string, constructor interface{}) error {
	if callable, ok := constructor.(Callable); ok {
		c.types[name] = callable
		return nil
	}

	return ErrNonCallableType
}

func (c *Container) Resolve(name string) (interface{}, error) {
	if constructor, exists := c.types[name]; exists {
		return constructor(), nil
	}

	return nil, ErrNoConstructorFound
}
