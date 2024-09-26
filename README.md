<div align="center" id="top"> 
  
</div>

<div align="center" id="top" >
 <img style=" width: 120px; height: 120px;" src="https://go.dev/images/gophers/biplane.svg" alt="Go API" />
</div>


<h1 align="center">Go API Master Example</h1>


<h4 align="center">
 🚧  Go API 🚀 Under construction...  🚧
</h4>

<p align="center">
  <img alt="Github top language" src="https://img.shields.io/github/languages/top/LuizWeitz/Go-API?color=56BEB8">
  <img alt="License" src="https://img.shields.io/github/license/LuizWeitz/Go-API?color=56BEB8">
  <img alt="Go Card Report" src="https://goreportcard.com/badge/github.com/luizweitz/go-api"> 
</p>
<hr> 

<p align="center">
  <a href="#dart-about">About</a> &#xa0; | &#xa0; 
  <a href="#sparkles-features">Features</a> &#xa0; | &#xa0;
  <a href="#rocket-technologies">Technologies</a> &#xa0; | &#xa0;
  <a href="#white_check_mark-requirements">Requirements</a> &#xa0; | &#xa0;
  <a href="#checkered_flag-starting">Starting</a> &#xa0; | &#xa0;
  <a href="#memo-license">License</a> &#xa0; | &#xa0;
  <a href="https://github.com/LuizWetz" target="_blank">Author</a>
</p>

<br>

## :dart: About ##

This a project is a example of API Rest in Go using Gin, Gorm & PostgreSQL.
<br>
For good practies of coding i'm using [Twelve Go Best Practices](https://go.dev/talks/2013/bestpractices.slide#1)
<br>
For quality of code i'm using [Go Report Card](https://goreportcard.com/)
<br>
For folder struct this project i'm using [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
<br>
Link for <a href="https://goreportcard.com/report/github.com/luizweitz/go-api">Go Report Card</a>

## :sparkles: Features ##

The features of projet is
  * Create User: Returns the created user.
  * Update User: Returns a message indicating success or failure.
  * Find All Users: Returns a list of users or an empty list if no users are found.
  * Find User By ID: Returns the user if found, or an error message if not found or error.
  * Delete User By ID: Returns a message indicating success or failure.

## :rocket: Technologies ##

The following tools were used in this project:

- [Go](https://go.dev/)
- [Gin](https://gin-gonic.com/)
- [Gorm](https://gorm.io/index.html)
- [PostgreSQL](https://www.postgresql.org/)

## :white_check_mark: Requirements ##

Before starting :checkered_flag:, you need to have [Git](https://git-scm.com) and [Go](https://go.dev/) installed. Obs.: for PostgreSQL you can use a datababase provider, like AWS 

## :checkered_flag: Starting ##

```bash
# Clone this project
$ git clone https://github.com/LuizWetz/go-api

# Access
$ cd Go-API

# Install dependencies
$ go mod tiny

# Run the project
$ cd cmd
$ go run main.go

# The server will initialize in the <http://localhost:8080>
# The swagger docs will initialize in the <http://localhost:8080/swagger/index.html#/>

```

## :memo: License ##

This project is under license from MIT. For more details, see the [LICENSE](LICENSE) file.


Made with :heart: by <a href="https://github.com/LuizWetz" target="_blank">Luiz Weitz</a>

&#xa0;

<a href="#top">Back to top</a>
