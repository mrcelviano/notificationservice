package commons

import (
	"google.golang.org/grpc"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var StopTimeout = time.Minute * 2
var stopChan = make(chan bool)

// Обработчик сигналов операционной системы.
// Мягко (не принимает новые запросы и дожидается пока текущие обработаются) останавливает grpc сервер.
// Если в stopChan поступит false, приложение остановится через StopTimeout времени

func NewSignalHandler(s *grpc.Server) chan bool {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go handle(c)
	go func() {
		notStop := <-stopChan
		if s != nil {
			log.Println("STOP GRPC server")
			s.GracefulStop()
		}
		if !notStop {
			time.Sleep(StopTimeout)
			os.Exit(99)
		}
	}()
	return stopChan
}

func handle(sigChan <-chan os.Signal) {
	stopped := false
	for sig := range sigChan {
		log.Println("STOP TRIGGER: signal")
		log.Println(sig.String() + " signal!!!")
		if stopped {
			time.Sleep(time.Second)
			os.Exit(561)
		}
		stopped = true
		go func() {
			for {
				stopChan <- true
			}
		}()
	}
}
