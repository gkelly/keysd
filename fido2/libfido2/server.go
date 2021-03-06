package libfido2

import (
	"context"
	"sort"

	"github.com/keys-pub/go-libfido2"
	"github.com/keys-pub/keysd/fido2"
)

// Server ...
type Server struct{}

// NewAuthenticatorsServer creates an AuthenticatorsServer.
func NewAuthenticatorsServer() fido2.AuthenticatorsServer {
	return &Server{}
}

// Devices ...
func (s *Server) Devices(ctx context.Context, req *fido2.DevicesRequest) (*fido2.DevicesResponse, error) {
	devices, err := libfido2.DeviceLocations()
	if err != nil {
		return nil, err
	}
	sort.Slice(devices, func(i, j int) bool {
		return devices[i].Product < devices[j].Product
	})
	return &fido2.DevicesResponse{
		Devices: devicesToRPC(devices),
	}, nil
}

// findDevice returns a device from a name.
// You need to call Device.Close() when done.
func findDevice(name string) (*libfido2.Device, error) {
	device, err := libfido2.NewDevice(name)
	if err != nil {
		return nil, err
	}
	return device, nil
}

// DeviceInfo ...
func (s *Server) DeviceInfo(ctx context.Context, req *fido2.DeviceInfoRequest) (*fido2.DeviceInfoResponse, error) {
	device, err := findDevice(req.Device)
	if err != nil {
		return nil, err
	}
	defer device.Close()

	info, err := device.Info()
	if err != nil {
		return nil, err
	}

	return &fido2.DeviceInfoResponse{
		Info: deviceInfoToRPC(info),
	}, nil
}

// MakeCredential ...
func (s *Server) MakeCredential(ctx context.Context, req *fido2.MakeCredentialRequest) (*fido2.MakeCredentialResponse, error) {
	device, err := findDevice(req.Device)
	if err != nil {
		return nil, err
	}
	defer device.Close()

	typ, err := credTypeFromRPC(req.Type)
	if err != nil {
		return nil, err
	}
	extensions, err := extensionsFromRPC(req.Extensions)
	if err != nil {
		return nil, err
	}
	rk, err := optionValueFromRPC(req.RK)
	if err != nil {
		return nil, err
	}
	uv, err := optionValueFromRPC(req.UV)
	if err != nil {
		return nil, err
	}

	attestation, err := device.MakeCredential(
		req.ClientDataHash,
		rpFromRPC(req.RP),
		userFromRPC(req.User),
		typ,
		req.PIN,
		&libfido2.MakeCredentialOpts{
			Extensions: extensions,
			RK:         rk,
			UV:         uv,
		},
	)
	if err != nil {
		return nil, err
	}
	return &fido2.MakeCredentialResponse{
		Attestation: attestationToRPC(attestation),
	}, nil
}

// SetPIN ...
func (s *Server) SetPIN(ctx context.Context, req *fido2.SetPINRequest) (*fido2.SetPINResponse, error) {
	device, err := findDevice(req.Device)
	if err != nil {
		return nil, err
	}
	defer device.Close()

	if err := device.SetPIN(req.PIN, req.OldPIN); err != nil {
		return nil, err
	}

	return &fido2.SetPINResponse{}, nil
}

// Reset ...
func (s *Server) Reset(ctx context.Context, req *fido2.ResetRequest) (*fido2.ResetResponse, error) {
	device, err := findDevice(req.Device)
	if err != nil {
		return nil, err
	}
	defer device.Close()

	if err := device.Reset(); err != nil {
		return nil, err
	}

	return &fido2.ResetResponse{}, nil
}

// RetryCount ...
func (s *Server) RetryCount(ctx context.Context, req *fido2.RetryCountRequest) (*fido2.RetryCountResponse, error) {
	device, err := findDevice(req.Device)
	if err != nil {
		return nil, err
	}
	defer device.Close()

	count, err := device.RetryCount()
	if err != nil {
		return nil, err
	}

	return &fido2.RetryCountResponse{
		Count: int32(count),
	}, nil
}

// Assertion ...
func (s *Server) Assertion(ctx context.Context, req *fido2.AssertionRequest) (*fido2.AssertionResponse, error) {
	device, err := findDevice(req.Device)
	if err != nil {
		return nil, err
	}
	defer device.Close()

	extensions, err := extensionsFromRPC(req.Extensions)
	if err != nil {
		return nil, err
	}
	uv, err := optionValueFromRPC(req.UV)
	if err != nil {
		return nil, err
	}
	up, err := optionValueFromRPC(req.UP)
	if err != nil {
		return nil, err
	}

	assertion, err := device.Assertion(req.RPID, req.ClientDataHash, req.CredID, req.PIN, &libfido2.AssertionOpts{Extensions: extensions, UV: uv, UP: up})
	if err != nil {
		return nil, err
	}

	return &fido2.AssertionResponse{
		Assertion: assertionToRPC(assertion),
	}, nil
}

// CredentialsInfo ...
func (s *Server) CredentialsInfo(ctx context.Context, req *fido2.CredentialsInfoRequest) (*fido2.CredentialsInfoResponse, error) {
	device, err := findDevice(req.Device)
	if err != nil {
		return nil, err
	}
	defer device.Close()

	info, err := device.CredentialsInfo(req.PIN)
	if err != nil {
		return nil, err
	}

	return &fido2.CredentialsInfoResponse{
		Info: credentialsInfoToRPC(info),
	}, nil
}

// Credentials ...
func (s *Server) Credentials(ctx context.Context, req *fido2.CredentialsRequest) (*fido2.CredentialsResponse, error) {
	device, err := findDevice(req.Device)
	if err != nil {
		return nil, err
	}
	defer device.Close()

	out := []*fido2.Credential{}
	if req.RPID == "" {
		rps, err := device.RelyingParties(req.PIN)
		if err != nil {
			return nil, err
		}
		for _, rp := range rps {
			credentials, err := device.Credentials(rp.ID, req.PIN)
			if err != nil {
				return nil, err
			}
			out = append(out, credentialsToRPC(relyingPartyToRPC(rp), credentials)...)
		}
	} else {
		credentials, err := device.Credentials(req.RPID, req.PIN)
		if err != nil {
			return nil, err
		}
		rp := &fido2.RelyingParty{ID: req.RPID} // TODO: Name
		out = credentialsToRPC(rp, credentials)
	}

	return &fido2.CredentialsResponse{
		Credentials: out,
	}, nil
}

// RelyingParties ...
func (s *Server) RelyingParties(ctx context.Context, req *fido2.RelyingPartiesRequest) (*fido2.RelyingPartiesResponse, error) {
	device, err := findDevice(req.Device)
	if err != nil {
		return nil, err
	}
	defer device.Close()

	rps, err := device.RelyingParties(req.PIN)
	if err != nil {
		return nil, err
	}

	return &fido2.RelyingPartiesResponse{
		Parties: relyingPartiesToRPC(rps),
	}, nil
}
