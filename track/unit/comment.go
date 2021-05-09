package unit

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Comment - Comment line
type Comment string

// CommentTagType - Some comments can has tags like this #!TIMING:* and TIMING is a tag.
// Tags can be used to set specific tracker configuration or provide additional information for track file.
// Tags are case insensitive, means #!TIMING:* same as #!timing:*
type CommentTagType string

const (
	// CommentTagNone - Empty tag for comments without tag
	CommentTagNone CommentTagType = ""
	// CommentTagTiming - Sets tracker tick's delay in ms like: #!timing:120.
	// It helpful if you want to make your song more slower or faster.
	// If the app can't catch all of yours notes, you can increase this value to prevent notes throttling.
	// Usual, default value is 100
	CommentTagTiming CommentTagType = "timing"
	// CommentTagShift - Sets notes shifting by tones or octaves.
	// Usually the app automatically detects shifting and it's helpful for usability when you write your songs.
	// But if you want to play only in second octave and your song
	// describes only notes from one octave they will be decreased to first octave.
	// This behavior can be overridden by #!shift:1.
	// This comment forces increase tracker's octave.
	CommentTagShift CommentTagType = "shift"
)

// NewComment - Make new comment line and check it.
// Usual comments start with # (hash) like: # Page 1.
// You can use tagged comments started with #!tag:
// Those comments are using to change tracker behavior or add additional information into your tracks.
func NewComment(str string) (Comment, error) {
	if !strings.HasPrefix(str, "#") {
		return "", fmt.Errorf("comment must has # (hash) prefix: %s", str)
	}
	comment := Comment(str)

	// Check for comment tag
	switch comment.Tag() {
	case CommentTagTiming:
		timing := comment.ValueInt()
		if timing <= 0 {
			return Comment(""), fmt.Errorf("timing tag can't be lower or equal 0: %s", str)
		}
	case CommentTagShift:
		shift := comment.ValueInt()
		if shift < 0 {
			return Comment(""), fmt.Errorf("shift tag can't be lower 0: %s", str)
		}
	}

	return Comment(str), nil
}

// getTagValue - Fetch comment tag and value
func (c Comment) getTagValue() (CommentTagType, string) {
	// Trim hash and whitespaces
	cleanComment := strings.TrimSpace(strings.TrimPrefix(string(c), "#"))
	// Parse tag and value
	matches := regexp.MustCompile(`(?i)$?(?:!([^\s:]*):)?(.*)`).FindStringSubmatch(cleanComment)
	tag := strings.ToLower(strings.TrimSpace(matches[1]))
	val := strings.TrimSpace(matches[2])
	if len(tag) > 0 {
		return CommentTagType(tag), val
	}
	return CommentTagNone, val
}

// Tag - Return comment tag
func (c Comment) Tag() CommentTagType {
	tag, _ := c.getTagValue()
	return tag
}

// Value - Return comment value
func (c Comment) Value() string {
	_, val := c.getTagValue()
	return val
}

// ValueInt - Return comment value as an int
func (c Comment) ValueInt() int {
	val := c.Value()
	i, _ := strconv.Atoi(val)
	return i
}

// ValueBool - Return comment value as an bool
func (c Comment) ValueBool() bool {
	val := c.Value()
	val = strings.ToLower(val)
	switch val {
	case "1", "on", "true", "yes":
		return true
	}
	return false
}
