package aseprite_test

import (
	"embed"
	"path/filepath"
	"testing"

	"github.com/andygeiss/aseprite-go"
	"github.com/andygeiss/utils/assert"
)

func Test_DecodePath_Should_Return_Not_Nil_Given_Valid_File(t *testing.T) {
	path := filepath.Join("testdata", "file_valid.json")
	result, err := aseprite.DecodePath(path)
	assert.That("err should be nil", t, err, nil)
	assert.That("result should not be nil", t, result != nil, true)
}

func Test_DecodePath_Should_Return_An_Error_Given_Invalid_File(t *testing.T) {
	path := filepath.Join("testdata", "file_not_exists.json")
	_, err := aseprite.DecodePath(path)
	assert.That("err should not be nil", t, err != nil, true)
}

func Test_DecodePath_Should_Return_An_Error_Given_Invalid_JSON(t *testing.T) {
	path := filepath.Join("testdata", "file_with_invalid.json")
	_, err := aseprite.DecodePath(path)
	assert.That("err should not be nil", t, err != nil, true)
}

func Test_FrameCount_Should_Return_6_Given_Example_File(t *testing.T) {
	path := filepath.Join("testdata", "level0", "ball-bounce.json")
	data, _ := aseprite.DecodePath(path)
	count := aseprite.FrameCount(data)
	assert.That("count should be 9", t, count, 9)
}

func Test_FrameAt_Should_Return_Frame_1_Given_Example_File_With_Valid_Index(t *testing.T) {
	path := filepath.Join("testdata", "level0", "ball-bounce.json")
	data, _ := aseprite.DecodePath(path)
	frame := aseprite.FrameAt(1, data)
	assert.That("frame should not be nil", t, frame != nil, true)
	if frame != nil {
		assert.That("DurationMs should be 100", t, frame.DurationMs, 100)
		assert.That("PosX should be 32", t, frame.PosX, 32)
		assert.That("PosY should be 0", t, frame.PosY, 0)
		assert.That("SizeX should be 32", t, frame.SizeX, 32)
		assert.That("SizeY should be 32", t, frame.SizeY, 32)
	}
}

func Test_FrameAt_Should_Return_Nil_Given_Example_File_With_Invalid_Index(t *testing.T) {
	path := filepath.Join("testdata", "level0", "ball-bounce.json")
	data, _ := aseprite.DecodePath(path)
	frame := aseprite.FrameAt(10, data)
	assert.That("frame should be nil", t, frame == nil, true)
}

func Test_Frames_Should_Return_3_Frames_Given_Example_File(t *testing.T) {
	path := filepath.Join("testdata", "level0", "ball-bounce.json")
	data, _ := aseprite.DecodePath(path)
	frames := aseprite.Frames(data)
	assert.That("frame count should be 9", t, len(frames), 9)
}

func Test_DecodePath_Should_Return_PNG_Given_Example_File(t *testing.T) {
	path := filepath.Join("testdata", "level0", "ball-bounce.json")
	data, _ := aseprite.DecodePath(path)
	imageName := aseprite.ImageName(data)
	assert.That("imageName should be correct", t, imageName, "ball-bounce.png")
}

func Test_LoadSpritesheet_Should_Return_2_Sprites_Given_Project_Path(t *testing.T) {
	path := filepath.Join("testdata")
	data, err := aseprite.LoadSpritesheet(path)
	assert.That("error should be nil", t, err, nil)
	assert.That("data length should be 2", t, len(data), 2)
	assert.That("background-static length should be 1", t, len(data[filepath.Join("level0", "background-static")]), 1)
	assert.That("ball-bounce length should be 9", t, len(data[filepath.Join("level0", "ball-bounce")]), 9)
}

func Test_LoadSpritesheet_Should_Return_8_Sprites_Given_Project_Path(t *testing.T) {
	path := filepath.Join("sprites")
	data, err := aseprite.LoadSpritesheet(path)
	assert.That("error should be nil", t, err, nil)
	assert.That("data length should be 8", t, len(data), 8)
	assert.That("player-idle-down length should be 4", t, len(data["player-idle-down"]), 4)
}

//go:embed sprites
var sprites embed.FS

func Test_LoadSpritesheetEmbed_Should_Return_8_Sprites_Given_Project_Path(t *testing.T) {
	path := filepath.Join("sprites")
	data, err := aseprite.LoadSpritesheetEmbed(path, sprites)
	assert.That("error should be nil", t, err, nil)
	assert.That("data length should be 8", t, len(data), 8)
	assert.That("player-idle-down length should be 4", t, len(data["player-idle-down"]), 4)
}
