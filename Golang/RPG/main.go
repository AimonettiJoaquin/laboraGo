package main

import (
	"fmt"
	"math"
	"math/rand/v2"
	"strings"
)

func main() {
	jugador := createCaracter()
	initialEquipment(jugador)
	for !randomEncounter(jugador) {
		wait()
	}

}

type Weapon struct {
	name        string
	cost        float64
	weight      float64
	description string
	nroDice     int
	damage      int
	oneHand     bool
	equiped     bool
	finesse     bool
	properties  WeaponsProperty
}
type WeaponsProperty struct {
	Ammunition bool
	Finesse    bool
	Heavy      bool
	Light      bool
	Loading    bool
	Range      bool
	Reach      bool
	Thrown     bool
	TwoHanded  bool
	Versatile  bool
}
type WeaponsProperties struct {
	weaponsProperties []WeaponsProperty
}
type Weapons struct {
	weapons []Weapon
}
type Armor struct {
	name        string
	cost        float64
	weight      float64
	description string
	armorClass  int
	shield      bool
	equiped     bool
}
type Armors struct {
	armors []Armor
}
type ArmorType string

const (
	lightArmor  ArmorType = "Light Armor"
	mediumArmor ArmorType = "Medium Armor"
	heavyArmor  ArmorType = "Heavy Armor"
)

type Creature struct {
	name             string
	size             Size
	race             Race
	hitDice          int
	currentHitDice   int
	Level            int
	ArmorClass       int
	MaxHitPoints     int
	CurrentHitPoints int
	Strength         int
	StrMod           int
	Dexterity        int
	DexMod           int
	Consitution      int
	ConMod           int
	Intelligence     int
	IntMod           int
	Wisdom           int
	WisMod           int
	Charisma         int
	ChaMod           int
	ProficiencyBonus int
	weapons          []Weapon
	armors           []Armor
	xp               int
	iniciative       int
}
type Creatures struct {
	creatures []Creature
}

type Size string

const (
	tiny   Size = "tiny"
	small  Size = "small"
	medium Size = "medium"
)

type Race string

const (
	humanoide   Race = "Humanoide"
	monstrosity Race = "Monstrosity"
)

// funciones de juego
func createCaracter() *Creature {
	fmt.Println("******************************************************************************************************************************")
	fmt.Println("Creación de personaje")
	fmt.Print("Ingrese el nombre de tu personaje: ")
	var name string
	fmt.Scan(&name)
	c := &Creatures{}

	c.createCreature(name, medium, humanoide, 10, 1, 12, 16, 10, 14, 8, 12, 12)
	jugador := c.findCreature(name)
	fmt.Println("*******************************************************")
	fmt.Println("En un reino de Labora, la princesa Juan ha sido raptada")
	fmt.Println("por una horda de goblins y llevada a una caverna")
	fmt.Println("oculta en el corazón del Bosque Oscuro. El rey,")
	fmt.Println("desesperado, promete una gran recompensa a quien la")
	fmt.Println("rescate. Un valiente guerrero, armado con su espada y")
	fmt.Println("su coraje, se adentra en la caverna de los goblins.")
	fmt.Println("Enfrentando trampas y criaturas peligrosas, se adentra")
	fmt.Println("en la oscuridad con un solo objetivo: rescatar a la")
	fmt.Println("princesa y devolver la esperanza al reino. ¿Logrará")
	fmt.Println("nuestro héroe superar los desafíos y salvar a la")
	fmt.Println("princesa? Solo su valentía y tu habilidad decidirán")
	fmt.Println("el final de esta historia.")
	fmt.Println("*******************************************************")
	wait()
	fmt.Println("Antes de adentrarte a la caverna donde esta la princesa, debes elegir un arma")
	fmt.Println("1. Espada de dos manos")
	fmt.Println("2. Espada de una mano + Escudo")
	return jugador
}

