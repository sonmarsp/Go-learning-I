package main

import (
	"fmt"
	"os"
	"testing"
)

// Можно сделать глобальную обработку для всего пакета в данном случае main
func TestMain(m *testing.M) {
	fmt.Println("setup")
	res := m.Run()
	fmt.Printf("tear-down")

	os.Exit(res)
}

// команда go test найдет все функции начинающиеся с тест
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

	//можно объединят в группы они запускаются последовательно а тесты в них параллельно

	t.Run("groupA", func(t *testing.T) {
		t.Log("A")
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

	})

	t.Run("groupB", func(t *testing.T) {
		t.Log("B")
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

	})

	t.Run("groupC", func(t *testing.T) {
		t.Log("C")
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

	})

	t.Run("groupD", func(t *testing.T) {

		t.Log("D")
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

	})

	//teer down выполняется последним даже при параллельном запуске
	//delete test data within db

}

// С v дает больше информации
//go test -v
//Можно запускать какой то тест конкретный
// go test -v -run TestMultiple/simple
