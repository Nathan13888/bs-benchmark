package main

import "time"

func GetShortTime() string {
	return time.Now().Format("20060102150405")
}
