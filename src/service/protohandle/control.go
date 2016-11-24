package protohandle

import (
	cm "library/core/controlmsg"
	dm "library/core/datamsg"
)

func (t *protoHandle) Start(bus *dm.DataMsgPipe) bool {
	t.Output = bus
	return true
}

func (t *protoHandle) Pause() bool {
	return true
}

func (t *protoHandle) Resume() bool {
	return true
}

func (t *protoHandle) Exit() bool {
	return true
}

func (t *protoHandle) ControlHandler(msg *cm.ControlMsg) (int, int) {
	//switch msg.MsgType {
	//case cm.ControlMsgExit:
	//	logger.Info("%s:ControlMsgPipe.Cmd Read %d", t.Name, msg.MsgType)
	//	t.Echo <- &cm.ControlMsg{MsgType: cm.ControlMsgExit}
	//	logger.Info("%s:exit", t.Name)
	//	//return service.Return, service.FunOK
	//	return cm.NextActionReturn, cm.ProcessStatOK
	//case cm.ControlMsgPause:
	//	logger.Info("%s:paused", t.Name)
	//	t.Echo <- &cm.ControlMsg{MsgType: cm.ControlMsgPause}
	//	for {
	//		var resume bool = false
	//		select {
	//		case msg, ok := <-t.Cmd:
	//			if !ok {
	//				logger.Info("%s:Cmd Read error", t.Name)
	//				break
	//			}
	//			switch msg.MsgType {
	//			case cm.ControlMsgResume:
	//				t.Echo <- &cm.ControlMsg{MsgType: cm.ControlMsgResume}
	//				resume = true
	//				break
	//			}
	//		}
	//		if resume {
	//			break
	//		}
	//	}
	//	logger.Info("%s:resumed", t.Name)
	//}
	return cm.NextActionContinue, cm.ProcessStatOK
}
