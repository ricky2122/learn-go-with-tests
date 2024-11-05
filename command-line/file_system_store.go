package poker

import (
	"cmp"
	"encoding/json"
	"fmt"
	"os"
	"slices"
)

type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
}

func FileSystemPlayerStoreFromFile(path string) (*FileSystemPlayerStore, func(), error) {
	db, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0o666)
	if err != nil {
		return nil, nil, fmt.Errorf("problem opening %s %v", path, err)
	}

	closeFunc := func() {
		db.Close()
	}

	store, err := NewFileSystemPlayerStore(db)
	if err != nil {
		return nil, nil, fmt.Errorf("problem creating file system player store, %v", err)
	}

	return store, closeFunc, nil
}

func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	if err := initializePlayerDBFile(file); err != nil {
		return nil, fmt.Errorf("problem initializing player db file, %v", err)
	}

	league, err := NewLeague(file)
	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}

	return &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{file}),
		league:   league,
	}, nil
}

func initializePlayerDBFile(file *os.File) error {
	if _, err := file.Seek(0, 0); err != nil {
		return fmt.Errorf("problem seeking file %s, %v", file.Name(), err)
	}

	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		if _, err := file.Write([]byte("[]")); err != nil {
			return fmt.Errorf("problem writing file %s, %v", file.Name(), err)
		}

		if _, err := file.Seek(0, 0); err != nil {
			return fmt.Errorf("problem seeking file %s, %v", file.Name(), err)
		}
	}

	return nil
}

func (f *FileSystemPlayerStore) GetLeague() League {
	slices.SortStableFunc(f.league, func(a, b Player) int {
		return cmp.Compare(b.Wins, a.Wins)
	})

	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.GetLeague().Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	_ = f.database.Encode(&f.league)
}
