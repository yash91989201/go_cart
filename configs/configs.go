package configs

func Init() {
	loadEnv()
	createDBConnection()
}
