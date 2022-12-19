package main

import "time"

//function to sleep for specified amoutn of time
func Sleep(sec int){
	time.Sleep(time.Duration(sec) * time.Second);
}
