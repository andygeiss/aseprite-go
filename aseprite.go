package aseprite

import (
	"embed"
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// Frame stores the data.
type Frame struct {
	DurationMs float64 `json:"duration_ms"`
	PosX       float32 `json:"pos_x"`
	PosY       float32 `json:"pos_y"`
	SizeX      float32 `json:"size_x"`
	SizeY      float32 `json:"size_y"`
}

// Spritesheet is a map of frame slices.
type Spritesheet map[string][]*Frame

// DecodePath ...
func DecodePath(path string) (data map[string]interface{}, err error) {
	// open file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// decode JSON to Go data structure
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return nil, err
	}
	return data, err
}

// FrameAt ...
func FrameAt(index int, data map[string]interface{}) (f *Frame) {
	// select frames
	frames := data["frames"].([]interface{})
	if index >= len(frames) {
		return nil
	}
	// return frame at given index
	frame := frames[index].(map[string]interface{})["frame"].(map[string]interface{})
	duration := frames[index].(map[string]interface{})["duration"].(float64)
	return &Frame{
		DurationMs: duration,
		PosX:       float32(frame["x"].(float64)),
		PosY:       float32(frame["y"].(float64)),
		SizeX:      float32(frame["w"].(float64)),
		SizeY:      float32(frame["h"].(float64)),
	}
}

// FrameCount ...
func FrameCount(data map[string]interface{}) (count int) {
	frames := data["frames"].([]interface{})
	return len(frames)
}

// Frames ...
func Frames(data map[string]interface{}) (frames []*Frame) {
	n := FrameCount(data)
	for i := 0; i < n; i++ {
		frames = append(frames, FrameAt(i, data))
	}
	return
}

// ImageName ...
func ImageName(data map[string]interface{}) string {
	return data["meta"].(map[string]interface{})["image"].(string)
}

// LoadSpritesheet ...
func LoadSpritesheet(path string) (spritesheet Spritesheet, err error) {
	// read directory
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	// create and return the spritesheet
	return createSpritesheetFromEntries(path, entries)
}

// LoadSpritesheetEmbed ...
func LoadSpritesheetEmbed(path string, fs embed.FS) (spritesheet Spritesheet, err error) {
	// read directory
	entries, err := fs.ReadDir(path)
	if err != nil {
		return nil, err
	}
	// create and return the spritesheet
	return createSpritesheetFromEntries(path, entries)
}

func createSpritesheetFromEntries(path string, entries []fs.DirEntry) (spritesheet Spritesheet, err error) {
	// check each entry
	spritesheet = make(Spritesheet)
	for _, entry := range entries {
		//
		name := entry.Name()
		// handle sub directories
		if entry.IsDir() {
			ss, err := LoadSpritesheet(filepath.Join(path, name))
			if err != nil {
				return nil, err
			}
			// add the sprites from the sub directory
			for k, v := range ss {
				spritesheet[filepath.Join(name, k)] = v
			}
		}
		// skip non valid JSON files
		if !strings.Contains(name, "-") || !strings.HasSuffix(name, "json") {
			continue
		}
		// extract the spritesheet
		prefixSuffix := strings.Split(name, ".")
		sheet, err := DecodePath(filepath.Join(path, name))
		if err != nil {
			return nil, err
		}
		spritesheet[prefixSuffix[0]] = Frames(sheet)
	}
	return spritesheet, nil
}
