package shortlink

import (
	"database/sql"
	"db"
	"errors"
	"fmt"
)

// checks if a id already exists in the database
func checkIDConflict(id string) (bool, error) {
	collision := 1
	err := db.Db.QueryRow("SELECT EXISTS(SELECT 1 FROM links WHERE id=?)", id).Scan(&collision)
	return collision != 0, err
}

// gets a unique id for a new game
func getUniqueID() (string, error) {
	var count uint
	var scale uint
	var addConst uint
	var newID uint
	var idString string

	conflict := true
	err := db.Db.QueryRow("SELECT count, scale, addConst FROM count WHERE type='links'").Scan(&count, &scale, &addConst)
	if err != nil {
		return "", err
	}

	for conflict {
		count += 1
		newID = (count*scale + addConst) % 65536
		idString = fmt.Sprintf("%x", newID)
		// makes 4 wide
		for len(idString) < 4 {
			idString = "0" + idString
		}
		conflict, err = checkIDConflict(idString)
		if err != nil {
			return "", err
		}
	}

	updateCount, err := db.Db.Prepare("UPDATE count SET count=? WHERE type='links'")
	if err != nil {
		return idString, err
	}

	_, err = updateCount.Exec(count)
	if err != nil {
		return idString, err
	}

	return idString, nil
}

// adds the game into the database
func Add(link string) (string, error) {
	err := db.Db.Ping()
	if err != nil {
		return "", err
	}

	idString, err := getUniqueID()
	if err != nil {
		return "", err
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
