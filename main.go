package main

import (
	"fmt"
	"math/bits"

	"github.com/69b33ebd29f/gipa-hyper/cm"
	"github.com/69b33ebd29f/gipa-hyper/utils"
	"github.com/69b33ebd29f/mcl-wrapper"
)

func main() {
	fmt.Println("Hello, World!")
	mcl.InitFromString("bls12-381")
	GenKeys()
}

const MAX_AGG_SIZE = 1 << 19

func GenKeys() {
	alpha, beta, G, H := utils.RunMPC()
	mn := uint64(MAX_AGG_SIZE) // short circuiting things
	folderPath := fmt.Sprintf("ck-%02d", bits.Len(MAX_AGG_SIZE)-1)
	ck, kzg1, kzg2 := cm.IPPSetupKZG(mn, alpha, beta, G, H)
	cm.IPPSaveCmKzg(ck, kzg1, kzg2, folderPath)
}
