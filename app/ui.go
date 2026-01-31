// Package app - графический интерфейс приложения
// Использует библиотеку Fyne для создания кроссплатформенного GUI
package app

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"trialtaskmanager/internal"
)

// AppUI - структура графического интерфейса
type AppUI struct {
	fyneApp     fyne.App
	mainWindow  fyne.Window
	taskManager *TaskManager
	taskList    *widget.List
	statusLabel *widget.Label
	statsLabel  *widget.Label

	// Флаги для защиты
	firstAction bool // Первое действие выполнено
	trialActive bool // Триал активен
}

// NewAppUI создаёт новый интерфейс приложения
func NewAppUI() *AppUI {
	ui := &AppUI{
		fyneApp:     app.New(),
		taskManager: NewTaskManager(),
		firstAction: false,
		trialActive: true,
	}

	ui.mainWindow = ui.fyneApp.NewWindow("Менеджер Задач - Триал версия")
	ui.mainWindow.Resize(fyne.NewSize(700, 500))

	// Проверяем триал при создании
	ui.trialActive = internal.Ch3ckL1m1t()

	return ui
}

// Run запускает приложение
func (ui *AppUI) Run() {
	ui.createUI()
	ui.mainWindow.ShowAndRun()
}

// createUI создаёт элементы интерфейса
func (ui *AppUI) createUI() {
	// Статус триала
	ui.statusLabel = widget.NewLabel("")
	ui.updateTrialStatus()

	// Статистика
	ui.statsLabel = widget.NewLabel("")
	ui.updateStats()

	// Список задач
	ui.taskList = widget.NewList(
		func() int {
			return len(ui.taskManager.GetAllTasks())
		},
		func() fyne.CanvasObject {
			return ui.createTaskItem()
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			ui.updateTaskItem(id, obj)
		},
	)

	// Кнопки управления
	addButton := widget.NewButtonWithIcon("Добавить задачу", theme.ContentAddIcon(), func() {
		ui.showAddTaskDialog()
	})

	clearButton := widget.NewButtonWithIcon("Очистить всё", theme.DeleteIcon(), func() {
		ui.showClearConfirmDialog()
	})

	// Панель инструментов
	toolbar := container.NewHBox(
		addButton,
		widget.NewSeparator(),
		clearButton,
	)

	// Компоновка
	header := container.NewVBox(
		ui.statusLabel,
		widget.NewSeparator(),
		toolbar,
		container.NewHBox(widget.NewLabel("Задачи:"), widget.NewLabel(" "), ui.statsLabel),
	)

	content := container.NewBorder(
		header, // top
		nil,    // bottom
		nil,    // left
		nil,    // right
		ui.taskList,
	)

	ui.mainWindow.SetContent(content)
}

// updateTrialStatus обновляет статус триала
func (ui *AppUI) updateTrialStatus() {
	remaining := internal.G3tR3m41n1ng()
	current := internal.G3tCurr3nt()

	if remaining > 0 {
		ui.statusLabel.SetText(fmt.Sprintf(
			"Триал-версия | Использовано запусков: %d из 4 | Осталось: %d",
			current, remaining,
		))
	} else {
		ui.statusLabel.SetText("Триал-период истёк! Приобретите полную версию.")
	}
}

// updateStats обновляет статистику
func (ui *AppUI) updateStats() {
	total := ui.taskManager.GetTaskCount()
	completed := ui.taskManager.GetCompletedCount()
	ui.statsLabel.SetText(fmt.Sprintf("Всего: %d | Выполнено: %d", total, completed))
}

// onAction вызывается при действии пользователя
// Инкрементирует счётчик при первом действии
func (ui *AppUI) onAction() bool {
	if !ui.firstAction {
		ui.firstAction = true
		internal.Incr3m3ntC0unt3r()
		ui.trialActive = internal.Ch3ckL1m1t()
		ui.updateTrialStatus()
	}

	if !ui.trialActive {
		ui.showTrialExpiredDialog()
		return false
	}
	return true
}

// showTrialExpiredDialog показывает сообщение об истечении триала
func (ui *AppUI) showTrialExpiredDialog() {
	dialog.ShowInformation(
		"Триал-период истёк",
		"Вы исчерпали лимит бесплатных запусков (4 запуска).\n\n"+
			"Для продолжения использования приобретите полную версию программы.",
		ui.mainWindow,
	)
}

