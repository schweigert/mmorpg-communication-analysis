FROM golang:1.12

ADD . /go/src/github.com/schweigert/mmorpg-communication-analysis

RUN go install github.com/schweigert/mmorpg-communication-analysis/infrastructure/services/rauthservice
RUN go install github.com/schweigert/mmorpg-communication-analysis/infrastructure/services/rcrudservice
RUN go install github.com/schweigert/mmorpg-communication-analysis/infrastructure/services/rgameservice
RUN go install github.com/schweigert/mmorpg-communication-analysis/infrastructure/services/rwebservice
RUN go install github.com/schweigert/mmorpg-communication-analysis/infrastructure/services/sauthservice
RUN go install github.com/schweigert/mmorpg-communication-analysis/infrastructure/services/schatservice
RUN go install github.com/schweigert/mmorpg-communication-analysis/infrastructure/services/sgameservice
RUN go install github.com/schweigert/mmorpg-communication-analysis/infrastructure/services/sglobalservice
RUN go install github.com/schweigert/mmorpg-communication-analysis/infrastructure/services/swebservice
RUN go install github.com/schweigert/mmorpg-communication-analysis/infrastructure/services/wauthservice
RUN go install github.com/schweigert/mmorpg-communication-analysis/infrastructure/services/wgameservice
RUN go install github.com/schweigert/mmorpg-communication-analysis/infrastructure/services/wglobalservice
RUN go install github.com/schweigert/mmorpg-communication-analysis/infrastructure/services/wwebservice

RUN go install github.com/schweigert/mmorpg-communication-analysis/clients/wclient
