package main

import (
	"io"
	"log"
)

type FileSystemPlayerStore struct {
	database io.ReadSeeker
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	f.database.Seek(0, 0)
	league, err := NewLeague(f.database)
	if err != nil {
		// Handle the error here. For example, you might log it and return an empty slice.
		log.Printf("Error reading league from database: %v", err)
		return []Player{}
	}
	return league
}
