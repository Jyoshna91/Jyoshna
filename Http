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

	// Construct nping command for HTTPS traffic
	httpsCommand := fmt.Sprintf("sudo nping --tcp -c 10 --rate 100 --source-ip %s --data-string 'GET / HTTP/1.1\nHost: example.com\n\n' --dest-port 443 %s", sourceIP, destinationIP)

	// Execute nping command for HTTPS traffic
	httpsCmd := exec.Command("sh", "-c", httpsCommand)

	httpsOutput, err := httpsCmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Error starting nping HTTPS command: %v. Output: %s", err, httpsOutput)
	}

	t.Logf("nping HTTPS output: %s", string(httpsOutput))

	// Wait for the specified duration
	time.Sleep(duration)

	// Stop nping command
	if err := httpsCmd.Process.Kill(); err != nil {
		t.Fatalf("Error stopping nping HTTPS command: %v", err)
	}

	t.Log("nping HTTPS traffic generation completed.")
}
