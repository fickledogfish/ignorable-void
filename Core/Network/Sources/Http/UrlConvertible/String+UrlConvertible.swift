import struct Foundation.URL

extension String: UrlConvertible {
    public func asURL() throws(HttpError.InvalidUrl) -> URL {
        do {
            return try URL(self, strategy: .url)
        } catch {
            throw .unknown(error)
        }
    }
}
