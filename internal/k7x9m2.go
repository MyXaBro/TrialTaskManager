// Package internal - модуль k7x9m2
// Этот файл содержит основную логику защиты и проверки триал-периода
// Реализует распределённую систему проверок с отвлекающими операциями
package internal

import (
	"math/rand"
	"time"

	"golang.org/x/sys/windows/registry"
)

// Глобальные переменные с обфусцированными именами
// Хранят промежуточные результаты проверок
var (
	q7w3e9 int = -1 // Значение из первого файла
	r8t4y2 int = -1 // Значение из второго файла
	u5i1o6 int = -1 // Отвлекающее значение из реестра
	p9a3s7 int = -1 // Финальное значение счётчика
)

// Маркеры состояния (обфусцированные)
var (
	z1x2c3 bool = false // Первый файл прочитан
	v4b5n6 bool = false // Второй файл прочитан
	m7k8j9 bool = false // Проверка пройдена
	l0h1g2 bool = false // Инкремент выполнен
)

// Максимальное количество запусков
const mX_L4unch3s = 4

// Инициализация защиты - вызывается при старте программы
// Читает данные из всех источников, но НЕ проверяет их сразу
func In1t14l1z3() {
	// Отвлекающая операция - случайная задержка
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(50)+10))

	// Читаем первый реальный файл
	go func() {
		p4th1 := G3tH1dd3nP4th1()
		k3y1 := P4thK3y(p4th1)

		if F1l33x1sts(p4th1) {
			d4t4, err := S4f3R34d(p4th1)
			if err == nil {
				q7w3e9 = D3crypt(d4t4, k3y1)
			}
		} else {
			q7w3e9 = 0
		}
		z1x2c3 = true
	}()

	// Отвлекающее чтение из реестра
	go func() {
		d3c0yR34d()
	}()

	// Читаем второй реальный файл
	go func() {
		// Небольшая случайная задержка для маскировки
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(30)+5))

		p4th2 := G3tH1dd3nP4th2()
		k3y2 := P4thK3y(p4th2)

		if F1l33x1sts(p4th2) {
			d4t4, err := S4f3R34d(p4th2)
			if err == nil {
				r8t4y2 = D3crypt2(d4t4, k3y2)
			}
		} else {
			r8t4y2 = 0
		}
		v4b5n6 = true
	}()

	// Создаём/обновляем отвлекающий файл
	go func() {
		d3c0yF1l3()
	}()
}

// Отвлекающее чтение из реестра Windows
// Данные в реестре - ложные, используются только для отвлечения
func d3c0yR34d() {
	// Пытаемся прочитать из отвлекающего места в реестре
	k3y, err := registry.OpenKey(
		registry.CURRENT_USER,
		`Software\Microsoft\Windows\CurrentVersion\Explorer\Advanced`,
		registry.QUERY_VALUE,
	)
	if err != nil {
		u5i1o6 = rand.Intn(10) // Случайное значение
		return
	}
	defer k3y.Close()

	// Читаем несуществующее значение (или существующее - не важно)
	val, _, err := k3y.GetIntegerValue("TaskManagerRuns")
	if err != nil {
		u5i1o6 = rand.Intn(10)
	} else {
		u5i1o6 = int(val)
	}
}

// Создание/обновление отвлекающего файла
func d3c0yF1l3() {
	d3c0yP4th := G3tD3c0yP4th()

	// Записываем случайные данные в отвлекающий файл
	fak3D4t4 := make([]byte, 64)
	rand.Read(fak3D4t4)

	// Не важно, успешно ли - это отвлечение
	S4f3Wr1t3(d3c0yP4th, fak3D4t4)
}

// Ожидание завершения чтения всех файлов
func W41tF0rR34d() {
	// Ждём пока оба флага не станут true
	for !z1x2c3 || !v4b5n6 {
		time.Sleep(time.Millisecond * 10)
	}

	// Ещё одна отвлекающая задержка
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(20)+5))
}

