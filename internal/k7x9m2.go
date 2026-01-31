// Package internal - модуль k7x9m2
// Этот файл содержит основную логику защиты и проверки триал-периода
// Реализует распределённую систему проверок с отвлекающими операциями
package internal

import (
	"encoding/binary"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/sys/windows/registry"
)

// Структура для косвенного хранения результатов
type v9z3x7 struct {
	a1 int    // Зашифрованное значение 1
	b2 int    // Зашифрованное значение 2
	c3 int    // Зашифрованное значение 3
	d4 []byte // Отвлекающие данные
	e5 int64  // Метка времени
	f6 bool   // Флаг валидации
}

// Глобальные обфусцированные структуры
var (
	h8k2m4 = &v9z3x7{a1: -1, b2: -1, c3: -1} // Файл 1
	j3n5p7 = &v9z3x7{a1: -1, b2: -1, c3: -1} // Файл 2
	l6q8s1 = &v9z3x7{a1: -1, b2: -1, c3: -1} // Файл 3
	r9g4k2 = &v9z3x7{a1: -1, b2: -1, c3: -1} // Реестр (резервный!)
	n7a3d5 = &v9z3x7{a1: -1, b2: -1, c3: -1} // ADS (скрытый поток!)
	w2y4t6 = &v9z3x7{a1: -1, b2: -1, c3: -1} // Отвлекающий
)

// Маркеры состояния через битовые операции
var (
	x5c7v9 int = 0 // Биты: 0-файл1, 1-файл2, 2-файл3, 3-реестр, 4-ADS, 5-проверка
)

// Магические числа для косвенных вычислений
const (
	m4g1c1 = 0x5A3B
	m4g1c2 = 0x7C9D
	m4g1c3 = 0x2E4F
	m4g1c4 = 0x8B1A // Для реестра
	m4g1c5 = 0x3F6C // Для ADS
)

// Максимальное количество запусков (замаскировано вычислением)
func g3tM4xL4unch() int {
	r := ((m4g1c1 ^ m4g1c2) % 10) - 3
	if r < 4 {
		r = 4
	}
	if r > 4 {
		r = 4
	}
	return r
}

// Случайная задержка для защиты от анализа времени выполнения
func r4nd0mD3l4y() {
	d3l4y := time.Duration(rand.Intn(50)+10) * time.Millisecond
	time.Sleep(d3l4y)
}

// Анти-дебаг: проверка времени выполнения
func ch3ckT1m1ng() bool {
	st4rt := time.Now()
	x := 0
	for i := 0; i < 1000; i++ {
		x += i
	}
	_ = x
	e14ps3d := time.Since(st4rt)
	return e14ps3d < time.Millisecond*100
}

// ============ РЕЗЕРВНОЕ ХРАНЕНИЕ В РЕЕСТРЕ ============

