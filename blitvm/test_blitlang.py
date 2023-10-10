import pytest
import sys

import blitlang

# blitlang.WHITLIST_FUNCTIONS += ["test_blit.*"]

import math
from blitlang import *


def blit_is_still_python(code, blit_scope=None):
    """
    Here we run the code in python exec, then in blit_eval, then compare the `scope['result']`
    """
    print("===code")
    print(code.strip())
    print("===code")
    py_scope = {}
    blit_scope = blit_scope or {}
    exec(code, py_scope)

    print("python result", py_scope.get("result"))
    print("blit_eval output", blit_eval(code, scope=blit_scope))
    print("blit result", blit_scope.get("result"))
    assert (
        py_scope["result"] == blit_scope["result"]
    ), f'{code}\n{py_scope["result"]} != {blit_scope["result"]}'


def test_blit_comprehension_python():
    CASES = [
        """
non_flat= [ [1,2,3], [4,5,6], [7,8] ]
result = [y for x in non_flat for y in x]
""",
        """
a = 'foo'
l = [a for a in [1,2,3]]
result = a, l
""",
        """
a = {'b':['c', 'd', 'e']}
a['b'][1] =  'D'
result = a
""",
        """
a = [1,2,3,4,5,6]
a[2:5] = ['a', 'b']
result = a
""",
        """
a = [1,2,3,4,5,6,7,8]
a[::2] = [0,0,0,0]
result = a
""",
        """
a = [1,2,3,4,5,6,7,8]
del a[3:6:-1]
result = a
""",
        """
a = [1]
a = [a, a, a, a]
result = a[2]""",
        """
a = [(1,1), (2,2), (3,3)]
result = [(x,y) for x,y in a]
""",
        """
a = [1,2,3,4,5,6,7,8,9,10]
result = [x for x in a if x % 2]
""",
        """
d = {**{"a": 1}, **{"b":2}, "z": 3}
result = list(d.items())
result.sort()
""",
        """
a = {"a":1}
result = {None: 1, **a, **{"a":2, **{"b":2}}, **{"b":3}}
""",
    ]
    for code in CASES:
        blit_is_still_python(code)


def test_blit_assertions():
    code = """
assert True, "no"
result = 1"""
    blit_is_still_python(code)

    code = """
result = 1
try:
    assert False, "no"
except:
    pass
result = 2"""
    blit_is_still_python(code)

    code = """
result = 1
try:
    assert False
except:
    pass
result = 2"""
    blit_is_still_python(code)


def test_assignment():
    code = """
a = b,c = 1,2
result = a,b,c """
    blit_is_still_python(code)


def test_starred():
    code = """
a, *b, c,(*d,e) = (1,2,3,4,5,6,7,8,(9,10,))
result = a, b, c, d, e"""
    blit_is_still_python(code)


def test_augassign():
    for operator in ["+=", "-=", "/=", "//=", "%=", "!=", "*=", "^="]:
        code = f"""
result = 110
result {operator} 21
    """
        blit_is_still_python(code)


def test_blit_delete():
    code = """
a = 1
b = [a, a]
del a
result = [b, b]"""
    blit_is_still_python(code)


def test_blit_kwargs():
    code = """
def foo(a,b,c,d,e,f,g):
    return (a,b,c,d,e,f,g)

result = foo(1, *[2,3], *[4], e=5, **{'f':6}, **{'g':7} )
"""
    blit_is_still_python(code)


def test_args():
    code = """
def standard_arg(arg):
    return arg
result = standard_arg("a") + standard_arg(arg="b")"""
    blit_is_still_python(code)

    code = """
def standard_arg(arg=1):
    return arg
result = standard_arg("a") + standard_arg(arg="b")"""
    blit_is_still_python(code)


@pytest.mark.skipif(sys.version_info < (3, 8), reason="requires python3.8 or higher")
def test_pos_args():
    code = """
def pos_only_arg(arg, /):
    return arg

result = pos_only_arg("a")"""
    blit_is_still_python(code)

    code = """
def pos_only_arg(arg=1, /):
    return arg

result = pos_only_arg("a")"""
    blit_is_still_python(code)


def test_kwd_only_arg():
    code = """
def kwd_only_arg(*, arg):
    return arg

result = kwd_only_arg(arg=42)"""
    blit_is_still_python(code)

    code = """
def kwd_only_arg(*, arg="foo"):
    return arg

result = kwd_only_arg(arg=42)"""
    blit_is_still_python(code)


