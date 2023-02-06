package ssz

import (
	"fmt"
)

var _ HashWalker = (*walker)(nil)

func Walk(obj HashRoot) {
	w := &walker{
		compute: NewHasher(),
		depth:   -1,
	}
	obj.HashTreeRootWith(w)

	for i, f := range w.fields {
		fmt.Println("==>", i, f)
	}
}

type walker struct {
	// hasher for sub-elements
	compute *Hasher

	// list of computed fields
	fields [][]byte

	// current depth of the hash computation
	depth int64
}

func (w *walker) Hash() []byte {
	return nil
}

func (w *walker) AppendUint8(i uint8) {
	if w.depth != 0 {
		w.compute.AppendUint8(i)
		return
	}

	panic("TODO")
}

func (w *walker) AppendUint64(i uint64) {
	if w.depth != 0 {
		w.compute.AppendUint64(i)
		return
	}
	panic("TODO")
}

func (w *walker) AppendBytes32(b []byte) {
	if w.depth != 0 {
		w.compute.AppendBytes32(b)
		return
	}
	panic("TODO")
}

func (w *walker) PutUint64(i uint64) {
	w.compute.PutUint64(i)

	if w.depth == 0 {
		w.fields = append(w.fields, w.compute.Hash())
	}
}

func (w *walker) PutUint32(i uint32) {
	if w.depth != 0 {
		w.compute.PutUint32(i)
		return
	}
	panic("TODO")
}

func (w *walker) PutUint16(i uint16) {
	if w.depth != 0 {
		w.compute.PutUint16(i)
		return
	}
	panic("TODO")
}

func (w *walker) PutUint8(i uint8) {
	if w.depth != 0 {
		w.compute.PutUint8(i)
		return
	}
	panic("TODO")
}

func (w *walker) FillUpTo32() {
	if w.depth != 0 {
		w.compute.FillUpTo32()
		return
	}
	panic("TODO")
}

func (w *walker) Append(i []byte) {
	if w.depth != 0 {
		w.compute.Append(i)
		return
	}
	panic("TODO")
}

func (w *walker) PutBitlist(bb []byte, maxSize uint64) {
	if w.depth != 0 {
		w.compute.PutBitlist(bb, maxSize)
		return
	}
	panic("TODO")
}

func (w *walker) PutBool(b bool) {
	if w.depth != 0 {
		w.compute.PutBool(b)
		return
	}
	panic("TODO")
}

func (w *walker) PutBytes(b []byte) {
	w.compute.PutBytes(b)

	if w.depth == 0 {
		w.fields = append(w.fields, w.compute.Hash())
	}
}

func (w *walker) Index() int {
	var indx int
	if w.depth >= 0 {
		// sub-items
		indx = w.compute.Index()
	}

	w.depth++
	return indx
}

func (w *walker) Merkleize(indx int) {
	if w.depth == 0 {
		// the main object ends with this merkleize call
		return
	}

	// merkleize a sub-item and gather the fields
	w.depth--
	w.compute.Merkleize(indx)

	if w.depth == 0 {
		w.fields = append(w.fields, w.compute.Hash())
	}
}

func (w *walker) MerkleizeWithMixin(indx int, num, limit uint64) {
	// merkleize a sub-item and gather the fields
	w.depth--
	w.compute.MerkleizeWithMixin(indx, num, limit)

	if w.depth == 0 {
		w.fields = append(w.fields, w.compute.Hash())
	}
}
