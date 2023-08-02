# KUTI

Kuti is a static site generator written in go, aims to facilate the flow of buildings mstatic sites. It allow you to use markdown for your contents and golang template language for your design, at the end, a simple yml file does the routing for you.

## Development mode
```bash
kuti serve
```

## Build html, css output
```bash
kuti build
```


## Simple blog yaml file
```yml
name: "Blog of debuqer"
author:
  name: "MohammadBagher Abbasi"
  username: "debuqer"
  description: "PHP Developer, trying to focus on DEBUGGING"
  profile: "assets/img/profile.jpeg"
server:
  host: "localhost"
  port: "3334"
  url: ""
  ext: "index.html"
template:
  dir: "templates/"
source:
  dir: "source/"
  ext: "md"
routes:
  "/":
    type: "index"
    dir: ""
    template: "index.html"
  "/about":
    type: "index"
    dir: ""
    template: "about.html"
  "/cv":
    type: "index"
    dir: ""
    template: "cv.html"
  "/blog":
    type: "index"
    dir: ""
    template: "blog.html"
  "/blog/post/:filename":
    type: "post"
    dir: "/blog"
    template: "article.html"
```


## Template 

Templates should be designed by go html/template package

Templates functions:

### contentof

will replace the content of given file name 