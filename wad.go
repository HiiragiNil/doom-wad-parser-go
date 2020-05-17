package wasdparser

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
)

// WAD main struct
type WAD struct {
	Header      *Header
	Directories []*Directory
}

// Header WAD Header...
type Header struct {
	WADType         [4]byte // WAD type. "IWAD" or "PWAD"
	DirectoryCount  uint32  // The number of entries in the directory.
	DirectoryOffset uint32  // Offset in bytes to the directory in the WAD file.
}

// Directory WAD directory struct
type Directory struct {
	Offset uint32  // offset value to the start of the lump data in the WAD file.
	Size   uint32  // The size of the lump in bytes.
	Name   [9]byte // ACSII holding the name of the lump
}

// LoadWAD loads WAD file.
func LoadWAD(wadpath string) (*WAD, error) {
	// Open wad file...
	file, err := os.Open(wadpath)
	if err != nil {
		return nil, fmt.Errorf("Failed to open wad file")
	}
	defer file.Close()

	// Initialize WAD struct
	wad := &WAD{
		Header:      &Header{},
		Directories: make([]*Directory, 0),
	}

	// Initialize file reader
	br := bufio.NewReader(file)

	if err := binary.Read(br, binary.LittleEndian, wad.Header); err != nil {
		return nil, fmt.Errorf("Failed to read WAD Header: %v", err)
	}

	for i := uint32(0); i < wad.Header.DirectoryCount; i++ {
		// TODO
	}

	return wad, nil
}
