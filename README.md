# Introduction

This is a starter project for Go API server applications. It has full dockerized and debugging capabilities using VSCode dev-containers. 

Any changes in the code is automatically recompiled and debugger is restarted

# Developer Enviroment Setup
Following steps are to be performed once on the developer's machine

- Install docker
- Download and install VSCode
- Install Dev-container extensions

# Setup
Following steps are required to be done to run the dev-container

- Create `.env` file by copying the sample file using this command `cp .env.sample .env`
- Edit `.env` file. `TARGET` can be either `prod` or `debug`

- `docker-compose build --build-args target=debug`