@pytest.mark.skipif(sys.version_info < (3, 8), reason="requires python3.8 or higher")
def test_func_combined():
    code = """
def combined_example(pos_only, /, standard, *, kwd_only, **kwgs):
    return (pos_only, standard, kwd_only, kwgs)
result = combined_example(1, 2, kwd_only=3, foo=[1,2,3])"""
    blit_is_still_python(code)


def test_blit_lambda():
    code = """
foo = lambda a,b,c,d,e,f,g: (a,b,c,d,e,f,g)
result = foo(1, *[2,3], *[4], e=5, **{'f':6}, **{'g':7} )
"""
    blit_is_still_python(code)


def test_fstring():
    code = """
result = f'{"😇"!a:^40}'
"""
    blit_is_still_python(code)


def test_blit_is_python_closure():
    code = """
a = 0
def nest(b):
    a = 4
    def _inner(c):
        a = 5
        def __inner(d):
            def ___inner(a):
                # shadow a
                return a,b,c,d
            return ___inner
        return __inner
    return _inner
result = nest(1)(2)(3)(4)
result
"""
    blit_is_still_python(code)


def test_blit_is_python_ifs():
    blit_is_still_python(
        """
s = 'the answer is '
def foo(i):
    if i < 0:
        res = s + 'too low'
    elif i > 0:
        res = s + 'too high'
    else:
        res = s + 'just right'
    return res

result = [foo(i) for i in [-1,0,1]]
            """
    )


def test_blit_is_python_while():
    blit_is_still_python(
        """
result = []
n = 10
while n:
    n = n -1
    result = result + [n]
"""
    )
    blit_is_still_python(
        """
result = []
n = 0
while n:
    n = n -1
    result = result + [n]
else:
    result = 'nonono'
"""
    )
    blit_is_still_python(
        """
result = []
n = 0
while n < 10:
    n = n + 1
    result = result + [n]
    if n > 5:
        break
"""
    )

    blit_is_still_python(
        """
result = []
n = 0
while n < 10:
    n = n + 1
    result = result + [n]
    if n == 5:
        continue
"""
    )


def test_blit_is_python_for():
    blit_is_still_python(
        """
def evens(s):
    out = ''
    try:
        for i, c in s:
            if i % 2 == 0:
                out = out + c
        else:
            out = -1
    except:
        return False
    return out

result = [evens(s) for s in
[[(0, 'a'),
 (1, 'b'),
 (2, 'c'),
 (3, 'd'),
 (4, 'e'),
 (5, 'f'),
 (6, 'g'),
 (7, 'h')],[], 'asdf'] ]
"""
    )
    blit_is_still_python(
        """
result = []
for n in [1,2,3,4,5,6,7,8,10]:
    result = result + [n]
    if n > 5:
        break
"""
    )

    blit_is_still_python(
        """
result = []
for n in [1,2,3,4,5,6,7,8,10]:
    result = result + [n]
    if n == 5:
        continue
"""
    )


def test_decorators():
    blit_is_still_python(
        """
result = []
a = 12
def my_decorator(log):
    a = 13
    result.append(f"d1 {log} {a}")
    def _wrapper(func):
        a = 14
        result.append(f"_wrapper {log} {a}")
        def _inner(*args, **kwargs):
            a = 15
            result.append(f"inner {log} {a}")
            func(*args, **kwargs)

        a = 16
        return _inner
    a = 17
    return _wrapper
a = 18

@my_decorator(1)
@my_decorator(2)
def say(word=None):
    result.append(word or "hi world!")
say()
say('ola mundo')
"""
    )


def test_blit_is_python_try():
    blit_is_still_python(
        """
def foo(a):
    res = 'try|'
    try:
        try:
            res = res + str(1/a) + "|"
        except ArithmeticError as e:
            res= res + 'ArithmeticError|'
    except Exception as e2:
        return 'oops|'
    else:
        res = res + 'else|'
    finally:
        res = res + 'fin|'
    return a, res
result = [foo(i) for i in [-1,0,1, 'a']]
""",
        blit_scope={"Exception": Exception, "ArithmeticError": ArithmeticError},
    )


def problem_func(*args, **kwargs):
    return [type]


blitlang.WHITELIST_FUNCTIONS.add("test_blitlang.problem_func")

