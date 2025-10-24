package main

type COWBuffer struct {
	data []byte
	refs *int
	// need to implement
}

func NewCOWBuffer(data []byte) COWBuffer {
	return COWBuffer{} // need to implement
}

func (b *COWBuffer) Clone() COWBuffer {
	return COWBuffer{} // need to implement
}

func (b *COWBuffer) Close() {
	// need to implement
}

func (b *COWBuffer) Update(index int, value byte) bool {
	return false // need to implement
}

func (b *COWBuffer) String() string {
	return "" // need to implement
}
