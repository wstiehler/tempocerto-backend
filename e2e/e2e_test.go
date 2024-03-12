//go:build e2e
// +build e2e

package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/wstiehler/tempocerto-backend/e2e/cmd"
// 	"github.com/wstiehler/tempocerto-backend/internal/domain/role"
// )

// var url string
// var timeout int

// func readEnv() (string, int) {
// 	url = os.Getenv("APPLICATION_URL")
// 	timeoutString := os.Getenv("TEST_TIMEOUT")

// 	timeout, err := strconv.Atoi(timeoutString)
// 	if err != nil {
// 		timeout = 3
// 	}
// 	return url, timeout
// }

// func TestApiHealth(t *testing.T) {
// 	assert := assert.New(t)
// 	readEnv()

// 	t.Run("Health status", func(t *testing.T) {
// 		client := cmd.NewProjectApi(url)
// 		health, err := client.ApiHealth()

// 		fmt.Printf("health: %+v\n", health)
// 		assert.Equal(err, nil)
// 	})
// }

// func TestMethodCompany(t *testing.T) {
// 	assert := assert.New(t)
// 	readEnv()

// 	roleCreated := role.RoleEntity{
// 		Role: "aadmin",
// 		Permissions: []role.PermissionEntity{
// 			{
// 				Name: "modifier",
// 			},
// 			{
// 				Name: "watcher",
// 			},
// 			{
// 				Name: "deleter",
// 			},
// 		},
// 	}

// 	var roleId uint

// 	t.Run("CreateRole_When_return_must_be_success", func(t *testing.T) {
// 		client := cmd.NewProjectApi(url)
// 		role, err := client.CreateRole(roleCreated)

// 		roleId = role.Id

// 		assert.Equal(role.Role, "aadmin")
// 		assert.Equal(err, nil)
// 	})

// 	t.Run("GetRoleByID_When_return_must_be_one", func(t *testing.T) {
// 		client := cmd.NewProjectApi(url)

// 		role, err := client.GetRoleByID(roleId)

// 		assert.Equal(role.Role, "aadmin")
// 		assert.Equal(err, nil)

// 	})

// }
