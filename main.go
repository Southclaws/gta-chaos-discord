package main

import (
	"fmt"
	"time"

	"github.com/natefinch/npipe"
	"k8s.io/apimachinery/pkg/util/rand"
)

const PIPE = `\\.\pipe\GTATrilogyChaosModPipe`
const COOLDOWN = time.Minute / 2

// time %d,%d,
// set_seed:957271
// effect:remove_all_weapons:30000:Remove All Weapons:N/A:0

// FunctionEffect
// WeatherEffect
// SpawnVehicleEffect
// TeleportationEffect

func main() {
	ln, err := npipe.Dial(PIPE)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected!")

	effects := Effects()

	last := time.Now()
	tc := time.NewTicker(time.Millisecond * 100)
	for range tc.C {
		message := fmt.Sprintf("time:%d,%d:-1:N/A:N/A:0", int64(COOLDOWN.Milliseconds())-time.Since(last).Milliseconds(), COOLDOWN.Milliseconds())

		n, err := ln.Write([]byte(message))
		fmt.Println("writing", message, n, err)

		if time.Since(last) > COOLDOWN {
			idx := rand.Intn(len(effects))
			effect := effects[idx]
			message := fmt.Sprintf("%s:N/A:0", EffectToMessage(effect))

			n, err := ln.Write([]byte(message))
			fmt.Println("writing", message, n, err)

			last = time.Now()
		}

	}
}

// func main() {
// 	fmt.Println("starting to listen to", PIPE)
// 	ln, err := npipe.Listen(PIPE)
// 	if err != nil {
// 		panic(err)
// 	}
// 	for {
// 		conn, err := ln.Accept()
// 		if err != nil {
// 			panic(err)
// 		}
// 		go handleConnection(conn)
// 	}
// }

// func handleConnection(conn net.Conn) {
// 	// fmt.Println("handleConnection", conn.LocalAddr())

// 	// buf := make([]byte, 0, 4096)
// 	tmp := make([]byte, 256)
// 	for {
// 		_, err := conn.Read(tmp)
// 		if err != nil {
// 			if err != io.EOF {
// 				panic(err)
// 			}
// 			break
// 		}
// 		stmp := string(tmp)

// 		if !strings.HasPrefix(stmp, "time") {
// 			fmt.Println(string(tmp))
// 		}
// 	}
// }
