
# Golang Docker sample project

## Docker registry

```
docker run -d --name registry \
	--restart always \
	-p 5000:5000 \
	registry
```

---

## Links

* https://docs.docker.com/develop/sdk/examples/
* https://stackoverflow.com/questions/45805563/pull-a-file-from-a-docker-image-in-golang-to-local-file-system
* https://stackoverflow.com/questions/45429276/how-to-run-docker-run-using-go-sdk-for-docker
* https://medium.com/backendarmy/controlling-the-docker-engine-in-go-d25fc0fe2c45

### Issues

* https://github.com/docker/distribution/blob/master/reference/normalize.go

### Docker Registry UI

* https://github.com/Joxit/docker-registry-ui
* https://hub.docker.com/r/joxit/docker-registry-ui
