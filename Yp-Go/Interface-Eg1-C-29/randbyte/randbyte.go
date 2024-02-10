package randbyte

import (
	"io"
	"math/rand"
)

type generator struct {

	//Генератор случайных чисел. Вообще rand уже реализует генератор, присвоеный интерфейсу io.Rader, сама структура
	//generator неэкпортируемая
	rnd rand.Source
}

//New обратите внимание, что мы возвращаем generator, присвоенный интерфейсу, io.Reader
//сама структура неэкспортируемая
//Мы скрыли внутри пакета все детали

func New(seed int64) io.Reader {
	return &generator{
		rnd: rand.NewSource(seed),
	}
}

// Read  реализация io.Reader
func (g *generator) Read(bytes []byte) (n int, err error) { //error это тип ощибки
	for i := range bytes {
		randInt := g.rnd.Int63()  //функция возвращает положительное число от 0 до 2^63
		randByte := byte(randInt) //приводит к типу byte
		bytes[i] = randByte
	}
	return len(bytes), nil
}
