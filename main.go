package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//create gin api
	r := gin.Default()

	//create a new router group
	root := r.Group("/")

	//respond to GET requests on / with a HTML page
	root.GET("/", func(c *gin.Context) {
		const site string = `<!DOCTYPE html><html><head><title>iFrame Timer</title></head><body><script>;function goto(n){var o='/';if(window.location.pathname.endsWith('/')){o=''};const url=window.location.origin+window.location.pathname+o+'timer?site='+document.getElementById('url').value;window.location.replace(url)};</script><input type="text" id="url" placeholder="URL..." /><button type="submit" onClick="goto()">Goto</button></body></html>`;
		c.Data(200, "text/html; charset=utf-8", []byte(site))
	})

	root.GET("/timer", func(c *gin.Context) {
		const site string = `<!DOCTYPE html><html><head><title>iFrame Timer</title><link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.3.1/dist/css/bootstrap.min.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous" /><style>body{margin:0;padding:0}</style></head><body><div class="row no-gutters"><div class="col-9" style="display:flex;height:100vh"><iframe id="iframe" width="100%" height="100%" style="border:none"></iframe><script>;if(window.location.search==''){window.location.replace(window.location.origin+window.location.pathname.split('/').slice(0,-1).join('/'))};document.getElementById('iframe').src=window.location.search.split('&')[0].split('=').slice(1).join('=');</script></div><div class="col-3"><script>;var time=120,set=0,timer=0;setInterval(function(){if(timer==1){document.getElementById('audio').play();set=set+1;document.getElementById('set').innerHTML=set};if(timer>0){timer=timer-1;document.getElementById('timer').innerHTML=timer}},1000);function start(){timer=time;document.getElementById('set').innerHTML=set};function reset(){timer=time;set=0;document.getElementById('set').innerHTML=set};</script><h1>Timer: <span id="timer"></span></h1><h1>Current set: <span id="set"></span></h1><button type="button" class="btn btn-primary" onclick="start()"> Pause </button><button type="button" class="btn btn-primary" onclick="reset()"> Next/Start </button><audio id="audio" src="https://cdn.pixabay.com/download/audio/2022/03/10/audio_2cd9a13a55.mp3?filename=buzzer-dog-39284.mp3"></audio></div></div><script src="https://code.jquery.com/jquery-3.3.1.slim.min.js" integrity="sha384-q8i/X+965DzO0rT7abK41JStQIAqVgRVzpbzo5smXKp4YfRvH+8abtTE1Pi6jizo" crossorigin="anonymous"></script><script src="https://cdn.jsdelivr.net/npm/popper.js@1.14.7/dist/umd/popper.min.js" integrity="sha384-UO2eT0CpHqdSJQ6hJty5KVphtPhzWj9WO1clHTMGa3JDZwrnQq4sF86dIHNDz0W1" crossorigin="anonymous"></script><script src="https://cdn.jsdelivr.net/npm/bootstrap@4.3.1/dist/js/bootstrap.min.js" integrity="sha384-JjSmVgyd0p3pXB1rRibZUAYoIIy6OrQ6VrjIEaFf/nJGzIxFDsf4x0xIM+B07jRM" crossorigin="anonymous"></script></body></html>`;
		c.Data(200, "text/html; charset=utf-8", []byte(site))
	})

	//listen and serve on port 4000
	r.Run(":4000")
}