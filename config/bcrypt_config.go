package config

type bcryptConfig struct {
	MIN_COST     int
	MAX_COST     int
	DEFAULT_COST int
}

func GetBcryptConfig() bcryptConfig {
	config := bcryptConfig{
		MIN_COST:     5,
		MAX_COST:     15,
		DEFAULT_COST: 10,
	}
	
	return config
}
