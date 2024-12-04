import class Foundation.JSONDecoder
import struct Foundation.Data

extension JSONDecoder: ResponseBodyDecoder {
    public func decode<Body: Decodable>(
        _ data: Data,
        as type: Body.Type
    ) throws(HttpError.Decoding) -> Body {
        do {
            return try self.decode(type, from: data)
        } catch {
            throw .unknown(error)
        }
    }
}
