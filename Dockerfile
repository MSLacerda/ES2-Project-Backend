# Build Stage
FROM lacion/alpine-golang-buildimage:1.9.7 AS build-stage

LABEL app="ES2-Project-Backend"
LABEL REPO="https://github.com/MSLacerda/ES2-Project-Backend"

ENV PROJPATH=/go/src/github.com/MSLacerda/ES2-Project-Backend

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/MSLacerda/ES2-Project-Backend
WORKDIR /go/src/github.com/MSLacerda/ES2-Project-Backend

RUN apk add --update libxslt \
                     alpine-sdk \
                     libxml2-dev \
                     libxslt-dev \
                     python-dev \
                     openssl-dev \
                     libffi-dev \
                     zlib-dev \
                     py-pip
RUN apk update \
    && apk upgrade \
    && apk --no-cache add --update libxml2-utils tcl apache2 ca-certificates \ 
    apk-tools curl build-base supervisor cups-client dcron bind-tools rsync

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN make setup
RUN make get-deps

RUN env GO111MODULE=on go install ./...
RUN make build

# Final Stage
FROM lacion/alpine-base-image:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/MSLacerda/ES2-Project-Backend"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/ES2-Project-Backend/bin

WORKDIR /opt/ES2-Project-Backend/bin

COPY --from=build-stage /go/src/github.com/MSLacerda/ES2-Project-Backend/bin/ES2-Project-Backend /opt/ES2-Project-Backend/bin/
RUN chmod +x /opt/ES2-Project-Backend/bin/ES2-Project-Backend

# Create appuser
RUN adduser -D -g '' ES2-Project-Backend
USER ES2-Project-Backend

ENTRYPOINT ["/usr/bin/dumb-init", "--"]
EXPOSE 8080

CMD ["/opt/ES2-Project-Backend/bin/ES2-Project-Backend", "run"]
