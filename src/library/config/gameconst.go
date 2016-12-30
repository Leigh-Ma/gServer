package config

func IsValidHeroName(hero_name string) bool {
	return true
}

func IsValidHeroSkill(skill string) bool {
	return true
}

const (
	RoomSideMid = 0
	RoomSideOne = 1
	RoomSideTwo = 2

	RoomSideMidName = "mid"
	RoomSideOneName = "1"
	RoomSideTwoName = "2"
)

func GetRoomSideNo(side string) int8 {
	switch side {
	case "one":
		return int8(RoomSideOne)
	case "two":
		return int8(RoomSideTwo)
	}
	return int8(RoomSideMid)
}

func GetRoomSideName(side int8) string {
	switch side {
	case RoomSideOne:
		return RoomSideOneName
	case RoomSideTwo:
		return RoomSideTwoName
	}
	return RoomSideMidName
}
