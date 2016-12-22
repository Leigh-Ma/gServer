package play

import (
	"sync"
	"types"
)

const (
	screenObjTypePlayer = 1
	screenObjTypeBullet = 2
)

const (
	actDirectionUp    = 1
	actDirectionDown  = 2
	actDirectionLeft  = 3
	actDirectionRight = 4
)

type ScrMove struct {
	DirectX int16
	DirectY int16
	Speed   int16
	PosX    int16
	PosY    int16
}

type ScrActive struct {
	Id        types.IdString
	BelongsTo types.IdString
	ScrMove

	Name    string
	Type    int8
	SubType int8
	Skin    int8

	Hp     int32
	FullHp int32
}

func (sa *ScrActive) ToClient() *types.ScreenActive {
	return &types.ScreenActive{
		Id: string(sa.Id),

		DirectX: int32(sa.DirectX),
		DirectY: int32(sa.DirectY),
		Speed:   int32(sa.Speed),
		PosX:    int32(sa.PosX),
		PosY:    int32(sa.PosY),

		Name:    sa.Name,
		Type:    int32(sa.Type),
		SubType: int32(sa.SubType),
		Skin:    int32(sa.Skin),

		Hp:     sa.Hp,
		FullHp: sa.FullHp,
	}
}

type ScrDecoration struct {
	AssetId int8
	X       int16
	Y       int16
}

func (sd *ScrDecoration) ToClient() *types.ScreenDecoration {
	return &types.ScreenDecoration{
		AssetId: int32(sd.AssetId),
		X:       int32(sd.X),
		Y:       int32(sd.Y),
	}
}

type screenElements struct {
	Decorations []*ScrDecoration
	Actives     map[types.IdString]*ScrActive
	sync.Mutex
}

func NewScreenElements() *screenElements {
	return &screenElements{
		Decorations: make([]*ScrDecoration, 0),
		Actives:     make(map[types.IdString]*ScrActive, 0),
	}
}

func (se *screenElements) Blank() bool {
	return len(se.Decorations) == 0 && len(se.Actives) == 0
}

func (se *screenElements) addOrReplaceActive(act *ScrActive) {
	se.Lock()
	se.Actives[act.Id] = act
	se.Unlock()
}

func (se *screenElements) delActive(actId types.IdString) {
	se.Lock()
	delete(se.Actives, actId)
	se.Unlock()
}

func (se *screenElements) addDecoration(decoration *ScrDecoration) {
	se.Decorations = append(se.Decorations, decoration)
}

func toClientActives(actives map[types.IdString]*ScrActive) []*types.ScreenActive {
	if len(actives) == 0 {
		return nil
	}
	sas := make([]*types.ScreenActive, len(actives))
	idx := 0
	for _, act := range actives {
		sas[idx] = act.ToClient()
		idx += 1
	}
	return sas
}

func toClientActivesIds(actives map[types.IdString]*ScrActive) []string {
	if len(actives) == 0 {
		return nil
	}
	sas := make([]string, len(actives))
	idx := 0
	for id := range actives {
		sas[idx] = string(id)
		idx += 1
	}
	return sas
}

func toClientDecorations(decorations []*ScrDecoration) []*types.ScreenDecoration {
	if len(decorations) == 0 {
		return nil
	}
	sds := make([]*types.ScreenDecoration, len(decorations))
	for idx, decoration := range decorations {
		sds[idx] = decoration.ToClient()
	}
	return sds
}
