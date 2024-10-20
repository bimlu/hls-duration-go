package hlsduration

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/grafov/m3u8"
)

func Calculate(masterManifestURI string) float64 {
	buffer, err := loadManifest(masterManifestURI)
	if err != nil {
		fmt.Printf("failed to load master manifest: %v", err)
		panic(err)
	}

	masterpl, err := decodeMasterManifest(buffer)
	if err != nil {
		fmt.Printf("failed to decode master manifest: %v", err)
		panic(err)
	}

	mediapls, err := loadMediaManifests(masterpl, masterManifestURI)
	if err != nil {
		fmt.Printf("failed to load media manifests: %v", err)
		panic(err)
	}
	// fmt.Printf("mediaplist: %v\n", mediapls)
	duration := getVideoSequencesDuration(*mediapls[0])
	return duration
}

func getVideoSequencesDuration(mediapl m3u8.MediaPlaylist) float64 {
	duration := 0.0
	for _, seg := range mediapl.Segments {
		if seg != nil {
			duration += seg.Duration
		}
	}
	return duration
}

func loadMediaManifests(masterpl *m3u8.MasterPlaylist, masterManifestURI string) ([]*m3u8.MediaPlaylist, error) {
	var playlists []*m3u8.MediaPlaylist
	mediaManifestURIPrefix, found := strings.CutSuffix(masterManifestURI, "index.m3u8")
	if !found {
		return nil, fmt.Errorf("invalid master manifest URI found")
	}

	for _, variant := range masterpl.Variants {
		buffer, err := loadManifest(mediaManifestURIPrefix + variant.URI)
		if err != nil {
			return nil, fmt.Errorf("failed to load media manifest: %v", err)
		}

		mediapl, err := decodeMediaManifest(buffer)
		if err != nil {
			return nil, fmt.Errorf("failed to decode media manifest: %v", err)
		}
		playlists = append(playlists, mediapl)
	}

	return playlists, nil
}

func decodeMediaManifest(buffer *bytes.Buffer) (*m3u8.MediaPlaylist, error) {
	playlist, listType, err := m3u8.Decode(*buffer, true)
	if err != nil {
		return nil, fmt.Errorf("failed to decode manifest: %v", err)
	}

	if listType != m3u8.MEDIA {
		return nil, fmt.Errorf("the provided manifest is not a media playlist")
	}
	mediapl := playlist.(*m3u8.MediaPlaylist)
	return mediapl, nil
}

func decodeMasterManifest(buffer *bytes.Buffer) (*m3u8.MasterPlaylist, error) {
	playlist, listType, err := m3u8.Decode(*buffer, true)
	if err != nil {
		return nil, fmt.Errorf("failed to decode manifest: %v", err)
	}

	if listType != m3u8.MASTER {
		return nil, fmt.Errorf("the provided manifest is not a master playlist")
	}
	masterpl := playlist.(*m3u8.MasterPlaylist)
	return masterpl, nil
}

func loadManifest(manifestURI string) (*bytes.Buffer, error) {
	resp, err := http.Get(manifestURI)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch VOD manifest from the URI: %v", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}
	buffer := bytes.NewBuffer(data)

	// fmt.Printf("buffer: %v\n", buffer)
	return buffer, nil
}
