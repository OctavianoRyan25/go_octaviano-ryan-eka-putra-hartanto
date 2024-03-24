#!/bin/bash

# Meminta input dari pengguna
read -p "Masukkan endpoint API: " endpoint
read -p "Masukkan method HTTP (GET/POST/PUT/DELETE): " method
read -p "Masukkan request body: " requestBody
read -p "Masukkan tipe request body (application/json, application/x-www-form-urlencoded, dll.): " contentType

# Mengirim request ke API
response=$(curl -X "$method" -H "Content-Type: $contentType" -d "$requestBody" "$endpoint")

# Menampilkan respons dari API
echo "Respons dari API:"
echo "$response"

sleep 15
