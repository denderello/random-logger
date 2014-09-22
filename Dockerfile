FROM golang:1.3-onbuild

ADD logger.go /logger.go

CMD ["go", "run", "/logger.go"]
