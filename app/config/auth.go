package config

type AuthConfig struct {
	Time    uint32
	Memory  uint32
	Threads uint8
	KeyLen  uint32
	SaltLen int32
}

var Auth = &AuthConfig{
	Time:    1,
	Memory:  64 * 1024,
	Threads: 4,
	KeyLen:  32,
	SaltLen: 32,
}
