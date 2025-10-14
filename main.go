package main

// Импортируем необходимые пакеты
import (
	"bufio"   // Пакет для буферизованного ввода
	"fmt"     // Пакет для форматированного ввода и вывода
	"os"      // Пакет для работы с операционной системой
	"strconv" // Пакет для преобразования строк в числа
	"strings" // Пакет для работы со строками

	"gorm.io/driver/mysql" // Драйвер для работы с MySQL/MariaDB
	"gorm.io/gorm"         // ORM библиотека для работы с базой данных

	"connect/Secret" // Пакет с конфигурацией базы данных
)

// Определяем структуру usergovnas, которая представляет таблицу в базе данных
type usergovnas struct {
	ID       uint   `gorm:"primarykey"` // Поле ID - основной ключ таблицы
	Name     string `gorm:"size:255"`   // Поле Name - строка длиной до 255 символов
	Sex      string `gorm:"size:10"`    // Поле Sex - строка длиной до 10 символов
	Sumgavna uint   // Поле Sumgavna - целое число без знака
}

// Основная функция программы
func main() {
	// Получаем строку подключения к базе данных из конфигурации
	dsn := config.GetDSN()

	// Устанавливаем соединение с базой данных используя GORM и MySQL драйвер
	// db - объект для работы с базой данных, err - переменная для хранения ошибки
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Проверяем успешность подключения к базе данных
	if err != nil {
		println("❌ FAILED TO CONNECT DATABASE") // Выводим сообщение об ошибке если подключение не удалось
		return                                  // Завершаем выполнение функции main
	}
	println("✅ DATABASE CONNECTED SUCCESSFULLY") // Выводим сообщение об успешном подключении

	// Выполняем автоматическую миграцию структуры usergovnas в базу данных
	// AutoMigrate создаст таблицу если её нет или обновит её структуру
	err = db.AutoMigrate(&usergovnas{})

	// Проверяем успешность выполнения миграции
	if err != nil {
		println("❌ FAILED TO MIGRATE DATABASE") // Выводим сообщение об ошибке миграции
		return                                  // Завершаем выполнение функции main
	}
	println("✅ DATABASE MIGRATED SUCCESSFULLY") // Выводим сообщение об успешной миграции

	// Создаем ридер для ввода данных с консоли
	reader := bufio.NewReader(os.Stdin)

	// Бесконечный цикл для добавления записей
	for {
		// Выводим меню выбора действия
		fmt.Println("\n=== МЕНЮ ДОБАВЛЕНИЯ ЗАПИСЕЙ ===")
		fmt.Println("1 - Добавить новую запись")
		fmt.Println("2 - Показать все записи")
		fmt.Println("0 - Выйти из программы")
		fmt.Print("Выберите действие: ")

		// Читаем выбор пользователя
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice) // Убираем символы переноса строки и пробелы

		// Обрабатываем выбор пользователя
		switch choice {
		case "1":
			addNewRecord(db, reader) // Вызываем функцию добавления новой записи
		case "2":
			showAllRecords(db) // Вызываем функцию показа всех записей
		case "0":
			fmt.Println("👋 Выход из программы...") // Сообщение о выходе
			return                                 // Завершаем программу
		default:
			fmt.Println("❌ Неверный выбор, попробуйте снова") // Сообщение об ошибке выбора
		}
	}
}

// Функция для добавления новой записи в таблицу
func addNewRecord(db *gorm.DB, reader *bufio.Reader) {
	fmt.Println("\n--- Добавление новой записи ---")

	// Запрашиваем имя пользователя
	fmt.Print("Введите имя: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name) // Убираем символы переноса строки и пробелы

	// Запрашиваем пол пользователя
	fmt.Print("Введите пол (male/female): ")
	sex, _ := reader.ReadString('\n')
	sex = strings.TrimSpace(sex) // Убираем символы переноса строки и пробелы

	// Запрашиваем значение Sumgavna
	fmt.Print("Введите значение Sumgavna: ")
	sumgavnaStr, _ := reader.ReadString('\n')
	sumgavnaStr = strings.TrimSpace(sumgavnaStr) // Убираем символы переноса строки и пробелы

	// Преобразуем строку в число
	sumgavna, err := strconv.Atoi(sumgavnaStr)
	if err != nil {
		fmt.Println("❌ Ошибка: Sumgavna должно быть числом!") // Сообщение об ошибке преобразования
		return
	}

	// Создаем новую запись пользователя
	newUser := usergovnas{
		Name:     name,
		Sex:      sex,
		Sumgavna: uint(sumgavna), // Преобразуем int в uint
	}

	// Сохраняем запись в базу данных
	result := db.Create(&newUser)
	if result.Error != nil {
		fmt.Println("❌ Ошибка при добавлении записи:", result.Error) // Сообщение об ошибке сохранения
		return
	}

	// Выводим сообщение об успешном добавлении
	fmt.Printf("✅ Запись успешно добавлена! ID: %d\n", newUser.ID)
}

// Функция для отображения всех записей из таблицы
func showAllRecords(db *gorm.DB) {
	fmt.Println("\n--- Все записи в таблице ---")

	// Создаем слайс для хранения всех записей
	var user []usergovnas

	// Получаем все записи из базы данных
	result := db.Find(&user)
	if result.Error != nil {
		fmt.Println("❌ Ошибка при получении записей:", result.Error) // Сообщение об ошибке чтения
		return
	}

	// Проверяем есть ли записи в таблице
	if len(user) == 0 {
		fmt.Println("📭 Таблица пуста") // Сообщение если таблица пуста
		return
	}

	// Выводим все записи в форматированном виде
	for _, user := range user {
		fmt.Printf("ID: %d | Name: %s | Sex: %s | Sumgavna: %d\n",
			user.ID, user.Name, user.Sex, user.Sumgavna)
	}

	// Выводим общее количество записей
	fmt.Printf("Всего записей: %d\n", len(user))
}
