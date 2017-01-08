package config

func IsValidHeroSkill(skill string) bool {
	return true
}

const (
	RoomSideMid  = 0
	RoomSideBlue = 1
	RoomSideRed  = 2

	RoomSideMidName  = "mid"
	RoomSideBlueName = "r"
	RoomSideRedName  = "b"
)

func GetRoomSideNo(side string) int8 {
	switch side {
	case RoomSideBlueName:
		return int8(RoomSideBlue)
	case RoomSideRedName:
		return int8(RoomSideRed)
	}
	return int8(RoomSideMid)
}

func GetRoomSideName(side int8) string {
	switch side {
	case RoomSideBlue:
		return RoomSideBlueName
	case RoomSideRed:
		return RoomSideRedName
	}
	return RoomSideMidName
}
