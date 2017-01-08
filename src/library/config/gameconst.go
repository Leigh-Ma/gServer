package config

func IsValidHeroSkill(skill string) bool {
	return true
}

const (
	RoomSideMid = 0
	RoomSideA   = 1
	RoomSideB   = 2

	RoomSideMidName = "mid"
	RoomSideAName   = "a"
	RoomSideBName   = "b"
)

func GetRoomSideNo(side string) int8 {
	switch side {
	case RoomSideAName:
		return int8(RoomSideA)
	case RoomSideBName:
		return int8(RoomSideB)
	}
	return int8(RoomSideMid)
}

func GetRoomSideName(side int8) string {
	switch side {
	case RoomSideA:
		return RoomSideAName
	case RoomSideB:
		return RoomSideBName
	}
	return RoomSideMidName
}
