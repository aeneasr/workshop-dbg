# Workshop Deutsche Börse Group

[![Build Status](https://travis-ci.org/ory/workshop-dbg.svg?branch=master)](https://travis-ci.org/ory/workshop-dbg)
[![Coverage Status](https://coveralls.io/repos/github/ory/workshop-dbg/badge.svg?branch=master)](https://coveralls.io/github/ory/workshop-dbg?branch=master)

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Prerequisites](#prerequisites)
- [Setting up your development environment](#setting-up-your-development-environment)
  - [Git](#git)
    - [Installing Git on Windows](#installing-git-on-windows)
    - [Installing Git on OSX](#installing-git-on-osx)
  - [Google's Go Language](#googles-go-language)
    - [Installing Go on Windows](#installing-go-on-windows)
    - [Installing Go on OSX](#installing-go-on-osx)
  - [JetBrains IntelliJ IDEA](#jetbrains-intellij-idea)
    - [Installing JetBrains IntelliJ IDEA on Windows](#installing-jetbrains-intellij-idea-on-windows)
    - [Installing JetBrains IntelliJ IDEA on OSX](#installing-jetbrains-intellij-idea-on-osx)
- [Wiring it all together](#wiring-it-all-together)
  - [Windows Environment](#windows-environment)
  - [Mac OSX Environment](#mac-osx-environment)
  - [IntelliJ on all platforms](#intellij-on-all-platforms)
  - [Clone this Repository](#clone-this-repository)
- [You made it!](#you-made-it)
- [Heroku](#heroku)
- [Docker](#docker)
  - [Basics](#basics)
  - [Build image using Dockerfile and start container](#build-image-using-dockerfile-and-start-container)
  - [Show layer configuration](#show-layer-configuration)
  - [List running Containers](#list-running-containers)
  - [Kill Container](#kill-container)
  - [List Docker Images](#list-docker-images)
  - [Build container and make it public](#build-container-and-make-it-public)
  - [Use Docker Hub to download images from the cloud](#use-docker-hub-to-download-images-from-the-cloud)
  - [Run Wordpress using Docker Hub](#run-wordpress-using-docker-hub)
  - [Build project and run it on Google Container Engine](#build-project-and-run-it-on-google-container-engine)
  - [Build Docker image and push it to heroku (beta!)](#build-docker-image-and-push-it-to-heroku-beta)
- [References](#references)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Prerequisites

Every developer needs a set of tools to work with. These may vary depending on the eco system, but the following steps are
quite common amongst today's cloud developers. Before we jump into action, there are a few tools that need
to be downloaded and installed to and on your PC:

1. Download [Git](https://git-scm.com/downloads). Git works on Windows, Linux and Mac OSX.
2. Download [Google's Go Language](https://golang.org/dl/). Go works on Windows, Linux, Mac OSX, Android and other platforms.
3. Download [JetBrains IntelliJ IDEA Community Edition](https://www.jetbrains.com/idea/download/download-thanks.html?code=IIC).
   If the link provided does not work, try [this one](https://www.jetbrains.com/idea/). IntelliJ works on Windows, Linux and Mac OSX
   and is one of the most common Integrated Developer Environments (IDE) today.

While these tools are downloading, [set up an account on GitHub](https://github.com/join). You can use any email address you wish. If
you do not want to provide your personal email, you can use a "trash mailer", for example [byom.de](https://www.byom.de/).

## Setting up your development environment

### Git

Git is a Versioning Control System (VCS) developed by Linux founder Linus Torvalds and others. Git was designed because
the Linux Kernel Development got out of hand and a better VCS was needed.

#### Installing Git on Windows

Unfortunately, Microsoft *might* block the installation of Git. This is due to missing certificates, as most open source
projects do not have the resources to buy such an certificate from Microsoft. If you see an error like this one

![](docs/win-prevents-git-install.png)

there is no need to worry. Git has been around for years and is used by hundreds of thousands of developers every day
and is completely open source and peer reviewed. To override the faulty Windows SmartScreen, click on **More info**

![](docs/win-prevents-git-install-override.png)

and press **Run anyway**. Now, the installer should start.

![](docs/git-install-windows.png)

There is no need for customization while installing and you can simply use the default settings
by pressing **Next >**. When the installer is done, you should see this:

![](docs/git-install-windows-success.png)

Congratulations, you have now installed the tool that every developer has on your machine. You can now skip ahead to
[Installing Go on Windows](#installing-go-on-windows).

#### Installing Git on OSX

Once you have downloaded the git installer, open the file with right-click -> open and confirm
the following dialogue:

![](docs/mac-git-warning.png)

Everything else should work per default and you should end up with this screen:

![](docs/mac-git-success.png)

You can now skip ahead to [Installing Go on OSX](#installing-go-on-osx).

### Google's Go Language

Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. It is
developed by Google and was introduced because Google had significant problems maintaining their code base.

Over time, the large scale projects at Google got significantly more complex were increasingly hard to maintain and
scale in aspects of human and computational resources. Go is therefore primarily aimed at companies that process huge
amounts of data in a large, distributed and cloud native way.

Today, six years after release, Go is the backbone of almost every modern, scalable cloud application. Companies like Amazon,
Cloudflare, Spotify or IBM, just to name a few, are using Go in their production systems. Modern infrastructure systems
like Docker, Cloud Foundry or Kubernetes are written primarily in Go.

#### Installing Go on Windows

Installing Go on Windows is straight forward. You can leave all defaults as-is. Once Go is installed, you should
see a screen similar to this one:

![](docs/mac-go-success.png)

You can now skip ahead to [Installing JetBrains IntelliJ IDEA on Windows](#installing-jetbrains-intellij-idea-on-windows).

#### Installing Go on OSX

Installing Go on OSX is straight forward. You can leave all defaults as-is. Once Go is installed, you should
see a screen similar to this one:

![](docs/mac-go-success.png)

You can now skip ahead to [Installing JetBrains IntelliJ IDEA on OSX](#installing-jetbrains-intellij-idea-on-osx).

### JetBrains IntelliJ IDEA

JetBrains IntelliJ IDEA is an IDE aimed at Java developers. JetBrains has various IDEs for all sorts of programming
languages including Ruby, JavaScript, PHP and others. IntelliJ is my personal favorite and has superb support for Google
Go.

#### Installing JetBrains IntelliJ IDEA on Windows

The set up is straight forward. You can leave all defaults as-is. Once IntelliJ is installed, you should see a screen similar to this one:

![](docs/win-intellij-install.png)

You can now skip ahead to [Windows Environment](#windows-environment).

#### Installing JetBrains IntelliJ IDEA on OSX

The set up is straight forward. You can leave all defaults as-is. Once IntelliJ is installed, you should see a screen similar to this one:

![](docs/mac-intellij-install.png)

## Wiring it all together

The hardest part is wiring it all together because each environment (your PC) is unique in it's configuration. There
are a couple of things that need to be done now.

### Windows Environment

There are a few things we need to do on Windows to get things running. First, we need to set up your workspace.
To do so, create a `workspace` directory anywhere on your disk. I keep mine in my home directory at
`C:\Users\aeneas\workspace`. Because we are working with go, it is a good idea to create a subdirectory called `go` as
well `C:\Users\aeneas\workspace\go`.

Now we need to tell Go where this directory is by setting up an environment variable. Open the Windows start menu,
search for "environment variables" and click on "Edit the system environment variables":

![](docs/win-env-path-1.png)

![](docs/win-env-path-2.png)

Next, click on "Environment Variables" and then on "New" in the "System variables" section:

![](docs/win-env-path-3.png)

![](docs/win-env-path-4.png)

Now we will set the GOPATH by using `GOPATH` as "Variable name" and the path to your go workspace as "Variable value":

![](docs/win-env-path-5.png)

Congratulations! You just completed setting up Go! Next we will initialize IntelliJ and once that is done, we are ready
to run and modify some code!

### Mac OSX Environment

There are a few things we need to do on OSX to get things running. First, open
the terminal. To do so, open the terminal by using spotlight or launchpad search

![](docs/mac-open-terminal.png)

When using developer tools on OSX, Apple forces to use the so called XCode Tools
which require an Apple account and are 2 GB large. Because only git is needed, we are going
to to a little hack by typing:

```
echo "PATH=/usr/local/git/bin:\$PATH" >> ~/.bash_profile
source ~/.bash_profile
```

in the console. You can use **copy and paste** in the OSX terminal.
To verify that git is set up properly, type `git` in the console. You should see
something like:

```
usage: git [--version] [--help] [-C <path>] [-c name=value]
           [--exec-path[=<path>]] [--html-path] [--man-path] [--info-path]
           [-p | --paginate | --no-pager] [--no-replace-objects] [--bare]
           [--git-dir=<path>] [--work-tree=<path>] [--namespace=<name>]
           <command> [<args>]

...
```

Now let's set up your workspace.
To do so, create a `workspace` directory anywhere on your disk. I keep mine in my home directory at
`/Users/aeneas/workspace`. Because we are working with go, it is a good idea to create a subdirectory called `go` as
well `/Users/aeneas/workspace/go`. OSX has very good terminal support and we can create
a workspace in your home directory directly from commandline.

```
mkdir -p ~/workspace/go
```

Now we need to tell Go where your workspace is located. To do so, type

```
echo "GOPATH=~/workspace/go" >> ~/.bash_profile
source ~/.bash_profile
```

If you used a different directory for your workspace, you need to replace `~/workspace/go`
with your directory path.

### IntelliJ on all platforms

Next, we need to set up IntelliJ and Go. The following screens should guide you through the process

![](docs/win-intellij-1.png)

![](docs/win-intellij-2.png)

![](docs/win-intellij-3.png)

Now hit install and "Restart IntelliJ" once the set up has completed.

### Clone this Repository

We are going to make modification to this repository. To do so, we need to check out this repository. Do you remember
where you created your workspace at? We are to clone this repository into that directory. Open IntelliJ, click on *Checkout
from version control*

The following screens should be similar to what you see on your screen:

![](docs/intellij-home.png)

use this link as repository url `https://github.com/ory/workshop-dbg.git` and your workspace directory (e.g. `/Users/aeneas/workspace/go`)
as parent directory.

![](docs/intellij-clone.png)

You will now see some screens which you can confirm as-is until you get to this one:

![](docs/intellij-sdk.png)

IntelliJ now asks us to set up a software development kit (SDK). We are going to do so by clicking *configure* and
choosing the location where Go is installed. This should be detected automatically. If not, it will be located in
`C:\Go` on windows and in `/usr/local/go` on mac:

![](docs/intellij-home.png)

Once you have confirmed this dialogue as well and confirmed with *Finish*, you should end up with a screen similar to this one:

![](docs/intellij-done.png)

## You made it!

Congratulations! You mastered one of the trickiest parts modern developers face: setting up your development environment.
We are now going to look at some code and collaborative improve our application, deploy it to the cloud, review it and much more.

## Heroku

**Deploy latest changes to the cloud**
```
git push heroku <branch>:master
```

**Scale process up to a hundred instances**
```
heroku ps:scale web=100 --app dbg-contacts
```

**Scale process down**
```
heroku ps:scale web=1 --app dbg-contacts
```

**List releases**
```
heroku releases --app dbg-contacts
```

**Rollback**
```
heroku rollback --app dbg-contacts <id>
```

## Docker

### Basics

**Run `echo "hello world"`**
```
docker run ubuntu /bin/echo 'Hello world'
```

**Check all containers, including stopped ones**
```
docker ps -a
```

**Rerun and remove the container**
```
docker start -a <id>
docker rm <id>
```

**List available images**
```
docker images
```

**Remove the ubuntu image**
```
docker rmi -f ubuntu
```

**Rerun and the container**
```
docker run ubuntu /bin/echo 'Hello world'
```

**Remove all images**
```
docker rmi -f $(docker images -q)
```

**Rerun and the container**
```
docker run ubuntu /bin/echo 'Hello world'
```

**Run bash**
```
docker run -t -i ubuntu /bin/bash
```

**Daemonize and cleanup**
```
docker run -d ubuntu /bin/sh -c "while true; do echo hello world; sleep 1; done"
docker logs <name>
docker kill <name>
docker rm <name>
```

### Build image using Dockerfile and start container

```
cd $GOPATH/src/github.com/ory/workshop-dbg
docker build -t workshop-dbg .
docker run -d --publish 5678:5678 workshop-dbg
open http://$(docker-machine ip default):5678/memory/contacts
```

### Show layer configuration

```
docker history workshop-dbg
```

### List running Containers

```
docker ps
```

### Kill Container

```
docker rm -f <id>
```

### List Docker Images

```
docker images
```

### Build container and make it public

```
docker build -t oryam/workshop-dbg .
docker push oryam/workshop-dbg
```

### Use Docker Hub to download images from the cloud

```
docker run -d --publish 5679:5678 oryam/workshop-dbg-auto
open http://$(docker-machine ip default):5679/memory/contacts
```

### Run Wordpress using Docker Hub

```
docker run --name some-mysql -e MYSQL_ROOT_PASSWORD=my-secret-pw -d mysql:latest
docker run --name some-wordpress --link some-mysql:mysql -p 9090:80 -d wordpress
open http://$(docker-machine ip default):9090
```

### Build project and run it on Google Container Engine

**Build docker image and run it locally**

```
cd $GOPATH/src/github.com/ory/workshop-dbg
docker build -t gcr.io/dbg-workshop-1298/workshop-dbg .
docker run -d --publish 9090:5678 gcr.io/dbg-workshop-1298/workshop-dbg
open http://$(docker-machine ip default):9090/memory/contacts
```

**Push docker image to private repository**
```
gcloud docker push gcr.io/dbg-workshop-1298/workshop-dbg
```

**Run docker image on kubernetes**
```
kubectl cluster-info
kubectl run hello-workshop --image=gcr.io/dbg-workshop-1298/workshop-dbg --port=5678
kubectl get deployments
kubectl get pods
```

**Expose image to the web**
```
kubectl expose deployment hello-workshop --type="LoadBalancer"
kubectl get services hello-workshop
open http://<external_ip>:5678/memory/contacts
```

**Scaling**
```
kubectl scale deployment hello-workshop --replicas=4
kubectl get deployments
kubectl get pods
```

Learn more about [autoscaling](http://kubernetes.io/docs/user-guide/horizontal-pod-autoscaling/) and
[rolling updates](http://kubernetes.io/docs/user-guide/rolling-updates/).

**Cleanup**
```
kubectl delete service,deployment hello-workshop
```

### Build Docker image and push it to heroku (beta!)

**Run locally**
```
cd ~
git clone https://github.com/dbg-workshop/go-websocket-chat-demo
cd go-websocket-chat-demo/
docker-compose up -d web
open http://$(docker-machine ip default):8080
```

**Publish to the cloud**
```
heroku create
heroku docker:release
heroku apps:info
open <url>
```

## References

Microservices
https://en.wikipedia.org/wiki/Microservices

Representational state transfer (REST)
https://en.wikipedia.org/wiki/Representational_state_transfer

JavaScript Object Notation (JSON)
https://en.wikipedia.org/wiki/JSON

Github
https://github.com/about
