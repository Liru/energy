// Package energy provides a concurrent energy system, useful for games and other applications.
package energy

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// ErrMaxed is returned when stuff happens.
var ErrMaxed = errors.New("energy is at or over maximum")

// Energy is a consumable and recoverable resource in games.
type Energy struct {
	mtx sync.RWMutex

	usedEnergy int
	usedAt     time.Time
	max        int

	recoveryInterval time.Duration
	recoveryQuantity int
}

// XXX(Liru): Implement these, you idiot https://golang.org/pkg/bytes/

// New creates and returns a new Energy instance.
func New(initEnergy int, maxEnergy int, interval time.Duration) *Energy {

	used := maxEnergy - initEnergy
	var usedAt time.Time
	// if 0 < used {
	// 	usedAt = time.Now().Add(-interval * time.Duration(used))
	// }
	usedAt = time.Now()

	return &Energy{
		max:        maxEnergy,
		usedEnergy: used,

		usedAt:           usedAt,
		recoveryInterval: interval,
		// recoveryQuantity: quantity,
	}
}

// CurrentEnergy returns the current energy available for use.
func (e *Energy) CurrentEnergy() int {
	// if e.usedAt.IsZero() {
	// 	return e.max
	// }
	return e.max - e.usedEnergy + e.recovered()
}

// Use uses a single instance of energy.
func (e *Energy) Use() bool {
	e.mtx.Lock()
	defer e.mtx.Unlock()

	return e.use(1, false)
}

// UseEnergy uses the specified amount of energy.
func (e *Energy) UseEnergy(i int) bool {
	e.mtx.Lock()
	defer e.mtx.Unlock()

	return e.use(i, false)
}

// RecoversIn calculates when the next instance of energy will be recovered.
//
// If energy cannot be recovered, RecoversIn returns 0 and an error
// explaining why it cannot be recovered.
func (e *Energy) RecoversIn() time.Duration {
	if e.usedAt.IsZero() {
		return 0
	}

	p := e.passed()

	ticks := int(p / e.recoveryInterval)
	if ticks >= e.usedEnergy {
		return 0
	}

	return e.recoveryInterval - (p % e.recoveryInterval)

}

// FullyRecoversIn calculates when the energy instance's capacity will be reached.
//
// If energy cannot be recovered, RecoversIn returns 0 and an error
// explaining why it cannot be recovered.
func (e *Energy) FullyRecoversIn() time.Duration {
	nextRecover := e.RecoversIn()
	if nextRecover == 0 {
		return 0
	}

	ttr := e.max - e.CurrentEnergy() - 1
	return nextRecover + e.recoveryInterval*time.Duration(ttr)
}

// String satisfies the fmt.Stringer interface.
func (e *Energy) String() string {
	s := fmt.Sprintf("<Energy %d/%d", e.CurrentEnergy(), e.max)
	if e.CurrentEnergy() < e.max {
		//TODO: get recovery time and input below
		nextRecover := e.RecoversIn()
		mins, secs := int(nextRecover.Minutes()), int(nextRecover.Seconds())%60
		s += fmt.Sprintf(" recover in %02d:%02d", mins, secs)
	}
	s += ">"

	return s
}

// SetEnergy sets the current energy to the given value.
//
// It can be used to set the energy over the given maximum.
func (e *Energy) SetEnergy(i int) {
	e.mtx.Lock()
	defer e.mtx.Unlock()

	if i >= e.max {
		e.usedEnergy = e.max - i
		e.usedAt = time.Time{}
	} else {
		e.use(e.CurrentEnergy()-i, true)
	}
}

// ResetEnergy resets the energy to its current maximum.
func (e *Energy) ResetEnergy() {
	e.SetEnergy(e.max)
}

// SetMax sets the maximum energy to the parameter.
func (e *Energy) SetMax(i int) {
	e.mtx.Lock()
	defer e.mtx.Unlock()

	if i < e.CurrentEnergy() {
		// if e.RecoversIn() != 0 {
		// 	e.usedEnergy += i - e.max
		// }
		e.usedAt = time.Time{}
		e.usedEnergy = 0
	}
	e.max = i
}

// SetInterval sets the interval at which energy is regained.
func (e *Energy) SetInterval(i time.Duration) {
	e.mtx.Lock()
	defer e.mtx.Unlock()

	e.recoveryInterval = i
}

// ===== Helper functions ahoy.  =====

// passed indicates how much time has passed since the last time this energy was used.
func (e *Energy) passed() time.Duration {
	if e.usedAt.IsZero() {
		return 0
	}
	return time.Since(e.usedAt)
}

// recovered indicates how much energy has been recovered since the last use.
func (e *Energy) recovered() int {
	p := e.passed()
	if p == 0 {
		return 0
	}

	intervals := int(p / e.recoveryInterval)
	rec := intervals * e.recoveryQuantity

	if rec > e.usedEnergy {
		return e.usedEnergy
	}
	return rec
}

// use is a backend helper that is wrapped by Use and UseEnergy.
func (e *Energy) use(i int, force bool) bool {
	if e.CurrentEnergy() < i && !force {
		return false
	}

	if ((e.CurrentEnergy()-i < e.max) && (e.max <= e.CurrentEnergy())) || force {
		e.usedEnergy = i - e.CurrentEnergy() + e.max
		e.usedAt = time.Now()
	} else {
		e.usedEnergy = e.max - e.CurrentEnergy() + i
	}

	return true
}
