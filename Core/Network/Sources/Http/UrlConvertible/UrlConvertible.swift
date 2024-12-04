import struct Foundation.URL

public protocol UrlConvertible: Sendable {
    func asURL() throws(HttpError.InvalidUrl) -> URL
}
