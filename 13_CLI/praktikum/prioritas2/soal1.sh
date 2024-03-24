#!/bin/bash

# Meminta input nama project
read -p "Masukkan nama project (tanpa spasi, boleh mengandung hypen): " projectName

# Meminta input jenis project
read -p "Pilih jenis project (simple/api): " projectType

# Membuat folder baru dengan nama project
mkdir "$projectName"

# Inisialisasi Go module
cd "$projectName" || exit
go mod init "$projectName"

if [ "$projectType" = "simple" ]; then
    # Buat file main.go untuk jenis project simple
    touch main.go
elif [ "$projectType" = "api" ]; then
    # Buat folder dan file untuk jenis project api
    mkdir routes models repositories services configs controllers
    touch main.go
fi

echo "Project berhasil dibuat!"
