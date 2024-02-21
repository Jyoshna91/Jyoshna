package otg

import (
	"fmt"
	"os/exec"
	"testing"
	"time"
)

func TestConnectToRouter(t *testing.T) {
	const (
		sourceIP      = "10.133.35.158"
		destinationIP = "10.133.35.143"
		duration      = 30 * time.Second
	)

	// Construct nping command
	command := fmt.Sprintf("sudo nping --udp -c 10 --rate 100 --source-ip %s %s", sourceIP, destinationIP)

	// Execute nping command
	cmd := exec.Command("sh", "-c", command)

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Error starting nping command: %v", err)
	}

	t.Logf("nping output: %s", string(output))

	// Wait for the specified duration
	time.Sleep(duration)

	// Stop nping command
	err = cmd.Process.Kill()
	if err != nil {
		t.Fatalf("Error stopping nping command: %v", err)
	}

	t.Log("nping traffic generation completed.")
}
