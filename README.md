# DBO TEST
- This project created for test from DBO and build with gonic, postgres, and docker

# ERD
- Please click this link https://drive.google.com/file/d/1wbxBGevLXqXG_Hu9VLQIccy6PkH0vw10/view?usp=sharing

# Doc API
- Please click this link https://documenter.getpostman.com/view/8144274/2sAXqp8PEd

# Requirment
- Docker

# Set up with docker
- Create file .env in /dbo-test
  `touch .env`
- Copy env variable needed, example : <br>
  `SECRET_JWT=<your jwt secret>`<br>
  `POSTGRES_DSN=<your database DSN>`<br>
  `POSTGRES_MAX_IDLE_CONNS=<max idle connection database>`<br>
  `POSTGRES_OPEN_CONNS=<open connection database>`<br>
  `PORT=<port apps>`<br>
  `ALLOW_CORS=<domain allow cors>`

- Build image and run container
  `docker-compose up -d`

# Stop service
- Run command
  `docker-compose down`