// createTaskItem создаёт элемент списка задач
func (ui *AppUI) createTaskItem() fyne.CanvasObject {
	titleLabel := widget.NewLabel("Title")
	titleLabel.TextStyle.Bold = true

	descLabel := widget.NewLabel("Description")
	descLabel.Wrapping = fyne.TextWrapWord

	priorityLabel := widget.NewLabel("[Priority]")
	statusLabel := widget.NewLabel("")

	checkButton := widget.NewButtonWithIcon("", theme.ConfirmIcon(), nil)
	deleteButton := widget.NewButtonWithIcon("", theme.DeleteIcon(), nil)

	buttons := container.NewHBox(checkButton, deleteButton)

	info := container.NewVBox(
		container.NewHBox(titleLabel, priorityLabel, statusLabel),
		descLabel,
	)

	return container.NewBorder(nil, nil, nil, buttons, info)
}

// updateTaskItem обновляет элемент списка
func (ui *AppUI) updateTaskItem(id widget.ListItemID, obj fyne.CanvasObject) {
	tasks := ui.taskManager.GetAllTasks()
	if id >= len(tasks) {
		return
	}

	task := tasks[id]

	// Получаем элементы контейнера
	border := obj.(*fyne.Container)
	info := border.Objects[0].(*fyne.Container)
	buttons := border.Objects[1].(*fyne.Container)

	// Заголовок и информация
	headerBox := info.Objects[0].(*fyne.Container)
	titleLabel := headerBox.Objects[0].(*widget.Label)
	priorityLabel := headerBox.Objects[1].(*widget.Label)
	statusLabel := headerBox.Objects[2].(*widget.Label)

	descLabel := info.Objects[1].(*widget.Label)

	// Кнопки
	checkButton := buttons.Objects[0].(*widget.Button)
	deleteButton := buttons.Objects[1].(*widget.Button)

	// Устанавливаем значения
	titleLabel.SetText(task.Title)
	descLabel.SetText(task.Description)

	// Приоритет
	switch task.Priority {
	case PriorityHigh:
		priorityLabel.SetText("[!!! Высокий]")
	case PriorityMedium:
		priorityLabel.SetText("[!! Средний]")
	default:
		priorityLabel.SetText("[! Низкий]")
	}

	// Статус выполнения
	if task.Completed {
		statusLabel.SetText(" [Выполнено]")
		checkButton.SetIcon(theme.CheckButtonCheckedIcon())
	} else {
		statusLabel.SetText("")
		checkButton.SetIcon(theme.CheckButtonIcon())
	}

	// Обработчики кнопок
	taskID := task.ID
	checkButton.OnTapped = func() {
		if ui.onAction() {
			ui.taskManager.ToggleComplete(taskID)
			ui.taskList.Refresh()
			ui.updateStats()
		}
	}

	deleteButton.OnTapped = func() {
		if ui.onAction() {
			ui.taskManager.DeleteTask(taskID)
			ui.taskList.Refresh()
			ui.updateStats()
		}
	}
}

// showAddTaskDialog показывает диалог добавления задачи
func (ui *AppUI) showAddTaskDialog() {
	if !ui.onAction() {
		return
	}

	titleEntry := widget.NewEntry()
	titleEntry.SetPlaceHolder("Название задачи")

	descEntry := widget.NewMultiLineEntry()
	descEntry.SetPlaceHolder("Описание задачи")
	descEntry.SetMinRowsVisible(3)

	prioritySelect := widget.NewSelect(
		[]string{"Низкий", "Средний", "Высокий"},
		nil,
	)
	prioritySelect.SetSelected("Средний")

	form := widget.NewForm(
		widget.NewFormItem("Название", titleEntry),
		widget.NewFormItem("Описание", descEntry),
		widget.NewFormItem("Приоритет", prioritySelect),
	)

	dialog.ShowCustomConfirm(
		"Добавить задачу",
		"Добавить",
		"Отмена",
		form,
		func(confirm bool) {
			if confirm && titleEntry.Text != "" {
				var priority Priority
				switch prioritySelect.Selected {
				case "Высокий":
					priority = PriorityHigh
				case "Средний":
					priority = PriorityMedium
				default:
					priority = PriorityLow
				}
				ui.taskManager.AddTask(titleEntry.Text, descEntry.Text, priority)
				ui.taskList.Refresh()
				ui.updateStats()
			}
		},
		ui.mainWindow,
	)
}

// showClearConfirmDialog показывает диалог подтверждения очистки
func (ui *AppUI) showClearConfirmDialog() {
	if !ui.onAction() {
		return
	}

	dialog.ShowConfirm(
		"Подтверждение",
		"Вы уверены, что хотите удалить все задачи?",
		func(confirm bool) {
			if confirm {
				ui.taskManager.ClearAllTasks()
				ui.taskList.Refresh()
				ui.updateStats()
			}
		},
		ui.mainWindow,
	)
}

// GetWindow возвращает главное окно
func (ui *AppUI) GetWindow() fyne.Window {
	return ui.mainWindow
}
