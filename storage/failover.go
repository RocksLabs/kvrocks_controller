package storage

import (
	"errors"

	"github.com/KvrocksLabs/kvrocks_controller/storage/persistence/etcd"
)

func (s *Storage) UpdateDoingFailOverTask(task *etcd.FailOverTask) error {
	s.rw.Lock()
	defer s.rw.Unlock()

	if task == nil {
		return errors.New("nil fail over task")
	}
	return s.instance.UpdateDoingFailOverTask(task)
}

func (s *Storage) GetDoingFailOverTask(ns, cluster string) (*etcd.FailOverTask, error) {
	s.rw.RLock()
	defer s.rw.RUnlock()

	return s.instance.GetDoingFailOverTask(ns, cluster)
}

func (s *Storage) AddFailOverHistory(task *etcd.FailOverTask) error {
	s.rw.Lock()
	defer s.rw.Unlock()

	return s.instance.AddFailOverHistory(task)
}

func (s *Storage) GetFailOverHistory(ns, cluster string) ([]*etcd.FailOverTask, error) {
	s.rw.RLock()
	defer s.rw.RUnlock()

	return s.instance.GetFailOverHistory(ns, cluster)
}
