package main

import (
	"fmt"
	"math"
	"math/rand/v2"
)

func main() {
	fmt.Println("******************************************************************************************************************************")
	fmt.Println("Creación de personaje")
	fmt.Print("Ingrese el nombre de tu personaje: ")
	var name string
	fmt.Scan(&name)
	c := &Creatures{}
	w := &Weapons{}
	a := &Armors{}
	c.createCreature(name, medium, humanoide, 1, 12, 16, 10, 14, 8, 12, 12)
	fmt.Println("******************************************************************************************************************************")
	fmt.Println("En un reino lejano, la princesa Elena ha sido raptada por una horda de goblins y llevada a una caverna oculta en el corazón del Bosque Oscuro. El rey, desesperado, promete una gran recompensa a quien la rescate.Un valiente guerrero, armado con su espada y su coraje, se adentra en la caverna de los goblins. Enfrentando trampas y criaturas peligrosas, se adentra en la oscuridad con un solo objetivo: rescatar a la princesa y devolver la esperanza al reino. ¿Logrará nuestro héroe superar los desafíos y salvar a la princesa? Solo su valentía y tu habilidad decidirán el final de esta historia.")
	fmt.Println("******************************************************************************************************************************")
	fmt.Println("Antes de adentrarte a la caverna donde esta la princesa, debes elegir un arma")
	fmt.Println("1. Espada de dos manos")
	fmt.Println("2. Espada de una mano + Escudo")
	var equipoInicial int
	for {
		fmt.Print("Ingrese la opción que desea elegir: ")
		fmt.Scanf("%d", &equipoInicial)
		fmt.Println("******************************************************************************************************************************")
		fmt.Println(equipoInicial)
		fmt.Println("******************************************************************************************************************************")
		if equipoInicial == 1 || equipoInicial == 2 {
			fmt.Println("Eligio la espada a dos manos")
			w.createWeapon("Zweihänder", 50, 6, "La Zweihänder es una gran espada a dos manos, usada por infantería medieval para romper formaciones.", 2, 6, false)
			c.equipWeapon(c.findCreature(name), w.findWeapon("Zweihänder"))
			a.createArmor("Cota de Malla", 75, 55, "La cota de malla es una armadura de anillos de metal entrelazados, ofreciendo flexibilidad y protección.", 16, false)
			c.equipArmor(c.findCreature(name), a.findArmor("Cota de Malla"))
			c.printCreature(name)
			break
		} else if equipoInicial == 2 {
			fmt.Println("Eligio la espada + escudo")
			w.createWeapon("Espada larga", 15, 3, "La espada larga es una espada medieval de doble filo, usada con una o ambas manos.", 1, 8, true)
			c.equipWeapon(c.findCreature(name), w.findWeapon("Espada larga"))
			a.createArmor("Cota de Malla", 75, 55, "La cota de malla es una armadura de anillos de metal entrelazados, ofreciendo flexibilidad y protección.", 18, false)
			c.printCreature(name)
			break
		} else {
			fmt.Println("Opción no valida")
		}
	}

}

type Creature struct {
	name             string
	size             Size
	race             Race
	Level            int
	ArmorClass       int
	MaxHitPoints     int
	CurrentHitPoints int
	Strength         int
	Dexterity        int
	Consitution      int
	Intelligence     int
	Wisdom           int
	Charisma         int
	ProficiencyBonus int
	attackBonus      int
	damageBonus      int
	weapon           Weapon
	armor            Armor
	xp               int
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
	humanoide Race = "Humanoide"
)

