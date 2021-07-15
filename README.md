# filter-function-configs

Build the function:

```shell
docker build -t gcr.io/foo/filter-local-configs:unstable .
```

Test locally:

```shell
kpt fn source | kpt fn eval - -i gcr.io/foo/filter-local-configs:unstable --image-pull-policy=never | kpt fn sink DEST_DIR
```
