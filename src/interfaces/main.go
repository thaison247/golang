package main

import "fmt"

type User struct {
	name     string
	email    string
	notifier UserNotifier
}

type UserNotifier interface {
	sendMessage(user *User, msg string)
}

type SmsNotifier struct {
}

type EmailNotifier struct {
}

func main() {
	user1 := User{"Son", "nguyenthaison247@gmail.com", SmsNotifier{}}
	user2 := User{"Hoang", "nguyenduyhoang247@gmail.com", EmailNotifier{}}

	user1.notifier.sendMessage(&user2, "Please check your sms")
	user2.notifier.sendMessage(&user1, "Please check your email")

}

func (notifier SmsNotifier) sendMessage(user *User, msg string) {
	fmt.Printf("Send SMS to %s, message: %s\n", user.name, msg)
}

func (notifier EmailNotifier) sendMessage(user *User, msg string) {
	fmt.Printf("Send EMAIL to %s, message: %s\n", user.name, msg)
}
