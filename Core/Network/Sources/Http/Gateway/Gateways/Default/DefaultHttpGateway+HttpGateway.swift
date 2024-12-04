import struct Foundation.Data
import struct Foundation.URL
import struct Foundation.URLRequest

import enum Alamofire.AFError

extension DefaultHttpGateway: HttpGateway {
    // swiftlint:disable:next function_parameter_count
    public func request<
        RequestBody: HttpRequestBody,
        ResponseBody: HttpResponseBody
    >(
        _ url: UrlConvertible,
        method: HttpMethod,
        headers: [HttpHeader],
        body: RequestBody?,
        withResponseBodyOf responseType: ResponseBody.Type,
        responseBodyDecoder: ResponseBodyDecoder
    ) async throws(HttpError) {
        let request = try await makeRequest(to: url)

        let response = await withCheckedContinuation { continuation in
            session.request(request).response {
                continuation.resume(returning: $0)
            }
        }

        switch response.result {
        case .success(let data):
            do {
                _ = try decode(
                    data: data,
                    as: responseType,
                    decoder: responseBodyDecoder
                )
            } catch {
                throw HttpError.decoding(.unknown(error))
            }

        case .failure(let error):
            throw mapAFError(error)
        }
    }

    private func decode<Response: HttpResponseBody>(
        data: Data?,
        as responseType: Response.Type,
        decoder: ResponseBodyDecoder
    ) throws(HttpError.Decoding) -> Response {
        guard let data else {
            return if
                let optionalResponseType = responseType as? OptionalResponse.Type,
                let optionalResponseValue = optionalResponseType.emptyValue() as? Response {
                optionalResponseValue
            } else {
                throw .nilBody
            }
        }

        return try decoder.decode(data, as: responseType)
    }

    private func mapAFError(_ error: AFError) -> HttpError {
        .unknown(error)
    }
}
