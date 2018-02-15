def gcd(a, b):
    assert a > 0 and b > 0
    if a == b:
        return a
    elif a > b:
        return gcd(b, a - b)
    else:
        return gcd(a, b - a)
