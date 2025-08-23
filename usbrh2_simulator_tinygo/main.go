// digital output singal 4ch for xiao-rp2040
package main

import (
	"bufio"
	"fmt"
	"machine"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {

	ver_string := "usbrh2, 1.0"
	ver_serial := "123456789"

	led_blue := machine.LED_BLUE
	led_blue.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})
	d0 := machine.D0
	d0.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})
	d0.Set(false)

	d1 := machine.D1
	d1.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})
	d1.Set(false)

	time.Sleep(2 * time.Second) // wait for USB-CDC init

	var cmd string
	var param []string
	for {
		reader := bufio.NewReader(os.Stdin)
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		//		fmt.Println(line)

		if len(line) > 0 {
			line2 := string(line)
			line2 = strings.Replace(line2, "=", " ", 1) // LED[12]=on/off/0/1 replace "=" to " "
			arr1 := strings.SplitN(line2, " ", 2)       // 分割数を指定して、先頭とそれ以降の2分割に制限
			cmd = strings.TrimSpace(arr1[0])
			if len(arr1) > 1 {
				param = strings.Split(arr1[1], ",")
			}
		} else {
			break
		}

		cmd = strings.ToLower(cmd)
		param_0 := ""
		if len(param) > 0 {
			param_0 = strings.ToLower(param[0])
		}
		switch cmd {
		case "auto":
			// nop, auto is always off
			break
		case "echo":
			// nop, echo is always off
			break
		case "led1":
			if param_0 == "1" || param_0 == "on" || param_0 == "true" {
				d0.Set(true)
			} else {
				//			} else if param_0 == "0" || param_0 == "off" || param_0 == "false" {
				d0.Set(false)
			}
			break
		case "led2":
			if param_0 == "1" || param_0 == "on" || param_0 == "true" {
				d1.Set(true)
			} else {
				//			} else if param_0 == "0" || param_0 == "off" || param_0 == "false" {
				d1.Set(false)
			}
			break
		case "heater":
			break
			// nop, led1, led2 no response
		case "getrh":
			fmt.Printf(":%.2f, %.2f, %02x\n", 23.0+rand.Float64(), 45+rand.Float64(), 0xFF) // Tmpr, humid, crc TODO: add random noise
		case "ver":
			fmt.Printf("%s\n", ver_string)
		case "serial":
			fmt.Printf("%s\n", ver_serial)
		default:
		}
		//		led_blue.Set(sw != 1)

	}
}
