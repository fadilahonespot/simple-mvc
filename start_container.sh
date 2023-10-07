#!/bin/bash

# Cek apakah ada kontainer yang berjalan
if [ "$(docker ps -q)" ]; then
  docker rm -f $(docker ps -a -q)
fi

# Selanjutnya, jalankan perintah docker pull dan docker run
docker pull fadilahonespot/simple-mvc:1.1.0
docker run -d -p 8000:8000 --name simple-mvc-apps fadilahonespot/simple-mvc:1.1.0
