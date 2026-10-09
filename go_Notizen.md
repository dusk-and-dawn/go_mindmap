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

## Session 9

Aufräumen der Modelle - hier Namensgebung anpassen, keine ungleichen Namen für gleichen Inhalt (Name - Content). 
Farbe und zwei Nodes via Endpunkt (POST) vergeben. Dafür ein neues Struct erstellen Namens HTTPRelationship mit color, und den beiden zu verbindenden IDs. 
Optional: noch X, Y hinzufügen. 

## Session 10 

Wie queried man relationships richtig - rausfinden. 

## Session 11 
Script aufsetzen für Constraints - einmalig ausführen --> Erledigt 
    Inhalt: Für nodes node ID als requirement 
            Für relationships relationship ID als requirement 
Mit den Constraints arbeiten: https://neo4j.com/docs/cdc/current/procedures/elementids-key-properties/ 
Dafür brauchen wir ein Doover der DB und müssen sämtliche POST funktionen darauf anpassen, dass sie die ID welche wir selbst anlegen mitgeben  
Neue Query die sowohl nodes als auch relationships zurück gibt 
Daraus dann den Graphen laden 

## Session 12 
Frontend selbst schreiben -> ah doch nicht 

## Session 13 
Relationship connection auch auf eigen generierte ID umstellen - backend und frontend 


## Insgesamt offenes 
unit tests 
selector fuer mehrere Mindmaps 
Error handling (error wrapping und aussagekraeftige Namensgebung) 
einheitliche Namenskonventionen  