package hw4

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestCircularQueue(t *testing.T) {
	data := NewOrderedMap[Int, Int]()
	assert.Zero(t, data.Size())

	data.Insert(10, 10)
	data.Insert(5, 5)
	data.Insert(15, 15)
	data.Insert(2, 2)
	data.Insert(4, 4)
	data.Insert(12, 12)
	data.Insert(14, 14)

	assert.Equal(t, 7, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(3))
	assert.False(t, data.Contains(13))

	var keys []Int
	expectedKeys := []Int{2, 4, 5, 10, 12, 14, 15}
	data.ForEach(func(key, _ Int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))

	data.Erase(15)
	data.Erase(14)
	data.Erase(2)

	assert.Equal(t, 4, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(2))
	assert.False(t, data.Contains(14))

	keys = nil
	expectedKeys = []Int{4, 5, 10, 12}
	data.ForEach(func(key, _ Int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))
}

func TestPersons(t *testing.T) {
	persons := NewOrderedMap[Person, int]()
	persons.Insert(Person{10}, 10)
	persons.Insert(Person{5}, 5)
	persons.Insert(Person{15}, 15)
	persons.Insert(Person{2}, 2)
	persons.Insert(Person{4}, 4)
	persons.Insert(Person{12}, 12)
	persons.Insert(Person{14}, 14)

	assert.Equal(t, 7, persons.Size())
	assert.True(t, persons.Contains(Person{4}))
	assert.True(t, persons.Contains(Person{10}))
	assert.False(t, persons.Contains(Person{3}))
	assert.False(t, persons.Contains(Person{13}))

	var ages []int
	expectedAges := []int{2, 4, 5, 10, 12, 14, 15}
	persons.ForEach(func(person Person, _ int) {
		ages = append(ages, person.age)
	})
	assert.True(t, reflect.DeepEqual(expectedAges, ages))

	persons.Erase(Person{15})
	persons.Erase(Person{14})
	persons.Erase(Person{2})

	assert.Equal(t, 4, persons.Size())
	assert.True(t, persons.Contains(Person{4}))
	assert.True(t, persons.Contains(Person{12}))
	assert.False(t, persons.Contains(Person{2}))
	assert.False(t, persons.Contains(Person{14}))

	ages = nil
	expectedAges = []int{4, 5, 10, 12}
	persons.ForEach(func(person Person, _ int) {
		ages = append(ages, person.age)
	})

	assert.True(t, reflect.DeepEqual(expectedAges, ages))
}

type Int int

func (i Int) Compare(other Comparable) int {
	if other, ok := other.(Int); ok {
		return int(i) - int(other)
	}
	return -1
}

type Person struct {
	age int
}

func (m Person) Compare(other Comparable) int {
	if other, ok := other.(Person); ok {
		return m.age - other.age
	}
	return -1
}
