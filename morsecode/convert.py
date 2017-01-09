

letters = """A
B
C
D
E
F
G
H
I
J
K
L
M
N
O
P
Q
R
S
T
U
V
W
X
Y
Z
0
1
2
3
4
5
6
7
8
9"""


codes = """.-
-...
-.-.
-..
.
..-.
--.
....
..
.---
-.-
.-..
--
-.
---
.--.
--.-
.-.
...
-
..-
...-
.--
-..-
-.--
--..
-----
.----
..---
...--
....-
.....
-....
--...
---..
----."""


unformatted = (zip(codes.split('\n'), letters.split('\n')))

for i in unformatted:
    print('"{0}": "{1}",'.format(i[0], i[1]))
