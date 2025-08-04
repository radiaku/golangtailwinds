package setup

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
 	"golang_tailwinds/internal/constant"
)

func SetUpIt() {

	_ = godotenv.Load()
	portStr := os.Getenv("PORTPOSTGRE")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Invalid PORT env value: %v", err)
	}
	constant.DbNamePostGres = os.Getenv("DBNAMEPOSTGRE")
	constant.HostPostGres = os.Getenv("HOSTPOSTGRE")
	constant.PortPostGres = port
	constant.UserPostGres = os.Getenv("USERPOSTGRE")
	constant.PasswordPostGres = os.Getenv("PASSWORDPOSTGRE")
	fmt.Println(os.Getenv("USERPOSTGRE"))

	// ConnectDb()

}
