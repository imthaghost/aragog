FROM golang:alpine3.15

# Enviornment variables
ENV APP_NAME aragog
ENV PORT 8080
ENV github_user imthaghost
ENV github_personal_token 0

RUN apk update
RUN apk add git

RUN git config \
    --global \
    url."https://${github_user}:${github_personal_token}@github.com".insteadOf \
    "https://github.com"

# Open system port
EXPOSE ${PORT}

# Working directory
WORKDIR /go/src/${APP_NAME}

COPY . /go/src/${APP_NAME}

# Install dependecies from mod file
RUN go mod download

RUN apk add git

# Build application
RUN go build -o ${APP_NAME} ./cmd/aragog

# Run application
CMD ./${APP_NAME}