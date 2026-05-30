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
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	//Разделяем строку на слайс строк
	parts := strings.Split(datastring, ",")
	//Проверяем длину слайса
	if len(parts) != 3 {
		return errors.New("неверный формат строки: ожидается 3 части")
	}
	//Преобразуем первый элемент слайса в int
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return errors.New("неверный формат количества шагов: " + err.Error())
	}
	if steps <= 0 {
		return errors.New("количество шагов должно быть положительным")
	}
	t.Steps = steps
	//Сохраняем поля структуры
	t.TrainingType = parts[1]
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return errors.New("неверный формат продолжительности: " + err.Error())
	}
	if duration <= 0 {
		return errors.New("продолжительность должна быть положительной")
	}
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	// Проверка, что данные о тренировке заполнены
	if t.Steps <= 0 {
		return "", errors.New("количество шагов должно быть положительным")
	}

	if t.TrainingType == "" {
		return "", errors.New("тип тренировки не указан")
	}

	if t.Duration <= 0 {
		return "", errors.New("продолжительность должна быть положительной")
	}
	// Проверка, что данные о пользователе заполнены
	if t.Weight <= 0 {
		return "", errors.New("вес пользователя не указан")
	}

	if t.Height <= 0 {
		return "", errors.New("рост пользователя не указан")
	}
	//Вычисляем дистанцию
	distance := spentenergy.Distance(t.Steps, t.Height)
	//Вычисляем среднюю скорость
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	//Расчитываем калории в зависимости от типа тренировок
	var calories float64
	var err error
	switch t.TrainingType {
	case "Ходьба", "ходьба", "Walking", "walking":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Бег", "бег", "Running", "running":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %s", t.TrainingType)
	}
	if err != nil {
		return "", fmt.Errorf("ошибка при расчёте калорий: %w", err)
	}
	//Формируем строку с результатами
	result := fmt.Sprintf("Тип тренировки: %s\n", t.TrainingType)
	result += fmt.Sprintf("Длительность: %.2f ч.\n", t.Duration.Hours())
	result += fmt.Sprintf("Дистанция: %.2f км.\n", distance)
	result += fmt.Sprintf("Скорость: %.2f км/ч\n", speed)
	result += fmt.Sprintf("Сожгли калорий: %.2f\n", calories)
	return result, nil
}
