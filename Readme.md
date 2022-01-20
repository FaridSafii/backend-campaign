<p align="center"><a href="https://laravel.com" target="_blank"><img src="https://raw.githubusercontent.com/laravel/art/master/logo-lockup/5%20SVG/2%20CMYK/1%20Full%20Color/laravel-logolockup-cmyk-red.svg" width="400"></a></p>

### GoLang Backend - Campaign 

Campaign system with GoLanguage, Database MySQL, Web Framework Gin, and Gorm library for database

#### How to run

Use the terminal to run project.

```bash
go run main.go
```

#### Endpoint create user
User needs to enter name, email, occupation, and password to register at the endpoint

```bash
localhost:8088/api/v1/users
```

#### Endpoint login user
User need to enter data email and password to login. the system will store user session data encrypted using JWT Authorization
```bash
localhost:8088/api/v1/sessions
```