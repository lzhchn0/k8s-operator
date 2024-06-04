Setup Bazel Environment

```shell
export WATCH_NAMESPACE=default
```{{exec}}

> First, update your system's package list:
```shell
sudo apt update
```{{exec}}

> Install the necessary dependencies:
```shell
sudo apt install curl gnupg
```{{exec}}


```shell
echo "deb [arch=amd64] https://storage.googleapis.com/bazel-apt stable jdk1.8" | sudo tee /etc/apt/sources.list.d/bazel.list
```{{exec}}

> Import Bazel's GPG key:
```shell
curl https://bazel.build/bazel-release.pub.gpg | sudo apt-key add -
```{{exec}}

> Update your package list again:
```shell
sudo apt update
```{{exec}}

> Install Bazel:
```shell
sudo apt install bazel
```{{exec}}

> Verify that Bazel is installed correctly:
```shell
bazel version
```{{exec}}


```shell

```{{exec}}


```shell

```{{exec}}


```shell

```{{exec}}

