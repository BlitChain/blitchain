#!/usr/bin/env python3

import ast
import random
import re
from functools import wraps

import forge
from freezegun import freeze_time

try:
    import re2

    re2.set_fallback_notification(re2.FALLBACK_EXCEPTION)
    re2_dict = {
        "ASCII": re2.ASCII,
        "BackreferencesException": re2.BackreferencesException,
        "CharClassProblemException": re2.CharClassProblemException,
        "DEBUG": re2.DEBUG,
        "DOTALL": re2.DOTALL,
        "I": re2.I,
        "IGNORECASE": re2.IGNORECASE,
        "L": re2.L,
        "LOCALE": re2.LOCALE,
        "M": re2.M,
        "MULTILINE": re2.MULTILINE,
        "RegexError": re2.RegexError,
        "S": re2.S,
        "U": re2.U,
        "UNICODE": re2.UNICODE,
        "VERBOSE": re2.VERBOSE,
        "X": re2.X,
        "compile": re2.compile,
        "contains": re2.contains,
        "count": re2.count,
        "error": re2.error,
        "escape": re2.escape,
        "findall": re2.findall,
        "finditer": re2.finditer,
        "fullmatch": re2.fullmatch,
        "match": re2.match,
        "re": re2.re,
        "search": re2.search,
        "split": re2.split,
        "sub": re2.sub,
        "subn": re2.subn,
    }
except ImportError:
    re2_dict = {}
    print("WARNING: Could not load re2")


import io
import sys
from contextlib import redirect_stdout

import requests
import simplejson
import simplejson as json

import blitlang

# This will affect the gas cost of the eval
NODE_GAS_COST = 0
SCOPE_MEMORY_COST = 0.001


def main():
    (
        port,
        caller_address,
        script_address,
        block_info,
        function_name,
        json_kwargs,
        extra_code,
        code,
    ) = sys.argv[1:]

    return eval_script(
        port,
        caller_address,
        script_address,
        block_info,
        function_name,
        json_kwargs,
        extra_code,
        code,
    )


