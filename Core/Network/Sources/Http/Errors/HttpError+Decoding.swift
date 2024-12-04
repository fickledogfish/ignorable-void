import protocol Foundation.LocalizedError

extension HttpError {
    public enum Decoding {
        case nilBody

        case unknown(Error)
    }
}

// MARK: - Error

extension HttpError.Decoding: Error {
}
