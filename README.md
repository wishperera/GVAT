# GVAT
German vat id validator

[![lint and test](https://github.com/wishperera/GVAT/actions/workflows/github-actions.yml/badge.svg?branch=master)](https://github.com/wishperera/GVAT/actions/workflows/github-actions.yml)
### Introduction

GVAT is a REST service for validating the German VAT ID numbers against the EU/VIES online database.

### How to build and run

#### Using docker

- make sure the latest version of docker is installed
- clone the repository using git 

```shell
git clone https://github.com/wishperera/GVAT.git
```
- enter the root directory and run the following command. Edit the `.env` file if you need to change the configurations
  

```shell
cd GVAT
sh cmd/docker_build_and_run.sh
```

- verify the container is running 

```shell
docker ps
```
- if everyting goes smooth  you should see something like below

![include](docs/img/docker-container.png)


#### Using go build

- make sure go 1.17.1 or above is installed
- clone the repository using git

```shell
git clone https://github.com/wishperera/GVAT.git
```
- enter the root directory and run the following command. Edit the `.env` file if you need to change the configurations
- edit the ``
- 
```shell
cd GVAT
sh cmd/build_and_run.sh
```
