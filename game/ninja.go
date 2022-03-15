package game

import (
	"log"
	"sync"
	"time"

	"github.com/sflewis2970/ninja-game/common"
	"github.com/sflewis2970/ninja-game/ninja"
)

const WG_GO_Rountine_Cnt int = 1

var wg sync.WaitGroup
var once sync.Once
var targetsSize int = 0
var nTargetsEliminated int = 0
var lastIDUsed int = 0

func StartGame() {
	startTime := time.Now()

	superNinja := createNinja("Shawn Lewis", "Sword")
	log.Println()

	startMission(startTime)
	defer endMission(startTime)

	// Create Targets
	targetList := []ninja.Ninja{}

	// Target 1 - Jon Doe
	target := createNinja("Jon Doe", "Staff")
	log.Println("Created ninja target:", target.Name, "ID:", target.ID)
	targetList = addTarget(target, targetList)

	// Target 2 - Perry Mason
	target = createNinja("Perry Mason", "Kunai")
	log.Println("Created ninja target:", target.Name, "ID:", target.ID)
	targetList = addTarget(target, targetList)

	// Target 3 - Jon Smith
	target = createNinja("Jon Smith", "Star")
	log.Println("Created ninja target:", target.Name, "ID:", target.ID)
	targetList = addTarget(target, targetList)

	// Target 4 - Curly Joe
	target = createNinja("Curly Joe", "Star")
	log.Println("Created ninja target:", target.Name, "ID:", target.ID)
	targetList = addTarget(target, targetList)

	targetsSize = len(targetList)
	manageTargets(superNinja, targetList)

	wg.Wait()

	if nTargetsEliminated == targetsSize {
		log.Println("All targets eliminated -- Mission complete!")
	}
}

func addTarget(target *ninja.Ninja, targetList interface{}) []ninja.Ninja {
	tl := targetList.([]ninja.Ninja)
	tl = append(tl, *target)

	return tl
}

func manageTargets(sn *ninja.Ninja, targetList []ninja.Ninja) {
	// Attack targets
	targetIdx := 0

	for targetIdx < targetsSize {
		// Get current target info
		target := targetList[targetIdx]

		// Update wg counter
		wg.Add(WG_GO_Rountine_Cnt)

		// launch thread-like function to handling attacks
		go underTakeTarget(sn, &target)

		// Advance to the next target
		targetIdx++
	}
}

func underTakeTarget(sn *ninja.Ninja, target *ninja.Ninja) {
	defer wg.Done()

	snOnce := func() {
		log.Println("Super Ninja", sn.Name, "has been eliminated.")
		sn.Eliminated = true
	}

	for {
		if !target.Acquired {
			target.Acquired = acquire()

			if target.Acquired {
				log.Println("Target:", target.Name, "has been acquired")
			} else {
				log.Println("Searching for", target.Name+"...")
			}
		}

		// Handle targets that have yet to be eliminated
		if !target.Eliminated {
			if target.Acquired {
				log.Println()
				log.Println(sn.Name, "has", sn.Health, "units of health")
				log.Println(sn.Name, "attacking", target.Name)
				log.Println(target.Name, "has", target.Health, "units of health")
				log.Println()

				// Attack target once acquired
				attack(sn, target)

				log.Println()
				log.Println(sn.Name, "attack resulted in the following updated stats")
				log.Println(target.Name, "has", target.Health, "units of health")
				log.Println(sn.Name, "has", sn.Health, "units of health")
				log.Println()

				// Determine if super ninja has been eliminated
				if sn.Health <= 0 {
					once.Do(snOnce)
					break
				}

				// Increment targets eliminated counter
				if target.Health <= 0 {
					log.Println("Target:", target.Name, "has been eliminated.")
					target.Eliminated = true
					nTargetsEliminated++
					if nTargetsEliminated == targetsSize {
						break
					}
				}
			} else {
				log.Println("Waiting for", target.Name, "to be acquired...")
				time.Sleep(time.Millisecond * time.Duration(ninja.ACQUIRE_PAUSE))
			}
		} else {
			log.Println(target.Name, "has been eliminated.")
			break
		}
	}
}

func acquire() bool {
	// Acquire target
	acquireTargetVal := common.GenerateFloat64Vals(float64(ninja.ACQUIRE_TARGET_RANGE), float64(ninja.ACQUIRE_TARGET_MIN))
	if ninja.AcquisitionType(acquireTargetVal) > ninja.ACQUIRE_TARGET_SUCCESS_MIN && ninja.AcquisitionType(acquireTargetVal) < ninja.ACQUIRE_TARGET_SUCCESS_MAX {
		// Target has been acquired
		return true
	}

	return false
}

func attack(sn *ninja.Ninja, target *ninja.Ninja) {
	// Super Ninja attacks target
	weaponDamage := common.GenerateFloat64Vals(sn.Weapon_Strength, ninja.WEAPON_STRENGTH_MIN)

	// Pause in the action
	log.Println("Pause after attack...")
	time.Sleep(time.Millisecond * time.Duration(ninja.ATTACK_PAUSE))

	// Target response
	attackDamage, snAttackDamage := targetResponse(target, sn, weaponDamage)

	// Pause in the action
	log.Println("Pause after response...")
	time.Sleep(time.Millisecond * time.Duration(ninja.RESPONSE_PAUSE))

	// Update Health for compatents
	sn.Health += snAttackDamage
	target.Health += attackDamage
}

func targetResponse(tn *ninja.Ninja, sn *ninja.Ninja, weaponDamage float64) (float64, float64) {
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

func createNinja(name string, weapon string) *ninja.Ninja {
	en := new(ninja.Ninja)

	// Since Golang does NOT have a pre-increment operator,
	// We increment the variable before assigning the variable
	lastIDUsed++

	en.ID = lastIDUsed
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
