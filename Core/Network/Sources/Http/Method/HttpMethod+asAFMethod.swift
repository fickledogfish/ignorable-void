import struct Alamofire.HTTPMethod

extension HttpMethod {
    func asAFMethod() -> HTTPMethod {
        switch self {
        case .get:
            HTTPMethod.get

        case .post:
            HTTPMethod.post
        }
    }
}
