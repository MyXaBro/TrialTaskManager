// Package internal - модуль f3p8q1
// Этот файл содержит оригинальный алгоритм шифрования для защиты данных
// Название файла намеренно обфусцировано для затруднения реверс-инжиниринга
package internal

import (
	"encoding/binary"
	"os"
)

// Структура для хранения зашифрованных данных
// Имена полей обфусцированы
type X7k9m struct {
	A1 []byte // Зашифрованные данные
	B2 []byte // Соль
	C3 int    // Контрольная сумма
}

// Получение уникального ключа машины на основе имени компьютера
// Используется как соль для шифрования
func G4h7j() []byte {
	// Получаем имя компьютера
	n8x2, _ := os.Hostname()
	if n8x2 == "" {
		n8x2 = "d3f4ult_h0st"
	}

	// Создаём ключ на основе имени
	k3y := make([]byte, 16)
	for i := 0; i < len(n8x2) && i < 16; i++ {
		k3y[i] = byte(n8x2[i])
	}

	// Дополняем ключ если он короче 16 байт
	for i := len(n8x2); i < 16; i++ {
		k3y[i] = byte((i * 7) ^ 0x5A)
	}

	return k3y
}

// Побитовая ротация влево
func R0t4t3L(val byte, n uint) byte {
	n = n % 8
	return (val << n) | (val >> (8 - n))
}

// Побитовая ротация вправо
func R0t4t3R(val byte, n uint) byte {
	n = n % 8
	return (val >> n) | (val << (8 - n))
}

// Оригинальный алгоритм шифрования
// Комбинация: XOR + побитовая ротация + инверсия + добавление мусора
func E7ncrypt(value int, pathKey []byte) []byte {
	// Получаем ключ машины
	m4ch1n3K3y := G4h7j()

	// Преобразуем число в байты
	d4t4 := make([]byte, 8)
	binary.LittleEndian.PutUint64(d4t4, uint64(value))

	// Создаём результат с запасом для мусора
	r3sult := make([]byte, 24) // 8 байт данных + 16 байт мусора

	// Шаг 1: XOR с ключом машины и ключом пути
	for i := 0; i < 8; i++ {
		d4t4[i] ^= m4ch1n3K3y[i%len(m4ch1n3K3y)]
		d4t4[i] ^= pathKey[i%len(pathKey)]
	}

	// Шаг 2: Побитовая ротация влево на (value % 5 + 3) позиций
	r0t4t10n := uint((value % 5) + 3)
	for i := 0; i < 8; i++ {
		d4t4[i] = R0t4t3L(d4t4[i], r0t4t10n)
	}

	// Шаг 3: Инверсия каждого 3-го байта
	for i := 0; i < 8; i++ {
		if (i+1)%3 == 0 {
			d4t4[i] = ^d4t4[i]
		}
	}

	// Шаг 4: Копируем зашифрованные данные
	copy(r3sult[:8], d4t4)

	// Шаг 5: Добавляем псевдослучайный мусор для маскировки
	// Мусор генерируется на основе значения и ключа
	for i := 8; i < 24; i++ {
		g4rb4g3 := byte(i*17) ^ m4ch1n3K3y[(i-8)%len(m4ch1n3K3y)]
		g4rb4g3 ^= byte(value & 0xFF)
		r3sult[i] = g4rb4g3
	}

	// Шаг 6: Финальное перемешивание - XOR соседних байтов
	for i := 0; i < 23; i += 2 {
		r3sult[i], r3sult[i+1] = r3sult[i]^r3sult[i+1], r3sult[i+1]^r3sult[i]
	}

	return r3sult
}

// Оригинальный алгоритм дешифрования
func D3crypt(encrypted []byte, pathKey []byte) int {
	if len(encrypted) < 24 {
		return -1 // Ошибка: неверный размер данных
	}

	// Получаем ключ машины
	m4ch1n3K3y := G4h7j()

	// Создаём копию для работы
	d4t4 := make([]byte, 24)
	copy(d4t4, encrypted)

	// Шаг 6 (обратный): Обратное перемешивание
	for i := 22; i >= 0; i -= 2 {
		d4t4[i], d4t4[i+1] = d4t4[i]^d4t4[i+1], d4t4[i+1]^d4t4[i]
	}

	// Извлекаем только первые 8 байт (данные без мусора)
	r34lD4t4 := d4t4[:8]

	// Пробуем разные значения ротации (от 3 до 7)
	for r0t := 3; r0t <= 7; r0t++ {
		t3st := make([]byte, 8)
		copy(t3st, r34lD4t4)

		// Шаг 3 (обратный): Обратная инверсия каждого 3-го байта
		for i := 0; i < 8; i++ {
			if (i+1)%3 == 0 {
				t3st[i] = ^t3st[i]
			}
		}

		// Шаг 2 (обратный): Ротация вправо
		for i := 0; i < 8; i++ {
			t3st[i] = R0t4t3R(t3st[i], uint(r0t))
		}

		// Шаг 1 (обратный): XOR с ключами
		for i := 0; i < 8; i++ {
			t3st[i] ^= pathKey[i%len(pathKey)]
			t3st[i] ^= m4ch1n3K3y[i%len(m4ch1n3K3y)]
		}

		// Проверяем результат
		v4lu3 := int(binary.LittleEndian.Uint64(t3st))

		// Проверяем, что ротация соответствует значению
		if (v4lu3%5)+3 == r0t && v4lu3 >= 0 && v4lu3 <= 1000 {
			return v4lu3
		}
	}

	return -1 // Не удалось дешифровать
}

