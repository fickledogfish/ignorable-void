import class Foundation.JSONEncoder
import struct Foundation.Data

extension JSONEncoder: RequestBodyEncoder {
    public func encode<Body: Encodable>(
        requestBody: Body
    ) throws(HttpError.Encoding) -> Data {
        do {
            return try encode(requestBody)
        } catch {
            throw .unknown(error)
        }
    }
}
