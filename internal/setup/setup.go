package setup

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
 	"golang_tailwinds/internal/constant"
)

func SetUpIt() {

	err := godotenv.Load()
	if err != nil {
		panic("No env")
	}

	portStr := os.Getenv("PORTPOSTGRE")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic("Invalid PORT env value")
	}

	constant.DbNamePostGres = os.Getenv("DBNAMEPOSTGRE")
	constant.HostPostGres = os.Getenv("HOSTPOSTGRE")
	constant.PortPostGres = port
	constant.UserPostGres = os.Getenv("USERPOSTGRE")
	constant.PasswordPostGres = os.Getenv("PASSWORDPOSTGRE")
	fmt.Println(os.Getenv("USERPOSTGRE"))

	// ConnectDb()

}
