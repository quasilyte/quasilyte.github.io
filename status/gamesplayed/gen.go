package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"text/template"
)

func main() {
	var gamesPlayed = []gameInfo{
		{"SWAT 4", "FPS", 4.5},
		{"Splinter Cell", "FPS", 5.0},
		{"Quake 3 Arena", "FPS", 7.0},
		{"The Elder Scrolls III: Morrowind", "RPG", 8.0},
		{"The Elder Scrolls IV: Oblivion", "RPG", 6.5},
		{"Blood Bowl", "Turn-based Strategy, Sport", 6.0},
		{"Deus Ex", "FPS, RPG, Action", 8.5},
		{"Robinson's Requiem", "Adventure, Survival", 7.0},
		{"Return to Castle Wolfenstein", "FPS, Horror", 6.5},
		{"Postal", "Action, Sandbox", 6.0},
		{"Civilization 3", "Turn-based Strategy", 7.5},
		{"Civilization 4", "Turn-based Strategy", 7.0},
		{"Heroes of Might And Magic", "Turn-based Strategy", 5.5},
		{"Heroes of Might And Magic II", "Turn-based Strategy", 6.5},
		{"Heroes of Might And Magic III", "Turn-based Strategy", 7.0},
		{"Heroes of Might And Magic IV", "Turn-based Strategy", 7.0},
		{"Heroes of Might And Magic V", "Turn-based Strategy", 6.5},
		{"Star Control II (Ur-Quan Masters)", "Adventure, Arcade, RPG", 9.0},
		{"Space Rangers", "Adventure, Arcade, RPG", 7.5},
		{"Space Rangers 2", "Adventure, Arcade, RPG", 8.0},
		{"Warzone 2100", "Adventure, Arcade, RPG", 8.5},
		{"WarCraft II", "RTS", 7.0},
		{"WarCraft III", "RTS", 8.5},
		{"Praetorians", "RTS", 7.0},
		{"Dungeons & Dragons: Dragonshard", "RTS", 7.5},
		{"Dune 2000", "RTS", 7.0},
		{"Outlive", "RTS", 6.0},
		{"Command and Conquer: Red Alert", "RTS", 7.5},
		{"Command and Conquer: Red Alert 2", "RTS", 7.5},
		{"Command and Conquer: Tiberian Dawn", "RTS", 7.0},
		{"Command and Conquer: Generals", "RTS", 6.0},
		{"Sonic the Hedgehog", "Arcade, Platformer", 5.5},
		{"Syndicate Wars (psx)", "Arcade, RTS", 6.5},
		{"Fallout", "RPG, Adventure, Sandbox", 7.5},
		{"Fallout II", "RPG, Adventure, Sandbox", 7.0},
		{"X-Com", "Turn-based Strategy", 7.0},
		{"X-Com II", "Turn-based Strategy", 7.0},
		{"Extraterrestrials", "Turn-based Strategy", 6.0},
		{"Darkstone", "Rogue-like, RPG, Action", 6.5},
		{"Diablo", "Rogue-like, RPG, Action", 7.0},
		{"Diablo II", "Rogue-like, RPG, Action", 7.5},
		{"Path Of Exile", "Rogue-like, RPG, Action", 7.0},
		{"Hatoful Boyfriend", "Visual Novel", 4.5},
		{"Danganronpa", "Visual Novel, Detective", 7.5},
		{"Danganronpa 3", "Visual Novel, Detective", 8.0},
		{"Mount & Blade", "Action, RPG, Sandbox", 7.5},
		{"Desert Strike", "Arcade", 5.5},
		{"Urban Strike", "Arcade", 6.0},
		{"Nuclear Strike", "Arcade", 6.5},
		{"Age of Wonders", "Turn-based Strategy", 6.5},
		{"Dragon Shards", "Turn-based Strategy", 6.5},
		{"Arcanum", "RPG", 6.5},
		{"The Stanley Parable", "Walking Simulator", 7.0},
		{"Bloodline Champions", "Online, Action", 6.0},
		{"Hogs of War (PSX)", "Action, Strategy", 7.5},
		{"Worms: World Party", "Action, Strategy", 7.5},
		{"Worms: Armageddon", "Action, Strategy", 7.5},
		{"Worms 3D", "Action, Strategy", 7.0},
		{"Worms Ultimate Mayhem", "Action, Strategy", 7.5},
		{"Worms Revolution", "Action, Strategy", 7.0},
		{"Worms Reloaded", "Arcade", 7.5},
		{"Worms Clan Wars", "Action, Strategy", 6.5},
		{"Worms Battle Islands", "Action, Strategy", 6.5},
		{"Worms Battle Islands 2", "Action, Strategy", 6.5},
		{"Worms Blast", "Arcade", 5.5},
		{"Serious Sam", "FPS", 6.0},
		{"Serious Sam 2", "FPS", 5.5},
		{"The Entente", "RTS", 6.5},
		{"Mage Knight: Apocalypse", "Action, RPG", 5.5},
		{"Silverfall", "Action, RPG", 6.0},
		{"Sacred", "Action, RPG", 6.0},
		{"Sacred: Underworld", "Action, RPG", 7.0},
		{"Titan Quest", "Action, RPG", 6.5},
		{"Divine Divinity", "RPG", 8.0},
		{"Theme Hospital", "RTS, Puzzle", 5.5},
		{"Warhammer: Shadow of the Horned Rat", "RTS", 7.0},
		{"Warhammer: Dark Omen", "RTS", 8.0},
		{"Tanktics", "RTS, Puzzle", 7.0},
		{"Battle Realms", "RTS", 5.5},
		{"RPG Maker (PSX)", "Sandbox, RPG", 6.5},
		{"Army Men 3D (PSX)", "FPS", 6.5},
		{"Army Men Air Attack", "Arcade", 6.5},
		{"Carnage Heart EXA (PSP)", "Strategy, Engineering, Programming", 8.0},
		{"Robot Arena 2", "Arcade, Engineering", 7.5},
		{"Monster Hunter: Freedom Unite (PSP)", "Action, RPG, Souls-like", 7.5},
		{"Gods Eater Burst (PSP)", "Action, RPG, Souls-like", 7.0},
		{"Dungeon Maker: Hunting Grounds", "RPG", 6.0},
		{"Lineage II", "Online, MMO, RPG", 7.0},
		{"Guild Wars", "Online, MMO, RPG", 8.0},
		{"Neverwinter Nights", "RPG", 7.0},
		{"Neverwinter Nights 2", "RPG", 6.0},
		{"Carmageddon", "Racing", 6.5},
		{"Carmageddon 2", "Racing", 6.5},
		{"Fate", "RPG", 5.5},
		{"Dominions 5", "Global Strategy, Turn-based Strategy", 7.0},
		{"WarGame: Defcon (PSX)", "Arcade, FPS", 6.5},
		{"Star Wars: Battlefront II", "FPS", 6.0},
		{"Knights of Honor", "Global Strategy, RTS", 7.0},
		{"Need For Speed II", "Racing", 5.0},
		{"Need For Speed: Hot Pursuit 2", "Racing", 6.5},
		{"Need For Speed: Most Wanted", "Racing", 8.0},
		{"Need For Speed: Underground", "Racing", 7.0},
		{"Need For Speed: Underground 2", "Racing", 7.5},
		{"FlatOut", "Racing", 6.0},
		{"Gran Turismo (PSP)", "Racing", 6.5},
		{"GTA 1", "Arcade", 6.0},
		{"GTA 2", "Arcade", 6.5},
		{"GTA Vice City", "Action, Sandbox", 7.5},
		{"GTA San Andreas", "Action, Sandbox", 7.5},
		{"GTA Chinatown Wars (PSP)", "Action, Sandbox", 6.5},
		{"Jagged Alliance 2", "Turn-based Strategy", 7.5},
		{"Z (1996)", "RTS", 6.5},
		{"Resident Evil 2", "Arcade, Horror, Survival", 6.0},
		{"Resident Evil 3", "Arcade, Horror, Survival", 6.0},
		{"Doki Doki Literature Club", "Visual Novel, Horror", 8.5},
		{"Vampire: The Masquerade", "RPG, Horror", 7.0},
		{"Warhammer 40000: Dawn of War", "RTS", 8.0},
		{"Lord of the Rings: War of the Ring", "RTS", 7.5},
		{"Lord of the Rings: Battle for Middle Earth", "RTS", 8.0},
		{"Lord of the Rings: Battle for Middle Earth 2", "RTS", 8.5},
		{"Lord of the Rings: The Return of the King", "Action", 6.5},
		{"Mafia", "Action", 6.5},
		{"Oddworld: Abe's Oddysee", "Platformer, Puzzle", 7.0},
		{"Star Wars: Knights of the Old Republic 2", "RPG", 7.5},
		{"Counter Strike 1.6", "FPS", 6.5},
		{"Counter Strike Source", "FPS", 6.0},
		{"Breath of Fire IV", "Turn-based RPG", 6.5},
		{"Final Fantasy IX", "Turn-based RPG", 7.0},
		{"Spider-Man (PSX)", "Action", 6.5},
		{"SpellForce: The Order of Dawn", "RPG, RTS", 8.5},
		{"SpellForce: Breath of Winter", "RPG, RTS", 8.5},
		{"SpellForce: Shadow of the Phoenix", "RPG, RTS", 7.5},
		{"SpellForce 2", "RPG, RTS", 7.0},
		{"Disciples", "Turn-based Strategy", 7.0},
		{"Disciples 2", "Turn-based Strategy", 6.5},
		{"Warlords Battlecry", "RTS", 6.5},
		{"Warlords Battlecry III", "RTS", 7.5},
		{"Warlords IV: Heroes of Etheria", "Turn-based Strategy", 7.0},
		{"Test Drive: Overdrive", "Racing", 6.5},
		{"Delta Force: Land Warrior", "FPS", 5.0},
		{"MotorStorm: Arctic Edge (PSP)", "Racing", 7.0},
		{"Every Extend Extra", "Arcade", 7.0},
		{"What Did I Do to Deserve This, My Lord? (PSP)", "RTS", 7.0},
		{"What Did I Do to Deserve This, My Lord? 2 (PSP)", "RTS", 7.5},
		{"Twisted Metal 2", "Racing, Action", 6.0},
		{"Twisted Metal 4", "Racing, Action", 5.5},
		{"Twisted Metal: Head-On (PSP)", "Racing, Action", 6.0},
		{"Overlord", "Strategy, Action", 6.5},
		{"Rise of Nations", "RTS", 7.0},
		{"Empire Earth", "RTS", 6.5},
		{"Empire Earth II", "RTS", 7.5},
		{"Naruto: Ultimate Ninja Heroes 2 (PSP)", "Fighting", 6.5},
		{"Naruto: Ultimate Ninja Heroes 3 (PSP)", "Fighting", 7.0},
		{"Dungeon Siege", "RPG", 7.0},
		{"Dungeon Siege: Legends of Aranna", "RPG", 7.0},
		{"Dungeon Siege 2", "RPG", 7.5},
		{"Dungeon Siege 2: Broken World", "RPG", 7.5},
		{"Dungeon Siege: Throne of Agony (PSP)", "RPG", 5.5},
		{"Commandos: Behind Enemy Lines", "Strategy", 6.5},
		{"Stronghold: Crusader", "RTS", 7.5},
		{"Perimeter", "RTS", 6.5},
		{"Summoner", "RPG", 6.0},
		{"Command and Conquer: Renegade", "FPS", 6.0},
		{"Half-Life", "FPS", 6.0},
		{"Evil Genius", "RTS", 6.5},
		{"Blitzkrieg: Rolling Thunder", "Strategy", 6.0},
		{"Fable: The Lost Chapters", "RPG", 7.5},
		{"PuzzleQuest: Challenge of the Warlords", "RPG, Puzzle, Match-Three", 7.5},
		{"Prince of Persia (GBA)", "Platformer", 6.5},
		{"Sims 2", "Sandbox", 5.5},
		{"DCS: Black Shark", "Arcade", 6.0},
		{"League of Legends", "Online, MOBA", 7.0},
		{"DoTA 2", "Online, MOBA", 6.0},
		{"Baba Is You", "Puzzle, Programming", 8.5},
		{"Team Buddies", "RTS, Action", 8.0},
		{"Cladun", "Action, RPG", 7.0},
		{"Cladun Returns", "Action, RPG", 7.0},
		{"Roboden", "RTS", 6.5},
		{"No Heroes Allowed", "Puzzle, RTS", 7.5},
		{"Monster Hunter Rise", "Action, RPG, Souls-like", 8.5},
		{"Dragon Age: Origins", "RPG", 8.0},
		{"King's Bounty (2008)", "Turn-based Strategy", 7.5},
	}

	var gamesWatched = []gameInfo{
		{"Galerians: Ash", "Action, Horror", 6.0},
		{"The Orion Conspiracy", "Point and Click", 6.5},
		{"The Bad Rats", "Puzzle", 6.0},
		{"The Elder Scrolls II: Daggerfall", "RPG", 6.0},
		{"The Elder Scrolls V: Skyrim", "RPG", 7.0},
		{"Shadow of the Colossus", "RPG", 7.0},
		{"BioShock Infinite", "FPS", 6.5},
		{"Tales of Vesperia", "Action, RPG", 8.0},
		{"Fort Zombie", "FPS, Survival", 5.5},
		{"Sanitarium", "Point and Click", 7.5},
		{"Primordia", "Point and Click", 7.0},
		{"Beneath a Steel Sky", "Point and Click", 5.5},
		{"Phantasmagoria", "FMV, Point and Click", 6.0},
		{"Recettear", "RPG, Vendor Simulator", 5.5},
		{"Blade Runner: The Game", "Adventure, Point and Click", 6.5},
		{"Deus Ex: Human Revolution", "FPS, RPG, Action", 7.5},
		{"Deus Ex: Mankind Devided", "FPS, RPG, Action", 7.0},
		{"Hitman: Blood Money", "FPS, Puzzle", 7.5},
		{"Metal Gear Solid 3: Snake Eater", "Action, Stealth", 7.0},
		{"Tokobot", "Puzzle, Platformer", 7.5},
		{"Way of the Samurai 4", "Action, RPG, Sandbox", 6.5},
		{"SOMA", "Adventure, Horror", 7.5},
		{"Alan Wake", "FPS, Horror", 6.0},
		{"Mystical Ninja Starring Goemon", "Action", 5.5},
		{"Dead Space", "FPS, Horror, Survival", 7.0},
		{"Dead Space 2", "FPS, Horror, Survival", 6.5},
		{"Dead Space 3", "FPS, Horror, Survival", 6.5},
		{"Dishonored", "Action, RPG, Adventure, Sandbox", 7.0},
		{"Dishonored 2", "Action, RPG, Adventure, Sandbox", 6.5},
		{"Genuine Blast Corps", "Racing, Action", 5.0},
		{"Mega Man Legends", "Action", 7.5},
		{"Mega Man Legends 2", "Action", 7.5},
		{"The Misadventures of Tron Bonn", "Action", 7.0},
		{"Brave Fencer Musashi", "Action", 5.0},
		{"The Turing Test", "Puzzle", 6.0},
		{"Contact (NDS)", "RPG", 6.5},
		{"Made In Abyss: Binary Star Falling Into Darkness", "Action, Adventure, RPG, Sandbox", 6.0},
		{"Endoparasitic", "Action, Horror", 6.5},
		{"Backpack Hero", "RPG, Rogue-like", 7.0},
		{"Master of Orion", "Turn-based Strategy", 6.5},
		{"Master of Orion II", "Turn-based Strategy", 7.0},
		{"5 Days A Stranger", "Point and Click", 4.5},
		{"7 Days A Skeptic", "Point and Click", 5.0},
		{"Trilby's Notes", "Point and Click", 5.0},
		{"6 Days A Sacrifice", "Point and Click", 5.0},
		{"The Sandbox of God", "Sandbox", 3.5},
		{"I Have No Mouth, and I Must Scream", "Point and Click, Horror", 8.0},
		{"The Dig", "Point and Click, Horror", 6.5},
		{"Myth II: Soulblighter", "RTS", 8.0},
		{"System Shock", "FPS, RPG", 7.5},
		{"System Shock 2", "FPS, RPG", 7.0},
		{"Harvester", "Point and Click, Horror", 6.5},
		{"Alone In the Dark", "Point and Click, Horror", 5.0},
		{"Faster Than Light", "Rogue-like, Strategy, Action", 7.0},
		{"Tex Murphy: Pandora Directive", "Point and Click", 6.5},
		{"Danganronpa 2", "Visual Novel, Detective", 7.0},
		{"To The Moon", "Point and Click", 7.5},
		{"Finding Paradise", "Point and Click", 7.0},
		{"Impostor Factory", "Point and Click", 6.5},
		{"Monster Loves You", "Sandbox", 5.0},
		{"Unholy Heights", "Strategy", 6.5},
		{"Miasmata", "Adventure, Horror", 5.0},
		{"Little Inferno", "Sandbox, Puzzle", 5.5},
		{"Doom", "FPS", 5.5},
		{"Might And Magic 7", "RPG", 5.5},
		{"Little Big Adventure", "Point and Click", 6.0},
		{"Crusader: No Remorse", "Action", 6.0},
		{"King's Bounty", "Turn-based Strategy", 6.5},
		{"Late Shift", "FMV", 6.0},
		{"Until Dawn", "Action, Walking Simulator, Horror", 7.0},
		{"The Quarry", "Action, Walking Simulator, Horror", 6.5},
		{"Gone Rogue", "Strategy, Puzzle", 6.0},
		{"Resident Evil", "Arcade, Horror, Survival", 5.5},
		{"Resident Evil 4", "FPS, Horror, Survival", 6.0},
		{"Resident Evil VII", "Action, Horror, Survival", 7.5},
		{"L.A. Noire", "Detective, Action", 8.0},
		{"TUNIC", "Arcade, RPG, Action, Puzzle", 7.5},
		{"Undertale", "RPG, Puzzle", 8.5},
		{"Deltarune", "RPG, Puzzle", 7.0},
		{"Life Is Strange", "Walking Simulator", 8.0},
		{"Life Is Strange: Before The Storm", "Walking Simulator", 6.5},
		{"Life Is Strange 2", "Walking Simulator", 6.5},
		{"Her Story", "FMV, Puzzle", 6.5},
		{"Tales from the Borderlands", "Walking Simulator", 7.5},
		{"Wolf Among Us", "Walking Simulator", 7.5},
		{"Psychonauts", "Arcade, Platformer", 6.5},
		{"LISA: The First", "RPG", 5.0},
		{"LISA: The Painful RPG", "RPG", 8.0},
		{"LISA: The Joyful", "RPG", 7.0},
		{"Oxenfree", "Puzzle, Adventure", 7.0},
		{"Firewatch", "Walking Simulator", 6.5},
		{"The Walking Dead", "Walking Simulator", 7.0},
		{"Fahrenheit: Indigo Prophecy", "Interactive Movie, Detective", 6.5},
		{"Heavy Rain", "Interactive Movie, Detective", 7.5},
		{"Detroit: Become Human", "Interactive Movie, Detective", 8.5},
		{"Beyond: Two Souls", "Interactive Movie, Detective", 6.5},
		{"The Last of Us", "FPS, Horror", 8.0},
		{"OneShot", "Point and Click", 6.0},
		{"Hidden Agenda", "Interactive Movie", 6.0},
		{"The Witcher", "RPG", 6.5},
		{"The Witcher 2", "RPG", 7.0},
		{"Night in the Woods", "Platformer, Puzzle", 6.5},
		{"The Cat Lady", "Point and Click, Horror", 8.5},
		{"Downfall", "Point and Click, Horror", 7.0},
		{"Lorelai", "Point and Click, Horror", 6.5},
		{"The Council", "RPG, Talking Simulator, Puzzle", 7.0},
		{"Forgotton Anne", "Point and Click", 6.5},
		{"State of Mind", "Walking Simulator", 6.5},
		{"Unawoved", "Point and Click", 7.0},
		{"The Hex", "Arcade, Platformer, Puzzle", 8.0},
		{"Distraint", "Point and Click, Platformer", 6.5},
		{"Distraint 2", "Point and Click, Platformer", 7.0},
		{"Return of the Obra Dinn", "Detective, Puzzle", 9.0},
		{"Draugen", "Walking Simulator", 5.5},
		{"Outer Wilds", "Adventure, Puzzle", 8.5},
		{"Outer Wilds: Echoes of the Eye", "Adventure, Puzzle", 8.0},
		{"The Sinking City", "FPS, Adventure, Detective", 6.5},
		{"Little Misfortune", "Point and Click", 7.0},
		{"Telling Lies", "FMV", 6.5},
		{"Fran Bow", "Point and Click", 7.0},
		{"Afterparty", "Point and Click", 6.0},
		{"Not For Broadcast", "FMV", 7.5},
		{"Hades", "Action, RPG, Rogue-like", 7.5},
		{"Twelve Minutes", "Puzzle", 7.0},
		{"Inscryption", "Strategy, Deckbuilder", 7.0},
		{"The Forgotten City", "Talking Simulator, Detective", 7.5},
		{"As Dusk Falls", "Walking Simulator", 7.0},
		{"Burnhouse Lane", "Puzzle, Horror", 7.0},
		{"Signalis", "Arcade, Horror, Survival", 7.5},
		{"Papers, Please", "Puzzle", 8.0},
		{"Darkest Dungeon", "Turn-based RPG", 7.0},
		{"Stray", "Platformer, Arcade", 6.0},
		{"McPixel", "Puzzle", 6.0},
		{"Broken Age", "Point and Click", 7.0},
		{"South Park: The Stick of Truth", "RPG", 6.5},
		{"South Park: Fractured But Whole", "RPG", 7.0},
		{"Yesterday", "Point and Click", 7.0},
		{"Yesterday Origins", "Point and Click", 7.0},
		{"Shovel Knight", "Platformer, Arcade", 7.0},
		{"The Shapeshifting Detective", "FMV, Detective", 6.5},
		{"Contradiction", "FMV, Detective", 7.0},
		{"The Complex", "FMV", 5.5},
		{"Dark Nights With Poe & Munro", "FMV", 6.5},
		{"Murdered: Soul Suspect", "Detective", 6.5},
		{"Observer", "FPS, Horror", 7.5},
		{"Bulb Boy", "Point and Click", 6.0},
		{"The Infectious Madness of Doctor Dekker", "FMV", 7.5},
		{"Tesla Effect: A Tex Murphy Adventure", "Puzzle, Adventure, FMV, Detective", 7.0},
		{"World of Horror", "Turn-based RPG, Horror, Rogue-like", 7.0},
		{"Twin Mirror", "Walking Simulator", 6.5},
		{"The X-Files", "FMV, Detective", 6.5},
		{"Fear Effect", "Action", 7.0},
		{"The Hobbit", "Action", 6.5},
		{"Double Switch", "FMV", 5.0},
		{"Bully", "Action, Sandbox", 7.5},
		{"The Bouncer", "Fighting", 5.0},
		{"Steambot Chronicles", "Action, RPG", 7.0},
		{"Max Payne", "FPS", 7.5},
		{"Spore", "Action, Adventure", 7.0},
		{"Raft", "Sandbox, Adventure, Survival", 6.5},
		{"Vangers", "Racing, Adventure, Arcade, Sandbox", 6.5},
		{"Gothic", "RPG, Adventure", 7.0},
		{"Gothic 2", "RPG, Adventure", 7.0},
		{"Gothic 3", "RPG, Adventure", 5.5},
		{"Risen", "RPG, Adventure", 6.0},
		{"Call of Cthulhu: Shadow of the Comet", "Point and Click", 5.5},
		{"Tokyo Jungle", "Action, Survival", 6.0},
		{"Among the Sleep", "Walking Simulator", 4.5},
		{"This Means Warp", "Rogue-like, Strategy, Action", 5.5},
		{"Cult of the Lamb", "Rogue-like, RPG, Action", 7.0},
		{"Devil in Me", "FMV, Horror, Walking Simulator", 6.5},
		{"House of Ashes", "FMV, Horror, Walking Simulator", 7.0},
		{"Little Hope", "FMV, Horror, Walking Simulator", 5.5},
		{"Man of Medan", "FMV, Horror, Walking Simulator", 5.5},
		{"Beginners Guide", "Walking Simulator", 6.5},
		{"Ghost Trick: Phantom Detective", "Puzzle, Point and Click", 7.5},
	}

	var gamesBarelyKnow = []gameInfo{
		{"Dark Souls", "Action, RPG, Souls-like", 8.0},
		{"Minecraft", "Action, RPG, Souls-like", 6.0},
		{"Dwarf Fortress", "Sandbox", 5.5},
	}

	sort.SliceStable(gamesPlayed, func(i, j int) bool {
		return gamesPlayed[i].Rating > gamesPlayed[j].Rating
	})
	sort.SliceStable(gamesWatched, func(i, j int) bool {
		return gamesWatched[i].Rating > gamesWatched[j].Rating
	})
	sort.SliceStable(gamesBarelyKnow, func(i, j int) bool {
		return gamesBarelyKnow[i].Rating > gamesBarelyKnow[j].Rating
	})

	var allGames []fullGameInfo
	for _, g := range gamesPlayed {
		allGames = append(allGames, fullGameInfo{
			Name:      g.Name,
			GenreTags: splitTags(g.GenreTags),
			Rating:    g.Rating,
			Status:    statusPlayedMyself,
		})
	}
	for _, g := range gamesWatched {
		allGames = append(allGames, fullGameInfo{
			Name:      g.Name,
			GenreTags: splitTags(g.GenreTags),
			Rating:    g.Rating,
			Status:    statusWatchedLetsplay,
		})
	}
	for _, g := range gamesBarelyKnow {
		allGames = append(allGames, fullGameInfo{
			Name:      g.Name,
			GenreTags: splitTags(g.GenreTags),
			Rating:    g.Rating,
			Status:    statusWatchedReview,
		})
	}

	tagCounters := map[string]int{}
	tagRatings := map[string]float64{}
	knownGenres := map[string]bool{
		"Action":              true,
		"Adventure":           true,
		"Arcade":              true,
		"Online":              true,
		"MMO":                 true,
		"RPG":                 true,
		"Turn-based RPG":      true,
		"FPS":                 true,
		"RTS":                 true,
		"Visual Novel":        true,
		"Detective":           true,
		"Turn-based Strategy": true,
		"Strategy":            true,
		"Sandbox":             true,
		"Rogue-like":          true,
		"Souls-like":          true,
		"Survival":            true,
		"Horror":              true,
		"Walking Simulator":   true,
		"Sport":               true,
		"Platformer":          true,
		"Puzzle":              true,
		"Point and Click":     true,
		"FMV":                 true,
		"Vendor Simulator":    true,
		"Stealth":             true,
		"Racing":              true,
		"Interactive Movie":   true,
		"Talking Simulator":   true,
		"Deckbuilder":         true,
		"Fighting":            true,
		"Global Strategy":     true,
		"Engineering":         true,
		"Programming":         true,
		"Match-Three":         true,
		"MOBA":                true,
	}
	gamesListed := map[string]bool{}
	for i := range allGames {
		g := &allGames[i]
		if gamesListed[g.Name] {
			panic(fmt.Sprintf("%q listed more than once", g.Name))
		}
		gamesListed[g.Name] = true
		for _, tag := range g.GenreTags {
			if !knownGenres[tag] {
				panic(fmt.Sprintf("%q: unexpected genre tag: %q", g.Name, tag))
			}
			tagCounters[tag]++
			tagRatings[tag] += g.Rating
		}
		sort.Strings(g.GenreTags)
		g.GenreTagString = strings.Join(g.GenreTags, ", ")
	}

	type ratedTagItem struct {
		Name   string
		Count  int
		Rating float64
	}
	var ratedTagList []ratedTagItem
	for name, count := range tagCounters {
		ratedTagList = append(ratedTagList, ratedTagItem{
			Name:   name,
			Count:  count,
			Rating: tagRatings[name] / float64(count),
		})
	}
	sort.Slice(ratedTagList, func(i, j int) bool {
		return ratedTagList[i].Count > ratedTagList[j].Count
	})

	var buf bytes.Buffer
	resultTemplate.Execute(&buf, map[string]any{
		"AllGames":     allGames,
		"Tags":         ratedTagList,
		"GamesPlayed":  len(gamesPlayed),
		"GamesWatched": len(gamesWatched),
		"GamesKnown":   len(gamesBarelyKnow),
	})
	fmt.Println(buf.String())
}

