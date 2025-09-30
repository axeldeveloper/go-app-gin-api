# Project
  - go 
  - gin
  - postgreSQL


## Getting Started

Follow the steps below to deploy and run the Go Gin application on your Koyeb account.


# Docker Run

```sh
# build imagens
$ docker build -f Dockerfile.dev -t app-gin .

# run container
$ docker run -d -p 8001:8001 --name api-gin app-gin

```
# Run 
To run the project, you need to have Go installed. Then, you can run the following command:

```sh

## make help 
## demostrate cmd

make run 
# ->  equals  go run ./server.go

```

# SWAGGER
```sh

go install github.com/swaggo/swag/cmd/swag@latest

echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.zshrc
source ~/.zshrc



   // @Summary      Descrição breve
   // @Description  Descrição detalhada
   // @Tags         exemplo
   // @Accept       json
   // @Produce      json
   // @Success      200  {object}  domain.Exemplo
   // @Router       /exemplo [get]
```

# Using 

$ go mod init example/mymodule

go get .
go get go.mongodb.org/mongo-driver/mongo    



## Contributing

If you have any questions, ideas or suggestions regarding this application sample, feel free to open an [issue](//github.com/koyeb/example-go-gin/issues) or fork this repository and open a [pull request](//github.com/koyeb/example-go-gin/pulls).

## Contact

[Koyeb](https://www.koyeb.com) - [@gokoyeb](https://twitter.com/gokoyeb) - [Slack](http://slack.koyeb.com/)
