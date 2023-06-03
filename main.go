package main

import ( "fmt" "io/ioutil" "log"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
import "handlers"

type Config struct {
	LogLevel     string `yaml:"log_level"`
	DBConnection string `yaml:"db_connection"`
	Port         int    `yaml:"port"`
}

type User struct {
	gorm.Model
	Name  string
	Email string
	Age   int
}

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Read the configuration file
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Failed to read configuration file: %v", err)
	}

	// Create a Config struct
	var config Config

	// Unmarshal the YAML data into the Config struct
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Failed to unmarshal configuration: %v", err)
	}

	// Use the configuration values
	fmt.Println("Log Level:", config.LogLevel)
	fmt.Println("DB Connection:", config.DBConnection)
	fmt.Println("Port:", config.Port)

	// Establish a database connection
	db, err := gorm.Open(mysql.Open(config.DBConnection), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// AutoMigrate the user table
	db.AutoMigrate(&User{})

	fmt.Println("Database connection established")

	// Create a new instance of the UserService
	userService := services.NewUserService(db)

	err = userService.CreateUser("John Doe", "johndoe@example.com", "password123")
	if err != nil {
		panic("Failed to create user")
	}

    // Define your API routes and handlers
    router.GET("/api/users", getUsersHandler)
    router.POST("/api/users", createUserHandler)
    router.GET("/api/users/:id", getUserHandler)
    router.PUT("/api/users/:id", updateUserHandler)
    router.DELETE("/api/users/:id", deleteUserHandler)
    router.GET("/healthz",healthzHandler)
    router.GET("/", rootHandler)
    router.GET("/hello/:name", helloHandler)
    // Run the server
    router.Run(":8080")
}
func getUsersHandler(c *gin.Context) {
    // Your logic to retrieve and return users
    c.JSON(200, gin.H{"message": "Get all users"})
}

func createUserHandler(c *gin.Context) {
    // Your logic to create a new user
    c.JSON(200, gin.H{"message": "Create a user"})
}

func getUserHandler(c *gin.Context) {
    // Your logic to retrieve and return a user by ID
    userID := c.Param("id")
    c.JSON(200, gin.H{"message": "Get user", "id": userID})
}

func updateUserHandler(c *gin.Context) {
    // Your logic to update a user by ID
    userID := c.Param("id")
    c.JSON(200, gin.H{"message": "Update user", "id": userID})
}

func deleteUserHandler(c *gin.Context) {
    // Your logic to delete a user by ID
    userID := c.Param("id")
    c.JSON(200, gin.H{"message": "Delete user", "id": userID})
}
func healthzHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "OK"})
}

func rootHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "OK"})
}
func helloHandler(c *gin.Context) {
	name := c.Param("name")
	message := fmt.Sprintf("Hello %s", name)
	c.String(200, message)
}
