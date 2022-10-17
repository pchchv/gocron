package timechecker

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/pchchv/gocron/types"
)

// Check the necessity of running the command at this moment
func NeedToRunNow(element types.Task) bool {
	period := element.Period
	current := int32(time.Now().Unix())
	currentDate := time.Now().Format("2006-01-02")
	currentTime := time.Now().Format("15:04:05")
	currentWeekday := time.Now().Format("Mon")
	need := true
	// If a period is set, check if the current time is suitable for the task
	if period > 0 {
		if !checkPeriod(period, current) {
			need = false
		}
	}
	// Check if the current time is within the "sleeping" period of the task
	if need {
		for _, sTime := range element.SleepTime {
			if checkInSleepTime(sTime, currentTime) {
				need = false
			}
		}
	}
	// Check if the current date is within the "sleep" period of the task
	if need {
		for _, sDay := range element.SleepDays {
			if checkInSleepingDays(sDay, currentWeekday) {
				need = false
			}
		}
	}
	// If it is set that the task must be performed at a certain time, then check this condition
	if need && len(element.Time) > 0 {
		tempNeed := false
		for _, sTime := range element.Time {
			if checkInTime(sTime, currentTime) {
				tempNeed = true
			}
		}
		if !tempNeed {
			need = false
		}
	}
	// If it is set that the task must be performed on a certain day, then check this condition
	if need && len(element.DateTime) > 0 {
		tempNeedDateTime := false
		for _, sDateTime := range element.DateTime {
			if checkInTime(sDateTime, currentDate+" "+currentTime) {
				tempNeedDateTime = true
			}
		}
		if !tempNeedDateTime {
			need = false
		}
	}
	return need
}

// Check the period belonging to the period in seconds
func checkPeriod(period int32, current int32) bool {
	if current%period == 0 {
		return true
	}
	return false
}

// Check if the current time is within the "sleep" period for the task
func checkInSleepTime(sleeprange string, currentTime string) bool {
	sleepRange := strings.Split(sleeprange, "-")
	sleepRangeStart := toInteger(sleepRange[0])
	sleepRangeEnd := toInteger(sleepRange[1])
	currentSeconds := toInteger(currentTime)
	if currentSeconds >= sleepRangeStart && currentSeconds <= sleepRangeEnd {
		return true
	}
	return false
}

// Check if the current time is not included in the "sleeping" days of the task
func checkInSleepingDays(day string, currentDay string) bool {
	if day == currentDay {
		return true
	}
	return false
}

// Check if the current time is the moment when the task starts
func checkInTime(time string, currentTime string) bool {
	if time == currentTime {
		log.Println(time)
		return true
	}
	return false
}

// Convert time to seconds from midnight
func toInteger(time string) int {
	timeArr := strings.Split(time, ":")
	hours, err := strconv.Atoi(timeArr[0])
	if err != nil {
		log.Panic(err)
	}
	minutes, err := strconv.Atoi(timeArr[1])
	if err != nil {
		log.Panic(err)
	}
	seconds, err := strconv.Atoi(timeArr[2])
	if err != nil {
		log.Panic(err)
	}
	return hours*3600 + minutes*60 + seconds
}
