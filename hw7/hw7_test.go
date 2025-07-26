package hw7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDIContainer(t *testing.T) {
	container := NewContainer()
	err := container.RegisterType("UserService", func() interface{} {
		return &UserService{}
	})
	assert.NoError(t, err)

	err = container.RegisterType("MessageService", func() interface{} {
		return &MessageService{}
	})
	assert.NoError(t, err)

	err = container.RegisterType("MessageService", &MessageService{})
	assert.ErrorIs(t, err, ErrNonCallableType)

	userService1, err := container.Resolve("UserService")
	assert.NoError(t, err)
	userService2, err := container.Resolve("UserService")
	assert.NoError(t, err)
	_, err = container.Resolve("PostService")
	assert.ErrorIs(t, err, ErrNoConstructorFound)

	u1 := userService1.(*UserService)
	u2 := userService2.(*UserService)
	assert.False(t, u1 == u2)

	messageService, err := container.Resolve("MessageService")
	assert.NoError(t, err)
	assert.NotNil(t, messageService)

	paymentService, err := container.Resolve("PaymentService")
	assert.Error(t, err)
	assert.Nil(t, paymentService)
}
