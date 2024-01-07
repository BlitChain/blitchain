
if __name__ == "__main__":
    assert str(object()) == '<object object at 0x1234>', "blit-python not found on PATH, make sure it is installed with pyenv and first on PATH"

    import blitlang
    import blitwsgi
    code = '''assert str(obj) == '<object object at 0x1234>', "blit-python not found on PATH"'''
    blitlang.blit_eval(code, scope={"obj": object()})
