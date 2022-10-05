package bootstrap

import "os"

type driverSupabase struct {
	User     string
	Password string
	Host     string
	Port     string
	DbName   string
}

func readEnvSupabase() (driverSupabase, error) {
	return driverSupabase{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DbName:   os.Getenv("DB_NAME"),
	}, nil
}

type AppDriver struct {
	Port    string
	BaseURL string
	DB      driverSupabase
}

func GetAppDriver() (AppDriver, error) {
	driverDB, err := readEnvSupabase()
	if err != nil {
		return AppDriver{}, nil
	}
	return AppDriver{DB: driverDB, Port: os.Getenv("PORT"), BaseURL: os.Getenv("BASE_URL")}, nil
}
