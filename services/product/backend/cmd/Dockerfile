    # Использование golang:alpine в качестве билдера.
    # golang:alpine означает линукс с установленным golang
    # builder означает, что после того, как скомпилируется приложение,
    # можно будет скопировать бинарный файл в другой образ без установленного golang
    # это необходимо для того, чтобы образ не занимал много места
    # 300 мб с установленным golang либо 15 мб без golang
    FROM golang:alpine as builder

    # Директория, путь, в котором находится командная строка
    WORKDIR /app/services/product/backend/cmd

    # Копирование исходного кода в контейнер
    # Путь начинается с корня проекта, так как сборка происходит
    # с помощью docker-compose.yml файла, который находится в корне проекта

    # Очень важно соблюдать последовательность папок, так как в коде часто используются
    # относительные пути, в том числе и в import командах

    # Копирование backend части
    COPY /services/product/backend /app/services/product/backend
    # Копирование пакетов, используемых в backend 
    COPY /services/product/domain /app/services/product/domain
    COPY /services/product/dto /app/services/product/dto
    # Копирование базовых утилит для всех сервисов
    COPY /services/base /app/services/base
    # Копирование файлов модуля
    COPY /go.mod /app
    COPY /go.sum /app

    # Установка библиотек "RUN go mod download" пропущена, так как
    # вместо этого использутся механизм "Volume" — общее хранилище между контейнерами
    # В docker-compose.yml файле прописано, что содержимое директории "vendor" копируется
    # в /go/pkg/mod — директории исходного кода для библиотек
    # Это позволяет не скачивать библиотеки каждый раз 

    # Компиляция проекта
    RUN go build -o main .

    FROM alpine

    # Использование чистого Alpine (дистрибутив Linux). Характеризуется легковесностью
    # Это новый образ, который возьмет только исходный код
    WORKDIR /app/services/product/backend/cmd

    # Скопировать из builder исходный код в текущую рабочую директорию — из /app/services/product/backend/cmd/main в .
    COPY --from=builder /app/services/product/backend/cmd/main .

    # Установка порта
    EXPOSE 7071

    # Запуск бинарного файла
    CMD ["./main"]
