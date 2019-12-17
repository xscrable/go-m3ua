// Copyright 2018-2019 go-m3ua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package params

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParams(t *testing.T) {
	cases := []struct {
		name       string
		structured *Param
		serialized []byte
	}{
		{
			"AspIdentifier",
			NewAspIdentifier(1),
			[]byte{0x00, 0x11, 0x00, 0x08, 0x00, 0x00, 0x00, 0x01},
		},
		{
			"TrafficModeType",
			NewTrafficModeType(1),
			[]byte{0x00, 0x0b, 0x00, 0x08, 0x00, 0x00, 0x00, 0x01},
		},
		{
			"NetworkAppearance",
			NewNetworkAppearance(1),
			[]byte{0x02, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x01},
		},
		{
			"RoutingContext-single",
			NewRoutingContext(1),
			[]byte{0x00, 0x06, 0x00, 0x08, 0x00, 0x00, 0x00, 0x01},
		},
		{
			"RoutingContext-multiple",
			NewRoutingContext(1, 2, 3),
			[]byte{
				0x00, 0x06, 0x00, 0x10, 0x00, 0x00, 0x00, 0x01,
				0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x03,
			},
		},
		{
			"HeartbeatData",
			NewHeartbeatData([]byte("some information")),
			[]byte{
				0x00, 0x09, 0x00, 0x14, 0x73, 0x6f, 0x6d, 0x65,
				0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
				0x74, 0x69, 0x6f, 0x6e,
			},
		},
		{
			"ErrorCode",
			NewErrorCode(ErrInvalidVersion),
			[]byte{0x00, 0x0c, 0x00, 0x08, 0x00, 0x00, 0x00, 0x01},
		},
		{
			"UserCause",
			NewUserCause(UserIdentityUnknown, SCCP),
			[]byte{0x02, 0x04, 0x00, 0x08, 0x00, 0x01, 0x00, 0x00},
		},
		{
			"Status",
			NewStatus(AsStateActive),
			[]byte{0x00, 0x0d, 0x00, 0x08, 0x00, 0x01, 0x00, 0x03},
		},
		{
			"AffectedPointCode",
			NewAffectedPointCode(1, 2, 3),
			[]byte{
				0x00, 0x12, 0x00, 0x10, 0x00, 0x00, 0x00, 0x01,
				0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x03,
			},
		},
		{
			"ConcernedDestination",
			NewConcernedDestination(1),
			[]byte{0x02, 0x06, 0x00, 0x08, 0x00, 0x00, 0x00, 0x01},
		},
		{
			"CorrelationID",
			NewCorrelationID(1),
			[]byte{0x00, 0x13, 0x00, 0x08, 0x00, 0x00, 0x00, 0x01},
		},
		{
			"InfoString",
			NewInfoString("some information"),
			[]byte{
				0x00, 0x04, 0x00, 0x14, 0x73, 0x6f, 0x6d, 0x65,
				0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
				0x74, 0x69, 0x6f, 0x6e,
			},
		},
		{
			"DiagnosticInformation",
			NewDiagnosticInformation([]byte("some information")),
			[]byte{
				0x00, 0x07, 0x00, 0x14, 0x73, 0x6f, 0x6d, 0x65,
				0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
				0x74, 0x69, 0x6f, 0x6e,
			},
		},
		{
			"CongestionIndications",
			NewCongestionIndications(1),
			[]byte{0x02, 0x05, 0x00, 0x08, 0x00, 0x00, 0x00, 0x01},
		},
		{
			"LocalRoutingKeyIdentifier",
			NewLocalRoutingKeyIdentifier(1),
			[]byte{0x02, 0x0a, 0x00, 0x08, 0x00, 0x00, 0x00, 0x01},
		},
		{
			"DestinationPointCode",
			NewDestinationPointCode(1),
			[]byte{0x02, 0x0b, 0x00, 0x08, 0x00, 0x00, 0x00, 0x01},
		},
		{
			"OriginatingPointCodeList",
			NewOriginatingPointCodeList(1, 2, 3),
			[]byte{
				0x02, 0x0e, 0x00, 0x10, 0x00, 0x00, 0x00, 0x01,
				0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x03,
			},
		},
		{
			"ServiceIndicators",
			NewServiceIndicators(1, 2, 3),
			[]byte{0x02, 0x0c, 0x00, 0x08, 0x01, 0x02, 0x03, 0x00},
		},
		{
			"RegistrationStatus",
			NewRegistrationStatus(1),
			[]byte{0x02, 0x12, 0x00, 0x08, 0x00, 0x00, 0x00, 0x01},
		},
		{
			"DeregistrationStatus",
			NewDeregistrationStatus(1),
			[]byte{0x02, 0x13, 0x00, 0x08, 0x00, 0x00, 0x00, 0x01},
		},
		{
			"Generic",
			NewParam(1, []byte{0xde, 0xad, 0xbe, 0xef}),
			[]byte{0x00, 0x01, 0x00, 0x08, 0xde, 0xad, 0xbe, 0xef},
		},
		{
			"ProtocolData",
			NewProtocolData(
				1, // OriginatingPointCode
				2, // DestinationPointCode
				3, // ServiceIndicator
				1, // NetworkIndicator
				0, // MessagePriority
				1, // SignalingLinkSelection
				[]byte{ // Data
					0xde, 0xad, 0xbe, 0xef,
				},
			),
			[]byte{
				// Param Header
				0x02, 0x10, 0x00, 0x14,
				// OPC
				0x00, 0x00, 0x00, 0x01,
				// DPC
				0x00, 0x00, 0x00, 0x02,
				// SI
				0x03,
				// NI
				0x01,
				// MP
				0x00,
				// SLS
				0x01,
				// Data
				0xde, 0xad, 0xbe, 0xef,
			},
		},
		{
			"RegistrationResult",
			NewRegistrationResult(
				NewRegistrationResultPayload(
					NewLocalRoutingKeyIdentifier(1),
					NewRegistrationStatus(1),
					NewRoutingContext(1),
				),
			),
			[]byte{
				// Param Header
				0x02, 0x08, 0x00, 0x1c,
				// LocalRoutingKeyIdentifier
				0x02, 0x0a, 0x00, 0x08, 0x00, 0x00, 0x00, 0x01,
				// RegistrationStatus
				0x02, 0x12, 0x00, 0x08, 0x00, 0x00, 0x00, 0x01,
				// RoutingContext
				0x00, 0x06, 0x00, 0x08, 0x00, 0x00, 0x00, 0x01,
			},
		},
		{
			"DeregistrationResult",
			NewDeregistrationResult(
				NewDeregResultPayload(
					NewRoutingContext(1),
					NewDeregistrationStatus(1),
				),
			),
			[]byte{
				// Param Header
				0x02, 0x09, 0x00, 0x14,
				// RoutingContext
				0x00, 0x06, 0x00, 0x08, 0x00, 0x00, 0x00, 0x01,
				// DeregistrationStatus
				0x02, 0x13, 0x00, 0x08, 0x00, 0x00, 0x00, 0x01,
			},
		},
	}

	for _, c := range cases {
		t.Run("encode/"+c.name, func(t *testing.T) {
			got, err := c.structured.MarshalBinary()
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(got, c.serialized); diff != "" {
				t.Error(diff)
			}
		})

		t.Run("decode/"+c.name, func(t *testing.T) {
			got, err := Parse(c.serialized)
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(got, c.structured); diff != "" {
				t.Error(diff)
			}
		})
	}

}

