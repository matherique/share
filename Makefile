TARGET = bin/main

# Comando para compilar o código
build:
	go build -o $(TARGET) cmd/api/main.go

# Comando para executar o código
run:
	go build -o $(TARGET) cmd/api/main.go
	./$(TARGET)

# Alvo "all" realiza o build e executa o código em sequência
all: build run

# Alvo "clean" remove o arquivo executável gerado
clean:
	rm -f $(TARGET)
