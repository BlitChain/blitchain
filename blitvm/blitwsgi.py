import sys
import json
import base64
import io
from wsgiref.simple_server import WSGIRequestHandler, WSGIServer, ServerHandler
from io import BytesIO
from blitvm_server import eval_script
from contextlib import redirect_stdout
from freezegun import freeze_time


class BetterServerHandler(ServerHandler):
    def error_output(self, environ, start_response):
        start_response(self.error_status, self.error_headers[:], sys.exc_info())
        _, exc, _ = sys.exc_info()
        lineno = getattr(exc, "lineno", None)
        col = getattr(exc, "col", None)
        return [
            f"Error: {exc} on line: {lineno} col: {col}\n\n".encode(),
            b"Logs:\n",
            buf.getvalue().encode(),
        ]


class SimpleWSGIRequestHandler(WSGIRequestHandler):
    def finish(self):
        pass

    def setup(self):
        self.rfile = self.request[0]
        self.wfile = self.request[1]

    def handle(self):
        """Handle a single HTTP request"""

        self.raw_requestline = self.rfile.readline(65537)
        if len(self.raw_requestline) > 65536:
            self.requestline = ""
            self.request_version = ""
            self.command = ""
            self.send_error(414)
            return

        if not self.parse_request():  # An error code has been sent, just exit
            return

        handler = BetterServerHandler(
            self.rfile,
            self.wfile,
            self.get_stderr(),
            self.get_environ(),
            multithread=False,
        )
        handler.os_environ = dict()  # the key pa
        handler.request_handler = self  # backpointer for logging
        handler.run(self.server.get_app())


class SimpleWSGIServer(WSGIServer):
    def setup(self):
        super().setup()

    def server_bind(self):
        """Override server_bind to store the server name."""
        self.server_port = ""
        self.server_name = "blitchain"
        self.setup_environ()

    def handle_request(self, request_text, output):
        self.input = BytesIO(request_text.encode())
        self.output = output

        return self._handle_request_noblock()

    def get_request(self):
        return (self.input, self.output), ["0.0.0.0", ""]

    def shutdown_request(self, request):
        pass

    def close_request(self, request):
        pass


if __name__ == "__main__":
    port, address, block_info, code, http_request = sys.argv[1:]

    with freeze_time(json.loads(block_info)["time"]):
        with io.StringIO() as buf, redirect_stdout(buf):
            try:
                sandbox, ret = eval_script(
                    port=port,
                    caller_address=None,
                    script_address=address,
                    block_info=block_info,
                    function_name=None,
                    json_kwargs=None,
                    extra_code="",
                    code=code,
                    json_msgs=None,
                    json_msg_results=None,
                )
                app = None
                if sandbox and (
                    app := sandbox.scope.get("wsgi", None)
                ):
                    s = SimpleWSGIServer("0.0.0.0", SimpleWSGIRequestHandler)
                    s.set_app(app)
                    output = BytesIO()
                    s.handle_request(http_request, output)
                    wsgiout = output.getvalue()
                elif app is None:
                    wsgiout = f"""HTTP/1.1 404\ncontent-type: text/plain\n\nOops! No WSGI Application defined on this Blit script.\nCheck out the documentation to learn how to host your website on-chain.\n\n{ret["stdout"]}""".encode() 
                else:
                    wsgiout = f"""HTTP/1.1 500\ncontent-type: text/plain\n\n{ret['exception']}\n\nstdout:\n{ret['stdout']}""".encode()
            except Exception as e:
                import traceback

                print(traceback.format_exc())
                wsgiout = f"""HTTP/1.1 500\ncontent-type: text/plain\n\nExc: {e}""".encode()
            out = buf.getvalue()

    import sys

    sys.stderr.write(out)
    print(base64.b64encode(wsgiout).decode(), end="")
