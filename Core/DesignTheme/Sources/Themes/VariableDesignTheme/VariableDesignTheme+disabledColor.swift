import struct SwiftUI.Color

extension VariableDesignTheme {
    func disabledColor(_ token: Token.ColorFamily.Color) -> Color {
        switch token {
        case .background:
            themeData.colors.disabled.background

        case .border:
            themeData.colors.disabled.border

        case .placeholderText:
            themeData.colors.disabled.placeholderText

        case .text:
            themeData.colors.disabled.text
        }
    }
}
