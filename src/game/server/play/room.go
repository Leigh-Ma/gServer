package play

import (
	. "game/com"
	"time"
	. "types"
)

const (
	screenSyncFrameRatio = 20
)

type Room struct {
	*BrdCastGroup
	*Screen

	Name string
	stop chan int
}

func NewRoom(creatorId IdString, name string) *Room {
	r := &Room{Name: name}
	r.BrdCastGroup = NewBrdCastGroup4Server(getMyServerId(), creatorId)
	r.Screen = NewScreen()

	return r
}

func (r *Room) BeginFrameSync() {
	go func() {
		ticker := time.Tick(time.Microsecond * (1000 / screenSyncFrameRatio))
		for {
			select {
			case <-r.stop:
				break
			case <-ticker:
				if notify := r.ScreenChangeNotify(); notify != nil {
					AsyncSender.InstantSendBroadCastNotify(MT_ScreenChangeNotify, r.BrdCastGroup.Id, notify)
				}
			}
		}
	}()
}

func (r *Room) Destroy() {
	r.stop <- 1
}