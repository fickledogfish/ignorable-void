struct NilResponseBody {
}

// MARK: - Decodable

extension NilResponseBody: Decodable {
}

// MARK: - OptionalResponse

extension NilResponseBody: OptionalResponse {
    static func emptyValue() -> Self {
        NilResponseBody()
    }
}
