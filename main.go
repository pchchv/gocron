package main

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {}

func runCommand(command string, output string) {
	cmd := command + " >> " + output
	gocronLog("START", cmd)
	runcmd(cmd, true)
}

func runcmd(cmd string, shell bool) []byte {
	startTime := int(time.Now().Unix())
	if shell {
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			gocronLog("ERROR", cmd+"("+err.Error()+")")
		}
		finishTime := int(time.Now().Unix())
		t := strconv.Itoa(finishTime - startTime)
		gocronLog("FINISH in "+t+"s", cmd)
		return out
	}
	out, err := exec.Command(cmd).Output()
	if err != nil {
		gocronLog("ERROR", cmd+"("+err.Error()+")")

	}
	finishTime := int(time.Now().Unix())
	t := strconv.Itoa(finishTime - startTime)
	gocronLog("FINISH in "+t, cmd)
	return out
}

func gocronLog(messageType string, message string) {
	formatMessage := "[" + time.Now().String() + "]" + " " + messageType + ": " + message + "\n"
	log.Print(formatMessage)
	f, _ := os.OpenFile("./logs/gocron.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	f.WriteString(formatMessage)
	f.Close()
	if messageType == "ERROR" {
		f, _ := os.OpenFile("./logs/error.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		f.WriteString(formatMessage)
		f.Close()
	}
}
