import struct Foundation.Data

public protocol RequestBodyEncoder {
    func encode<Body: Encodable>(
        requestBody: Body
    ) throws(HttpError.Encoding) -> Data
}
