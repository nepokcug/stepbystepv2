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
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	//Разделяем строку на слайс строк
	parts := strings.Split(datastring, ",")
	//Проверяем длину слайса
	if len(parts) != 2 {
		return errors.New("неверный формат строки: ожидается 2 части")
	}
	//Преобразуем первый элемент слайса в int
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return errors.New("неверный формат количества шагов: " + err.Error())
	}
	if steps <= 0 {
		return errors.New("количество шагов должно быть положительным")
	}
	ds.Steps = steps
	//Записываем длительность
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return errors.New("неверный формат продолжительности: " + err.Error())
	}
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	// Проверяем, что данные о шагах корректны
	if ds.Steps <= 0 {
		return "", errors.New("количество шагов должно быть положительным")
	}

	// Проверяем, что длительность корректна
	if ds.Duration <= 0 {
		return "", errors.New("продолжительность должна быть положительной")
	}

	// Проверяем, что данные о пользователе заполнены
	if ds.Weight <= 0 {
		return "", errors.New("вес пользователя не указан")
	}

	if ds.Height <= 0 {
		return "", errors.New("рост пользователя не указан")
	}
	//Вычисляем дистанцию
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	//Вычисляем количество калорий
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("ошибка при расчёте калорий: %w", err)
	}
	//Формируем строку с информацией
	result := fmt.Sprintf("Количество шагов: %d.\n", ds.Steps)
	result += fmt.Sprintf("Дистанция составила %.2f км.\n", distance)
	result += fmt.Sprintf("Вы сожгли %.2f ккал.\n", calories)

	return result, nil
}
