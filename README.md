# Bozo Bot 2 - The Sequel
## Made In Golang. Designed For Shits And Giggles
#### (Now with docker support)

## About
BB2 (Bozo Bot 2) is a Golang based discord bot for a friend group server. This is bot 2 of 3 in the collection of "Bozo Bots"

## Running
### Running Via Docker
```
docker run -d \
  --name=discord \
  --net=host \
  -e VERSION=docker \
  -e GUILD="" \
  -e TOKEN="" \
  --restart unless-stopped \
  PapaMilky/Bozobot2-The-Sequal-Golang```
