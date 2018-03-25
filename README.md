# Welcome to MyVol!

This is a WIP project that is mainly aimed at learning to build Web applications using [http://gobuffalo.io](Buffalo).

[![Build Status](https://travis-ci.org/piffio/myvol.svg?branch=master)](https://travis-ci.org/piffio/myvol)
[![Maintainability](https://api.codeclimate.com/v1/badges/9baeaee7897432fa5303/maintainability)](https://codeclimate.com/github/piffio/myvol/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/9baeaee7897432fa5303/test_coverage)](https://codeclimate.com/github/piffio/myvol/test_coverage)
[![Go Report Card](https://goreportcard.com/badge/github.com/piffio/myvol)](https://goreportcard.com/report/github.com/piffio/myvol)

## Database Setup

In order to run a local instance, you will need a PostgreSQL db up and running.
The simplest way to achieve this is via one of the offcial Docker images:

    $ docker run --rm -it --publish 0.0.0.0:5432:5432 --name pg -e POSTGRES_PASSWORD=postgres postgres:9-alpine

Caveat: the container will be destroyed when you stop it.

### Initialise the local enviroment

Once the DB is up and running, run the following command that will take care of setting up the whole enviroment

	$ buffalo setup

You will need to run this command only once

## Starting the Application

To run the application, use the standard buffalo way:

	$ buffalo dev

The application will be available at [http://127.0.0.1:3000](http://127.0.0.1:3000).