func initialEquipment(jugador *Creature) {
	w := &Weapons{}
	c := &Creatures{}
	a := &Armors{}
	var equipoInicial string
	for {
		fmt.Print("Ingrese 'Mandoble' o 'Espada y escudo': ")
		fmt.Scanf("%s", &equipoInicial)
		fmt.Println("******************************************")
		if equipoInicial == "Mandoble" {
			fmt.Println("El guerrero se pone su cota de malla y toma su Zweihander")
			wait()
			property := createProperty(false, false, false, false, false, false, false, false, false)
			w.createWeapon("Zweihänder", 50, 6, "La Zweihänder es una gran espada a dos manos, usada por infantería medieval para romper formaciones.", 2, 6, false, false, property)
			Zweihänder := w.findWeapon("Zweihänder")
			c.equipWeapon(jugador, Zweihänder)
			a.createArmor("Cota de Malla", 75, 55, "La cota de malla es una armadura de anillos de metal entrelazados, ofreciendo flexibilidad y protección.", 16, false)
			armor := a.findArmor("Cota de Malla")
			c.equipArmor(jugador, armor)
			c.printCreature(jugador)

			//c.printEqupedWeapon(jugador)
			break
		} else if equipoInicial == "Espada" {
			fmt.Println("El guerrero se pone su cota de malla, toma su escudo y su espada larga.")
			wait()
			property := createProperty(false, false, false, false, false, false, false, false, false)
			w.createWeapon("Espada larga", 15, 3, "La espada larga es una espada medieval de doble filo, usada con una o ambas manos.", 1, 8, true, false, property)
			espadaLarga := w.findWeapon("Espada larga")
			c.equipWeapon(jugador, espadaLarga)
			a.createArmor("Cota de Malla", 75, 55, "La cota de malla es una armadura de anillos de metal entrelazados, ofreciendo flexibilidad y protección.", 16, false)
			armor := a.findArmor("Cota de Malla")
			c.equipArmor(jugador, armor)
			a.createArmor("Escudo", 10, 6, "Escudo medieval: arma defensiva de madera, decorado con emblemas heráldicos, utilizado en combates y torneos para protección contra ataques enemigos.", 2, true)
			shield := a.findArmor("Escudo")
			c.equipShield(jugador, shield)

			c.printCreature(jugador)
			break
		} else {
			fmt.Println("Opción no valida")
		}
	}
}
func wait() {
	wait := ""
	fmt.Scanf("%s", &wait)
}

func randomEncounter(jugador *Creature) bool {
	encuentro := rollD(10)
	encuentro = 1
	if encuentro >= 3 {
		goblinEncounter(jugador)
		return false
	} else if encuentro == 1 {

		chest := chestEncounter(jugador)
		switch chest {
		case 1:
			mimicEncounter(jugador)
		case 2:
			if jugador.Strength >= 20 {
				fmt.Println("Tu fuerza no puede subir mas.")
			} else {
				jugador.Strength += 2
				fmt.Println("Tu fuerza ha aumentado en 2 puntos.")
				fmt.Println("Fuerza actual:", jugador.Strength)
			}
			checkLvlUp(jugador, 50)

		case 3:
			jugador.CurrentHitPoints = jugador.MaxHitPoints
			fmt.Println("Te sientes recuperado.")
			fmt.Println("HP:", jugador.CurrentHitPoints, "/", jugador.CurrentHitPoints)
			checkLvlUp(jugador, 20)
		case 4:
			checkLvlUp(jugador, 10)
		}
		return false
	} else if encuentro == 2 {
		fmt.Println("***************************")
		fmt.Println("No has encontrado nada")
		wait()
		return false
	}

	return false
}

