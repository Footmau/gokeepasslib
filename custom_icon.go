package gokeepasslib

import (
	"encoding/base64"
	"encoding/xml"
)

// IconData stores custom icon content
type IconData []byte

// CustomIcon is the structure for custom user icons
type CustomIcon struct {
	XMLName xml.Name `xml:"Icon"`
	UUID    UUID     `xml:"UUID"`
	Data    IconData `xml:"Data"`
}

// MarshalText is a marshaler method to encode icon data as base 64 and return it
func (icon IconData) MarshalText() ([]byte, error) {
	if len(icon) == 0 {
		return make([]byte, 0), nil
	}

	text := make([]byte, base64.StdEncoding.EncodedLen(len(icon)))
	base64.StdEncoding.Encode(text, icon[:])
	return text, nil
}

// UnmarshalText unmarshals a byte slice into a icon data by decoding the given from base64
func (icon *IconData) UnmarshalText(text []byte) error {
	data := make([]byte, base64.StdEncoding.DecodedLen(len(text)))
	_, err := base64.StdEncoding.Decode(data, text)

	if err != nil {
		return err
	}

	*icon = append(data[:0:0], data...)
	return nil
}
