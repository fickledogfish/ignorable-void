extension DefaultHttpGateway.Builder {
    public mutating func pinning(certificates: [HttpCertificate]) -> Self {
        pinnedCertificates.append(contentsOf: certificates)

        return self
    }

    @inlinable
    public mutating func pinning(certificates: HttpCertificate...) -> Self {
        pinning(certificates: certificates)
    }
}
