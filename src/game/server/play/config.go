package play

import "math/rand"

var playerName []string = []string{"Fast-an", "Bahew", "Vinric", "Bertdon", "Evera", "Laupauljo", "Hesmac",
	"Orahbert", "Helmrie", "Ichet", "Eadchar", "Kev", "Ferdha", "Tala",
	"Cuthtol", "Ridvell", "Nieldan", "Fred'ne", "Chardmark", "Kimka", "Riedo",
	"Elimac", "Sylna", "Riechel", "Ordever", "Egelsa", "Mondtol", "Hewlau",
	"Vid", "Linnaldni", "Sylwulf", "Dun", "Linhal", "Rymeri", "Kev", "Rywig",
	"Frid'ja", "Frithruth", "Mondmit", "Soward", "Brytman", "Mes'ric", "Daca",
	"Markne", "Sulau", "Cynchet", "Bandomann", "Cuthferth", "Ly", "Geor", "Riskim",
	"Rycen", "Thyda", "Ridnan", "Crowphie", "Subando", "Chellbur", "Maxmann",
	"Charfer", "Cy-meri", "Ceoltho", "Doegel", "Rojo", "Ceolcen", "Timchard",
	"Drape", "Isummarma", "Nas", "Roley", "Lacrow", "Grimand", "Bertbeth",
	"Hildimar", "Escpe", "Jackwaru", "Paulrich", "Stanken", "Ciada", "Egelvia",
	"Ordlac", "Betan", "Saha", "Dona", "Lincon", "Ric'anne", "Joanmarma", "Jen-frith",
	"Ferthann", "Mitsan", "Men", "Acarnan", "Muel-grim", "Dun", "Hol", "Leof", "Esc-za",
	"Isum", "Beorthle", "Baneth", "Theodhe", "Kered", "Uric", "Theodchel", "Ja", "Leofubeth",
	"Kedo", "Rahdic", "Marry", "Cawise", "Danegel", "Lafwine", "Vin", "Dinorah", "Burchell",
	"Ferdrydoc", "Ciajo", "Vinan", "Sothae", "Wilthony", "Cynja", "Ethelwen", "Kenri", "Cyndun",
	"Grimrahham", "Syl", "Dathy", "Ledryt", "Elfne", "Reginethel", "Frithria", "Laccla", "Menham",
	"Ealsa", "And'can", "Geret", "Jobar", "Gardloc", "Kenpe", "Berkim", "Riaroy", "Stana", "Soja",
	"Saret", "Donmi", "Laury", "Fridcy", "Roladelhes", "Fer", "Fortinrah", "Lamark", "He'isen", "Dasga",
	"Ferum-red", "Tontineethel", "Censan", "Wisesean", "Doald", "Ceolchard", "Andgrim", "Theodnya", "Beorthwaru",
	"Sonhes", "Kimisum", "Ordan", "Sani", "Rideverlas", "Chetu", "Ichar", "Lafsa", "Laf-seph", "Thasan", "Vellu",
	"Waldmuel", "Annebawulf", "Gardchell", "Anja", "Donpher", "Riscas", "Nalddon", "Brytbet", "Joan", "Fridmar",
	"Lidun", "Hewna", "Nandic", "Dodcla", "Herenas", "Nethton", "Karob", "Sala", "Dosa", "Danmas", "Nanter",
	"Sacon", "Mesdon", "Rahton", "Cuthwig", "Karo", "Roymit", "Dadino"}

var playerNameSize = len(playerName)

func RandomPlayerName() string {
	idx := rand.Intn(playerNameSize)
	return playerName[idx]
}

var heroType = []string{
	"swordman", "archer", "craftsman", "wizard",
}

var heroTypeNameSize = len(heroType)

func RandomHeroType() string {
	idx := rand.Intn(heroTypeNameSize)
	return heroType[idx]
}

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
