FROM golang:1.23.3-alpine3.20 as builder

ARG GOLANGCI_LINT_VERSION=1.55.2

WORKDIR /develop

RUN apk add git make wget curl bash fzf vim jq

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v${GOLANGCI_LINT_VERSION} && \
    go install github.com/golang/mock/mockgen@v1.6.0

RUN curl -o $(go env GOPATH)/bin/swagger -L'#' "$(curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/latest | jq -r '.assets[] | select(.name | contains("'"$(uname | tr '[:upper:]' '[:lower:]')"'_amd64")) | .browser_download_url')"
RUN chmod 755 $(go env GOPATH)/bin/swagger

RUN git config --global --add safe.directory /develop
RUN cp /usr/share/fzf/key-bindings.bash /root/.bashrc
RUN echo "complete -W \"\`grep -oE '^[a-zA-Z0-9_.-]+:([^=]|$)' ?akefile | sed 's/[^a-zA-Z0-9_.-]*$//'\`\" make" >>/root/.bashrc

ENTRYPOINT [ "bash" ]