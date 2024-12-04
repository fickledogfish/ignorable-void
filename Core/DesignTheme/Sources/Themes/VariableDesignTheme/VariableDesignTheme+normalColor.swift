import struct SwiftUI.Color

extension VariableDesignTheme {
    func normalColor(_ token: Token.ColorFamily.Color) -> Color {
        switch token {
        case .background:
            themeData.colors.normal.background

        case .border:
            themeData.colors.normal.border

        case .placeholderText:
            themeData.colors.normal.placeholderText

        case .text:
            themeData.colors.normal.text
        }
    }
}
