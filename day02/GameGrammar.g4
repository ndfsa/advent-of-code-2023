grammar GameGrammar;

initial: game+;

game: 'Game' INT ':' turn (';' turn)*;

turn: color (',' color)*;

color: INT ID;

INT: [0-9]+;
ID: ('red'|'green'|'blue');
WS: [ \t\r\n]+ -> skip;
