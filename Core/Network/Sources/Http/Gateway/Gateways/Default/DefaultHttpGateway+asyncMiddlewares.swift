extension DefaultHttpGateway {
    var asyncMiddlewareStream: AsyncStream<HttpMiddleware> {
        AsyncStream { continuation in
            middlewares.forEach { continuation.yield($0) }

            continuation.finish()
        }
    }
}
