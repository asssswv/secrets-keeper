# secrets-keeper

# Run Redis 
    docker-compose up --build redis
# Run Server
    go run app/cmd/main.go

# PS(clearing the docker cache)
    docker system prune --volumes  