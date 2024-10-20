# hls-duration-go
`hls-duration-go` is a Go library that fetches and calculates the total duration of an HLS (HTTP Live Streaming) VOD (Video on Demand). It loads the master and media playlists, processes the segments, and calculates the cumulative duration.

## Features
- Fetches the HLS master and media playlists.
- Calculates the total duration of a VOD based on the media segments in the playlists.
- Supports multiple media variants from the master playlist.

## Installation
Install the package using go get:

```
go get github.com/bimlu/hls-duration-go
```

## Usage
To use the hls-duration library in your Go project, follow the example below.

### Example
```
package main

import (
	"fmt"
	"github.com/bimlu/hls-duration-go"
)

func main() {
	// Replace the URL with your master manifest URI
	duration := hlsduration.Calculate("https://d1hsxynlvbyrp1.cloudfront.net/videos/big_bunny/hls/index.m3u8")
	fmt.Printf("Total duration of the VOD: %.2f seconds\n", duration)
}
```

### Output
```
Total duration of the VOD: 62.28 seconds
```

## API
`Calculate(masterManifestURI string) float64`
Calculates the total duration of an HLS VOD.
* Parameters:
    * masterManifestURI: The URL of the master manifest .m3u8 file.
* Returns:
    * The total duration of the VOD in seconds.

## Testing
You can run tests using Go’s built-in testing framework. The project includes a basic test case to validate the duration calculation.
```
go test ./...
```

### Example test case
```
package hlsduration

import (
	"fmt"
	"testing"
)

func TestCalculateSimple(t *testing.T) {
	t.Run("simple test", func(t *testing.T) {
		duration := Calculate("https://d1hsxynlvbyrp1.cloudfront.net/videos/big_bunny/hls/index.m3u8")
		fmt.Printf("duration: %v\n", duration)
		if duration != 62.280000 {
			t.Errorf("Calculate() = %v, want %v", duration, 62.280000)
		}
	})
}
```

## License
This project is licensed under the MIT License
