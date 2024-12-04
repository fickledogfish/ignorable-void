import struct CoreFoundation.CGFloat

import class Foundation.NSRecursiveLock

import enum SwiftUI.ColorScheme
import struct SwiftUI.Color

open class DesignTheme {
    // MARK: - Public members

    public private(set) var colorScheme: ColorScheme {
        get {
            colorSchemeLock.withLock {
                currentColorScheme
            }
        }
        set {
            colorSchemeLock.withLock {
                currentColorScheme = newValue
            }
        }
    }

    // MARK: - Private members

    private var colorSchemeLock = NSRecursiveLock()
    private var currentColorScheme: ColorScheme = .light

    // MARK: - Initializers

    public init() {
    }

    // MARK: - Public functions

    open func didChange(colorScheme: ColorScheme) {
        objectWillChange.send()

        self.colorScheme = colorScheme
    }

    open func cgFloat(spacing token: Token.Spacing) -> CGFloat {
        0
    }

    open func color(_ token: Token.ColorFamily) -> Color {
        .white
    }
}