// Проверка консистентности данных
// Вызывается позже, НЕ сразу после чтения
func V3r1fyC0ns1st3ncy() {
	// Отвлекающие вычисления
	_ = u5i1o6 * 2
	_ = rand.Intn(100)

	// Реальная проверка консистентности
	if q7w3e9 != r8t4y2 {
		// Данные не совпадают - возможно взлом
		// Берём максимальное значение как защита
		if q7w3e9 > r8t4y2 {
			p9a3s7 = q7w3e9
		} else {
			p9a3s7 = r8t4y2
		}
	} else {
		p9a3s7 = q7w3e9
	}

	// Если оба файла новые (значение 0 или -1)
	if p9a3s7 < 0 {
		p9a3s7 = 0
	}

	// Отвлекающая операция
	_ = u5i1o6 + p9a3s7
}

// Инкремент счётчика запусков
// Вызывается при первом значимом действии пользователя
func Incr3m3ntC0unt3r() bool {
	if l0h1g2 {
		return true // Уже инкрементировали
	}

	// Отвлекающие операции перед инкрементом
	_ = rand.Intn(50)
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))

	// Увеличиваем счётчик
	p9a3s7++

	// Записываем в первый файл
	go func() {
		p4th1 := G3tH1dd3nP4th1()
		k3y1 := P4thK3y(p4th1)
		encr1pt3d := E7ncrypt(p9a3s7, k3y1)
		S4f3Wr1t3(p4th1, encr1pt3d)
	}()

	// Записываем во второй файл
	go func() {
		p4th2 := G3tH1dd3nP4th2()
		k3y2 := P4thK3y(p4th2)
		encr1pt3d := E7ncrypt2(p9a3s7, k3y2)
		S4f3Wr1t3(p4th2, encr1pt3d)
	}()

	// Обновляем отвлекающий файл
	go func() {
		d3c0yF1l3()
	}()

	l0h1g2 = true

	// Отвлекающая запись в реестр
	go func() {
		d3c0yR3g1stryWr1t3()
	}()

	return true
}

// Отвлекающая запись в реестр
func d3c0yR3g1stryWr1t3() {
	k3y, _, err := registry.CreateKey(
		registry.CURRENT_USER,
		`Software\TaskManagerApp`,
		registry.SET_VALUE,
	)
	if err != nil {
		return
	}
	defer k3y.Close()

	// Записываем случайное (ложное) значение
	k3y.SetDWordValue("RunCount", uint32(rand.Intn(100)))
}

// Проверка, не исчерпан ли лимит запусков
// НЕ возвращает результат напрямую - использует goto для усложнения анализа
func Ch3ckL1m1t() bool {
	// Множество отвлекающих переходов и вычислений
	var r3sult bool
	tmp1 := rand.Intn(10)
	tmp2 := u5i1o6 % 7

	if tmp1 > 5 {
		goto l4b3l1
	}
	goto l4b3l2

l4b3l1:
	_ = tmp2 * 2
	goto l4b3l3

l4b3l2:
	_ = tmp1 + tmp2
	goto l4b3l3

l4b3l3:
	// Отвлекающая проверка
	if u5i1o6 > 100 {
		_ = u5i1o6
	}

	// Реальная проверка (но не сразу возвращаем результат)
	if p9a3s7 > mX_L4unch3s {
		r3sult = false
		goto l4b3l4
	}
	r3sult = true
	goto l4b3l4

l4b3l4:
	// Ещё отвлекающие операции
	_ = rand.Intn(5)
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(5)))

	return r3sult
}

// Получение оставшегося количества запусков
func G3tR3m41n1ng() int {
	r3m := mX_L4unch3s - p9a3s7
	if r3m < 0 {
		r3m = 0
	}
	return r3m
}

// Получение текущего количества запусков
func G3tCurr3nt() int {
	return p9a3s7
}

// Принудительная инвалидация триала (для тестирования)
func F0rc31nv4l1d4t3() {
	p9a3s7 = mX_L4unch3s + 1

	p4th1 := G3tH1dd3nP4th1()
	k3y1 := P4thK3y(p4th1)
	S4f3Wr1t3(p4th1, E7ncrypt(p9a3s7, k3y1))

	p4th2 := G3tH1dd3nP4th2()
	k3y2 := P4thK3y(p4th2)
	S4f3Wr1t3(p4th2, E7ncrypt2(p9a3s7, k3y2))
}
