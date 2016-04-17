# The Shorty Project

## Introduction

I have came across one of my work colleague [urlshortener](https://github.com/jarisoft/urlshortener), then I decided to make my own version using [Google Go](https://golang.org/) and [MongoDB](https://www.mongodb.org/).

I thougth I use this as an opportunity to prove to myself, that I can build a Symfony style web application without using a full-stack web framework.

It's also a good opportunity to showcase my libraries I have built in [cjtoolkit](https://github.com/cjtoolkit), see it all work in action.
*  [The working prototype version of form 3.0](https://github.com/cjtoolkit/form/tree/epic/form_three)
    * Automated form validation, it's doesn't do rendering, like it did in the previous version, it's was a nightmare to maintain.  I think it's better that the end  user build the template themselves.
*  [cli](https://github.com/cjtoolkit/cli)
    * Command Line Interface Builder inspired by [Symfonys' Console Component](http://symfony.com/doc/current/components/console/introduction.html), it's uses pointers and transformers.
*  [context](https://github.com/cjtoolkit/context)
    * An user context holder, it's uses gos' net/http ResponseWriter interface as a trojan horse, just to avoid using a centralised hash table and mutual exclusion (which may cause scalability issues).
*  [groot](https://github.com/cjtoolkit/groot)
    * An Object Oriented Style router built on top of the speeedy httprouter (https://github.com/julienschmidt/httprouter)

## Why use Google Go?

Because I like it, it's very simplistic, the fact that it's does not provide elegant on silver plate like for example Haskell and Lisp does,  so therefore I'm forced to be inventive with my code and I really like that.  If you want elegant on a silver plate, I'm not going to judge.  PHP is kinda like Go in that sense.

The other thing I like about Google Go, it call do asynchronous programing without the ridiculous callbacks like it's does with JavaScript (and node.JS), instead I use channels to to keep asynchronous programming under control and I can even catch the errors with channels, just think of channels as a high level version of mutual exclusion.

## How long did it's take to build this project?

It's took me about a week, well I'm employed and I took my sweet little time testing the project, I even built my own mock manually. (It's not as bad as writing dependency injection in yml, or worse xml)

## Why did you build the mocks manually? Rather than have than it done automatically as it is in phpspec?

I built my own mock manually because I find it nice to have full control over the flow, I even manage to test the html output elegantly and yes I have made use of asynchronous programing and channels, I only have to worry about what method is actually going to get called, as for the method that get called while it was not suppose to, will cause a deadlock (because channels block), Go is very good at detecting that and will throw a stack trace and that how I find out.  It's that simple.

## Any other usage of asynchronous programming?

I have used it to convert middleware into a easy to manage dependency, for dependency injection.  I'm not a fan of middleware.

## What approach did you take with dependency injection?

It's quite simple, yet clever.  It's just series of hierarchical function calls, each of those function building a data struct (or constructing an object in PHP terms) and than returning the struct; It's does the entire thing without reflection (which is slow) and caching (which comes at the cost of complexity), plus because it's pure Go you get the added benefit of compilation time type checking, as Go is strongly and statically type, which is a good thing, a compilation time error is more likely to happen in the development environment, but less likely to happen to production environment.

As for reflection, no matter what language you use (including Go), it is dynamically type, it's does the type checking at run time, but it's does have it's benefits, for example html template, serialization and testing.

As for business logic and dependency injection, using reflection is a very bad idea, anything can happen with dynamic typing, it's better to have it locked down static and strong typing.

Also because Go is compiled, you don't need to use yml or xml, otherwise you'll just end up adding complexity, so my advice, don't do it.  I have used json but only for configuration,  because obviously it's usualy a bad idea to hard code configuration.

## Demo

Can be found at https://shorty.cj-jackson.com

## Installation

*  First you need to prepare the configuration. Copy and paste (and make adjustment) the below to either `/home/username/.shorty/parameters.json` (Unix) or `C:\Users\Username\.shorty\parameters.json` (Windows)

```json
{
	"MongoDial": "127.0.0.1",
	"MongoDbName": "shorty",
	"FilePath": "/path/to/static/file",
	"CsrfKey": "47b8ef6ea7e92486d14dbecf6a921700",
	"Domain": "https://www.example.com"
}
```

*  Assuming you got golang tooling setup, type the below into your terminal

```sh
$ go get github.com/CJ-Jackson/shorty
```

*  Now to start the engine (it's runs on port 8080 by default, you can change it with `--address=:8080` option.). Assuming you got MongoDB up and running (it will create the database and collection on the fly).

```sh
$ shorty http:start
```