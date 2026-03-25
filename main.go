package main

import (
	"fmt"

	"github.com/Pycckuu1024/KpymouProject/app/calendar"
)

//	func main() {
//		event1, err1 := calendar.AddEvent("Встреча", "2025/06/12 16:33")
//		if err1 != nil {
//			fmt.Println("Ошибка:", err1)
//			return
//		}
//
//		event2, err2 := calendar.AddEvent("Еще одна встреча", "2025/06/12 15:00")
//		if err2 != nil {
//			fmt.Println("Ошибка:", err2)
//			return
//		}
//
//		calendar.ShowEvents()
//		calendar.DeleteEvent(event1.ID)
//
//		calendar.EditEvent(event2.ID, "Заменённое событие", "1111/11/11 11:11")
//		calendar.ShowEvents()
//	}
func main() {
	c := calendar.NewCalendar()

	event1, err1 := c.AddEvent("Встреча номер 1", "2025/06/12")
	if err1 != nil {
		fmt.Println("Ошибка:", err1)
	} else {
		fmt.Println(event1.Title, "добавлено")
	}

	event2, err2 := c.AddEvent("Встреча номер 2", "2025/06/12")
	if err2 != nil {
		fmt.Println("Ошибка:", err2)
	} else {
		fmt.Println(event2.Title, "добавлено")
	}

	err := c.EditEvent(event2.ID, "Созвон", "2025/06/12 16:50")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Событие обновлено")
	}

	c.ShowEvents()
}
