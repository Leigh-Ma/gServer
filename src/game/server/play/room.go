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
	oneFightDuration     = 900
	roomStatusWait       = "wait"
	roomStatusFighting   = "fight"
	roomStatusStopped    = "stop"
)

//room_id -> brdcastgroup.id
type Room struct {
	*BrdCastGroup
	*Screen

	Status        string
	FightStartAt  UnixTS
	FightDuration UnixTS
	Name          string
	stop          chan int
}

func NewRoom(id IdString, name string) *Room {
	r := &Room{Name: name}
	r.BrdCastGroup = NewBrdCastGroup4Server(getMyServerId(), id)
	r.Screen = NewScreen()
	r.Status = roomStatusWait
	r.FightDuration = UnixTS(oneFightDuration)

	r.stop = make(chan int)

	return r
}

func (r *Room) RoomInfo() *RoomInfo {
	return &RoomInfo{
		RoomId:        string(r.Id),
		Status:        r.Status,
		FightDuration: int32(r.FightDuration),
		FightStartAt:  int32(r.FightStartAt),
	}
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

				if PlayFrame.FrameTime()-oneFightDuration > r.FightStartAt {
					r.FightEnd()
					r.Destroy()
				}
			}
		}
	}()
}

func (r *Room) FightBegin() {
	r.Status = roomStatusFighting
	r.FightStartAt = PlayFrame.FrameTime()

	notify := &StartFightNotify{
		Room: r.RoomInfo(),
		Screen: r.ScreenInfo(),
	}

	AsyncSender.InstantSendBroadCastNotify(MT_StartFightNotify, r.BrdCastGroup.Id, notify)
}

func (r *Room) FightEnd() {
	r.Status = roomStatusStopped
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