EXCEPTION_CASES = [
    (
        "*a, *b = c",
        {},
        'BlitRuntimeError(SyntaxError(',
    ),
    (
        "*a, b, c = [1]",
        {},
        "BlitRuntimeError(ValueError('not enough values to unpack (expected at least 2, got 1)'))",
    ),
    ("nope", {}, "BlitRuntimeError(NameError(\"\'nope\' is not defined\"))"),
    (
        "a=1; a.b",
        {},
        "BlitRuntimeError(AttributeError(\"\'int\' object has no attribute \'b\'\"))",
    ),
    ("1/0", {}, "BlitRuntimeError(ZeroDivisionError('division by zero'))"),
    (
        "len(str(10000 ** 10001))",
        {},
        "BlitRuntimeError(MemoryError(\"Sorry! I don\'t want to evaluate 10000 ** 10001\"))",
    ),
    (
        "'aaaa' * 200000",
        {},
        "BlitRuntimeError(MemoryError('Sorry, I will not evalute something that long.'))",
    ),
    (
        "200000 * 'aaaa'",
        {},
        "BlitRuntimeError(MemoryError('Sorry, I will not evalute something that long.'))",
    ),
    (
        "(10000 * 'world!') + (10000 * 'world!')",
        {},
        "BlitRuntimeError(MemoryError('Sorry, adding those two together would make something too long.'))",
    ),
    (
        "4 @ 3",
        {},
        "BlitRuntimeError(NotImplementedError('Sorry, MatMult is not available in this evaluator'))",
    ),
    (
        "",
        {"open": open},
        # dfferences in how python version report the error
        "DangerousValue",
    ),
    (
        "a @= 3",
        {},
        "BlitRuntimeError(NotImplementedError('Sorry, MatMult is not available in this evaluator'))",
    ),
    ("int.mro()", {}, 'BlitRuntimeError(DangerousValue'),
    (repr("a" * 100001), {}, 'BlitRuntimeError(MemoryError'),
    (
        repr({hex(i): 1 for i in range(1000001)}),
        {},
        "BlitRuntimeError(MemoryError('Dict in statement is too long!'))",
    ),
    (
        repr(tuple(range(1000001))),
        {},
        "BlitRuntimeError(MemoryError('Tuple in statement is too long!'))",
    ),
    (
        repr(set(range(1000001))),
        {},
        "BlitRuntimeError(MemoryError('Set in statement is too long!'))",
    ),
    ("b'" + ("a" * 100001) + "'", {}, 'BlitRuntimeError(MemoryError('),
    (("1" + "0" * blitlang.MAX_STRING_LENGTH), {}, 'BlitRuntimeError(MemoryError('),
    (
        repr(list("a" * 100001)),
        {},
        "BlitRuntimeError(MemoryError('List in statement is too long!'))",
    ),
    ("1()", {}, "BlitRuntimeError(TypeError('Sorry, int type is not callable'))"),
    (
        "problem_func()[0]()",
        {"problem_func": problem_func},
        "BlitRuntimeError(DangerousValue('This function is forbidden: builtins.type'))",
    ),
    ("a[1]", {"a": []}, "BlitRuntimeError(IndexError('list index out of range'))"),
    (
        "a.__length__",
        {"a": []},
        "BlitRuntimeError(NotImplementedError('Sorry, access to this attribute is not available. (__length__)'))",
    ),
    (
        "'say{}'.format('hi') ",
        {},
        "BlitRuntimeError(DangerousValue('Sorry, this method is not available. (str.format)'))",
    ),
    (
        "class A: 1",
        {},
        "BlitRuntimeError(NotImplementedError('Sorry, ClassDef is not available in this evaluator'))",
    ),
    (
        "a.b",
        {"a": object()},
        "BlitRuntimeError(AttributeError(\"\'object\' object has no attribute \'b\'\"))",
    ),
    (
        "'a' + 1",
        {},
        "TypeError",
    ),
    (
        "import non_existant",
        {},
        "BlitRuntimeError(ModuleNotFoundError('non_existant'))",
    ),
    (
        "from nowhere import non_existant",
        {},
        "BlitRuntimeError(ModuleNotFoundError('nowhere'))",
    ),
    ('assert False, "no"', {}, "BlitRuntimeError(AssertionError('no'))"),
    (
        "del a,b,c",
        {},
        "BlitRuntimeError(NotImplementedError('Sorry, cannot delete 3 targets.'))",
    ),
    (
        "del a.c",
        {},
        "BlitRuntimeError(NotImplementedError('Sorry, cannot delete Attribute'))",
    ),
    (
        "[1,2,3][[]]",
        {},
        "BlitRuntimeError(TypeError('list indices must be integers or slices, not list'))",
    ),
    (
        "1<<1",
        {},
        "BlitRuntimeError(NotImplementedError('Sorry, LShift is not available in this evaluator'))",
    ),
    ("assert False", {}, "BlitRuntimeError(AssertionError())"),
    ("assert False, 'oh no'", {}, "BlitRuntimeError(AssertionError('oh no'))"),
    (
        "(a for a in a)",
        {},
        "BlitRuntimeError(NotImplementedError('Sorry, GeneratorExp is not available in this evaluator'))",
    ),
    (
        "vars(object)",
        {"vars": vars},
        "DangerousValue(\"This function 'vars' in scope is <built-in function vars> and is in DISALLOW_FUNCTIONS\")",
    ),
    (
        "a, b = 1",
        {},
        "BlitRuntimeError(TypeError('cannot unpack non-iterable int object'))",
    ),
    (
        "a, b = 1, 2, 3",
        {},
        "BlitRuntimeError(ValueError('too many values to unpack (expected 2)'))",
    ),
    (
        "a, b, c = 1, 2",
        {},
        "BlitRuntimeError(ValueError('not enough values to unpack (expected 3, got 2)'))",
    ),
    (
        "f'{1:<1000}'",
        {},
        "BlitRuntimeError(MemoryError('Sorry, this format width is too long.'))",
    ),
    (
        "f'{1/3:.1000}'",
        {},
        "BlitRuntimeError(MemoryError('Sorry, this format precision is too long.'))",
    ),
    (
        "f'{1:a}'",
        {},
        "Unknown format code",
    ),
]


