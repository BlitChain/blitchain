from setuptools import setup

__version__ = "0.0.1"

setup(
    name="blitlang",
    py_modules=["blitlang"],
    version=__version__,
    description="",
    long_description=open("README.rst", "r").read(),
    long_description_content_type="text/x-rst",
    test_suite="test_blit",
    install_requires=["pytest", "ConfigArgParse", "python-forge"],
    classifiers=[
        "Development Status :: 4 - Beta",
        "Intended Audience :: Developers",
        "Topic :: Software Development :: Libraries :: Python Modules",
        "Programming Language :: Python",
    ],
)
