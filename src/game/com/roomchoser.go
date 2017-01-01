package com

import (
    . "types"
)

func ChoseARoom() (serverId, roomId IdString) {
    for s, m := range brdCastGroupsByServer {
        serverId = s
        for id := range m.groups {
            roomId = id
            break
        }
        break
    }
    return
}