

Basic Usage
-----------

``blit_eval`` returns a list of all the expressions in the provided code.
Generally you care about the last one. 


To get very simple evaluating:

.. code-block:: python

    from blitlang import blit_eval

    blit_eval("'Hi!' + ' world!'")

returns ``[Hi! World!]``.

Expressions can be as complex and convoluted as you want:

.. code-block:: python

    blit_eval("21 + 19 / 7 + (8 % 3) ** 9")

returns ``[535.714285714]``.

You can add your own functions in as well.

.. code-block:: python

    blit_eval("square(11)", scope={"square": lambda x: x*x})

returns ``[121]``.