func TestParseMultiParams(t *testing.T) {
	cases := []struct {
		name       string
		structured []*Param
		serialized []byte
	}{
		{
			"rc-generic",
			[]*Param{
				NewRoutingContext(1),
				NewParam(1, []byte{0xde, 0xad, 0xbe, 0xef}),
			},
			[]byte{
				// Routing Context
				0x00, 0x06, 0x00, 0x08, 0x00, 0x00, 0x00, 0x01,
				// Something with String
				0x00, 0x01, 0x00, 0x08, 0xde, 0xad, 0xbe, 0xef,
			},
		},
	}

	for _, c := range cases {
		got, err := ParseMultiParams(c.serialized)
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(got, c.structured); diff != "" {
			t.Error(diff)
		}
	}
}

func TestParseMalformed(t *testing.T) {
	cases := []struct {
		data []byte
		err  error
	}{
		{[]byte{0x00}, ErrTooShortToParse},
		{[]byte{0x00, 0x00}, ErrTooShortToParse},
		{[]byte{0x00, 0x00, 0x00}, ErrTooShortToParse},
		{[]byte{0x00, 0x00, 0x00, 0x00}, ErrInvalidLength},
	}

	for _, c := range cases {
		if _, err := Parse(c.data); err != c.err {
			t.Errorf("Parse/unexpected error: got: %v, want: %v", err, c.err)
		}
		if _, err := ParseMultiParams(c.data); err != c.err {
			t.Errorf("ParseMulti/unexpected error: got: %v, want: %v", err, c.err)
		}
	}
}
