package batch

import (
	"github.com/69b33ebd29f/gipa-hyper/gipakzg"
	"github.com/69b33ebd29f/mcl-wrapper"
)

type Proof struct {
	T            mcl.GT
	GipaKzgProof gipakzg.Proof
}