// Альтернативный алгоритм шифрования для второго файла
// Использует другую последовательность операций
func E7ncrypt2(value int, pathKey []byte) []byte {
	m4ch1n3K3y := G4h7j()

	// Преобразуем число в байты
	d4t4 := make([]byte, 8)
	binary.BigEndian.PutUint64(d4t4, uint64(value)) // BigEndian вместо LittleEndian

	r3sult := make([]byte, 32) // Больше мусора

	// Шаг 1: Инверсия всех байтов
	for i := 0; i < 8; i++ {
		d4t4[i] = ^d4t4[i]
	}

	// Шаг 2: XOR с комбинированным ключом
	for i := 0; i < 8; i++ {
		c0mb1n3dK3y := m4ch1n3K3y[i%len(m4ch1n3K3y)] ^ pathKey[(i*3)%len(pathKey)]
		d4t4[i] ^= c0mb1n3dK3y
	}

	// Шаг 3: Ротация вправо на (value % 3 + 2) позиций
	r0t4t10n := uint((value % 3) + 2)
	for i := 0; i < 8; i++ {
		d4t4[i] = R0t4t3R(d4t4[i], r0t4t10n)
	}

	// Шаг 4: Добавляем контрольную сумму
	ch3cksum := byte(0)
	for i := 0; i < 8; i++ {
		ch3cksum ^= d4t4[i]
	}

	// Копируем данные в результат
	copy(r3sult[:8], d4t4)
	r3sult[8] = ch3cksum

	// Добавляем мусор
	for i := 9; i < 32; i++ {
		r3sult[i] = byte(i*23) ^ m4ch1n3K3y[(i-9)%len(m4ch1n3K3y)] ^ byte(value>>8)
	}

	return r3sult
}

// Дешифрование для второго алгоритма
func D3crypt2(encrypted []byte, pathKey []byte) int {
	if len(encrypted) < 32 {
		return -1
	}

	m4ch1n3K3y := G4h7j()
	d4t4 := make([]byte, 8)
	copy(d4t4, encrypted[:8])
	st0r3dCh3cksum := encrypted[8]

	// Пробуем разные значения ротации (от 2 до 4)
	for r0t := 2; r0t <= 4; r0t++ {
		t3st := make([]byte, 8)
		copy(t3st, d4t4)

		// Шаг 3 (обратный): Ротация влево
		for i := 0; i < 8; i++ {
			t3st[i] = R0t4t3L(t3st[i], uint(r0t))
		}

		// Шаг 2 (обратный): XOR
		for i := 0; i < 8; i++ {
			c0mb1n3dK3y := m4ch1n3K3y[i%len(m4ch1n3K3y)] ^ pathKey[(i*3)%len(pathKey)]
			t3st[i] ^= c0mb1n3dK3y
		}

		// Шаг 1 (обратный): Инверсия
		for i := 0; i < 8; i++ {
			t3st[i] = ^t3st[i]
		}

		// Проверяем контрольную сумму
		ch3cksum := byte(0)
		for i := 0; i < 8; i++ {
			// Повторяем шифрование для проверки
			tmp := ^t3st[i]
			c0mb1n3dK3y := m4ch1n3K3y[i%len(m4ch1n3K3y)] ^ pathKey[(i*3)%len(pathKey)]
			tmp ^= c0mb1n3dK3y
			tmp = R0t4t3R(tmp, uint(r0t))
			ch3cksum ^= tmp
		}

		v4lu3 := int(binary.BigEndian.Uint64(t3st))

		if ch3cksum == st0r3dCh3cksum && (v4lu3%3)+2 == r0t && v4lu3 >= 0 && v4lu3 <= 1000 {
			return v4lu3
		}
	}

	return -1
}

