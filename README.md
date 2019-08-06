```sh
dev_appserver.py app
```

```sh
tar -C /usr/local -xzf go1.9.2.linux-amd64.tar.gz
echo "export PATH=\$PATH:/usr/local/go/bin" >> /etc/profile.d/go.sh
export PATH=$PATH:$HOME/go/bin
```

```sh
export PATH=$PATH:/usr/lib64/google-cloud-sdk/bin:/home/tahomatx/download/go_appengine


sudo dnf -y install google-cloud-sdk-app-engine-go
```

```sh
go get -u github.com/golang/dep/cmd/dep
```

```sh
export GOPATH=$HOME/go

goapp deploy -application gcp-project -version 4 app
```
