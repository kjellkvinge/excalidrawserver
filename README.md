# excalidrawserver

Proof of concept.

Bundle excalidraw in a simple binary

All you have to do is to start the binary

```
./excalidrawserver-linux
# and go to localhost or an IP you host is listening to
```

## Hacking on excalidraw
```
# create patch-set
git format-patch master --stdout > ../mypatch.patch

# apply patches
git am ../mypatch.patch
```
see `.drone.ytml`
