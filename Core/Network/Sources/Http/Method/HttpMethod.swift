public enum HttpMethod {
    case get
    case post
}

// MARK: - Sendable

extension HttpMethod: Sendable {
}
