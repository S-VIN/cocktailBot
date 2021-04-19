package main

import (
//"fmt"

)

type Database struct {
	likes map[int64]Set
}

func NewDatabase() *Database {
	var database Database
	database.likes = make(map[int64]Set)
	return &database
}

func (database *Database) like(chatID int64, cocktailID string) {
	temp := database.likes[chatID]
	if temp == nil{
		
	}
	temp.Add(cocktailID)
}

func (database Database) isLike(chatID int64, cocktailID string) bool {
	_, ok := database.likes[chatID][cocktailID]
	return !ok
}

func (database Database) getRangeOfLikes(chatID int64) *[]string {
	var result []string
	for key := range database.likes[chatID] {
		result = append(result, key)
	}
	return &result
}

func (database Database) getLikedByIndex(chatID int64, index int) string {
	//hahahaha you need to have set instead of map in likes, loser
	return ""
}