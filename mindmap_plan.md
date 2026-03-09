
# architecture
Wir implementieren eine hexagonale Architektur als Ausbaustufe 2. Nachdem eine simpe JSON Erzeugung/Auslese Version implementiert wurde.  

## frontend
Für das Frontend wird cytoscape.js (Hier)[https://js.cytoscape.org] benutzt, da dies einen Fokus auf das Backend erlaubt. 

## backend
Im Backend werde ich in go nodes als Struktur erfassen. 
Dabei werde  ich pro node die folgenden Werte etablieren:
1. ID 
2. Inhalt
3. X / Y Koordinaten auf der Oberfläche 
4. Children der Node 

Außerdem baue ich einen Mindmap Container der pro Mindmap jeweils alle Nodes enthält, dabei konzentriere ich mich auf die folgenden Elemente: 
1. Titel der Mindmap 
2. Nodes 

Meine Go API wird dann die jeweiligen Daten im JSON Format zur Verfügung stellen. Die Nodedaten müssen abgefragt werden können via GET request und geschrieben werden können via POST request. 

## database
Am Anfang lege ich meine Daten einfach als Einträge in einer SQLite Datenbank ab. 
