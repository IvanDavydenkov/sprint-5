package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || height <= 0 || weight <= 0 || duration <= 0 {
		return float64(0), errors.New("некорректный формат")
	}

	speed := MeanSpeed(steps, height, duration)

	durationInMinutes := duration.Minutes()

	return (weight * speed * durationInMinutes) / minInH * walkingCaloriesCoefficient, nil
}
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || height <= 0 || weight <= 0 || duration <= 0 {
		return float64(0), errors.New("некорректный формат")
	}
	speed := MeanSpeed(steps, height, duration)

	durationInMinutes := duration.Minutes()

	return (weight * speed * durationInMinutes) / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return float64(0)
	}
	d := Distance(steps, height)
	return d / duration.Hours()
}
func Distance(steps int, height float64) float64 {
	return height * stepLengthCoefficient * float64(steps) / mInKm
}
