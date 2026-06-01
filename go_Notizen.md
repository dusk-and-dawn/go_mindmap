Hausaufgaben 

Aufschreiben was wir gemacht haben 
OpenAPI spec entwerfen und ersten Endpunkt implementieren 

## Session 1 

In der ersten Session haben wir einen http handler erstellt und in die controller.go Datei ausgelagert. 
Wir haben ein package erstellt und die .go files darunter gezogen, damit die Dateien gegenseitig alle Elemente benutzen können. 
Ausserdem erzeugen wir einen Server mit der ListenAndServe Funktion. 
Zuletzt haben wir auch noch das anlegen von Methoden unter Strukturen und deren Aufruf geübt. 

## Session 2 

HA: Statt Prefix eine Hashmap an nodes reingegeben, dann durch business logik die richtige node finden und ausgeben 

## Session 3 

HA: Post Handler implementieren, sodass eine node von Hand angelegt werden kann + SQLight db implementieren und hier nodes wegschreiben 

## Session 4 

HA: return die ID vom post endpunkt, abruf basierend auf der selben id von der db --> Hexagonale Architektur im Kopf behalten, db Model definieren und in das Domain Model uebersetzen 

## Session 5 

HA: Einmal refactoren und besser benennen plus durchdringen hier Bezug nehmen auf die Hexagonale Architektur - Achtung Domain soll die Logik ausfuehren nicht der Http adapter 

## Session 6 

Funktionalität zur Wiedergabe von eine Liste an Nodes finalisieren. Node-Verbindung via Endpunkt ermöglichen. 

## Session 7 

ORdentliche Dynamisierung der Relationship-Property-Create Funktion. 

## Session 8

Liste der Nodes im FE ausgeben nicht im Terminal (inklusive Struct Anpassung). Farbe und zwei Nodes via Endpunkt (POST) vergeben. Dafür ein neues Struct erstellen Namens HTTPRelationship mit color, und den beiden zu verbindenden IDs. 

