package steps

import (
	"bufio"
	"fmt"
	"net"
	"sync"
	"time"

	testrunners "github.com/buildium-org/codium-harness/testrunners"
)

func TestRedisPort6379Step(config *testrunners.ServerTestConfig) error {
	config.Logger.LogTitle("Test Redis Port 6379")
	config.Logger.LogInfo("Testing if the Redis port 6379 is open")
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

func TestRedisPingStep(config *testrunners.ServerTestConfig) error {
	config.Logger.LogTitle("Test Redis Ping")
	config.Logger.LogInfo("Testing if the Redis server is responding to PING")
	time.Sleep(1000 * time.Millisecond)
	defer config.Server.Stop()

	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		config.Logger.LogError("Failed to connect to Redis port 6379")
		return err
	}
	defer conn.Close()
	conn.Write([]byte("PING\r\n"))
	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		config.Logger.LogError("Failed to read response from Redis server")
		return err
	}
	config.Logger.LogInfo(fmt.Sprintf("Response from Redis server: %s", response))
	if response != "+PONG\r\n" {
		config.Logger.LogError("Redis server did not respond with PONG")
		return fmt.Errorf("redis server did not respond with PONG")
	}
	config.Logger.LogSuccess("Redis server is responding to PING")
	return nil
}

func TestMultiplePingsStep(config *testrunners.ServerTestConfig) error {
	config.Logger.LogTitle("Test Multiple Pings")
	config.Logger.LogInfo("Testing if the Redis server is responding to multiple PINGs")
	time.Sleep(1000 * time.Millisecond)
	defer config.Server.Stop()

	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		config.Logger.LogError("Failed to connect to Redis port 6379")
		return err
	}
	defer conn.Close()
	responses := []string{}
	for range 2 {
		conn.Write([]byte("PING\n"))
		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			config.Logger.LogError("Failed to read response from Redis server")
			return err
		}
		responses = append(responses, response)
		config.Logger.LogInfo(fmt.Sprintf("Response from Redis server: %s", response))
	}
	if responses[0] != "+PONG\r\n" || responses[1] != "+PONG\r\n" {
		config.Logger.LogError("Redis server did not respond with PONG")
		return fmt.Errorf("redis server did not respond with PONG")
	}
	config.Logger.LogSuccess("Redis server is responding to multiple PINGs")
	return nil
}

func TestConcurrentPingsStep(config *testrunners.ServerTestConfig) error {
	config.Logger.LogTitle("Test Concurrent Pings")
	config.Logger.LogInfo("Testing if the Redis server is responding to concurrent PINGs")
	time.Sleep(1000 * time.Millisecond)
	defer config.Server.Stop()

	wg := sync.WaitGroup{}
	responses := []string{}
	respMutex := sync.Mutex{}
	wg.Add(2)
	for range 2 {
		go func() {
			defer wg.Done()
			conn, err := net.Dial("tcp", "localhost:6379")
			if err != nil {
				config.Logger.LogError("Failed to connect to Redis port 6379")
				return
			}
			defer conn.Close()
			conn.Write([]byte("PING\n"))
			response, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				config.Logger.LogError("Failed to read response from Redis server")
				return
			}
			respMutex.Lock()
			responses = append(responses, response)
			respMutex.Unlock()
			config.Logger.LogInfo(fmt.Sprintf("Response from Redis server: %s", response))
		}()
	}
	wg.Wait()
	respMutex.Lock()
	defer respMutex.Unlock()
	if len(responses) != 2 || responses[0] != "+PONG\r\n" || responses[1] != "+PONG\r\n" {
		config.Logger.LogError("Redis server did not respond with PONG")
		return fmt.Errorf("redis server did not respond with PONG")
	}
	config.Logger.LogSuccess("Redis server is responding to concurrent PINGs")
	return nil
}
