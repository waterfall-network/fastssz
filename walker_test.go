package ssz

import (
	"fmt"
	"testing"
)

func TestWalkerXXX(t *testing.T) {
	w := &fieldWalker{
		compute: NewHasher(),
		depth:   -1,
	}

	indx1 := w.Index()

	w.PutUint64(1)

	w.PutBytes([]byte{0x1})

	val := make([]byte, 33)
	w.PutBytes(val)

	indx2 := w.Index()
	indx3 := w.Index()

	w.Merkleize(indx3)
	w.Merkleize(indx2)

	indx2 = w.Index()
	w.Merkleize(indx2)

	w.Merkleize(indx1)

	fmt.Println(w.fields)
}
