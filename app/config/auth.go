package config

type AuthConfig struct {
	Time    uint32
	Memory  uint32
	Threads uint8
	KeyLen  uint32
	SaltLen int32
}

var Auth = &AuthConfig{
	Time:    3,
	Memory:  32 * 1024,
	Threads: 4,
	KeyLen:  256,
	SaltLen: 32,
}
