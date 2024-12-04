public struct HttpCertificate {
    // MARK: - Public members

    public let host: String
    public let certificate: String

    // MARK: - Initializers

    public init(
        host: String,
        certificate: String
    ) {
        self.host = host
        self.certificate = certificate
    }
}
