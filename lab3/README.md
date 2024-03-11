
## Python
### 1 (Bookworm)
[Dockerfile](./python.Dockerfile)
```bash
bash prepare.bash 1 # if not already done for versions [1-4]
docker build -t py:1 -f ./python.Dockerfile .
docker run -it --rm -p 8080:8080 py:1
```

### 2 (Bookworm)
[Dockerfile](./python.Dockerfile)
```bash
bash prepare.bash 2 # if not already done for versions [1-4]
docker build -t py:2 -f ./python.Dockerfile .
docker run -it --rm -p 8080:8080 py:1
```

### 3 (Bookworm)
[Dockerfile](./3.python.Dockerfile)
```bash
bash prepare.bash 3 # if not already done for versions [1-4]
docker build -t py:3 -f ./3.python.Dockerfile .
docker run -it --rm -p 8080:8080 py:3
```

### 4 (Alpine)
[Dockerfile](./4.python.Dockerfile)
```bash
bash prepare.bash 4 # if not already done for versions [1-4]
docker build -t py:4 -f ./4.python.Dockerfile .
docker run -it --rm -p 8080:8080 py:4
```

### 5
[Dockerfile](./5.python.Dockerfile)
```bash
bash prepare.bash 5
docker build -t py:5.book-w -f ./python.Dockerfile .
docker build -t py:5.alpine -f ./4.python.Dockerfile .
```

## GO
### 1 (Alpine)
[Dockerfile](./1.go.Dockerfile)
```bash
docker build -t go:1 -f ./1.go.Dockerfile .
docker run --rm -p 8080:8080 go:1
```

### 2 (Alpine & Scratch)
[Dockerfile](./2.go.Dockerfile)
```bash
docker build -t go:2 -f ./2.go.Dockerfile .
docker run --rm -p 8080:8080 go:2
```

### 3 (Alpine & gcr.io/distroless/static-debian11)
[Dockerfile](./3.go.Dockerfile)
```bash
docker build -t go:3 -f ./3.go.Dockerfile .
docker run --rm -p 8080:8080 go:3
```
