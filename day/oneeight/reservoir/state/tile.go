package state

import (
	"github.com/phyrwork/goadvent/day/oneeight/reservoir/display"
	"log"
)

type Tile interface {
	Draw() rune
	Permeable() bool
}

type Source rune
func (t Source) Draw() rune { return display.MarkerSource }
func (t Source) Permeable() bool { return false }

type Clay rune
func (t Clay) Draw() rune { return display.MarkerClay }
func (t Clay) Permeable() bool {return false }

const (
	Dry = iota
	Flow
	Still
	Default = Dry
)

type Sand int

func (t Sand) Draw() rune {
	switch int(t) {
	case Dry:   return display.MarkerSand
	case Flow:  return display.MarkerFlow
	case Still: return display.MarkerStill
	default:    log.Panicf("invalid sand state: %v", t)
	}
	return 0
}

func (t Sand) Permeable() bool {
	switch int(t) {
	case Still: return false
	default:    return true
	}
}