type Weapon struct {
	name        string
	cost        float64
	weight      float64
	description string
	nroDice     int
	damage      int
	oneHand     bool
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

func (c *Creatures) equipWeapon(creature *Creature, weapon *Weapon) {
	creature.weapon = *weapon
}

func (w *Weapons) createWeapon(name string, cost float64, weight float64, description string, nroDice int, damage int, oneHand bool) {
	weapon := Weapon{
		name:        name,
		cost:        cost,
		weight:      weight,
		description: description,
		nroDice:     nroDice,
		damage:      damage,
		oneHand:     oneHand,
	}
	w.weapons = append(w.weapons, weapon)
}

func (w *Weapons) findWeapon(name string) *Weapon {
	for _, weapon := range w.weapons {
		if weapon.name == name {
			return &weapon
		}
	}
	return nil
}

func (c *Creatures) createCreature(name string, size Size, race Race, level float64, HP int, STR int, DEX int, CON int, INT int, WIS int, CHA int) {
	pb := math.Ceil((level / 4) + 1)
	attB := ((STR - 10) / 2) + int(pb)
	dmgB := (STR - 10) / 2
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
		Dexterity:        DEX,
		Consitution:      CON,
		Intelligence:     INT,
		Wisdom:           WIS,
		Charisma:         CHA,
		ProficiencyBonus: int(pb),
		attackBonus:      attB,
		damageBonus:      dmgB,
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

func (c *Creatures) ataque(attackerName string, defenderName string) {
	attacker := c.findCreature(attackerName)
	defender := c.findCreature(defenderName)
	fmt.Println("**************************")
	fmt.Println(attacker, "Ataca")
	attackerToHit := rand.IntN(20) + attacker.attackBonus
	fmt.Println("El ataque fue un: ", attackerToHit)
	fmt.Println("La armadura del ", defender.name, "es: ", defender.ArmorClass)
	if attackerToHit >= defender.ArmorClass {
		weaponDmg := attacker.weapon.damage
		weaponDice := attacker.weapon.nroDice
		var damage int
		for weaponDice > 0 {
			damage += rand.IntN(weaponDmg)
			weaponDice--
		}
		damage += attacker.damageBonus
		fmt.Println("La vida de ", defender.name, "es: ", defender.CurrentHitPoints)
		defender.CurrentHitPoints = defender.CurrentHitPoints - damage
		if defender.CurrentHitPoints <= 0 {
			fmt.Println("El daño fue de: ", damage)
			fmt.Println(defender, " ha muerto")
			fmt.Println("**************************")
		} else {
			fmt.Println("El daño fue de: ", damage)
			fmt.Println("La vida de ", defender.name, "es: ", defender.CurrentHitPoints)
			fmt.Println("**************************")
		}
	} else {
		fmt.Println("El ataque fallo")
	}
}
func (a *Armors) createArmor(name string, cost float64, weight float64, description string, armorClass int, shield bool) {
	armor := Armor{
		name:        name,
		cost:        cost,
		weight:      weight,
		description: description,
		armorClass:  armorClass,
		shield:      shield,
	}
	a.armors = append(a.armors, armor)
}

func (c *Creatures) equipArmor(creature *Creature, armor *Armor) {
	creature.armor = *armor
	creature.ArmorClass = creature.armor.armorClass
}

func (a *Armors) findArmor(name string) *Armor {
	for _, armor := range a.armors {
		if armor.name == name {
			return &armor
		}
	}
	return nil
}

func (c *Creatures) printCreature(name string) {
	creature := c.findCreature(name)
	if creature != nil {
		fmt.Println("**************************")
		fmt.Println("Caracteristicas de ", creature.name)
		fmt.Println("Tamaño: ", creature.size)
		fmt.Println("Raza: ", creature.race)
		fmt.Println("Nivel: ", creature.Level)
		fmt.Println("Armor Class: ", creature.ArmorClass)
		fmt.Println("Max Hit Points: ", creature.MaxHitPoints)
		fmt.Println("Current Hit Points: ", creature.CurrentHitPoints)
		fmt.Println("Strength: ", creature.Strength)
		fmt.Println("Dexterity: ", creature.Dexterity)
		fmt.Println("Consitution: ", creature.Consitution)
		fmt.Println("Intelligence: ", creature.Intelligence)
		fmt.Println("Wisdom: ", creature.Wisdom)
		fmt.Println("Charisma: ", creature.Charisma)
		fmt.Println("Proficiency Bonus: ", creature.ProficiencyBonus)
		fmt.Println("Attack Bonus: ", creature.attackBonus)
		fmt.Println("Damage Bonus: ", creature.damageBonus)
		fmt.Println("Weapon: ", creature.weapon.name)
		fmt.Println("**************************")
	}
}
