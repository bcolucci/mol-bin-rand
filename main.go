package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Path struct {
	values []int
}

type Map struct {
	w, h	  int
	paths 	[]Path
}

func NewMap(w int) *Map {
	return &Map{w:w, h:w+1}
}

func (m *Map) BuildPaths(nbPaths int) {
	m.paths = make([]Path, nbPaths)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i:= 0; i < nbPaths; i += 1 {
		p := Path{values:make([]int, m.w)}
		p.values[0] = int(m.h/2) * m.w
		for i := 1; i < m.w; i += 1 {
			if r.Intn(2) == 0 {
				p.values[i] = p.values[i-1] - m.w + 1
			} else {
				p.values[i] = p.values[i-1] + m.w + 1
			}
		}
		m.paths[i] = p
	}
}

func (m *Map) Print() {
	l := m.w*m.h
	var found bool
	for i := 0; i < l; i += 1 {
		if i%m.w == 0 {
			fmt.Printf("\n")
		}
		found = false
		Loop:
			for _, p := range m.paths {
				for _, k := range p.values {
					if k == i {
						found = true
						break Loop
					}
				}
			}
		if found {
			fmt.Print("[X]")
		} else {
			fmt.Print("[ ]")
		}
	}
}

func main() {
	m := NewMap(66)
	m.BuildPaths(10)
	m.Print()
}
