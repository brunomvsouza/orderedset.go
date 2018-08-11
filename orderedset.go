// Copyright (c) 2018, Bruno M V Souza <github@b.bmvs.io>. All rights reserved.
// Use of this source code is governed by a BSD-2-Clause license that can be
// found in the LICENSE file.

// Package orderedset implements a thread safe insertion ordered set
// backed by an ordered map.
package orderedset // import "go.bmvs.io/orderedset"

import (
	"sync"

	"go.bmvs.io/orderedmap"
)

// New return a new Set implemented by OrderedSet
func New() *OrderedSet {
	return &OrderedSet{
		store: orderedmap.New(),
		index: make(map[interface{}]int),
	}
}

// OrderedSet insertion ordered Set implementation
type OrderedSet struct {
	sync.Mutex

	// the underlying store of the Set
	store *orderedmap.OrderedMap
	index map[interface{}]int
	// currentIndex keeps track of the keys of the underlying store
	currentIndex int
}

// Add add items to the OrderedSet
func (s *OrderedSet) Add(items ...interface{}) {
	for _, item := range items {
		if _, found := s.index[item]; found {
			continue
		}

		s.put(item)
	}
}

// Remove remove items from the OrderedSet
func (s *OrderedSet) Remove(items ...interface{}) {
	for _, item := range items {
		index, found := s.index[item]
		if !found {
			return
		}

		s.remove(item, index)
	}
}

// Contains return if OrderedSet contains the specified items or not
func (s *OrderedSet) Contains(items ...interface{}) bool {
	for _, item := range items {
		if _, found := s.index[item]; !found {
			return false
		}
	}
	return true
}

// Empty return if the OrderedSet in empty or not
func (s *OrderedSet) Empty() bool {
	return s.store.Empty()
}

// Values return the values of the OrderedSet in insertion order
func (s *OrderedSet) Values() []interface{} {
	return s.store.Values()
}

// Size return the size of the OrderedSet
func (s *OrderedSet) Size() int {
	return s.store.Size()
}

func (s *OrderedSet) put(item interface{}) {
	s.Lock()
	defer s.Unlock()

	s.store.Put(s.currentIndex, item)
	s.index[item] = s.currentIndex
	s.currentIndex++
}

func (s *OrderedSet) remove(item interface{}, index int) {
	s.Lock()
	defer s.Unlock()

	s.store.Remove(index)
	delete(s.index, item)
}
