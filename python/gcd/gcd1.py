def gcd(a, b):
    assert a > 0 and b > 0
    while a != b:
        if a > b:
            a -= b
        else:
            b -= a
    return a
