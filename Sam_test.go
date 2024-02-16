package otg

import (
	"strings"
	"testing"

	"golang.org/x/net/context"

	"github.com/open-traffic-generator/snappi/gosnappi"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/fakebind"
	"github.com/openconfig/testt"
	"google.golang.org/grpc"
)

type OTG struct {
	ate *fakebind.ATE
}

var (
	fakeSnappi = new(fakeGosnappi)
	fakeATE    = &fakebind.ATE{
		AbstractATE: &binding.AbstractATE{&binding.Dims{
			Ports: map[string]*binding.Port{"port1": {}},
		}},
		DialOTGFn: func(context.Context, ...grpc.DialOption) (gosnappi.Api, error) {
			return fakeSnappi, nil
		},
	}
	otgAPI = &OTG{ate: fakeATE}
)

func (otg *OTG) FetchConfig(t *testing.T) gosnappi.Config {
	wantCfg := gosnappi.NewConfig()
	fakeSnappi.config = wantCfg
	gotCfg := otgAPI.FetchConfig(t)
	if wantCfg != gotCfg {
		t.Errorf("FetchConfig got unexpected config %v, want %v", gotCfg, wantCfg)
	}
	return gotCfg
}

func (otg *OTG) PushConfig(t *testing.T, cfg gosnappi.Config) {
	otg.ate.SetConfig(cfg)
}

func (otg *OTG) StartProtocols(t *testing.T) {
	fakeSnappi.controlState = nil
	otgAPI.StartProtocols(t)
	if got, want := fakeSnappi.controlState.Protocol().All().State(), gosnappi.StateProtocolAllState.START; got != want {
		t.Errorf("StartProtocols got unexpected protocol state %v, want %v", got, want)
	}
}

func (otg *OTG) StopProtocols(t *testing.T) {
	fakeSnappi.controlState = nil
	otgAPI.StopProtocols(t)
	if got, want := fakeSnappi.controlState.Protocol().All().State(), gosnappi.StateProtocolAllState.STOP; got != want {
		t.Errorf("StopProtocols got unexpected protocol state %v, want %v", got, want)
	}
}

func (otg *OTG) StartTraffic(t *testing.T) {
	fakeSnappi.controlState = nil
	otgAPI.StartTraffic(t)
	if got, want := fakeSnappi.controlState.Traffic().FlowTransmit().State(), gosnappi.StateTrafficFlowTransmitState.START; got != want {
		t.Errorf("StartTraffic got unexpected transmit state %v, want %v", got, want)
	}
}

func (otg *OTG) StopTraffic(t *testing.T) {
	fakeSnappi.controlState = nil
	otgAPI.StopTraffic(t)
	if got, want := fakeSnappi.controlState.Traffic().FlowTransmit().State(), gosnappi.StateTrafficFlowTransmitState.STOP; got != want {
		t.Errorf("StopTraffic got unexpected transmit state %v, want %v", got, want)
	}
}

func (otg *OTG) SetControlState(t *testing.T, state gosnappi.ControlState) {
	otgAPI.SetControlState(t, state)
	if got := fakeSnappi.controlState; got != state {
		t.Errorf("SetControlState got unexpected control state %v, want %v", got, state)
	}
}

func (otg *OTG) SetControlAction(t *testing.T, action gosnappi.ControlAction) {
	otgAPI.SetControlAction(t, action)
	if got := fakeSnappi.controlAction; got != action {
		t.Errorf("SetControlAction got unexpected control action %v, want %v", got, action)
	}
}

func (otg *OTG) GetCapture(t *testing.T, request gosnappi.CaptureRequest) {
	otgAPI.GetCapture(t, request)
	if got := fakeSnappi.captureReq; got != request {
		t.Errorf("GetCapture got unexpected request %v, want %v", got, request)
	}
}

type fakeGosnappi struct {
	gosnappi.Api
	config        gosnappi.Config
	controlState  gosnappi.ControlState
	controlAction gosnappi.ControlAction
	captureReq    gosnappi.CaptureRequest
}

func (fg *fakeGosnappi) GetConfig() (gosnappi.Config, error) {
	return fg.config, nil
}

func (fg *fakeGosnappi) SetConfig(cfg gosnappi.Config) (gosnappi.Warning, error) {
	fg.config = cfg
	return gosnappi.NewWarning(), nil
}

func (fg *fakeGosnappi) SetControlState(state gosnappi.ControlState) (gosnappi.Warning, error) {
	fg.controlState = state
	return gosnappi.NewWarning(), nil
}

func (fg *fakeGosnappi) SetControlAction(action gosnappi.ControlAction) (gosnappi.ControlActionResponse, error) {
	fg.controlAction = action
	return gosnappi.NewControlActionResponse(), nil
}

func (fg *fakeGosnappi) GetCapture(request gosnappi.CaptureRequest) ([]byte, error) {
	fg.captureReq = request
	return nil, nil
}
