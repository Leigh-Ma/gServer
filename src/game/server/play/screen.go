package play

import (
	. "types"
)

const (
	scrWidth  = 1920
	scrHeight = 1080
)

type Screen struct {
	*screenElements
	Width   int16
	Height  int16
	FrameId int64
	add     *screenElements
	del     *screenElements
}

func NewScreen() *Screen {
	scr := &Screen{Width: scrWidth, Height: scrHeight, FrameId: 0}

	scr.screenElements = NewScreenElements()
	scr.add = NewScreenElements()
	scr.del = NewScreenElements()

	scr.initDecorations()
	return scr
}

func (se *Screen) initDecorations() {

}

func (scr *Screen) GetExistActive(actId IdString) *ScrActive {
	act := scr.Actives[actId]
	return act
}

func (scr *Screen) AddOrReplaceActive(act *ScrActive) {
	scr.addOrReplaceActive(act)
	scr.add.addOrReplaceActive(act)
}

func (scr *Screen) DelActive(actId IdString) {
	scr.delActive(actId)
	scr.del.delActive(actId)
}

func (scr *Screen) AddDecoration(decoration *ScrDecoration) {
	scr.addDecoration(decoration)
	scr.add.addDecoration(decoration)
}

func (scr *Screen) ScreenInfo() *ScreenInfo {
	si := &ScreenInfo{FrameId: scr.FrameId, Width: int32(scr.Width), Height: int32(scr.Height)}
	scr.Lock()
	si.Roles = toClientScreenRoleInfo(scr.Actives)
	si.Decorations = toClientScreenDecorations(scr.Decorations)
	si.Arrows = toClientScreenArrowInfo(scr.Actives)
	scr.Unlock()

	return si
}

func (scr *Screen) ScreenChangeNotify() *ScreenChangeNotify {
	scr.Lock()
	add := scr.add
	del := scr.del
	if add.Blank() || del.Blank() {
		scr.Unlock()
		return nil
	}

	scr.add = NewScreenElements()
	scr.del = NewScreenElements()
	scr.Unlock()

	si := &ScreenChangeNotify{FrameId: scr.FrameId}
	si.Heroes = toChangedClientRoleInfo(add.Actives)
	si.HeroesMove = toClientRoleMove(add.Actives)
	si.Arrows = toClientScreenArrowInfo(add.Actives)
	si.DelArrowsIds = toClientActivesIds(del.Actives, screenObjTypeArrow)
	si.DelHeroIds = toClientActivesIds(del.Actives, screenObjTypePlayer)
	si.AddDecorations = toClientScreenDecorations(add.Decorations)
	si.DelDecorations = toClientScreenDecorations(del.Decorations)

	return si
}
