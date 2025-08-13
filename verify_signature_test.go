package kutils_test

import (
	"github.com/KyberNetwork/kutils"
	"github.com/stretchr/testify/require"
	"regexp"
	"testing"
	"time"
)

func TestVerifySignature(t *testing.T) {
	// Common test regexp and duration
	tenMinutes := 10 * time.Minute

	type args struct {
		eip191            kutils.EIP191
		authMessageRegexp *regexp.Regexp
		authExpiry        time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid signature and message",
			args: args{
				// Replace with actual valid signature, message, and address
				eip191: kutils.EIP191{
					Signature: "0x308efb5a82c888da39c77c1d946bb85ace75b7981bee631acda11d1d5afd4b2d1ae49c3325163d447fa4a88010f5d7bc13175f997530e187d5f626a34fdacca61c",
					Msg:       "Click sign to add favorite tokens at Kyberswap.com without logging in.\nThis request won’t trigger any blockchain transaction or cost any gas fee. Expires in 7 days. \n\nIssued at: 2024-11-13T07:22:42.790Z",
					Address:   "0x63FaC9201494f0bd17B9892B9fae4d52fe3BD377",
				},
				authMessageRegexp: kutils.DefaultAuthRegexp,
				authExpiry:        10000 * time.Hour,
			},
			wantErr: false,
		},
		{
			name: "Invalid signature format",
			args: args{
				eip191: kutils.EIP191{
					Signature: "invalid-sig",
					Msg:       "Signing this message at: 2024-03-20T10:00:00Z",
					Address:   "0xabcd...",
				},
				authMessageRegexp: kutils.DefaultAuthRegexp,
				authExpiry:        tenMinutes,
			},
			wantErr: true,
		},
		{
			name: "Invalid last byte",
			args: args{
				// Signature with invalid last byte
				eip191: kutils.EIP191{
					Signature: "0x12345...",
					Msg:       "Signing this message at: 2024-03-20T10:00:00Z",
					Address:   "0xabcd...",
				},
				authMessageRegexp: kutils.DefaultAuthRegexp,
				authExpiry:        tenMinutes,
			},
			wantErr: true,
		},
		{
			name: "Mismatched addresses",
			args: args{
				eip191: kutils.EIP191{
					Signature: "0x12345...",
					Msg:       "Signing this message at: 2024-03-20T10:00:00Z",
					Address:   "0xdifferentAddress...",
				},
				authMessageRegexp: kutils.DefaultAuthRegexp,
				authExpiry:        tenMinutes,
			},
			wantErr: true,
		},
		{
			name: "Expired message",
			args: args{
				eip191: kutils.EIP191{
					Signature: "0x4586bdcd4fa68e248ab1e5383c85e5d68550a88b500f96eefa5c65a56584bd003af5ea34f5127ec66fa461dd47b67864aa4b92bb6104e665b1e59a2446ef52f701",
					Msg:       "Click sign to add favorite tokens at Kyberswap.com without logging in.\nThis request won’t trigger any blockchain transaction or cost any gas fee. Expires in 7 days. \n\nIssued at: 2024-11-01T07:22:42.790Z",
					Address:   "0x63FaC9201494f0bd17B9892B9fae4d52fe3BD377",
				},
				authMessageRegexp: kutils.DefaultAuthRegexp,
				authExpiry:        tenMinutes,
			},
			wantErr: true,
		},
		{
			name: "Invalid message format",
			args: args{
				eip191: kutils.EIP191{
					Signature: "0x12345...",
					Msg:       "Invalid message format",
					Address:   "0xabcd...",
				},
				authMessageRegexp: kutils.DefaultAuthRegexp,
				authExpiry:        tenMinutes,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			if tt.name == "Expired message" {
				err = kutils.VerifyEIP191SignatureWithDefaults(&tt.args.eip191)
			} else {
				err = kutils.VerifyEIP191Signature(&tt.args.eip191, tt.args.authMessageRegexp, tt.args.authExpiry)
			}
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
