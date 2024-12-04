import struct Foundation.URL

extension UrlConvertible {
    func throwingHttpError() throws(HttpError) -> URL {
        do {
            return try asURL()
        } catch {
            throw HttpError.invalidUrl(error)
        }
    }
}
