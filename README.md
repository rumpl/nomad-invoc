# nomad-invoc

This is a test to run a [docker application](https://github.com/docker/app) orchestrated by nomad.

# Prerequisites

You need to have [consul](https://www.consul.io/) and [nomad](https://www.nomadproject.io/) installed and running:
```bash
$ consul agent -dev
$ nomad agent -dev
```

Install [docker app](https://github.com/docker/app) (comes bundled with Docker Desktop)

You must change the base invocation image for docker app, in `~/.docker/config.json` add this:
```json
"plugins": {
    "app": {
        "base-invocation-image": "nomad-invoc"
    }
}
```

You can then make the image with make: `make invocation-image`.

With all that setup, you can now use `docker app` to install an image:

```bash
$ docker app build examples/hello-world -t my-app
$ docker app run my-app --name my-app
```

The example here comes from the [docker app repository](https://github.com/docker/app/tree/master/examples/hello-world).