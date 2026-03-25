package events

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/araddon/dateparse"
	"github.com/google/uuid"
)

type Event struct {
	ID      string
	Title   string
	StartAt time.Time
}

func getNextID() string {
	return uuid.New().String()
}

func IsValidTitle(title string) bool {
	pattern := "^[a-zA-Z0-9а-яёА-ЯЁ ]{3,50}$"
	matched, err := regexp.MatchString(pattern, title)
	if err != nil {
		return false
	}
	return matched
}

func NewEvent(title string, date string) (*Event, error) {
	t, err := dateparse.ParseAny(date)
	if err != nil {
		return &Event{}, errors.New("неверный формат даты")
	}
	checkTitle := IsValidTitle(title)
	if checkTitle == false {
		return &Event{}, errors.New("неверное имя задачи")
	}
	return &Event{
		ID:      getNextID(),
		Title:   title,
		StartAt: t,
	}, nil

}

func (e *Event) Print() {
	fmt.Println(e.Title, e.StartAt)
}

func (e *Event) Update(title string, date string) error {
	d, _ := dateparse.ParseAny(date)
	e.Title = title
	e.StartAt = d
	return nil
}
