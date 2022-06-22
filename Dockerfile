
# Download golang
FROM golang:latest

# Geef aan waar directory word aangemaakt
RUN mkdir /Website

# Alles wat in dir staat toevoegen
ADD . /Website

# Alles wat hier in staat begint hier te werken
WORKDIR /Website

# Download mysql
RUN go get -u github.com/go-sql-driver/mysql 

# Maakt main file
RUN go build -o main .

CMD ["/Website/main"]
# Mainfile gaat die openen
