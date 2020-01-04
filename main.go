package main

import (
	qcloud "GoDDNS/Qcloud"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"GoDDNS/util"

	"github.com/robfig/cron"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	go Watchdog()
	osChannel := make(chan os.Signal)
	signal.Notify(osChannel, os.Interrupt, os.Kill, syscall.SIGUSR1, syscall.SIGUSR2)
	<-osChannel
}

func Watchdog() {
	setting := util.Setting()
	watch := cron.New()
	spec := setting.Cron
	watch.AddFunc(spec, func() {
		QcloudCommon := new(qcloud.QcloudCommon)
		QcloudCommon.DDNS()
	})
	watch.Start()
	select {}
}
