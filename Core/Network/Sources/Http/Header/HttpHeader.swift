public struct HttpHeader {
    // MARK: - Public members

    public let key: String
    public let value: String

    // MARK: - Initializers

    public init<Value: LosslessStringConvertible>(
        key: String,
        value: Value
    ) {
        self.key = key
        self.value = String(value)
    }
}

// MARK: - Sendable

extension HttpHeader: Sendable {
}
