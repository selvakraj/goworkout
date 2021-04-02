set up redis

docker run --name my-redis -p 6379:6379 --restart always --detach redis

Check:
docker exec -it my-redis sh

execute redis-cli in the shell
ping

--------

Setuo ENV Variables

export URL_DB=redis

export REDIS_URL=redis://localhost:6379
