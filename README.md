Jarvis - Feature Switch Slack Bot developed in Golang
==============
[![Build Status](https://travis-ci.org/felipecruz91/feature-switch-slack-bot.svg?branch=master)](https://travis-ci.org/felipecruz91/feature-switch-slack-bot) [![codecov](https://codecov.io/gh/felipecruz91/feature-switch-slack-bot/branch/master/graph/badge.svg?maxAge=0)](https://codecov.io/gh/felipecruz91/feature-switch-slack-bot/) [![Go Report Card](https://goreportcard.com/badge/github.com/felipecruz91/feature-switch-slack-bot)](https://goreportcard.com/report/github.com/felipecruz91/feature-switch-slack-bot) [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

## Setup the environment variables in my_env.list
`SLACK_TOKEN=token`
`Z_AUTH_KEY_QA=authKey`
`EXP_API_URL_QA=expApiUrl`

## Compile go source code
`CGO_ENABLED=0 GOOS=linux go build -a --installsuffix cgo --ldflags="-s" -o app`

## Build docker image
`docker build -t feature-switch-slack-bot -f Dockerfile .`

## Run docker image in a container
`docker run --rm -d --env-file ./my_env.list feature-switch-slack-bot`
