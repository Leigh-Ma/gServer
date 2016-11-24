package main

import (
	"math/rand"
	"runtime"
	"time"
)

import (
	cm "library/core/controlmsg"
	"library/idgen"
	"library/logger"
	"service"
	//"service/mongo"
	"service/protohandle"
	"service/prototcp"
	"service/timer"
)

func stopAndCleanMemory() {
	memstat := &runtime.MemStats{}
	runtime.ReadMemStats(memstat)
	logger.Info("before gc:memstat.Alloc:%d K", memstat.Alloc/1024)
	runtime.GC()
	runtime.ReadMemStats(memstat)
	logger.Info("after gc:memstat.Alloc:%d K", memstat.Alloc/1024)
}

func main() {
	logger.EnableDebugLog(true)
	logger.Info("main start")
	if runtime.GOOS != "linux" && runtime.GOOS != "darwin" {
		logger.Warn("Only tested in linux and mac os operating system, be carefally using in %v", runtime.GOOS)
	}

	rand.Seed(time.Now().UTC().Unix())

	idgen.InitIDGen("1")

	signalMod := initSystemSignalHandler()
	// good idea to stop the world and clean memory before get job
	stopAndCleanMemory()

	distributor := service.NewEngine("engine")
	distributor.Start(nil)

	timerSrv := timer.NewTimer("timer")
	service.StartService(timerSrv, distributor.BUS)

	protoDealer := protohandle.NewProtoHandle(protohandle.ServiceName)
	service.StartService(protoDealer, distributor.BUS)

	//mongosrv := mongo.NewMongo("mongo", "127.0.0.1", "27017")
	//service.StartService(mongosrv, distributor.BUS)

	protoSvr := prototcp.NewProtoTCPServer(prototcp.ServiceName, "0.0.0.0", "7788")
	service.StartService(protoSvr, distributor.BUS)

	for {
		select {
		case sigMsg, ok := <-signalMod.Echo:
			if !ok {
				logger.Error("sigint echo error %t", ok)
				continue
			}
			logger.Info("receive:signal echo:%+v", sigMsg)
			distributor.Cmd <- &cm.ControlMsg{MsgType: cm.ControlMsgExit}
			<-distributor.Echo
			return
		case sigMsg, ok := <-signalMod.Echo:
			if !ok {
				logger.Error("sigterm echo error %t", ok)
				continue
			}
			logger.Info("receive:signal echo:%+v", sigMsg)
			distributor.Cmd <- &cm.ControlMsg{MsgType: cm.ControlMsgExit}
			<-distributor.Echo
			return
		}
	}
}
