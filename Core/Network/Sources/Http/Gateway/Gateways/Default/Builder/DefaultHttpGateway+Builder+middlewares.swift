extension DefaultHttpGateway.Builder {
    public mutating func middlewares(_ middlewares: [HttpMiddleware]) -> Self {
        self.middlewares.append(contentsOf: middlewares)

        return self
    }

    @inlinable
    public mutating func middlewares(_ middlewares: HttpMiddleware...) -> Self {
        self.middlewares(middlewares)
    }
}
