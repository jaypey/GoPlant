#include <ESP8266WiFi.h>
#include <WiFiUdp.h>

#define WIFI_SSID "*"
#define WIFI_PASS "*"
#define UDP_PORT 8080
#define SENSOR_ID "1"

WiFiUDP UDP;
char packet[255];
char reply[] = "";
const char *remoteip = "*";
IPAddress remote;

int sensorValue;
int sensorPin = A0;
int powerPin = 5;
int limit = 300;
int latestValue = -1;

void setup() {
  // put your setup code here, to run once:
  Serial.begin(115200);
  Serial.println();

  //Humidity sensor
  pinMode(powerPin, OUTPUT);
  digitalWrite(powerPin, LOW);
  
  //Begin WiFi
  WiFi.begin(WIFI_SSID, WIFI_PASS);

  // Connecting to WiFi...
  Serial.print("Connecting to ");
  Serial.print(WIFI_SSID);
  // Loop continuously while WiFi is not connected
  while (WiFi.status() != WL_CONNECTED)
  {
    delay(100);
    Serial.print(".");
  }

    // Connected to WiFi
  Serial.println();
  Serial.print("Connected! IP address: ");
  Serial.println(WiFi.localIP());

  if (remote.fromString(remoteip)) { // try to parse into the IPAddress
      Serial.println(remote);
      Serial.println("Ready to receive from ip");// print the parsed IPAddress 
  } else {
      Serial.println("UnParsable IP");
  }
  Serial.print("Receiving on UDP port ");
  Serial.println(UDP_PORT);
}

void loop() {

    latestValue = logSoilMoistureSensor();
    // Send return packet
     char latestValueString[5];
     sprintf(latestValueString, "%d", latestValue);
     strcat(reply, SENSOR_ID);
     strcat(reply,"SoilMoist:");
     strcat(reply, latestValueString);
    
    if(UDP.beginPacket(remote, UDP_PORT)){
      Serial.println("Adequate IP and Port");
      }
    UDP.write(reply);
    Serial.println(reply);
    if(UDP.endPacket()){
      Serial.println("Packet sent successfully");
    }
    
    reply[0] = 0;
    delay(60000);
}

int logSoilMoistureSensor(){
    digitalWrite(powerPin, HIGH);
    delay(100);
    sensorValue = analogRead(sensorPin);
    digitalWrite(powerPin, LOW);
    sensorValue = map(sensorValue,420,50,0,100);
    return sensorValue;
}
