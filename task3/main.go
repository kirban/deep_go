package main

import (
	"unsafe"
)

/*
Идея подхода copy-on-write заключается в том, что при чтении данных
используется общая копия данных буффера, но в случае изменения данных
— создается новая копия данных буффера.
Для реализации такого подхода можно использовать разделяемый счетчик ссылок
- если при изменении данных буффера кто-то еще ссылается на этот буффер,
то нужно будет сначала произвести копию данных буффера,
изменить счетчик ссылок и только затем произвести изменение
(*если никто не ссылается на буффер, то копировать данные буффера не нужно при изменении данных*).

Дополнительно еще нужно реализовать метод конвертации данных буффера в строку без копирования
и дополнительного выделения памяти.
*/

type COWBuffer struct {
	data []byte
	refs *int
	// need to implement
}

// NewCOWBuffer - создать буффер с определенными данными
func NewCOWBuffer(data []byte) COWBuffer {
	refs := new(int)
	*refs = 1

	return COWBuffer{
		data: data,
		refs: refs,
	}
}

// Clone - создать новую копию буфера
func (b *COWBuffer) Clone() COWBuffer {
	*b.refs++
	return COWBuffer{
		data: b.data,
		refs: b.refs,
	}
}

// Close - перестать использовать копию буффера
func (b *COWBuffer) Close() {
	if *b.refs > 0 {
		*b.refs--
	}

	b.data = nil
}

// Update - изменить определенный байт в буффере
func (b *COWBuffer) Update(index int, value byte) bool {
	if index < 0 || index >= len(b.data) {
		return false
	}

	if *(b.refs) > 1 {
		newData := make([]byte, len(b.data))
		copy(newData, b.data)
		b.data = newData
		*b.refs--
	}

	b.data[index] = value

	return true
}

// String - сконвертировать буффер в строку
func (b *COWBuffer) String() string {
	return unsafe.String(unsafe.SliceData(b.data), len(b.data))
}

func main() {

}
