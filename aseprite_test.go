package aseprite_test

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/andygeiss/aseprite-go"
)

func assertThat(t *testing.T, value, expected interface{}) {
	if fmt.Sprintf("%v", value) != fmt.Sprintf("%v", expected) {
		t.Fatalf("value should be %v, but got %v", expected, value)
	}
}

func Test_DecodePath_Should_Return_Not_Nil_Given_Valid_File(t *testing.T) {
	path := filepath.Join("testdata", "file_valid.json")
	result, err := aseprite.DecodePath(path)
	assertThat(t, err, nil)
	assertThat(t, result != nil, true)
}

func Test_DecodePath_Should_Return_An_Error_Given_Invalid_File(t *testing.T) {
	path := filepath.Join("testdata", "file_not_exists.json")
	_, err := aseprite.DecodePath(path)
	assertThat(t, err != nil, true)
}

func Test_DecodePath_Should_Return_An_Error_Given_Invalid_JSON(t *testing.T) {
	path := filepath.Join("testdata", "file_with_invalid.json")
	_, err := aseprite.DecodePath(path)
	assertThat(t, err != nil, true)
}

func Test_FrameCount_Should_Return_6_Given_Example_File(t *testing.T) {
	path := filepath.Join("testdata", "level0", "ball-bounce.json")
	data, _ := aseprite.DecodePath(path)
	count := aseprite.FrameCount(data)
	assertThat(t, count, 9)
}

func Test_FrameAt_Should_Return_Frame_1_Given_Example_File_With_Valid_Index(t *testing.T) {
	path := filepath.Join("testdata", "level0", "ball-bounce.json")
	data, _ := aseprite.DecodePath(path)
	frame := aseprite.FrameAt(1, data)
	assertThat(t, frame != nil, true)
	if frame != nil {
		assertThat(t, frame.DurationMs, 100)
		assertThat(t, frame.DurationMs, 100)
		assertThat(t, frame.PosX, 32)
		assertThat(t, frame.PosY, 0)
		assertThat(t, frame.SizeX, 32)
		assertThat(t, frame.SizeY, 32)
	}
}

func Test_FrameAt_Should_Return_Nil_Given_Example_File_With_Invalid_Index(t *testing.T) {
	path := filepath.Join("testdata", "level0", "ball-bounce.json")
	data, _ := aseprite.DecodePath(path)
	frame := aseprite.FrameAt(10, data)
	assertThat(t, frame, nil)
}

func Test_Frames_Should_Return_3_Frames_Given_Example_File(t *testing.T) {
	path := filepath.Join("testdata", "level0", "ball-bounce.json")
	data, _ := aseprite.DecodePath(path)
	frames := aseprite.Frames(data)
	assertThat(t, len(frames), 9)
}

func Test_DecodePath_Should_Return_PNG_Given_Example_File(t *testing.T) {
	path := filepath.Join("testdata", "level0", "ball-bounce.json")
	data, _ := aseprite.DecodePath(path)
	imageName := aseprite.ImageName(data)
	assertThat(t, imageName, "ball-bounce.png")
}

func Test_LoadSpritesheet_Should_Return_2_Sprites_Given_Project_Path(t *testing.T) {
	path := filepath.Join("testdata")
	data, err := aseprite.LoadSpritesheet(path)
	assertThat(t, err, nil)
	assertThat(t, len(data), 2)
	assertThat(t, len(data[filepath.Join("level0", "background-static")]), 1)
	assertThat(t, len(data[filepath.Join("level0", "ball-bounce")]), 9)
}

func Test_LoadSpritesheet_Should_Return_8_Sprites_Given_Project_Path(t *testing.T) {
	path := filepath.Join("sprites")
	data, err := aseprite.LoadSpritesheet(path)
	assertThat(t, err, nil)
	assertThat(t, len(data), 8)
	assertThat(t, len(data["player-idle-down"]), 4)
}
