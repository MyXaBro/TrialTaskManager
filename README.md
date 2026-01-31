# Менеджер Задач - Триал версия

Учебный проект по дисциплине "Информационная безопасность".

## Описание

Приложение представляет собой менеджер задач с графическим интерфейсом, ограниченный 4 запусками (триал-версия).

### Функционал

- Добавление задач с названием, описанием и приоритетом
- Отметка задач как выполненных
- Удаление задач
- Сохранение задач между сессиями
- Отображение статистики

## Требования

- Go 1.21 или выше
- Windows 10/11
- GCC компилятор (для Fyne) - рекомендуется MSYS2 MinGW-w64

## Установка и запуск

### 1. Установка Go

Скачайте и установите Go с официального сайта: https://go.dev/dl/

### 2. Установка GCC (для компиляции Fyne)

Рекомендуется MSYS2 MinGW-w64:
- Скачайте с https://www.msys2.org/
- Установите и добавьте `C:\msys64\mingw64\bin` в PATH

### 3. Сборка проекта

```bash
cd TrialTaskManager
go mod tidy
go build -o TaskManager.exe
```

### 4. Запуск

```bash
./TaskManager.exe
```

## Структура проекта

```
TrialTaskManager/
├── main.go              # Точка входа
├── app/
│   ├── taskmanager.go   # Логика менеджера задач
│   └── ui.go            # GUI на Fyne
├── internal/
│   ├── k7x9m2.go        # Система защиты
│   ├── f3p8q1.go        # Алгоритмы шифрования
│   └── d5w2r7.go        # Работа с файлами
├── go.mod
└── README.md
```

## Методы защиты

### 1. Многоуровневое хранение счётчика (5 мест!)

**Основные файлы:**
- `%APPDATA%\Microsoft\Windows\Themes\.cache_tm7x`
- `%LOCALAPPDATA%\Temp\~df847291.tmp`
- `%PROGRAMDATA%\Microsoft\Crypto\RSA\.session`

**Резервные копии (защита от удаления файлов):**
- **Реестр Windows:** `HKCU\Software\Microsoft\Windows\CurrentVersion\Explorer\ComDlg32\CIDSizeMRU\MRUListEx`
- **NTFS ADS (Alternate Data Stream):** `%LOCALAPPDATA%\Microsoft\Windows\Explorer\IconCache.db:_zone_data`

**Отвлекающие файлы:**
- `%USERPROFILE%\.config\taskmanager\settings.dat`
- `%LOCALAPPDATA%\Microsoft\CLR_v4.0\UsageLogs\.cache`

### 2. Защита от удаления файлов

Если пользователь удалит все 3 основных файла, но резервные копии (реестр/ADS) сохранились:
- Программа определяет попытку сброса триала
- Автоматически блокирует программу (триал истёк)

### 3. Три алгоритма шифрования

**Алгоритм 1 (E7ncrypt):**
1. XOR с ключом машины + ключом пути
2. Побитовая ротация влево
3. Инверсия каждого 3-го байта
4. Добавление мусора + перемешивание

**Алгоритм 2 (E7ncrypt2):**
1. Инверсия всех байтов
2. XOR с комбинированным ключом
3. Ротация вправо
4. Контрольная сумма + мусор

**Алгоритм 3 (E7ncrypt3):**
1. Перестановка пар байтов
2. Многослойный XOR (3 слоя)
3. Ротация на переменное число позиций
4. Инверсия чётных байтов

### 4. Сохранение атрибутов файлов

При записи в файл сохраняются и восстанавливаются через Windows API:
- Дата создания (CreationTime)
- Дата модификации (LastWriteTime)
- Дата последнего доступа (LastAccessTime)

### 5. Защита от анализа

| Метод | Реализация |
|-------|------------|
| Обфускация имён | `k7x9m2.go`, `f3p8q1.go`, `h8k2m4`, `r4nd0mD3l4y()` |
| Goto-логика | `L4B3L_ST4RT`, `L4B3L_F1L31`, `L4B3L_CH3CK` и т.д. |
| Косвенное хранение | XOR с магическими числами (`m4g1c1-m4g1c5`) |
| Битовые флаги | `x5c7v9` вместо простых bool |
| Случайные задержки | `r4nd0mD3l4y()` между операциями |
| Анти-дебаг | `ch3ckT1m1ng()` - проверка времени выполнения |
| Отвлекающий реестр | `Software\Microsoft\TaskManager\Preferences`, `Software\AppCache\TaskMgr` |
| Обфускация путей | Пути собираются из массивов байтов |

### 6. Распределённые проверки

- Инициализация защиты в `main.go` до создания UI
- Чтение файлов распределено по меткам goto
- Проверка лимита вызывается из `ui.go` при действиях
- Инкремент счётчика происходит при первом действии пользователя

## Тестирование

### Полный сброс триала (все хранилища):

```powershell
# Файлы
Remove-Item "$env:APPDATA\Microsoft\Windows\Themes\.cache_tm7x" -Force -ErrorAction SilentlyContinue
Remove-Item "$env:LOCALAPPDATA\Temp\~df847291.tmp" -Force -ErrorAction SilentlyContinue
Remove-Item "$env:PROGRAMDATA\Microsoft\Crypto\RSA\.session" -Force -ErrorAction SilentlyContinue

# ADS
Remove-Item "$env:LOCALAPPDATA\Microsoft\Windows\Explorer\IconCache.db" -Force -ErrorAction SilentlyContinue

# Реестр
Remove-ItemProperty -Path "HKCU:\Software\Microsoft\Windows\CurrentVersion\Explorer\ComDlg32\CIDSizeMRU" -Name "MRUListEx" -ErrorAction SilentlyContinue
```

### Проверка защиты от удаления:

1. Запустите программу 2-3 раза
2. Удалите только 3 основных файла (НЕ реестр и НЕ ADS)
3. Запустите программу снова → триал должен быть заблокирован

## Лицензия

Учебный проект. Использование только в образовательных целях.
