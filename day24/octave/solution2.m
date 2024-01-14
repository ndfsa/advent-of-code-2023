pkg load symbolic

syms t1 t2 t3
syms k0x k0y k0z vx vy vz
syms p10x p10y p10z v1x v1y v1z
syms p20x p20y p20z v2x v2y v2z
syms p30x p30y p30z v3x v3y v3z

eq1 = k0x + vx * t1 == 19 + -2 * t1;
eq2 = k0y + vy * t1 == 13 + 1 * t1;
eq3 = k0z + vz * t1 == 30 + -2 * t1;

eq4 = k0x + vx * t2 == 18 + -1 * t2;
eq5 = k0y + vy * t2 == 19 + -1 * t2;
eq6 = k0z + vz * t2 == 22 + -2 * t2;

eq7 = k0x + vx * t3 == 20 + -2 * t3;
eq8 = k0y + vy * t3 == 25 + -2 * t3;
eq9 = k0z + vz * t3 == 34 + -4 * t3;

[kx, ky, kz, Vx, Vy, Vz, T1, T2, T3] = solve(eq1, eq2, eq3, eq4, eq5, eq6, eq7, eq8, eq9, k0x, k0y, k0z, vx, vy, vz, t1, t2, t3)

