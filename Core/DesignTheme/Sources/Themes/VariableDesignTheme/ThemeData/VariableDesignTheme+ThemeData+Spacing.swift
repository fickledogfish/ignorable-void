import struct CoreFoundation.CGFloat

extension VariableDesignTheme.ThemeData {
    public struct Spacing {
        // MARK: - Public members

        public let extraSmall: CGFloat
        public let small: CGFloat
        public let medium: CGFloat
        public let large: CGFloat

        // MARK: - Initializers

        public init(
            extraSmall: CGFloat,
            small: CGFloat,
            medium: CGFloat,
            large: CGFloat
        ) {
            self.extraSmall = extraSmall
            self.small = small
            self.medium = medium
            self.large = large
        }
    }
}

// MARK: - Sendable

extension VariableDesignTheme.ThemeData.Spacing: Sendable {
}
