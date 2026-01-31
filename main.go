// Менеджер Задач - Триал версия
// Программа для управления списком задач с ограничением на 4 запуска
package main

import (
	"trialtaskmanager/app"
	"trialtaskmanager/internal"
)

func main() {
	// Инициализация защиты ДО создания UI
	internal.In1t14l1z3()
	internal.V3r1fyC0ns1st3ncy()

	// Создаём и запускаем GUI
	ui := app.NewAppUI()
	ui.Run()
}
