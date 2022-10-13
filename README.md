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



