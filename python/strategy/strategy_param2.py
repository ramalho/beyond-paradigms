# strategy.py
# Strategy pattern -- function-based implementation

"""
# BEGIN STRATEGY_TESTS

    >>> joe = Customer('John Doe', 0)  # <1>
    >>> ann = Customer('Ann Smith', 1100)
    >>> cart = [LineItem('banana', 4, .5),
    ...         LineItem('apple', 10, 1.5),
    ...         LineItem('watermellon', 5, 5.0)]
    >>> Order(joe, cart, FidelityPromo(10))  # <2>
    <Order total: 42.00 due: 42.00>
    >>> Order(ann, cart, FidelityPromo(10))
    <Order total: 42.00 due: 37.80>
    >>> banana_cart = [LineItem('banana', 30, .5),
    ...                LineItem('apple', 10, 1.5)]
    >>> Order(joe, banana_cart, BulkItemPromo(10))  # <3>
    <Order total: 30.00 due: 28.50>
    >>> long_order = [LineItem(str(item_code), 1, 1.0)
    ...               for item_code in range(10)]
    >>> Order(joe, long_order, LargeOrderPromo(7))
    <Order total: 10.00 due: 9.30>
    >>> Order(joe, cart, LargeOrderPromo(7))
    <Order total: 42.00 due: 42.00>

# END STRATEGY_TESTS
"""
# BEGIN STRATEGY

from collections import namedtuple

Customer = namedtuple('Customer', 'name fidelity')


class LineItem:

    def __init__(self, product, quantity, price):
        self.product = product
        self.quantity = quantity
        self.price = price

    def total(self):
        return self.price * self.quantity


class Order:  # the Context

    def __init__(self, customer, cart, promotion=None):
        self.customer = customer
        self.cart = list(cart)
        self.promotion = promotion

    def total(self):
        if not hasattr(self, '__total'):
            self.__total = sum(item.total() for item in self.cart)
        return self.__total

    def due(self):
        if self.promotion is None:
            discount = 0
        else:
            discount = self.promotion(self)  # <1>
        return self.total() - discount

    def __repr__(self):
        fmt = '<Order total: {:.2f} due: {:.2f}>'
        return fmt.format(self.total(), self.due())

# <2>

class FidelityPromo():
    """discount for customers with 1000 or more fidelity points"""

    def __init__(self, percent):
        self.percent = percent

    def __call__(self, order):    
        return order.total() * self.percent/100.0 if order.customer.fidelity >= 1000 else 0


class BulkItemPromo():
    """discount for each LineItem with 20 or more units"""

    def __init__(self, percent):
        self.percent = percent

    def __call__(self, order):    
        discount = 0
        for item in order.cart:
            if item.quantity >= 20:
                discount += item.total() * self.percent/100.0
        return discount


class LargeOrderPromo():
    """discount for orders with 10 or more distinct items"""

    def __init__(self, percent):
        self.percent = percent

    def __call__(self, order):    
        distinct_items = {item.product for item in order.cart}
        if len(distinct_items) >= 10:
            return order.total() * self.percent / 100.0
        return 0

# END STRATEGY