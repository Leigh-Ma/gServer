package play

type Hero struct {
	HeroName string
	HeroType string
	Skin     string
	Skills   []string
}

func NewHero() *Hero {
	hero := &Hero{
		HeroType: RandomHeroType(),
		Skin:     "default_skin",
		Skills:   RandHeroSkillN(2),
	}
	hero.HeroName = hero.HeroType
	return hero
}
