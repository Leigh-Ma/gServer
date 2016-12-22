package types

const (
    OK              = int32(0)
    ERR_INVALID_REQ = int32(1)
    ERR_GROUP_NOT_FOUND = int32(2)
    ERR_ROOM_NOT_FOUND = int32(3)
    ERR_PLAYER_NOT_FOUND = int32(4)
)

var ErrDesc = map[int32]string{
    OK:              "success",
    ERR_INVALID_REQ: "invalid req message type",
    ERR_GROUP_NOT_FOUND: "boradcast group not found",
    ERR_ROOM_NOT_FOUND: "room not found",
    ERR_PLAYER_NOT_FOUND: "player not found",
}