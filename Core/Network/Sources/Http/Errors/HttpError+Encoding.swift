extension HttpError {
    public enum Encoding {
        case unknown(Error)
    }
}

// MARK: - Error

extension HttpError.Encoding: Error {
}
