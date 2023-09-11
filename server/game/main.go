package main

import (
	"fmt"
	"os"
	"os/signal"
	"server/game/core/log"
	"server/game/game"
	"server/game/gate"
	"syscall"
	"time"
)

var (
	SigChan   = make(chan os.Signal, 1)
	CloseChan chan struct{}
)

func main() {
	log.Info("game start ...")

	signal.Notify(SigChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	/*
		HUP 1	控制台中的终端/程序中断
		INT 2	键盘的插入指令（同Ctrl+C）
		QUIT 3	键盘的中断指令（同Ctrl+\）
		TERM 15	程序的终止指令
		KILL 9	程序的强制终止指令（暴力砍掉）
		CONT 18	程序的再启动指令
		STOP 19	程序的停止指令（同Ctrl+Z）
	*/
	logicTicker := time.NewTicker(10 * time.Millisecond)

	CloseChan = make(chan struct{})
	exitChan := make(chan struct{})
	go func() {
		defer func() {
			exitChan <- struct{}{}
		}()

		log.Info("start listen ...")
		gs := gate.NewServer("127.0.0.1", 10086)
		go gs.Run()

		gm := game.NewGame()

		for {
			select {
			case <-CloseChan:
				goto QUIT
			case sig := <-SigChan:
				log.Info(fmt.Sprintf("收到信号：%v", sig))
				close(CloseChan)
			case <-logicTicker.C:
				gm.Loop()
			case client := <-gs.AcceptChan:
				gm.OnAccept(client)
			}
		}

	QUIT:
		log.Info("start stop")
		gs.Stop()

		log.Info("stop all done")
		return
	}()

	log.Info("running ...")
	<-exitChan

	log.Info("exit ...")
}
