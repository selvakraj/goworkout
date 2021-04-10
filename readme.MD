set up redis

Setup docker desktop in you laptop then do the following steps:

install docker container images
1. docker run --name my-redis -p 6379:6379 --restart always --detach redis

login to container
2. docker exec -it my-redis sh

execute redis-cli in the container
3  # redis-cli 

4 type ping you should get pong respone

it should  hardly take 10 min to do the entire setup

--------

Setuo ENV Variables

export URL_DB=redis

export REDIS_URL=redis://localhost:6379
