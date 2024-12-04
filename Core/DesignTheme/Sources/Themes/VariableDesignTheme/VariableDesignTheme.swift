import struct CoreFoundation.CGFloat

import struct SwiftUI.Color

public final class VariableDesignTheme: DesignTheme {
    // MARK: - Internal members

    var themeData: ThemeData {
        switch colorScheme {
        case .dark:
            darkThemeData

        case .light:
            lightThemeData

        @unknown
        default:
            unknownThemeData
        }
    }

    // MARK: - Private members

    private let darkThemeData: ThemeData
    private let lightThemeData: ThemeData
    private let unknownThemeData: ThemeData

    // MARK: - Initializers

    public init(
        darkThemeData: ThemeData = .dark,
        lightThemeData: ThemeData = .light,
        unknownThemeData: ThemeData? = nil
    ) {
        self.darkThemeData = darkThemeData
        self.lightThemeData = lightThemeData
        self.unknownThemeData = unknownThemeData ?? darkThemeData

        super.init()
    }

    // MARK: - Public functions

    override public func cgFloat(spacing token: Token.Spacing) -> CGFloat {
        switch token {
        case .extraSmall:
            themeData.spacing.extraSmall

        case .small:
            themeData.spacing.extraSmall

        case .medium:
            themeData.spacing.extraSmall

        case .large:
            themeData.spacing.extraSmall
        }
    }

    override public func color(_ token: DesignTheme.Token.ColorFamily) -> Color {
        switch token {
        case .disabled(let token):
            disabledColor(token)

        case .error(let token):
            errorColor(token)

        case .normal(let token):
            normalColor(token)
        }
    }
}
