package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

var(
	AxeHero *hero
	Io *hero
	Juggernaut *hero

	smallCampOne int
	smallCampTwo int

	mediumCampOne int
	mediumCampTwo int
	mediumCampThree int
	mediumCampFour int
	mediumCampFive int
	mediumCampSix int

	largeCampOne int
	largeCampTwo int
	largeCampThree int
	largeCampFour int
	largeCampFive int
	largeCampSix int

	ancientCampOne int
	ancientCampTwo int
	ancientCampThree int
	ancientCampFour int
)
func main(){
	//builder for hero creation
	//axehero()

	//creep factory
	//creeps()
	//startgame with bounty and creeps




	//Dire
	fmt.Println("\nGame started:")
	direCreepsFactory := NewCreepsFactory("Dire",true)
	direCreepsMelee := direCreepsFactory("Melee creep",3,10,2,150)
	direCreepsRanged := direCreepsFactory("Ranged creep",1,12,0,100)

	//Radiant
	radiantCreepsFactory := NewCreepsFactory("Radiant",true)
	radiantCreepsMelee := radiantCreepsFactory("Melee creep",3,10,2,150)
	radiantCreepsRanged := radiantCreepsFactory("Ranged creep",1,12,0,100)

	//Neutrals
	neutralCreepsFactory := NewCreepsFactory("Neutrals", true)
	//small camps
	neutralCreepsKobolds := neutralCreepsFactory("Kobold",3, 10, 0, 240)
	neutralCreepsKoboldSoldier := neutralCreepsFactory("Kobold Soldier",1, 15, 1, 325)
	neutralCreepsKoboldForeman := neutralCreepsFactory("Kobold Foreman",1, 15, 1, 400)

	neutralCreepsHillTrollBerserkers := neutralCreepsFactory("Hill Troll Berserkers",2, 37, 1, 500)
	neutralCreepsHillTrollPriest := neutralCreepsFactory("Hill Troll Priest",1, 32, 0, 450)

	neutralCreepsHillTrollBerserkersVolume2 := neutralCreepsFactory("Hill Troll Berserkers",2, 37, 1, 500)
	neutralCreepsKoboldForemanVolume2 := neutralCreepsFactory("Kobold Foreman",1, 15, 1, 400)

	neutralCreepsVhoulAssassins := neutralCreepsFactory("Vhoul Assassins",3, 36, 1, 370)

	neutralCreepsFellSpirits := neutralCreepsFactory("Fell Spirits",2, 15, 1, 400)
	neutralCreepsGhost := neutralCreepsFactory("Ghost",1, 50, 1, 500)

	neutralCreepsHarpyScouts := neutralCreepsFactory("Harpy Scouts",2, 35, 1, 400)
	neutralCreepsHarpyStormcrafter := neutralCreepsFactory("Harpy Stormcrafter",1, 37, 2, 550)

	//medium camps
	neutralCreepsGiantWolfs := neutralCreepsFactory("Giant Wolfs",2,30,1,500)
	neutralCreepsAlphaWolf := neutralCreepsFactory("Alpha Wolf",1,33,3,600)

	neutralCreepsCentaurCourser := neutralCreepsFactory("Centaur Courser", 1, 21, 1, 320)
	neutralCreepsCentaurConqueror := neutralCreepsFactory("Centaur Conqueror", 1, 55, 4, 1100)

	neutralCreepsSatyrBanishers := neutralCreepsFactory("Satyr Banishers", 2, 9, 0, 300)
	neutralCreepsSatyrMindstealers := neutralCreepsFactory("Satyr Mindstealers", 2, 26, 1, 600)

	neutralCreepsMudGolems := neutralCreepsFactory("Mud Golems", 2, 33, 0, 800)
	neutralCreepsShardGolems := neutralCreepsFactory("Shard Golems", 4, 10, 0, 240)

	neutralCreepsOgreBruisers := neutralCreepsFactory("Ogre Bruisers", 2, 26, 1, 850)
	neutralCreepsOgreFrostmage := neutralCreepsFactory("Ogre Frostmage", 1, 26, 0, 600)

	//large camps
	neutralCreepsCentaurCoursersVolume2 := neutralCreepsFactory("Centaur Coursers", 2, 21, 1, 320)
	neutralCreepsCentaurConquerorVolume2 := neutralCreepsFactory("Centaur Conqueror", 1, 55, 4, 1100)

	neutralCreepsSatyrBanisherVolume2 := neutralCreepsFactory("Satyr Banisher", 1, 9, 0, 300)
	neutralCreepsSatyrMindstealerVolume2 := neutralCreepsFactory("Satyr Mindstealer", 1, 26, 1, 600)
	neutralCreepsSatyrTormenter := neutralCreepsFactory("Satyr Tormenter", 1, 55, 0, 1100)

	neutralCreepsHellbear := neutralCreepsFactory("Hellbear", 1, 44, 3, 700)
	neutralCreepsHellbearSmasher := neutralCreepsFactory("Hellbear Smasher", 1, 55, 4, 950)

	neutralCreepsWildwings := neutralCreepsFactory("Wildwings", 2, 25, 2, 350)
	neutralCreepsWildwingRipper := neutralCreepsFactory("Wildwing Ripper", 1, 56, 4, 950)

	neutralCreepsHillTrolls := neutralCreepsFactory("Hill Trolls", 2, 27, 0, 500)
	neutralCreepsDarkTrollSummoner := neutralCreepsFactory("Dark Troll Summoner", 1, 45, 1, 1100)

	//ancient camps
	neutralCreepsAncientBlackDrakes := neutralCreepsFactory("Ancient Black Drakes", 2, 45, -1, 950)
	neutralCreepsAncientBlackDragon := neutralCreepsFactory("Ancient Black Dragon", 1, 81, 3, 2000)

	neutralCreepsAncientRockGolems := neutralCreepsFactory("Ancient Rock Golems", 2, 33, 4, 800)
	neutralCreepsAncientGraniteGolem := neutralCreepsFactory("Ancient Granite Golem", 1, 87, 8, 1700)

	neutralCreepsAncientRumblehides := neutralCreepsFactory("Ancient Rumblehides", 2, 48, 2, 800)
	neutralCreepsAncientThunderhide := neutralCreepsFactory("Ancient Thunderhide", 1, 65, 2, 1400)

	neutralCreepsAncientProwlerAcolytes := neutralCreepsFactory("Ancient Prowler Acolytes", 2, 38, 10, 600)
	neutralCreepsAncientProwlerShaman := neutralCreepsFactory("Ancient Prowler Shaman", 1, 65, 11, 1200)

	//roshPit
	neutralCreepsRoshan := neutralCreepsFactory("Roshan", 1, 75, 20, 6000)

	rand.Seed(time.Now().UnixNano())
	smallCampOne = random(0, 100)
	smallCampTwo = random(0, 100)

	mediumCampOne = random(0, 100)
	mediumCampTwo = random(0, 100)
	mediumCampThree = random(0, 100)
	mediumCampFour = random(0, 100)
	mediumCampFive = random(0, 100)
	mediumCampSix = random(0, 100)

	largeCampOne = random(0, 100)
	largeCampTwo = random(0, 100)
	largeCampThree = random(0, 100)
	largeCampFour = random(0, 100)
	largeCampFive = random(0, 100)
	largeCampSix = random(0, 100)

	ancientCampOne = random(0, 100)
	ancientCampTwo = random(0, 100)
	ancientCampThree = random(0, 100)
	ancientCampFour = random(0, 100)

	start := GameStartFacade()

	start.CreepsActivity(neutralCreepsRoshan)

	start.CreepsActivity(direCreepsMelee)
	start.CreepsActivity(direCreepsRanged)

	start.CreepsActivity(radiantCreepsMelee)
	start.CreepsActivity(radiantCreepsRanged)

	start.BountyActivity()

	switch {
	case smallCampOne >= 0 && smallCampOne < 16:
		start.CreepsActivity(neutralCreepsKobolds)
		start.CreepsActivity(neutralCreepsKoboldSoldier)
		start.CreepsActivity(neutralCreepsKoboldForeman)
	case smallCampOne >= 16 && smallCampOne < 32:
		start.CreepsActivity(neutralCreepsHillTrollBerserkers)
		start.CreepsActivity(neutralCreepsHillTrollPriest)
	case smallCampOne >= 32 && smallCampOne < 48:
		start.CreepsActivity(neutralCreepsHillTrollBerserkersVolume2)
		start.CreepsActivity(neutralCreepsKoboldForemanVolume2)
	case smallCampOne >= 48 && smallCampOne < 64:
		start.CreepsActivity(neutralCreepsVhoulAssassins)
	case smallCampOne >= 64 && smallCampOne < 80:
		start.CreepsActivity(neutralCreepsFellSpirits)
		start.CreepsActivity(neutralCreepsGhost)
	case smallCampOne >= 80 && smallCampOne <= 100:
		start.CreepsActivity(neutralCreepsHarpyScouts)
		start.CreepsActivity(neutralCreepsHarpyStormcrafter)
	}

	switch {
	case smallCampTwo >= 0 && smallCampTwo < 16:
		start.CreepsActivity(neutralCreepsKobolds)
		start.CreepsActivity(neutralCreepsKoboldSoldier)
		start.CreepsActivity(neutralCreepsKoboldForeman)
	case smallCampTwo >= 16 && smallCampTwo < 32:
		start.CreepsActivity(neutralCreepsHillTrollBerserkers)
		start.CreepsActivity(neutralCreepsHillTrollPriest)
	case smallCampTwo >= 32 && smallCampTwo < 48:
		start.CreepsActivity(neutralCreepsHillTrollBerserkersVolume2)
		start.CreepsActivity(neutralCreepsKoboldForemanVolume2)
	case smallCampTwo >= 48 && smallCampTwo < 64:
		start.CreepsActivity(neutralCreepsVhoulAssassins)
	case smallCampTwo >= 64 && smallCampTwo < 80:
		start.CreepsActivity(neutralCreepsFellSpirits)
		start.CreepsActivity(neutralCreepsGhost)
	case smallCampTwo >= 80 && smallCampTwo <= 100:
		start.CreepsActivity(neutralCreepsHarpyScouts)
		start.CreepsActivity(neutralCreepsHarpyStormcrafter)
	}









	switch {
	case mediumCampOne >= 0 && mediumCampOne < 20:
		start.CreepsActivity(neutralCreepsGiantWolfs)
		start.CreepsActivity(neutralCreepsAlphaWolf)
	case mediumCampOne >= 20 && mediumCampOne < 40:
		start.CreepsActivity(neutralCreepsCentaurCourser)
		start.CreepsActivity(neutralCreepsCentaurConqueror)
	case mediumCampOne >= 40 && mediumCampOne < 60:
		start.CreepsActivity(neutralCreepsSatyrBanishers)
		start.CreepsActivity(neutralCreepsSatyrMindstealers)
	case mediumCampOne >= 60 && mediumCampOne < 80:
		start.CreepsActivity(neutralCreepsMudGolems)
		start.CreepsActivity(neutralCreepsShardGolems)
	case mediumCampOne >= 80 && mediumCampOne <= 100:
		start.CreepsActivity(neutralCreepsOgreBruisers)
		start.CreepsActivity(neutralCreepsOgreFrostmage)
	}

	switch {
	case mediumCampTwo >= 0 && mediumCampTwo < 20:
		start.CreepsActivity(neutralCreepsGiantWolfs)
		start.CreepsActivity(neutralCreepsAlphaWolf)
	case mediumCampTwo >= 20 && mediumCampTwo < 40:
		start.CreepsActivity(neutralCreepsCentaurCourser)
		start.CreepsActivity(neutralCreepsCentaurConqueror)
	case mediumCampTwo >= 40 && mediumCampTwo < 60:
		start.CreepsActivity(neutralCreepsSatyrBanishers)
		start.CreepsActivity(neutralCreepsSatyrMindstealers)
	case mediumCampTwo >= 60 && mediumCampTwo < 80:
		start.CreepsActivity(neutralCreepsMudGolems)
		start.CreepsActivity(neutralCreepsShardGolems)
	case mediumCampTwo >= 80 && mediumCampTwo <= 100:
		start.CreepsActivity(neutralCreepsOgreBruisers)
		start.CreepsActivity(neutralCreepsOgreFrostmage)
	}

	switch {
	case mediumCampThree >= 0 && mediumCampThree < 20:
		start.CreepsActivity(neutralCreepsGiantWolfs)
		start.CreepsActivity(neutralCreepsAlphaWolf)
	case mediumCampThree >= 20 && mediumCampThree < 40:
		start.CreepsActivity(neutralCreepsCentaurCourser)
		start.CreepsActivity(neutralCreepsCentaurConqueror)
	case mediumCampThree >= 40 && mediumCampThree < 60:
		start.CreepsActivity(neutralCreepsSatyrBanishers)
		start.CreepsActivity(neutralCreepsSatyrMindstealers)
	case mediumCampThree >= 60 && mediumCampThree < 80:
		start.CreepsActivity(neutralCreepsMudGolems)
		start.CreepsActivity(neutralCreepsShardGolems)
	case mediumCampThree >= 80 && mediumCampThree <= 100:
		start.CreepsActivity(neutralCreepsOgreBruisers)
		start.CreepsActivity(neutralCreepsOgreFrostmage)
	}

	switch {
	case mediumCampFour >= 0 && mediumCampFour < 20:
		start.CreepsActivity(neutralCreepsGiantWolfs)
		start.CreepsActivity(neutralCreepsAlphaWolf)
	case mediumCampFour >= 20 && mediumCampFour < 40:
		start.CreepsActivity(neutralCreepsCentaurCourser)
		start.CreepsActivity(neutralCreepsCentaurConqueror)
	case mediumCampFour >= 40 && mediumCampFour < 60:
		start.CreepsActivity(neutralCreepsSatyrBanishers)
		start.CreepsActivity(neutralCreepsSatyrMindstealers)
	case mediumCampFour >= 60 && mediumCampFour < 80:
		start.CreepsActivity(neutralCreepsMudGolems)
		start.CreepsActivity(neutralCreepsShardGolems)
	case mediumCampFour >= 80 && mediumCampFour <= 100:
		start.CreepsActivity(neutralCreepsOgreBruisers)
		start.CreepsActivity(neutralCreepsOgreFrostmage)
	}

	switch {
	case mediumCampFive >= 0 && mediumCampFive < 20:
		start.CreepsActivity(neutralCreepsGiantWolfs)
		start.CreepsActivity(neutralCreepsAlphaWolf)
	case mediumCampFive >= 20 && mediumCampFive < 40:
		start.CreepsActivity(neutralCreepsCentaurCourser)
		start.CreepsActivity(neutralCreepsCentaurConqueror)
	case mediumCampFive >= 40 && mediumCampFive < 60:
		start.CreepsActivity(neutralCreepsSatyrBanishers)
		start.CreepsActivity(neutralCreepsSatyrMindstealers)
	case mediumCampFive >= 60 && mediumCampFive < 80:
		start.CreepsActivity(neutralCreepsMudGolems)
		start.CreepsActivity(neutralCreepsShardGolems)
	case mediumCampFive >= 80 && mediumCampFive <= 100:
		start.CreepsActivity(neutralCreepsOgreBruisers)
		start.CreepsActivity(neutralCreepsOgreFrostmage)
	}

	switch {
	case mediumCampSix >= 0 && mediumCampSix < 20:
		start.CreepsActivity(neutralCreepsGiantWolfs)
		start.CreepsActivity(neutralCreepsAlphaWolf)
	case mediumCampSix >= 20 && mediumCampSix < 40:
		start.CreepsActivity(neutralCreepsCentaurCourser)
		start.CreepsActivity(neutralCreepsCentaurConqueror)
	case mediumCampSix >= 40 && mediumCampSix < 60:
		start.CreepsActivity(neutralCreepsSatyrBanishers)
		start.CreepsActivity(neutralCreepsSatyrMindstealers)
	case mediumCampSix >= 60 && mediumCampSix < 80:
		start.CreepsActivity(neutralCreepsMudGolems)
		start.CreepsActivity(neutralCreepsShardGolems)
	case mediumCampSix >= 80 && mediumCampSix <= 100:
		start.CreepsActivity(neutralCreepsOgreBruisers)
		start.CreepsActivity(neutralCreepsOgreFrostmage)
	}









	switch {
	case largeCampOne >= 0 && largeCampOne < 20:
		start.CreepsActivity(neutralCreepsCentaurCoursersVolume2)
		start.CreepsActivity(neutralCreepsCentaurConquerorVolume2)
	case largeCampOne >= 20 && largeCampOne < 40:
		start.CreepsActivity(neutralCreepsSatyrBanisherVolume2)
		start.CreepsActivity(neutralCreepsSatyrMindstealerVolume2)
		start.CreepsActivity(neutralCreepsSatyrTormenter)
	case largeCampOne >= 40 && largeCampOne < 60:
		start.CreepsActivity(neutralCreepsHellbear)
		start.CreepsActivity(neutralCreepsHellbearSmasher)
	case largeCampOne >= 60 && largeCampOne < 80:
		start.CreepsActivity(neutralCreepsWildwings)
		start.CreepsActivity(neutralCreepsWildwingRipper)
	case largeCampOne >= 80 && largeCampOne <= 100:
		start.CreepsActivity(neutralCreepsHillTrolls)
		start.CreepsActivity(neutralCreepsDarkTrollSummoner)
	}

	switch {
	case largeCampTwo >= 0 && largeCampTwo < 20:
		start.CreepsActivity(neutralCreepsCentaurCoursersVolume2)
		start.CreepsActivity(neutralCreepsCentaurConquerorVolume2)
	case largeCampTwo >= 20 && largeCampTwo < 40:
		start.CreepsActivity(neutralCreepsSatyrBanisherVolume2)
		start.CreepsActivity(neutralCreepsSatyrMindstealerVolume2)
		start.CreepsActivity(neutralCreepsSatyrTormenter)
	case largeCampTwo >= 40 && largeCampTwo < 60:
		start.CreepsActivity(neutralCreepsHellbear)
		start.CreepsActivity(neutralCreepsHellbearSmasher)
	case largeCampTwo >= 60 && largeCampTwo < 80:
		start.CreepsActivity(neutralCreepsWildwings)
		start.CreepsActivity(neutralCreepsWildwingRipper)
	case largeCampTwo >= 80 && largeCampTwo <= 100:
		start.CreepsActivity(neutralCreepsHillTrolls)
		start.CreepsActivity(neutralCreepsDarkTrollSummoner)
	}

	switch {
	case largeCampThree >= 0 && largeCampThree < 20:
		start.CreepsActivity(neutralCreepsCentaurCoursersVolume2)
		start.CreepsActivity(neutralCreepsCentaurConquerorVolume2)
	case largeCampThree >= 20 && largeCampThree < 40:
		start.CreepsActivity(neutralCreepsSatyrBanisherVolume2)
		start.CreepsActivity(neutralCreepsSatyrMindstealerVolume2)
		start.CreepsActivity(neutralCreepsSatyrTormenter)
	case largeCampThree >= 40 && largeCampThree < 60:
		start.CreepsActivity(neutralCreepsHellbear)
		start.CreepsActivity(neutralCreepsHellbearSmasher)
	case largeCampThree >= 60 && largeCampThree < 80:
		start.CreepsActivity(neutralCreepsWildwings)
		start.CreepsActivity(neutralCreepsWildwingRipper)
	case largeCampThree >= 80 && largeCampThree <= 100:
		start.CreepsActivity(neutralCreepsHillTrolls)
		start.CreepsActivity(neutralCreepsDarkTrollSummoner)
	}

	switch {
	case largeCampFour >= 0 && largeCampFour < 20:
		start.CreepsActivity(neutralCreepsCentaurCoursersVolume2)
		start.CreepsActivity(neutralCreepsCentaurConquerorVolume2)
	case largeCampFour >= 20 && largeCampFour < 40:
		start.CreepsActivity(neutralCreepsSatyrBanisherVolume2)
		start.CreepsActivity(neutralCreepsSatyrMindstealerVolume2)
		start.CreepsActivity(neutralCreepsSatyrTormenter)
	case largeCampFour >= 40 && largeCampFour < 60:
		start.CreepsActivity(neutralCreepsHellbear)
		start.CreepsActivity(neutralCreepsHellbearSmasher)
	case largeCampFour >= 60 && largeCampFour < 80:
		start.CreepsActivity(neutralCreepsWildwings)
		start.CreepsActivity(neutralCreepsWildwingRipper)
	case largeCampFour >= 80 && largeCampFour <= 100:
		start.CreepsActivity(neutralCreepsHillTrolls)
		start.CreepsActivity(neutralCreepsDarkTrollSummoner)
	}

	switch {
	case largeCampFive >= 0 && largeCampFive < 20:
		start.CreepsActivity(neutralCreepsCentaurCoursersVolume2)
		start.CreepsActivity(neutralCreepsCentaurConquerorVolume2)
	case largeCampFive >= 20 && largeCampFive < 40:
		start.CreepsActivity(neutralCreepsSatyrBanisherVolume2)
		start.CreepsActivity(neutralCreepsSatyrMindstealerVolume2)
		start.CreepsActivity(neutralCreepsSatyrTormenter)
	case largeCampFive >= 40 && largeCampFive < 60:
		start.CreepsActivity(neutralCreepsHellbear)
		start.CreepsActivity(neutralCreepsHellbearSmasher)
	case largeCampFive >= 60 && largeCampFive < 80:
		start.CreepsActivity(neutralCreepsWildwings)
		start.CreepsActivity(neutralCreepsWildwingRipper)
	case largeCampFive >= 80 && largeCampFive <= 100:
		start.CreepsActivity(neutralCreepsHillTrolls)
		start.CreepsActivity(neutralCreepsDarkTrollSummoner)
	}

	switch {
	case largeCampSix >= 0 && largeCampSix < 20:
		start.CreepsActivity(neutralCreepsCentaurCoursersVolume2)
		start.CreepsActivity(neutralCreepsCentaurConquerorVolume2)
	case largeCampSix >= 20 && largeCampSix < 40:
		start.CreepsActivity(neutralCreepsSatyrBanisherVolume2)
		start.CreepsActivity(neutralCreepsSatyrMindstealerVolume2)
		start.CreepsActivity(neutralCreepsSatyrTormenter)
	case largeCampSix >= 40 && largeCampSix < 60:
		start.CreepsActivity(neutralCreepsHellbear)
		start.CreepsActivity(neutralCreepsHellbearSmasher)
	case largeCampSix >= 60 && largeCampSix < 80:
		start.CreepsActivity(neutralCreepsWildwings)
		start.CreepsActivity(neutralCreepsWildwingRipper)
	case largeCampSix >= 80 && largeCampSix <= 100:
		start.CreepsActivity(neutralCreepsHillTrolls)
		start.CreepsActivity(neutralCreepsDarkTrollSummoner)
	}









	switch {
	case ancientCampOne >= 0 && ancientCampOne < 25:
		start.CreepsActivity(neutralCreepsAncientBlackDrakes)
		start.CreepsActivity(neutralCreepsAncientBlackDragon)
	case ancientCampOne >= 25 && ancientCampOne < 50:
		start.CreepsActivity(neutralCreepsAncientRockGolems)
		start.CreepsActivity(neutralCreepsAncientGraniteGolem)
	case ancientCampOne >= 50 && ancientCampOne < 75:
		start.CreepsActivity(neutralCreepsAncientRumblehides)
		start.CreepsActivity(neutralCreepsAncientThunderhide)
	case ancientCampOne >= 75 && ancientCampOne <= 100:
		start.CreepsActivity(neutralCreepsAncientProwlerAcolytes)
		start.CreepsActivity(neutralCreepsAncientProwlerShaman)
	}

	switch {
	case ancientCampTwo >= 0 && ancientCampTwo < 25:
		start.CreepsActivity(neutralCreepsAncientBlackDrakes)
		start.CreepsActivity(neutralCreepsAncientBlackDragon)
	case ancientCampTwo >= 25 && ancientCampTwo < 50:
		start.CreepsActivity(neutralCreepsAncientRockGolems)
		start.CreepsActivity(neutralCreepsAncientGraniteGolem)
	case ancientCampTwo >= 50 && ancientCampTwo < 75:
		start.CreepsActivity(neutralCreepsAncientRumblehides)
		start.CreepsActivity(neutralCreepsAncientThunderhide)
	case ancientCampTwo >= 75 && ancientCampTwo <= 100:
		start.CreepsActivity(neutralCreepsAncientProwlerAcolytes)
		start.CreepsActivity(neutralCreepsAncientProwlerShaman)
	}

	switch {
	case ancientCampThree >= 0 && ancientCampThree < 25:
		start.CreepsActivity(neutralCreepsAncientBlackDrakes)
		start.CreepsActivity(neutralCreepsAncientBlackDragon)
	case ancientCampThree >= 25 && ancientCampThree < 50:
		start.CreepsActivity(neutralCreepsAncientRockGolems)
		start.CreepsActivity(neutralCreepsAncientGraniteGolem)
	case ancientCampThree >= 50 && ancientCampThree < 75:
		start.CreepsActivity(neutralCreepsAncientRumblehides)
		start.CreepsActivity(neutralCreepsAncientThunderhide)
	case ancientCampThree >= 75 && ancientCampThree <= 100:
		start.CreepsActivity(neutralCreepsAncientProwlerAcolytes)
		start.CreepsActivity(neutralCreepsAncientProwlerShaman)
	}

	switch {
	case ancientCampFour >= 0 && ancientCampFour < 25:
		start.CreepsActivity(neutralCreepsAncientBlackDrakes)
		start.CreepsActivity(neutralCreepsAncientBlackDragon)
	case ancientCampFour >= 25 && ancientCampFour < 50:
		start.CreepsActivity(neutralCreepsAncientRockGolems)
		start.CreepsActivity(neutralCreepsAncientGraniteGolem)
	case ancientCampFour >= 50 && ancientCampFour < 75:
		start.CreepsActivity(neutralCreepsAncientRumblehides)
		start.CreepsActivity(neutralCreepsAncientThunderhide)
	case ancientCampFour >= 75 && ancientCampFour <= 100:
		start.CreepsActivity(neutralCreepsAncientProwlerAcolytes)
		start.CreepsActivity(neutralCreepsAncientProwlerShaman)
	}

	//visitor for places in dota

	//create axe














	action:=func(axe *heroBuilder) {
		axe.Name("Axe").
			MainAttribute("Strength").
			Tank(true).
			Skills([]string{"Berserker's call", "Battle hunger", "Counter helix", "Culling blade"}).
			Items([]string{"DESOLATOR", "Devine Rapier", "Devine Rapier", "Travel boots", "Devine Rapier", "Devine Rapier"}).
			IsRadiant("Radiant").
			Health(1000).
			Attack(50).
			Defence(10)
		AxeHero=&axe.hero
	}

	fmt.Println("")
	fmt.Println("")
	CreateHero(action)
	//create io


	action2:=func(io *heroBuilder) {
		io.Name("IO").
			MainAttribute("Strength").
			Tank(true).
			Skills([]string{"Tether", "Spirits", "Spirits Movement", "Overcharge"}).
			Items([]string{"DESOLATOR", "Devine Rapier", "Devine Rapier", "Travel boots", "Devine Rapier", "Devine Rapier"}).
			IsRadiant("Dire").
			Health(1000).
			Attack(50).
			Defence(10)
		Io=&io.hero

	}
	CreateHero(action2)


	fmt.Println("\nPlaces:")
	dotaland := new(DotaLand)

	dotaland.Add(&MainShop{"Radiant"})
	dotaland.Add(&HiddenShop{"Radiant"})

	//dotaland.Add(&SideShop{"Radiant"})
	dotaland.Add(&MainShop{"Dire"})
	dotaland.Add(&HiddenShop{"Dire"})

	//dotaland.Add(&SideShop{"Dire"})
	dotaland.Add(&Outpost{"Radiant"})
	dotaland.Add(&Outpost{"Dire"})
	RadiantMid:=Line{"Radiant","Mid"}//direCreepsBlyshniki.health*3+direCreepsDalniki.health
	dotaland.Add(&RadiantMid)
	dotaland.Add(&Forest{"Dire","Big"})//,neutralCreepsWolfs.health+neutralCreepsVozhak.health}
	Axe:=dotaland.Accept(AxeHero)

	fmt.Println("Axe Actions\n" + Axe)


	dotaland1 := new(DotaLand)

	Io:=dotaland1.Accept(Io)


	fmt.Println("Io\n"+Io)
	//jugger()

	fmt.Println(AxeHero)
	root := NewCreatureModifier(AxeHero)
	root.Add(NewDoubleDamageModifier(AxeHero))
	root.Add(NewFrostShieldModifier(AxeHero))
	root.Add(NewDoubleDamageModifier(AxeHero))
	root.Add(NewFrostShieldModifier(AxeHero))
	root.Apply()
	fmt.Println(AxeHero)
}