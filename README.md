# Go-testing
este repo empezo siendo un projecto para probar go, estoy aprendiendo este lenguaje 
pero termino siendo un proyecto que quiero deployar en Heroku en el cual permita al equipo
Fecha Libre F.C. administrar a los jugadores asi como tambien a manejar la asistencia de los partidos.

## Funcionalidad basica

La funcionalidad basica seria que cada jugador pueda loguearse a la aplicacion, una vez alli
ver los partidos futuros que tiene fecha libre y poder listarse o borrarse de dichos partidos.
A su vez, existiran usuarios "ADMIN" que podran cargar partidos como asi tambien agregar nuevos jugadores.

## Buildear el proyecto.

Buscar las dependencias

```
go mod tidy
'''

El proyecto se buildea con el siguiente commando/

```
go build
```

Ese comando bajara todas las dependencias necesarias leyendo el archivo go.mod.
Una vez buildeado se corre con el siguiente commando

```
./tincho.example
```

correrlo localmente 
```
./tincho.example conf-dev.yam
```