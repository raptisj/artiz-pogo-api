# 🎹 Artiz PoGo API 🥁

_Still WIP_

An API that lists artists and their songs. Authenticated users can perform actions such as liking a song, following an artist and creating playlists.

### Technologies
**Po**stgres, **Go**lang and Docker.

### Libaries
- [chi](https://github.com/go-chi/chi/tree/master) - lightweight router for building Go HTTP services
- [golang-jwt](https://github.com/golang-jwt/jwt?tab=readme-ov-file) - to generate and verify JWTs
- [godotenv](https://github.com/joho/godotenv) - to read environments variables
- [gorm](https://github.com/go-gorm/gorm) - an popular ORM for Golang
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) - for hashing and comparing passwords


### Usage and info

```
git clone git@github.com:raptisj/artiz-pogo-api.git
cd ./artiz-pogo-api
```

To build and run your images: 
```
make run
```

To connect to psql: 
```
make psql
```



### Features

- custom auth
- list single or all artists
- follow artist
- song list per artists
- liked songs
- create playlists
- user profile
- pagination

### Resources
Some resources to help you learn Golang
- [A Tour of Go](https://go.dev/tour/basics/1)
- [Effective Go](https://go.dev/doc/effective_go)
- [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests)

### Author
John Raptis