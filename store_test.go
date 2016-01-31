package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockedBackend struct {
	mock.Mock
}

func (m *MockedBackend) StoreMap(kvmap map[string]string) {
}

func (m *MockedBackend) ReadMap() map[string]string {
	return make(map[string]string)
}




func Test_Factory(t *testing.T) {
	assert := assert.New(t)
	assert.NotNil(NewStore( new(MockedBackend) ))
}

func Test_WriteOp(t *testing.T) {

	op := Op{
		key: "test",
		value: "testval",
		mode: OP_WRITE,
	}

	store := NewStore( new(MockedBackend) )
	op.handle(store)

	assert := assert.New(t)
	val := store.get("test")
	assert.Equal("testval",val)
}

func Test_ReadOp(t *testing.T) {

	op := Op{
		key: "test",
		mode: OP_READ,
	}

	store := NewStore( new(MockedBackend) )
	op.handle(store)

}