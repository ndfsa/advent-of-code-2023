pkg load symbolic

syms x1 x2 x3 y1 y2 y3 a b c

eqn1 = a*x1^2 + b*x1 + c == y1;
eqn2 = a*x2^2 + b*x2 + c == y2;
eqn3 = a*x3^2 + b*x3 + c == y3;
eqn4 = x1 == 65;
eqn5 = y1 == 3726;
eqn6 = x2 == 65+131;
eqn7 = y2 == 33086;
eqn8 = x3 == 65+131*2;
eqn9 = y3 == 91672;

[ra, rb, rc] = solve(eqn1, eqn2, eqn3, eqn4, eqn5, eqn6, eqn7, eqn8, eqn9);

syms x y;
eqn10 = ra*x^2 + rb*x + rc == y;
eqn11 = x == 26501365;

[rx, ry] = solve(eqn10, eqn11)
