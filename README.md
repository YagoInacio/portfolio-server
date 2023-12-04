<p align="center">
  <img alt="GitHub language count" src="https://img.shields.io/github/languages/count/yagoinacio/portfolio-server?color=353949">

  <img alt="Repository size" src="https://img.shields.io/github/repo-size/yagoinacio/portfolio-server">

  <a href="https://github.com/yagoinacio/portfolio-server/commits/main">
    <img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/yagoinacio/portfolio-server">
  </a>

   <img alt="License" src="https://img.shields.io/badge/license-MIT-brightgreen">

  <a href="https://yagofaran.dev">
    <img alt="Made by Yago Faran" src="https://img.shields.io/badge/made_by-Yago_Faran-353949">
  </a>
</p>

<h1 align="center">
    <img alt="Logo Yago Faran" title="#YagoFaran" src="./assets/logo.svg" />
</h1>

<p align="center">
 <a href="#-portfolio-server">About</a> •
 <a href="#-tech-stack">Tech Stack</a> • 
 <a href="#-features">Features</a> •
 <!-- <a href="#-layout">Layout</a> •  -->
 <a href="#-how-it-works">How it works</a> • 
 <!-- <a href="#-contributors">Contributors</a> •  -->
 <a href="#-author">Author</a> • 
 <a href="#-license">License</a>
</p>

## 💻 Portfolio Server

This API is designed to improve the management of project and experience information for your portfolio website, ensuring a seamless rendering and effortless updating process. Additionally, it facilitates interaction with Firebase Storage, allowing you to efficiently load and display icons and images, enhancing the overall visual experience of your portfolio.

## 🛠 Tech Stack

-   **[Echo](https://echo.labstack.com)**
-   **[MongoDB](https://www.mongodb.com/docs/drivers/go/current/)**
-   **[Firebase Storage](https://firebase.google.com/docs/storage)**
-   **[Testify](https://github.com/stretchr/testify)**
-   **[Viper](https://github.com/spf13/viper)**

## ✨ Features

- [x] Technologies:
  - [x] Create
  - [x] List
  - [x] Enable/Disable

- [x] Experiences:
  - [x] Create
  - [x] List
  - [x] Update

- [ ] Projects
  - [x] Create
  - [x] List
  - [ ] Disable
  - [ ] Update

- [x] Images:
  - [x] Get Images
  - [x] Get Icons

## 🚀 How it works

This instructions will allow you to run a functional version of the project on your local machine.

### 📋 Pre-requisites

Before you begin, you will need to have the following tools installed on your machine:
[Git](https://git-scm.com), [Go](https://go.dev).
In addition, it is good to have an editor to work with the code like [VSCode](https://code.visualstudio.com/)

#### 🔧 Instalation

```bash
# Clone this repository
$ git clone git@github.com:yagoinacio/portfolio-server.git

# Access the project folder cmd/terminal
$ cd project

# install the dependencies
$ go mod tidy
```

#### 🔧 Configuration

To be able to run the application you need to set up the environment variables.

For that, create the file ```.env``` on the project's root folder.

You can follow the example bellow:

```bash
# db.env.properties:
API_PORT=9000
DATABASE_URL="mongodb+srv://[USERNAME]:[PASSWORD]@[HOST]/[DATABASE]?retryWrites=true&w=majority"
DATABASE_NAME=mongo_database
```

You will also need to add your firebase credentials file to project's root folder with the name ```firebase_credentials.json```.

#### 🎲 Running the application

```bash
# Run the application in development mode
$ go run ./cmd/portfolio/main.go

# The server will start at port: 9000 - go to http://localhost:9000
```

You can try out the API using its swagger documentation on http://localhost:9000/swagger-ui/index.html

#### ✅ Running automated tests

```bash
# Run automated tests
$ go test ./...

# The test automation will run for unit tests
```

## 🦸 Author

<a href="https://yagofaran.dev">
 <img style="border-radius: 50%;" src="https://avatars.githubusercontent.com/yagoinacio" width="100px;" alt=""/>
 <br />
 <sub><b>Yago Faran 💧</b></sub>
</a>

[![Github Badge](https://img.shields.io/badge/-YagoInacio-gray?style=flat-square&labelColor=gray&logo=github&logoColor=white&link=https://github.com/yagoinacio)](https://github.com/yagoinacio)
[![Linkedin Badge](https://img.shields.io/badge/-Yago-blue?style=flat-square&logo=Linkedin&logoColor=white&link=https://www.linkedin.com/in/yagoinacio/)](https://www.linkedin.com/in/yagoinacio/) 
[![Gmail Badge](https://img.shields.io/badge/-yagofaran@gmail.com-c14438?style=flat-square&logo=Gmail&logoColor=white&link=mailto:yagofaran@gmail.com)](mailto:yagofaran@gmail.com)

## 📝 License

This project is under the license [MIT](./LICENSE).

Made with ❤️ by Yago Faran 👋🏽 [Get in touch!](https://www.linkedin.com/in/yagoinacio/)