func splitTags(s string) []string {
	tags := strings.Split(s, ",")
	for i := range tags {
		tags[i] = strings.TrimSpace(tags[i])
	}
	return tags
}

type gameInfo struct {
	Name      string
	GenreTags string
	Rating    float64
}

type playStatus int

const (
	statusWatchedReview playStatus = iota
	statusWatchedLetsplay
	statusPlayedMyself
)

func (s playStatus) String() string {
	switch s {
	case statusWatchedReview:
		return "Watched review/know about"
	case statusWatchedLetsplay:
		return "Completed via YouTube"
	default:
		return "Completed myself"
	}
}

type fullGameInfo struct {
	Name           string
	GenreTags      []string
	GenreTagString string
	Rating         float64
	Status         playStatus
}

var resultTemplate = template.Must(template.New("").Parse(`
## Games Played

There are 3 possible game title status here:

1. Completed myself.
2. Completed via let's play (e.g. YouTube).
3. Extra category for the games I can discuss, but I haven't completed them.

I usually try not to add an entry unless I either completed the game or feeling competent enough to include it at least in the 3rd category.

Also, this list is never complete. I can't possibly remember every game I have ever played, but I'm doing my best here.

| Game | Genre Tags | Status | Rating |
|---|---|---|---|
{{- range $.AllGames}}
| {{.Name}} | {{.GenreTagString}} | {{.Status}} | {{.Rating}} |
{{- end}}

Here are some tables for stats.

| Tag | Count | Avg. Rating |
|---|---|---|
{{- range $.Tags}}
| {{.Name}} | {{.Count}} | {{printf "%.2f" .Rating}} |
{{- end}}

| Status | Count |
|---|---|
| Completed myself | {{$.GamesPlayed}} |
| Completed via YouTube | {{$.GamesWatched}} |
| Watched review/know about | {{$.GamesKnown}} |

### [itch.io](https://itch.io/) indie games

Sometimes I [play random games](https://www.youtube.com/playlist?list=PLAjl-3QkinHsHaEnaie_FFG1sI8KU2Eoq) from [itch.io](https://itch.io/). I wonder, how many games from this list you ever heard of.

| Game | Short summary |
|---|---|
| [Mandrake](https://osakitsukiko.itch.io/mandrake) | A turn-based RPG game with a very simple combat mechanics and counter-intuitive character progression. The artstyle is great. |
| [Samoai](https://reyoruga.itch.io/samoai) | A 2D fighting platformer; it's not a finished game, but it can be entertaining for 30 minutes or so. |
| [Nebula Mining Corp](https://lightningspell.itch.io/nebula-mining-corp) | A game with a font that is very hard to read. The game itself lacks difficulty progression. The graphics are good. |
| [Soul of Forest](https://blayz-by-fill.itch.io/soul-of-the-forest) | A very simple 2D platformer focused around jumping puzzles. |
| [Striving For Light: Survival](https://ignitingsparkgames.itch.io/strivingforlight-survival) | A well-made survival-style game, but its graphical style is not my cup of tea. |
| [Astra Defender](https://yokereba.itch.io/astra-defender) | An interesting top-down shooter combined with an RTS. Try it out. |
| [Verdant](https://justinmullin.itch.io/verdant) | A good tower defense game; it's a very solid game jam entry, it feels quite polished too. |
| [Huck and Hubble](https://chefpancake.itch.io/hucknhubble) | A well-drawn game with OK concepts, but the execution is quite annoying. This game is hard to navigate, the terrain is sticky. |
| [Love At First Sight](https://mrmathur.itch.io/love-at-first-sight) | A very simple and short game focused around "don't look here" mechanic. |
| [FiBunnyCci into the Yilon-Verse](https://piccoloplay.itch.io/fibunnycci-into-the-yilon-verse) | This game is weird, but maybe for its own good. The author put a lot of efforts to create things around the game, like wikis, etc. |
| [Graveyard Shift](https://badcop.itch.io/graveyard-shift) | Now this is a good stuff. A nice card puzzle game with great presentation. |
| [Elephant In The Room](https://epicramen3d.itch.io/elephantintheroom) | A Hotline Miami-like game. It could be better if collision boxes were more precise; the instant shooting reaction is weird too. |
| [Golden Warden](https://xpii.itch.io/golden-warden) | A nice game that plays like a turn-based combat game. It's not unique, but it's still refreshing enough. |
| [KittInn](https://pastel-cherry-games.itch.io/kittinn) | This game is too buggy to be described, but the idea is that you're running a kitty hotel. It's an OK start. |
| [Weak Sbot](https://blek.itch.io/weak-sbot) | A game where you fight several bosses. It's interesting and good-looking, but the difficulty balance might be slightly off. |
| [Cliche Hunter](https://njitram.itch.io/cliche-hunter) | Another game from a Cliche-themed game jam. It's a game full of gags. Gameplay-wise, it's mostly a 2D platformer game. |
| [Blasteroids](https://kharibidus.itch.io/blasteroids) | Remember a classic Asteroids game? Well, in this game, you play as asteroids; you destroy spaceships. |
| [Shoot Into The Void](https://gameboymarcus.itch.io/shoot-into-the-void) | A game where you have to memorize your target locations and then direct a projectile through the darkness to get them. |
| [Escort Mission](https://sinisterstuf.itch.io/escort-mission) | A top-down shooter build around the "escort mission" cliche concept. It's good for a single run. |
| [It's But a Scratch](https://blazingedgegames.itch.io/tis-but-a-scratch) | A well-styled top-down shooter game, if a bit too easy. The sound design is great. |
| [You Catch My Drift?](https://beanfox.itch.io/you-catch-my-drift) | A simple rally-style racing game with a split-screen support. I played a couple of races with my friend. |
| [Space Survivors](https://roaming-maelstrom.itch.io/space-survivors) | Like a Vampire Survivors, but with space vessels. It's OK. |
| [Uphill Battle](https://jaredn.itch.io/uphill-battle) | A card battling game. It's quite nice for a game jam submission. |
| [In the Bag](https://sandramoen.itch.io/in-the-bag) | I can actually recommend this game. It could use some post-jam polish to become a good game, it has the potential. |
| [It's not a bug it's a feature](https://partyvaper.itch.io/its-not-a-bug-its-a-feature) | A game that required me to go through its sources as a gameplay mechanic. The graphics are neat too. |
| [Curious Cats](https://e-potapova.itch.io/curious-cats) | It's like Lemmings, but with cats. It's good for a jam game. |
| [Space Defenders](https://irywo.itch.io/space-defenders) | Probably one of the most simplistic games I played on the record. But hey, it's a game jam submission. |
| [Specter X-terminators](https://kblm-studio.itch.io/specter-x-terminators) | A game inspired by Undertale. The execution is good (for a jam prototype), but the game is really-really short. |
| [Zero Zero UFO](https://kronoman.itch.io/zero-zero-ufo) | A cool (but silent) retro arcade made for DOS! I played this game on Linux using DosBox software. |
| [Glass Jungle](https://polanski.itch.io/glass-jungle) | A cheesy shooter made for a "Cliche" game jam. It's full of cliches. |
| [Geometry Battles](https://hyperstellar-games.itch.io/geometry-battles) | A well-made RTS game with minimalistic graphics. The AI could be better though. |
| [Danmaku](https://zergon321.itch.io/touhou-game-in-go) | A Touhou-inspired bullet hell game. |
| [Violet's Awakening](https://rafaskb.itch.io/violets-awakening) | A game that reminds me of TUNIC, but it's quite broken and unfinished. |
| [Fairyside](https://kiyummi.itch.io/fairyside) | Plays like a cute and lighter version of Binding of Isaac. |
`))
