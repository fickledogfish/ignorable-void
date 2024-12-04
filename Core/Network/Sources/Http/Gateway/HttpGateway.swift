public protocol HttpGateway {
    // swiftlint:disable:next function_parameter_count
    func request<
        RequestBody: HttpRequestBody,
        ResponseBody: HttpResponseBody
    >(
        _ url: UrlConvertible,
        method: HttpMethod,
        headers: [HttpHeader],
        body: RequestBody?,
        withResponseBodyOf responseType: ResponseBody.Type,
        responseBodyDecoder: ResponseBodyDecoder
    ) async throws(HttpError)
}
