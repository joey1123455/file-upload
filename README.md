# File Upload Application
This repository contains the source code for a file upload application. The application is built using Docker, Docker Compose, and Make. It uses Go as the programming language, Gin as the web framework, and MongoDB as the database.

## Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites
Before you can run this application, you need to have the following installed on your machine:
- **Docker:** Docker is an open-source platform that allows you to automate the deployment, scaling, and management of applications. You can download Docker from the official website.
- **Docker Compose:** Docker Compose is a tool for defining and running multi-container Docker applications. It uses a YAML file to configure the application's services, and then, with a single command, you create and start all the services from your configuration. You can download Docker Compose from the official website.
- **Make:** Make is a build automation tool that automatically builds executable programs and libraries from source code by reading files called Makefiles. It's typically available on Unix-based systems. If you're on Windows, you can install it through a package manager like Chocolatey.
- **Go:** Go is an open-source programming language that makes it easy to build simple, reliable, and efficient software.
- **Gin:** Gin is a web framework written in Go. It features a martini-like API with much better performance, up to 40 times faster thanks to httprouter. If you need performance and good productivity.
- **MongoDB:** MongoDB, also known as Mongo, is a free and open-source document database management system. It emphasizes extensibility and NoSQL compliance.

### Installing
#### Clone the repository:
``` bash 
git clone https://github.com/joey1123455/file-upload.git
```

#### Change the current directory
``` bash 
cd file-upload
```


### Running the Application
The Makefile in this repository provides several commands for managing the application:
When using a unix based operating system rember to add sudo before all make commands i.e.
```bash
sudo make up
```

- ``make build:`` This command builds the Docker image for the application.
- ``make dev:`` This command starts the application in development mode. The application will run in the foreground, and logs will be displayed in the terminal.
- ``make prod:`` This command starts the application in production mode. The application will run in the background.
- ``make down:`` This command stops the application.
- ``make logs:`` This command displays the logs for the application.
- ``make test:`` This command runs the tests for the application.

## Built With
- **Docker** - The containerization platform used.
- **Docker Compose** - The tool used for defining and running multi-container Docker applications.
- **Make** - The build automation tool used.
- **Go** - The programming language used.
- **Gin** - The web framework used.
- **MongoDB** - The database used.
- **Cloudinary** - CDN used.
