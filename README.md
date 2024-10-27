# cwk
"[C]trl-[W]... [K]URWA" is what happens when I accidentally press vim's 'delete previous word' keybind in a non-neovim window such as Firefox. This is a very simple API for counting these events. It is my very first actual program in Go, so feel free to open an issue if you see some obvious possible improvement.

# Installation
## Using Docker
Fill the Basic Auth username and password in `docker-compose.yml` and then run
```sh
docker compose up -d
```
## Directly
```sh
go build
./cwk
```

# TODO
- Move trusted proxies to a variable so they're not hardcoded
