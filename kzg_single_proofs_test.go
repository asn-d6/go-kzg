// +build !bignum_pure,!bignum_hol256

package kzg

import (
	"github.com/protolambda/go-kzg/bls"
	"testing"
)

func TestKZGSettings_CheckProofSingle(t *testing.T) {
	fs := NewFFTSettings(4)
	s1, s2 := generateSetup("1927409816240961209460912649124", 16+1)
	ks := NewKZGSettings(fs, s1, s2)
	for i := 0; i < len(ks.secretG1); i++ {
		t.Logf("secret g1 %d: %s", i, bls.StrG1(&ks.secretG1[i]))
	}

	polynomial := testPoly(1, 2, 3, 4, 7, 7, 7, 7, 13, 13, 13, 13, 13, 13, 13, 13)
	for i := 0; i < len(polynomial); i++ {
		t.Logf("poly coeff %d: %s", i, bls.BigStr(&polynomial[i]))
	}

	commitment := ks.CommitToPoly(polynomial)
	t.Log("commitment\n", bls.StrG1(commitment))

	proof := ks.ComputeProofSingle(polynomial, 17)
	t.Log("proof\n", bls.StrG1(proof))

	var x bls.Big
	bls.AsBig(&x, 17)
	var value bls.Big
	bls.EvalPolyAt(&value, polynomial, &x)
	t.Log("value\n", bls.BigStr(&value))

	if !ks.CheckProofSingle(commitment, proof, &x, &value) {
		t.Fatal("could not verify proof")
	}
}

func testPoly(polynomial ...uint64) []bls.Big {
	n := len(polynomial)
	polynomialBig := make([]bls.Big, n, n)
	for i := 0; i < n; i++ {
		bls.AsBig(&polynomialBig[i], polynomial[i])
	}
	return polynomialBig
}

func generateSetup(secret string, n uint64) ([]bls.G1, []bls.G2) {
	var s bls.Big
	bls.BigNum(&s, secret)

	var sPow bls.Big
	bls.CopyBigNum(&sPow, &bls.ONE)

	s1Out := make([]bls.G1, n, n)
	s2Out := make([]bls.G2, n, n)
	for i := uint64(0); i < n; i++ {
		bls.MulG1(&s1Out[i], &bls.GenG1, &sPow)
		bls.MulG2(&s2Out[i], &bls.GenG2, &sPow)
		var tmp bls.Big
		bls.CopyBigNum(&tmp, &sPow)
		bls.MulModBig(&sPow, &tmp, &s)
	}
	return s1Out, s2Out
}
