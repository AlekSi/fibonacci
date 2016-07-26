# Fibonacci

## Assignment

1. The project should provide a RESTful web service.
    * The web service accepts a number, n, as input and returns the first n Fibonacci numbers, starting from 0.
      I.e. given n  = 5, appropriate output would represent the sequence [0, 1, 1, 2, 3].
    * Given a negative number, it will respond with an appropriate error.
2. Include whatever instructions are necessary to build and deploy/run the project, where "deploy/run" means
   the web service is accepting requests and responding to them as appropriate.
3. Include some unit and/or functional tests
4. Use any language that you know well

While this project is admittedly trivial, approach it as representing a more complex problem that you'll
have to put into production and maintain for 5 years.
Providing a link to a github/bitbucket repo with the project would probably be the easiest way to submit.

## How to run

```
go get github.com/AlekSi/fibonacci
fibonacci
```

## Rationale

1. Choosing a good web framework is essential: a good one can significantly reduce an amount of boilerplate,
   a bad one will not be pleasant to use. While it's tempting not to use any and just build everything oneself
   on top of `net/http` (I did it several times), the result will be just another half-baked web framework
   to maintain. Some time ago we examined several frameworks and chose Echo for its simplicity,
   good middleware system, and sane centralized error handling. Then I influenced some design decisions
   during Echo v2 development, so it becomes even closer to my idea of the good web framework.
2. The same about testing library. While `testing` package is simple, it's too simplistic. `testify` provides
   helpful assertions and suites, and works on top of `testing`, not replaces it.
3. I'm yet to find an input validation library I would be 100% happy to use. For that reason,
   I check N manually for now. I would definitely switch to a third-party library as a project gets bigger, though.
4. I wasn't given information about maximum N. I assumed that `uint` is enough, so we can just precompute
   all 94 numbers that fit. If support for bigger numbers is required, implementation can be changed.
   Original one used on-demand memoization; that's why unit test uses goroutines to check thread-safety.
5. In a real project, I would add some functional/integration/system tests (after all,
   [most unit tests are waste](http://rbcs-us.com/documents/Why-Most-Unit-Testing-is-Waste.pdf)).
   I have a [special project](https://github.com/go-gophers/gophers) for writing them.
