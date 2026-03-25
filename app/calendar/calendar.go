package calendar

import (
	"errors"

	"fmt"

	"github.com/Pycckuu1024/KpymouProject/app/events"
)

type Calendar struct {
	eventsMap map[string]*events.Event
}

var eventsMap = make(map[string]*events.Event)

func NewCalendar() *Calendar {
	return &Calendar{}
}

func (*Calendar) AddEvent(title string, date string) (*events.Event, error) {
	e, err := events.NewEvent(title, date)
	if err != nil {
		return &events.Event{}, errors.New("неверное имя задачи")
	}
	fmt.Printf("\nДобавлено событие - ID : %s, Имя : %s, Дата : %s\n", e.ID, e.Title, e.StartAt)
	eventsMap[e.ID] = e

	return e, nil
}

func (*Calendar) EditEvent(id string, title string, date string) error {
	e, exists := eventsMap[id]
	if !exists {
		return fmt.Errorf("ивент с таким ключом %q не найден", id)
	}
	fmt.Printf("\nОбновлено событие '%s' на '%s'\n ", e.Title, title)
	err := e.Update(title, date)
	return err
}

func (*Calendar) DeleteEvent(key string) {
	fmt.Printf("\nУдалено Событие : %s\n", eventsMap[key])
	delete(eventsMap, key)
}
func (*Calendar) ShowEvents() {
	for _, v := range eventsMap {
		fmt.Printf("\n ID : %s  Событие : %s.  Дата : %v.\n", v.ID, v.Title, v.StartAt)
	}
}
