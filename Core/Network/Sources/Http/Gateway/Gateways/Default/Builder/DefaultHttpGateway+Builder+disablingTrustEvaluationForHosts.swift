extension DefaultHttpGateway.Builder {
    public mutating func disablingTrustEvaluation(for hosts: [String]) -> Self {
        self.hostsWithDisabledTrustEvaluation.append(contentsOf: hosts)

        return self
    }

    public mutating func disablingTrustEvaluation(for hosts: String...) -> Self {
        disablingTrustEvaluation(for: hosts)
    }
}
