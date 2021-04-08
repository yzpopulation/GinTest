package jobs

import "github.com/bamzi/jobrunner"

func Init() {
	jobrunner.Start()
	jobrunner.Schedule("@every 5s", ReminderEmails{})
}
