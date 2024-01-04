<p align="center" >
    <h2 align="center">anime-repository-web</h2>
    <h3 align="center">a http server based on go</h3>
    <p align="center">just a practice work for self-entertainment. </p>
<p>

---
developing...  
**Due to file operations involved, you will need to resolve the related permission issues on your own during runtime.**

### Installation & Usage


<details>
<summary>compile manually</summary>

```bash
go build . 
cd ./frontend 
npm run build  
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