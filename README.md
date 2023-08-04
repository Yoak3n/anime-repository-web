<p align="center" >
    <h2 align="center">anime-repository-web</h2>
    <h3 align="center">a http server based on go</h3>
    <p align="center">just a practice work for self-entertainment. </p>
<p>

---
developing...  
preparing for last one step (still a toy)

### Installation


<details>
<summary>compile manually</summary>

```bash
go build . && cd ./frontended && npm run build && cp ./dist/ ../resource/
```

</details>

<details>
<summary>Docker</summary>

```bash
docker run \
  -v '/app/anime-repository-web/config.yml:/app/anime-repository-web/config.yml'\
  -p 8080:8080 \ 
  -name arw \
  yoaken/anime-repository-web:latest
```
</details>

<details>
<summary>docker-compose</summary>

Change to the work directory where is the `docker-compose.yml`file and run:
```bash
docker-compose up -d 
```

</details>

### Usage

### Reference
根据[Anime-Repository](https://github.com/Chikage0o0/Anime-Repository)用go实现的一个功能类似的前后端不分离的web项目