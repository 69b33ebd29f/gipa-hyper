package batchplain

import (
	"fmt"
	"testing"

	"github.com/69b33ebd29f/gipa-hyper/cm"
	"github.com/69b33ebd29f/gipa-hyper/utils"
	"github.com/69b33ebd29f/mcl-wrapper"
)

func testsetup() []uint8 {
	return []uint8{3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
}
func BenchmarkBatchPlain(b *testing.B) {

	folderPath := "../ck-19"
	rows := testsetup()
	M := uint32(30)
	// maxN := (uint32(1) << (rows[len(rows)-1]))            // Not sure why this does not work
	// p, q, a_i, b_i := utils.GenerateBatchingData(M, maxN) // Not sure why this does not work
	// fmt.Println(len(p), len(a_i))
	for _, ell := range rows {
		N := uint32(1) << ell
		MN := utils.NextPowOf2(uint64(N * M))
		nPad, mnPad := utils.ComputePadding(M, N)

		p, q, a_i, b_i := utils.GenerateBatchingData(M, N)

		pPad := make([]mcl.G1, nPad)
		qPad := make([]mcl.G2, nPad)
		P := append(p[:N], pPad...)
		Q := append(q[:N], qPad...)

		aPad := make([]mcl.G1, mnPad)
		bPad := make([]mcl.G2, mnPad)
		A := append(a_i[:(N*M)], aPad...)
		B := append(b_i[:(N*M)], bPad...)

		ck, _, _ := cm.LoadKeys(MN, folderPath)

		prover, verifier := AssembleProverVerifier(M,
			&ck,
			P, Q,
			A, B)
		// fmt.Println(folderPath, M, N, M*N, utils.NextPowOf2(MN), nPad, mnPad)

		var proofs []Proof
		b.Run(fmt.Sprintf("%d/Prove;%d", ell, MN), func(b *testing.B) {
			var proverLocal Prover
			for bn := 0; bn < b.N; bn++ {
				proverLocal = Prover{}
				proverLocal.Clone(&prover)
				b.StartTimer()
				proof := proverLocal.Prove()
				b.StopTimer()
				proofs = append(proofs, proof)
			}
		})

		b.Run(fmt.Sprintf("%d/Verifier;%d", ell, MN), func(b *testing.B) {
			var verifierLocal Verifier
			var proof Proof
			var status bool
			for bn := 0; bn < b.N; bn++ {
				verifierLocal = Verifier{}
				verifierLocal.Clone(&verifier)
				proof, proofs = proofs[0], proofs[1:]
				b.StartTimer()
				status = verifierLocal.Verify(proof)
				b.StopTimer()
				if !status {
					b.Errorf("Batch Verification failed")
				}
			}
		})
	}
}
