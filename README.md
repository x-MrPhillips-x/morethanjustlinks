[![Go](https://github.com/x-MrPhillips-x/morethanjustlinks/actions/workflows/go.yml/badge.svg)](https://github.com/x-MrPhillips-x/morethanjustlinks/actions/workflows/go.yml) [![codecov](https://codecov.io/gh/x-MrPhillips-x/morethanjustlinks/branch/main/graph/badge.svg?token=7YN9SBDGR1)](https://codecov.io/gh/x-MrPhillips-x/morethanjustlinks)

# >>>justlinks 🔗
is intended to share a free alternative to all my links, linktree, and the rest

go + gin + nextjs to display your links dynamically

## quick start guide

Guide supposes you: 
- have installed docker. If not please see https://docs.docker.com/engine/install/

Clone repo
```sh
git clone git@github.com:x-MrPhillips-x/morethanjustlinks.git
```

📍 Starting service locally 
```sh
# in one terminal run
go run main.go
# in another terminal 
cd nextjs-frontend/
npm run dev
```

🐳 Starting service with docker 
```sh
docker compose up
# Front end  http://localhost:3000/
```

Learn More about the Project
- [Nextjs Frontend](/nextjs-frontend/README.md) 
- Session Auth [gin auth](/handler/auth.go)