// Обфусцированный путь реестра (выглядит как системный)
func g3tR3g1stryP4th() string {
	// HKCU\Software\Microsoft\Windows\CurrentVersion\Explorer\ComDlg32\CIDSizeMRU
	p := []byte{0x53, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65} // Software
	p2 := []byte{0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x6f, 0x66, 0x74}
	p3 := []byte{0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73}
	p4 := []byte{0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e}
	p5 := []byte{0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72}
	p6 := []byte{0x43, 0x6f, 0x6d, 0x44, 0x6c, 0x67, 0x33, 0x32}
	p7 := []byte{0x43, 0x49, 0x44, 0x53, 0x69, 0x7a, 0x65, 0x4d, 0x52, 0x55}

	return string(p) + `\` + string(p2) + `\` + string(p3) + `\` + string(p4) + `\` + string(p5) + `\` + string(p6) + `\` + string(p7)
}

// Обфусцированное имя значения в реестре
func g3tR3g1stryV4lu3N4m3() string {
	// MRUListEx - выглядит как системное
	return string([]byte{0x4d, 0x52, 0x55, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x78})
}

// Читаем счётчик из реестра
func r34dR3g1stry() int {
	k3y, err := registry.OpenKey(
		registry.CURRENT_USER,
		g3tR3g1stryP4th(),
		registry.QUERY_VALUE,
	)
	if err != nil {
		return -1
	}
	defer k3y.Close()

	d4t4, _, err := k3y.GetBinaryValue(g3tR3g1stryV4lu3N4m3())
	if err != nil || len(d4t4) < 8 {
		return -1
	}

	// Дешифруем - простой XOR с magic
	v4l := binary.LittleEndian.Uint64(d4t4[:8])
	d3crypt3d := int(v4l ^ uint64(m4g1c4)<<32 ^ uint64(m4g1c4))

	if d3crypt3d < 0 || d3crypt3d > 1000 {
		return -1
	}
	return d3crypt3d
}

// Записываем счётчик в реестр
func wr1t3R3g1stry(value int) {
	k3y, _, err := registry.CreateKey(
		registry.CURRENT_USER,
		g3tR3g1stryP4th(),
		registry.SET_VALUE,
	)
	if err != nil {
		return
	}
	defer k3y.Close()

	// Шифруем - XOR с magic
	e3ncrypt3d := uint64(value) ^ uint64(m4g1c4)<<32 ^ uint64(m4g1c4)

	d4t4 := make([]byte, 16) // Добавляем мусор
	binary.LittleEndian.PutUint64(d4t4[:8], e3ncrypt3d)
	rand.Read(d4t4[8:]) // Мусор

	k3y.SetBinaryValue(g3tR3g1stryV4lu3N4m3(), d4t4)
}

// ============ ХРАНЕНИЕ В NTFS ADS (Alternate Data Stream) ============

// Получаем путь к ADS - прикрепляем к существующему системному файлу
func g3tADSP4th() string {
	// %LOCALAPPDATA%\Microsoft\Windows\WebCache\WebCacheV01.dat:Zone.Identifier
	// Это выглядит как системный Zone.Identifier
	b4s3 := os.Getenv("LOCALAPPDATA")
	if b4s3 == "" {
		b4s3 = filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Local")
	}

	// Собираем путь из байтов
	p1 := []byte{0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x6f, 0x66, 0x74} // Microsoft
	p2 := []byte{0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73}             // Windows
	p3 := []byte{0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72}       // Explorer
	// Имя потока (после :)
	str34m := []byte{0x3a, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61} // :_zone_data

	// Создаём директорию если не существует
	d1r := filepath.Join(b4s3, string(p1), string(p2), string(p3))
	os.MkdirAll(d1r, 0755)

	// Создаём базовый файл если не существует
	b4s3F1l3 := filepath.Join(d1r, "IconCache.db")
	if _, err := os.Stat(b4s3F1l3); os.IsNotExist(err) {
		// Создаём файл с мусором
		f, _ := os.Create(b4s3F1l3)
		if f != nil {
			junk := make([]byte, 256)
			rand.Read(junk)
			f.Write(junk)
			f.Close()
		}
	}

	// Возвращаем путь к ADS
	return b4s3F1l3 + string(str34m)
}

// Читаем счётчик из ADS
func r34dADS() int {
	p4th := g3tADSP4th()

	d4t4, err := os.ReadFile(p4th)
	if err != nil || len(d4t4) < 8 {
		return -1
	}

	// Дешифруем
	v4l := binary.LittleEndian.Uint64(d4t4[:8])
	d3crypt3d := int(v4l ^ uint64(m4g1c5)<<32 ^ uint64(m4g1c5))

	if d3crypt3d < 0 || d3crypt3d > 1000 {
		return -1
	}
	return d3crypt3d
}

// Записываем счётчик в ADS
func wr1t3ADS(value int) {
	p4th := g3tADSP4th()

	// Шифруем
	e3ncrypt3d := uint64(value) ^ uint64(m4g1c5)<<32 ^ uint64(m4g1c5)

	d4t4 := make([]byte, 24)
	binary.LittleEndian.PutUint64(d4t4[:8], e3ncrypt3d)
	rand.Read(d4t4[8:]) // Мусор

	os.WriteFile(p4th, d4t4, 0644)
}

// ============ ИНИЦИАЛИЗАЦИЯ С ЗАЩИТОЙ ОТ УДАЛЕНИЯ ============

func In1t14l1z3() {
	rand.Seed(time.Now().UnixNano())

	st3p := 0
	c0unt3r := 0
	t3mp := 0

	// Счётчики для обнаружения удаления
	f1l3sM1ss1ng := 0
	b4ckupsF0und := 0

L4B3L_ST4RT:
	if st3p == 0 {
		r4nd0mD3l4y()
		st3p = 1
		goto L4B3L_CH3CK
	}

L4B3L_F1L31:
	if st3p == 1 {
		p4th1 := G3tH1dd3nP4th1()
		k3y1 := P4thK3y(p4th1)

		if F1l33x1sts(p4th1) {
			d4t4, err := S4f3R34d(p4th1)
			if err == nil {
				t3mp = D3crypt(d4t4, k3y1)
				h8k2m4.a1 = t3mp ^ m4g1c1
				h8k2m4.e5 = time.Now().UnixNano()
			}
		} else {
			h8k2m4.a1 = 0 ^ m4g1c1
			f1l3sM1ss1ng++
		}
		x5c7v9 |= 1
		st3p = 2
		goto L4B3L_D3C0Y1
	}

L4B3L_D3C0Y1:
	d3c0yR34d()
	r4nd0mD3l4y()
	c0unt3r++
	if c0unt3r < 2 {
		st3p = 3
		goto L4B3L_F1L32
	}

L4B3L_F1L32:
	if st3p == 2 || st3p == 3 {
		p4th2 := G3tH1dd3nP4th2()
		k3y2 := P4thK3y(p4th2)

		if F1l33x1sts(p4th2) {
			d4t4, err := S4f3R34d(p4th2)
			if err == nil {
				t3mp = D3crypt2(d4t4, k3y2)
				j3n5p7.a1 = t3mp ^ m4g1c2
				j3n5p7.e5 = time.Now().UnixNano()
			}
		} else {
			j3n5p7.a1 = 0 ^ m4g1c2
			f1l3sM1ss1ng++
		}
		x5c7v9 |= 2
		st3p = 4
		goto L4B3L_F1L33
	}

L4B3L_F1L33:
	if st3p == 4 {
		p4th3 := G3tH1dd3nP4th3()
		k3y3 := P4thK3y(p4th3)

		if F1l33x1sts(p4th3) {
			d4t4, err := S4f3R34d(p4th3)
			if err == nil {
				t3mp = D3crypt(d4t4, k3y3)
				l6q8s1.a1 = t3mp ^ m4g1c3
				l6q8s1.e5 = time.Now().UnixNano()
			}
		} else {
			l6q8s1.a1 = 0 ^ m4g1c3
			f1l3sM1ss1ng++
		}
		x5c7v9 |= 4
		st3p = 5
		goto L4B3L_R3G1STRY
	}

L4B3L_R3G1STRY:
	// Читаем резервную копию из реестра
	if st3p == 5 {
		r3gV4l := r34dR3g1stry()
		if r3gV4l >= 0 {
			r9g4k2.a1 = r3gV4l ^ m4g1c4
			b4ckupsF0und++
		} else {
			r9g4k2.a1 = 0 ^ m4g1c4
		}
		x5c7v9 |= 8
		st3p = 6
		goto L4B3L_ADS
	}

L4B3L_ADS:
	// Читаем резервную копию из ADS
	if st3p == 6 {
		a4dsV4l := r34dADS()
		if a4dsV4l >= 0 {
			n7a3d5.a1 = a4dsV4l ^ m4g1c5
			b4ckupsF0und++
		} else {
			n7a3d5.a1 = 0 ^ m4g1c5
		}
		x5c7v9 |= 16
		st3p = 7
		goto L4B3L_T4MP3R_CH3CK
	}

L4B3L_T4MP3R_CH3CK:
	// ЗАЩИТА ОТ УДАЛЕНИЯ: если все 3 файла отсутствуют, но бэкапы есть
	if f1l3sM1ss1ng >= 3 && b4ckupsF0und > 0 {
		// Пользователь удалил файлы! Восстанавливаем из бэкапа
		// и устанавливаем максимальное значение (триал истёк)
		m4xV4l := r9g4k2.a1 ^ m4g1c4
		a4dsV4l := n7a3d5.a1 ^ m4g1c5

		// Берём максимум из бэкапов
		if a4dsV4l > m4xV4l {
			m4xV4l = a4dsV4l
		}

		// Если было использовано хотя бы 1 раз - блокируем полностью
		if m4xV4l > 0 {
			m4xV4l = g3tM4xL4unch() + 1 // Триал истёк!
		}

		// Восстанавливаем значения
		h8k2m4.a1 = m4xV4l ^ m4g1c1
		j3n5p7.a1 = m4xV4l ^ m4g1c2
		l6q8s1.a1 = m4xV4l ^ m4g1c3
	}
	st3p = 8
	goto L4B3L_D3C0Y2

L4B3L_D3C0Y2:
	d3c0yF1l3()
	w2y4t6.d4 = make([]byte, rand.Intn(32)+16)
	rand.Read(w2y4t6.d4)
	st3p = 9
	goto L4B3L_3ND

L4B3L_CH3CK:
	if !ch3ckT1m1ng() {
		r4nd0mD3l4y()
		r4nd0mD3l4y()
	}
	st3p = 1
	goto L4B3L_F1L31

L4B3L_3ND:
	x5c7v9 |= 32
	_ = st3p
	_ = f1l3sM1ss1ng
	_ = b4ckupsF0und
	return

	goto L4B3L_ST4RT
}

// Отвлекающее чтение из реестра Windows
func d3c0yR34d() {
	k3y, err := registry.OpenKey(
		registry.CURRENT_USER,
		`Software\Microsoft\Windows\CurrentVersion\Explorer\Advanced`,
		registry.QUERY_VALUE,
	)
	if err != nil {
		w2y4t6.a1 = rand.Intn(100)
		return
	}
	defer k3y.Close()

	val, _, err := k3y.GetIntegerValue("TaskbarSmallIcons")
	if err != nil {
		w2y4t6.a1 = rand.Intn(100)
	} else {
		w2y4t6.a1 = int(val) ^ 0x55
	}
}

// Создание/обновление отвлекающего файла
func d3c0yF1l3() {
	d3c0yP4th := G3tD3c0yP4th()
	fak3D4t4 := make([]byte, 64+rand.Intn(64))
	rand.Read(fak3D4t4)
	S4f3Wr1t3(d3c0yP4th, fak3D4t4)

	d3c0yP4th2 := G3tD3c0yP4th2()
	fak3D4t42 := make([]byte, 128+rand.Intn(64))
	rand.Read(fak3D4t42)
	S4f3Wr1t3(d3c0yP4th2, fak3D4t42)
}

// Проверка консистентности
func V3r1fyC0ns1st3ncy() {
	for (x5c7v9 & 31) != 31 {
		r4nd0mD3l4y()
	}

	// Собираем значения из всех источников
	v4l1 := h8k2m4.a1 ^ m4g1c1
	v4l2 := j3n5p7.a1 ^ m4g1c2
	v4l3 := l6q8s1.a1 ^ m4g1c3
	v4lR3g := r9g4k2.a1 ^ m4g1c4
	v4lADS := n7a3d5.a1 ^ m4g1c5

	// Берём МАКСИМУМ из всех источников (защита от манипуляций)
	f1n4l := v4l1
	st3p := 0

L4B3L_CMP1:
	if st3p == 0 {
		if v4l2 > f1n4l {
			f1n4l = v4l2
		}
		st3p = 1
		goto L4B3L_CMP2
	}

L4B3L_CMP2:
	if st3p == 1 {
		if v4l3 > f1n4l {
			f1n4l = v4l3
		}
		st3p = 2
		goto L4B3L_CMP3
	}

L4B3L_CMP3:
	if st3p == 2 {
		if v4lR3g > f1n4l {
			f1n4l = v4lR3g
		}
		st3p = 3
		goto L4B3L_CMP4
	}

L4B3L_CMP4:
	if st3p == 3 {
		if v4lADS > f1n4l {
			f1n4l = v4lADS
		}
		st3p = 4
		goto L4B3L_S4V3
	}

L4B3L_S4V3:
	if f1n4l < 0 {
		f1n4l = 0
	}

	h8k2m4.b2 = f1n4l ^ m4g1c1
	j3n5p7.b2 = f1n4l ^ m4g1c2
	l6q8s1.b2 = f1n4l ^ m4g1c3
	r9g4k2.b2 = f1n4l ^ m4g1c4
	n7a3d5.b2 = f1n4l ^ m4g1c5

	h8k2m4.f6 = true
	j3n5p7.f6 = true
	l6q8s1.f6 = true

	x5c7v9 |= 64
	_ = st3p
	return

	goto L4B3L_CMP1
}

// Инкремент счётчика - записываем во ВСЕ хранилища
func Incr3m3ntC0unt3r() bool {
	if (x5c7v9 & 128) != 0 {
		return true
	}

	r4nd0mD3l4y()

	curr3nt := h8k2m4.b2 ^ m4g1c1
	if curr3nt < 0 {
		curr3nt = 0
	}

	n3wV4l := curr3nt + 1
	st3p := 0

L4B3L_WR1:
	if st3p == 0 {
		p4th1 := G3tH1dd3nP4th1()
		k3y1 := P4thK3y(p4th1)
		encr1pt3d := E7ncrypt(n3wV4l, k3y1)
		S4f3Wr1t3(p4th1, encr1pt3d)
		st3p = 1
		r4nd0mD3l4y()
		goto L4B3L_WR2
	}

L4B3L_WR2:
	if st3p == 1 {
		p4th2 := G3tH1dd3nP4th2()
		k3y2 := P4thK3y(p4th2)
		encr1pt3d2 := E7ncrypt2(n3wV4l, k3y2)
		S4f3Wr1t3(p4th2, encr1pt3d2)
		st3p = 2
		r4nd0mD3l4y()
		goto L4B3L_WR3
	}

L4B3L_WR3:
	if st3p == 2 {
		p4th3 := G3tH1dd3nP4th3()
		k3y3 := P4thK3y(p4th3)
		encr1pt3d3 := E7ncrypt(n3wV4l, k3y3)
		S4f3Wr1t3(p4th3, encr1pt3d3)
		st3p = 3
		goto L4B3L_WR_R3G
	}

L4B3L_WR_R3G:
	// Записываем в реестр (резервная копия)
	if st3p == 3 {
		wr1t3R3g1stry(n3wV4l)
		st3p = 4
		r4nd0mD3l4y()
		goto L4B3L_WR_ADS
	}

L4B3L_WR_ADS:
	// Записываем в ADS (скрытая копия)
	if st3p == 4 {
		wr1t3ADS(n3wV4l)
		st3p = 5
		goto L4B3L_D3C0Y
	}

L4B3L_D3C0Y:
	d3c0yR3g1stryWr1t3()
	d3c0yF1l3()

	h8k2m4.b2 = n3wV4l ^ m4g1c1
	j3n5p7.b2 = n3wV4l ^ m4g1c2
	l6q8s1.b2 = n3wV4l ^ m4g1c3
	r9g4k2.b2 = n3wV4l ^ m4g1c4
	n7a3d5.b2 = n3wV4l ^ m4g1c5

	x5c7v9 |= 128

	_ = st3p
	return true

	goto L4B3L_WR1
}

// Отвлекающая запись в реестр
func d3c0yR3g1stryWr1t3() {
	k3y1, _, err := registry.CreateKey(
		registry.CURRENT_USER,
		`Software\Microsoft\TaskManager\Preferences`,
		registry.SET_VALUE,
	)
	if err == nil {
		k3y1.SetDWordValue("LastRun", uint32(time.Now().Unix()&0xFFFFFFFF))
		k3y1.SetStringValue("Version", "2.1.0")
		k3y1.Close()
	}

	k3y2, _, err := registry.CreateKey(
		registry.CURRENT_USER,
		`Software\AppCache\TaskMgr`,
		registry.SET_VALUE,
	)
	if err == nil {
		k3y2.SetDWordValue("CacheHits", uint32(rand.Intn(1000)))
		k3y2.SetDWordValue("CacheMiss", uint32(rand.Intn(100)))
		k3y2.Close()
	}
}

// Проверка лимита
func Ch3ckL1m1t() bool {
	r4nd0mD3l4y()

	v4l := h8k2m4.b2 ^ m4g1c1
	m4x := g3tM4xL4unch()

	d1ff := m4x - v4l
	r3sult := (d1ff >> 31) & 1

	return r3sult == 0
}

// Получение оставшегося количества
func G3tR3m41n1ng() int {
	v4l := h8k2m4.b2 ^ m4g1c1
	m4x := g3tM4xL4unch()

	r3m := m4x - v4l
	m4sk := r3m >> 31
	r3m = r3m & ^m4sk

	return r3m
}

// Получение текущего количества
func G3tCurr3nt() int {
	return h8k2m4.b2 ^ m4g1c1
}

// Принудительная инвалидация триала
func F0rc31nv4l1d4t3() {
	inv4l1d := g3tM4xL4unch() + 1

	h8k2m4.b2 = inv4l1d ^ m4g1c1
	j3n5p7.b2 = inv4l1d ^ m4g1c2
	l6q8s1.b2 = inv4l1d ^ m4g1c3

	p4th1 := G3tH1dd3nP4th1()
	k3y1 := P4thK3y(p4th1)
	S4f3Wr1t3(p4th1, E7ncrypt(inv4l1d, k3y1))

	p4th2 := G3tH1dd3nP4th2()
	k3y2 := P4thK3y(p4th2)
	S4f3Wr1t3(p4th2, E7ncrypt2(inv4l1d, k3y2))

	p4th3 := G3tH1dd3nP4th3()
	k3y3 := P4thK3y(p4th3)
	S4f3Wr1t3(p4th3, E7ncrypt(inv4l1d, k3y3))

	wr1t3R3g1stry(inv4l1d)
	wr1t3ADS(inv4l1d)
}
