
# goping

Application uptime notifier written in Go.

## Running Locally

Make sure you have [Go](http://golang.org/doc/install) and the [Heroku Toolbelt](https://toolbelt.heroku.com/) installed.

```sh
$ go get -u github.com/blong14/goping-web
$ cd $GOPATH/src/github.com/blong14/goping-web
$ heroku local
```

Your app should now be running on [localhost:5000](http://localhost:5000/).

You should also install [govendor](https://github.com/kardianos/govendor) if you are going to add any dependencies to the sample app.

## Deploying to Heroku

```sh
$ heroku create
$ git push heroku master
$ heroku open
```

## Documentation

For more information about using Go on Heroku, see these Dev Center articles:

- [Go on Heroku](https://devcenter.heroku.com/categories/go)
