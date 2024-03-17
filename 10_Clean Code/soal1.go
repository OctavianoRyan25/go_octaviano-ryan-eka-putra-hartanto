package main

import "fmt"

type User struct { // Strurt menggunakan huruf besar agar dapat diakses package lain
	Email    string
	Password string
}

type UserRepo struct {
	DB []User
}

func (r UserRepo) Register(u User) {
	if u.Email == "" || u.Password == "" {
		fmt.Println("register failed") // seharusnya pesannya lebih spesifik seperti "email tidak boleh kosong" atau "password tidak boleh kosong"
	}

	r.DB = append(r.DB, u)
}

func (r UserRepo) Login(u User) string {
	if u.Email == "" || u.Password == "" {
		fmt.Println("login failed") // seharusnya pesannya lebih spesifik seperti "email tidak boleh kosong" atau "password tidak boleh kosong"
	}

	//Belum ada kondisi dimana email dan password tidak ada di DB(salah satu atau keduanya)

	for _, us := range r.DB {
		if us.Email == u.Email && us.Password == u.Password {
			return "auth token"
		}
	}

	return ""
}
