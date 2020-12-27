# nrcnewsapi
### For experimental purpose
<br/>

### Description
NRC = dutch news website

A simple NRC Scraper API written in GO with Colly (scraping framework) and Gin (web framework).

This API can be used for building NRC-news mobile and web app.
<br/>

### How-to
[Build a docker image]<br/>
docker build -f Dockerfile -t nrcnewsapi .

or

docker pull docker.pkg.github.com/ciccic/nrcnewsapi/nrcnewsapi:latest

[Run docker container]<br/>
docker run -i -t --memory=50m --memory-swap=50m --cpus=".5" --name nrcnewsapi -d -p 5011:5011 nrcnewsapi:latest

[Open webbrowser and run localhost:5011](localhost:5011)<br/>
<br/>

### Swagger
[Open webbrowser and run localhost:5011/swagger/index.html](localhost:5011/swagger/index.html)<br/>

<img src="../master/welcomedragon.png"
alt="drawing" width="800" height="400"/>
