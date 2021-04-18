package main

import ()

type Database struct {
	likes map[int64]string
}

func (database Database) like(chatID int64, cocktailID string) {
	database.likes[chatID] = cocktailID
}

func (database Database) isLike(chatID int64, cocktailID string) bool {
	_, ok := database.likes[chatID]
	return ok
}

func (database Database) getRangeOfLikes(chatID int64) *[]string{
	var result []string
	for _, value := range(database.likes){
		result = append(result, value)
	}
	return &result
} 
