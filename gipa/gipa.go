package gipa

import (
	"github.com/69b33ebd29f/gipa-hyper/cm"
	"github.com/69b33ebd29f/mcl-wrapper"
)

type Proof struct {
	L []cm.Com
	R []cm.Com
	A [1]mcl.G1
	B [1]mcl.G2
}

func (self *Proof) Append(ComL cm.Com, ComR cm.Com) {

	self.L = append(self.L, ComL)
	self.R = append(self.R, ComR)
}

func (self *Proof) At(i uint64) (cm.Com, cm.Com) {

	return self.L[i], self.R[i]
}
