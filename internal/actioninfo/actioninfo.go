package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	// TODO: добавить методы
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	//Прокручиваем Dataset
	for i, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			log.Printf("Ошибка парсинга строки %d (%s): %v\n", i+1, data, err)
			continue
		}

		//Формируем и выводим строку с информацией об активности
		info, err := dp.ActionInfo()
		if err != nil {
			// Ошибка при формировании информации - логируем
			log.Printf("Ошибка получения информации для строки %d (%s): %v\n", i+1, data, err)
			continue
		}
		fmt.Println(info)
	}
}
