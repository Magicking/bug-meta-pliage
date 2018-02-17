FROM golang
MAINTAINER Sylvain Laurent

ENV GOBIN $GOPATH/bin
ENV PROJECT_DIR github.com/Magicking/dao_tracabilite
ENV PROJECT_NAME d-a-o-t-server

ADD vendor /usr/local/go/src
ADD cmd /go/src/${PROJECT_DIR}/cmd
ADD models /go/src/${PROJECT_DIR}/models
ADD restapi /go/src/${PROJECT_DIR}/restapi
ADD internal /go/src/${PROJECT_DIR}/internal

WORKDIR /go/src/${PROJECT_DIR}

RUN go build -v -o /go/bin/main /go/src/${PROJECT_DIR}/cmd/${PROJECT_NAME}/main.go
ADD run.sh /go/src/${PROJECT_DIR}/
ENTRYPOINT /go/src/${PROJECT_DIR}/run.sh

EXPOSE 8090
