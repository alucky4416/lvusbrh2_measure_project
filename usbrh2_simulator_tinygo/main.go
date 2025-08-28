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

	ver_string := "usbrh2, ver 1.0.0"
	serial_string := "123456789"
	//	table := crc8.MakeTable(crc8.CRC8)

	led_blue := machine.LED_BLUE
	led_blue.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})
	led1 := machine.D0
	led1.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})
	led1.Set(false)

	led2 := machine.D1
	led2.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})
	led2.Set(false)

	time.Sleep(2 * time.Second) // wait for USB-CDC init

	var cmd string
	var param []string
	sw := false
	for {
		reader := bufio.NewReader(os.Stdin)
		line, _, err := reader.ReadLine() // ReadLine()は改行コードの直前まで読み込み
		if err != nil {
			break
		}
		//		fmt.Println(line)

		if len(line) > 0 {
			line2 := string(line)
			//			line2 := strings.TrimLeft(string(line), " \r\n") // 末尾の改行コードを削除
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
				led1.Set(true)
			} else {
				//			} else if param_0 == "0" || param_0 == "off" || param_0 == "false" {
				led1.Set(false)
			}
			break
		case "led2":
			if param_0 == "1" || param_0 == "on" || param_0 == "true" {
				led2.Set(true)
			} else {
				//			} else if param_0 == "0" || param_0 == "off" || param_0 == "false" {
				led2.Set(false)
			}
			break
		case "heater":
			// nop, heater no response
			break
		case "status":
			// nop, status no response
			break
		case "list":
			// nop, list no response
			break
		case "getrh":
			resp := fmt.Sprintf(":%.2f,%.2f,", 23.0+rand.Float64(), 46.0+rand.Float64()) // :<Tmpr>,<humid>,
			// table := crc8.MakeTable(crc8.CRC8)
			crc := 0xFF
			//			crc := crc8.Checksum([]byte(resp), table)
			fmt.Printf("%s%02X\n", resp, crc) // :<Tmpr,humid,>crc8(HEX)
		case "ver":
			fmt.Printf("%s\n", ver_string)
		case "serial":
			fmt.Printf("%s\n", serial_string)
		default:
		}
		led_blue.Set(sw)
		sw = !sw

	}
}
