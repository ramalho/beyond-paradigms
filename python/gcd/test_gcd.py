from math import gcd
from gcd1 import gcd as gcd1
from gcd2 import gcd as gcd2

pairs = set()

for i in range(1, 50):
    for j in range(1, 50):
        t = tuple(sorted((i, j)))
        if t in pairs:
            continue
        res = gcd(*t)
        #print(t, res)
        assert res == gcd1(*t)
        assert res == gcd2(*t)
        pairs.add(t)
