[![Go](https://github.com/x-MrPhillips-x/morethanjustlinks/actions/workflows/go.yml/badge.svg)](https://github.com/x-MrPhillips-x/morethanjustlinks/actions/workflows/go.yml) [![codecov](https://codecov.io/gh/x-MrPhillips-x/morethanjustlinks/branch/main/graph/badge.svg?token=7YN9SBDGR1)](https://codecov.io/gh/x-MrPhillips-x/morethanjustlinks)

# morethanjustlinks
is intended to share a free alternative to all my links, linktree, and the rest

ngnix + go + gin + frontend to display your links dynamically

## quick start guide

Guide supposes you: 
- have installed docker. If not please see https://docs.docker.com/engine/install/

Clone repo
```
git clone git@github.com:x-MrPhillips-x/morethanjustlinks.git
```

Start service locally

```
docker compose up
```

## App instructions

To demostrate the app restAPI, send the curl below to setup db tables to support new accouts and links


```
curl http://localhost/setup \
    --include \
    --header "Content-Type: application/json" \
    --request "GET" \
```

output
```
{"msg":"created user tables succesfully"}%
```

Then navigate to http://localhost and click Create new account.