func chestEncounter(j *Creature) int {
	fmt.Println("****************************")
	fmt.Println("Has encontrado un cofre")
	wait()
	chestArt()
	chest := rollD(100)
	chest = 90
	if chest >= 95 {
		fmt.Println("NO ERA UN COFRE, ERA UN MIMIC!!! .")
		wait()
		mimicArt()
		wait()
		return 1
	} else if chest >= 90 {
		fmt.Println("El jugador encuentra un cofre en la oscuridad de la caverna.")
		fmt.Println("Con cuidado, abre el cofre y dentro descubre una poción mágica.")
		fmt.Println("Es una poción de fuerza, brillando con una luz mística.")
		fmt.Println("La toma, sintiendo el poder fluir a través de su cuerpo.")
		wait()
		potionArt()
		wait()
		return 2
	} else if chest >= 85 {
		fmt.Println("El jugador encuentra un cofre en la oscuridad de la caverna.")
		fmt.Println("Con cuidado, abre el cofre y dentro descubre una poción mágica.")
		fmt.Println("Es una poción de vida, brillando con una luz mística.")
		fmt.Println("La toma, sintiendo como recupera su vigor.")
		wait()
		potionArt()
		wait()

		return 3
	} else {
		fmt.Println("El cofre esta vacio.")
		fmt.Println("****************************")
		return 4
	}

}

func potionArt() {
	fmt.Println("       _____")
	fmt.Println("      `.___,'")
	fmt.Println("       (___)")
	fmt.Println("       <   >")
	fmt.Println("        ) (")
	fmt.Println("       /`-.\\")
	fmt.Println("      /     \\")
	fmt.Println("     / _    _\\")
	fmt.Println("    :,' `-.' `:")
	fmt.Println("    |         |")
	fmt.Println("    :         ;")
	fmt.Println("     \\       /")
	fmt.Println("      `.___.'")
}

func increaseStat(j *Creature) {
	fmt.Println("****************************************************")
	fmt.Println("Puedes subir una estadistica +2 o 2 estadisticas +1")
	fmt.Println("Fuerza", j.Strength)
	fmt.Println("Destreza", j.Dexterity)
	fmt.Println("Constitucion", j.Consitution)
	fmt.Println("Inteligencia", j.Intelligence)
	fmt.Println("Sabiduria", j.Wisdom)
	fmt.Println("Carisma", j.Charisma)
	fmt.Println("****************************************************")
	fmt.Println("Ejemplo: Fuerza y Destreza")
	fmt.Println("O: Fuerza +2")
	var stat string
	fmt.Scanf("%s", &stat)
	switch {
	case strings.Contains(stat, "Fuerza") && strings.Contains(stat, "Destreza"):
		j.Strength++
		j.Dexterity++
	case strings.Contains(stat, "Fuerza") && (strings.Contains(stat, "Constitucion") || strings.Contains(stat, "Constitución")):
		j.Strength++
		j.Consitution++
	case strings.Contains(stat, "Fuerza") && (strings.Contains(stat, "Inteligencia")):
		j.Strength++
		j.Intelligence++
	case strings.Contains(stat, "Fuerza") && (strings.Contains(stat, "Sabiduria") || strings.Contains(stat, "Sabiduría")):
		j.Strength++
		j.Wisdom++
	case strings.Contains(stat, "Fuerza") && strings.Contains(stat, "Carisma"):
		j.Strength++
		j.Charisma++
		//Destreza
	case strings.Contains(stat, "Destreza") && (strings.Contains(stat, "Constitucion") || strings.Contains(stat, "Constitución")):
		j.Dexterity++
		j.Consitution++
	case strings.Contains(stat, "Destreza") && strings.Contains(stat, "Inteligencia"):
		j.Dexterity++
		j.Intelligence++
	case strings.Contains(stat, "Destreza") && (strings.Contains(stat, "Sabiduria") || strings.Contains(stat, "Sabiduría")):
		j.Dexterity++
		j.Wisdom++
	case strings.Contains(stat, "Destreza") && strings.Contains(stat, "Carisma"):
		j.Dexterity++
		j.Charisma++
		//Constitucion
	case strings.Contains(stat, "Constitucion") && strings.Contains(stat, "Inteligencia"):
		j.Consitution++
		j.Intelligence++
	case strings.Contains(stat, "Constitucion") && (strings.Contains(stat, "Sabiduria") || strings.Contains(stat, "Sabiduría")):
		j.Consitution++
		j.Wisdom++
	case strings.Contains(stat, "Constitucion") && strings.Contains(stat, "Carisma"):
		j.Consitution++
		j.Charisma++
		//Inteligencia
	case strings.Contains(stat, "Inteligencia") && (strings.Contains(stat, "Sabiduria") || strings.Contains(stat, "Sabiduría")):
		j.Intelligence++
		j.Wisdom++
	case strings.Contains(stat, "Inteligencia") && strings.Contains(stat, "Carisma"):
		j.Intelligence++
		j.Charisma++
		//Sabiduria
	case strings.Contains(stat, "Sabiduria") && strings.Contains(stat, "Carisma"):
		j.Wisdom++
		j.Charisma++
	case strings.Contains(stat, "Fuerza"):
		j.Strength++
	case strings.Contains(stat, "Destreza"):
		j.Dexterity++
	case strings.Contains(stat, "Constitucion"):
		j.Consitution++
	case strings.Contains(stat, "Inteligencia"):
		j.Intelligence++
	case strings.Contains(stat, "Sabiduria"):
		j.Wisdom++
	case strings.Contains(stat, "Carisma"):
		j.Charisma++
	}
}

