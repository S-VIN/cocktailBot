package main

import (
	"database/sql"
	"errors"
	"strconv"
	_ "github.com/mattn/go-sqlite3"
)



// FAKE DATABASE
type FDatabase struct {
	likes map[int64]Set
}

func NewFDatabase() *FDatabase {
	var database FDatabase
	database.likes = make(map[int64]Set)
	return &database
}

func (database *FDatabase) Like(chatID int64, cocktailID string) error{
	_, ok := database.likes[chatID]
	if !ok {
		database.likes[chatID] = *NewSet()
	}
	temp := database.likes[chatID]
	temp.Add(cocktailID)
	return nil
}

func (database FDatabase) IsLike(chatID int64, cocktailID string) (bool, error) {
	return database.likes[chatID].Find(cocktailID), nil
}

func (database FDatabase) GetRangeOfLikes(chatID int64) ([]string, error) {
	var result []string
	for i := 0; i < database.likes[chatID].GetSize(); i++ {
		value, err := database.likes[chatID].GetByIndex(i)
		if(err != nil){
			return result, err
		}
		result = append(result, value)
	}
	return result, nil
}

func (database FDatabase) GetLikedByIndex(chatID int64, index int) (string, error) {
	res, err := database.likes[chatID].GetByIndex(index)
	return res, err
}

// NORMAL SQLITE DATABASE
type Database struct {
	db *sql.DB
}

func (database *Database) GetRangeOfLikes(chatID int64) (result []string, err error) {
	database.db, err = sql.Open("sqlite3", "likes.db")
	defer database.db.Close()
	if err != nil{
		return 
	}
	sqlResult, err := database.db.Query("select distinct likesStr from likes where chatID==" + strconv.FormatInt(chatID, 10))
	for sqlResult.Next() {
		result = append(result, "")
		err = sqlResult.Scan(&result[len(result)-1])
		if err != nil {
			return
		}
	}
	return
}

func (database *Database) Like(chatID int64, cocktailID string)(err error) {
	database.db, err = sql.Open("sqlite3", "likes.db")
	defer database.db.Close()
	if err != nil{
		return 
	}
	_, err = database.db.Exec("insert into likes values (" + strconv.FormatInt(chatID, 10) + ", '" + cocktailID + "')")
	return err
}

func (database *Database) IsLike(chatID int64, cocktailID string) (res bool, err error) {
	database.db, err = sql.Open("sqlite3", "likes.db")
	defer database.db.Close()
	if err != nil{
		return 
	}
	sqlResult, err := database.db.Query("select distinct likesStr from likes where chatID==" + strconv.FormatInt(chatID, 10) + " and likesStr==" + "'" + cocktailID + "'")
	if err != nil {
		return false, err
	}

	var temp string
	if !sqlResult.Next() {
		return false, nil
	}

	err = sqlResult.Scan(&temp)
	if err != nil {
		return false, err
	}

	if temp == cocktailID {
		return true, nil
	}
	return false, nil
}

func (database *Database) GetLikedByIndex(chatID int64, index int) (string, error) {
	res, err := database.GetRangeOfLikes(chatID)
	if(err != nil){
		return "", err
	}
	if(index >= len(res) || index < 0){
		return "", errors.New("index out of range, index = " + strconv.Itoa(index) + ", len = " + strconv.Itoa(len(res)))
	}
	return res[index], nil
}