def get_module_dict():
    import base64
    import datetime
    import decimal
    import hashlib
    import html
    import mimetypes
    import pathlib
    import random
    import string
    import typing
    import urllib

    @forge.copy(random.seed)
    def safe_random_seed(a=None, version=2):
        assert a is not None, "in Blit seed must not be None"
        return random.seed(a, version)

    safe_random_seed.__doc__ = random.Random.seed.__doc__
    safe_random_seed.__module__ = random.Random.seed.__module__
    safe_random_seed.__qualname__ = random.Random.seed.__qualname__
    allow_func(safe_random_seed)

    @forge.copy(json.dumps)
    def safe_json_dumps(**kwargs):
        if kwargs.get("separators", None):
            assert kwargs["separators"] == (
                ",",
                ":",
            ), "separators can only be (',', ': ')"
        if check_circular := kwargs.get("check_circular", None):
            assert check_circular is True, "check_circular must be True"
        if indent := kwargs.get("indent", None):
            if isinstance(indent, str):
                assert (
                    len(kwargs["indent"]) <= 4
                ), f"indent must be less than or equal 4, got: {indent}"
            if isinstance(indent, int):
                assert (
                    kwargs["indent"] <= 4
                ), f"indent must be less than or equal 4, got: {indent}"
        return json.dumps(**kwargs)

    safe_json_dumps.__doc__ = json.dumps.__doc__
    safe_json_dumps.__module__ = "simplejson"
    safe_json_dumps.__qualname__ = "dumps"
    allow_func(safe_json_dumps)

    mod_dict = {
        "datetime": {
            "date": datetime.date,
            "datetime": datetime.datetime,
            "time": datetime.time,
            "timedelta": datetime.timedelta,
            "timezone": datetime.timezone,
            "tzinfo": datetime.tzinfo,
        },
        "pathlib": {"PurePath": pathlib.PurePath},
        "mimetypes": {"guess_type": mimetypes.guess_type},
        "urllib": {
            "parse": {
                "parse_qs": urllib.parse.parse_qs,
                "parse_qsl": urllib.parse.parse_qsl,
                "urlsplit": urllib.parse.urlsplit,
                "urlunsplit": urllib.parse.urlunsplit,
                "urljoin": urllib.parse.urljoin,
                "urldefrag": urllib.parse.urldefrag,
                "quote": urllib.parse.quote,
                "quote_plus": urllib.parse.quote_plus,
                "unquote": urllib.parse.unquote,
                "unquote_plus": urllib.parse.unquote_plus,
                "unquote_to_bytes": urllib.parse.unquote_to_bytes,
                "quote_from_bytes": urllib.parse.quote_from_bytes,
            }
        },
        "base64": {
            "decodebytes": base64.decodebytes,
            "encodebytes": base64.encodebytes,
            "b64decode": base64.b64decode,
            "b64encode": base64.b64encode,
            "urlsafe_b64encode": base64.urlsafe_b64encode,
            "urlsafe_b64decode": base64.urlsafe_b64decode,
        },
        "decimal": {"Decimal": decimal.Decimal},
        "simplejson": {
            "dumps": safe_json_dumps,
            "loads": simplejson.loads,
        },
        "simplejson.errors": {
            "JSONDecodeError": simplejson.JSONDecodeError,
        },
        "json": {
            "dumps": safe_json_dumps,
            "loads": simplejson.loads,
            "JSONDecodeError": simplejson.JSONDecodeError,
        },
        "html": {"escape": html.escape, "unescape": html.unescape},
        "io": {"StringIO": io.StringIO, "BytesIO": io.BytesIO},
        "hashlib": {"sha1": hashlib.sha1, "sha256": hashlib.sha256},
        "typing": {
            "Callable": typing.Callable,
            "Any": typing.Any,
            "Union": typing.Union,
            "Optional": typing.Optional,
            "Literal": typing.Literal,
            "Annotated": typing.Annotated,
        },
        "string": {
            "Template": string.Template,
            "capwords": string.capwords,
            "ascii_letters": string.ascii_letters,
            "ascii_lowercase": string.ascii_lowercase,
            "ascii_uppercase": string.ascii_uppercase,
            "digits": string.digits,
            "hexdigits": string.hexdigits,
            "octdigits": string.octdigits,
            "punctuation": string.punctuation,
            "printable": string.printable,
            "whitespace": string.whitespace,
        },
        "random": {
            "betavariate": random.betavariate,
            "choice": random.choice,
            "choices": random.choices,
            "expovariate": random.expovariate,
            "gauss": random.gauss,
            "paretovariate": random.paretovariate,
            "random": random.random,
            "sample": random.sample,
            "shuffle": random.shuffle,
            "triangular": random.triangular,
            "seed": safe_random_seed,
            "uniform": random.uniform,
        },
        "re2": re2_dict,
    }

    def walk(node):
        for key, item in node.items():
            if isinstance(item, dict):
                walk(item)
            else:
                allow_func(item)

    walk(mod_dict)
    return mod_dict


