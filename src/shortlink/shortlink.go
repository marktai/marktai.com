package shortlink

import (
	"database/sql"
	"db"
	"errors"
	"fmt"
)

// checks if a id already exists in the database
func checkIDConflict(id uint) (bool, error) {
	collision := 1
	err := db.Db.QueryRow("SELECT EXISTS(SELECT 1 FROM links WHERE id=?)", id).Scan(&collision)
	return collision != 0, err
}

// gets a unique id for a new game
func getUniqueID() (uint, error) {
	var count uint
	var scale uint
	var addConst uint
	var newID uint

	conflict := true
	err := db.Db.QueryRow("SELECT count, scale, addConst FROM count WHERE type='links'").Scan(&count, &scale, &addConst)
	if err != nil {
		return 0, err
	}

	for conflict {
		count += 1
		newID = (count*scale + addConst) % 65536
		conflict, err = checkIDConflict(newID)
		if err != nil {
			return 0, err
		}
	}

	updateCount, err := db.Db.Prepare("UPDATE count SET count=? WHERE type='links'")
	if err != nil {
		return newID, err
	}

	_, err = updateCount.Exec(count)
	if err != nil {
		return newID, err
	}

	return newID, nil
}

// adds the game into the database
func Add(link string) (string, error) {
	err := db.Db.Ping()
	if err != nil {
		return "", err
	}

	idInt, err := getUniqueID()

	if err != nil {
		return "", err
	}

	idString := fmt.Sprintf("%x", idInt)

	// makes 4 wide
	for len(idString) < 4 {
		idString = " " + idString
	}

	addLink, err := db.Db.Prepare("INSERT INTO links VALUES(?, ?)")
	if err != nil {
		return "", err
	}
	_, err = addLink.Exec(idString, link)
	return idString, err
}

// adds the game into the database
func Get(id string) (string, error) {
	err := db.Db.Ping()
	if err != nil {
		return "", err
	}

	var link string

	//TODO: handle NULLS
	err = db.Db.QueryRow("SELECT link FROM links WHERE id=?", id).Scan(&link)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("Link not found")
		}
		return "", err
	}

	return link, nil
}
