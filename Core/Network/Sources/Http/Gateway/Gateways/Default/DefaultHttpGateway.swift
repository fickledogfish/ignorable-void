import class Alamofire.Session

public actor DefaultHttpGateway {
    // MARK: - Internal members

    let middlewares: [HttpMiddleware]

    let session: Session

    // MARK: - Initializers

    init(
        middlewares: [HttpMiddleware],
        session: Session
    ) {
        self.middlewares = middlewares
        self.session = session
    }
}
