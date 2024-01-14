pkg load symbolic

syms t1 t2 rx ry sx sy vx vy ux uy

eq1 = -2 * t1 - -1 * t2 == 18 - 19;
eq2 = 1 * t1 - -1 * t2 == 19 - 13;

[sol1, sol2] = solve(eq1, eq2, t1, t2)
