# Introduction

Our API is organized as a collection of [microseverices](https://en.wikipedia.org/wiki/Microservices), behind [Arbor](https://github.com/arbor-dev/arbor), a statically configured API framework, which acts as an API Gateway.

Each microservice is responsible for doing only one set of tasks. For example, our [user service](/reference/services/user-service) is only responsible for storing and processing data for the various kinds of users our API deals with - mentors, participants, staff, etc.

For authorization we use [JSON Web Tokens (JWTs)](https://jwt.io) that encode a user ID and some more information, in a system similar to [Bearer (or token-based) authentication](https://swagger.io/docs/specification/authentication/bearer-authentication/).

When a client makes an HTTP request to `api.hackillinois.org`, it is taken through several middleware. One of them is the *Authentication* middleware, which ensures the user is authenticated. Another one is the *Identification* middleware, and puts the user ID of the requesting user in the HackIllinois-Identity header, which can be used by the individual services. The *Error* middleware allows passing of errors to the client using standard HTTP mechanisms, such as status codes, and response bodies.

The authorized request is then forwarded to the relevant micro-service based on routes configured in the [gateway](/reference/gateway), where controllers present in each micro-service process the request, call various service funcitons, perform the action requested, and return the response, which is passed back to the user, via Arbor.

Our persistence layer consists of a [MongoDB](https://mongodb.com) database, which has collections storing data relevant to each service.

##  Errors

Setting the DEBUG_MODE to "true" in the config file allows raw error messages (if applicable) to be passed through to the client. Otherwise, the raw error is suppressed.

Errors are classified into the following types:

1. **DatabaseError** - When database operations, such as fetch / insert / update) doesn't work. These are usually returned when a document / record that was requested wasn't found, such as when an operation is performed on an inexistent user.

2. **MalformedRequestError** - When the request is invalid or missing some key information. Possible scenarios are, when field validation fails on a request body, or when an ID is missing for an endpoint that depends on it.

3. **AuthorizationError** - When an authentication / authorization attempt fails. Possible scenarios include when OAuth-related services fail, such as when an authorization code is incorrect, a token is invalid / has expired etc.

4. **AttributeMismatchError** - When an action is performed on a user who is missing some attribute, such as when a check-in (without override) is attempted for a user who doesn't have a registration or RSVP, modifying a decision on a candidate (hacker) whose decision has already been finalized by a senior staff member etc. 

5. **InternalError** - When there could be multiple possible causes of the error, this is what we use. Using DEBUG_MODE to get the raw error is highly recommended to expedite bug resolution.

6. **UnknownError** - When the cause of an error cannot be identified.
