import struct Foundation.URL

import class Alamofire.ServerTrustManager
import enum Alamofire.AFError
import protocol Alamofire.ServerTrustEvaluating

final class WildcardServerTrustManager: ServerTrustManager {
    // MARK: - Public functions

    override func serverTrustEvaluator(
        forHost host: String
    ) throws -> ServerTrustEvaluating? {
        if let evaluator = evaluators[host] {
            return evaluator
        }

        let wildcards = URL.wildcards(for: host)
        guard
            !wildcards.isEmpty,
            let evaluator = wildcards.compactMap({ evaluators[$0] }).first
        else {
            if allHostsMustBeEvaluated {
                throw AFError.serverTrustEvaluationFailed(
                    reason: .noRequiredEvaluator(host: host)
                )
            }

            return nil
        }

        return evaluator
    }
}

// MARK: - Sendable

extension WildcardServerTrustManager: @unchecked Sendable {
}
