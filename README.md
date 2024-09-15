# DBO TEST
- This project created for test from DBO and build with gonic, postgres, and docker

# ERD
- Please click this link `https://drive.google.com/file/d/1wbxBGevLXqXG_Hu9VLQIccy6PkH0vw10/view?usp=sharing`

# Doc API
- Please click this link https://documenter.getpostman.com/view/8144274/2sAXqp8PEd

# Requirment
- Docker

# Set up with docker
- Create file .env in /dbo-test
  `touch .env`
- Copy env variable needed, example : <br>
  `SECRET_JWT=<your stein jwt secret>`<br>
  `STEIN_HQ_URL = <your stein hq url>`<br>
  `STEIN_HQ_USERNAME = <your stein hq username>`<br>
  `STEIN_HQ_PASSWORD = <your stein hq password>`<br>
  `CURRENCY_API_KEY = <your currency api key>`
- Build image and run container
  `docker-compose up -d`

# Stop service
- Run command
  `docker-compose down`