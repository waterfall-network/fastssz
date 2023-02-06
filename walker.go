package ssz

var _ HashWalker = (*fieldWalker)(nil)

func Walk(obj HashRoot, h *Hasher) ([]byte, error) {
	w := &fieldWalker{
		compute: h,
		depth:   -1,
	}
	if err := obj.HashTreeRootWith(w); err != nil {
		return nil, err
	}

	// TODO: Generate the intermediate tree to generate the proofs
	root := w.compute.merkleizeImpl(nil, w.fields, 0)
	return root, nil
}

type fieldWalker struct {
	// hasher for sub-elements
	compute *Hasher

	// list of computed fields
	fields []byte

	// current depth of the hash computation
	depth int64
}

func (w *fieldWalker) Hash() []byte {
	return nil
}

func (w *fieldWalker) complete() {
	if w.depth == 0 {
		w.fields = append(w.fields, w.compute.Hash()...)
	}
}

func (w *fieldWalker) AppendUint8(i uint8) {
	w.compute.AppendUint8(i)
	w.complete()
}

func (w *fieldWalker) AppendUint64(i uint64) {
	w.compute.AppendUint64(i)
	w.complete()
}

func (w *fieldWalker) AppendBytes32(b []byte) {
	w.compute.AppendBytes32(b)
	w.complete()
}

func (w *fieldWalker) PutUint64(i uint64) {
	w.compute.PutUint64(i)
	w.complete()
}

func (w *fieldWalker) PutUint32(i uint32) {
	w.compute.PutUint32(i)
	w.complete()
}

func (w *fieldWalker) PutUint16(i uint16) {
	w.compute.PutUint16(i)
	w.complete()
}

func (w *fieldWalker) PutUint8(i uint8) {
	w.compute.PutUint8(i)
	w.complete()
}

func (w *fieldWalker) FillUpTo32() {
	w.compute.FillUpTo32()
	// FillUpTo32 is part of Merkleize calls, do not need to complete
}

func (w *fieldWalker) Append(i []byte) {
	w.compute.Append(i)
	// Append is part of Merkleize calls, do not need to complete
}

func (w *fieldWalker) PutBitlist(bb []byte, maxSize uint64) {
	w.compute.PutBitlist(bb, maxSize)
	w.complete()
}

func (w *fieldWalker) PutBool(b bool) {
	w.compute.PutBool(b)
	w.complete()
}

func (w *fieldWalker) PutBytes(b []byte) {
	w.compute.PutBytes(b)
	w.complete()
}

func (w *fieldWalker) Index() int {
	var indx int
	if w.depth >= 0 {
		// sub-items
		indx = w.compute.Index()
	}

	w.depth++
	return indx
}

func (w *fieldWalker) Merkleize(indx int) {
	if w.depth == 0 {
		// the main object ends with this merkleize call
		return
	}

	// merkleize a sub-item and gather the fields
	w.depth--
	w.compute.Merkleize(indx)

	w.complete()
}

func (w *fieldWalker) MerkleizeWithMixin(indx int, num, limit uint64) {
	// merkleize a sub-item and gather the fields
	w.depth--
	w.compute.MerkleizeWithMixin(indx, num, limit)

	w.complete()
}
