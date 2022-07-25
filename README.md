<img src="https://github.com/jaypey/GoPlant/blob/master/static/goplantlogo.png?raw=true" alt="drawing" width="200"/>

GoPlant is a website/software that receives UDP packets sent by ESP8266s while profiting from Go concurrency features and shows the data in real time through the Gin web framework.
Similar to Home Assistant, GoPlant plans to offer a front end to monitor the data. It is an open source IoT solution for your projects.

## API

The website offers an interface to retrieve the data received and stored in a PostgreSQL database. 

## Controllers

A controller can be linked to many sensors. The .ino provided in this repository shows the format in which the data must be sent : {SensorName}:{value}. 

## Dependencies
GoPlant requires the following packages:
- Gin
- Gorm
