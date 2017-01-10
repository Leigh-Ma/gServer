package play

import (
	. "game/com"
	"library/config"
	"time"
	. "types"
)

const (
	screenSyncFrameRatio = 20
	maxRoomMemberNum     = 8
)

//room_id -> brdcastgroup.id
type Room struct {
	*BrdCastGroup
	*Screen

	Name string
	stop chan int
}

func NewRoom(id IdString, name string) *Room {
	r := &Room{Name: name}
	r.BrdCastGroup = NewBrdCastGroup4Server(getMyServerId(), id)
	r.Screen = NewScreen()

	return r
}

func (r *Room) BeginFrameSync() {
	go func() {
		//create broadcast group on gateway first
		AsyncSender.InstantSendServerNotify(MT_BrdCastSyncReq, r.BcgGroupDetail())
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

func (r *Room) SideRoles() (red, blue []*RoleInfo) {
	for _, act := range r.Actives {
		if act.Type == screenObjTypePlayer {
			switch act.Side {
			case config.RoomSideBlue:
				blue = append(blue, act.RoleInfo())
			case config.RoomSideRed:
				red = append(red, act.RoleInfo())
			}
		}
	}
	return
}