def build_sandbox(port, caller_address, script_address, block_info):
    url = f"http://localhost:{port}/rpc"

    def _chain(chain_method, **params):
        """
        The main way to interact with the chain from a script.

        :param method: the command to call on the chain, see TxBuilder for a list of possible commands
        `**kwargs` will depend on the command being called

        :returns: the response of the command or error

        """

        # params = {snake(k): v for k, v in params.items()}
        payload = {
            "method": f"RpcService.{chain_method}",
            "params": [params],
            "jsonrpc": "2.0",
            "id": 0,
        }

        # print(f"rpc {payload}")
        res = requests.post(url, json=payload)
        # print(f"rpc response: {res.text}")
        try:
            ret_json = res.json()
            if ret_json["error"] and ret_json["error"].startswith(
                "CHAIN ERROR: types.ErrorOutOfGas"
            ):
                raise MemoryError("Out of Gas")
            try:
                # some rpc responses are json encoded
                ret_json["result"] = json.loads(str(ret_json["result"]).encode())
            except json.JSONDecodeError:
                pass
        except json.JSONDecodeError:
            ret_json = {"error": res.text, "result": None}

        if ret_json["error"] is not None:
            raise Exception(ret_json["error"])
        return ret_json["result"]

    _chain.__qualname__ = "_chain"
    allow_blit_func(_chain)

    gas_state = {
        "unconsumed_gas": 1,
        "gas_consumed": 0,
        "gas_limit": 0,
        "nodes_called": 0,
    }

    class ScopedBlitEval(blitlang.BlitEval):
        def gas_state(self):
            return gas_state

        def _eval(self, node):
            # print("scopse ", [s.scope.dicts[-1] for s in blitlang.sandboxes])
            gas_state["nodes_called"] += 1
            gas_state["unconsumed_gas"] += NODE_GAS_COST + int(
                get_scope_size() * SCOPE_MEMORY_COST
            )
            if gas_state["unconsumed_gas"] > 100000 or isinstance(node, ast.Module):
                self.consume_gas()
            if gas_state["nodes_called"] > blitlang.MAX_NODE_CALLS:
                raise MemoryError("This program has too many evaluations")
            return super()._eval(node)

        def consume_gas(self):
            resp = _chain("Consumegas", amount=gas_state["unconsumed_gas"])
            gas_state["unconsumed_gas"] = 1

            gas_state["gas_consumed"] = resp.get("GasConsumed", None)
            gas_state["gas_limit"] = resp.get("GasLimit", None)
            if (
                gas_state["gas_limit"] > 0
                and gas_state["gas_consumed"] > gas_state["gas_limit"]
            ):
                raise MemoryError("Out of Gas")

    scope = {}
    sandbox = ScopedBlitEval(
        scope=scope,
    )

    @forge.copy(print)
    def blit_print(*args, sep=" ", end="\n", file=None, flush=False):
        assert len(sep) <= 5, AssertionError("sep length greater than 5")
        assert len(end) <= 5, AssertionError("end length greater than 5")
        print(*args, end=end)

    blit_print.__doc__ = print.__doc__
    blit_print.__module__ = "builtins"
    blit_print.__qualname__ = "print"
    blit_print = allow_func(blit_print)

    sandbox.scope.dicts[0]["print"] = blit_print
    # sandbox.scope.dicts[0]["globals"] = sandbox.scope.globals
    # sandbox.scope.dicts[0]["locals"] = sandbox.scope.locals

    @allow_blit_func
    def _sendMsg(typeUrl, value=None, **kwargs):
        """
        Send a message to the chain.

        :param typeUrl: the type of message to send (camlCase to match the protobuf usage)
        :param value: the value of the message to send
        :param **kwargs: the params of the message

        Only value or kwargs can be used.

        :returns: the response of the message or error

        """
        # only value or kwargs can be used
        assert value is None or kwargs == {}, AssertionError(
            "value and kwargs cannot both be used"
        )

        if value is None:
            value = kwargs
        value["@type"] = typeUrl
        res = _chain("Msg", json_msg=json.dumps(value))
        if "data" in res:
            del res["data"]
        if "events" in res:
            res["event"] = res["events"][0]
            del res["events"]
        if "msg_responses" in res:
            res["msg_response"] = res["msg_responses"][0]
            del res["msg_responses"]
        return res

    @allow_blit_func
    def sendQuery(method, params=None, **kwargs):
        """
        Send a query to the chain.

        :param method: the query method to call
        :param params: the params of the query
        :param **kwargs: the params of the query

        Only params or kwargs can be used.

        :returns: the response of the query or error

        """
        # only params or kwargs can be used
        assert params is None or kwargs == {}, AssertionError(
            "params and kwargs cannot both be used"
        )

        if params is None:
            params = kwargs

        return _chain("Query", method=method, json_args=json.dumps(params))

    @allow_blit_func
    def get_gas_consumed():
        """
        The total amount of gas consumed so far.
        """
        return gas_state["gas_consumed"] + gas_state["unconsumed_gas"]

    @allow_blit_func
    def get_gas_limit():
        """
        The maximum amount of gas that can be used in this query or transaction
        """
        return gas_state["gas_limit"]

    @allow_blit_func
    def get_nodes_called():
        """
        The number of Python AST nodes evaluated in this query or transaction
        """
        return gas_state["nodes_called"]

    @allow_blit_func
    def get_script_address() -> str:
        """
        Returns the address of this current script.
        """
        return script_address

    @allow_blit_func
    def get_scope_size() -> int:
        """
        Returns the size of the current scope.
        """
        return sum(s.size for s in blitlang.sandboxes)

    @allow_blit_func
    def get_max_scope_size() -> int:
        """
        Returns the maximum size of the scope.
        """
        return blitlang.MAX_SCOPE_SIZE

    @allow_blit_func
    def get_caller_address() -> str:
        """
        Returns the address of the caller of this script.
        """
        return caller_address

    @allow_blit_func
    def get_block_info() -> dict:
        """
        Returns a dictionary of the current block info
        """
        return block_info

    @allow_blit_func
    def _blit_eval(
        code, scope=None, track_func=None, max_node_calls=None, max_scope_size=None
    ):
        """
        Evaluate a string of blit code.

        :param code: the code to evaluate
        :param scope: the scope to evaluate the code in
        :param track_func: a function to call after each node is evaluated use to track gas or scope size

        :returns: the result of the evaluation

        """

        sandbox = ScopedBlitEval(
            scope=scope,
            local_track_func=track_func,
            max_node_calls=max_node_calls,
            max_scope_size=max_scope_size,
        )

        result = sandbox.eval(code)
        sandbox.consume_gas()
        if not result:
            return None
        return result[-1]

    module_dict = get_module_dict()

    module_dict["blit"] = {
        "_chain": _chain,
        "_sendMsg": _sendMsg,
        "sendQuery": sendQuery,
        "get_gas_consumed": get_gas_consumed,
        "get_gas_limit": get_gas_limit,
        "get_scope_size": get_scope_size,
        "get_max_scope_size": get_max_scope_size,
        "get_script_address": get_script_address,
        "get_caller_address": get_caller_address,
        "get_block_info": get_block_info,
        "get_nodes_called": get_nodes_called,
        "_blit_eval": _blit_eval,
    }

    sandbox.modules = blitlang.make_modules(module_dict)
    return sandbox


