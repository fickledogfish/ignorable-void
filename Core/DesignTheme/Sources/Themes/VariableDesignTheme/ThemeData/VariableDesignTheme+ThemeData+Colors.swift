import struct SwiftUI.Color

extension VariableDesignTheme.ThemeData {
    public struct Colors {
        // MARK: - Public members

        public let disabled: Color
        public let error: Color
        public let normal: Color

        // MARK: - Initializers

        public init(
            disabled: Color,
            error: Color,
            normal: Color
        ) {
            self.disabled = disabled
            self.error = error
            self.normal = normal
        }
    }
}

// MARK: - Sendable

extension VariableDesignTheme.ThemeData.Colors: Sendable {
}
