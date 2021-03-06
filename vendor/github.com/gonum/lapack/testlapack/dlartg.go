// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testlapack

import (
	"math"
	"math/rand"
	"testing"
)

type Dlartger interface {
	Dlartg(f, g float64) (cs, sn, r float64)
}

func DlartgTest(t *testing.T, impl Dlartger) {
	rnd := rand.New(rand.NewSource(1))
	for i := 0; i < 100; i++ {
		f := rnd.NormFloat64()
		g := rnd.NormFloat64()
		cs, sn, r := impl.Dlartg(f, g)

		rTest := cs*f + sn*g
		zeroTest := -sn*f + cs*g
		if math.Abs(rTest-r) > 1e-14 {
			t.Errorf("R result mismatch. Computed %v, found %v", r, rTest)
		}
		if math.Abs(zeroTest) > 1e-14 {
			t.Errorf("Zero result mismatch. Found %v", zeroTest)
		}
		if math.Abs(f) > math.Abs(g) && cs < 0 {
			t.Errorf("Unexpected negative cs %v", cs)
		}
	}
	for i := 0; i < 100; i++ {
		cs, sn, _ := impl.Dlartg(rnd.NormFloat64(), 0)
		if cs != 1 {
			t.Errorf("Unexpected cs for g=0. Want 1, got %v", cs)
		}
		if sn != 0 {
			t.Errorf("Unexpected sn for g=0. Want 0, got %v", sn)
		}
	}
	for i := 0; i < 100; i++ {
		cs, sn, _ := impl.Dlartg(0, rnd.NormFloat64())
		if cs != 0 {
			t.Errorf("Unexpected cs for f=0. Want 0, got %v", cs)
		}
		if sn != 1 {
			t.Errorf("Unexpected sn for f=0. Want 1, got %v", sn)
		}
	}
}
