package config

type EnvConfig struct {
	AppPort   int    `env:"APP_PORT,required,notEmpty"`
	DBUser    string `env:"DB_USER,required,notEmpty"`
	DBPass    string `env:"DB_PASS,required,notEmpty"`
	DBHost    string `env:"DB_HOST,required,notEmpty"`
	DBPort    int    `env:"DB_PORT,required,notEmpty"`
	DBName    string `env:"DB_NAME,required,notEmpty"`
	JwtSecret string `env:"JWT_SECRET,required,notEmpty"`
}
