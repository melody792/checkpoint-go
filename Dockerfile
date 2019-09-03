FROM golang:1.12.9
COPY . /go/src/checkpoint-go
WORKDIR /go/src/checkpoint-go
RUN go get github.com/astaxie/beego && go get github.com/beego/bee && go get github.com/go-sql-driver/mysql
EXPOSE 5000
CMD ["bee", "run"]