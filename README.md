# boilr-registry

This is a [boilr template](https://github.com/tmrts/boilr) for creating a registry credentials extension. Use a registry credential extension to provide authentication credentials for pulling images from container registries. Get started by cloning the project and installing the template:

```console
$ cd boilr-registry
$ boilr template save . drone-registry -f
```

create a project in directory my-registry:

```console
$ boilr template use drone-registry my-registry
```

enter the docker registry name for this project:

```text
[?] Please choose a value for "DockerRepository" [default: "foo/bar"]:
```

enter the go module name:

```text
[?] Please choose a value for "GoModule" [default: "github.com/foo/bar":
```
