# Desafío Web


## Objetivo
El objetivo de esta guía práctica es poder afianzar y profundizar los conceptos vistos en Go Web. Para esto, vamos a plantear un desafío integrador que nos permitirá repasar los temas que estudiamos. 

## Planteo
Una aerolínea utiliza un programa para calcular diversas informaciones sobre vuelos que ocurren en la misma. Si bien su programa se encuentra funcional y sin dificultades, han decidido convertir el mismo en una API Rest.

### *_Requerimiento 1_*

Crear los packages correspondientes para que la arquitectura de la aplicación cumpla con la definición de *REST* y el *Diseño Orientado a Dominio*.

> Tip: Los paquetes internal y cmd son fundamentales

### *_Requerimiento 2_*

Ubicar los archivos de manera tal que cada uno quede en su respectivo package.

> Tip: Recordar chequear los *imports*.


### *_Requerimiento 3_*

Desarrollar los métodos correspondientes a la estructura Ticket. Uno de ellos debe devolver la cantidad de tickets de un destino. El  otro método debe devolver el promedio de personas que viajan a un país determinado en un dia


### *_Requerimiento 4_*

Una vez desarrollado el servicio y el repositorio, deben desarrollar los controladores para los siguientes endpoints:

- Endpoint `GET - “/ticket/getByCountry/:dest”`
- Endpoint `GET - “/ticket/getAverage/:dest”`

### *_Requerimiento 5_*

Correr el main.go y testear los endpoints con la herramienta de Postman.

---

## Run Test

*Window 1*

> Correr el servidor

```bash
$ go run cmd/server/main.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /ping                     --> main.main.func1 (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on :8080
[GIN] 2022/10/13 - 16:20:59 | 200 |      15.958µs |       127.0.0.1 | GET      "/ping"
...

```

*Window 0*

*  Asegurate de que en tu navegador esté corriendo

```bash
$ curl --write-out "%{http_code}\n" --silent --output /dev/null http://127.0.0.1:8080/ping
200
```
> Este endpoint `/ping` está diseñado para indicarte el estado de funcionamiento del servidor web

---

## Run Endpoints

*Window 1*

> Correr el servidor

```bash
$ go run cmd/server/main.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /ping                     --> desafio-goweb-juanurbano/cmd/server/handler.(*Service).Status.func1 (3 handlers)
[GIN-debug] GET    /ticket/getByCountry/:dest --> desafio-goweb-juanurbano/cmd/server/handler.(*Service).GetTicketsByCountry.func1 (3 handlers)
[GIN-debug] GET    /ticket/getAverage/:dest  --> desafio-goweb-juanurbano/cmd/server/handler.(*Service).AverageDestination.func1 (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on :8080

...

```

*Window 0*

*  Asegurate de que en tu navegador esté corriendo