@pytest.mark.filterwarnings("ignore::SyntaxWarning")
def test_exceptions():
    for i, (code, scope, ex_repr) in enumerate(EXCEPTION_CASES):
        try:
            out = blit_eval(code, scope=scope)
        except Exception as e:
            exc = e
            assert ex_repr in repr(
                exc
            ), f"{repr(repr(exc))}\nFailed {code} \nin CASE {i}"
            continue
        pytest.fail("{}\nneeded to raise: {}\nreturned: {}".format(code, ex_repr, out))


@pytest.mark.skipif(sys.version_info >= (3, 8), reason="Old way of checking functions")
def test_old_dangerous_values():
    with pytest.raises(blitlang.DangerousValue) as excinfo:
        blit_eval("a", scope={"a": {}.keys})
    assert (
        repr(excinfo.value)
        == "DangerousValue(\"This function 'a' in scope might be a bad idea.\")"
    )


def test_smoketests():
    CASES = [
        ("1 + 1", [2]),
        ("1 and []", [[]]),
        ("None or []", [[]]),
        ("3 ** 3", [27]),
        ("len(str(1000 ** 1000))", [3001]),
        ("True != False", [True]),
        ("None is None", [True]),
        ("True is not None", [True]),
        ("'a' in 'abc'", [True]),
        ("'d' not in 'abc'", [True]),
        ("- 1 * 2", [-2]),
        ("True or False", [True]),
        ("1 > 2 > 3", [False]),
        ("'abcd'[1]", ["b"]),
        ("'abcd'[1:3]", ["bc"]),
        ("'abcd'[:3]", ["abc"]),
        ("'abcd'[2:]", ["cd"]),
        ("'abcdefgh'[::3]", ["adg"]),
        ("('abc' 'xyz')", ["abcxyz"]),
        ("(b'abc' b'xyz')", [b"abcxyz"]),
        ("f'{1 + 2}'", ["3"]),
        ("{'a': 1}['a']", [1]),
        (repr([1] * 100), [[1] * 100]),
        (repr(set([1, 2, 3, 3, 3])), [set([1, 2, 3])]),
        ("[a + 1 for a in [1,2,3]]", [[2, 3, 4]]),
        ("[a + 1 for a in [1,2,3]]", [[2, 3, 4]]),
        ("{'a': 1}.get('a')", [1]),
        ("{'a': 1}.items()", [{"a": 1}.items()]),
        ("{'a': 1}.keys()", [{"a": 1}.keys()]),
        ("list({'a': 1}.values())", [[1]]),
        ("[a for a in [1,2,3,4,5,6,7,8,9,10] if a % 2 if a % 5]", [[1, 3, 7, 9]]),
    ]
    for code, out in CASES:
        assert blit_eval(code) == out, f"{code} should equal {out}"
        # Verify evaluates same as python
        assert [eval(code)] == out, code


