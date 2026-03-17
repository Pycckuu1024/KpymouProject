package events

import (
	"errors"
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
	return uuid.New().String() // создает и возвращает уникальный айди
}

func isValidTitle(title string) bool {
	patern := "^[a-zA-Z0-9а-яёА-ЯЁ ]{3,50}$"
	matched, err := regexp.MatchString(patern, title)
	if err != nil {
		return false
	}
	return matched
}

func NewEvent(title string, dateStr string) (Event, error) {
	t, err := dateparse.ParseAny(dateStr)
	if err != nil {
		return Event{}, errors.New("неверный формат даты")
	}
	checkTitle := isValidTitle(title)
	if checkTitle == false {
		return Event{}, errors.New("неверное имя задачи")
	}
	return Event{
		ID:      getNextID(),
		Title:   title,
		StartAt: t,
	}, nil

}
