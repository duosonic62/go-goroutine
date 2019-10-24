package main

import "time"

func connectToService() interface{}  {
		time.Sleep( 1 * time.Second)
		return struct {}{}
}
