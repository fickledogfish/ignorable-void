import class Foundation.URLResponse
import struct Foundation.Data
import struct Foundation.URLRequest

public protocol HttpMiddleware: Sendable {
    func before(request: URLRequest) async throws(HttpError) -> URLRequest

    func after(
        response: URLResponse,
        for request: URLRequest,
        with data: Data?
    ) async throws(HttpError) -> URLResponse
}

// MARK: - Default implementations

extension HttpMiddleware {
    public func before(
        request: URLRequest
    ) async throws(HttpError) -> URLRequest {
        request
    }

    public func after(
        response: URLResponse,
        for request: URLRequest,
        with data: Data?
    ) async throws(HttpError) -> URLResponse {
        response
    }
}
