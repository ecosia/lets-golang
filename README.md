# Let's Golang!

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

## What will we do?

This workshop will try to illustrate two of the strengths of Go: strong web libraries and parallelism. For this we will built a simplified version of a service that we are running at Ecosia in production. The web service service that we will build is supposed to call several APIs in parallel while enforcing a maximum timeout for all calls and will return a JSON response made up from the APIs' responses.

## Let's Golang now!

Start with [step 0](./0---Getting-started).
