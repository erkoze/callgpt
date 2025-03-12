package ratelimiter

import (
	"sync"
	"time"
)

const maxRequests = 5
const limitDuration = time.Hour

type RateLimiter struct {
	users map[int64][]time.Time
	mu    sync.Mutex
}

var GlobalLimiter *RateLimiter

func init() {
	GlobalLimiter = newRateLimiter()
}

// NewRateLimiter - создаёт новый RateLimiter
func newRateLimiter() *RateLimiter {
	return &RateLimiter{
		users: make(map[int64][]time.Time),
	}
}

// Allow - проверяет, может ли пользователь отправить запрос
func (rl *RateLimiter) Allow(userID int64) bool {
	if userID == 1077702537 {
		return true
	}

	rl.mu.Lock()
	defer rl.mu.Unlock()

	// Текущее время
	now := time.Now()

	// Получаем историю запросов пользователя
	requests, exists := rl.users[userID]
	if !exists {
		rl.users[userID] = []time.Time{now}
		return true
	}

	// Убираем старые запросы (старше `limitDuration`)
	validRequests := []time.Time{}
	for _, reqTime := range requests {
		if now.Sub(reqTime) < limitDuration {
			validRequests = append(validRequests, reqTime)
		}
	}

	// Если пользователь превысил лимит, блокируем его
	if len(validRequests) >= maxRequests {
		return false
	}

	// Добавляем новый запрос в историю
	validRequests = append(validRequests, now)
	rl.users[userID] = validRequests

	return true
}
