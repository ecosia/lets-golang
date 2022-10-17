# Let's Golang!

⚠️ ***Please note that the content may not be representative of Ecosia's current engineering standards.***

## Parallel programming for python developers

👋 Hi there!

We are [Jéssica](https://github.com/jessicalins) and [Dominik](https://github.com/DoHe) from [Ecosia](https://www.ecosia.org/) 🌳

In our life before Ecosia we were both mostly working with Python 🐍 but at Ecosia the main backend language is Go (sometimes called Golang) ![golang mascot](./gopher.png). In this workshop we want to show you what Go is about and how things work on the other side of the language divide from the perspective of two long-time Python developers.

## Why should you even care about Go?

Well, when you're reading/listening to this, you probably already have some initial interest, but here are some points that we really like about Go and why we are using it at Ecosia:

* It's fast 🐎 and can do system-level programming 🔧 but is still much more accessible 🚪 than C
* It has web 🌐 and parallelism ⏸️ built right into the language
* It's a typed ⌨️ language with an interesting concept of interfaces 👥
* It has an extensive standard library 📚, including things like a html template rendering engine 📄
* It's heavily inspired by Python 🐍

The title promised parallel programming, so why do we think Go is a good choice for that, especially when coming from Python?

**Python** sadly has some shortcomings when trying do things in parallel. The most commonly used Python implementation, CPython, has a [Global Interpreter Lock](https://en.wikipedia.org/wiki/Global_interpreter_lock) 🔒 which effectively limits parallelism – although less so when waiting for external resources like a filesystem or a HTTP call as these happen mostly outside the GIL. This can be avoided by using `multiprocessing` (using system processes) instead of `threading` (using OS threads), but at the cost of slower startup time 🐢 and higher memory consumption 💾. Additionally, the built-in functionality for synchronizing between threads (using `queue.Queue` and `threading.Lock`) can be rather intimidating and hard to understand and read ⚗️. Another issue is that a lot of objects in Python aren't meant to be shared between threads or processes, so "manual" serialization and access locking is necessary 🗝️.

**Go** on the other hand has parallel execution as one of its core language features. That manifests in three ways. Firstly, Go uses a very efficient implementation of threads called `goroutines` that can run several light-weight 🦢 Go threads inside a single OS thread and manages the threads for you, neatly combining concurrency and parallelism. Secondly, there is a core language feature to start execution in a goroutine (the `go` keyword) 🎢 and a very accessible communication model to share memory between goroutines without caring about race conditions (channels) ⛵. And lastly, most of the Go's standard library is thread-safe 🧵, so there is no need for extra locks when sharing objects among goroutines.

## What will we do?

This workshop will try to illustrate two of the strengths of Go: strong web libraries and parallelism. For this we will built a simplified version of a service that we are running at Ecosia in production. The web service that we will build is supposed to call several APIs in parallel while enforcing a maximum timeout for all calls and will return a JSON response made up from the APIs' responses.

If you are attending a live session: there will be two people running the workshop, please ask the one _not_ presenting and they will walk over to you. If they feel your question is relevant for everyone they will share it with everyone.

## Let's Golang now!

Start with [step 0](./0&#32;&#45;&#32;Getting&#32;started).
