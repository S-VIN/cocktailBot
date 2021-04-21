package main

type Database struct {
	likes map[int64]Set
}

func NewDatabase() *Database {
	var database Database
	database.likes = make(map[int64]Set)
	return &database
}

func (database *Database) like(chatID int64, cocktailID string) {
	_, ok := database.likes[chatID]
	if !ok {
		database.likes[chatID] = *NewSet()
	}
	temp := database.likes[chatID]
	temp.Add(cocktailID)
}

func (database Database) isLike(chatID int64, cocktailID string) bool {
	return database.likes[chatID].Find(cocktailID)
}

func (database Database) getRangeOfLikes(chatID int64) []string {
	var result []string
	for i := 0; i < database.likes[chatID].GetSize(); i++ {
		value, _ := database.likes[chatID].GetByIndex(i)
		result = append(result, value)
	}
	return result
}

func (database Database) getLikedByIndex(chatID int64, index int) string {
	res, _ := database.likes[chatID].GetByIndex(index)
	return res
}