func lvlUp(j *Creature) {
	j.Level++
	fmt.Println("Felicidades, subiste al nivel", j.Level)
	wait()
	hp := rollD(j.hitDice) + j.ConMod
	j.currentHitDice += hp
	j.MaxHitPoints += hp
	if j.Level == 4 {
		increaseStat(j)
	}
	if j.Level == 5 {
		j.ProficiencyBonus = 3
	}
}

func checkLvlUp(j *Creature, xp int) bool {
	j.xp += xp
	fmt.Println("Has ganado", xp, "puntos de experiencia.")

	switch j.Level {
	case 1:
		fmt.Println("Experiencia:", j.xp, "/300")
		if j.xp >= 300 {
			return true
		}
	case 2:
		if j.xp >= 900 {
			fmt.Println("Experiencia:", j.xp, "/900")
			return true
		}
	case 3:
		if j.xp >= 2700 {
			fmt.Println("Experiencia:", j.xp, "2700")
			return true
		}
	case 4:
		if j.xp >= 6500 {
			fmt.Println("Experiencia:", j.xp, "/6500")
			return true
		}
	case 5:
		if j.xp >= 14000 {
			fmt.Println("Experiencia:", j.xp, "/14000")
			return true
		}
	default:
		return false
	}
	return false
}
func mimicEncounter(jugador *Creature) bool {
	w := &Weapons{}
	c := &Creatures{}
	a := &Armors{}
	c.createCreature("Mimic", medium, monstrosity, 8, 9, 58, 17, 12, 15, 5, 13, 8)
	mimic := c.findCreature("Mimic")
	property := createProperty(false, false, false, false, false, false, false, false, false)
	w.createWeapon("Bite", 0, 0, "Bite.", 2, 8, true, false, property)
	bite := w.findWeapon("Bite")
	c.equipWeapon(mimic, bite)
	a.createArmor("Armadura natural", 0, 0, "El mimic posee una piel dura y una gran resistencia, pero su inmoviliad lo vuelve un objetivo facil.", 12, false)
	armor := a.findArmor("Armadura natural")
	c.equipArmor(mimic, armor)
	fmt.Println("El mimic te sorprendio, va primero")
	wait()
	for {
		if c.attack(mimic, jugador) {
			gameOver()
			return true
		} else if c.attack(jugador, mimic) {
			goblinDeadArt()
			checkLvlUp(jugador, 450)

			return false
		}
	}
}
func goblinEncounter(jugador *Creature) bool {
	w := &Weapons{}
	c := &Creatures{}
	a := &Armors{}
	fmt.Println("****************************************************")
	fmt.Println("El protagonista se adentra más en la caverna,")
	fmt.Println("el eco de sus pasos resonando en las paredes.")
	fmt.Println("La luz de la antorcha revela un pasaje oscuro.")
	fmt.Println("De repente, escucha un ruido adelante.")
	fmt.Println("Un goblin aparece, listo para atacarlo.")
	fmt.Println("****************************************************")
	wait()
	goblinArt()
	wait()
	c.createCreature("Goblin", small, humanoide, 6, 1, 7, 8, 14, 10, 10, 8, 8)
	goblin := c.findCreature("Goblin")
	property := createProperty(false, true, false, false, false, false, false, false, false)
	w.createWeapon("Scimitarra", 15, 3, "Scimitarra.", 1, 6, true, false, property)
	scimitarra := w.findWeapon("Scimitarra")
	c.equipWeapon(goblin, scimitarra)
	a.createArmor("Armadura de cuero", 10, 10, "Armadura de cuero.", 13, false)
	armor := a.findArmor("Armadura de cuero")
	c.equipArmor(goblin, armor)
	a.createArmor("Escudo", 10, 6, "Escudo medieval: arma defensiva de madera, decorado con emblemas heráldicos, utilizado en combates y torneos para protección contra ataques enemigos.", 2, true)
	shield := a.findArmor("Escudo")
	c.equipShield(goblin, shield)
	var jugadorIni int
	var goblinIni int
	for {
		goblinIni = rollIniciative(goblin)
		jugadorIni = rollIniciative(jugador)
		if goblinIni != jugadorIni {
			fmt.Println("******************************************************************************************************************************")
			fmt.Println("Iniciativa de", goblin.name, "es", goblinIni)
			fmt.Println("******************************************************************************************************************************")
			fmt.Println("Iniciativa de", jugador.name, "es", jugadorIni)
			break
		}
	}

	if jugadorIni > goblinIni {
		fmt.Println("******************************************************************************************************************************")
		fmt.Println(jugador.name, "empieza")
		wait()
		for {
			if c.attack(jugador, goblin) {
				goblinDeadArt()
				checkLvlUp(jugador, 50)
				return false

			} else if c.attack(goblin, jugador) {
				gameOver()
				return true

			}
		}
	} else {
		for {
			if c.attack(goblin, jugador) {

				gameOver()
				return true

			} else if c.attack(jugador, goblin) {
				goblinDeadArt()
				checkLvlUp(jugador, 50)
				return false

			}
		}
	}
}

