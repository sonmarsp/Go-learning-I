package main

import "testing"

//команда go test найдет все функции начинающиеся с тест
func TestMultiple(t *testing.T) {

	//простой тест
	//var x, y, result = 2, 2, 4
	//realresult := Multiple(x, y)
	//if realresult != result {
	//	t.Errorf("%d != %d", result, realresult)
	//}

	//можно делать иерархию тестов
	//можно написать тест подготовки обычно вставка данных в базу если не мокируете
	// insert data in db

	//

	t.Run("simple", func(t *testing.T) {
		//Тесты можно выполнять параллельно
		t.Parallel()
		t.Log("simple")

		var x, y, result = 2, 2, 4

		realResult := Multiple(x, y)

		if realResult != result {

			t.Errorf("expected result %d != %d real result", result, realResult)

		}

	})

	t.Run("medium", func(t *testing.T) {
		t.Parallel()
		t.Log("medium")

		var x, y, result = 222, 222, 49284

		realResult := Multiple(x, y)

		if realResult != result {

			t.Errorf("expected result %d != %d real result", result, realResult)

		}

	})

	t.Run("negative", func(t *testing.T) {
		t.Parallel()
		t.Log("negative")

		var x, y, result = -2, 4, -8

		realResult := Multiple(x, y)

		if realResult != result {

			t.Errorf("expected result %d != %d real result", result, realResult)

		}

	})
	//delete test data within db

}

// С v дает больше информации
//go test -v
//Можно запускать какой то тест конкретный
// go test -v -run TestMultiple/simple
