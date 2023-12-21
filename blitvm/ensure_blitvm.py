import blitlang
import blitwsgi

if __name__ == "__main__":
    code = '''assert str(obj) == '<object object at 0x1234>', "blit-python not found on PATH"'''

    blitlang.blit_eval(code, scope={"obj": object()})
