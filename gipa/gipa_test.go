package gipa

import (
	"testing"

	"github.com/69b33ebd29f/gipa-hyper/utils"
)

func TestGIPA(t *testing.T) {

	M := uint64(1) << 10
	alpha, beta, g, h := utils.RunMPC()
	prover, verifier := GipaTestSetup(M, alpha, beta, g, h)
	proof := prover.Prove()
	status := verifier.Verify(proof)
	if status == false {
		t.Errorf("GIPA Test: Failed")
	}
}
