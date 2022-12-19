package main

import (
	"io"
	"os"

	"github.com/hajimehoshi/oto"

	"github.com/hajimehoshi/go-mp3"
)

var AudioFile string

// function to play a sound
func Play_sound(AudioFile string) error {
	f, err := os.Open(AudioFile)
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		return err
	}

	c, err := oto.NewContext(d.SampleRate(), 2, 2, 8192)
	if err != nil {
		return err
	}
	defer c.Close()

	p := c.NewPlayer()
	defer p.Close()

	if _, err := io.Copy(p, d); err != nil {
		return err
	}
	return nil
}

func Char_hello_sound(char string) error {
	if char == "The Archer" {
		AudioFile = ".\\sounds\\lets-frickin-go-man.mp3"
	} else if char == "Dark Knight" {
		AudioFile = ".\\sounds\\lets-do-this.mp3"
	} else if char == "Barbara" {
		AudioFile = ".\\sounds\\okayyy-lets-go.mp3"
	}
	Play_sound(AudioFile)

	return nil
}

func Start_game_sound() {
	Play_sound(".\\sounds\\are-you-ready.mp3")
}

func Alarm_sound() {
	Play_sound(".\\sounds\\alarm-clock.mp3")
}

func Yawn_sound() {
	Play_sound(".\\sounds\\yawn.mp3")
}

func Bird_sound() {
	Play_sound(".\\sounds\\polly-parrot.mp3")
}

func Ouch_sound() {
	if char == "The Archer" {
		Play_sound(".\\sounds\\ouch-male.mp3")
	} else if char == "Dark Knight" {
		Play_sound(".\\sounds\\ouch.mp3")
	} else if char == "Barbara" {
		Play_sound(".\\sounds\\ouch-girl.mp3")
	}
}

func Bite_sound() {
	Play_sound(".\\sounds\\bite.mp3")
}

func Crowd_shocked_sound() {
	Play_sound(".\\sounds\\crowd-shocked.mp3")
}

func Dead_sound() {
	Play_sound(".\\sounds\\undertaker.mp3")
}

func Bells_sound() {
	Play_sound(".\\sounds\\bells.mp3")
}
