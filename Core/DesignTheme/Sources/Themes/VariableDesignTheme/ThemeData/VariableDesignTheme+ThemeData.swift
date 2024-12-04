extension VariableDesignTheme {
    public struct ThemeData {
        // MARK: - Public members

        public let colors: Colors
        public let spacing: Spacing

        // MARK: - Initializers

        public init(
            colors: Colors,
            spacing: Spacing
        ) {
            self.colors = colors
            self.spacing = spacing
        }
    }
}

// MARK: - Sendable

extension VariableDesignTheme.ThemeData: Sendable {
}
