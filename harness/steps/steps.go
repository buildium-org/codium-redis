package steps

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	testrunners "github.com/buildium-org/codium-harness/testrunners"
)

func readResponse(r *bufio.Reader) (string, error) {
	line, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}
	if len(line) == 0 || line[0] != '$' {
		return line, nil
	}

	length, err := strconv.Atoi(strings.TrimSpace(line[1:]))
	if err != nil {
		return "", err
	}
	if length < 0 {
		return line, nil
	}

	data := make([]byte, length+2)
	if _, err := io.ReadFull(r, data); err != nil {
		return "", err
	}
	return line + string(data), nil
}

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
	conn.Write([]byte("+PING\r\n"))
	response, err := readResponse(bufio.NewReader(conn))
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
		conn.Write([]byte("+PING\r\n"))
		response, err := readResponse(bufio.NewReader(conn))
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
			conn.Write([]byte("+PING\r\n"))
			response, err := readResponse(bufio.NewReader(conn))
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

func TestEchoStep(config *testrunners.ServerTestConfig) error {
	config.Logger.LogTitle("Test Echo")
	config.Logger.LogInfo("Testing if the Redis server is responding to ECHO")
	time.Sleep(1000 * time.Millisecond)
	defer config.Server.Stop()

	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		config.Logger.LogError("Failed to connect to Redis port 6379")
		return err
	}
	defer conn.Close()
	conn.Write([]byte("*2\r\n$4\r\nECHO\r\n$3\r\nhey\r\n"))
	response, err := readResponse(bufio.NewReader(conn))
	if err != nil {
		config.Logger.LogError("Failed to read response from Redis server")
		return err
	}
	config.Logger.LogInfo(fmt.Sprintf("Response from Redis server: %s", response))
	if response != "$3\r\nhey\r\n" {
		config.Logger.LogError("Redis server did not respond with ECHO")
		return fmt.Errorf("redis server did not respond with ECHO")
	}
	config.Logger.LogSuccess("Redis server is responding to ECHO")
	return nil
}

func TestSetAndGetStep(config *testrunners.ServerTestConfig) error {
	config.Logger.LogTitle("Test Set and Get")
	config.Logger.LogInfo("Testing if the Redis server is responding to SET and GET")
	time.Sleep(1000 * time.Millisecond)
	defer config.Server.Stop()

	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		config.Logger.LogError("Failed to connect to Redis port 6379")
		return err
	}
	defer conn.Close()
	config.Logger.LogInfo("Setting key1 to value1")
	conn.Write([]byte("*3\r\n$3\r\nSET\r\n$4\r\nkey1\r\n$6\r\nvalue1\r\n"))
	response, err := readResponse(bufio.NewReader(conn))
	if err != nil {
		config.Logger.LogError("Failed to read response from Redis server")
		return err
	}
	config.Logger.LogInfo(fmt.Sprintf("Response from Redis server: %s", response))
	if response != "+OK\r\n" {
		config.Logger.LogError("Redis server did not respond with OK")
		return fmt.Errorf("redis server did not respond with OK")
	}
	conn.Write([]byte("*2\r\n$3\r\nGET\r\n$4\r\nkey1\r\n"))
	response, err = readResponse(bufio.NewReader(conn))
	if err != nil {
		config.Logger.LogError("Failed to read response from Redis server")
		return err
	}
	config.Logger.LogInfo(fmt.Sprintf("Response from Redis server: %s", response))
	if response != "$6\r\nvalue1\r\n" {
		config.Logger.LogError("Redis server did not respond with value1")
		return fmt.Errorf("redis server did not respond with value1")
	}
	config.Logger.LogSuccess("Redis server is responding to SET and GET")
	return nil
}

func TestSetAndGetWithExpireTimeStep(config *testrunners.ServerTestConfig) error {
	config.Logger.LogTitle("Test Set and Get with Expire Time")
	config.Logger.LogInfo("Testing if the Redis server is responding to SET and GET with expire time")
	time.Sleep(1000 * time.Millisecond)
	defer config.Server.Stop()

	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		config.Logger.LogError("Failed to connect to Redis port 6379")
		return err
	}
	defer conn.Close()
	config.Logger.LogInfo("Setting key1 to value1 with expire time 1000ms")
	conn.Write([]byte("*5\r\n$3\r\nSET\r\n$4\r\nkey1\r\n$6\r\nvalue1\r\n$2\r\nPX\r\n:1000\r\n"))
	response, err := readResponse(bufio.NewReader(conn))
	if err != nil {
		config.Logger.LogError("Failed to read response from Redis server")
		return err
	}
	config.Logger.LogInfo(fmt.Sprintf("Response from Redis server: %s", response))
	if response != "+OK\r\n" {
		config.Logger.LogError("Redis server did not respond with OK")
		return fmt.Errorf("redis server did not respond with OK")
	}
	conn.Write([]byte("*2\r\n$3\r\nGET\r\n$4\r\nkey1\r\n"))
	response, err = readResponse(bufio.NewReader(conn))
	if err != nil {
		config.Logger.LogError("Failed to read response from Redis server")
		return err
	}
	config.Logger.LogInfo(fmt.Sprintf("Response from Redis server: %s", response))
	if response != "$6\r\nvalue1\r\n" {
		config.Logger.LogError("Redis server did not respond with value1")
		return fmt.Errorf("redis server did not respond with value1")
	}
	time.Sleep(1000 * time.Millisecond)
	conn.Write([]byte("*2\r\n$3\r\nGET\r\n$4\r\nkey1\r\n"))
	response, err = readResponse(bufio.NewReader(conn))
	if err != nil {
		config.Logger.LogError("Failed to read response from Redis server")
		return err
	}
	config.Logger.LogInfo(fmt.Sprintf("Response from Redis server: %s", response))
	if response != "$-1\r\n" {
		config.Logger.LogError("Redis server did not respond with -1")
		return fmt.Errorf("redis server did not respond with -1")
	}
	config.Logger.LogSuccess("Redis server is responding to SET and GET with expire time")
	return nil
}
