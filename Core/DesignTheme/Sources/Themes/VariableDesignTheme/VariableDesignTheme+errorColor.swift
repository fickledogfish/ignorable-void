import struct SwiftUI.Color

extension VariableDesignTheme {
    func errorColor(_ token: Token.ColorFamily.Color) -> Color {
        switch token {
        case .background:
            themeData.colors.error.background

        case .border:
            themeData.colors.error.border

        case .placeholderText:
            themeData.colors.error.placeholderText

        case .text:
            themeData.colors.error.text
        }
    }
}
