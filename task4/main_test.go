package main

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// go test -v main_test.go

func TestOrderedMap(t *testing.T) {
	data := NewOrderedMap[int, int]()
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

	var keys []int
	expectedKeys := []int{2, 4, 5, 10, 12, 14, 15}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))

	assert.True(t, data.Erase(15))
	assert.True(t, data.Erase(14))
	assert.True(t, data.Erase(2))
	assert.False(t, data.Erase(42))

	assert.Equal(t, 4, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(2))
	assert.False(t, data.Contains(14))

	keys = nil
	expectedKeys = []int{4, 5, 10, 12}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))
}

func TestOrderedMap_Insert(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		expRoot   *Node[int, int]
		keysToAdd []int
	}{
		"insert to empty tree": {
			expRoot: &Node[int, int]{
				left: &Node[int, int]{
					left:  nil,
					right: nil,
					key:   1,
					value: 2,
				},
				right: &Node[int, int]{
					left: &Node[int, int]{
						key:   9,
						value: 4,
						left:  nil,
						right: nil,
					},
					right: &Node[int, int]{
						left:  nil,
						right: nil,
						key:   13,
						value: 3,
					},
					key:   12,
					value: 1,
				},
				key:   2,
				value: 0,
			},
			keysToAdd: []int{2, 12, 1, 13, 9},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			data := NewOrderedMap[int, int]()

			for i, key := range test.keysToAdd {
				data.Insert(key, i)
			}

			assert.True(t, reflect.DeepEqual(test.expRoot, data.root))

		})
	}
}

func TestOrderedMap_Erase(t *testing.T) {
	t.Parallel()

	nilNode := &Node[int, int]{
		left:  nil,
		right: nil,
		key:   2,
		value: 0,
	}

	tests := map[string]struct {
		expRemoved       bool
		expRoot          *Node[int, int]
		removedLeftChild bool
		keysToAdd        []int
		keyToRemove      int
	}{
		"erase root that equals nil": {
			expRemoved:       false,
			expRoot:          nil,
			removedLeftChild: false,
			keysToAdd:        []int{},
			keyToRemove:      42,
		},
		"erase root node": {
			expRemoved:       true,
			expRoot:          nil,
			removedLeftChild: false,
			keysToAdd:        []int{2},
			keyToRemove:      2,
		},
		"erase unknown node": {
			expRemoved: false,
			expRoot: &Node[int, int]{
				left: nil,
				right: &Node[int, int]{
					left: nil,
					right: &Node[int, int]{
						left:  nil,
						right: nil,
						key:   252,
						value: 2,
					},
					key:   152,
					value: 1,
				},
				key:   52,
				value: 0,
			},
			removedLeftChild: false,
			keysToAdd:        []int{52, 152, 252},
			keyToRemove:      42,
		},
		"erase root right (children nils)": {
			expRemoved:       true,
			expRoot:          nilNode,
			removedLeftChild: false,
			keysToAdd:        []int{2, 12},
			keyToRemove:      12,
		},
		"erase root left (children nils)": {
			expRemoved:       true,
			expRoot:          nilNode,
			removedLeftChild: true,
			keysToAdd:        []int{2, 1},
			keyToRemove:      1,
		},
		"erase root left (right children == nil)": {
			expRemoved: true,
			expRoot: &Node[int, int]{
				left:  &Node[int, int]{left: &Node[int, int]{left: nil, right: nil, key: 1, value: 3}, right: nil, key: 3, value: 2},
				right: nil,
				key:   33,
				value: 0,
			},
			removedLeftChild: true,
			keysToAdd:        []int{33, 5, 3, 1},
			keyToRemove:      5,
		},
		"erase root left (left children == nil)": {
			expRemoved: true,
			expRoot: &Node[int, int]{
				left: &Node[int, int]{
					left: nil,
					right: &Node[int, int]{
						left:  nil,
						right: nil,
						key:   24,
						value: 3,
					},
					key:   22,
					value: 2,
				},
				right: nil,
				key:   33,
				value: 0,
			},
			removedLeftChild: true,
			keysToAdd:        []int{33, 5, 22, 24},
			keyToRemove:      5,
		},
		"erase root left (both children != nil)": {
			expRemoved: true,
			expRoot: &Node[int, int]{
				left: &Node[int, int]{
					left: &Node[int, int]{left: &Node[int, int]{
						left: nil, right: nil, key: 1, value: 4,
					},
						right: &Node[int, int]{
							left: nil, right: nil, key: 4, value: 5,
						},
						key:   3,
						value: 2,
					},
					right: &Node[int, int]{
						left:  nil,
						right: &Node[int, int]{left: nil, right: nil, key: 24, value: 7},
						key:   22,
						value: 3,
					},
					key:   11,
					value: 6,
				},
				right: nil,
				key:   33,
				value: 0,
			},
			removedLeftChild: true,
			keysToAdd:        []int{33, 5, 3, 22, 1, 4, 11, 24},
			keyToRemove:      5,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			om := NewOrderedMap[int, int]()

			for i, key := range test.keysToAdd {
				om.Insert(key, i)
			}

			removed := om.Erase(test.keyToRemove)

			assert.Equal(t, removed, test.expRemoved)
			assert.True(t, reflect.DeepEqual(test.expRoot, om.root))

			if test.expRemoved && om.root != nil {
				if test.removedLeftChild {
					assert.True(t, reflect.DeepEqual(test.expRoot.left, om.root.left))
				} else {
					assert.True(t, reflect.DeepEqual(test.expRoot.right, om.root.right))
				}
			}
		})
	}
}

func TestOrderedMap_ForEach(t *testing.T) {

}

func TestOrderedMap_Size(t *testing.T) {

}

func TestOrderedMap_Contains(t *testing.T) {

}
