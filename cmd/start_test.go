package cmd_test

import (
	"os"
	"testing"

	"github.com/nabhdeep/gateway-cli/cmd"
	"github.com/stretchr/testify/assert"
)

func TestStartServer(t *testing.T) {
	t.Run("Check if server already exist", func(t *testing.T) {

		// change the path to base
		os.Chdir("..")
		path := "tmp/gateway.pid"

		// create and delete
		os.Create(path)
		os.WriteFile(path, []byte("TEST"), 0644)
		defer os.Remove(path)

		get := cmd.Check_server_already_exist()
		assert.True(t, get, "Sever file should exist")
	})
}
