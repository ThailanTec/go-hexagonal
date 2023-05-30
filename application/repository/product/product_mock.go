package repository

import "github.com/stretchr/testify/mock"

type ProductMock struct {
	mock.Mock
}

func (m *ProductMock) IsValid() (bool, error) {
	args := m.Called()

	return args.Bool(0), args.Error(1)
}

func (m *ProductMock) Enable() error {
	args := m.Called()

	return args.Error(0)
}

func (m *ProductMock) Disable() error {
	args := m.Called()

	return args.Error(0)
}

func (m *ProductMock) GetID() string {
	args := m.Called()

	return args.String(0)
}

func (m *ProductMock) GetName() string {
	args := m.Called()

	return args.String(0)
}

func (m *ProductMock) GetPrice() string {
	args := m.Called()

	return args.String(0)
}

func (m *ProductMock) GetStatus() string {
	args := m.Called()

	return args.String(0)
}
