#!/usr/bin/env bash

gometalinter --config gometalinter.json ./... |grep -v 'exported method\|function'