package db

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"gotest.tools/assert"
)

func TestListAllTags(t *testing.T) {
	os.Setenv("DATABASE_URL", "postgresql://user2020:pass2020@localhost:5432/stop-analyzing-api")

	dbCli, err := Connect()
	if err != nil {
		fmt.Println("Error at test List All Tags - Connect to Database:", err)
	}
	defer dbCli.Disconnect()
	tags, err := dbCli.GetAllTags()
	if err != nil {
		fmt.Println("Error at test List All Tags - Get All Tags:", err)
	}

	expectedResult := []TagModel{}
	assert.Equal(t, true, reflect.DeepEqual(tags, expectedResult))
}
