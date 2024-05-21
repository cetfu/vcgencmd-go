package main

import (
	"errors"
	"log"
	"os/exec"
	"slices"
	"strconv"
	"strings"
)

func doCommand(command []string) string {
	out, err := exec.Command("bash", "-c", strings.Join(command, " ")).Output()
	if err != nil {
		log.Fatal("Command run error: " + err.Error())
	}
	return strings.Replace(string(out), "\n", "", -1)
}

func lookup(command string, srcList []string, src string) string {
	if !slices.Contains(srcList, src) {
		log.Fatal("Invalid argument error: " + src + " not in " + strings.Join(srcList, ", "))
	}
	var cmd []string

	cmd = append(cmd, "vcgencmd", command, src)
	output := doCommand(cmd)
	return output
}

func MeasureTemp() (float64, error) {
	output := lookup("measure_temp", []string{""}, "")
	temp := strings.Split(output, "=")[1]
	temp = strings.Replace(temp, "'C", "", -1)
	floatTemp, err := strconv.ParseFloat(temp, 64)
	if err != nil {
		return 0, errors.New("temp convert problem")
	}
	return floatTemp, nil
}

var kMemorySources = []string{"arm", "gpu"}

func MemorySources() []string {
	return kMemorySources
}

func GetMemory(src string) (int, error) {
	output := lookup("get_mem", kMemorySources, src)
	mem := strings.Split(output, "=")[1]
	num, _ := strconv.Atoi(mem[:len(mem)-1])
	lastChar := mem[len(mem)-1:]
	if lastChar == "M" {
		return num * 1024 * 1024, nil
	} else if lastChar == "G" {
		return num * 1024 * 1024, nil
	}
	return 0, errors.New("memory problem")
}

var kCodecSources = []string{"h264", "mpg2", "wvc1", "mpg4", "mjpg", "wmv9", "hevc"}

func CodecSources() []string {
	return kCodecSources
}

func CodecEnabled(src string) bool {
	output := lookup("codec_enabled", kCodecSources, src)
	status := strings.Split(output, "=")[1]
	if status == "enabled" {
		return true
	} else {
		return false
	}
}

var kVoltageSources = []string{"core", "sdram_c", "sdram_i", "sdram_p"}

func GetVoltageSources() []string {
	return kVoltageSources
}

func measureVolts(src string) (float64, error) {
	output := lookup("measure_volts", kVoltageSources, src)
	volt := strings.Split(output, "=")[1]

	volt = strings.Replace(volt, "V", "", -1)
	float, err := strconv.ParseFloat(volt, 64)
	if err != nil {
		return 0, errors.New("measure voltage problem")
	}
	return float, nil
}

var kFreqSources = []string{"arm", "core", "h264", "isp", "v3d", "uart", "pwm", "emmc",
	"pixel", "vec", "hdmi", "dpi"}

func GetFrequencySources() []string {
	return kFreqSources
}

func MeasureClock(src string) (int, error) {
	output := lookup("measure_clock", kFreqSources, src)
	value := strings.Split(output, "=")[1]
	freq, err := strconv.Atoi(value)

	if err != nil {
		return 0, errors.New("freq convert error")
	}

	return freq, nil
}

func main() {
	/*
		func main() {
			if temp, err := MeasureTemp(); err == nil {
				log.Println("Temperature: " + strconv.FormatFloat(temp, 'g', -1, 64))
			}

			if memory, err := GetMemory("gpu"); err == nil {
				log.Println("GPU Memory: " + strconv.Itoa(memory))
			}
			if memory, err := GetMemory("arm"); err == nil {
				log.Println("ARM Memory: " + strconv.Itoa(memory))
			}
			isSupportsHevc := CodecEnabled("hevc")
			log.Println("hevc: ", isSupportsHevc)
			isSupportsH264 := CodecEnabled("h264")
			log.Println("h264: ", isSupportsH264)
			if volts, err := measureVolts("core"); err == nil {
				log.Println("Core Volts: " + strconv.FormatFloat(volts, 'g', -1, 64))
			}

			if clock, err := MeasureClock("core"); err == nil {
				log.Println("Core freq: " + strconv.Itoa(clock))
			}
		}
	*/
}
