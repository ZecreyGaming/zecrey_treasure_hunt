# Zecrey Browser Quest Server

This is an example server project of Zecrey BrowserQuest.

## Getting Started

If you do not have docker installed, [install docker](https://dockerdocs.cn/desktop/#download-and-install).


Clone the repo and create the `config.yaml` file

```bash
  cd ./game/api/server/etc/
  
  cp config.yaml.example config.yaml
```

Modify the `config.yaml` file to configure your information

```yaml
  Seed: "your Zecrey account seed"

  NftPrefix: "prefix of nft name"

  CollectionId: "ID of the collection you created"

```

Then,run the development server:
```bash
  cd ./game/api/server/etc/
  
  docker run -d -p 5566:5566 -v $(pwd)/config.yaml:/app/etc/config.yaml zecrey/browser-quest:0.0.1
  
  #Starting server at 0.0.0.0:5566...
```

