import struct Foundation.URL

extension URL {
    static func wildcards(for host: String) -> [String] {
        let components = host.split(separator: ".")
        guard components.count > 2 else {
            return []
        }

        return (2 ..< components.count)
            .reversed()
            .map {
                ["*"] + components.suffix(from: $0).joined(separator: ".")
            }
    }
}
