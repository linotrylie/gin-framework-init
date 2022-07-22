FROM golang:1.18-alpine
 
RUN mkdir /app
 
WORKDIR /app
 
ADD . /app

RUN apk --update add tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && \
    apk del tzdata && \
    rm -rf /var/cache/apk/*
 
RUN go build -o main ./main.go
 
EXPOSE 8080

CMD /app/main