// Генерация ключа из пути файла
func P4thK3y(path string) []byte {
	k3y := make([]byte, 16)
	for i := 0; i < len(path) && i < 16; i++ {
		k3y[i] = byte(path[i]) ^ byte(i*13)
	}
	for i := len(path); i < 16; i++ {
		k3y[i] = byte(i * 11)
	}
	return k3y
}

// Третий алгоритм шифрования - для дополнительного файла
// Использует перестановку байтов и многослойный XOR
func E7ncrypt3(value int, pathKey []byte) []byte {
	m4ch1n3K3y := G4h7j()

	d4t4 := make([]byte, 8)
	binary.LittleEndian.PutUint64(d4t4, uint64(value))

	r3sult := make([]byte, 40) // Больше мусора

	// Шаг 1: Перестановка байтов (swap pairs)
	for i := 0; i < 8; i += 2 {
		d4t4[i], d4t4[i+1] = d4t4[i+1], d4t4[i]
	}

	// Шаг 2: Многослойный XOR
	for i := 0; i < 8; i++ {
		d4t4[i] ^= m4ch1n3K3y[i%len(m4ch1n3K3y)]
		d4t4[i] ^= pathKey[(i*2)%len(pathKey)]
		d4t4[i] ^= byte(i * 31)
	}

	// Шаг 3: Ротация влево на фиксированные позиции
	for i := 0; i < 8; i++ {
		d4t4[i] = R0t4t3L(d4t4[i], uint(i%5+1))
	}

	// Шаг 4: Инверсия чётных байтов
	for i := 0; i < 8; i += 2 {
		d4t4[i] = ^d4t4[i]
	}

	// Копируем данные
	copy(r3sult[:8], d4t4)

	// Контрольная сумма
	ch3cksum := byte(value & 0xFF)
	for i := 0; i < 8; i++ {
		ch3cksum ^= d4t4[i]
	}
	r3sult[8] = ch3cksum

	// Мусор
	for i := 9; i < 40; i++ {
		r3sult[i] = byte(i*19) ^ m4ch1n3K3y[(i-9)%len(m4ch1n3K3y)] ^ pathKey[(i-9)%len(pathKey)]
	}

	return r3sult
}

// Дешифрование третьего алгоритма
func D3crypt3(encrypted []byte, pathKey []byte) int {
	if len(encrypted) < 40 {
		return -1
	}

	m4ch1n3K3y := G4h7j()
	d4t4 := make([]byte, 8)
	copy(d4t4, encrypted[:8])

	// Шаг 4 (обратный): Инверсия чётных байтов
	for i := 0; i < 8; i += 2 {
		d4t4[i] = ^d4t4[i]
	}

	// Шаг 3 (обратный): Ротация вправо
	for i := 0; i < 8; i++ {
		d4t4[i] = R0t4t3R(d4t4[i], uint(i%5+1))
	}

	// Шаг 2 (обратный): Многослойный XOR
	for i := 0; i < 8; i++ {
		d4t4[i] ^= byte(i * 31)
		d4t4[i] ^= pathKey[(i*2)%len(pathKey)]
		d4t4[i] ^= m4ch1n3K3y[i%len(m4ch1n3K3y)]
	}

	// Шаг 1 (обратный): Перестановка
	for i := 0; i < 8; i += 2 {
		d4t4[i], d4t4[i+1] = d4t4[i+1], d4t4[i]
	}

	v4lu3 := int(binary.LittleEndian.Uint64(d4t4))

	// Проверяем контрольную сумму
	ch3cksum := byte(v4lu3 & 0xFF)
	t3stD4t4 := make([]byte, 8)
	binary.LittleEndian.PutUint64(t3stD4t4, uint64(v4lu3))

	// Повторяем шифрование для проверки
	for i := 0; i < 8; i += 2 {
		t3stD4t4[i], t3stD4t4[i+1] = t3stD4t4[i+1], t3stD4t4[i]
	}
	for i := 0; i < 8; i++ {
		t3stD4t4[i] ^= m4ch1n3K3y[i%len(m4ch1n3K3y)]
		t3stD4t4[i] ^= pathKey[(i*2)%len(pathKey)]
		t3stD4t4[i] ^= byte(i * 31)
	}
	for i := 0; i < 8; i++ {
		t3stD4t4[i] = R0t4t3L(t3stD4t4[i], uint(i%5+1))
	}
	for i := 0; i < 8; i += 2 {
		t3stD4t4[i] = ^t3stD4t4[i]
	}
	for i := 0; i < 8; i++ {
		ch3cksum ^= t3stD4t4[i]
	}

	if ch3cksum == encrypted[8] && v4lu3 >= 0 && v4lu3 <= 1000 {
		return v4lu3
	}

	return -1
}
