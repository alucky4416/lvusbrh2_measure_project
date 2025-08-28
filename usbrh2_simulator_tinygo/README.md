# USBRH2 simulator for xiao-rp2040 (or pico, rp2040 controller)

## About
This firmware is USBRH2 simulator for Seeed xiao-rp2040 or pico, RP2040 controller.  
USBRH2 is Temperature and Relative humidity sensor (using Sensirion SHT-31).  
This firmware is developed using TinyGo.  

USBRH2  
https://strawberry-linux.com/catalog/items?code=52031

Seeed xiao-rp2040  
https://wiki.seeedstudio.com/XIAO-RP2040/

TinyGo  
https://tinygo.org/

## Feature
This firmware supports the following commands.  

* getrh:  
When the USBRH2 temperature and humidity acquisition command (getrh) is received, it returns simulated temperature and relative humidity data.  
CRC is always 0xFF.  
* serial, ver:  
For the serial number acquisition command (serial) and version information acquisition command (ver), it responds with simulated serial numbers and version information, respectively. 
* led1, led2:  
  Turn on/off LED on xiao-rp2040 (LED1 is D0, LED2 is D1)
* echo, auto, heater and other:  
It does not respond to other commands such as echo, auto, or heater.  
always echo off, heater off, auto off.  
* status, list, help:  
not supported, no response.

### Restrict
* delimiter is LF only. not support CR.

## How to build, and bake

### build 
```
\> tinygo build -o usbrh2_sim_xiao-rp2040.uf2 -target xiao-rp2040 main.go
```
### bake
While pressing the BOOT switch on the rp2040, connect the rp2040 to a USB port.  
then the RP2040 folder will then appear.  

copy uf2 file to RP2040 USB mass storage  

```
\> copy usbrh2_sim_xiao-rp2040.uf2 <RP2040 USB mass storage>
```

another way  

### flash (build and bake)
Without pressing the BOOT switch. connect the rp2040 to a USB port.
```
\> tinygo flash -target xiao-rp2040 main.go
```

### Reconnect
After flashing the firmware, reconnect the RP2040 and it will start working automatically. It should be recognized as a virtual COM port.

* On Windows  
COM1 or COM2, COM3, ...

* On Linux  
/dev/ttyACM0 or /dev/ttyACM1, /dev/ttyACM2,...


## How to port to pico or RP2040 controllers
main.go change the machine.<IOname> part below.  
```
	led_blue := machine.LED_BLUE
	led1 := machine.D0
	led2 := machine.D1
```
example pico:
```
	led_blue := machine.LED
	led1 := machine.GP0
	led2 := machine.GP1
```

To search for I/O names online, do the following:  
```
"tinygo <controller product name>"
```

change build option "-target"  
example pico: 
```
\> tinygo flash -target pico main.go
```


## License
This software is released under the MIT License
