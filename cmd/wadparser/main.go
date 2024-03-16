package main

import (
	"bytes"
	"log"
)

func main() {
	reader, err := readBytesFromFile("file")
	if err != nil {
		log.Fatal(err)
	}

	wadType, err := extractWadType(reader)
	if err != nil {
		log.Fatal(err)
	}

	_, err = extractNumLumps(reader)
	if err != nil {
		log.Fatal(err)
	}

	data, err := extractData(reader)
	if err != nil {
		log.Fatal(err)
	}

	lumps, err := extractLumps(reader, data)
	if err != nil {
		log.Fatal(err)
	}

	levels, err := extractLevels(lumps)
	if err != nil {
		log.Fatal(err)
	}

	wad := Wad{
		WadType: wadType,
		Levels:  levels,
	}

	log.Println(wad)
}

func readBytesFromFile(filename string) (*bytes.Reader, error) {
	return nil, nil
}

func extractWadType(reader *bytes.Reader) (string, error) {
	return "", nil
}

func extractNumLumps(reader *bytes.Reader) (string, error) {
	return "", nil
}

func extractData(reader *bytes.Reader) ([]byte, error) {
	return nil, nil
}

func extractLumps(reader *bytes.Reader, data []byte) ([]Lump, error) {
	return nil, nil
}

func extractLevels(lumps []Lump) ([]Level, error) {
	return nil, nil
}
