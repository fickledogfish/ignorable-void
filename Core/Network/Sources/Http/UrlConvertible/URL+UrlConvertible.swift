import struct Foundation.URL

extension URL: UrlConvertible {
    public func asURL() throws(HttpError.InvalidUrl) -> URL {
        self
    }
}