// Criaturas
func (c *Creatures) createCreature(name string, size Size, race Race, hitDice int, level float64, HP int, STR int, DEX int, CON int, INT int, WIS int, CHA int) {
	pb := math.Ceil((level / 4) + 1)
	strMod := (STR - 10) / 2
	dexMod := (DEX - 10) / 2
	conMod := (CON - 10) / 2
	intMod := (INT - 10) / 2
	wisMod := (WIS - 10) / 2
	chaMod := (CHA - 10) / 2
	calcAC := 10 + ((DEX - 10) / 2)
	creature := Creature{
		name:             name,
		size:             size,
		race:             race,
		Level:            int(level),
		ArmorClass:       calcAC,
		MaxHitPoints:     HP,
		CurrentHitPoints: HP,
		Strength:         STR,
		StrMod:           strMod,
		Dexterity:        DEX,
		DexMod:           dexMod,
		Consitution:      CON,
		ConMod:           conMod,
		Intelligence:     INT,
		IntMod:           intMod,
		Wisdom:           WIS,
		WisMod:           wisMod,
		Charisma:         CHA,
		ChaMod:           chaMod,
		xp:               0,
		ProficiencyBonus: int(pb),
		iniciative:       dexMod,
		hitDice:          hitDice,
		currentHitDice:   int(level),
	}
	c.creatures = append(c.creatures, creature)
}
func (c *Creatures) findCreature(name string) *Creature {
	for _, creature := range c.creatures {
		if creature.name == name {
			return &creature
		}
	}
	return nil
}

func rollIniciative(c *Creature) int {
	iniciative := rollD(20) + c.iniciative
	return iniciative
}

func rollD(dice int) int {

	return rand.IntN(dice) + 1
}

