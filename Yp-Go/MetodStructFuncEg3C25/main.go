package main

import (
	"fmt"
)

// Circular buffer реализует структуру кольцевой буфер для значений float64
type CircularBuffer struct {
	values  []float64 //текущие значения буфера
	headIdx int       // индекс головы (первый непустой элемент)
	tailIdx int       //индекс хвоста (первый пустой элемент)
}

// GetCurrentSize возращает текущею длину буфера
func (b CircularBuffer) GetCurrentSize() int {

	if b.tailIdx < b.headIdx {
		return b.tailIdx + cap(b.values) - b.headIdx
	}

	return b.tailIdx - b.headIdx
}

// GetValues возвращает слайс текущих занчений буфера, сохраняя порядок записи
func (b CircularBuffer) GetValues() (retValues []float64) {

	for i := b.headIdx; i != b.tailIdx; i = (i + 1) % cap(b.values) {
		retValues = append(retValues, b.values[i])
	}

	return
}

// AddValue добавляет новое значение в буфер
func (b *CircularBuffer) AddValue(v float64) {

	b.values[b.tailIdx] = v
	b.tailIdx = (b.tailIdx + 1) % cap(b.values)
	if b.tailIdx == b.headIdx {
		b.headIdx = (b.headIdx + 1) % cap(b.values)
	}
}

// NewCircularBuffer - конструктор типа CircularBuffer
func NewCircularBuffer(size int) CircularBuffer {
	return CircularBuffer{values: make([]float64, size+1)}
}

func main() {

	fmt.Println("Hello")
	//вызов метода
	buf := NewCircularBuffer(4)
	for i := 0; i < 6; i++ {
		if i > 0 {
			buf.AddValue(float64(i))
		}
		fmt.Printf("[%d]: %v\n", buf.GetCurrentSize(), buf.GetValues())
	}

}
