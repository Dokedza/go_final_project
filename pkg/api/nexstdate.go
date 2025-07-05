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
	if repeatS[0] == "y" {
		for {
			tPars = tPars.AddDate(1, 0, 0)
			if tPars.After(now) {
				break
			}
		}
	} else if repeatS[0] == "d" {
		if len(repeatS) != 2 {
			return "", errors.New("неподдерживаемый формат")
		}
		dayCount, err := strconv.Atoi(repeatS[1])
		if err != nil {
			return "", err
		}

		if dayCount <= 0 || dayCount > 400 {
			return "", errors.New("Указанное число находится в неподдерживаемом интервале")
		}
		for {

			tPars = tPars.AddDate(0, 0, dayCount)
			if tPars.After(now) {
				break
			}
		}

	} else {
		return "", errors.New("неподдерживаемый формат")
	}
	return tPars.Format(DateFormat), nil
}
func nextDayHandler(w http.ResponseWriter, r *http.Request) {
	nowString := r.FormValue("now")
	dateString := r.FormValue("date")
	repeat := r.FormValue("repeat")

	var now time.Time
	var err error

	if nowString == "" {
		now = time.Now()
	} else {
		now, err = time.Parse(DateFormat, nowString)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	val, err := NextDate(now, dateString, repeat)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write([]byte(val))
}
