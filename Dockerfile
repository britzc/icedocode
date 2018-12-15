FROM golang:1.10

WORKDIR /app

ADD index.html .
ADD icedoapp .

CMD ["/icedoapp"]
