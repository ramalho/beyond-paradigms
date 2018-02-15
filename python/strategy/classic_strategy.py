# classic_strategy.py
# Strategy pattern — classic implementation

"""
Create two customers, with and without "fidelity points":

    >>> ann = Customer('Ann Smith', 1100)  # ➊
    >>> joe = Customer('John Doe', 0)

Create a shopping cart with some fruits:

    >>> cart = [LineItem('banana', 4, .5),  # ➋
    ...         LineItem('apple', 10, 1.5),
    ...         LineItem('watermellon', 5, 5.0)]

The `FidelityPromo` only gives a discount to Ann:

    >>> Order(joe, cart, FidelityPromo())  # ➌
    <Order total: 42.00 due: 42.00>
    >>> Order(ann, cart, FidelityPromo())  # ➍
    <Order total: 42.00 due: 39.90>

The `BulkItemPromo` gives a discount for items with 20+ units:

    >>> banana_cart = [LineItem('banana', 30, .5),  # ➎
    ...                LineItem('apple', 10, 1.5)]
    >>> Order(joe, banana_cart, BulkItemPromo())  # ➏
    <Order total: 30.00 due: 28.50>

`LargeOrderPromo` gives a discount for orders with 10+ items:

    >>> long_order = [LineItem(str(item_code), 1, 1.0) # ➐
    ...               for item_code in range(10)]
    >>> Order(joe, long_order, LargeOrderPromo())  # ➑
    <Order total: 10.00 due: 9.30>
    >>> Order(joe, cart, LargeOrderPromo())
    <Order total: 42.00 due: 42.00>

"""

from abc import ABC, abstractmethod
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
            discount = self.promotion.discount(self)
        return self.total() - discount

    def __repr__(self):
        fmt = '<Order total: {:.2f} due: {:.2f}>'
        return fmt.format(self.total(), self.due())


class Promotion(ABC):  # the Strategy: an Abstract Base Class

    @abstractmethod
    def discount(self, order):
        """Return discount as a positive dollar amount"""


class FidelityPromo(Promotion):  # first Concrete Strategy
    """5% discount for customers with 1000 or more fidelity points"""

    def discount(self, order):
        return order.total() * .05 if order.customer.fidelity >= 1000 else 0


class BulkItemPromo(Promotion):  # second Concrete Strategy
    """10% discount for each LineItem with 20 or more units"""

    def discount(self, order):
        discount = 0
        for item in order.cart:
            if item.quantity >= 20:
                discount += item.total() * .1
        return discount


class LargeOrderPromo(Promotion):  # third Concrete Strategy
    """7% discount for orders with 10 or more distinct items"""

    def discount(self, order):
        distinct_items = {item.product for item in order.cart}
        if len(distinct_items) >= 10:
            return order.total() * .07
        return 0
