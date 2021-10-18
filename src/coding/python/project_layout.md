# PYTHON Project Layout
[[A recommended layout:][layout-recommend]]
1. README.rst
1. LICENSE
1. Makefile
1. setup.py
1. requirements.txt
1. package/__init__.py
1. package/*.py
1. docs/conf.py (`shpinx-quickstart`)
1. docs/index.rst
1. test/test_*.py ([pytest][pytest-howto])

---

## Sample Apache-2.0 license
[requests/LICENSE](https://github.com/psf/requests/blob/main/LICENSE)

## Setup.py
[requests/setup.py](https://github.com/psf/requests/blob/main/setup.py)

## Makefile
[requests/Makefile](https://github.com/psf/requests/blob/main/Makefile)
```txt
.PHONY: docs
init:
    pip install -e .[socks]
    pip install -r requirements-dev.txt
test:
    detox
ci:
test-readme:
flake8:
coverage:
publish:
    pip install 'twine>=1.5.0'
    python setup.py sdist bdist_wheel
    twine upload dist/*
    rm -fr build dist .egg requests.egg-info
docs:
    cd docs && make html
```

[layout-recommend]: https://docs.python-guide.org/writing/structure/
[pytest-howto]: https://realpython.com/pytest-python-testing/