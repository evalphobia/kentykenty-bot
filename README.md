kentykenty-bot
----

[![GoDoc][1]][2] [![License: MIT][3]][4] [![Release][5]][6] [![Build Status][7]][8] [![Co decov Coverage][11]][12] [![Go Report Card][13]][14] [![Code Climate][19]][20] [![BCH compliance][21]][22] [![Downloads][15]][16]

[1]: https://godoc.org/github.com/evalphobia/kentykenty-bot?status.svg
[2]: https://godoc.org/github.com/evalphobia/kentykenty-bot
[3]: https://img.shields.io/badge/License-MIT-blue.svg
[4]: LICENSE.md
[5]: https://img.shields.io/github/release/evalphobia/kentykenty-bot.svg
[6]: https://github.com/evalphobia/kentykenty-bot/releases/latest
[7]: https://travis-ci.org/evalphobia/kentykenty-bot.svg?branch=master
[8]: https://travis-ci.org/evalphobia/kentykenty-bot
[9]: https://coveralls.io/repos/evalphobia/kentykenty-bot/badge.svg?branch=master&service=github
[10]: https://coveralls.io/github/evalphobia/kentykenty-bot?branch=master
[11]: https://codecov.io/github/evalphobia/kentykenty-bot/coverage.svg?branch=master
[12]: https://codecov.io/github/evalphobia/kentykenty-bot?branch=master
[13]: https://goreportcard.com/badge/github.com/evalphobia/kentykenty-bot
[14]: https://goreportcard.com/report/github.com/evalphobia/kentykenty-bot
[15]: https://img.shields.io/github/downloads/evalphobia/kentykenty-bot/total.svg?maxAge=1800
[16]: https://github.com/evalphobia/kentykenty-bot/releases
[17]: https://img.shields.io/github/stars/evalphobia/kentykenty-bot.svg
[18]: https://github.com/evalphobia/kentykenty-bot/stargazers
[19]: https://codeclimate.com/github/evalphobia/kentykenty-bot/badges/gpa.svg
[20]: https://codeclimate.com/github/evalphobia/kentykenty-bot
[21]: https://bettercodehub.com/edge/badge/evalphobia/kentykenty-bot?branch=master
[22]: https://bettercodehub.com/



Slack Bot for killing time.


# Install

```bash
$ go get -u github.com/evalphobia/kentykenty-bot
```

# Build

```bash
$ make build
```

for Raspberry Pi

```bash
$ make build-arm6
```

# Run

```bash
SLACK_RTM_TOKEN=xoxb-0000... ./bin/kentykenty-bot
```

## Environment variables

|Name|Description|
|:--|:--|
| `SLACK_RTM_TOKEN` | [Slack Bot Token](https://slack.com/apps/A0F7YS25R-bots) |
| `SLACK_BOT_TOKEN` | [Slack Bot Token](https://slack.com/apps/A0F7YS25R-bots) |
| `SLACK_TOKEN` | [Slack Bot Token](https://slack.com/apps/A0F7YS25R-bots) |
| `BOBO_DEBUG` | Flag for debug logging. Set [boolean like value](https://golang.org/pkg/strconv/#ParseBool). |
| `AWS_ACCESS_KEY_ID` | [AWS Access Key ID](https://github.com/aws/aws-sdk-go/blob/bef02444773a49eaf30cdd615920b56896827c06/aws/credentials/env_provider.go) |
| `AWS_SECRET_ACCESS_KEY` | [AWS Secret Access Key](https://github.com/aws/aws-sdk-go/blob/bef02444773a49eaf30cdd615920b56896827c06/aws/credentials/env_provider.go) |
| `FACEPP_API_KEY` | [API Key of Face++](https://github.com/evalphobia/go-face-plusplus). |
| `FACEPP_API_SECRET` | [API Secret of Face++](https://github.com/evalphobia/go-face-plusplus). |
