package nping_test

import (
	"testing"
	"time"

	"github.com/openconfig/ondatra"
	"github.com/openconfig/ondatra/nping"
)

func TestSendUDPTraffic(t *testing.T) {
	// Set up the test environment
	ate := ondatra.ATE(t, "ate")
	srcRouter := ate.Device(t, "source_router")
	dstRouter := ate.Device(t, "destination_router")

	// Define parameters for packet generation
	srcIP := "192.0.2.1"
	dstIP := "198.51.100.2"
	udpPort := 5000

	// Set up nping configuration
	cfg := nping.NewConfig()
	cfg.SetSourceDevice(srcRouter.ID()).
		SetDestinationDevice(dstRouter.ID())

	// Send UDP packets
	udpFlow := cfg.Flows().Add().
		SetName("udp_flow").
		SetProtocol(nping.UDP).
		SetSourceIP(srcIP).
		SetDestinationIP(dstIP).
		SetDestinationPort(udpPort)
	udpFlow.SetRate(nping.PacketsPerSecond(100))

	// Push config and start traffic
	ate.Nping().PushConfig(t, cfg)
	ate.Nping().StartTraffic(t)

	// Wait for traffic to stabilize
	time.Sleep(30 * time.Second)

	// Stop traffic
	ate.Nping().StopTraffic(t)

	// Clean up
	ate.Nping().RemoveConfig(t)
}
