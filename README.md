
### Build

```bash
docker build -t go-db-populate-script .
```

### Run

```bash
docker run --rm -v $(pwd):/output go-db-populate-script
```
