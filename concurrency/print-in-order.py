""" The simple stupid slowest way
using shared variables as semaphores
class Foo:
    def __init__(self):
        self.second_sema = False
        self.third_sema = False
        pass


    def first(self, printFirst: 'Callable[[], None]') -> None:
        # printFirst() outputs "first". Do not change or remove this line.
        printFirst()
        self.second_sema = True


    def second(self, printSecond: 'Callable[[], None]') -> None:
        while not self.second_sema:
            pass
        # printSecond() outputs "second". Do not change or remove this line.
        printSecond()
        self.third_sema = True


    def third(self, printThird: 'Callable[[], None]') -> None:
        while not self.third_sema:
            pass
        # printThird() outputs "third". Do not change or remove this line.
        printThird()
"""


from threading import Lock

class Foo:
    """ This solution uses thread locks (threading.Lock)
    40ms / 14MB / Accepted
    """
    def __init__(self):
        self.second_locs = Lock()
        self.third_lock = Lock()
        self.second_locs.acquire()
        self.third_lock.acquire()
        pass


    def first(self, printFirst: 'Callable[[], None]') -> None:
        # printFirst() outputs "first". Do not change or remove this line.
        printFirst()
        self.second_locs.release()


    def second(self, printSecond: 'Callable[[], None]') -> None:
        self.second_locs.acquire()
        # printSecond() outputs "second". Do not change or remove this line.
        printSecond()
        self.third_lock.release()


    def third(self, printThird: 'Callable[[], None]') -> None:
        self.third_lock.acquire()
        # printThird() outputs "third". Do not change or remove this line.
        printThird()