def test_call_stack():
    scope = {}
    blit_eval("def foo(x): return x, x > 0 and foo(x-1) or 0", scope=scope)
    scope["foo"](3)
    with pytest.raises(BlitRuntimeError) as excinfo:
        scope["foo"](50)
    assert (
        repr(excinfo.value)
        == "BlitRuntimeError(\"RecursionError('Sorry, stack is to large')\")"
    )
    # test fake call stack
    blit_eval(
        "def foo(x): return foo(x - 1) if x > 0 else 0",
        scope=scope,
        call_stack=30 * [1],
    )
    with pytest.raises(BlitRuntimeError) as excinfo:
        scope["foo"](3)
    assert (
        repr(excinfo.value)
        == "BlitRuntimeError(\"RecursionError('Sorry, stack is to large')\")"
    )


def test_settings():
    code = """
i=0
a=[]
while i < 10:
    a=[a, a]
    i+=1
"""
    orig = blitlang.MAX_NODE_CALLS
    with pytest.raises(BlitRuntimeError) as excinfo:
        scope = {}
        blitlang.MAX_NODE_CALLS = 20
        blit_eval(code, scope=scope)
    assert (
        repr(excinfo.value)
        == "BlitRuntimeError(\"TimeoutError('This program has too many evaluations')\")"
    )
    blitlang.MAX_NODE_CALLS = orig

    orig = blitlang.MAX_SCOPE_SIZE
    with pytest.raises(BlitRuntimeError) as excinfo:
        scope = {}
        blitlang.MAX_SCOPE_SIZE = 500
        blit_eval(code, scope=scope)
    assert (
        repr(excinfo.value)
        == "BlitRuntimeError(\"MemoryError('Scope has used too much memory')\")"
    )
    blitlang.MAX_SCOPE_SIZE = orig


def test_importing():
    assert (
        blit_eval("import a as c; c", module_dict={"a": {"b": "123"}})[
            1
        ].__class__.__name__
        == "module"
    )
    assert blit_eval("from a import b as c; c", module_dict={"a": {"b": "123"}}) == [
        None,
        "123",
    ]
    with pytest.raises(BlitRuntimeError) as excinfo:
        blit_eval("from a import d", module_dict={"a": {"b": "123"}})
    assert repr(excinfo.value) == "BlitRuntimeError(\"ImportError('d')\")"


def test_dissallowed_functions():
    blit_eval("", scope={"thing": {}})
    with pytest.raises(DangerousValue):
        blit_eval("", scope={"open": open})


def test_undefined():
    with pytest.raises(BlitRuntimeError, match="NameError"):
        blit_eval("a += 3")


def test_undefined_local():
    with pytest.raises(BlitRuntimeError, match="UnboundLocalError"):
        blit_eval(
            """
a =1
def foo():
    a += 3
    return a
foo()
"""
        )


def test_return_nothing():
    assert (
        blit_eval(
            """
def foo():
    return
foo()"""
        )
        == [None, None]
    )


def test_exception_variable_assignment():
    blit_is_still_python(
        """
e = 1
try:
    try: 1/0
    except ZeroDivisionError as e: pass
    e
except NameError as e2:
    result = True
"""
    )


def test_return_in_exception():
    assert (
        blit_eval(
            """
def foo():
    try:
        return 1
    except Exception as e:
        return e
foo()"""
        )
        == [None, 1]
    )


def test_eval_keyword():
    assert (
        blit_eval(
            """
def foo(a,b):
    return a,b
foo(1,b=2)"""
        )
        == [None, (1, 2)]
    )


def test_eval_functiondef_does_nothing():
    assert (
        blit_eval(
            """
def foo(a,b):
    pass
foo(1,b=2)"""
        )
        == [None, None]
    )


def test_eval_joinedstr():
    with pytest.raises(BlitRuntimeError):
        blitlang.MAX_SCOPE_SIZE = 10000000
        blit_eval(
            """
a='a' * 50000
f"{a} {a}"
"""
        )
    blitlang.MAX_SCOPE_SIZE = 100000

    assert (
        blit_eval(
            """
width = 10
precision = 4
value = 12.345
f"result: {value:{width}.{precision}}"  # nested fields
"""
        )[-1]
        == "result:      12.35"
    )


def test_coverage():
    src = """
def bar():
    return untested_variable

def foo(a):
    b = 1
    return (b if a else 2)

def test_foo():
    [foo(i) for i in [False, [], 0]]

    """
    coverage = blit_test_coverage(src)

    assert (
        ascii_format_coverage(coverage, src)
        == """Missing Return on line: 3 col: 4
    return untested_variable
----^
Missing Name on line: 3 col: 11
    return untested_variable
-----------^
Missing Name on line: 7 col: 12
    return (b if a else 2)
------------^
87% coverage
"""
    )