func chestArt() {
	fmt.Println("       _________________________")
	fmt.Println("      /                        /\\")
	fmt.Println("     /________________________/  \\")
	fmt.Println("    |                        |    |")
	fmt.Println("    |  _____      _____      |    |")
	fmt.Println("    |  \\    \\    /    /      |    |")
	fmt.Println("    |   \\____\\  /____/       |    |")
	fmt.Println("    |   /____/  \\____\\       |    |")
	fmt.Println("    |  /\\    \\  /    /\\      |    |")
	fmt.Println("    | /  \\____\\/____/  \\     |    |")
	fmt.Println("    |/____________________\\  |    |")
	fmt.Println("    [______________________\\ |    |")
	fmt.Println("    |                        |    |")
	fmt.Println("    |________________________|____|")
	fmt.Println("    /                        /    /")
	fmt.Println("   /________________________/____/")
}
func mimicArt() {
	fmt.Println("       _________________________")
	fmt.Println("      /                        /\\")
	fmt.Println("     /________________________/  \\")
	fmt.Println("    |                        |    |")
	fmt.Println("    |  _____      _____      |    |")
	fmt.Println("    |  \\ O /\\    /\\ O /      |    |")
	fmt.Println("    |   \\_/  \\  /  \\_/       |    |")
	fmt.Println("    |   / \\   \\/   / \\       |    |")
	fmt.Println("    |  /___\\  /\\  /___\\      |    |")
	fmt.Println("    | /     \\/  \\/     \\     |    |")
	fmt.Println("    |/____.___/\\___.____\\    |    |")
	fmt.Println("    [______________________\\ |    |")
	fmt.Println("    |                        |    |")
	fmt.Println("    |   /\\ /\\ /\\ /\\ /\\ /\\   |____|")
	fmt.Println("    |  /  V  V  V  V  V  \\  /    /")
	fmt.Println("   /________________________/____/")
}
func goblinDeadArt() {
	fmt.Println("        ,      ,")
	fmt.Println("       /(.-\"\"-.)\\")
	fmt.Println("   |\\  \\/      \\/  /|")
	fmt.Println("   | \\ / =.  .= \\ / |")
	fmt.Println("   \\( \\   x\\/x   / )/")
	fmt.Println("    \\_, '-/  \\-' ,_/")
	fmt.Println("      /   \\__/   \\")
	fmt.Println("      \\ \\__/\\__/ /")
	fmt.Println("    ___\\ \\|--|/ /___")
	fmt.Println("  /`    \\      /    `\\")
	fmt.Println(" /       '----'       \\")
}

func goblinArt() {
	fmt.Println("        ,      ,")
	fmt.Println("       /(.-\"\"-.)\\")
	fmt.Println("   |\\  \\/      \\/  /|")
	fmt.Println("   | \\ / =.  .= \\ / |")
	fmt.Println("   \\( \\   o\\/o   / )/")
	fmt.Println("    \\_, '-/  \\-' ,_/")
	fmt.Println("      /   \\__/   \\")
	fmt.Println("      \\ \\__/\\__/ /")
	fmt.Println("    ___\\ \\|--|/ /___")
	fmt.Println("  /`    \\      /    `\\")
	fmt.Println(" /       '----'       \\")
	fmt.Println("****************************************************")
}
func gameOver() {
	fmt.Println("       _____")
	fmt.Println("      //    \\\\")
	fmt.Println("     || X  X ||")
	fmt.Println("     ||  __  ||")
	fmt.Println("      \\\\____//")

	fmt.Println("  ____                         ___                 ")
	fmt.Println(" / ___| __ _ _ __ ___   ___   / _ \\__   _____ _ __ ")
	fmt.Println("| |  _ / _` | '_ ` _ \\ / _ \\ | | | \\ \\ / / _ \\ '__|")
	fmt.Println("| |_| | (_| | | | | | |  __/ | |_| |\\ V /  __/ |   ")
	fmt.Println(" \\____|\\__,_|_| |_| |_|\\___|  \\___/  \\_/ \\___|_|   ")
}
func (c *Creatures) attack(attacker *Creature, defender *Creature) bool {

	fmt.Println("**************************")
	attackerWeapon := c.findEquipedWeapon(attacker)
	fmt.Println(attacker.name, "Usa su", attackerWeapon.name)
	d20 := rollD(20)
	var attackerMod int
	if attackerWeapon.properties.Finesse && attacker.DexMod > attacker.StrMod {
		attackerMod = attacker.DexMod
	} else {
		attackerMod = attacker.StrMod
	}
	attackerToHit := d20 + attackerMod + attacker.ProficiencyBonus
	fmt.Println("La tirada de ataque:", d20, "Ataque", "+ Modificador de ataque", attackerMod, "+ Proficiencia", attacker.ProficiencyBonus, "=", (d20 + attackerMod + attacker.ProficiencyBonus))
	fmt.Println("La armadura del", defender.name, " es: ", defender.ArmorClass)
	wait()
	if attackerToHit >= defender.ArmorClass {

		weaponDice := attackerWeapon.nroDice
		weaponDice2 := attackerWeapon.nroDice
		nroDice := attackerWeapon.nroDice
		attackerWDamage := attackerWeapon.damage
		var damage int
		for weaponDice > 0 {
			damage += rollD(attackerWDamage)
			weaponDice--
		}
		damageNoMod := damage
		damage += attackerMod
		if damage == 0 {
			damage++
		}
		fmt.Println(defender.name, "tiene:", defender.CurrentHitPoints, "Puntos de vida")
		defender.CurrentHitPoints = defender.CurrentHitPoints - damage
		if defender.CurrentHitPoints <= 0 {
			fmt.Println("Daño:", nroDice, "d", weaponDice2, damageNoMod, "+", attackerMod, "=", damage)
			fmt.Println(defender.name, " ha muerto")
			fmt.Println("**************************")
			wait()
			return true
		} else {
			fmt.Println("El daño fue de: ", damage)
			fmt.Println("La vida actual de ", defender.name, "es: ", defender.CurrentHitPoints)
			fmt.Println("**************************")
			wait()
			return false
		}
	} else {
		fmt.Println("El ataque fallo")
		wait()
		return false
	}
}

