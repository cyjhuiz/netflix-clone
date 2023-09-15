Netflix Clone - App Architecture
Frontend:
- React

Backend (Microservices with REST and GRPC endpoints):
- User
- Show
- Notification

Database:
- Postgres

Cache:
- Redis

Messaging Queue:
- RabbitMQ

Steps to run:
1. Ensure that "netflix_clone_db" is created in the Postgres database
2. Run the docker compose file by going to "/backend" directory and running "docker-compose up" in the terminal. This will create the Redis and RabbitMQ containers (created with default authentication details)
3. Run user service by going to "/user" directory and running "make run"
4. Run show service by going to "/show" directory and running "make run"
6. Run notification service by going to "/notification" directory and running "make run"
7. Run Frontend by going to "/frontend" directory and running "npm install" (if required) and "npm start" to start the App

Notes:
- the database tables for User, Show and Notification will be automatically created upon running the services, but the database "netflix_clone_db" has to be created first.
- To edit the database details look for the "NewStore()" function in the "dao.go" file


