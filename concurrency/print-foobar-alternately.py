""" Using threading.Event """
from threading import Event
class FooBar:
    def __init__(self, n):
        self.n = n
        self.foo_lock = Event()
        self.bar_lock = Event()
        self.foo_lock.set()


    def foo(self, printFoo: 'Callable[[], None]') -> None:
        for i in range(self.n):
            self.foo_lock.wait()
            # printFoo() outputs "foo". Do not change or remove this line.
            printFoo()
            self.foo_lock.clear()
            self.bar_lock.set()


    def bar(self, printBar: 'Callable[[], None]') -> None:
        for i in range(self.n):
            self.bar_lock.wait()
            # printBar() outputs "bar". Do not change or remove this line.
            printBar()
            self.bar_lock.clear()
            self.foo_lock.set()


""" Using threading.Lock """
from threading import Lock
class FooBar:
    def __init__(self, n):
        self.n = n
        self.foo_lock = Lock()
        self.bar_lock = Lock()
        self.bar_lock.acquire()


    def foo(self, printFoo: 'Callable[[], None]') -> None:
        for i in range(self.n):
            self.foo_lock.acquire()
            # printFoo() outputs "foo". Do not change or remove this line.
            printFoo()
            self.bar_lock.release()


    def bar(self, printBar: 'Callable[[], None]') -> None:
        for i in range(self.n):
            self.bar_lock.acquire()
            # printBar() outputs "bar". Do not change or remove this line.
            printBar()
            self.foo_lock.release()