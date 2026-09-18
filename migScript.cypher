
CREATE CONSTRAINT node_id
FOR (a:Node) REQUIRE (a.id) IS NODE KEY; 

CREATE CONSTRAINT relationship_id
FOR ()-[a:relationship]-() REQUIRE (a.id) IS RELATIONSHIP KEY; 

CREATE CONSTRAINT node_name
FOR (a:Node) REQUIRE a.name IS NOT NULL;

CREATE CONSTRAINT relationship_name
FOR ()-[r:relationship]-() REQUIRE r.kind IS NOT NULL;