// FUNCIONES DE WEAPON
func (w *Weapons) createWeapon(name string, cost float64, weight float64, description string, nroDice int, damage int, oneHand bool, finesse bool, properties *WeaponsProperty) {
	weapon := Weapon{
		name:        name,
		cost:        cost,
		weight:      weight,
		description: description,
		nroDice:     nroDice,
		damage:      damage,
		oneHand:     oneHand,
		equiped:     false,
		finesse:     finesse,
		properties:  *properties,
	}
	w.weapons = append(w.weapons, weapon)
}
func createProperty(ammo, finesse, heavy, light, loading, raange, reach, throw, twoHanded bool) *WeaponsProperty {
	property := WeaponsProperty{
		Ammunition: ammo,
		Finesse:    finesse,
		Heavy:      heavy,
		Light:      light,
		Loading:    loading,
		Range:      raange,
		Reach:      reach,
		Thrown:     throw,
		TwoHanded:  twoHanded,
	}
	return &property
}

func (c *Creatures) equipWeapon(creature *Creature, weapon *Weapon) {
	for _, w := range creature.weapons {
		w.equiped = false
	}
	weapon.equiped = true
	creature.weapons = append(creature.weapons, *weapon)

}

func (w *Weapons) findWeapon(name string) *Weapon {
	for _, weapon := range w.weapons {
		if weapon.name == name {
			return &weapon
		}
	}
	return nil
}
func (c *Creatures) findEquipedWeapon(creature *Creature) *Weapon {
	for _, weapon := range creature.weapons {
		if weapon.equiped {
			return &weapon
		}
	}
	return nil
}

func (c *Creatures) printEquipedWeapon(creature *Creature) {
	equipedWeapon := c.findEquipedWeapon(creature)
	if equipedWeapon != nil {
		fmt.Println("Arma equipada: ", equipedWeapon.name)
	} else {
		fmt.Println("No hay arma equipada")
	}
}

func (c *Creatures) printAllWeapons(creature *Creature) {
	weapons := creature.weapons
	if len(weapons) > 0 {
		fmt.Println("Armas:")
		for _, weapon := range weapons {
			if weapon.equiped {
				fmt.Println(weapon.name, "Daño:", weapon.nroDice, "d", weapon.damage, "*")
			} else {
				fmt.Println(weapon.name, " Daño: ", weapon.nroDice, "d", weapon.damage)
			}
		}
	} else {
		fmt.Println("No hay armas equipadas")
	}
}

