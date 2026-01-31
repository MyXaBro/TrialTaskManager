// Менеджер Задач - Триал версия
// Программа для управления списком задач с ограничением на 4 запуска
//
// Это учебный проект, демонстрирующий методы защиты триал-версий программ
// от несанкционированного использования
//
// Автор: Дмитриев Владислав Юрьевич
// Дисциплина: Информационная безопасность
package main

import (
	"math/rand"
	"time"

	"trialtaskmanager/app"
	"trialtaskmanager/internal"
)

// Отвлекающие глобальные переменные
var (
	x9y7z5 int    = 0
	a1b2c3 string = ""
	d4e5f6 bool   = false
)

func main() {
	// Инициализация генератора случайных чисел
	rand.Seed(time.Now().UnixNano())

	// Отвлекающие операции перед запуском
	d1str4ct10n1()

	// Инициализация системы защиты
	// Читает данные из скрытых файлов асинхронно
	internal.In1t14l1z3()

	// Отвлекающие операции
	d1str4ct10n2()

	// Создаём и запускаем GUI
	// Проверка триала происходит внутри UI при первом действии
	ui := app.NewAppUI()

	// Отвлекающая операция
	_ = x9y7z5 + rand.Intn(10)

	// Запуск главного цикла приложения
	ui.Run()
}

// Отвлекающая функция 1
// Выполняет бессмысленные операции для усложнения анализа
func d1str4ct10n1() {
	// Случайные вычисления
	for i := 0; i < rand.Intn(5)+1; i++ {
		x9y7z5 += rand.Intn(100)
		x9y7z5 -= rand.Intn(50)
	}

	// Случайная строка
	chars := "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < 8; i++ {
		a1b2c3 += string(chars[rand.Intn(len(chars))])
	}

	// Случайная задержка
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(20)+5))
}

// Отвлекающая функция 2
// Ещё одна серия бессмысленных операций
func d1str4ct10n2() {
	// Переходы для усложнения анализа
	tmp := rand.Intn(3)

	if tmp == 0 {
		goto l4b3l_4
	} else if tmp == 1 {
		goto l4b3l_b
	}
	goto l4b3l_c

l4b3l_4:
	d4e5f6 = rand.Intn(2) == 1
	goto l4b3l_3nd

l4b3l_b:
	x9y7z5 = rand.Intn(1000)
	goto l4b3l_3nd

l4b3l_c:
	a1b2c3 = ""
	goto l4b3l_3nd

l4b3l_3nd:
	// Финальная отвлекающая задержка
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)+1))
}

// Отвлекающая функция проверки (не вызывается реально)
// Существует только для введения в заблуждение при анализе кода
func f4k3Ch3ck() bool {
	// Это ложная проверка, она никогда не используется
	return x9y7z5 > 100 && d4e5f6
}