def eval_script(
    port,
    caller_address,
    script_address,
    block_info,
    function_name,
    json_kwargs,
    extra_code,
    code,
):
    result = None
    stdout = None
    exception = None
    sandbox = None

    block_info = json.loads(block_info)
    with freeze_time(block_info["time"]):
        with io.StringIO() as buf, redirect_stdout(buf):
            try:
                sandbox = build_sandbox(
                    port,
                    caller_address,
                    script_address,
                    block_info,
                )
                sandbox.consume_gas()

                random.seed(
                    (str(block_info) or "")
                    + (caller_address or "")
                    + (script_address or "")
                    + (function_name or "")
                    + (json_kwargs or "")
                    + (extra_code or "")
                    + str(sandbox.gas_state() or "")
                )
                result = (
                    sandbox.eval(
                        code + "\n" + extra_code,
                    )
                    or [None]
                )[-1]
                scope = sandbox.scope.flatten()
                public_scope_all = scope.get(
                    "__all__",
                    [
                        k
                        for k, v in scope.items()
                        if v is not sandbox.modules.blit._chain
                        and v is not sandbox.modules.blit._sendMsg
                        and not k.startswith("_")
                        and k not in []
                    ],
                )
                if function_name:
                    if function_name in scope:
                        if function_name in public_scope_all:
                            kwargs = json.loads(json_kwargs or "{}")
                            # print("kwargs", kwargs)
                            result = scope[function_name](**kwargs)
                        else:
                            raise Exception(f"function not public: {function_name}")
                    else:
                        raise Exception(f"function not defined: {function_name}")
                # consume final gas
                sandbox.consume_gas()
            except blitlang.BlitRuntimeError as e:
                exception = e
                print("ERR:", repr(e.__context__))
            except Exception as e:
                exception = e
                print("ERROR: ", e)
            finally:
                stdout = buf.getvalue()

    if exception is not None:
        exception = {
            "class": exception.__class__.__name__,
            "context": exception.__context__.__class__.__name__,
            "msg": str(exception),
            "lineno": getattr(exception, "lineno", 0),
            "col": getattr(exception, "col", 0),
        }

    return sandbox, {
        "result": result,
        "stdout": stdout,
        "exception": exception,
        "nodes_called": sandbox.modules.blit.get_nodes_called() if sandbox else None,
        "gas_limit": sandbox.modules.blit.get_gas_limit() if sandbox else None,
        "script_gas_consumed": sandbox.modules.blit.get_gas_consumed()
        if sandbox
        else None,
    }


