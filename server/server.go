package main

import (
	"bufio"
	cmd2 "gkd/cmd"
	db2 "gkd/db"
	"gkd/utils"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type Server struct {
	Addr string
	Down bool
	db   *db2.Database
}

func NewServer(addr string, db *db2.Database) *Server {
	return &Server{
		Addr: addr,
		Down: false,
		db:   db,
	}
}

func (s *Server) handleCmd(data string) string {
	cmd := strings.Split(data, " ")
	cmdFunc := cmd2.CommandMap[cmd[0]]
	if cmdFunc == nil {
		log.Println("command not found")
	} else {
		return cmdFunc(s.db, cmd[1:])
	}
	return "error"
}

func (s *Server) handle(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(nil)
	for {
		reader.Reset(conn)
		if message, err := utils.NewMessageFromReader(reader); err != nil {
			log.Println(err)
			return
		} else {
			res := s.handleCmd(message.Data)
			reply := utils.NewMessage(uint32(len(res)), res)
			conn.Write(reply.ToBytes())
		}
	}
}

func (s *Server) listen() {
	listener, err := net.Listen("tcp", s.Addr)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		if s.Down {
			break
		}
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go func() {
			s.handle(conn)
		}()
	}
}

func listenSignal() chan os.Signal {
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	return sigChan
}

func (s *Server) Stop() {
	log.Println("save ... ")
	err := s.db.SaveData(".", "test.data")
	if err != nil {
		return
	}
	log.Println("close ... ")
}

func (s *Server) Start() {
	sigChan := listenSignal()
	go s.listen()
	<-sigChan
	s.Stop()
}

func main() {
	addr := ":8000"
	database, err := db2.OpenDatabase(".", "test.data")
	if err != nil {
		log.Println(err)
	}
	gkdServer := NewServer(addr, database)
	gkdServer.Start()
}
