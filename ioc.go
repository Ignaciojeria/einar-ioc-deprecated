package ioc

import (
	"errors"

	"github.com/google/uuid"
)

type container[T any] struct {
	loadableDependency func() (T, error)
	isLoaded           bool
	Dependency         T
}

func (c *container[T]) Load() (any, error) {
	if c.isLoaded {
		return nil, errors.New("dependency already loaded")
	}
	instance, err := c.loadableDependency()
	c.Dependency = instance
	c.isLoaded = true
	return instance, err
}

type loadable[T any] interface {
	Load() (any, error)
}

var installations = make(map[string]loadable[any])

func Installation[T any](loadableDependency func() (T, error)) *container[T] {
	adapter := container[T]{loadableDependency: loadableDependency}
	installations[uuid.NewString()] = &adapter
	return &adapter
}

var useCases = make(map[string]loadable[any])

func UseCase[T any](loadableDependency func() (T, error)) *container[T] {
	adapter := container[T]{loadableDependency: loadableDependency}
	useCases[uuid.NewString()] = &adapter
	return &adapter
}

var inboundAdapters = make(map[string]loadable[any])

func InboundAdapter[T any](loadableDependency func() (T, error)) *container[T] {
	adapter := container[T]{loadableDependency: loadableDependency}
	inboundAdapters[uuid.NewString()] = &adapter
	return &adapter
}

var outBoundAdapters = make(map[string]loadable[any])

func OutBoundAdapter[T any](loadableDependency func() (T, error)) *container[T] {
	adapter := container[T]{loadableDependency: loadableDependency}
	outBoundAdapters[uuid.NewString()] = &adapter
	return &adapter
}

func LoadDependencies() error {
	for _, v := range installations {
		_, err := v.Load()
		if err != nil {
			return err
		}
	}
	for _, v := range outBoundAdapters {
		_, err := v.Load()
		if err != nil {
			return err
		}
	}
	for _, v := range useCases {
		_, err := v.Load()
		if err != nil {
			return err
		}
	}
	for _, v := range inboundAdapters {
		_, err := v.Load()
		if err != nil {
			return err
		}
	}
	return nil
}
