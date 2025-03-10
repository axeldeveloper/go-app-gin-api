# Definindo variáveis
CONTAINER_NAME := mss_api
IMAGE_NAME := api:mss
DOCKERFILE_PATH := ./Dockerfile.dev  # Defina o caminho para o seu Dockerfile
PORT := 8001

# Comandos
run:
	@echo "Executando o servidor Go..."
	go run ./server.go

cl:
	@echo "Limpar o terminal..."
	clear

rund: build
	@echo "Executando o binário..."
	@./bin/api

build:
	@echo "Construindo a aplicação Go..."
	@go build -o bin/api

test:
	@echo "Executando os testes Go..."
	@go test -v ./...

dc:
	@echo "Construindo a imagem Docker..."
	@docker build -f $(DOCKERFILE_PATH) -t $(IMAGE_NAME) .

cr:
	@echo "Criando e executando o container Docker..."
	@docker run --rm -p $(PORT):$(PORT) --name $(CONTAINER_NAME) $(IMAGE_NAME)

tr: 
	@echo "Construindo a imagem e executando o container Docker..."
	@docker build -f $(DOCKERFILE_PATH) -t $(IMAGE_NAME) . && docker run --rm -p $(PORT):$(PORT) --name $(CONTAINER_NAME) $(IMAGE_NAME)

help:
	@echo "Comandos disponíveis no Makefile:"
	@echo "  make run       - Executa a aplicação Go"
	@echo "  make cl        - Limpa o terminal"
	@echo "  make rund      - Executa o binário da aplicação"
	@echo "  make build     - Compila a aplicação Go"
	@echo "  make test      - Executa os testes da aplicação Go"
	@echo "  make dc        - Constrói a imagem Docker"
	@echo "  make cr        - Cria e executa o container Docker"
	@echo "  make tr        - Constrói e executa a imagem Docker em sequência"
	@echo "  make help      - Exibe este resumo de comandos"
