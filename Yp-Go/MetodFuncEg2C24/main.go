package main

import (
	"fmt"
)

/* Метод представляет собой функцию, привязанную к конкретному типу. Методы позволяют
связывать поведение и данные типа в самом типе, обеспечивая инкапсуляцию.

*/
// объявление типа
type DeliveryState string

//Возможные значения перечисления

const (
	DeliveryStatePending   DeliveryState = "pending"      // сообщение отправлено
	DeliveryStateAck       DeliveryState = "acknowledged" // сообщение получено
	DeliveryStateProcessed DeliveryState = "processed"    // сообщение обработано успешно
	DeliveryStateCanceled  DeliveryState = "canceled"     // обработка сообщения прервана
)

/* объявление метода Синатксис метода типа похож на синтаксис обычной функции, но добавляется получатель (receiver)
после ключевого слова func. Можно сказать что получатель - это еще один аргумент функции */

// IsValid проверяет валидность текущего значения типа DeliveryState.
func (s DeliveryState) IsValid() bool {

	switch s {
	case DeliveryStatePending, DeliveryStateAck, DeliveryStateProcessed, DeliveryStateCanceled:
		return true
	default:
		return false

	}
}

// String возвращает строковое представление типа DeliveryState.
func (s DeliveryState) String() string {
	return string(s)
}

func HandleMsgDeliveryStatus(status DeliveryState) error {
	// проверка корректности enum-значения через вызов метода типа DeliveryState
	if !status.IsValid() {
		return fmt.Errorf("status: invalid")
	}

	return nil
}

func main() {

	fmt.Println("Hello")
	//вызов метода

	if err := HandleMsgDeliveryStatus(DeliveryState("fake")); err != nil {
		panic(err)
	}

}
