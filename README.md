# slack-cli
Simple command-line client for slack by golang.

You can use curl as alternative

```
curl -X POST --data-urlencode 'payload={"channel": "<channel>", "username": "<username>", "text": "<message>", "icon_emoji": ":ghost:"}' https://hooks.slack.com/services/<webhook>
```

## Features
* Post message only
* Use Incoming Web hook service on slack


## Install
### via Go
```
git clone <this repo>
go install
```

or

```
$ go get <this repo>
```

## Usage
```shell
$ slack-cli -webhook https://hooks.slack.com/services/<yourhook>  -message <message>
```

## Remove
### via Go
```shell
$ rm $GOPATH/bin/slack-cli
```
### via Binary
Remove your slack-cli binary in `$PATH`.  
```
## License
MIT
