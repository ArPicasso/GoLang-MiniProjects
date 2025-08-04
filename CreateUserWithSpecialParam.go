package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Структура Юзера
type User struct {
	Name           string
	Age            int
	Rating         float32
	RatingLocation *float32
}

// Рандомное создания Юзера
func createUserName() string {
	names := []string{"alice", "kate", "steve", "max"}
	arrayIndex := rand.Intn(4)
	return names[arrayIndex]
}
func createUserAge() int {
	ages := []int{15, 23, 93, 43}
	arrayIndex := rand.Intn(4)
	return ages[arrayIndex]
}

func createUserRating() float32 {
	return rand.Float32() * 10.0
}

func (user *User) addRatingLocation() {
	user.RatingLocation = &user.Rating
	fmt.Println("Указана локация в памяти для Rating")
}

// Вызываем от юзера
func (user User) greeting() {
	fmt.Println("Hallo, ich bin", user.Name)
}

func (user User) sayAge() {
	fmt.Println("Ich bin", user.Age, "Jahre alt")
}

func (user User) sayRating() {
	fmt.Println("Meine Bewertung:", user.Rating)
}

func addRating(user *User) {
	freeRating := rand.Float32() * 5
	fmt.Println("Будем добавлять:", freeRating)
	if user.Rating+freeRating <= 10.0 {
		fmt.Println("Обновляю рейтинг User'a с никнеймом:", user.Name)
		time.Sleep(2 * time.Second)
		user.Rating = user.Rating + 3.0
		user.RatingLocation = &user.Rating
		fmt.Println("Рейтинг User'a с никнеймом:", user.Name, "обновлен.")
	} else {
		fmt.Println("User:", user.Name, "не сможет получить столько рейтинга, его рейтинг уже:", user.Rating)
	}
}

// Основная логика программы
func main() {
	user := User{
		Name:           createUserName(),
		Age:            createUserAge(),
		Rating:         createUserRating(),
		RatingLocation: nil,
	}
	fmt.Println("User rating:", user.Rating)
	user.addRatingLocation()
	addRating(&user)
	fmt.Println("User created.")
	fmt.Println("-----------")
	user.greeting()
	user.sayAge()
	user.sayRating()
	fmt.Println("-----------")

}
