import typealias Foundation.TimeInterval

extension DefaultHttpGateway {
    public struct Builder {
        // MARK: - Internal members

        var middlewares: [HttpMiddleware] = []

        var timeoutIntervalForRequest: TimeInterval = 10
        var timeoutIntervalForResource: TimeInterval = 30

        var acceptsSelfSignedCertificates = false

        var hostsWithDisabledTrustEvaluation: [String] = []
        var pinnedCertificates: [HttpCertificate] = []

        // MARK: - Initializers

        init() {
        }
    }
}
