package energy

import (
	"testing"
	"time"
)

var defaultEnergy struct {
	max      int
	interval time.Duration
}

func init() {
	defaultEnergy.max = 10
	defaultEnergy.interval = time.Second
}

func TestNew(t *testing.T) {

	// TODO(Liru): Add more test cases for TestNew.
	var tests = []struct {
		initEnergy int
		maxEnergy  int
	}{
		{10, 10},
		{15, 20},
	}

	for i, test := range tests {
		e := New(test.initEnergy, test.maxEnergy, defaultEnergy.interval)
		currentEnergy := e.CurrentEnergy()
		if currentEnergy != test.initEnergy {
			t.Errorf("#%d: New(%d,%d): expected %d for currentEnergy, got %d",
				i, test.initEnergy, test.maxEnergy,
				test.initEnergy, currentEnergy)
		}
	}
}

func TestUse(t *testing.T) {
	var tests = []struct {
		initEnergy int
		expected   int
	}{
		{10, 9},
		{5, 4},
		{0, 0},
		{-5, -5},
	}

	for i, test := range tests {
		e := New(test.initEnergy, defaultEnergy.max, defaultEnergy.interval)
		e.Use()
		current := e.CurrentEnergy()
		if current != test.expected {
			t.Errorf("#%d: Use(): expected %d for currentEnergy, got %d",
				i, test.expected, current,
			)
		}

	}
}

func TestSetEnergy(t *testing.T) {
	var tests = []struct {
		set int
	}{
		{0}, {1}, {2}, {3}, {4}, {5},
		{-1}, {-2}, {-3}, {-4}, {-5},
		{10}, {11}, {12}, {13}, {14}, {15},
	}

	for i, test := range tests {
		e := New(defaultEnergy.max, defaultEnergy.max, defaultEnergy.interval)
		e.SetEnergy(test.set)
		current := e.CurrentEnergy()
		if current != test.set {
			t.Errorf("#%d: SetEnergy(%d): expected %d for CurrentEnergy, got %d",
				i, test.set, test.set, current,
			)
		}
	}
}

func TestResetEnergy(t *testing.T) {
	var tests = []struct {
		set int
		max int
	}{
		{0, 10},
		{1, 10},
		{9, 10},
		{10, 10},
		{11, 10},
		{100, 10},
		{0, 1000},
		{1000, 10},
		{1000000, 1000},
	}

	for i, test := range tests {
		e := New(test.set, test.max, defaultEnergy.interval)
		e.ResetEnergy()
		current := e.CurrentEnergy()
		if current != test.max {
			t.Errorf("#%d: ResetEnergy(): expected %d for CurrentEnergy, got %d",
				i, test.max, current,
			)
		}
	}
}

func TestRecoverEnergy(t *testing.T) {
	t.Skip("test not written yet")
}

func TestUseEnergyWhileRecovering(t *testing.T) {
	t.Skip("test not written yet")
}

func TestSetMaxEnergy(t *testing.T) {
	t.Skip("test not written yet")
}

func TestExtraEnergy(t *testing.T) {
	t.Skip("test not written yet")
}

func TestRecoveryQuantity(t *testing.T) {
	t.Skip("test not written yet")
}

func TestFloatRecoveryInterval(t *testing.T) {
	t.Skip("test not supported yet: must switch e.interval to float first")
}

// TestExtraEnergyRecovery is to see if energy will recover after setting
// it over the max limit and then using it. It shouldn't recover at that point.
func TestExtraEnergyRecovery(t *testing.T) {
	t.Skip("test not written yet")
}

// ===== Additional stuff =====

func TestString(t *testing.T) {
	t.Skip("test not written yet")
}

func TestJSON(t *testing.T) {
	t.Skip("test not written yet")
}

func TestEnergySort(t *testing.T) {
	t.Skip("test not written yet")
}
