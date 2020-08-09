# nrcnewsapi
### For experimental purpose
<br/>

NRC = dutch news website

A simple NRC Scraper API written in GO with Colly (scraping framework) and Gin (web framework).

This API can be used for building NRC-news mobile and web app.

[Build a docker image]<br/>
docker build -f Dockerfile -t nrcnewsapi .

or

docker pull docker.pkg.github.com/ciccic/nrcnewsapi/nrcnewsapi:latest

[Run docker container]<br/>
docker run -i -t --name nrcnewsapi -d -p 5011:5011 nrcnewsapi:latest

[Open webbrowser and run localhost:5011/]<br/>
docker run -i -t -p 5011:5011 nrcnewsapi:latest

<img src="../master/welcomedragon.png"
alt="drawing" width="800" height="400"/>
