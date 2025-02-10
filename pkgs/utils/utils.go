package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// setInterval имитирует поведение setInterval из JavaScript
func SetInterval(callback func(), interval time.Duration) chan bool {
	ticker := time.NewTicker(interval) // Создаём тикер с заданным интервалом
	stop := make(chan bool)            // Канал для остановки интервала

	go func() {
		for {
			select {
			case <-ticker.C: // Срабатывает каждые `interval` времени
				callback() // Выполняем переданную функцию
			case <-stop: // Если получен сигнал остановки
				ticker.Stop() // Останавливаем тикер
				return        // Выходим из горутины
			}
		}
	}()

	return stop // Возвращаем канал для остановки
}

func GenerateNewId() string {
	rand.Seed(time.Now().UnixNano())
	// Используем максимальное значение для int64
	maxInt := int64(^uint64(0) >> 1)        // Это эквивалент Number.MAX_SAFE_INTEGER
	randomNumber := rand.Int63n(maxInt) + 1 // Генерация числа от 1 до maxInt
	return fmt.Sprintf("%d", randomNumber)  // Преобразуем число в строку
}
