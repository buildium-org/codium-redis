package steps

import (
	"net"
	"time"

	testrunners "github.com/buildium-org/codium-harness/testrunners"
)

func TestRedisPort6379Step(config *testrunners.ServerTestConfig) error {
	config.Logger.LogTitle("Test Redis Port 6379")
	config.Logger.LogInfo("Testing if the Redis port 6379 is open")
	config.Server.Start()
	time.Sleep(1000 * time.Millisecond)
	defer config.Server.Stop()

	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		config.Logger.LogError("Failed to connect to Redis port 6379")
		return err
	}
	defer conn.Close()
	config.Logger.LogSuccess("Redis port 6379 is open")
	return nil
}
