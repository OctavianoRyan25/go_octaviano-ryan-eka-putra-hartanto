#!/bin/bash

# Membuat tiga folder: foods, drinks, dan snacks
mkdir foods drinks snacks

# Membuat file menu.txt di dalam folder foods, drinks, dan snacks
touch foods/menu.txt
touch drinks/menu.txt
touch snacks/menu.txt

# Menambahkan item ke dalam file menu.txt
echo "nasi goreng, ayam goreng, bubur ayam" >> foods/menu.txt
echo "kopi susu, susu oat, es teh" >> drinks/menu.txt
curl -s  https://gist.githubusercontent.com/nadirbslmh/c84ee3527272e52a8e6257d46e627f91/raw/74593cc3fe61ca4ff36e11ed8e3fffcfb1d0c602/snacks.json > snacks/menu.txt