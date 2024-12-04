extension HttpError {
    public enum InvalidUrl {
        case unknown(Error)
    }
}

// MARK: - Error

extension HttpError.InvalidUrl: Error {
}
