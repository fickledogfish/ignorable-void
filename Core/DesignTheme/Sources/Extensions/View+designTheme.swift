import protocol SwiftUI.View

extension View {
    public func designTheme(_ theme: DesignTheme) -> some View {
        environmentObject(theme)
    }
}
