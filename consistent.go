package main

import (
	"errors"
	"hash/crc32"
	"hash/fnv"
	"sort"
	"strconv"
	"sync"
)

type uints []uint32
type Nodes map[uint32]*Node

// Len returns the length of the uints array.
func (x uints) Len() int { return len(x) }

// Less returns true if element i is less than element j.
func (x uints) Less(i, j int) bool { return x[i] < x[j] }

// Swap exchanges elements i and j.
func (x uints) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

// ErrEmptyCircle is the error returned when trying to get an element when nothing has been added to hash.
var ErrEmptyCircle = errors.New("empty circle")

type Node struct {
	id   uint32
	addr string
}

type Ring struct {
	Nodes Nodes
}

func New() *consistent {
	c := new(consistent)
	c.NumberOfReplicas = 20
	c.Ring.Nodes = map[uint32]*Node{}
	return c
}

type consistent struct {
	Ring             Ring
	sortedHashes     uints
	NumberOfReplicas int
	UseFnv           bool
	sync.RWMutex
}

func (c *consistent) AddNode(addr string) {
	c.Lock()
	defer c.Unlock()
	for i := 0; i < c.NumberOfReplicas; i++ {
		id := c.hashKey(addr + strconv.Itoa(i))
		c.Ring.Nodes[id] = NewNode(addr, id)
		// update sorted hash
		c.updateSortedHashes(id)

	}
}

func (c *consistent) updateSortedHashes(id uint32) {
	c.sortedHashes = append(c.sortedHashes, id)
	sort.Sort(c.sortedHashes)
}

func (c *consistent) RemoveNode(addr string) {
	c.Lock()
	defer c.Unlock()
	for i := 0; i < c.NumberOfReplicas; i++ {
		id := c.hashKey(addr + strconv.Itoa(i))
		delete(c.Ring.Nodes, id)
		idx := c.searchEquality(id)
		c.sortedHashes = append(c.sortedHashes[:idx], c.sortedHashes[idx+1:]...)
	}
}

func (c *consistent) Get(key string) (*Node, error) {
	if len(c.sortedHashes) == 0 {
		return nil, ErrEmptyCircle
	}
	hashKey := c.hashKey(key)
	k := c.search(hashKey)
	l := c.sortedHashes[k]
	node := c.Ring.Nodes[l]
	if node == nil {
		return nil, errors.New("invalid key ")
	}
	return node, nil
}

func (c *consistent) searchEquality(key uint32) (i int) {
	f := func(x int) bool {
		return c.sortedHashes[x] == key
	}
	i = sort.Search(len(c.sortedHashes), f)
	if i >= len(c.sortedHashes) {
		i = 0
	}
	return
}

func (c *consistent) search(key uint32) (i int) {
	f := func(x int) bool {
		return c.sortedHashes[x] > key
	}
	i = sort.Search(len(c.sortedHashes), f)
	if i >= len(c.sortedHashes) {
		i = 0
	}
	return
}

func NewNode(addr string, id uint32) *Node {
	return &Node{
		id:   id,
		addr: addr,
	}
}

func (c *consistent) hashKey(key string) uint32 {
	if c.UseFnv {
		return c.hashKeyFnv(key)
	}
	return c.hashKeyCRC32(key)
}

func (c *consistent) hashKeyFnv(key string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return h.Sum32()
}

func (c *consistent) hashKeyCRC32(key string) uint32 {
	if len(key) < 64 {
		var scratch [64]byte
		copy(scratch[:], key)
		return crc32.ChecksumIEEE(scratch[:len(key)])
	}
	return crc32.ChecksumIEEE([]byte(key))
}
