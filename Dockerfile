FROM golang:1.19
RUN mkdir -p src/app
WORKDIR /src/app/
COPY . /src/app
EXPOSE 8080
CMD make serv