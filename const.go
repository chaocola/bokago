package bokago

var (
	BOKA                   *Config
	privateExpire          int64 = 5000
	privateReferer               = `https://s3.boka.vc/`
	privateTokenTaskSignal       = make(chan bool, 1)
)
