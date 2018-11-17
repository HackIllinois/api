# Introduction

Our API is organized as a collection of [microseverices](https://en.wikipedia.org/wiki/Microservices), behind [Arbor](https://github.com/arbor-dev/arbor), a statically configured API framework, which acts as an API Gateway.

Each microservice is responsible for doing only one set of tasks. For example, our [user service](/reference/services/user-service) is only responsible for storing and processing data for the various kinds of users our API deals with - mentors, participants, staff, etc.

For authorization we use [JSON Web Tokens (JWTs)](https://jwt.io) that encode a user ID and some more information, in a system similar to [Bearer (or token-based) authentication](https://swagger.io/docs/specification/authentication/bearer-authentication/).

When a client makes an HTTP request to `api.hackillinois.org`, it is taken through several middleware. One of them is the *Authentication* middleware, which ensure the user is authenticated. Another one is the *Identification* middleware, and puts the user ID of the requesting user in the HackIllinois-Identity header, which can be used by the individual services.

The authorized request is then forwarded to the relevant micro-service based on routes configured in the [gateway](/reference/gateway), where controllers present in each micro-service process the request, call various service funcitons, perform the action requested, and return the response, which is passed back to the user, via Arbor.

Our persistence layer consists of a [MongoDB](https://mongodb.com) database, which has collections storing data relevant to each service.