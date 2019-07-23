# MMORPG Communication Analysis
This repository is part of my bachelor thesis.
I implemented Rudy (Tibia), Salz (Albion) and Willson (GuildWars 1/2) microsservices architectures for a communication analisys using Golang. You can use for some tests :)

# Go Dep Graph

```sh
    go get github.com/kisielk/godepgraph
    sudo pacman -S graphviz
```

```sh
    godepgraph github.com/schweigert/mmorpg-communication-analysis/clients/wclient
    godepgraph github.com/schweigert/mmorpg-communication-analysis/clients/wclient | dot -Tpng -o wclient.png
    # or other
```