package hlsduration

import (
	"fmt"
	"testing"
)

func TestCalculate(t *testing.T) {
	type args struct {
		masterManifestURI string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{name: "simple", args: args{masterManifestURI: "https://d1hsxynlvbyrp1.cloudfront.net/videos/big_bunny/hls/index.m3u8"}, want: 62.280000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Calculate(tt.args.masterManifestURI)
			fmt.Printf("got: %v\n", got)
			if got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateSimple(t *testing.T) {
	t.Run("simple test", func(t *testing.T) {
		duration := Calculate("https://d1hsxynlvbyrp1.cloudfront.net/videos/big_bunny/hls/index.m3u8")
		fmt.Printf("duration: %v\n", duration)
		if duration != 62.280000 {
			t.Errorf("Calculate() = %v, want %v", duration, 62.280000)
		}
	})
}
