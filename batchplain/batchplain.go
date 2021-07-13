package batchplain

import (
	"github.com/69b33ebd29f/gipa-hyper/gipa"
	"github.com/69b33ebd29f/mcl-wrapper"
)

type Proof struct {
	T         mcl.GT
	GipaProof gipa.Proof
}
