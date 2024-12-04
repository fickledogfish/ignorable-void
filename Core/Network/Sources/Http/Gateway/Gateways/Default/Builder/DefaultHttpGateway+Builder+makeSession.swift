import class Foundation.URLSessionConfiguration
import struct Foundation.TimeInterval

import class Alamofire.Session

extension DefaultHttpGateway.Builder {
    func makeSession() -> Session {
        let configuration = URLSessionConfiguration.ephemeral

        configuration.timeoutIntervalForRequest = timeoutIntervalForRequest
        configuration.timeoutIntervalForResource = timeoutIntervalForResource

        configuration.urlCache = nil
        configuration.requestCachePolicy = .reloadIgnoringLocalCacheData

        configuration.httpCookieStorage = nil
        configuration.httpCookieAcceptPolicy = .never

        return Session(
            configuration: configuration,
            delegate: SessionDelegate(),
            serverTrustManager: WildcardServerTrustManager(
                evaluators: makeEvaluators()
            )
        )
    }
}
