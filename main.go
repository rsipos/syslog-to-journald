package main

import (
	"context"
	"fmt"
	"github.com/coreos/go-systemd/daemon"
	"github.com/coreos/go-systemd/journal"
	"log"
	"log/syslog"
	"net"
	"os"
	"os/signal"
	"strings"
	"time"
)

type syslogPrio struct {
	name string
	val  syslog.Priority
}

var priorities = [11]syslogPrio{
	{"alert", syslog.LOG_ALERT},
	{"crit", syslog.LOG_CRIT},
	{"debug", syslog.LOG_DEBUG},
	{"emerg", syslog.LOG_EMERG},
	{"err", syslog.LOG_ERR},
	{"error", syslog.LOG_ERR},
	{"info", syslog.LOG_INFO},
	{"notice", syslog.LOG_NOTICE},
	{"panic", syslog.LOG_EMERG},
	{"warn", syslog.LOG_WARNING},
	{"warning", syslog.LOG_WARNING},
}

var count int

func main() {
	// Lets prepare a address at any address at port 8080
	ServerAddr, _ := net.ResolveUDPAddr("udp", ":514")

	// Now listen at selected port
	ServerConn, _ := net.ListenUDP("udp", ServerAddr)

	daemon.SdNotify(false, fmt.Sprintf("%s\nSTATUS=Listening for syslog input...", daemon.SdNotifyReady))

	go func() {
		for {
			handleUDPConnection(ServerConn)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := ServerConn.Close(); err != nil {
		log.Fatal(err)
	}
	daemon.SdNotify(false, daemon.SdNotifyStopping)
	log.Print("UDP Shutdown...")
}

func handleUDPConnection(ServerConn *net.UDPConn) {
	buf := make([]byte, 1024)
	prio := syslog.LOG_INFO
	n, addr, err := ServerConn.ReadFromUDP(buf)
	if err != nil {
		daemon.SdNotify(false, daemon.SdNotifyStopping)
		log.Println("Error: ", err)
	}
	buf = buf[:n]

	// parse priority
	for _, k := range priorities {
		if strings.Contains(strings.Split(string(buf), " ")[0], k.name) {
			prio = k.val
		}
	}

	journal.Send(string(buf), journal.Priority(prio), map[string]string{"SYSLOG_IDENTIFIER": addr.IP.String()})

	count++
	daemon.SdNotify(false, fmt.Sprintf("%s\nSTATUS=Forwarded %d syslog messages.", daemon.SdNotifyReady, count))
}
