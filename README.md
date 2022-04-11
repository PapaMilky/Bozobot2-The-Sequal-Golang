# Bozo Bot 2 - The Sequel
## Made In Golang. Designed For Shits And Giggles
#### (Now with docker support)

## About
BB2 (Bozo Bot 2) is a Golang based discord bot for a friend group server.  
This is bot 2 of 3 in the collection of "Bozo Bots"

## Running
### Running Via Docker
```
docker run -d \
  --name=discord \
  --net=host \
  -e VERSION=docker \
  -e GUILD="<SINGULAR Guild ID>" \
  -e TOKEN="<Bot Token>" \
  --restart unless-stopped \
  papamilky/bozo-bot-2
  ```
### Running Via CLI
Don't.

## Credits
Discord Bindings: [This Chad](https://github.com/diamondburned/arikawa)

## Other Bozo's
Nullcode: [Github](https://github.com/NullCode1337) | [Bozo Bot](NaN)  
Dottik: [Github](https://github.com/usrDottik) | [Bozo Bot](https://github.com/usrDottik/Bozos-Bot)
