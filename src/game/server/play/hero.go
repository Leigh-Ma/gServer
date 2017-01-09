package play

import (
	"library/idgen"
	"math/rand"
	. "types"
)

type Hero struct {
	UserName string
	HeroType string
	Skin     string
	Skills   []string
	Hp       int32
	FullHp   int32
}

func NewHero() (h *Hero) {
	for _, hero := range AllHeroes {
		hero.Skin = "default_skin"
		hero.Hp = hero.FullHp
		hero.Skills = RandHeroSkillN(2)
		h = &hero
		break
	}

	return
}

func (h *Hero) debugRoleInfo() *RoleInfo {
	hi := &RoleInfo{
		UserId:   string(idgen.NewObjectID().ToIdString()),
		UserName: h.UserName,
		HeroType: h.HeroType,
		FullHp:   h.FullHp,
		Hp:       h.Hp,
		Skin:     h.Skin,
	}
	hi.Skills = append(hi.Skills, h.Skills...)

	return hi
}

var SwordMan = Hero{HeroType: "swordman", FullHp: 1000}
var Archer = Hero{HeroType: "archer", FullHp: 1000}
var CraftsMan = Hero{HeroType: "craftsman", FullHp: 1000}
var Wizard = Hero{HeroType: "wizard", FullHp: 1000}
var AllHeroes = map[string]Hero{
	"swordman":  SwordMan,
	"archer":    Archer,
	"craftsman": CraftsMan,
	"wizard":    Wizard,
}

var AllHeroesNum = 4

var heroSkillName = map[int]string{
	0: "hs_no_1",
	1: "hs_no_2",
	2: "hs_no_3",
}

var heroSkillNameSize = len(heroSkillName)

func RandomHeroSkillName() string {
	idx := rand.Intn(heroSkillNameSize)
	return heroSkillName[idx]
}

func RandHeroSkillN(n int) []string {
	skills := make([]string, n)
	idx := 0
	for _, skill := range heroSkillName {
		skills[idx] = skill
		idx += 1
		if idx >= n {
			break
		}
	}
	return skills
}
