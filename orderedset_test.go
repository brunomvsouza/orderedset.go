// Copyright (c) 2018, Bruno M V Souza <github@b.bmvs.io>. All rights reserved.
// Use of this source code is governed by a BSD-2-Clause license that can be
// found in the LICENSE file.

package orderedset_test

import (
	"testing"

	"go.bmvs.io/orderedset"
)

func TestOrderedSet_Add(t *testing.T) {
	s := orderedset.New()
	s.Add("e", "f", "g", "c", "d", "x", "b", "a")
	s.Add("b") //overwrite
	structValue := complexType{"svalue"}
	s.Add(structValue)
	s.Add(&structValue)
	s.Add(true)

	actualOutput := s.Values()
	expectedOutput := []interface{}{"e", "f", "g", "c", "d", "x", "b", "a", structValue, &structValue, true}
	if !sameElements(actualOutput, expectedOutput) {
		t.Errorf("Got %v expected %v", actualOutput, expectedOutput)
	}
}

func TestOrderedSet_Remove(t *testing.T) {
	s := orderedset.New()
	s.Add("e", "f", "g", "c", "d", "x", "b", "a")
	s.Add("b") //overwrite
	structValue := complexType{"svalue"}
	s.Add(structValue)
	s.Add(&structValue)
	s.Add(true)

	s.Remove("f", "g", &structValue, true)

	actualOutput := s.Values()
	expectedOutput := []interface{}{"e", "c", "d", "x", "b", "a", structValue}
	if !sameElements(actualOutput, expectedOutput) {
		t.Errorf("Got %v expected %v", actualOutput, expectedOutput)
	}

	// already removed
	s.Remove("f", "g", &structValue, true)
	actualOutput = s.Values()
	expectedOutput = []interface{}{"e", "c", "d", "x", "b", "a", structValue}
	if !sameElements(actualOutput, expectedOutput) {
		t.Errorf("Got %v expected %v", actualOutput, expectedOutput)
	}
}

func TestOrderedSet_Contains(t *testing.T) {
	s := orderedset.New()
	s.Add("e", "f", "g", "c", "d", "x", "b", "a")
	s.Add("b") //overwrite
	structValue := complexType{"svalue"}
	s.Add(structValue)
	s.Add(&structValue)
	s.Add(true)

	table := []struct {
		input          []interface{}
		expectedOutput bool
	}{
		{[]interface{}{"c", "d", &structValue}, true},
		{[]interface{}{"c", "d", nil}, false},
		{[]interface{}{true}, true},
		{[]interface{}{structValue}, true},
	}

	for _, test := range table {
		actualOutput := s.Contains(test.input...)
		if actualOutput != test.expectedOutput {
			t.Errorf("Got %v expected %v", actualOutput, test.expectedOutput)
		}
	}
}

func TestOrderedSet_Empty(t *testing.T) {
	s := orderedset.New()
	if empty := s.Empty(); !empty {
		t.Errorf("Got %v expected %v", empty, true)
	}
	s.Add("e", "f", "g", "c", "d", "x", "b", "a")
	if empty := s.Empty(); empty {
		t.Errorf("Got %v expected %v", empty, false)
	}
	s.Remove("e", "f", "g", "c", "d", "x", "b", "a")
	if empty := s.Empty(); !empty {
		t.Errorf("Got %v expected %v", empty, true)
	}
}

func TestOrderedSet_Values(t *testing.T) {
	s := orderedset.New()
	s.Add("e", "f", "g", "c", "d", "x", "b", "a")
	s.Add("b") //overwrite
	structValue := complexType{"svalue"}
	s.Add(structValue)
	s.Add(&structValue)
	s.Add(true)

	actualOutput := s.Values()
	expectedOutput := []interface{}{"e", "f", "g", "c", "d", "x", "b", "a", structValue, &structValue, true}
	if !sameElements(actualOutput, expectedOutput) {
		t.Errorf("Got %v expected %v", actualOutput, expectedOutput)
	}
}

func TestOrderedSet_Size(t *testing.T) {
	s := orderedset.New()
	if size := s.Size(); size != 0 {
		t.Errorf("Got %v expected %v", size, 0)
	}
	s.Add("e", "f", "g", "c", "d", "x", "b", "a")
	s.Add("e", "f", "g", "c", "d", "x", "b", "a", "z") // overwrite
	if size := s.Size(); size != 9 {
		t.Errorf("Got %v expected %v", size, 9)
	}
	s.Remove("e", "f", "g", "c", "d", "x", "b", "a", "z")
	if size := s.Size(); size != 0 {
		t.Errorf("Got %v expected %v", size, 0)
	}
}

func BenchmarkOrderedSet_Add(b *testing.B) {
	s := orderedset.New()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Add(i ^ 2)
	}
}

func BenchmarkOrderedSet_Remove(b *testing.B) {
	s := orderedset.New()
	for i := 0; i < b.N; i++ {
		s.Add(i ^ 2)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Remove(i ^ 2)
	}
}

var resultBenchmarkOrderedSetContains bool

func BenchmarkOrderedSet_Contains(b *testing.B) {
	s := orderedset.New()
	for i := 0; i < b.N; i++ {
		s.Add(i ^ 2)
	}

	var contains bool

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		contains = s.Contains(i ^ 2)
	}

	resultBenchmarkOrderedSetContains = contains
}

var resultBenchmarkOrderedSetEmpty bool

func BenchmarkOrderedSet_Empty(b *testing.B) {
	s := orderedset.New()
	for i := 0; i < b.N; i++ {
		s.Add(i ^ 2)
	}

	var empty bool

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		empty = s.Empty()
	}

	resultBenchmarkOrderedSetEmpty = empty
}

var resultBenchmarkOrderedSetSize int

func BenchmarkOrderedSet_Size(b *testing.B) {
	s := orderedset.New()
	for i := 0; i < b.N; i++ {
		s.Add(i ^ 2)
	}

	var size int

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		size = s.Size()
	}

	resultBenchmarkOrderedSetSize = size
}

var resultBenchmarkOrderedSetValues []interface{}

func BenchmarkOrderedSet_Values(b *testing.B) {
	s := orderedset.New()
	for i := 0; i < b.N; i++ {
		s.Add(i ^ 2)
	}

	var values []interface{}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		values = s.Values()
	}

	resultBenchmarkOrderedSetValues = values
}

type complexType struct {
	foo string
}

func sameElements(a []interface{}, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for _, av := range a {
		found := false
		for _, bv := range b {
			if av == bv {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
