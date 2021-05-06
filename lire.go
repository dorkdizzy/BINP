package main

import (
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func mp3player() {

	contents, err := getResource("alarme.mp3")
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile("alarme.mp3", contents, 0644)

	f, err := os.Open("alarme.mp3")
	if err != nil {
		log.Fatal(err)
	}
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	speaker.Play(streamer)
	select {}
}
