package config

// Объявляем и инициализируем переменные с параметрами подключения к базе данных
var (
	// DbHost - хост базы данных (адрес сервера)
	DbHost = "192.168.1.92"
	// DbPort - порт базы данных (стандартный порт MySQL/MariaDB)
	DbPort = "3306"
	// DbUser - имя пользователя для подключения к базе данных
	DbUser = "demid"
	// DbPassword - пароль пользователя для подключения к базе данных
	DbPassword = "Gandon345"
	// DbName - название базы данных к которой подключаемся
	DbName = "testgovna"
)

// GetDSN возвращает строку подключения к базе данных
func GetDSN() string {
	return DbUser + ":" + DbPassword + "@tcp(" + DbHost + ":" + DbPort + ")/" + DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
}