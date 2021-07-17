package main

import "log"
import "fmt"
import "strings"
import "go.bug.st/serial"
import "go.bug.st/serial/enumerator"

// import "github.com/itchyny/volume-go"

/**
 * VOLUME
 * get volume
 *  vol, err := volume.GetVolume()
 *  if err != nil {
 *    log.Fatal(err)
 *  }
 * mute volume
 *  volume.Mute()
 * set volume - we'll use this to set to the adc value of the slider (as a perc of its max)
 *  volume.SetVolume(50)
 */

func main() {
	list()
	mode := &serial.Mode{
		BaudRate: 115200,
	}
	port, err := serial.Open("COM3", mode)
	if err != nil {
		log.Fatal(err)
	}

	for {
		cmd := strings.Trim(getCmd(port), "\n")
		fmt.Printf("cmd: %s\n", cmd)
		if strings.EqualFold(strings.Trim(cmd, "\n "), "END") {
			break
		} else if strings.EqualFold(strings.Trim(cmd, "\n "), "B01") {

		} else if strings.EqualFold(strings.Trim(cmd, "\n "), "B00") {

		}
	}
	fmt.Println("end of prog")
}

func getCmd(port serial.Port) string {
	buff := make([]byte, 100)
	cmd := ""
	for {
		// Reads up to 100 bytes
		n, err := port.Read(buff)
		if err != nil {
			log.Fatal(err)
		}
		if n == 0 {
			break
		}
		cmd += string(buff[:n])
		// If we receive a newline stop reading
		if strings.Contains(string(buff[:n]), "\n") {
			return string(cmd)
		}
	}
	return "EOF"
}

func list() {
	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		fmt.Println("No serial ports found!")
		return
	}
	for _, port := range ports {
		fmt.Printf("Found port: %s\n", port.Name)
		if port.IsUSB {
			fmt.Printf("   USB ID     %s:%s\n", port.VID, port.PID)
			fmt.Printf("   USB serial %s\n", port.SerialNumber)
		}
	}
}
