package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const DateFormat = "20060102"

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if repeat == "" {
		return "", errors.New("Нулевое чисно повторений")
	}
	tPars, err := time.Parse(DateFormat, dstart)
	if err != nil {
		return "", err
	}
	repeatS := strings.Split(repeat, " ")
	//проверка в repeat на прибавление года
	if len(repeatS) == 1 {
		if repeatS[0] != "y" {
			return "", errors.New("Неверный формат записи")
		} else {
			//дата для напоминания через год
			for {
				tPars = tPars.AddDate(1, 0, 0)
				if tPars.After(now) {
					break
				}
			}
		}
		//проверка в repeat на прибавление дней или дня
	} else if len(repeatS) == 2 {
		if repeatS[0] != "d" {
			return "", errors.New("Неверный формат записи")
		} else {
			//проверка пройдена, вычисляется колличество дней для напоминания
			repeat = repeatS[1]
			val, err := strconv.Atoi(repeatS[1])
			if err != nil {
				return "", err
			}
			if val > 400 {
				return "", errors.New("Число не может быть больше 400")
			} else {
				tPars = tPars.AddDate(0, 0, val)
			}
		}
	} else {
		return "", errors.New("Неверный формат записи")
	}
	return tPars.Format(DateFormat), nil
}
func nextDayHandler(w http.ResponseWriter, r *http.Request) {
	nowString := r.FormValue("now")
	dateString := r.FormValue("date")
	repeat := r.FormValue("repeat")
	if nowString == "" {
		nowString = time.Now().Format(DateFormat)
	}
	now, err := time.Parse(DateFormat, nowString)
	if err != nil {
		fmt.Println(err)
		return
	}
	val, err := NextDate(now, dateString, repeat)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write([]byte(val))
}
