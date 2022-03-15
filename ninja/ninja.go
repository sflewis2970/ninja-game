package ninja

type Ninja struct {
	ID              int
	Name            string
	Weapon          string
	Acquired        bool
	Eliminated      bool
	Weapon_Strength float64
	Health          float64
}

const (
	ATTACK_PAUSE         int     = 3000
	RESPONSE_PAUSE       int     = 1500
	WEAPON_STRENGTH_MIN  float64 = 1
	HEALTH_REDUCTION_MIN float64 = 1
	MIN_HEALTH           float64 = 100
	MAX_HEALTH           float64 = 301
	MIN_WEAPON           float64 = 10
	MAX_WEAPON           float64 = 101
)

type AcquisitionType float32

const ACQUIRE_PAUSE int = 5000

const (
	ACQUIRE_TARGET_MIN         AcquisitionType = 1
	ACQUIRE_TARGET_RANGE       AcquisitionType = 250
	ACQUIRE_TARGET_SUCCESS_MIN AcquisitionType = 100
	ACQUIRE_TARGET_SUCCESS_MAX AcquisitionType = 150
)

type ResponseType int32

const (
	RT_Block     ResponseType = 0
	RT_Dodge     ResponseType = 1
	RT_Attack    ResponseType = 2
	RT_Count     int          = 3
	RT_Count_Min int          = int(RT_Block)
)

type ResponseAttempt int32

const (
	Block_Attempt_Min  ResponseAttempt = 175
	Block_Attempt_Max  ResponseAttempt = 225
	Dodge_Attempt_Min  ResponseAttempt = 150
	Dodge_Attempt_Max  ResponseAttempt = 250
	Attack_Attempt_Min ResponseAttempt = 110
	Attack_Attempt_Max ResponseAttempt = 290
)

type Response float64

const (
	// Blocking Response
	Resp_Block_Attempt_Min Response = 100
	Resp_Block_Attempt_Max Response = 300
	Resp_Block_Adjustment  Response = 0.1

	// Dodge Response
	Resp_Dodge_Attempt_Min Response = 100
	Resp_Dodge_Attempt_Max Response = 300
	Resp_Dodge_Adjustment  Response = 0

	// Attack Response
	Resp_Attack_Attempt_Min Response = 100
	Resp_Attack_Attempt_Max Response = 300
	Resp_Attack_Adjustment  Response = 0
)
