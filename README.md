<div align="center">

# **gocron — minimalistic solution for running tasks**

</div>

# **Running:**

## 1. You must have the current version of Go installed, and set $GOPATH

## 2. Enter the directory where we want to install the script `cd /home`

## 3. Clone the repository `git clone https://github.com/pchchv/gocron`

## 4. `cd gocron`

## 5. The tasks are described in the file cronjob.json

## 6. After describing the tasks we want `go build main.go

## 7. ./main &

#

## **Task description:**

## The simplest task

### In this case we will call the command `ping -c 4 google.com >> ./logs.txt` every 15 seconds

```json
{
    "Tasks":[
        {
            "Period":15,
            "Command":"ping -c 4 google.com",
            "Output":"./logs.txt"
        }
    ]
}   
```

## Let's complicate the example. Let's add periods when the script will not be called

### This script will call the command every 84 seconds, between 1 p.m. and 7 p.m., every day except Monday

```json
{
       "Period": 84,
       "SleepTime": [
         "19:00:00-23:59:59",
         "00:00:00-13:00:00"
       ],
       "SleepDays": [
         "Mon"
       ],
       "Command": "ping -c 4 google.com",
       "Output": "./logs2.txt"
}
```

## It is also possible to call the command at a strictly specified time

### This script will run twice a day, at a strictly specified time

```json
{
    "Time": [
        "14:00:00",
        "14:30:00",
    ],
    "Command": "ping -c 4 google.com",
    "Output": "./logs3.txt"
}
```

## You can also set a specific date and time to run the script

```json
{
       "DateTime": [
         "2018-05-13 14:07:22",
         "2018-05-13 13:55:10",
         "2018-05-14 15:00:00",
         "2018-05-15 17:00:00"
       ],
       "Command": "ping -c 4 google.com",
       "Output": "./logs4.txt"
}
 ```
