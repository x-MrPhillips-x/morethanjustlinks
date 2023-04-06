# morethanjustlinks
is intended to share a free alternative to all my links, linktree, and the rest

ngnix + go + gin + frontend to display your links dynamically

## quick start guide

Clone repo
```
git clone git@github.com:x-MrPhillips-x/morethanjustlinks.git
```

Start service locally

```
docker compose up
```

Then navigate to http://localhost

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

