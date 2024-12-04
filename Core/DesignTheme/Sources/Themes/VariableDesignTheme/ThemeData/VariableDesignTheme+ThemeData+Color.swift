import struct SwiftUI.Color

extension VariableDesignTheme.ThemeData.Colors {
    public struct Color {
        // MARK: - Public members

        public let background: SwiftUI.Color
        public let text: SwiftUI.Color
        public let placeholderText: SwiftUI.Color
        public let border: SwiftUI.Color

        // MARK: - Initializers

        public init(
            background: SwiftUI.Color,
            border: SwiftUI.Color,
            placeholderText: SwiftUI.Color,
            text: SwiftUI.Color
        ) {
            self.background = background
            self.border = border
            self.placeholderText = placeholderText
            self.text = text
        }
    }
}

// MARK: - Sendable

extension VariableDesignTheme.ThemeData.Colors.Color: Sendable {
}