```bash
$ curl http://127.0.0.1:8080/ticket/getByCountry/Brazil -X GET -vvv ; echo "" | cat -e

[{"Id":"7","Name":"Lorne Gaukrodge","Email":"lgaukrodge6@cyberchimps.com","Country":"Brazil","Time":"8:12","Price":963},{"Id":"31","Name":"Cindelyn Butterfield","Email":"cbutterfieldu@chicagotribune.com","Country":"Brazil","Time":"10:03","Price":1401},{"Id":"67","Name":"Mareah Torres","Email":"mtorres1u@ow.ly","Country":"Brazil","Time":"6:54","Price":921},{"Id":"68","Name":"Reinald M'Chirrie","Email":"rmchirrie1v@webnode.com","Country":"Brazil","Time":"7:39","Price":933},{"Id":"160","Name":"Bent Hayhow","Email":"bhayhow4f@gov.uk","Country":"Brazil","Time":"20:08","Price":1436},{"Id":"215","Name":"Wendel Feldhuhn","Email":"wfeldhuhn5y@digg.com","Country":"Brazil","Time":"13:41","Price":1006},{"Id":"221","Name":"Levey Fursey","Email":"lfursey64@tripod.com","Country":"Brazil","Time":"23:41","Price":1265},{"Id":"232","Name":"Creigh Ferretti","Email":"cferretti6f@wix.com","Country":"Brazil","Time":"19:26","Price":1457},{"Id":"236","Name":"Debee Jaimez","Email":"djaimez6j@delicious.com","Country":"Brazil","Time":"19:13","Price":1082},{"Id":"252","Name":"Aloysia Pre","Email":"apre6z@themeforest.net","Country":"Brazil","Time":"4:14","Price":542},{"Id":"325","Name":"Derry Eccleshare","Email":"deccleshare90@spotify.com","Country":"Brazil","Time":"16:52","Price":1458},{"Id":"331","Name":"Dionisio Crother","Email":"dcrother96@ox.ac.uk","Country":"Brazil","Time":"4:10","Price":797},{"Id":"339","Name":"Dom MacPaden","Email":"dmacpaden9e@answers.com","Country":"Brazil","Time":"10:45","Price":1422},{"Id":"399","Name":"Danette Tran","Email":"dtranb2@deviantart.com","Country":"Brazil","Time":"2:38","Price":1062},{"Id":"421","Name":"Melany Jasiak","Email":"mjasiakbo@slate.com","Country":"Brazil","Time":"1:12","Price":1051},{"Id":"469","Name":"Frazer Lauritzen","Email":"flauritzend0@dmoz.org","Country":"Brazil","Time":"4:12","Price":904},{"Id":"486","Name":"Bibbye Warlawe","Email":"bwarlawedh@berkeley.edu","Country":"Brazil","Time":"20:23","Price":894},{"Id":"506","Name":"Tomkin Kleint","Email":"tkleinte1@aol.com","Country":"Brazil","Time":"16:54","Price":1212},{"Id":"537","Name":"Ariella Bramont","Email":"abramontew@spotify.com","Country":"Brazil","Time":"4:51","Price":714},{"Id":"561","Name":"Bunnie Yewdale","Email":"byewdalefk@spotify.com","Country":"Brazil","Time":"22:55","Price":1309},{"Id":"614","Name":"Staffard Coverly","Email":"scoverlyh1@google.es","Country":"Brazil","Time":"10:39","Price":516},{"Id":"635","Name":"Sandra Jados","Email":"sjadoshm@mapy.cz","Country":"Brazil","Time":"13:14","Price":552},{"Id":"640","Name":"Elmo Molden","Email":"emoldenhr@photobucket.com","Country":"Brazil","Time":"5:45","Price":1249},{"Id":"656","Name":"Perceval Gunthorpe","Email":"pgunthorpei7@google.com.hk","Country":"Brazil","Time":"9:10","Price":1432},{"Id":"676","Name":"Rolland Bennetts","Email":"rbennettsir@mozilla.org","Country":"Brazil","Time":"20:25","Price":1050},{"Id":"686","Name":"Hurley Goodin","Email":"hgoodinj1@pbs.org","Country":"Brazil","Time":"9:15","Price":946},{"Id":"717","Name":"Miles Jouandet","Email":"mjouandetjw@tinypic.com","Country":"Brazil","Time":"2:10","Price":870},{"Id":"718","Name":"Daveen Yerrell","Email":"dyerrelljx@marriott.com","Country":"Brazil","Time":"9:26","Price":937},{"Id":"728","Name":"Merrile Glasson","Email":"mglassonk7@alexa.com","Country":"Brazil","Time":"14:16","Price":1418},{"Id":"753","Name":"Baron Boutflour","Email":"bboutflourkw@smh.com.au","Country":"Brazil","Time":"10:11","Price":809},{"Id":"758","Name":"Pierre Sive","Email":"psivel1@google.com.hk","Country":"Brazil","Time":"23:07","Price":971},{"Id":"774","Name":"Sharla Thacke","Email":"sthackelh@ebay.com","Country":"Brazil","Time":"6:15","Price":1472},{"Id":"791","Name":"Geordie Labb","Email":"glabbly@tamu.edu","Country":"Brazil","Time":"0:53","Price":1302},{"Id":"795","Name":"Breena Stephen","Email":"bstephenm2@techcrunch.com","Country":"Brazil","Time":"14:42","Price":1484},{"Id":"797","Name":"Osgood Biskupiak","Email":"obiskupiakm4@baidu.com","Country":"Brazil","Time":"20:40","Price":678},{"Id":"817","Name":"Darlene Cattle","Email":"dcattlemo@ra* Connection #0 to host 127.0.0.1 left intactmbler.ru","Country":"Brazil","Time":"21:27","Price":501},{"Id":"845","Name":"Ceciley Caukill","Email":"ccaukillng@sakura.ne.jp","Country":"Brazil","Time":"6:56","Price":1345},{"Id":"849","Name":"Melamie Bouldstridge","Email":"mbouldstridgenk@mashable.com","Country":"Brazil","Time":"9:09","Price":1087},{"Id":"870","Name":"Carmella Bealing","Email":"cbealingo5@jugem.jp","Country":"Brazil","Time":"17:11","Price":1034},{"Id":"871","Name":"Eadmund Sambidge","Email":"esambidgeo6@wired.com","Country":"Brazil","Time":"23:48","Price":1226},{"Id":"906","Name":"Dniren Ryam","Email":"dryamp5@sohu.com","Country":"Brazil","Time":"7:10","Price":1466},{"Id":"912","Name":"Damaris Pelfer","Email":"dpelferpb@qq.com","Country":"Brazil","Time":"6:38","Price":1038},{"Id":"918","Name":"Ruthi Kohen","Email":"rkohenph@infoseek.co.jp","Country":"Brazil","Time":"6:03","Price":664},{"Id":"922","Name":"Devora Wickardt","Email":"dwickardtpl@naver.com","Country":"Brazil","Time":"19:00","Price":1077},{"Id":"955","Name":"Gratia Naish","Email":"gnaishqi@microsoft.com","Country":"Brazil","Time":"0:22","Price":634}]$

```

> Información de los vuelos que se refieren a `Brazil`

---


## Run Test

*Window 1*

> Correr el servidor

```bash
$ go run cmd/server/main.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /ping                     --> main.main.func1 (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on :8080
[GIN] 2022/10/13 - 16:20:59 | 200 |      15.958µs |       127.0.0.1 | GET      "/ping"
...

```

*Window 0*


```bash
$ curl http://127.0.0.1:8080/ticket/getAverage/Brazil -X GET -vvv ; echo "" | cat -e
45$
```

> Información de los vuelos que se refieren a `Brazil`
---


