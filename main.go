package main

import "back_agendamento/server"

func main() {
	s := server.NewServer()
	s.Run()
}
