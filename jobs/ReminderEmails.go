package jobs

import (
	"fmt"
	"github.com/colt3k/utils/file"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"os"
	"time"
)

type ReminderEmails struct {
	// filtered
}
func (e ReminderEmails) Run() {
	// Queries the DB
	// Sends some email
	mp3FileName:="FilePath"
	if !file.Available(mp3FileName){
		return
	}
	f, err := os.Open(mp3FileName)
	defer f.Close()
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		fmt.Printf(err.Error())
	}
	defer streamer.Close()
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))
	for {
		select {
		case <-done:
			return
		case <-time.After(time.Second):
			speaker.Lock()
			fmt.Println(format.SampleRate.D(streamer.Position()).Round(time.Second))
			speaker.Unlock()
		}
	}
	fmt.Printf("Every 5 sec send reminder emails \n")
}
