import SwiftUI

extension EnvironmentValues {
    @Entry
    var splashViewOnFinish: () -> Void = {}
}

extension View {
    public func splashViewOnFinish(
        perform action: @escaping () -> Void
    ) -> some View {
        environment(\.splashViewOnFinish, action)
    }
}
