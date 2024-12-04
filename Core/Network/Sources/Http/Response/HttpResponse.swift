public struct HttpResponse<Body: HttpResponseBody> {
    // MARK: - Public members

    public let status: Int

    // MARK: - Internal members

    let body: Body?
}
