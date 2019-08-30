package state

import (
	"github.com/phyrwork/goadvent/day/oneeight/reservoir/display"
	"github.com/phyrwork/goadvent/vector"
	"log"
)

type Vector2 [2]int

func NewVector2(d ...int) Vector2 {
	x, y := display.XY(d)
	return Vector2{x, y}
}

type Map map[Vector2]Tile

func (m Map) Get(v vector.Vector) Tile { return m[NewVector2(v...)] }

func (m Map) Set(v vector.Vector, t Tile) { m[NewVector2(v...)] = t }

func (m Map) Used() []vector.Vector {
	u := make([]vector.Vector, 0, len(m))
	for v := range m {
		u = append(u, v[:])
	}
	return u
}

func (m Map) Range() *vector.Range {
	u := m.Used()
	if len(u) == 0 {
		return nil
	}
	sx, sy := display.XY(u[0])
	ex, ey := sx, sy
	for _, v := range u {
		vx, vy := display.XY(v)
		switch {
		case vx < sx: sx = vx
		case vx > ex: ex = vx
		}
		switch {
		case vy < sy: sy = vy
		case vy > ey: ey = vy
		}
	}
	s, e := vector.Vector{sx, sy}, vector.Vector{ex, ey}
	r := vector.NewRange(s, e)
	return &r
}

func (m Map) Display() display.Offset {
	r := m.Range()
	w, h := display.XY(r.Size())
	g := display.NewGrid(w, h)
	d := display.Offset{g, r.Start()}
	_ = r.Each(func (v vector.Vector) error {
		t := m.Get(v)
		if t == nil {
			t = Sand(Default)
		}
		if err := d.Set(v[:], t.Draw()); err != nil {
			log.Panicf("map display error: %v", err)
		}
		return nil
	})
	return d
}

type State struct {
	Map
}

func (s State) Get(v vector.Vector) Tile {
	t := s.Map.Get(v)
	if t == nil {
		t = Sand(Default)
		s.Set(v, t)
	}
	return t
}