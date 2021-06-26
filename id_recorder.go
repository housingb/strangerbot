package main

import "sync"

type IdRecorder struct {
	data *sync.Map
}

func NewIdRecorder() *IdRecorder {
	return &IdRecorder{
		data: new(sync.Map),
	}
}

func (m *IdRecorder) IsSent(id int) bool {

	_, ok := m.data.Load(id)

	return ok
}

func (m *IdRecorder) SetSent(id int) {

	m.data.Store(id, true)

}
