package omikuji

import (
	"math/rand"
	"time"
)

type OmikujiType int

// 大吉・吉・中吉・小吉・末吉・凶
const (
	Daikichi OmikujiType = iota + 1
	Kichi
	Chukichi
	Shokichi
	Suekichi
	Kyo
)

type Result struct {
	Message string `json: "message"`
}

func (o OmikujiType) Result() *Result {
	var message string
	switch o {
	case Daikichi:
		message = "大吉"
	case Kichi:
		message = "吉"
	case Chukichi:
		message = "中吉"
	case Shokichi:
		message = "小吉"
	case Suekichi:
		message = "末吉"
	case Kyo:
		message = "凶"
	default:
		message = "不明"

	}
	return &Result{Message:message}
}

func init()  {
	rand.Seed(time.Now().UnixNano())
}

func getCandidates(now time.Time) []OmikujiType {
	if now.Month() == time.January && (now.Day() == 1 || now.Day() == 2 || now.Day() == 3) {
		return []OmikujiType {
			Daikichi,
		}
	} else {
		return []OmikujiType {
			Daikichi,
			Kichi,
			Chukichi,
			Shokichi,
			Suekichi,
			Kyo,
		}
	}
}

func Draw(now time.Time) *Result {
	candidates := getCandidates(now)
	index := rand.Intn(len(candidates))
	return candidates[index].Result()
}
