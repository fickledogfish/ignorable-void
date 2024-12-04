extension DefaultHttpGateway.Builder {
    public mutating func acceptsSelfSignedCertificates(_ value: Bool) -> Self {
        self.acceptsSelfSignedCertificates = value

        return self
    }
}
