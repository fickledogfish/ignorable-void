public enum HttpError {
    case decoding(Decoding)
    case encoding(Encoding)
    case invalidUrl(InvalidUrl)

    case unknown(Error)
}

// MARK: - Error

extension HttpError: Error {
}
