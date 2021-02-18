// Original: https://github.com/ethereum/research/blob/master/mimc_stark/fft.py

package kzg

import (
	"github.com/protolambda/go-kzg/bls"
	"math/bits"
)

// if not already a power of 2, return the next power of 2
func nextPowOf2(v uint64) uint64 {
	if v == 0 {
		return 1
	}
	return uint64(1) << bits.Len64(v-1)
}

// Expands the power circle for a given root of unity to WIDTH+1 values.
// The first entry will be 1, the last entry will also be 1,
// for convenience when reversing the array (useful for inverses)
func expandRootOfUnity(rootOfUnity *bls.Big) []bls.Big {
	rootz := make([]bls.Big, 2)
	rootz[0] = bls.ONE // some unused number in py code
	rootz[1] = *rootOfUnity
	for i := 1; !bls.EqualOne(&rootz[i]); {
		rootz = append(rootz, bls.Big{})
		this := &rootz[i]
		i++
		bls.MulModBig(&rootz[i], this, rootOfUnity)
	}
	return rootz
}

type FFTSettings struct {
	maxWidth uint64
	// the generator used to get all roots of unity
	rootOfUnity *bls.Big
	// domain, starting and ending with 1 (duplicate!)
	expandedRootsOfUnity []bls.Big
	// reverse domain, same as inverse values of domain. Also starting and ending with 1.
	reverseRootsOfUnity []bls.Big
}

func NewFFTSettings(maxScale uint8) *FFTSettings {
	width := uint64(1) << maxScale
	root := &bls.Scale2RootOfUnity[maxScale]
	rootz := expandRootOfUnity(&bls.Scale2RootOfUnity[maxScale])
	// reverse roots of unity
	rootzReverse := make([]bls.Big, len(rootz), len(rootz))
	copy(rootzReverse, rootz)
	for i, j := uint64(0), uint64(len(rootz)-1); i < j; i, j = i+1, j-1 {
		rootzReverse[i], rootzReverse[j] = rootzReverse[j], rootzReverse[i]
	}

	return &FFTSettings{
		maxWidth:             width,
		rootOfUnity:          root,
		expandedRootsOfUnity: rootz,
		reverseRootsOfUnity:  rootzReverse,
	}
}
