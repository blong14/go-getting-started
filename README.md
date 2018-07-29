![Build Status](https://travis-ci.org/blong14/goping-web.svg?branch=master)

# goping-web

Application uptime notifier written in Go.

## Setup
```sh
$ go get github.com/codegangsta/gin
$ dep ensure
```

## Running Locally

Hot reload
```sh
$ gin run main.go
```

## Deploying to Heroku

Make sure you have [Go](http://golang.org/doc/install) and the [Heroku Toolbelt](https://toolbelt.heroku.com/) installed.

```sh
$ go get -u github.com/blong14/goping-web
$ cd $GOPATH/src/github.com/blong14/goping-web
$ heroku local
```

Your app should now be running on [localhost:5000](http://localhost:5000/).

### Deploy
```sh
$ heroku create
$ git push heroku master
$ heroku open
```

### Documentation

For more information about using Go on Heroku, see these Dev Center articles:

- [Go on Heroku](https://devcenter.heroku.com/categories/go)
