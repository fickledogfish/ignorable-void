extension DefaultHttpGateway.Builder {
    public func build() -> DefaultHttpGateway {
        DefaultHttpGateway(
            middlewares: middlewares,
            session: makeSession()
        )
    }
}
