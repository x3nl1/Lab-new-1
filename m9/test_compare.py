from compare import py_sum, rust_sum

def test_equal():
    data = [1, 2, 3]
    assert py_sum(data) == rust_sum(data)

def test_empty():
    assert py_sum([]) == rust_sum([])

def test_negative():
    data = [-1, -2]
    assert py_sum(data) == rust_sum(data)