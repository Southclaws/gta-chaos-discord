package main

import "fmt"

type Effect interface {
	Name() string
	ID() string
}

func EffectToMessage(effect Effect) string {
	switch e := effect.(type) {
	case FunctionEffect:
		return fmt.Sprintf("effect:%s:%d:%s", e.id, 30000, e.name)
	case WeatherEffect:
		return fmt.Sprintf("weather:%d:%d:%s", e.weather, 30000, e.name)
	case SpawnVehicleEffect:
		return fmt.Sprintf("spawn_vehicle:%d:%d:%s", e.model, 30000, e.id)
	case TeleportationEffect:
		return fmt.Sprintf("teleport:%d,%d,%d", int(e.location.x), int(e.location.y), int(e.location.z))
	}
	return ""
}

type FunctionEffect struct {
	name  string
	name2 string
	id    string
	mult  float64
}

func (e FunctionEffect) Name() string { return e.name }
func (e FunctionEffect) ID() string   { return e.id }

type WeatherEffect struct {
	name    string
	id      string
	weather int
}

func (e WeatherEffect) Name() string { return e.name }
func (e WeatherEffect) ID() string   { return e.id }

type SpawnVehicleEffect struct {
	id    string
	model int
}

func (e SpawnVehicleEffect) Name() string { return e.id }
func (e SpawnVehicleEffect) ID() string   { return e.id }

type Location struct{ x, y, z float32 }

var GrooveStreet Location = Location{2493, -1670, 15}
var LSTower Location = Location{1544, -1353, 332}
var LSPier Location = Location{836, -2061, 15}
var LSAirport Location = Location{2109, -2544, 16}
var LSDocks Location = Location{2760, -2456, 16}
var MountChiliad Location = Location{-2233, -1737, 483}
var SFAirport Location = Location{-1083, 409, 17}
var SFBridge Location = Location{-2669, 1595, 220}
var Area52 Location = Location{213, 1911, 20}
var LVQuarry Location = Location{614, 856, -40}
var LVAirport Location = Location{1612, 1166, 17}
var LVSatellite Location = Location{-310, 1524, 78}

type TeleportationEffect struct {
	name     string
	id       string
	location Location
}

func (e TeleportationEffect) Name() string { return e.name }
func (e TeleportationEffect) ID() string   { return e.id }

