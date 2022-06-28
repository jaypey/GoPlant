#include <ESP8266WiFi.h>
#include <WiFiUdp.h>

#define WIFI_SSID "******"
#define WIFI_PASS "*******"
#define UDP_PORT 8080

WiFiUDP UDP;
char packet[255];
char reply[] = "Packet received!";
const char *remoteip = "**********";
IPAddress remote;


void setup() {
  // put your setup code here, to run once:
  Serial.begin(115200);
  Serial.println();

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

    // Send return packet
    if(UDP.beginPacket(remote, UDP_PORT)){
      Serial.println("Adequate IP and Port");
      }
    UDP.write(reply);
    if(UDP.endPacket()){
      Serial.println("Packet sent successfully");
    }
    
    delay(1000);
}
