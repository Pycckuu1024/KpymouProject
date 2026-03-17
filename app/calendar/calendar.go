package calendar

import (
	"errors"

	"fmt"

	"github.com/Pycckuu1024/KpymouProject/app/events"
)

var eventsMap = make(map[string]events.Event)

func AddEvent(title string, date string) (events.Event, error) {
	e, err := events.NewEvent(title, date)
	if err != nil {
		return events.Event{}, errors.New("неверное имя задачи")
	}
	fmt.Printf("\nДобавлено событие - ID : %s, Имя : %s, Дата : %s\n", e.ID, e.Title, e.StartAt)
	eventsMap[e.ID] = e

	return e, nil
}

func EditEvent(key string, e events.Event) {
	if _, ok := eventsMap[key]; ok {
		fmt.Printf("\nИзменено событие : %s", eventsMap[key])
		eventsMap[key] = e
		fmt.Printf(" на : %s\n", eventsMap[key])
	} else {
		fmt.Printf("\nНет события с таким именем, изменения невозожны !\n")
		fmt.Printf("\nНет события с таким именем, изменения невозожны !\n")

	}
}
func DeleteEvent(key string) {
	fmt.Printf("\nУдалено Событие : %s\n", eventsMap[key])
	delete(eventsMap, key)
}
func ShowEvents() {
	for _, v := range eventsMap {
		fmt.Printf("\n ID : %s  Событие : %s.  Дата : %v.\n", v.ID, v.Title, v.StartAt)
	}
}