blitlang.WHITELIST_FUNCTIONS.update(
    [
        "datetime.datetime.isoformat",
        "datetime.datetime.fromisoformat",
        "freezegun.api.FakeDatetime.now",
        "freezegun.api.FakeDatetime.time",
        "freezegun.api.FakeDatetime.date",
        "freezegun.api.FakeDatetime.timestamp",
        "datetime.time",
        "datetime.date",
        # re2.Match.re
        "contains",
        "count",
        "findall",
        "finditer",
        "fullmatch",
        "match",
        "scanner",
        "search",
        "split",
        # str
        "str.capitalize",
        "str.casefold",
        "str.count",
        "str.encode",
        "str.endswith",
        "str.find",
        "str.index",
        "str.isalnum",
        "str.isalpha",
        "str.isascii",
        "str.isdecimal",
        "str.isdigit",
        "str.isidentifier",
        "str.islower",
        "str.isnumeric",
        "str.isprintable",
        "str.isspace",
        "str.istitle",
        "str.isupper",
        "str.join",
        "str.lower",
        "str.lstrip",
        "str.partition",
        "str.removeprefix",
        "str.removesuffix",
        "str.rfind",
        "str.rindex",
        "str.rpartition",
        "str.rsplit",
        "str.rstrip",
        "str.split",
        "str.splitlines",
        "str.startswith",
        "str.strip",
        "str.swapcase",
        "str.title",
        "str.upper",
        # list
        "list.append",
        "list.clear",
        "list.copy",
        "list.count",
        "list.extend",
        "list.index",
        "list.insert",
        "list.pop",
        "list.remove",
        "list.reverse",
        "list.sort",
        # dict
        "dict.clear",
        "dict.copy",
        "dict.fromkeys",
        "dict.get",
        "dict.items",
        "dict.keys",
        "dict.pop",
        "dict.popitem",
        "dict.setdefault",
        "dict.update",
        "dict.values",
        # Template
        "string.Template.safe_substitute",
        "string.Template.substitute",
        # random
        "Random.random",
        "random.Random.uniform",
        "random.Random.expovariate",
        "random.Random.choice",
        "random.Random.shuffle",
        "random.Random.sample",
        # wsgi
        "wsgiref.handlers.BaseHandler.start_response",
        "wsgiref.handlers.BaseHandler.write",
    ]
)


def allow_func(func=None, mod=None, name=None):
    if callable(func):
        return _allow_func(func, mod=mod, name=name)
    return func


def allow_blit_func(function):
    function.__qualname__ = function.__name__
    function.__module__ = "blit"
    return wraps(function)(_allow_func)(function)


def _allow_func(func, mod=None, name=None):
    if not callable(func):
        return func

    modname = mod or getattr(func, "__module__", None)
    qualname = name or getattr(
        func, "__qualname__", getattr(func, "__name__", getattr(func, "_name", None))
    )

    if modname:
        fullname = modname + "." + qualname
    else:
        fullname = qualname

    blitlang.WHITELIST_FUNCTIONS.add(fullname)
    return func


if __name__ == "__main__":
    _, response = main()
    print(
        json.dumps(response, sort_keys=True, default=str, separators=(",", ":")),
        end="",
    )  # "end" is important for parsing in go
    if response["exception"] is not None:
        sys.exit(1)
