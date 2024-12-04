import struct Foundation.URL
import struct Foundation.URLRequest

extension DefaultHttpGateway {
    func makeRequest(
        to url: UrlConvertible
    ) async throws(HttpError) -> URLRequest {
        let url = try url.throwingHttpError()
        let baseRequest = URLRequest(url: url)

        do {
            return try await asyncMiddlewareStream.reduce(
                baseRequest,
                Self.middlewareReducer
            )
        } catch let error as HttpError {
            throw error
        } catch {
            // HttpMiddleware guarantees this will be an HttpError, so if it
            // isn't just fatal on it.
            fatalError("Unreasonable error: \(error)")
        }
    }

    // MARK: - Private functions

    private static func middlewareReducer(
        acc: URLRequest,
        next: HttpMiddleware
    ) async throws -> URLRequest {
        try await next.before(request: acc)
    }
}
