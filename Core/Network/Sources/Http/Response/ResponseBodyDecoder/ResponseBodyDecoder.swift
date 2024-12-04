import struct Foundation.Data

public protocol ResponseBodyDecoder: Sendable {
    func decode<Body: Decodable>(
        _ data: Data,
        as type: Body.Type
    ) throws(HttpError.Decoding) -> Body
}
