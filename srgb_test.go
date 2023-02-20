package colt

import "testing"

func TestStandardb(t *testing.T) {
	tests := []struct {
		linear   float32
		standard uint8
	}{
		{linear: 1, standard: 255},
		{linear: .999, standard: 255},
		{linear: .99, standard: 254},
		{linear: .75, standard: 225},
		{linear: 0.5, standard: 188},
		{linear: .212, standard: 127},
		{linear: 1.0 / 255, standard: 13},
		{linear: .001, standard: 3},
		{linear: .0003, standard: 1},
		{linear: 0, standard: 0},
	}
	for _, test := range tests {
		var standard = Standardb(test.linear)
		if standard != test.standard {
			t.Errorf("Standardb(%f) returned %d. Expected %d.", test.linear, standard, test.standard)
		}
	}
}

func TestLinearb(t *testing.T) {
	tests := []struct {
		linear   float32
		standard uint8
	}{
		{linear: 1, standard: 255},
		{linear: .991102, standard: 254},
		{linear: .752942, standard: 225},
		{linear: 0.502886, standard: 188},
		{linear: .212231, standard: 127},
		{linear: 0.004025, standard: 13},
		{linear: .000911, standard: 3},
		{linear: .000304, standard: 1},
		{linear: 0, standard: 0},
	}
	for _, test := range tests {
		var linear = Linearb(test.standard)
		if (linear+0.00001)*1.00001 < test.linear || linear > (test.linear+0.00001)*1.00001 {
			t.Errorf("Linearb(%d) returned %f. Expected %f.", test.standard, linear, test.linear)
		}
	}
}
