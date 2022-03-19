package game

import (
	"log"
	"time"

	"github.com/sflewis2970/ninja-game/common"
	"github.com/sflewis2970/ninja-game/ninja"
)

const waitDuration = 750

func StartGame() {
	startTime := time.Now()

	fileName := "ninja.json"
	superNinja, targetList, err := ninja.ReadAssignmentFile(fileName)
	if err != nil {
		log.Println("Error reading assignment file:", err.Error())
		return
	}

	// Update Super Ninja health
	if superNinja.Health == 0 {
		healthRange := ninja.MAX_HEALTH - ninja.MIN_HEALTH
		superNinja.Health = common.GenerateFloat64Vals(healthRange, ninja.MIN_HEALTH)
	}

	startMission(startTime)
	defer endMission(startTime)

	for _, target := range targetList {
		log.Println(superNinja.Name, "out to eliminate", target.Name)
		eliminateTarget(superNinja, &target)

		// If super ninja is eliminated end mission fails
		if superNinja.Eliminated {
			log.Println(superNinja.Name, "has been eliminated - mission failed")
			break
		}
	}

}

func eliminateTarget(sn *ninja.SuperNinja, target *ninja.TargetNinja) {
	log.Println(sn.Name, "found", target.Name)

	for {
		time.Sleep(time.Millisecond * time.Duration(waitDuration))

		// Attack target
		attack(sn, target)

		log.Println(sn.Name, "health =", sn.Health)
		if sn.Health <= 0 {
			sn.Eliminated = true
			break
		}

		log.Println(target.Name, "health =", target.Health)
		if target.Health <= 0 {
			log.Println(target.Name, "eliminated!")
			target.Eliminated = true
			break
		} else {
			time.Sleep(time.Millisecond * time.Duration(waitDuration))
		}
	}
}

func attack(sn *ninja.SuperNinja, target *ninja.TargetNinja) {
	// Super Ninja attacks target
	weaponDamage := common.GenerateFloat64Vals(sn.Weapon_Strength, ninja.WEAPON_STRENGTH_MIN)

	// Pause in the action
	time.Sleep(time.Millisecond * time.Duration(ninja.ATTACK_PAUSE))

	// Target response
	attackDamage, snAttackDamage := targetResponse(sn, target, weaponDamage)

	// Pause in the action
	time.Sleep(time.Millisecond * time.Duration(ninja.RESPONSE_PAUSE))

	// Update Health for compatents
	sn.Health += snAttackDamage
	target.Health += attackDamage
}

func targetResponse(sn *ninja.SuperNinja, tn *ninja.TargetNinja, weaponDamage float64) (float64, float64) {
	// Response type
	responseType := ninja.ResponseType(common.GenerateIntVals(ninja.RT_Count, ninja.RT_Count_Min))

	// Select response type
	respAttempt := float64(0)
	attackDamage := float64(0)
	snAttackDamage := float64(0)

	switch responseType {
	case ninja.RT_Block:
		log.Println("response type: block")
		respAttempt = common.GenerateFloat64Vals(float64(ninja.Resp_Block_Attempt_Max-ninja.Resp_Block_Attempt_Min), float64(ninja.Resp_Block_Attempt_Min))

		if ninja.ResponseAttempt(respAttempt) > ninja.Block_Attempt_Min && ninja.ResponseAttempt(respAttempt) < ninja.Block_Attempt_Max {
			// Attack block was successful
			attackDamage = tn.Health * -0.1
		} else {
			// Attack block was unsuccessful
			attackDamage = weaponDamage * -1
		}

	case ninja.RT_Dodge:
		log.Println("response type: dodge")
		respAttempt = common.GenerateFloat64Vals(float64(ninja.Resp_Dodge_Attempt_Max-ninja.Resp_Dodge_Attempt_Min), float64(ninja.Resp_Dodge_Attempt_Min))

		if ninja.ResponseAttempt(respAttempt) > ninja.Dodge_Attempt_Min && ninja.ResponseAttempt(respAttempt) < ninja.Dodge_Attempt_Max {
			// Attack dodge is successful
			attackDamage = (tn.Health - 1) * -1
		} else {
			// Attack dodge is unsuccessful
			attackDamage = weaponDamage * -1
		}

	case ninja.RT_Attack:
		log.Println("response type: attack")
		respAttempt = common.GenerateFloat64Vals(float64(ninja.Resp_Attack_Attempt_Max-ninja.Resp_Attack_Attempt_Min), float64(ninja.Resp_Attack_Attempt_Min))

		healthReduction := float64(0)
		if ninja.ResponseAttempt(respAttempt) > ninja.Attack_Attempt_Min && ninja.ResponseAttempt(respAttempt) < ninja.Attack_Attempt_Max {
			// Attack (hit while attacking) is successful
			attackDamage = tn.Health * -0.5
			healthReduction = sn.Health * -.25
		} else {
			// Attack (hit while attacking) is unsuccessful
			attackDamage = 0
			healthReduction = sn.Health * -.75
		}

		snAttackDamage = common.GenerateFloat64Vals(healthReduction, ninja.HEALTH_REDUCTION_MIN)
	}

	return attackDamage, snAttackDamage
}

func createNinja(name string, weapon string) *ninja.SuperNinja {
	en := new(ninja.SuperNinja)

	en.Name = name
	en.Weapon = weapon
	en.Weapon_Strength = common.GenerateFloat64Vals(ninja.MAX_WEAPON-ninja.MIN_WEAPON, ninja.MIN_WEAPON)
	en.Health = common.GenerateFloat64Vals(ninja.MAX_HEALTH-ninja.MIN_HEALTH, ninja.MIN_HEALTH)

	return en
}

func startMission(startTime time.Time) {
	log.Println("Mission started")
	log.Println("Start time: ", startTime)
	log.Println()
}

func endMission(startTime time.Time) {
	log.Println("End time: ", time.Since(startTime))
}
