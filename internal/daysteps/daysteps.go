package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps                 int           // кол-во шагов
	Duration              time.Duration // длительность прогулки
	personaldata.Personal               // перс-ые данные
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	arr := strings.Split(datastring, ",")
	if len(arr) != 2 {
		return errors.New("ошибка парсинга данных, некорректное число элементов после парсинга")
	}
	steps, err := strconv.Atoi(arr[0])
	if steps <= 0 {
		return errors.New("ошибка парсинга шагов")
	}
	if err != nil {
		return fmt.Errorf("некорректный формат шагов: %w", err)
	}
	duration, err := time.ParseDuration(arr[1])
	if duration <= 0 {
		return errors.New("время не может быть отрицательнысс")
	}
	if err != nil {
		return fmt.Errorf("некорректный формат времени: %w", err)
	}

	ds.Steps = steps
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {

	distance := spentenergy.Distance(ds.Steps, ds.Height)
	energy, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", errors.New("ошибка парсинга")
	}
	result := fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, energy,
	)

	return result, nil
}
