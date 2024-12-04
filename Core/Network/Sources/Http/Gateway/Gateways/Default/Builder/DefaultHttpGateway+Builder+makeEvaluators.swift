import class CoreFoundation.CFData

import struct Foundation.Data
import struct Foundation.URL

import class Security.SecCertificate
import func Security.SecCertificateCreateWithData

import class Alamofire.DisabledTrustEvaluator
import class Alamofire.PinnedCertificatesTrustEvaluator
import protocol Alamofire.ServerTrustEvaluating

extension DefaultHttpGateway.Builder {
    func makeEvaluators() -> [String: ServerTrustEvaluating] {
        makeTrustedEvaluators()
            .merging(
                makeDisabledEvaluators(),
                uniquingKeysWith: { $1 }
            )
    }

    // MARK: - Private functions

    private func makeTrustedEvaluators() -> [String: ServerTrustEvaluating] {
        let groupedCertificates = Dictionary(
            grouping: pinnedCertificates,
            by: \.host
        )

        return groupedCertificates.reduce(
            into: [String: ServerTrustEvaluating]()
        ) { acc, next in
            let (host, certificates) = next
            let wildcards = URL
                .wildcards(for: host)
                .flatMap {
                    groupedCertificates[$0] ?? []
                }

            acc[host] = PinnedCertificatesTrustEvaluator(
                certificates: (certificates + wildcards)
                    .map(\.certificate)
                    .compactMap(makeSecCertificate(for:)),
                acceptSelfSignedCertificates: acceptsSelfSignedCertificates
            )
        }
    }

    private func makeDisabledEvaluators() -> [String: ServerTrustEvaluating] {
        hostsWithDisabledTrustEvaluation.reduce(into: [:]) { acc, next in
            acc[next] = DisabledTrustEvaluator()
        }
    }

    private func makeSecCertificate(
        for certificate: String
    ) -> SecCertificate? {
        guard let certificateData = Data(base64Encoded: certificate) else {
            return nil
        }

        return SecCertificateCreateWithData(
            nil,
            certificateData as CFData
        )
    }
}
