# Using NGT(Neighborhood Graph and Tree) from WebApp

## How to use
When you execute `docker run`, the NGT API server starts written by golang.

```
$ git clone https://github.com/tamanobi/ngt_webapp.git
$ docker build -t tamanobi/ngt_webapp .
$ docker run -it -p 2300:8000 tamanobi/ngt_webapp
```
