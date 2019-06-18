# slack-cli
Simple command-line client for slack by golang.

You can use curl as alternative
```
curl -X POST --data-urlencode 'payload={"channel": "<channel>", "username": "<username>", "text": "<message>", "icon_emoji": ":ghost:"}' https://hooks.slack.com/services/<webhook>
```

but this tool can read Stdin and you can use unix pipe to pass message like:
```
tail -100 '/var/log/nginx/access.log' | slack-cli -webhook https://hooks.slack.com/services/<yourhook>
```

## Features
* Post message only
* Use Incoming Web hook service on slack
* Use unix pipe to pass message body

## Install
### via Go
```
git clone <this repo>
go mod init
go install ./cmd/...
```

or

```
$ go get <this repo>
```

## Usage
```shell
$ slack-cli -webhook https://hooks.slack.com/services/<yourhook>  -message <message>
```
or via exporting
```shell
export SLACK_CLI_WEBHOOK=https://hooks.slack.com/services/<yourhook>
# for csh
# set SLACK_CLI_WEBHOOK="https://hooks.slack.com/services/<yourhook>"
$ slack-cli -message <message>
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