func Effects() []Effect {
	return []Effect{
		FunctionEffect{"Weapon Set 1", "ThugsArmoury", "weapon_set_1", 0.0},
		FunctionEffect{"Weapon Set 2", "ProfessionalsKit", "weapon_set_2", 0.0},
		FunctionEffect{"Weapon Set 3", "NuttersToys", "weapon_set_3", 0.0},
		FunctionEffect{"Weapon Set 4", "MinigunMadness", "weapon_set_4", 0.0},
		FunctionEffect{"Health, Armor, $250k", "INeedSomeHelp", "health_armor_money", 0.0},
		FunctionEffect{"Suicide", "GoodbyeCruelWorld", "suicide", 0.0},
		FunctionEffect{"Infinite Ammo", "FullClip", "infinite_ammo", 0.0},
		FunctionEffect{"Infinite Health (Player)", "NoOneCanHurtMe", "infinite_health", 0.0},
		FunctionEffect{"Wanted Level +2 Stars", "TurnUpTheHeat", "wanted_plus_two", 0.0},
		FunctionEffect{"Clear Wanted Level", "TurnDownTheHeat", "clear_wanted", 0.0},
		FunctionEffect{"Never Wanted", "IDoAsIPlease", "never_wanted", 0.0},
		FunctionEffect{"Six Wanted Stars", "BringItOn", "wanted_six_stars", 0.0},
		FunctionEffect{"Get Parachute", "LetsGoBaseJumping", "get_parachute", 0.0},
		FunctionEffect{"Get Jetpack", "Rocketman", "get_jetpack", 0.0},
		FunctionEffect{"0.25x Game Speed", "MatrixMode", "quarter_gamespeed", 1.0 / 3.0},
		FunctionEffect{"0.5x Game Speed", "SlowItDown", "half_gamespeed", 2.0 / 3.0},
		FunctionEffect{"2x Game Speed", "SpeedItUp", "double_gamespeed", 0.0},
		FunctionEffect{"4x Game Speed", "YoureTooSlow", "quadruple_gamespeed", 0.0},
		FunctionEffect{"Always Midnight", "NightProwler", "always_midnight", 0.0},
		FunctionEffect{"Stop Game Clock", "DontBringOnTheNight", "stop_game_clock", 0.0},
		FunctionEffect{"Faster Clock", "TimeJustFliesBy", "faster_clock", 0.0},
		FunctionEffect{"Blow Up All Cars", "AllCarsGoBoom", "blow_up_all_cars", 0.0},
		FunctionEffect{"Pink Traffic", "PinkIsTheNewCool", "pink_traffic", 0.0},
		FunctionEffect{"Black Traffic", "SoLongAsItsBlack", "black_traffic", 0.0},
		FunctionEffect{"Cheap Cars", "EveryoneIsPoor", "cheap_cars", 0.0},
		FunctionEffect{"Expensive Cars", "EveryoneIsRich", "expensive_cars", 0.0},
		FunctionEffect{"Insane Handling", "StickLikeGlue", "insane_handling", 0.0},
		FunctionEffect{"All Green Lights", "DontTryAndStopMe", "all_green_lights", 0.0},
		FunctionEffect{"Cars On Water", "JesusTakeTheWheel", "cars_on_water", 0.0},
		FunctionEffect{"Boats Fly", "FlyingFish", "boats_fly", 0.0},
		FunctionEffect{"Cars Fly", "ChittyChittyBangBang", "cars_fly", 0.0},
		FunctionEffect{"Smash N' Boom", "TouchMyCarYouDie", "smash_n_boom", 0.0},
		FunctionEffect{"All Cars Have Nitro", "SpeedFreak", "all_cars_have_nitro", 0.0},
		FunctionEffect{"Cars Float Away When Hit", "BubbleCars", "cars_float_away_when_hit", 0.0},
		FunctionEffect{"All Taxis Have Nitrous", "SpeedyTaxis", "all_taxis_have_nitro", 0.0},
		FunctionEffect{"Invisible Vehicles (Only Wheels)", "WheelsOnlyPlease", "wheels_only_please", 0.0},
		FunctionEffect{"Peds Attack Each Other", "RoughNeighbourhood", "peds_attack_each_other", 0.0},
		FunctionEffect{"Have A Bounty On Your Head", "StopPickingOnMe", "have_a_bounty_on_your_head", 0.0},
		FunctionEffect{"Elvis Is Everywhere", "BlueSuedeShoes", "elvis_is_everywhere", 0.0},
		FunctionEffect{"Peds Attack You", "AttackOfTheVillagePeople", "peds_attack_you", 0.0},
		FunctionEffect{"Gang Members Everywhere", "OnlyHomiesAllowed", "gang_members_everywhere", 0.0},
		FunctionEffect{"Gangs Control The Streets", "BetterStayIndoors", "gangs_control_the_streets", 0.0},
		FunctionEffect{"Riot Mode", "StateOfEmergency", "riot_mode", 0.0},
		FunctionEffect{"Everyone Armed", "SurroundedByNutters", "everyone_armed", 0.0},
		FunctionEffect{"Aggressive Drivers", "AllDriversAreCriminals", "aggressive_drivers", 0.0},
		FunctionEffect{"Recruit Anyone (9mm)", "WannaBeInMyGang", "recruit_anyone_9mm", 0.0},
		FunctionEffect{"Recruit Anyone (AK-47)", "NoOneCanStopUs", "recruit_anyone_ak47", 0.0},
		FunctionEffect{"Recruit Anyone (Rockets)", "RocketMayhem", "recruit_anyone_rockets", 0.0},
		FunctionEffect{"Ghost Town", "GhostTown", "ghost_town", 0.0},
		FunctionEffect{"Beach Party", "LifesABeach", "beach_theme", 0.0},
		FunctionEffect{"Ninja Theme", "NinjaTown", "ninja_theme", 0.0},
		FunctionEffect{"Kinky Theme", "LoveConquersAll", "kinky_theme", 0.0},
		FunctionEffect{"Funhouse Theme", "CrazyTown", "funhouse_theme", 0.0},
		FunctionEffect{"Country Traffic", "HicksVille", "country_traffic", 0.0},
		FunctionEffect{"Weapon Aiming While Driving", "IWannaDriveBy", "weapon_aiming_while_driving", 0.0},
		FunctionEffect{"Huge Bunny Hop", "CJPhoneHome", "huge_bunny_hop", 0.0},
		FunctionEffect{"Mega Jump", "Kangaroo", "mega_jump", 0.0},
		FunctionEffect{"Infinite Oxygen", "ManFromAtlantis", "infinite_oxygen", 0.0},
		FunctionEffect{"Mega Punch", "StingLikeABee", "mega_punch", 0.0},
		FunctionEffect{"Fat Player", "WhoAteAllThePies", "fat_player", 0.0},
		FunctionEffect{"Max Muscle", "BuffMeUp", "muscle_player", 0.0},
		FunctionEffect{"Skinny Player", "LeanAndMean", "skinny_player", 0.0},
		FunctionEffect{"Max Stamina", "ICanGoAllNight", "max_stamina", 0.0},
		FunctionEffect{"No Stamina", "ImAllOutOfBreath", "no_stamina", 0.0},
		FunctionEffect{"Hitman Level For All Weapons", "ProfessionalKiller", "hitman_level_for_all_weapons", 0.0},
		FunctionEffect{"Beginner Level For All Weapons", "BabysFirstGun", "beginner_level_for_all_weapons", 0.0},
		FunctionEffect{"Max Driving Skills", "NaturalTalent", "max_driving_skills", 0.0},
		FunctionEffect{"No Driving Skills", "BackToDrivingSchool", "no_driving_skills", 0.0},
		FunctionEffect{"Never Get Hungry", "IAmNeverHungry", "never_get_hungry", 0.0},
		FunctionEffect{"Lock Respect At Max", "WorshipMe", "lock_respect_at_max", 0.0},
		FunctionEffect{"Lock Sex Appeal At Max", "HelloLadies", "lock_sex_appeal_at_max", 0.0},
		FunctionEffect{"Clear Active Effects", "ClearActiveEffects", "clear_active_effects", 0.0},
		FunctionEffect{"Remove All Weapons", "NoWeaponsAllowed", "remove_all_weapons", 0.0},
		FunctionEffect{"Get Busted", "GoToJail", "get_busted", 0.0},
		FunctionEffect{"Get Wasted", "Hospitality", "get_wasted", 0.0},
		FunctionEffect{"Set Everyone On Fire", "HotPotato", "set_everyone_on_fire", 0.0},
		FunctionEffect{"Kick Player Out Of Vehicle", "ThisAintYourCar", "kick_player_out_of_vehicle", 0.0},
		FunctionEffect{"Lock Player Inside Vehicle", "ThereIsNoEscape", "lock_player_inside_vehicle", 0.0},
		FunctionEffect{"Set Current Vehicle On Fire", "WayTooHot", "set_current_vehicle_on_fire", 0.0},
		FunctionEffect{"Pop Tires Of All Vehicles", "TiresBeGone", "pop_tires_of_all_vehicles", 0.0},
		FunctionEffect{"Send Vehicles To Space", "StairwayToHeaven", "send_vehicles_to_space", 0.0},
		FunctionEffect{"Turn Vehicles Around", "TurnAround", "turn_vehicles_around", 0.0},
		FunctionEffect{"To The Left, To The Right", "ToTheLeftToTheRight", "to_the_left_to_the_right", 0.0},
		FunctionEffect{"Timelapse Mode", "DiscoInTheSky", "timelapse", 0.0},
		FunctionEffect{"Where Is Everybody?", "ImHearingVoices", "where_is_everybody", 0.0},
		FunctionEffect{"Everybody Bleed Now!", "EverybodyBleedNow", "everybody_bleed_now", 0.5},
		FunctionEffect{"To Drive Or Not To Drive", "ToDriveOrNotToDrive", "to_drive_or_not_to_drive", 0.0},
		FunctionEffect{"One Hit K.O.", "ILikeToLiveDangerously", "one_hit_ko", 0.0},
		FunctionEffect{"Experience The Lag", "PacketLoss", "experience_the_lag", 0.0},
		FunctionEffect{"Ghost Rider", "GhostRider", "ghost_rider", 0.0},
		FunctionEffect{"Disable HUD", "FullyImmersed", "disable_hud", 0.0},
		FunctionEffect{"Disable Blips / Markers / Pickups", "INeedSomeInstructions", "disable_blips_markers_pickups", 0.0},
		FunctionEffect{"Disable All Weapon Damage", "TruePacifist", "disable_all_weapon_damage", 0.0},
		FunctionEffect{"Let's Take A Break", "LetsTakeABreak", "lets_take_a_break", 0.0},
		FunctionEffect{"Pride Cars", "AllColorsAreBeautiful", "pride_cars", 0.0},
		FunctionEffect{"High Suspension Damping", "VeryDampNoBounce", "high_suspension_damping", 0.0},
		FunctionEffect{"Zero Suspension Damping", "LowrideAllNight", "zero_suspension_damping", 0.0},         // Cars have almost zero suspension dampin},
		FunctionEffect{"Long Live The Rich!", "LongLiveTheRich", "long_live_the_rich", 0.0},                  // Money = Health, shoot people to gain money, take damage to lose i},
		FunctionEffect{"Inverted Controls", "InvertedControls", "inverted_controls", 0.0},                    // Inverts some control},
		FunctionEffect{"Disable One Movement Key", "DisableOneMovementKey", "disable_one_movement_key", 0.0}, // Disable one movement ke},
		FunctionEffect{"Fail Current Mission", "MissionFailed", "fail_current_mission", 0.0},
		FunctionEffect{"Night Vision", "NightVision", "night_vision", 0.0},
		FunctionEffect{"Thermal Vision", "ThermalVision", "thermal_vision", 0.0},
		FunctionEffect{"Pass Current Mission", "IllTakeAFreePass", "pass_current_mission", 0.0},
		FunctionEffect{"Infinite Health (Everyone)", "NoOneCanHurtAnyone", "infinite_health_everyone", 0.0},
		FunctionEffect{"Invisible Vehicles", "InvisibleVehicles", "invisible_vehicles", 0.0},
		FunctionEffect{"Powerpoint Presentation", "PowerpointPresentation", "framerate_15", 0.0},
		FunctionEffect{"Smooth Criminal", "SmoothCriminal", "framerate_60", 0.0},
		FunctionEffect{"Reload Autosave", "HereWeGoAgain", "reload_autosave", 0.0},
		FunctionEffect{"Quarter Gravity", "GroundControlToMajorTom", "quarter_gravity", 0.0},
		FunctionEffect{"Half Gravity", "ImFeelingLightheaded", "half_gravity", 0.0},
		FunctionEffect{"Double Gravity", "KilogramOfFeathers", "double_gravity", 0.0},
		FunctionEffect{"Quadruple Gravity", "KilogramOfSteel", "quadruple_gravity", 0.0},
		FunctionEffect{"Inverted Gravity", "BeamMeUpScotty", "inverted_gravity", 0.0},
		FunctionEffect{"Zero Gravity", "ImInSpaaaaace", "zero_gravity", 0.0},
		FunctionEffect{"Insane Gravity", "StraightToHell", "insane_gravity", 0.0},
		FunctionEffect{"Tunnel Vision", "TunnelVision", "tunnel_vision", 1.0 / 3.0},
		FunctionEffect{"High Pitched Audio", "CJAndTheChipmunks", "high_pitched_audio", 0.0},
		FunctionEffect{"Pitch Shifter", "VocalRange", "pitch_shifter", 0.0},
		FunctionEffect{"Pass Current Mission", "IWontTakeAFreePass", "fake_pass_current_mission", 0.0},
		FunctionEffect{"DVD Screensaver", "ItsGonnaHitTheCorner", "dvd_screensaver", 1.0 / 3.0},
		FunctionEffect{"Honk Boost", "GottaHonkFast", "honk_boost", 0.0},
		FunctionEffect{"Oh Hey, Tanks!", "OhHeyTanks", "oh_hey_tanks", 0.0},
		FunctionEffect{"Always Wanted", "ICanSeeStars", "always_wanted", 0.0},
		FunctionEffect{"Cinematic Vehicle Camera", "MachinimaMode", "cinematic_vehicle_camera", 0.0},
		FunctionEffect{"Your Car Floats Away When Hit", "ImTheBubbleCar", "your_car_floats_away_when_hit", 0.0},
		FunctionEffect{"Ring Ring !!", "RingRing", "ring_ring", 0.0},
		FunctionEffect{"Peds Explode When Run Over", "ExplosivePeds", "peds_explode_when_run_over", 0.0},
		FunctionEffect{"HONK!!!", "HONKHONK", "honk_vehicle", 5},
		FunctionEffect{"Quake FOV", "QuakeFOV", "quake_fov", 0.0},
		FunctionEffect{"Beyblade", "LetItRip", "beyblade", 0.0},
		FunctionEffect{"Weapon Roulette", "WeaponRoulette", "weapon_roulette", 0.0},
		FunctionEffect{"No Need To Hurry", "NoNeedToHurry", "no_need_to_hurry", 0.0},
		FunctionEffect{"Drunk Player", "DrunkPlayer", "drunk_player", 0.0},
		FunctionEffect{"Force Vehicle Mouse Steering", "ForceVehicleMouseSteering", "force_vehicle_mouse_steering", 0.0},
		FunctionEffect{"Upside-Down Screen", "WhatsUpIsDown", "upside_down_screen", 1.0 / 3.0},
		FunctionEffect{"Mirrored Screen", "WhatsLeftIsRight", "mirrored_screen", 1.0 / 3.0},
		FunctionEffect{"Mirrored World", "LetsTalkAboutParallelUniverses", "mirrored_world", 0.0},
		FunctionEffect{"Big Heads", "BigHeadsMode", "big_heads", 0.0},
		FunctionEffect{"Pedal To The Metal", "PedalToTheMetal", "pedal_to_the_metal", 0.0},
		FunctionEffect{"Unflippable Vehicles", "ThereGoesMyBurrito", "unflippable_vehicles", 0.0},
		FunctionEffect{"Freefall!", "WhereWeDroppingBoys", "freefall", 0.0},
		FunctionEffect{"Carmageddon", "Carmageddon", "carmageddon", 0.0},
		FunctionEffect{"Usain Bolt", "FastestManAlive", "usain_bolt", 0.0},
		FunctionEffect{"Roll Credits", "WaitItsOver", "roll_credits", 0.0},
		FunctionEffect{"Instantly Hungry", "IllHave2Number9s", "instantly_hungry", 0.0},
		FunctionEffect{"Vehicle Bumper Camera", "FrontRowSeat", "vehicle_bumper_camera", 0.0},
		FunctionEffect{"Random Teleport", "LetsGoSightseeing", "random_teleport", 0.0},

		WeatherEffect{"Sunny Weather", "PleasantlyWarm", 1},
		WeatherEffect{"Very Sunny Weather", "TooDamnHot", 0},
		WeatherEffect{"Overcast Weather", "DullDullDay", 4},
		WeatherEffect{"Rainy Weather", "StayInAndWatchTV", 16},
		WeatherEffect{"Foggy Weather", "CantSeeWhereImGoing", 9},
		WeatherEffect{"Thunderstorm", "ScottishSummer", 16},
		WeatherEffect{"Sandstorm", "SandInMyEars", 19},

		SpawnVehicleEffect{"TimeToKickAss", 432},
		SpawnVehicleEffect{"OldSpeedDemon", 504},
		SpawnVehicleEffect{"DoughnutHandicap", 489},
		SpawnVehicleEffect{"NotForPublicRoads", 502},
		SpawnVehicleEffect{"JustTryAndStopMe", 503},
		SpawnVehicleEffect{"WheresTheFuneral", 442},
		SpawnVehicleEffect{"CelebrityStatus", 409},
		SpawnVehicleEffect{"TrueGrime", 408},
		SpawnVehicleEffect{"18Holes", 457},
		SpawnVehicleEffect{"JumpJet", 520},
		SpawnVehicleEffect{"IWantToHover", 539},
		SpawnVehicleEffect{"OhDude", 425},
		SpawnVehicleEffect{"FourWheelFun", 471},
		SpawnVehicleEffect{"ItsAllBull", 486},
		SpawnVehicleEffect{"FlyingToStunt", 513},
		SpawnVehicleEffect{"MonsterMash", 556},
		SpawnVehicleEffect{"SurpriseDriver", -1},

		TeleportationEffect{"Teleport Home", "BringMeHome", GrooveStreet},
		TeleportationEffect{"Teleport To A Tower", "BringMeToATower", LSTower},
		TeleportationEffect{"Teleport To A Pier", "BringMeToAPier", LSPier},
		TeleportationEffect{"Teleport To The LS Airport", "BringMeToTheLSAirport", LSAirport},
		TeleportationEffect{"Teleport To The Docks", "BringMeToTheDocks", LSDocks},
		TeleportationEffect{"Teleport To A Mountain", "BringMeToAMountain", MountChiliad},
		TeleportationEffect{"Teleport To The SF Airport", "BringMeToTheSFAirport", SFAirport},
		TeleportationEffect{"Teleport To A Bridge", "BringMeToABridge", SFBridge},
		TeleportationEffect{"Teleport To A Secret Place", "BringMeToASecretPlace", Area52},
		TeleportationEffect{"Teleport To A Quarry", "BringMeToAQuarry", LVQuarry},
		TeleportationEffect{"Teleport To The LV Airport", "BringMeToTheLVAirport", LVAirport},
		TeleportationEffect{"Teleport To Big Ear", "BringMeToBigEar", LVSatellite},
	}
}
