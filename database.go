package main

import (
	//"fmt"

)

type Database struct {
	likes map[int64]map[string]bool
}

func  NewDatabase() *Database{
	var database Database
	database.likes = make(map[int64]map[string]bool)
	return &database
}

func (database *Database) like(chatID int64, cocktailID string) {
	if(database.likes[chatID] == nil){
		database.likes[chatID] = make(map[string]bool)
	}
	database.likes[chatID][cocktailID] = true
}

func (database Database) isLike(chatID int64, cocktailID string) bool {
	_, ok := database.likes[chatID][cocktailID]
	return !ok
}

func (database Database) getRangeOfLikes(chatID int64) *[]string{
	var result []string
	for key, _ := range(database.likes[chatID]){
		result = append(result, key)
	}
	return &result
} 
