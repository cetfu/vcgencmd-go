# A basically GoLang port of https://github.com/nicmcd/vcgencmd/

```go
package main

import (
	"log"
	"strconv"
)

func main() {

	isSupportsHevc := CodecEnabled("hevc")
	log.Println("hevc: ", isSupportsHevc)
	isSupportsH264 := CodecEnabled("h264")
	log.Println("h264: ", isSupportsH264)
	
	if temp, err := MeasureTemp(); err == nil {
		log.Println("Temperature: " + strconv.FormatFloat(temp, 'g', -1, 64))
	}
	if memory, err := GetMemory("gpu"); err == nil {
		log.Println("GPU Memory: " + strconv.Itoa(memory))
	}
	if memory, err := GetMemory("arm"); err == nil {
		log.Println("ARM Memory: " + strconv.Itoa(memory))
	}
	if volts, err := measureVolts("core"); err == nil {
		log.Println("Core Volts: " + strconv.FormatFloat(volts, 'g', -1, 64))
	}
	if clock, err := MeasureClock("core"); err == nil {
		log.Println("Core freq: " + strconv.Itoa(clock))
	}
}

```