// FUNCIONES DE ARMOR

func (a *Armors) createArmor(name string, cost float64, weight float64, description string, armorClass int, shield bool) {
	armor := Armor{
		name:        name,
		cost:        cost,
		weight:      weight,
		description: description,
		armorClass:  armorClass,
		shield:      shield,
		equiped:     false,
	}
	a.armors = append(a.armors, armor)
}

func (c *Creatures) equipArmor(creature *Creature, armor *Armor) {
	// Desequipar armadura actual si existe
	for _, arm := range creature.armors {
		if arm.equiped && !arm.shield {
			arm.equiped = false
		}
	}
	armor.equiped = true
	creature.ArmorClass = armor.armorClass
	creature.armors = append(creature.armors, *armor)
}

func (a *Creatures) equipShield(creature *Creature, armor *Armor) {
	for _, arm := range creature.armors {
		if arm.equiped && arm.shield {
			arm.equiped = false
		}
	}
	armor.equiped = true
	creature.ArmorClass = creature.ArmorClass + armor.armorClass
	creature.armors = append(creature.armors, *armor)
}
func (a *Armors) findArmor(name string) *Armor {
	for _, armor := range a.armors {
		if armor.name == name {
			return &armor
		}
	}
	return nil
}
func (c *Creatures) findEquipedArmor(creature *Creature) *Armor {
	for _, armor := range creature.armors {
		if armor.equiped {
			return &armor
		}
	}
	return nil
}

func (c *Creatures) printEquipedArmor(creature *Creature) {
	equipedArmor := c.findEquipedArmor(creature)
	if equipedArmor != nil {
		fmt.Println("Armadura equipada: ", equipedArmor.name)
	} else {
		fmt.Println("No hay armadura equipada")
	}
}
func (c *Creatures) printAllArmors(creature *Creature) {
	armors := creature.armors
	if len(armors) > 0 {
		fmt.Println("Armaduras:")
		for _, armor := range armors {
			if armor.equiped {
				fmt.Println(armor.name, "AC:", armor.armorClass, "*")
			} else {
				fmt.Println(armor.name, "AC:", armor.armorClass)
			}

		}
	} else {
		fmt.Println("No hay armaduras")
	}
}

func (c *Creatures) printCreature(creature *Creature) {
	if creature != nil {
		fmt.Println("**************************")
		fmt.Println("Hoja de", creature.name)
		fmt.Println("**************************")
		fmt.Println("Tamaño: ", creature.size)
		fmt.Println("Raza: ", creature.race)
		fmt.Println("Nivel: ", creature.Level)
		fmt.Println("Armadura: ", creature.ArmorClass)
		fmt.Println("Puntos de vida: ", creature.CurrentHitPoints, "/", creature.MaxHitPoints)
		fmt.Println("Fuerza: ", creature.Strength, "Modificador de Fuerza:", creature.StrMod)
		fmt.Println("Destreza: ", creature.Dexterity, "Modificador de Destreza:", creature.DexMod)
		fmt.Println("Consitucion: ", creature.Consitution, "Modificador de Constitución", creature.ConMod)
		fmt.Println("Inteligencia: ", creature.Intelligence, "Modificador de Inteligencia:", creature.IntMod)
		fmt.Println("Sabiduria: ", creature.Wisdom, "Modificador de Sabiduria:", creature.WisMod)
		fmt.Println("Carisma: ", creature.Charisma, "Modificador de Carisma:", creature.ChaMod)
		fmt.Println("Bono de proficiencia: ", creature.ProficiencyBonus)
		fmt.Println("Experiencia: ", creature.xp)
		fmt.Println("Iniciativa: ", creature.iniciative)
		fmt.Println("**************************")
		fmt.Println("Inventario (*equipado)")
		c.printAllArmors(creature)
		c.printAllWeapons(creature)
		fmt.Println("**************************")
		wait()
	}
}
