package test

import (
	"client/connect"
	"log"
	"testing"
)

func TestRestartFalco(t *testing.T) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	connect.RestartFalco()
}
