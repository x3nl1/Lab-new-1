import time
import rustlib

def py_sum(numbers):
    return sum(x * x for x in numbers)

def rust_sum(numbers):
    return rustlib.sum_squares(numbers)

def benchmark():
    data = list(range(1, 1000000))

    t1 = time.time()
    py_sum(data)
    t2 = time.time()

    t3 = time.time()
    rust_sum(data)
    t4 = time.time()

    return (t2 - t1, t4 - t3)