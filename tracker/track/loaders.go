package track

import (
	"bytes"
	"io"
	"io/fs"
	"io/ioutil"
	"net/http"
	"player/tracker/track/parser"
)

// LoadData - Load track from bytes
func (t *Track) LoadData(data []byte, parser parser.Interface) error {
	units, err := parser.Unmarshal(data)
	if err != nil {
		return err
	}
	t.Units = units
	return nil
}

// ToData - Save track as a bytes' array
func (t *Track) ToData(parser parser.Interface) ([]byte, error) {
	data, err := parser.Marshal(t.Units)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// LoadStream - Load track from stream bytes
func (t *Track) LoadStream(stream io.Reader, parser parser.Interface) error {
	data, err := ioutil.ReadAll(stream)
	if err != nil {
		return err
	}
	return t.LoadData(data, parser)
}

// ToStream - Write bytes' array data to stream
func (t *Track) ToStream(stream io.Writer, parser parser.Interface) error {
	data, err := t.ToData(parser)
	if err != nil {
		return err
	}
	if _, err := io.Copy(stream, bytes.NewReader(data)); err != nil {
		return err
	}
	return nil
}

// LoadFile - Load track from file by path
func (t *Track) LoadFile(filename string, parser parser.Interface) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	if err := t.LoadData(data, parser); err != nil {
		return err
	}
	return nil
}

// ToFile - Save track's data to file.
// It will create file if doesn't exist or replace existing.
func (t *Track) ToFile(filename string, parser parser.Interface) error {
	data, err := t.ToData(parser)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(filename, data, fs.ModePerm); err != nil {
		return err
	}
	return nil
}

// LoadURL - Load track from URL
func (t *Track) LoadURL(url string, parser parser.Interface) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return t.LoadData(data, parser)
}