package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int           // количество шагов, проделанных за тренировку
	TrainingType string        // тип тренировки(бег или ходьба)
	Duration     time.Duration // длительность тренировки
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")

	if len(parts) != 3 {
		return errors.New("ошибка парсинга")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("шаги не могут быть отрицательными")
	}
	duration, err := time.ParseDuration(parts[2])

	if err != nil {
		return fmt.Errorf("некорректный формат времени: %w", err)
	}
	if duration <= 0 {
		return errors.New("длительность не должна быть равно нулю или меньше его")
	}

	t.Steps = steps
	t.TrainingType = parts[1]
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {

	var energy float64
	switch t.TrainingType {
	case "Ходьба":
		value, err := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("ошибка вычисления калорий ходьбы: %w", err)
		}
		energy = value
	case "Бег":
		value, err := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("ошибка вычисления калорий бега: %w", err)
		}
		energy = value
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	result := fmt.Sprintf(
		"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		spentenergy.Distance(t.Steps, t.Height),
		spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration),
		energy,
	)
	return result, nil
}
