package play

import (
	"library/config"
	"sync"
	"types"
)

const (
	screenObjTypePlayer = 1
	screenObjTypeArrow  = 2
)

const (
	actDirectionUp    = 1
	actDirectionDown  = 2
	actDirectionLeft  = 3
	actDirectionRight = 4
)

type ActDetail struct {
	BelongsTo types.IdString
	Name      string
	SubType   string
	Hp        int32
	FullHp    int32
	Skin      string
	Skills    []string
	dirty     bool
}

func (ad *ActDetail) IsDetailChanged() bool {
	return ad.dirty
}

func (ad *ActDetail) SetDetailChanged() {
	ad.dirty = true
}

func (ad *ActDetail) ClearDetailChanged() {
	ad.dirty = false
}

type ScrMove struct {
	DirectX int16
	DirectY int16
	Speed   int16
	PosX    int16
	PosY    int16
}

type ScrActive struct {
	Id   types.IdString
	Type int8
	Side int8
	ScrMove
	ActDetail
}

func (sa *ScrActive) ToMove() *types.ActiveMove {
	return &types.ActiveMove{
		Id:      string(sa.Id),
		Type:    int32(sa.Type),
		DirectX: int32(sa.DirectX),
		DirectY: int32(sa.DirectY),
		Speed:   int32(sa.Speed),
		PosX:    int32(sa.PosX),
		PosY:    int32(sa.PosY),
		Side:    config.GetRoomSideName(sa.Side),
	}
}

func (sa *ScrActive) Detail() *types.ActiveDetail {
	ad := &types.ActiveDetail{
		Id:        string(sa.Id),
		BelongsTo: string(sa.BelongsTo),
		Name:      sa.Name,
		SubType:   sa.SubType,
		Skin:      sa.Skin,
		Hp:        int32(sa.Hp),
		FullHp:    int32(sa.FullHp),
	}
	if sa.Type == screenObjTypePlayer {
		for _, skill := range sa.Skills {
			ad.Skills = append(ad.Skills, skill)
		}
	}
	return ad
}

func (sa *ScrActive) RoleInfo() *types.RoleInfo {
	if sa.Type == screenObjTypePlayer {
		hi := &types.RoleInfo{
			UserId:   string(sa.Id),
			UserName: sa.Name,
			HeroType: sa.SubType,
			Skin:     sa.Skin,
			Hp:       int32(sa.Hp),
			FullHp:   int32(sa.FullHp),
		}
		hi.Skills = append(hi.Skills, sa.Skills...)
		return hi
	}

	return nil
}

func (sa *ScrActive) ArrowInfo() *types.ArrowInfo {
	if sa.Type == screenObjTypeArrow {
		return &types.ArrowInfo{}
	}

	return nil
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

func toClientRoleInfo(actives map[types.IdString]*ScrActive) (ris []*types.RoleInfo) {
	if len(actives) == 0 {
		return nil
	}

	for _, act := range actives {
		ri := act.RoleInfo()
		if ri != nil {
			ris = append(ris, ri)
		}
	}
	return ris
}

func toChangedClientRoleInfo(actives map[types.IdString]*ScrActive) (ris []*types.RoleInfo) {
	if len(actives) == 0 {
		return nil
	}

	for _, act := range actives {
		ri := act.RoleInfo()
		if ri != nil && act.IsDetailChanged() {
			ris = append(ris, ri)
		}
	}
	return ris
}

func toClientRoleMove(actives map[types.IdString]*ScrActive) (ams []*types.ActiveMove) {
	if len(actives) == 0 {
		return nil
	}

	for _, act := range actives {
		if act.Type == screenObjTypePlayer {
			ams = append(ams, act.ToMove())
		}
	}
	return ams
}

func toClientScreenRoleInfo(actives map[types.IdString]*ScrActive) (sis []*types.ScreenRoleInfo) {
	if len(actives) == 0 {
		return nil
	}

	for _, act := range actives {
		hero := act.RoleInfo()
		if hero != nil {
			sis = append(sis, &types.ScreenRoleInfo{Hero: hero, HeroMove: act.ToMove()})
		}
	}
	return sis
}

func toClientScreenArrowInfo(actives map[types.IdString]*ScrActive) (sas []*types.ScreenArrowInfo) {
	if len(actives) == 0 {
		return nil
	}

	for _, act := range actives {
		arrow := act.ArrowInfo()
		if arrow != nil {
			sas = append(sas, &types.ScreenArrowInfo{Arrow: arrow, ArrowMove: act.ToMove()})
		}
	}
	return sas
}

func toClientActivesIds(actives map[types.IdString]*ScrActive, tp int8) (ids []string) {
	if len(actives) == 0 {
		return nil
	}

	for _, act := range actives {
		if act.Type == tp {
			ids = append(ids, string(act.Id))
		}
	}
	return ids
}

func toClientScreenDecorations(decorations []*ScrDecoration) []*types.ScreenDecoration {
	if len(decorations) == 0 {
		return nil
	}
	sds := make([]*types.ScreenDecoration, len(decorations))
	for idx, decoration := range decorations {
		sds[idx] = decoration.ToClient()
	}
	return sds
}
